// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/**
 * @title IUsageBilling
 * @notice Interface for usage-based billing system with types, events, and errors
 */
interface IUsageBilling {
    // ============================================================================
    // Data Types
    // ============================================================================

    struct PeriodUsage {
        uint96 amount; // Total usage for the period
        bool settled; // Whether charges have been sent to BillingCore
        uint40 settledAt; // When it was settled
    }

    // ============================================================================
    // Events
    // ============================================================================

    event UsageRecorded(address indexed account, uint96 amount, uint40 indexed period);
    event PeriodSettled(address indexed account, uint40 indexed period, uint96 amount);
    event UsageReporterSet(address indexed reporter);

    // ============================================================================
    // Custom Errors
    // ============================================================================

    error UnauthorizedReporter();
    error InvalidPeriod();
    error PeriodAlreadySettled();
    error InvalidAmount();

    // ============================================================================
    // Functions
    // ============================================================================

    // Usage Recording
    function recordUsage(address account, uint96 amount) external;
    function recordUsageForPeriod(address account, uint96 amount, uint40 period) external;
    function recordUsageBatch(address[] calldata accounts, uint96[] calldata amounts, uint40 period) external;

    // Period Settlement
    function settlePeriod(address account, uint40 period) external;
    function settlePeriodBatch(address[] calldata accounts, uint40 period) external;

    // View Functions
    function getUsageForPeriod(address account, uint40 period)
        external
        view
        returns (uint96 amount, bool settled, uint40 settledAt);
    function getUnsettledUsage(address account) external view returns (uint96 total);
    function isPeriodSettled(address account, uint40 period) external view returns (bool);
}
