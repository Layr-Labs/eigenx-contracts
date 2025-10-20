// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ComputeBilling} from "../../../src/billing/ComputeBilling.sol";

/**
 * @title MockComputeBilling
 * @notice Mock implementation of ComputeBilling for testing
 */
contract MockComputeBilling is ComputeBilling {
    mapping(address => address) public appAccounts;
    mapping(address => uint16) public appSKUs;
    mapping(address => bool) public appActive;
    mapping(address => bool) public appRunning;

    constructor(address _billingCore) ComputeBilling(_billingCore) {}

    function _getAppAccount(address app) internal view override returns (address) {
        return appAccounts[app];
    }

    function _getAppSKU(address app) internal view override returns (uint16) {
        return appSKUs[app];
    }

    function _isAppActive(address app) internal view override returns (bool) {
        return appActive[app];
    }

    function _isAppRunning(address app) internal view override returns (bool) {
        return appRunning[app];
    }

    // Test helpers
    function setAppAccount(address app, address account) external {
        appAccounts[app] = account;
    }

    function setAppSKU(address app, uint16 skuId) external {
        appSKUs[app] = skuId;
    }

    function setAppActive(address app, bool active) external {
        appActive[app] = active;
    }

    function setAppRunning(address app, bool running) external {
        appRunning[app] = running;
    }

    // Expose internal functions for testing
    function startBilling(address app, address account, uint16 skuId) external {
        _startBilling(app, account, skuId);
    }

    function removeBilling(address app, address account, uint16 skuId) external {
        _removeBilling(app, account, skuId);
    }

    function addSKU(
        uint16 skuId,
        string memory name,
        uint96 runningRate,
        uint96 stoppedRate,
        uint16 vcpus,
        uint96 minimumDeposit
    ) external {
        _addSKU(skuId, name, runningRate, stoppedRate, vcpus, minimumDeposit);
    }

    function setSettler(address settler, bool authorized) external {
        _setSettler(settler, authorized);
    }

    function setResourceCap(uint16 vcpuCap, uint16 vmInstanceCap) external {
        _setResourceCap(vcpuCap, vmInstanceCap);
    }
}
