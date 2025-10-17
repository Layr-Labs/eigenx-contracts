// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/**
 * @title IComputeBilling
 * @notice Interface for compute billing system with types, events, and errors
 */
interface IComputeBilling {
    // ============================================================================
    // Data Types
    // ============================================================================

    struct SKU {
        uint96 runningRate;
        uint96 stoppedRate;
        uint16 vcpus;
        uint96 minimumDeposit;
        string description;
    }

    struct SKUAppCounts {
        uint16 runningCount;
        uint16 stoppedCount;
    }

    struct AccountBilling {
        uint96 totalRate;
        uint40 lastSettled;
        uint40 lastRateUpdate;
    }

    struct Resources {
        uint16 vcpuCap; // Max vCPUs allowed
        uint16 vmInstanceCap; // Max instances allowed
        uint16 vcpuUsed; // Current vCPU usage
        uint16 vmInstancesUsed; // Current instance count
    }

    // ============================================================================
    // Events
    // ============================================================================

    event BillingStarted(address indexed app, address indexed account, uint16 skuId, uint96 rate);
    event BillingSettled(address indexed account, uint96 amount, uint40 period);
    event BillingRateChanged(address indexed app, bool isRunning, uint96 newRate);
    event BillingRemoved(address indexed app);
    event SKURatesSet(
        uint16 indexed skuID, uint40 effectivePeriod, string name, uint96 runningRate, uint96 stoppedRate, uint16 vcpus
    );
    event ResourceCapSet(uint16 vcpuCap, uint16 vmInstanceCap);

    // ============================================================================
    // Custom Errors
    // ============================================================================

    error AppAlreadyBilled();
    error InvalidSKU();
    error SKUNotConfigured();
    error VCPUCapacityExceeded();
    error InstanceCapacityExceeded();
    error AppNotBilled();
    error AlreadyRunning();
    error AlreadyStopped();
    error InvalidNewSKU();
    error InvalidAccount();
    error InsufficientDepositForSKU(uint96 required, int96 available);

    // ============================================================================
    // Functions
    // ============================================================================

    function settleAccount(address account) external;
    function getResourceUsage() external view returns (uint16 vcpuUsed, uint16 vmUsed, uint16 vcpuCap, uint16 vmCap);
}
