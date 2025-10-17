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
        int96 balance;          // Can go negative (debt)
        uint96 totalSpent;      // Lifetime spending tracker
        bool suspended;         // True when balance < 0
    }

    struct Product {
        string name;
        address module;         // Billing module address
        address revenueRecipient; // Who can withdraw revenue
        uint96 unclaimedRevenue;   // Unclaimed revenue for this product
        bool active;
    }

    struct ProductCharges {
        uint8 productId;
        string productName;
        uint96 amount;
        bool settled;
    }

    // ============================================================================
    // Events
    // ============================================================================

    event Deposit(address indexed account, uint96 amount);
    event Withdrawal(address indexed account, uint96 amount);
    event PeriodCharge(
        address indexed account,
        uint8 indexed productId,
        uint40 indexed period,
        uint96 amount,
        int96 newBalance,
        bool suspended
    );
    event ProductRegistered(uint8 indexed productId, string name, address module);
    event AccountSuspended(address indexed account);
    event AccountResumed(address indexed account);

    // ============================================================================
    // Custom Errors
    // ============================================================================

    error InsufficientBalance();
    error AccountIsSuspended();
    error UnknownProduct();
    error InvalidPeriod();
    error ZeroAmount();
    error InvalidAddress();
    error ModuleAlreadyRegistered();
    error ProductInactive();
    error NoRevenueRecipient();

    // ============================================================================
    // Functions
    // ============================================================================

    // Account Management
    function deposit(uint96 amount) external;
    function depositFor(address account, uint96 amount) external;
    function withdraw(uint96 amount) external;

    // Period Management
    function getCurrentPeriod() external view returns (uint40);
    function getPeriodForTimestamp(uint40 timestamp) external view returns (uint40);
    function getPeriodBounds(uint40 period) external view returns (uint40 start, uint40 end);

    // Product Management
    function registerProduct(string calldata name, address module) external returns (uint8);
    function setRevenueRecipient(uint8 productId, address recipient) external;
    function withdrawRevenue(uint8 productId) external;
    function deactivateProduct(uint8 productId) external;

    // Charging
    function chargePeriod(address account, uint96 amount, uint40 period) external returns (bool);

    // View Functions
    function genesisTime() external view returns (uint40);
    function getBalance(address account) external view returns (int96);
    function isAccountSuspended(address account) external view returns (bool);
    function getAccount(address account) external view returns (int96 balance, uint96 totalSpent, bool suspended);
    function getProduct(uint8 productId) external view returns (
        string memory name,
        address module,
        address revenueRecipient,
        uint96 unclaimedRevenue,
        bool active
    );
    function getChargesForPeriod(address account, uint40 period) external view returns (ProductCharges[] memory);
}
