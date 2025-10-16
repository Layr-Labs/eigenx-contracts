// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/**
 * @title IBillingCore
 * @notice Interface for central billing system with types, events, and errors
 */
interface IBillingCore {
    // ============================================================================
    // Data Types
    // ============================================================================

    struct Account {
        int96 balance; // Positive = credit, Negative = debt
        uint96 totalSpent;
        uint40 lastActivity;
        bool suspended;
    }

    struct Product {
        string name;
        address billingModule;
        address revenueRecipient;
        uint96 unclaimedRevenue;
        bool isActive;
    }

    struct ProductCharges {
        uint8 productId;
        string productName;
        uint96 amount;
    }

    struct EffectiveBalance {
        int96 balance;
        uint96 outstandingCharges;
        int96 effectiveBalance;
        bool willCoverCharges;
        ProductOutstanding[] breakdown;
    }

    struct ProductOutstanding {
        uint8 productId;
        string productName;
        uint96 outstanding;
    }

    // ============================================================================
    // Events
    // ============================================================================

    event Deposit(address indexed account, uint96 amount);
    event Withdrawal(address indexed account, uint96 amount);
    event Charge(address indexed account, uint8 indexed productId, uint96 amount, uint40 period);
    event ProductRegistered(uint8 indexed productId, string name, address module);
    event AccountSuspended(address indexed account);
    event AccountResumed(address indexed account);
    event DebtIncurred(address indexed account, uint96 amount);
    event DebtPaid(address indexed account, uint96 amount);
    event RevenueWithdrawn(uint8 indexed productId, address indexed recipient, uint96 amount);

    // ============================================================================
    // Custom Errors
    // ============================================================================

    error InsufficientBalance();
    error AccountIsSuspended();
    error AlreadyRegistered();
    error UnknownProduct();
    error ProductInactive();
    error CannotChargeFuturePeriods();
    error InsufficientRevenue();
    error NoDebt();

    // ============================================================================
    // Functions
    // ============================================================================

    function deposit(uint96 amount) external;
    function withdraw(uint96 amount) external;
    function depositFor(address account, uint96 amount) external;
    function getCurrentPeriod() external view returns (uint40);
    function getPeriodForTimestamp(uint40 timestamp) external view returns (uint40);
    function getCurrentPeriodSpending(address account) external view returns (uint96 total);
    function getSpendingBreakdown(address account, uint40 period)
        external
        view
        returns (uint8[] memory ids, string[] memory productNames, uint96[] memory amounts);
    function registerProduct(string calldata name, address module) external returns (uint8);
    function charge(address account, uint96 amount) external returns (bool);
    function chargePeriod(address account, uint96 amount, uint40 period) external returns (bool);
    function withdrawRevenue(uint8 productId) external;
    function setRevenueRecipient(uint8 productId, address recipient) external;
    function genesisTime() external view returns (uint40);
    function getBalance(address account) external view returns (int96);
    function getChargesForPeriod(address account, uint40 period) external view returns (ProductCharges[] memory);
    function getEffectiveBalance(address account) external view returns (EffectiveBalance memory);
}
