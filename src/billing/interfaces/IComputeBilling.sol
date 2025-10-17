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
        uint96 runningRate; // Tokens per second when running
        uint96 stoppedRate; // Tokens per second when stopped
        uint16 vcpus; // vCPUs required
        uint96 minimumDeposit; // Minimum deposit to start
        string name;
        bool active;
    }

    struct AccountState {
        uint16 totalVMCount; // Total number of VMs (running + stopped)
        uint96 totalRunningRate; // Total rate for running apps
        uint96 totalStoppedRate; // Total rate for stopped apps
        uint40 lastUpdate; // Last time charges were accrued
        uint96 accruedCharges; // Charges waiting to be settled
        uint40 chargesStartTime; // When current accrued charges started accumulating
    }

    struct SKUAppCounts {
        uint16 runningCount; // Number of running apps with this SKU
        uint16 stoppedCount; // Number of stopped apps with this SKU
    }

    struct Resources {
        uint16 vcpuUsed;
        uint16 vcpuCap;
        uint16 vmInstancesUsed;
        uint16 vmInstanceCap;
    }

    struct PeriodSettlement {
        bool settled;
        uint96 amount;
        uint40 settledAt;
    }

    // ============================================================================
    // Events
    // ============================================================================

    event BillingStarted(address indexed app, address indexed account, uint96 rate);
    event BillingStopped(address indexed app, address indexed account);
    event RateChanged(address indexed app, address indexed account, uint96 oldRate, uint96 newRate);
    event PeriodSettled(address indexed account, uint40 indexed period, uint96 amount);
    event SKUUpdated(uint16 indexed skuId, string name, uint96 rate);
    event SettlerAuthorizationChanged(address indexed settler, bool authorized);
    event ResourceCapUpdated(uint16 vcpuCap, uint16 vmInstanceCap);
    event SKUChangeApplied(uint16 indexed skuId, uint40 indexed effectivePeriod);

    // ============================================================================
    // Custom Errors
    // ============================================================================

    error InvalidSKU();
    error InvalidPeriod();
    error PeriodAlreadySettled();
    error UnauthorizedSettler();
    error InvalidAccount();
    error InsufficientDeposit(uint96 required, int96 available);
    error VCPUCapacityExceeded();
    error InstanceCapacityExceeded();

    // ============================================================================
    // Functions
    // ============================================================================

    // Charge Accrual
    function accrueCharges(address account) external;

    // Period Settlement
    function settlePeriod(address account, uint40 period) external;
    function settlePeriodBatch(address[] calldata accounts, uint40 period) external;

    // View Functions
    function getOutstandingCharges(address account) external view returns (uint96);
    function isPeriodSettled(address account, uint40 period) external view returns (bool);
    function getSettlement(address account, uint40 period)
        external
        view
        returns (bool settled, uint96 amount, uint40 settledAt);
    function estimateCurrentPeriodCharges(address account) external view returns (uint96);
    function getResourceUsage()
        external
        view
        returns (uint16 vcpuUsed, uint16 vcpuCap, uint16 vmInstancesUsed, uint16 vmInstanceCap);
}
