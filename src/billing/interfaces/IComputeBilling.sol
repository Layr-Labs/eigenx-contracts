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

    event AppStarted(address indexed app, address indexed account, uint16 skuId, uint96 rate);
    event AppStopped(address indexed app, address indexed account);
    event AppSettled(address indexed app, uint96 amount, uint40 period);
    event AppStateChanged(address indexed app, bool isRunning, uint96 newRate);
    event AppTerminated(address indexed app);
    event SKURatesSet(uint16 indexed skuID, string name, uint96 runningRate, uint96 stoppedRate, uint16 vcpus);
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

    // ============================================================================
    // Functions
    // ============================================================================

    function settleAccount(address account) external;
    function getResourceUsage() external view returns (uint16 vcpuUsed, uint16 vmUsed, uint16 vcpuCap, uint16 vmCap);
    function getAppDetails(address app)
        external
        view
        returns (address account, uint16 skuId, uint96 currentRate, bool isRunning, bool isActive);
    function getAccountBreakdown(address account)
        external
        view
        returns (uint96 runningRate, uint96 stoppedRate, uint16 runningCount, uint16 stoppedCount);
}
