// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IBillingCore {
    struct Account {
        uint128 balance;
        uint128 totalBurnRate;
        uint40 lastSettled;
    }

    struct Product {
        address implementation;
        uint128 earned;
    }

    // Errors
    error ProductAlreadyAuthorized(address implementation);
    error ProductNotAuthorized();
    error LengthMismatch();
    error NoDebt(address account);
    error InsufficientEarnings(uint128 available, uint128 requested);

    // Events
    event Deposited(address indexed user, uint128 amount);
    event AccountCharged(address indexed account, uint128 amount);
    event AccountDebtIncurred(address indexed account, uint128 debt);
    event AccountSuspended(address indexed account);
    event AccountResumed(address indexed account);
    event AccountRateAdjusted(address indexed account, uint128 newTotalRate);
    event ProductAuthorized(uint8 indexed id, address implementation);
    event ProductWithdrawal(uint8 indexed productId, uint128 amount);

    // Streaming pattern - continuous burn rate
    function adjustAccountRate(address account, uint128 delta, bool increase) external;
    function settleAccount(address account) external returns (uint128 charged);

    // Consumption pattern - one-time charges
    function chargeAmount(address account, uint128 amount) external returns (bool success);

    // Balance management
    function deposit(uint128 amount) external;
    function depositFor(address user, uint128 amount) external;

    // Views
    function checkAccountSolvency(address account) external view returns (bool solvent, uint256 totalOwed, uint256 balance);
}
