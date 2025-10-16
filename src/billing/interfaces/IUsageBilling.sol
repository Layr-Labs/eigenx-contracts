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

    struct PeriodData {
        uint96 amount;
        bool settled;
    }

    // ============================================================================
    // Events
    // ============================================================================

    event UsageRecorded(address indexed account, uint96 amount, uint40 indexed period);
    event PeriodSettled(address indexed account, uint40 indexed period, uint96 amount);

    // ============================================================================
    // Custom Errors
    // ============================================================================

    error NotAuthorized();
    error LengthMismatch();
    error PeriodAlreadySettled();
    error CannotRecordFutureUsage();
    error CannotSettleCurrentOrFuturePeriods();

    // ============================================================================
    // Functions
    // ============================================================================

    function recordUsage(address account, uint96 amount, uint40 period) external;
    function recordBatch(address[] calldata accounts, uint96[] calldata amounts, uint40 period) external;
    function settlePeriod(address account, uint40 period, uint96 finalAmount) external;
    function settlePeriodBatch(address[] calldata accounts, uint40 period, uint96[] calldata amounts) external;
    function isPeriodSettled(address account, uint40 period) external view returns (bool);
}
