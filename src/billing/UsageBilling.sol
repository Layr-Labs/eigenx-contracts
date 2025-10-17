// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {BillingCore} from "./BillingCore.sol";
import {IBillingModule} from "./interfaces/IBillingModule.sol";
import {IUsageBilling} from "./interfaces/IUsageBilling.sol";

/**
 * @title UsageBilling
 * @notice Usage-based billing with period settlement
 * @dev Records usage events and settles by period
 */
abstract contract UsageBilling is IBillingModule, IUsageBilling {
    // ============================================================================
    // State
    // ============================================================================

    BillingCore public immutable billingCore;

    mapping(address => mapping(uint40 => PeriodUsage)) public periodUsage;
    mapping(address => bool) public usageReporters;

    // ============================================================================
    // Modifiers
    // ============================================================================

    modifier onlyReporter() {
        require(usageReporters[msg.sender], UnauthorizedReporter());
        _;
    }

    // ============================================================================
    // Constructor
    // ============================================================================

    constructor(address _billingCore) {
        billingCore = BillingCore(_billingCore);
    }

    // ============================================================================
    // Admin Functions
    // ============================================================================

    /**
     * @notice Set usage reporter authorization
     */
    function _setUsageReporter(address reporter, bool authorized) internal {
        usageReporters[reporter] = authorized;
        emit UsageReporterSet(reporter);
    }

    // ============================================================================
    // Usage Recording
    // ============================================================================

    /**
     * @notice Record usage for an account in the current period
     * @dev Accumulates with existing usage
     */
    function recordUsage(address account, uint96 amount) external onlyReporter {
        uint40 currentPeriod = billingCore.getCurrentPeriod();
        _recordUsageForPeriod(account, amount, currentPeriod);
    }

    /**
     * @notice Record usage for a specific period
     * @dev Can record for current or past periods until settled
     */
    function recordUsageForPeriod(address account, uint96 amount, uint40 period) external onlyReporter {
        require(period <= billingCore.getCurrentPeriod(), InvalidPeriod());
        _recordUsageForPeriod(account, amount, period);
    }

    /**
     * @notice Batch record usage for multiple accounts
     */
    function recordUsageBatch(address[] calldata accounts, uint96[] calldata amounts, uint40 period)
        external
        onlyReporter
    {
        require(period <= billingCore.getCurrentPeriod(), InvalidPeriod());
        require(accounts.length == amounts.length, InvalidAmount());

        for (uint256 i = 0; i < accounts.length; i++) {
            _recordUsageForPeriod(accounts[i], amounts[i], period);
        }
    }

    /**
     * @notice Internal usage recording
     */
    function _recordUsageForPeriod(address account, uint96 amount, uint40 period) internal {
        PeriodUsage storage usage = periodUsage[account][period];

        // Cannot modify after settlement
        require(!usage.settled, PeriodAlreadySettled());

        // Accumulate usage
        usage.amount += amount;

        emit UsageRecorded(account, amount, period);
    }

    // ============================================================================
    // Period Settlement
    // ============================================================================

    /**
     * @notice Settle usage charges for a specific period
     * @dev Called at period end, same as compute billing
     */
    function settlePeriod(address account, uint40 period) external onlyReporter {
        _settlePeriod(account, period);
    }

    /**
     * @notice Batch settle multiple accounts
     */
    function settlePeriodBatch(address[] calldata accounts, uint40 period) external onlyReporter {
        require(period < billingCore.getCurrentPeriod(), InvalidPeriod());

        for (uint256 i = 0; i < accounts.length; i++) {
            if (!periodUsage[accounts[i]][period].settled) {
                _settlePeriod(accounts[i], period);
            }
        }
    }

    /**
     * @notice Internal settlement logic
     */
    function _settlePeriod(address account, uint40 period) internal {
        require(period < billingCore.getCurrentPeriod(), InvalidPeriod());

        PeriodUsage storage usage = periodUsage[account][period];
        require(!usage.settled, PeriodAlreadySettled());

        uint96 amount = usage.amount;

        // Charge through BillingCore
        if (amount > 0) {
            billingCore.chargePeriod(account, amount, period);
        }

        // Mark as settled
        usage.settled = true;
        usage.settledAt = uint40(block.timestamp);

        emit PeriodSettled(account, period, amount);
    }

    // ============================================================================
    // View Functions
    // ============================================================================

    /**
     * @notice Get usage for a specific period
     */
    function getUsageForPeriod(address account, uint40 period)
        external
        view
        returns (uint96 amount, bool settled, uint40 settledAt)
    {
        PeriodUsage memory usage = periodUsage[account][period];
        return (usage.amount, usage.settled, usage.settledAt);
    }

    /**
     * @notice Get total unsettled usage across all periods
     */
    function getUnsettledUsage(address account) external view returns (uint96 total) {
        uint40 currentPeriod = billingCore.getCurrentPeriod();

        // Look back through recent periods for unsettled usage
        // In practice, you'd limit how far back to look
        for (uint40 period = currentPeriod; period >= 0; period--) {
            PeriodUsage memory usage = periodUsage[account][period];

            if (!usage.settled && usage.amount > 0) {
                total += usage.amount;
            }

            // Prevent underflow and limit lookback
            if (period == 0 || period < currentPeriod - 12) break; // Max 12 months lookback
        }
    }

    /**
     * @notice Check if a period is settled
     */
    function isPeriodSettled(address account, uint40 period) external view returns (bool) {
        return periodUsage[account][period].settled;
    }

    /**
     * @notice Get charges for a specific period
     * @param account The account to query
     * @param period The period to query
     * @return amount The usage charges for this period
     * @return settled Whether the charges have been settled to BillingCore
     */
    function getChargesForPeriod(address account, uint40 period) external view returns (uint96 amount, bool settled) {
        PeriodUsage memory usage = periodUsage[account][period];
        return (usage.amount, usage.settled);
    }

    /**
     * @notice Get the count of active resources for an account
     * @dev Virtual function to be implemented by inheriting contracts
     * @param account The account to check
     * @return The number of active resources (e.g. VMs, API keys, etc)
     */
    function getActiveResourceCount(address account) external view virtual returns (uint256);

    /**
     * @notice Get the timestamp of the last resource change for an account
     * @dev Virtual function to be implemented by inheriting contracts
     * @dev Should return the timestamp of the last time resources were added/removed (e.g., API key changes)
     * @dev NOT for usage recording - only for resource state changes
     * @param account The account to check
     * @return The timestamp of the last resource modification
     */
    function getLastActivityTimestamp(address account) external view virtual returns (uint40);
}
