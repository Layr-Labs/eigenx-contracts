// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IBillingCore} from "./interfaces/IBillingCore.sol";
import {IBillingModule} from "./interfaces/IBillingModule.sol";
import {IComputeBilling} from "./interfaces/IComputeBilling.sol";
import {DateTimeLib} from "solady/utils/DateTimeLib.sol";

/**
 * @title ComputeBilling
 * @notice Continuous billing for compute resources with running/stopped rates and capacity management
 */
abstract contract ComputeBilling is IComputeBilling, IBillingModule {
    IBillingCore public immutable billing;

    // ============================================================================
    // State
    // ============================================================================

    mapping(uint16 skuId => SKU) public skus;

    // Pending SKU changes that take effect next period
    mapping(uint16 skuId => SKU) public pendingSKUs;
    mapping(uint16 skuId => uint40) public skuChangeEffectivePeriod;

    // Account billing state
    mapping(address account => mapping(uint16 skuId => SKUAppCounts)) public accountSKUCounts;
    mapping(address account => uint16[]) public accountSKUs;
    mapping(address account => AccountBilling) public accountBilling;

    // Resource tracking
    Resources public globalResources;

    // ============================================================================
    // Abstract Functions
    // ============================================================================

    function _getAppAccount(address app) internal view virtual returns (address);
    function _getAppSKU(address app) internal view virtual returns (uint16);
    function _isAppRunning(address app) internal view virtual returns (bool);
    function _isAppActive(address app) internal view virtual returns (bool);

    // ============================================================================
    // Constructor
    // ============================================================================

    constructor(address _billing) {
        billing = IBillingCore(_billing);
    }

    // ============================================================================
    // External Functions
    // ============================================================================

    /**
     * @notice Public settlement function
     */
    function settleAccount(address account) external {
        _settleAccount(account);
    }

    // ============================================================================
    // Admin Functions
    // ============================================================================

    /**
     * @notice Set SKU rates (takes effect next period)
     */
    function _setSKURate(
        uint16 skuID,
        string calldata name,
        uint96 runningRate,
        uint96 stoppedRate,
        uint16 vcpus,
        uint96 minimumDeposit
    ) internal {
        uint40 nextPeriod = billing.getCurrentPeriod() + 1;

        pendingSKUs[skuID] = SKU({
            runningRate: runningRate,
            stoppedRate: stoppedRate,
            vcpus: vcpus,
            minimumDeposit: minimumDeposit,
            description: name
        });

        skuChangeEffectivePeriod[skuID] = nextPeriod;

        emit SKURatesSet(skuID, nextPeriod, name, runningRate, stoppedRate, vcpus);
    }

    /**
     * @notice Set resource capacity limits
     */
    function _setResourceCap(uint16 vcpuCap, uint16 vmInstanceCap) internal {
        globalResources.vcpuCap = vcpuCap;
        globalResources.vmInstanceCap = vmInstanceCap;
        emit ResourceCapSet(vcpuCap, vmInstanceCap);
    }

    /**
     * @notice Check if account has sufficient deposit to cover SKU minimum
     * @param account The billing account to check
     * @param skuID The SKU ID to check minimum deposit for
     * @dev Checks effective balance (balance - outstanding charges) against SKU minimum
     */
    function _requireMinimumDeposit(address account, uint16 skuID) internal view {
        SKU memory sku = skus[skuID];

        // Skip check if no minimum required
        if (sku.minimumDeposit == 0) return;

        // Get current balance
        int96 balance = billing.getBalance(account);

        // Get outstanding charges for this account
        uint96 outstanding = getOutstandingCharges(account);

        // Calculate effective balance
        int96 effectiveBalance = balance - int96(outstanding);

        // Require effective balance meets minimum
        require(
            effectiveBalance >= int96(sku.minimumDeposit),
            InsufficientDepositForSKU(sku.minimumDeposit, effectiveBalance)
        );
    }

    // ============================================================================
    // Rate Management
    // ============================================================================

    /**
     * @notice Get effective SKU (current or pending based on period)
     */
    function _getEffectiveSKU(uint16 skuId) internal view returns (SKU memory) {
        uint40 currentPeriod = billing.getCurrentPeriod();

        // Check if pending changes should be applied
        if (skuChangeEffectivePeriod[skuId] > 0 && currentPeriod >= skuChangeEffectivePeriod[skuId]) {
            return pendingSKUs[skuId];
        }

        return skus[skuId];
    }

    /**
     * @notice Apply pending SKU changes if needed
     */
    function _applyPendingSKUChanges(uint16 skuId) internal returns (bool changed) {
        uint40 currentPeriod = billing.getCurrentPeriod();

        if (skuChangeEffectivePeriod[skuId] > 0 && currentPeriod >= skuChangeEffectivePeriod[skuId]) {
            skus[skuId] = pendingSKUs[skuId];
            delete pendingSKUs[skuId];
            delete skuChangeEffectivePeriod[skuId];
            return true;
        }

        return false;
    }

    /**
     * @notice Update account rates if period changed
     */
    function _updateAccountRatesIfNeeded(address account) internal returns (bool updated) {
        uint40 currentPeriod = billing.getCurrentPeriod();
        AccountBilling storage billing_ = accountBilling[account];

        // Only recalculate if we're in a new period
        if (billing_.lastRateUpdate >= currentPeriod) {
            return false;
        }

        uint96 newTotalRate = 0;
        uint16[] memory skuList = accountSKUs[account];

        for (uint256 i = 0; i < skuList.length; i++) {
            uint16 skuId = skuList[i];

            // Apply any pending changes
            _applyPendingSKUChanges(skuId);

            // Calculate rate with current SKU
            SKU memory sku = skus[skuId];
            SKUAppCounts memory counts = accountSKUCounts[account][skuId];

            newTotalRate += uint96(counts.runningCount) * sku.runningRate;
            newTotalRate += uint96(counts.stoppedCount) * sku.stoppedRate;
        }

        billing_.totalRate = newTotalRate;
        billing_.lastRateUpdate = currentPeriod;

        return true;
    }

    // ============================================================================
    // Billing Lifecycle
    // ============================================================================

    /**
     * @notice Start billing for an app
     */
    function _startBilling(address app, uint16 skuID, address account) internal {
        require(skuID != 0, InvalidSKU());

        // Apply any pending changes first
        _applyPendingSKUChanges(skuID);

        SKU memory sku = skus[skuID];
        require(sku.runningRate > 0 || sku.stoppedRate > 0, SKUNotConfigured());

        // Check resource limits
        require(globalResources.vcpuUsed + sku.vcpus <= globalResources.vcpuCap, VCPUCapacityExceeded());
        require(globalResources.vmInstancesUsed + 1 <= globalResources.vmInstanceCap, InstanceCapacityExceeded());

        // Update resources
        globalResources.vcpuUsed += sku.vcpus;
        globalResources.vmInstancesUsed++;

        // Initialize account if needed
        AccountBilling storage billing_ = accountBilling[account];
        if (billing_.lastSettled == 0) {
            billing_.lastSettled = uint40(block.timestamp);
            billing_.lastRateUpdate = billing.getCurrentPeriod();
        }

        // Track SKU usage
        SKUAppCounts storage counts = accountSKUCounts[account][skuID];
        if (counts.runningCount == 0 && counts.stoppedCount == 0) {
            accountSKUs[account].push(skuID);
        }
        counts.runningCount++;

        // Update rate
        billing_.totalRate += sku.runningRate;

        emit BillingStarted(app, account, skuID, sku.runningRate);
    }

    /**
     * @notice Set app to running state
     */
    function _setRunningRate(address app, address account) internal {
        require(_isAppActive(app), AppNotBilled());
        require(!_isAppRunning(app), AlreadyRunning());

        _settleAccount(account);

        uint16 skuId = _getAppSKU(app);
        _applyPendingSKUChanges(skuId);

        SKU memory sku = skus[skuId];

        // Update counts
        SKUAppCounts storage counts = accountSKUCounts[account][skuId];
        counts.stoppedCount--;
        counts.runningCount++;

        // Update rate
        AccountBilling storage billing_ = accountBilling[account];
        billing_.totalRate = billing_.totalRate - sku.stoppedRate + sku.runningRate;

        emit BillingRateChanged(app, true, sku.runningRate);
    }

    /**
     * @notice Set app to stopped state
     */
    function _setStoppedRate(address app, address account) internal {
        require(_isAppActive(app), AppNotBilled());
        require(_isAppRunning(app), AlreadyStopped());

        _settleAccount(account);

        uint16 skuId = _getAppSKU(app);
        _applyPendingSKUChanges(skuId);

        SKU memory sku = skus[skuId];

        // Update counts
        SKUAppCounts storage counts = accountSKUCounts[account][skuId];
        counts.runningCount--;
        counts.stoppedCount++;

        // Update rate
        AccountBilling storage billing_ = accountBilling[account];
        billing_.totalRate = billing_.totalRate - sku.runningRate + sku.stoppedRate;

        emit BillingRateChanged(app, false, sku.stoppedRate);
    }

    /**
     * @notice Remove billing for an app
     */
    function _removeBilling(address app, address account, uint16 skuID) internal {
        require(_isAppActive(app), AppNotBilled());

        _settleAccount(account);

        _applyPendingSKUChanges(skuID);
        SKU memory sku = skus[skuID];
        bool isRunning = _isAppRunning(app);

        // Free resources
        globalResources.vcpuUsed -= sku.vcpus;
        globalResources.vmInstancesUsed--;

        // Update counts and rate
        SKUAppCounts storage counts = accountSKUCounts[account][skuID];
        AccountBilling storage billing_ = accountBilling[account];
        if (isRunning) {
            counts.runningCount--;
            billing_.totalRate -= sku.runningRate;
        } else {
            counts.stoppedCount--;
            billing_.totalRate -= sku.stoppedRate;
        }

        // Clean up if no more apps using this SKU
        if (counts.runningCount == 0 && counts.stoppedCount == 0) {
            _removeSKUFromAccount(account, skuID);
        }

        emit BillingRemoved(app);
    }

    /**
     * @notice Change SKU for an app
     */
    function _changeSKU(address app, address account, uint16 oldSKUID, uint16 newSKUID, bool isRunning) internal {
        require(_isAppActive(app), AppNotBilled());

        _settleAccount(account);

        // Apply pending changes for both SKUs
        _applyPendingSKUChanges(oldSKUID);
        _applyPendingSKUChanges(newSKUID);

        SKU memory oldSKU = skus[oldSKUID];
        SKU memory newSKU = skus[newSKUID];
        require(newSKU.runningRate > 0 || newSKU.stoppedRate > 0, InvalidNewSKU());

        // Check resource delta
        int16 vcpuDelta = int16(newSKU.vcpus) - int16(oldSKU.vcpus);
        if (vcpuDelta > 0) {
            require(globalResources.vcpuUsed + uint16(vcpuDelta) <= globalResources.vcpuCap, VCPUCapacityExceeded());
            globalResources.vcpuUsed += uint16(vcpuDelta);
        } else if (vcpuDelta < 0) {
            globalResources.vcpuUsed -= uint16(-vcpuDelta);
        }

        // Update old SKU counts
        SKUAppCounts storage oldCounts = accountSKUCounts[account][oldSKUID];
        if (isRunning) {
            oldCounts.runningCount--;
        } else {
            oldCounts.stoppedCount--;
        }

        if (oldCounts.runningCount == 0 && oldCounts.stoppedCount == 0) {
            _removeSKUFromAccount(account, oldSKUID);
        }

        // Update new SKU counts
        SKUAppCounts storage newCounts = accountSKUCounts[account][newSKUID];
        if (newCounts.runningCount == 0 && newCounts.stoppedCount == 0) {
            accountSKUs[account].push(newSKUID);
        }

        if (isRunning) {
            newCounts.runningCount++;
        } else {
            newCounts.stoppedCount++;
        }

        // Update rate
        uint96 oldRate = isRunning ? oldSKU.runningRate : oldSKU.stoppedRate;
        uint96 newRate = isRunning ? newSKU.runningRate : newSKU.stoppedRate;
        AccountBilling storage billing_ = accountBilling[account];
        billing_.totalRate = billing_.totalRate - oldRate + newRate;
    }

    /**
     * @notice Change billing account for an app
     * TODO: We need an approval system for billing account transfers to be safe
     */
    function _changeAccount(address app, address oldAccount, address newAccount) internal {
        require(_isAppActive(app), AppNotBilled());
        require(newAccount != address(0), InvalidAccount());

        if (oldAccount == newAccount) return;

        // Settle both accounts first
        _settleAccount(oldAccount);
        _settleAccount(newAccount);

        uint16 skuId = _getAppSKU(app);
        _applyPendingSKUChanges(skuId);

        SKU memory sku = skus[skuId];
        bool isRunning = _isAppRunning(app);

        // Initialize new account if needed
        AccountBilling storage newBilling = accountBilling[newAccount];
        if (newBilling.lastSettled == 0) {
            newBilling.lastSettled = uint40(block.timestamp);
            newBilling.lastRateUpdate = billing.getCurrentPeriod();
        }

        // Move from old account
        AccountBilling storage oldBilling = accountBilling[oldAccount];
        SKUAppCounts storage oldCounts = accountSKUCounts[oldAccount][skuId];
        if (isRunning) {
            oldCounts.runningCount--;
            oldBilling.totalRate -= sku.runningRate;
        } else {
            oldCounts.stoppedCount--;
            oldBilling.totalRate -= sku.stoppedRate;
        }

        if (oldCounts.runningCount == 0 && oldCounts.stoppedCount == 0) {
            _removeSKUFromAccount(oldAccount, skuId);
        }

        // Add to new account
        SKUAppCounts storage newCounts = accountSKUCounts[newAccount][skuId];
        if (newCounts.runningCount == 0 && newCounts.stoppedCount == 0) {
            accountSKUs[newAccount].push(skuId);
        }

        if (isRunning) {
            newCounts.runningCount++;
            newBilling.totalRate += sku.runningRate;
        } else {
            newCounts.stoppedCount++;
            newBilling.totalRate += sku.stoppedRate;
        }
    }

    // ============================================================================
    // Settlement
    // ============================================================================

    /**
     * @notice Settle all charges for an account
     */
    function _settleAccount(address account) internal {
        AccountBilling storage billing_ = accountBilling[account];
        uint40 lastSettled = billing_.lastSettled;

        // Initialize if first time
        if (lastSettled == 0) {
            billing_.lastSettled = uint40(block.timestamp);
            return;
        }

        // Update rates if period changed
        _updateAccountRatesIfNeeded(account);

        uint96 rate = billing_.totalRate;
        if (rate == 0) return;

        uint40 currentTime = uint40(block.timestamp);
        if (currentTime <= lastSettled) return;

        uint40 currentPeriod = billing.getCurrentPeriod();
        uint40 lastSettledPeriod = billing.getPeriodForTimestamp(lastSettled);
        uint40 genesisTime = billing.genesisTime();

        // Handle same period (most common case)
        if (currentPeriod == lastSettledPeriod) {
            uint40 periodStart = uint40(DateTimeLib.addMonths(genesisTime, currentPeriod));
            uint40 periodEnd = uint40(DateTimeLib.addMonths(genesisTime, currentPeriod + 1));

            uint96 amount = _calculateChargesInPeriod(lastSettled, currentTime, rate, periodStart, periodEnd);

            if (amount > 0) {
                billing.chargePeriod(account, amount, currentPeriod);
                emit BillingSettled(account, amount, currentPeriod);
            }
        }
        // Handle crossing one period boundary (common case)
        else if (currentPeriod == lastSettledPeriod + 1) {
            // Charge for time in last period
            uint40 lastPeriodStart = uint40(DateTimeLib.addMonths(genesisTime, lastSettledPeriod));
            uint40 lastPeriodEnd = uint40(DateTimeLib.addMonths(genesisTime, lastSettledPeriod + 1));

            uint96 lastPeriodAmount =
                _calculateChargesInPeriod(lastSettled, currentTime, rate, lastPeriodStart, lastPeriodEnd);
            if (lastPeriodAmount > 0) {
                billing.chargePeriod(account, lastPeriodAmount, lastSettledPeriod);
                emit BillingSettled(account, lastPeriodAmount, lastSettledPeriod);
            }

            // Charge for time in current period
            uint40 currentPeriodStart = uint40(DateTimeLib.addMonths(genesisTime, currentPeriod));
            uint40 currentPeriodEnd = uint40(DateTimeLib.addMonths(genesisTime, currentPeriod + 1));

            uint96 currentPeriodAmount =
                _calculateChargesInPeriod(lastSettled, currentTime, rate, currentPeriodStart, currentPeriodEnd);
            if (currentPeriodAmount > 0) {
                billing.chargePeriod(account, currentPeriodAmount, currentPeriod);
                emit BillingSettled(account, currentPeriodAmount, currentPeriod);
            }
        }
        // Handle multiple periods (rare case)
        else {
            _settleMultiplePeriods(
                account, rate, lastSettled, currentTime, lastSettledPeriod, currentPeriod, genesisTime
            );
        }

        billing_.lastSettled = currentTime;
    }

    /**
     * @notice Handle settlement across multiple periods (rare case)
     * @dev Separated from main function to optimize for common cases
     */
    function _settleMultiplePeriods(
        address account,
        uint96 rate,
        uint40 startTime,
        uint40 endTime,
        uint40 startPeriod,
        uint40 endPeriod,
        uint40 genesisTime
    ) private {
        // Calculate charges for each period using shared helper
        for (uint40 period = startPeriod; period <= endPeriod; period++) {
            uint40 periodStart = uint40(DateTimeLib.addMonths(genesisTime, period));
            uint40 periodEnd = uint40(DateTimeLib.addMonths(genesisTime, period + 1));

            uint96 periodAmount = _calculateChargesInPeriod(startTime, endTime, rate, periodStart, periodEnd);

            if (periodAmount > 0) {
                billing.chargePeriod(account, periodAmount, period);
                emit BillingSettled(account, periodAmount, period);
            }
        }
    }

    function _removeSKUFromAccount(address account, uint16 skuId) internal {
        uint16[] storage skuList = accountSKUs[account];
        for (uint256 i = 0; i < skuList.length; i++) {
            if (skuList[i] == skuId) {
                skuList[i] = skuList[skuList.length - 1];
                skuList.pop();
                break;
            }
        }
    }

    /**
     * @notice Calculate charges that accrued in a specific period window
     * @param startTime The start of the billing window
     * @param endTime The end of the billing window
     * @param rate The rate in effect during this window
     * @param periodStart The start timestamp of the period
     * @param periodEnd The end timestamp of the period
     * @return charges The amount charged during the overlap of billing window and period
     */
    function _calculateChargesInPeriod(
        uint40 startTime,
        uint40 endTime,
        uint96 rate,
        uint40 periodStart,
        uint40 periodEnd
    ) internal pure returns (uint96 charges) {
        // Calculate overlap between [startTime, endTime] and [periodStart, periodEnd]
        uint40 effectiveStart = startTime > periodStart ? startTime : periodStart;
        uint40 effectiveEnd = endTime < periodEnd ? endTime : periodEnd;

        if (effectiveEnd <= effectiveStart) return 0;

        uint256 timeInPeriod = effectiveEnd - effectiveStart;
        return uint96(uint256(rate) * timeInPeriod);
    }

    // ============================================================================
    // View Functions
    // ============================================================================

    function getResourceUsage() external view returns (uint16 vcpuUsed, uint16 vmUsed, uint16 vcpuCap, uint16 vmCap) {
        return (
            globalResources.vcpuUsed,
            globalResources.vmInstancesUsed,
            globalResources.vcpuCap,
            globalResources.vmInstanceCap
        );
    }

    function getPendingChargesForApp(address app) public view returns (uint96) {
        if (!_isAppActive(app)) return 0;

        address account = _getAppAccount(app);
        AccountBilling memory billing_ = accountBilling[account];
        if (billing_.lastSettled == 0) return 0;

        uint16 skuId = _getAppSKU(app);
        SKU memory sku = _getEffectiveSKU(skuId);

        uint96 rate = _isAppRunning(app) ? sku.runningRate : sku.stoppedRate;
        uint256 elapsed = block.timestamp - billing_.lastSettled;

        return uint96(uint256(rate) * elapsed);
    }

    /**
     * @notice Get charges for an account in a specific period (IBillingModule interface)
     * @param account The account to query
     * @param period The period to query
     * @return amount Charges accrued in that period (whether settled or not)
     */
    function getChargesForPeriod(address account, uint40 period) external view override returns (uint96 amount) {
        AccountBilling memory billing_ = accountBilling[account];
        if (billing_.lastSettled == 0 || billing_.totalRate == 0) return 0;

        uint40 currentTime = uint40(block.timestamp);
        uint40 genesisTime = billing.genesisTime();

        // Calculate period boundaries
        uint40 periodStart = uint40(DateTimeLib.addMonths(genesisTime, period));
        uint40 periodEnd = uint40(DateTimeLib.addMonths(genesisTime, period + 1));

        // Use shared helper to calculate charges in this period
        return _calculateChargesInPeriod(billing_.lastSettled, currentTime, billing_.totalRate, periodStart, periodEnd);
    }

    /**
     * @notice Get total outstanding charges not yet charged to BillingCore (IBillingModule interface)
     * @param account The account to query
     * @return amount Total charges accumulated since last settlement
     */
    function getOutstandingCharges(address account) public view override returns (uint96 amount) {
        AccountBilling memory billing_ = accountBilling[account];
        if (billing_.lastSettled == 0 || billing_.totalRate == 0) return 0;

        uint40 currentTime = uint40(block.timestamp);
        if (currentTime <= billing_.lastSettled) return 0;

        uint256 elapsed = currentTime - billing_.lastSettled;
        return uint96(uint256(billing_.totalRate) * elapsed);
    }
}
