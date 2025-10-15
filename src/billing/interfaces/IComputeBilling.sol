// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IComputeBilling {
    struct SKU {
        uint96 runningRate;     // Cost per second when app is running
        uint96 stoppedRate;     // Cost per second when app is stopped
        uint16 vcpus;           // Virtual CPUs required
    }

    struct Resources {
        uint16 vcpuCap;         // Max vCPUs allowed globally
        uint16 vmInstanceCap;   // Max VM instances allowed globally
        uint16 vcpuUsed;        // Current global vCPU usage
        uint16 vmInstancesUsed; // Current global VM instance count
    }

    // Errors
    error InvalidSKU(uint16 skuID);
    error InvalidAccount();
    error LengthMismatch();
    error ResourceCapExceeded();

    // Events
    event AppBillingCreated(address indexed app, address indexed account, uint16 skuID);
    event AppBillingRunning(address indexed app);
    event AppBillingStopped(address indexed app);
    event AppBillingSKUChanged(address indexed app, uint16 oldSKUID, uint16 newSKUID);
    event AppBillingAccountChanged(address indexed app, address indexed oldAccount, address indexed newAccount);
    event AppBillingRemoved(address indexed app);
    event SKURatesSet(uint16 indexed skuID, string name, uint96 runningRate, uint96 stoppedRate, uint16 vcpus);
    event ResourceCapSet(uint16 vcpuCap, uint16 vmInstanceCap);
}
