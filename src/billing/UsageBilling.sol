// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IBillingCore} from "./interfaces/IBillingCore.sol";
import {IBillingModule} from "./interfaces/IBillingModule.sol";
import {IUsageBilling} from "./interfaces/IUsageBilling.sol";

/**
 * @title UsageBilling
 * @notice Billing for metered usage that accumulates token costs over a period
 */
abstract contract UsageBilling is IUsageBilling, IBillingModule {
    IBillingCore public immutable billing;

    // ============================================================================
    // State
    // ============================================================================

    // Period-specific usage tracking
    mapping(address account => mapping(uint40 period => PeriodData)) public periodData;

    address public usageReporter; // Authorized to report usage costs

    // ============================================================================
    // Constructor
    // ============================================================================

    constructor(address _billing) {
        billing = IBillingCore(_billing);
    }

    // ============================================================================
    // Modifiers
    // ============================================================================

    modifier onlyReporter() {
        require(msg.sender == usageReporter, NotAuthorized());
        _;
    }

    // ============================================================================
    // Usage Recording
    // ============================================================================

    /**
     * @notice Record usage for a specific period
     * @param account The account to record usage for
     * @param amount The total usage amount for this period
     * @param period The period this usage belongs to
     */
    function recordUsage(address account, uint96 amount, uint40 period) external onlyReporter {
        require(!periodData[account][period].settled, PeriodAlreadySettled());
        require(period <= billing.getCurrentPeriod(), CannotRecordFutureUsage());

        // Track usage for the specific period
        periodData[account][period].amount = amount;

        emit UsageRecorded(account, amount, period);
    }

    /**
     * @notice Record usage for multiple accounts in a specific period
     * @param accounts Array of account addresses
     * @param amounts Array of total usage amounts for this period
     * @param period The period this usage belongs to
     */
    function recordBatch(address[] calldata accounts, uint96[] calldata amounts, uint40 period) external onlyReporter {
        require(accounts.length == amounts.length, LengthMismatch());
        require(period <= billing.getCurrentPeriod(), CannotRecordFutureUsage());

        for (uint256 i = 0; i < accounts.length; i++) {
            if (!periodData[accounts[i]][period].settled) {
                periodData[accounts[i]][period].amount = amounts[i];
                emit UsageRecorded(accounts[i], amounts[i], period);
            }
        }
    }

    // ============================================================================
    // Period Finalization and Settlement
    // ============================================================================

    /**
     * @notice Settle a period for an account with final usage amount
     * @param account The account to settle
     * @param period The period to settle
     * @param finalAmount The final usage amount for this period
     */
    function settlePeriod(address account, uint40 period, uint96 finalAmount) external onlyReporter {
        require(period < billing.getCurrentPeriod(), CannotSettleCurrentOrFuturePeriods());
        require(!periodData[account][period].settled, PeriodAlreadySettled());

        // Set the final amount for this period (overwrites any previous recording)
        periodData[account][period].amount = finalAmount;

        // Automatically charge and mark as settled
        if (finalAmount > 0) {
            billing.chargePeriod(account, finalAmount, period);
        }

        periodData[account][period].settled = true;
        emit PeriodSettled(account, period, finalAmount);
    }

    /**
     * @notice Settle multiple accounts for a period
     * @param accounts Array of accounts
     * @param period The period to settle
     * @param amounts Final usage amounts for each account
     */
    function settlePeriodBatch(address[] calldata accounts, uint40 period, uint96[] calldata amounts)
        external
        onlyReporter
    {
        require(accounts.length == amounts.length, LengthMismatch());
        require(period < billing.getCurrentPeriod(), CannotSettleCurrentOrFuturePeriods());

        for (uint256 i = 0; i < accounts.length; i++) {
            if (!periodData[accounts[i]][period].settled) {
                periodData[accounts[i]][period].amount = amounts[i];

                if (amounts[i] > 0) {
                    billing.chargePeriod(accounts[i], amounts[i], period);
                }

                periodData[accounts[i]][period].settled = true;
                emit PeriodSettled(accounts[i], period, amounts[i]);
            }
        }
    }

    // ============================================================================
    // Admin Functions
    // ============================================================================

    /**
     * @notice Set the authorized usage reporter
     */
    function _setUsageReporter(address reporter) internal {
        usageReporter = reporter;
        emit UsageReporterSet(reporter);
    }

    // ============================================================================
    // View Functions
    // ============================================================================

    /**
     * @notice Check if a period is settled for an account
     */
    function isPeriodSettled(address account, uint40 period) external view returns (bool) {
        return periodData[account][period].settled;
    }

    /**
     * @notice Get charges for an account in a specific period (IBillingModule interface)
     * @param account The account to query
     * @param period The period to query charges for
     * @return amount Recorded usage amount for that period (whether or not settled)
     */
    function getChargesForPeriod(address account, uint40 period) external view override returns (uint96 amount) {
        return periodData[account][period].amount;
    }

    /**
     * @notice Get total outstanding charges not yet charged to BillingCore (IBillingModule interface)
     * @param account The account to query
     * @return amount Total charges across all unsettled periods
     */
    function getOutstandingCharges(address account) external view override returns (uint96 amount) {
        uint40 currentPeriod = billing.getCurrentPeriod();

        // Work backward from current period until we hit a settled period
        // This is more efficient as we only iterate through recent unsettled periods
        for (uint40 period = currentPeriod; period >= 0; period--) {
            PeriodData storage data = periodData[account][period];

            if (data.settled) {
                // Stop when we hit a settled period
                break;
            }

            if (data.amount > 0) {
                amount += data.amount;
            }

            // Prevent underflow on period = 0
            if (period == 0) break;
        }
    }
}
