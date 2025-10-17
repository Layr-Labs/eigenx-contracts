// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {BillingCore} from "./BillingCore.sol";
import {IBillingModule} from "./interfaces/IBillingModule.sol";
import {IComputeBilling} from "./interfaces/IComputeBilling.sol";

/**
 * @title ComputeBilling
 * @notice Compute billing with unified period-based settlement
 * @dev Accumulates charges continuously, settles by period
 */
abstract contract ComputeBilling is IBillingModule, IComputeBilling {
    // ============================================================================
    // State
    // ============================================================================

    BillingCore public immutable billingCore;

    mapping(uint16 => SKU) public skus;
    mapping(uint16 => SKU) public pendingSKUs;  // SKU changes waiting for next period
    mapping(uint16 => uint40) public skuChangeEffectivePeriod;  // When pending changes take effect
    mapping(address => AccountState) public accountState;
    mapping(address => mapping(uint16 => SKUAppCounts)) public accountSKUCounts;
    mapping(address => mapping(uint40 => PeriodSettlement)) public settlements;

    Resources public globalResources;

    // Who can settle periods (could be admin, automated keeper, etc.)
    mapping(address => bool) public settlers;

    // ============================================================================
    // Abstract Functions
    // ============================================================================

    function _getAppAccount(address app) internal view virtual returns (address);
    function _getAppSKU(address app) internal view virtual returns (uint16);
    function _isAppActive(address app) internal view virtual returns (bool);
    function _isAppRunning(address app) internal view virtual returns (bool);

    // ============================================================================
    // Modifiers
    // ============================================================================

    modifier onlySettler() {
        require(settlers[msg.sender], UnauthorizedSettler());
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
     * @notice Add or remove a settler
     */
    function _setSettler(address settler, bool authorized) internal {
        settlers[settler] = authorized;
        emit SettlerAuthorizationChanged(settler, authorized);
    }

    /**
     * @notice Add a new SKU (applies immediately)
     */
    function _addSKU(
        uint16 skuId,
        string memory name,
        uint96 runningRate,
        uint96 stoppedRate,
        uint16 vcpus,
        uint96 minimumDeposit
    ) internal {
        require(!skus[skuId].active, InvalidSKU());

        skus[skuId] = SKU({
            runningRate: runningRate,
            stoppedRate: stoppedRate,
            vcpus: vcpus,
            minimumDeposit: minimumDeposit,
            name: name,
            active: true
        });
        emit SKUUpdated(skuId, name, runningRate);
    }

    /**
     * @notice Update an existing SKU (applies at next period boundary)
     */
    function _updateSKU(
        uint16 skuId,
        string memory name,
        uint96 runningRate,
        uint96 stoppedRate,
        uint16 vcpus,
        uint96 minimumDeposit
    ) internal {
        require(skus[skuId].active, InvalidSKU());

        uint40 nextPeriod = billingCore.getCurrentPeriod() + 1;
        pendingSKUs[skuId] = SKU({
            runningRate: runningRate,
            stoppedRate: stoppedRate,
            vcpus: vcpus,
            minimumDeposit: minimumDeposit,
            name: name,
            active: true
        });
        skuChangeEffectivePeriod[skuId] = nextPeriod;
        emit SKUUpdated(skuId, name, runningRate);
    }

    /**
     * @notice Apply pending SKU changes if needed (lazy evaluation)
     */
    function _applyPendingSKUChanges(uint16 skuId) internal returns (bool changed) {
        uint40 currentPeriod = billingCore.getCurrentPeriod();
        if (skuChangeEffectivePeriod[skuId] > 0 && currentPeriod >= skuChangeEffectivePeriod[skuId]) {
            uint40 effectivePeriod = skuChangeEffectivePeriod[skuId];
            skus[skuId] = pendingSKUs[skuId];
            delete pendingSKUs[skuId];
            delete skuChangeEffectivePeriod[skuId];
            emit SKUChangeApplied(skuId, effectivePeriod);
            return true;
        }
        return false;
    }

    /**
     * @notice Get the effective SKU (with pending changes applied)
     */
    function _getEffectiveSKU(uint16 skuId) internal returns (SKU memory) {
        _applyPendingSKUChanges(skuId);
        return skus[skuId];
    }

    /**
     * @notice Set resource caps
     */
    function _setResourceCap(uint16 vcpuCap, uint16 vmInstanceCap) internal {
        globalResources.vcpuCap = vcpuCap;
        globalResources.vmInstanceCap = vmInstanceCap;
        emit ResourceCapUpdated(vcpuCap, vmInstanceCap);
    }

    // ============================================================================
    // Billing Lifecycle
    // ============================================================================

    /**
     * @notice Start billing for an app (initially in running state)
     */
    function _startBilling(address app, address account, uint16 skuId) internal {
        require(skuId != 0, InvalidSKU());
        SKU memory sku = _getEffectiveSKU(skuId);
        require(sku.active, InvalidSKU());

        // Check resource limits
        require(globalResources.vcpuUsed + sku.vcpus <= globalResources.vcpuCap, VCPUCapacityExceeded());
        require(globalResources.vmInstancesUsed + 1 <= globalResources.vmInstanceCap, InstanceCapacityExceeded());

        // Update resources
        globalResources.vcpuUsed += sku.vcpus;
        globalResources.vmInstancesUsed++;

        // Accrue any pending charges first
        _accrueCharges(account);

        // Update account state (app starts in running state)
        AccountState storage state = accountState[account];
        state.totalRunningRate += sku.runningRate;
        if (state.lastUpdate == 0) {
            state.lastUpdate = uint40(block.timestamp);
            state.chargesStartTime = uint40(block.timestamp);
        }

        // Track SKU usage
        accountSKUCounts[account][skuId].runningCount++;

        emit BillingStarted(app, account, sku.runningRate);
    }

    /**
     * @notice Set app to stopped state (lower rate)
     */
    function _setStoppedRate(address app, address account, uint16 skuId) internal {
        SKU memory sku = _getEffectiveSKU(skuId);

        // Accrue pending charges at running rate
        _accrueCharges(account);

        // Update rates and counts
        AccountState storage state = accountState[account];
        SKUAppCounts storage counts = accountSKUCounts[account][skuId];

        state.totalRunningRate -= sku.runningRate;
        state.totalStoppedRate += sku.stoppedRate;

        counts.runningCount--;
        counts.stoppedCount++;

        emit RateChanged(app, account, sku.runningRate, sku.stoppedRate);
    }

    /**
     * @notice Set app to running state (higher rate)
     */
    function _setRunningRate(address app, address account, uint16 skuId) internal {
        SKU memory sku = _getEffectiveSKU(skuId);

        // Accrue pending charges at stopped rate
        _accrueCharges(account);

        // Update rates and counts
        AccountState storage state = accountState[account];
        SKUAppCounts storage counts = accountSKUCounts[account][skuId];

        state.totalStoppedRate -= sku.stoppedRate;
        state.totalRunningRate += sku.runningRate;

        counts.stoppedCount--;
        counts.runningCount++;

        emit RateChanged(app, account, sku.stoppedRate, sku.runningRate);
    }

    /**
     * @notice Remove billing for an app completely
     */
    function _removeBilling(address app, address account, uint16 skuId) internal {
        SKU memory sku = _getEffectiveSKU(skuId);
        bool isRunning = _isAppRunning(app);

        // Free resources
        globalResources.vcpuUsed -= sku.vcpus;
        globalResources.vmInstancesUsed--;

        // Accrue pending charges
        _accrueCharges(account);

        // Update rates based on current state
        AccountState storage state = accountState[account];
        SKUAppCounts storage counts = accountSKUCounts[account][skuId];

        if (isRunning) {
            state.totalRunningRate -= sku.runningRate;
            counts.runningCount--;
        } else {
            state.totalStoppedRate -= sku.stoppedRate;
            counts.stoppedCount--;
        }

        emit BillingStopped(app, account);
    }

    // ============================================================================
    // Charge Accrual
    // ============================================================================

    /**
     * @notice Accrue charges up to current time
     * @dev Updates accruedCharges but doesn't charge BillingCore
     */
    function _accrueCharges(address account) internal {
        AccountState storage state = accountState[account];

        uint96 totalRate = state.totalRunningRate + state.totalStoppedRate;
        if (totalRate == 0 || state.lastUpdate == 0) {
            state.lastUpdate = uint40(block.timestamp);
            return;
        }

        uint40 currentTime = uint40(block.timestamp);
        if (currentTime <= state.lastUpdate) return;

        uint256 elapsed = currentTime - state.lastUpdate;
        uint96 charges = uint96(uint256(totalRate) * elapsed);

        state.accruedCharges += charges;
        state.lastUpdate = currentTime;
    }

    /**
     * @notice Public function to accrue charges (useful for view functions)
     */
    function accrueCharges(address account) external {
        _accrueCharges(account);
    }

    // ============================================================================
    // Period Settlement
    // ============================================================================

    /**
     * @notice Settle charges for an account for a specific period
     * @dev Called at the end of each period by authorized settler
     * @param account The account to settle
     * @param period The period to settle
     */
    function settlePeriod(address account, uint40 period) external onlySettler {
        require(period < billingCore.getCurrentPeriod(), InvalidPeriod());
        require(!settlements[account][period].settled, PeriodAlreadySettled());

        // First accrue any pending charges
        _accrueCharges(account);

        // Calculate what portion belongs to this period
        uint96 periodCharges = _calculatePeriodCharges(account, period);

        if (periodCharges > 0) {
            // Deduct from accrued charges
            AccountState storage state = accountState[account];
            if (state.accruedCharges >= periodCharges) {
                state.accruedCharges -= periodCharges;
            } else {
                // Use all accrued charges if calculation somehow exceeds
                periodCharges = state.accruedCharges;
                state.accruedCharges = 0;
            }

            // Charge through BillingCore
            billingCore.chargePeriod(account, periodCharges, period);

            // Update charges start time if we've fully settled up to this period
            (, uint40 periodEnd) = billingCore.getPeriodBounds(period);
            if (state.accruedCharges == 0) {
                state.chargesStartTime = periodEnd;
            }
        }

        // Mark as settled
        settlements[account][period] = PeriodSettlement({
            settled: true,
            amount: periodCharges,
            settledAt: uint40(block.timestamp)
        });

        emit PeriodSettled(account, period, periodCharges);
    }

    /**
     * @notice Batch settle multiple accounts for efficiency
     */
    function settlePeriodBatch(address[] calldata accounts, uint40 period) external onlySettler {
        require(period < billingCore.getCurrentPeriod(), InvalidPeriod());

        for (uint256 i = 0; i < accounts.length; i++) {
            if (!settlements[accounts[i]][period].settled) {
                _settlePeriodForAccount(accounts[i], period);
            }
        }
    }

    /**
     * @notice Internal settlement helper
     */
    function _settlePeriodForAccount(address account, uint40 period) internal {
        _accrueCharges(account);

        uint96 periodCharges = _calculatePeriodCharges(account, period);

        if (periodCharges > 0) {
            AccountState storage state = accountState[account];
            uint96 toCharge = periodCharges > state.accruedCharges ? state.accruedCharges : periodCharges;
            state.accruedCharges -= toCharge;

            billingCore.chargePeriod(account, toCharge, period);

            // Update charges start time if we've fully settled up to this period
            (, uint40 periodEnd) = billingCore.getPeriodBounds(period);
            if (state.accruedCharges == 0) {
                state.chargesStartTime = periodEnd;
            }
        }

        settlements[account][period] = PeriodSettlement({
            settled: true,
            amount: periodCharges,
            settledAt: uint40(block.timestamp)
        });

        emit PeriodSettled(account, period, periodCharges);
    }

    /**
     * @notice Calculate charges attributable to a specific period
     * @dev Calculates based on when charges started accumulating
     */
    function _calculatePeriodCharges(address account, uint40 period) internal view returns (uint96) {
        AccountState memory state = accountState[account];

        // Get period bounds
        (uint40 periodStart, uint40 periodEnd) = billingCore.getPeriodBounds(period);

        // No charges if account has no activity
        if (state.lastUpdate == 0 || state.chargesStartTime == 0) return 0;

        // No charges if billing started after this period ended
        if (state.chargesStartTime >= periodEnd) return 0;

        // Calculate the overlap between accrued charges time range and the period
        uint40 chargeStart = state.chargesStartTime > periodStart ? state.chargesStartTime : periodStart;
        uint40 chargeEnd = state.lastUpdate < periodEnd ? state.lastUpdate : periodEnd;

        // No charges if no overlap
        if (chargeStart >= chargeEnd) return 0;

        // Calculate total time span of accrued charges
        uint256 totalSpan = state.lastUpdate - state.chargesStartTime;
        if (totalSpan == 0) return 0;

        // Calculate time within this specific period
        uint256 periodSpan = chargeEnd - chargeStart;

        // Return proportional amount of accrued charges
        return uint96((uint256(state.accruedCharges) * periodSpan) / totalSpan);
    }

    // ============================================================================
    // View Functions
    // ============================================================================

    /**
     * @notice Check if account meets minimum deposit for SKU
     */
    function _requireMinimumDeposit(address account, uint16 skuId) internal {
        SKU memory sku = _getEffectiveSKU(skuId);
        if (sku.minimumDeposit == 0) return;

        int96 balance = billingCore.getBalance(account);
        uint96 outstanding = this.getOutstandingCharges(account);
        int96 effectiveBalance = balance - int96(outstanding);

        require(effectiveBalance >= int96(sku.minimumDeposit), InsufficientDeposit(sku.minimumDeposit, effectiveBalance));
    }

    /**
     * @notice Get current charges for an account (accrued but not settled)
     */
    function getOutstandingCharges(address account) external view returns (uint96) {
        AccountState memory state = accountState[account];

        uint96 totalRate = state.totalRunningRate + state.totalStoppedRate;
        if (totalRate == 0 || state.lastUpdate == 0) {
            return state.accruedCharges;
        }

        // Add charges since last update
        uint256 elapsed = block.timestamp > state.lastUpdate ? block.timestamp - state.lastUpdate : 0;
        uint96 additionalCharges = uint96(uint256(totalRate) * elapsed);

        return state.accruedCharges + additionalCharges;
    }

    /**
     * @notice Check if a period is settled for an account
     */
    function isPeriodSettled(address account, uint40 period) external view returns (bool) {
        return settlements[account][period].settled;
    }

    /**
     * @notice Get settlement details for a period
     */
    function getSettlement(address account, uint40 period) external view returns (
        bool settled,
        uint96 amount,
        uint40 settledAt
    ) {
        PeriodSettlement memory settlement = settlements[account][period];
        return (settlement.settled, settlement.amount, settlement.settledAt);
    }

    /**
     * @notice Estimate charges for current period (not yet settled)
     */
    function estimateCurrentPeriodCharges(address account) public view returns (uint96) {
        AccountState memory state = accountState[account];
        uint96 totalRate = state.totalRunningRate + state.totalStoppedRate;
        if (totalRate == 0) return 0;

        uint40 currentPeriod = billingCore.getCurrentPeriod();
        (uint40 periodStart, ) = billingCore.getPeriodBounds(currentPeriod);

        uint40 startTime = state.lastUpdate > periodStart ? state.lastUpdate : periodStart;
        uint256 elapsed = block.timestamp > startTime ? block.timestamp - startTime : 0;

        return uint96(uint256(totalRate) * elapsed);
    }

    /**
     * @notice Get resource usage
     */
    function getResourceUsage() external view returns (
        uint16 vcpuUsed,
        uint16 vcpuCap,
        uint16 vmInstancesUsed,
        uint16 vmInstanceCap
    ) {
        return (
            globalResources.vcpuUsed,
            globalResources.vcpuCap,
            globalResources.vmInstancesUsed,
            globalResources.vmInstanceCap
        );
    }

    /**
     * @notice Get charges for a specific period
     * @param account The account to query
     * @param period The period to query
     * @return amount The charges for this period
     * @return settled Whether the charges have been settled
     */
    function getChargesForPeriod(address account, uint40 period) external view returns (uint96 amount, bool settled) {
        PeriodSettlement memory settlement = settlements[account][period];

        if (settlement.settled) {
            // Return the settled amount
            return (settlement.amount, true);
        }

        // For unsettled periods, calculate the portion of accrued charges for this period
        if (period < billingCore.getCurrentPeriod()) {
            // Calculate what portion of accrued charges belongs to this period
            return (_calculatePeriodCharges(account, period), false);
        }

        // For current period, estimate charges
        if (period == billingCore.getCurrentPeriod()) {
            return (estimateCurrentPeriodCharges(account), false);
        }

        // Future periods have no charges
        return (0, false);
    }
}
