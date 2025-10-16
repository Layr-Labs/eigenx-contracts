// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {DateTimeLib} from "solady/utils/DateTimeLib.sol";
import {IBillingCore} from "./interfaces/IBillingCore.sol";
import {IBillingModule} from "./interfaces/IBillingModule.sol";

/**
 * @title BillingCore
 * @notice Central billing system for multi-product usage tracking and payment
 */
contract BillingCore is Initializable, OwnableUpgradeable, IBillingCore {
    using SafeERC20 for IERC20;

    // ============================================================================
    // Immutables & State Variables
    // ============================================================================

    IERC20 public immutable paymentToken;
    uint40 public immutable genesisTime;

    mapping(address account => Account) public accounts;
    mapping(uint8 productId => Product) public products;
    mapping(address module => uint8) public productIds;
    mapping(address account => mapping(uint8 productId => mapping(uint40 period => uint96))) public spending;

    uint8 public nextProductId;

    // ============================================================================
    // Constructor & Initializer
    // ============================================================================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor(address _token) {
        paymentToken = IERC20(_token);
        genesisTime = uint40(block.timestamp);
        _disableInitializers();
    }

    /**
     * @notice Initialize the BillingCore contract
     * @param _owner The owner address for the contract
     */
    function initialize(address _owner) external initializer {
        _transferOwnership(_owner);
        nextProductId = 1;
    }

    // ============================================================================
    // Period Management
    // ============================================================================

    /**
     * @notice Get current period based on calendar months elapsed since genesis
     */
    function getCurrentPeriod() public view returns (uint40) {
        return getPeriodForTimestamp(uint40(block.timestamp));
    }

    /**
     * @notice Get the period for a specific timestamp
     * @param timestamp The timestamp to calculate the period for
     * @return The period number (months since genesis)
     */
    function getPeriodForTimestamp(uint40 timestamp) public view returns (uint40) {
        if (timestamp < genesisTime) return 0;
        return uint40(DateTimeLib.diffMonths(genesisTime, uint256(timestamp)));
    }

    // ============================================================================
    // Customer Functions
    // ============================================================================

    /**
     * @notice Deposit tokens to account
     */
    function deposit(uint96 amount) external {
        _deposit(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Deposit tokens for another account
     */
    function depositFor(address account, uint96 amount) external {
        _deposit(msg.sender, account, amount);
    }

    /**
     * @notice Internal deposit logic
     */
    function _deposit(address from, address to, uint96 amount) internal {
        Account storage account = accounts[to];

        int96 previousBalance = account.balance;
        account.balance += int96(amount);
        account.lastActivity = uint40(block.timestamp);

        // Auto-unsuspend if debt is paid off
        if (previousBalance < 0 && account.balance >= 0 && account.suspended) {
            account.suspended = false;
            emit AccountResumed(to);
        }

        paymentToken.safeTransferFrom(from, address(this), amount);
        emit Deposit(to, amount);
    }

    /**
     * @notice Withdraw tokens from account
     */
    function withdraw(uint96 amount) external {
        Account storage account = accounts[msg.sender];

        require(!account.suspended, AccountIsSuspended());
        require(account.balance >= int96(amount), InsufficientBalance());

        account.balance -= int96(amount);
        account.lastActivity = uint40(block.timestamp);

        paymentToken.safeTransfer(msg.sender, amount);
        emit Withdrawal(msg.sender, amount);
    }

    // ============================================================================
    // Spending Visibility
    // ============================================================================

    /**
     * @notice Get total spending across all products for current period
     */
    function getCurrentPeriodSpending(address account) external view returns (uint96 total) {
        uint40 period = getCurrentPeriod();
        for (uint8 i = 1; i < nextProductId; i++) {
            total += spending[account][i][period];
        }
    }

    /**
     * @notice Get breakdown of spending by product for a period
     */
    function getSpendingBreakdown(address account, uint40 period)
        external
        view
        returns (uint8[] memory productIds_, string[] memory productNames, uint96[] memory amounts)
    {
        uint8 productCount = nextProductId - 1;
        productIds_ = new uint8[](productCount);
        productNames = new string[](productCount);
        amounts = new uint96[](productCount);

        for (uint8 i = 0; i < productCount; i++) {
            uint8 id = i + 1;
            productIds_[i] = id;
            productNames[i] = products[id].name;
            amounts[i] = spending[account][id][period];
        }
    }

    // ============================================================================
    // Product Functions
    // ============================================================================

    /**
     * @notice Register a new product/service
     */
    function registerProduct(string calldata name, address module) external onlyOwner returns (uint8) {
        require(productIds[module] == 0, AlreadyRegistered());

        uint8 id = nextProductId++;
        products[id] = Product({
            name: name,
            billingModule: module,
            revenueRecipient: address(0),
            unclaimedRevenue: 0,
            isActive: true
        });
        productIds[module] = id;

        emit ProductRegistered(id, name, module);
        return id;
    }

    /**
     * @notice Set the revenue recipient address for a product
     * @param productId The product ID
     * @param recipient The address that will receive withdrawn revenue
     */
    function setRevenueRecipient(uint8 productId, address recipient) external onlyOwner {
        require(productId > 0 && productId < nextProductId, UnknownProduct());
        products[productId].revenueRecipient = recipient;
    }

    /**
     * @notice Charge an account for product usage (attributes to current period)
     * @param account The account to charge
     * @param amount The amount to charge
     * @return fullyPaid Whether the full amount was paid
     */
    function charge(address account, uint96 amount) external returns (bool fullyPaid) {
        return chargePeriod(account, amount, getCurrentPeriod());
    }

    /**
     * @notice Charge an account for product usage in a specific period
     * @param account The account to charge
     * @param amount The amount to charge
     * @param period The period to attribute the charge to
     * @return fullyPaid Whether the full amount was paid
     */
    function chargePeriod(address account, uint96 amount, uint40 period) public returns (bool fullyPaid) {
        uint8 productId = productIds[msg.sender];
        require(productId != 0, UnknownProduct());
        require(products[productId].isActive, ProductInactive());
        require(period <= getCurrentPeriod(), CannotChargeFuturePeriods());

        Account storage acc = accounts[account];

        // Calculate what can actually be paid
        uint96 actualPaid;
        int96 newBalance = acc.balance - int96(amount);

        if (acc.balance >= int96(amount)) {
            // Full payment
            actualPaid = amount;
            fullyPaid = true;
        } else if (acc.balance > 0) {
            // Partial payment (use remaining balance)
            actualPaid = uint96(acc.balance);
            fullyPaid = false;
        } else {
            // No payment (already in debt)
            actualPaid = 0;
            fullyPaid = false;
        }

        // Update account
        acc.balance = newBalance;
        acc.lastActivity = uint40(block.timestamp);

        // Track actual payment
        if (actualPaid > 0) {
            acc.totalSpent += actualPaid;
            products[productId].unclaimedRevenue += actualPaid;
            spending[account][productId][period] += actualPaid;
            emit Charge(account, productId, actualPaid, period);
        }

        // Handle suspension
        if (!fullyPaid && !acc.suspended) {
            acc.suspended = true;
            emit AccountSuspended(account);
        }

        // Emit debt event if debt increased
        if (!fullyPaid) {
            uint96 debtIncurred = amount - actualPaid;
            emit DebtIncurred(account, debtIncurred);
        }

        return fullyPaid;
    }

    /**
     * @notice Get charges for an account across all products for a specific period
     * @param account The account to query
     * @param period The period to query
     * @return charges Array of charges broken down by product
     */
    function getChargesForPeriod(address account, uint40 period)
        external
        view
        returns (ProductCharges[] memory charges)
    {
        uint8 productCount = nextProductId - 1;
        charges = new ProductCharges[](productCount);

        for (uint8 i = 0; i < productCount; i++) {
            uint8 id = i + 1;

            Product storage product = products[id];
            charges[i] = ProductCharges({
                productId: id,
                productName: product.name,
                amount: IBillingModule(product.billingModule).getChargesForPeriod(account, period)
            });
        }
    }

    /**
     * @notice Withdraw all revenue for a product to its configured recipient
     * @dev Anyone can call this function to push revenue to the recipient
     * @param productId The product ID to withdraw revenue for
     */
    function withdrawRevenue(uint8 productId) external {
        require(productId > 0 && productId < nextProductId, UnknownProduct());

        Product storage product = products[productId];
        require(product.revenueRecipient != address(0), "No revenue recipient set");
        require(product.unclaimedRevenue > 0, InsufficientRevenue());

        uint96 amount = product.unclaimedRevenue;
        product.unclaimedRevenue = 0;

        paymentToken.safeTransfer(product.revenueRecipient, amount);
        emit RevenueWithdrawn(productId, product.revenueRecipient, amount);
    }

    // ============================================================================
    // View Functions
    // ============================================================================

    /**
     * @notice Get account's actual balance (can be negative if the account is in debt)
     */
    function getBalance(address account) external view returns (int96) {
        return accounts[account].balance;
    }

    /**
     * @notice Get effective balance after accounting for outstanding charges
     * @param account The account to query
     * @return result Comprehensive balance data including outstanding charges breakdown
     */
    function getEffectiveBalance(address account) external view returns (EffectiveBalance memory result) {
        result.balance = accounts[account].balance;

        uint8 productCount = nextProductId - 1;
        result.breakdown = new ProductOutstanding[](productCount);

        // Query each product for outstanding charges
        for (uint8 i = 0; i < productCount; i++) {
            uint8 id = i + 1;
            Product storage product = products[id];

            uint96 outstanding = 0;
            if (product.billingModule != address(0)) {
                outstanding = IBillingModule(product.billingModule).getOutstandingCharges(account);
            }

            result.breakdown[i] = ProductOutstanding({productId: id, productName: product.name, outstanding: outstanding});

            result.outstandingCharges += outstanding;
        }

        // Calculate effective balance
        result.effectiveBalance = result.balance - int96(result.outstandingCharges);
        result.willCoverCharges = result.effectiveBalance >= 0;
    }

    // ============================================================================
    // Storage Gap
    // ============================================================================

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
