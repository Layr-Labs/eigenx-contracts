// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IBillingCore} from "./interfaces/IBillingCore.sol";
import {IComputeBilling} from "./interfaces/IComputeBilling.sol";

/**
 * @title ComputeBilling
 * @notice Abstract base contract for compute resource billing with dual rates
 * @dev Compute resources consume at continuous rates (running vs stopped).
 *      Extend this contract to add product-specific logic and access control.
 */
abstract contract ComputeBilling is IComputeBilling {
    IBillingCore public immutable billingCore;

    mapping(uint16 skuID => SKU) public skus;  // SKU ID => SKU details
    mapping(address account => address[]) public billedApps;  // Billing account => apps under that account
    Resources public globalResources;  // Global resource tracking

    constructor(address _billingCore) {
        billingCore = IBillingCore(_billingCore);
    }

    // ============================================================================
    // Internal Admin Functions (to be called by extending contract)
    // ============================================================================

    function _setSKURate(uint16 skuID, string calldata name, uint96 runningRate, uint96 stoppedRate, uint16 vcpus) internal {
        skus[skuID] = SKU({
            runningRate: runningRate,
            stoppedRate: stoppedRate,
            vcpus: vcpus
        });
        emit SKURatesSet(skuID, name, runningRate, stoppedRate, vcpus);
    }

    function _setResourceCap(uint16 vcpuCap, uint16 vmInstanceCap) internal {
        globalResources.vcpuCap = vcpuCap;
        globalResources.vmInstanceCap = vmInstanceCap;
        emit ResourceCapSet(vcpuCap, vmInstanceCap);
    }

    function _hasResourceCapacity(uint16 skuID) internal view returns (bool) {
        SKU memory sku = skus[skuID];

        if (globalResources.vcpuUsed + sku.vcpus > globalResources.vcpuCap) return false;
        if (globalResources.vmInstancesUsed + 1 > globalResources.vmInstanceCap) return false;
        return true;
    }

    // ============================================================================
    // Internal Billing Lifecycle Functions
    // ============================================================================

    function _startBilling(address app, uint16 skuID, address account) internal {
        SKU memory sku = skus[skuID];
        require(sku.runningRate > 0 || sku.stoppedRate > 0, InvalidSKU(skuID));

        // Check global resource capacity before starting billing
        require(_hasResourceCapacity(skuID), ResourceCapExceeded());

        // Track app under billing account
        billedApps[account].push(app);

        // Increment global resource usage
        globalResources.vcpuUsed += sku.vcpus;
        globalResources.vmInstancesUsed += 1;

        // Register running rate with billing core (apps start in running state)
        billingCore.adjustAccountRate(account, sku.runningRate, true);

        emit AppBillingCreated(app, account, skuID);
    }

    // ============================================================================
    // Internal Billing Functions
    // ============================================================================

    function _setRunningRate(address app, address account, uint16 skuID) internal {
        SKU memory sku = skus[skuID];

        // Switch from stopped rate to running rate
        billingCore.adjustAccountRate(account, sku.stoppedRate, false);
        billingCore.adjustAccountRate(account, sku.runningRate, true);

        emit AppBillingRunning(app);
    }

    function _setStoppedRate(address app, address account, uint16 skuID) internal {
        SKU memory sku = skus[skuID];

        // Switch from running rate to stopped rate
        billingCore.adjustAccountRate(account, sku.runningRate, false);
        billingCore.adjustAccountRate(account, sku.stoppedRate, true);

        emit AppBillingStopped(app);
    }

    function _changeSKU(address app, address account, uint16 oldSKUID, uint16 newSKUID, bool isRunning) internal {
        SKU memory oldSKU = skus[oldSKUID];
        SKU memory newSKU = skus[newSKUID];
        require(newSKU.runningRate > 0 || newSKU.stoppedRate > 0, InvalidSKU(newSKUID));

        // Calculate net change in vCPUs
        int32 vcpuDelta = int32(uint32(newSKU.vcpus)) - int32(uint32(oldSKU.vcpus));

        // Check global capacity if vCPUs are increasing
        if (vcpuDelta > 0) {
            uint16 newVcpuUsed = globalResources.vcpuUsed + uint16(uint32(vcpuDelta));
            require(newVcpuUsed <= globalResources.vcpuCap, ResourceCapExceeded());
        }

        // Adjust global vCPU usage
        if (vcpuDelta > 0) {
            globalResources.vcpuUsed += uint16(uint32(vcpuDelta));
        } else if (vcpuDelta < 0) {
            globalResources.vcpuUsed -= uint16(uint32(-vcpuDelta));
        }

        // Remove old rate and add new rate based on current state
        if (isRunning) {
            billingCore.adjustAccountRate(account, oldSKU.runningRate, false);
            billingCore.adjustAccountRate(account, newSKU.runningRate, true);
        } else {
            billingCore.adjustAccountRate(account, oldSKU.stoppedRate, false);
            billingCore.adjustAccountRate(account, newSKU.stoppedRate, true);
        }

        emit AppBillingSKUChanged(app, oldSKUID, newSKUID);
    }

    function _changeAccount(address app, address oldAccount, address newAccount, uint16 skuID, bool isRunning) internal {
        require(newAccount != address(0), InvalidAccount());
        if (oldAccount == newAccount) return;

        SKU memory sku = skus[skuID];
        uint96 currentRate = isRunning ? sku.runningRate : sku.stoppedRate;

        // Move app from old billing account to new billing account
        _removeAppFromAccount(oldAccount, app);
        billedApps[newAccount].push(app);

        // Move billing from old account to new account
        // (Global resources stay the same - just changing who pays)
        billingCore.adjustAccountRate(oldAccount, currentRate, false);
        billingCore.adjustAccountRate(newAccount, currentRate, true);

        emit AppBillingAccountChanged(app, oldAccount, newAccount);
    }

    function _removeBilling(address app, address account, uint16 skuID, bool isRunning) internal {
        SKU memory sku = skus[skuID];
        uint96 currentRate = isRunning ? sku.runningRate : sku.stoppedRate;

        // Decrement global resource usage
        globalResources.vcpuUsed -= sku.vcpus;
        globalResources.vmInstancesUsed -= 1;

        // Remove billing from account
        billingCore.adjustAccountRate(account, currentRate, false);

        // Remove app from billing account's list
        _removeAppFromAccount(account, app);

        emit AppBillingRemoved(app);
    }

    // ============================================================================
    // View Functions
    // ============================================================================

    function getBilledApps(address account) external view returns (address[] memory) {
        return billedApps[account];
    }

    function getAccountBurnRate(
        address account,
        uint16[] calldata skuIDs,
        bool[] calldata runningStates
    ) external view returns (uint96 totalRate) {
        address[] memory appList = billedApps[account];
        require(appList.length == skuIDs.length, LengthMismatch());
        require(appList.length == runningStates.length, LengthMismatch());

        for (uint256 i = 0; i < appList.length; i++) {
            SKU memory sku = skus[skuIDs[i]];
            totalRate += runningStates[i] ? sku.runningRate : sku.stoppedRate;
        }
    }

    function getAccountAppsDetailed(
        address account,
        uint16[] calldata skuIDs,
        bool[] calldata runningStates
    ) external view returns (
        address[] memory appAddresses,
        uint96[] memory currentRates,
        uint96 totalRate
    ) {
        address[] memory appList = billedApps[account];
        require(appList.length == skuIDs.length, LengthMismatch());
        require(appList.length == runningStates.length, LengthMismatch());

        uint256 length = appList.length;
        appAddresses = new address[](length);
        currentRates = new uint96[](length);

        for (uint256 i = 0; i < length; i++) {
            appAddresses[i] = appList[i];

            SKU memory sku = skus[skuIDs[i]];
            uint96 currentRate = runningStates[i] ? sku.runningRate : sku.stoppedRate;
            currentRates[i] = currentRate;
            totalRate += currentRate;
        }
    }

    function getGlobalResourceUsage() external view returns (
        uint16 vcpuCap,
        uint16 vmInstanceCap,
        uint16 vcpuUsed,
        uint16 vmInstancesUsed,
        uint16 vcpuAvailable,
        uint16 vmInstancesAvailable
    ) {
        vcpuCap = globalResources.vcpuCap;
        vmInstanceCap = globalResources.vmInstanceCap;
        vcpuUsed = globalResources.vcpuUsed;
        vmInstancesUsed = globalResources.vmInstancesUsed;
        vcpuAvailable = globalResources.vcpuCap - globalResources.vcpuUsed;
        vmInstancesAvailable = globalResources.vmInstanceCap - globalResources.vmInstancesUsed;
    }

    // ============================================================================
    // Internal Functions
    // ============================================================================

    function _removeAppFromAccount(address account, address app) internal {
        address[] storage appList = billedApps[account];
        for (uint256 i = 0; i < appList.length; i++) {
            if (appList[i] == app) {
                appList[i] = appList[appList.length - 1];
                appList.pop();
                break;
            }
        }
    }
}
