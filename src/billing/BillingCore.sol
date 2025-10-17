// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {IBillingModule} from "./interfaces/IBillingModule.sol";
import {IBillingCore} from "./interfaces/IBillingCore.sol";
import {DateTimeLib} from "solady/utils/DateTimeLib.sol";

/**
 * @title BillingCore
 * @notice Central billing system for multi-product usage tracking and payment
 * @dev All products use the same settlement pattern: accumulate then settle by period
 */
contract BillingCore is Initializable, OwnableUpgradeable, IBillingCore {
    using SafeERC20 for IERC20;

    // ============================================================================
    // State
    // ============================================================================

    IERC20 public immutable paymentToken;
    uint40 public immutable genesisTime;

    mapping(address => Account) public accounts;
    mapping(uint8 => Product) public products;
    mapping(address => uint8) public moduleToProductId;

    uint8 public nextProductId;

    // ============================================================================
    // Constructor & Initialization
    // ============================================================================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor(address _token) {
        _disableInitializers();

        paymentToken = IERC20(_token);
        genesisTime = uint40(block.timestamp);
    }

    function initialize(address _owner) external initializer {
        _transferOwnership(_owner);
        nextProductId = 1; // Start at 1, 0 means unknown
    }

    // ============================================================================
    // Period Management
    // ============================================================================

    /**
     * @notice Get current period (months since genesis)
     */
    function getCurrentPeriod() public view returns (uint40) {
        return uint40(DateTimeLib.diffMonths(genesisTime, block.timestamp));
    }

    /**
     * @notice Get period for a specific timestamp
     */
    function getPeriodForTimestamp(uint40 timestamp) public view returns (uint40) {
        if (timestamp < genesisTime) return 0;
        return uint40(DateTimeLib.diffMonths(genesisTime, timestamp));
    }

    /**
     * @notice Get start and end timestamps for a period
     */
    function getPeriodBounds(uint40 period) public view returns (uint40 start, uint40 end) {
        start = uint40(DateTimeLib.addMonths(genesisTime, period));
        end = uint40(DateTimeLib.addMonths(genesisTime, period + 1));
    }

    // ============================================================================
    // Account Management
    // ============================================================================

    /**
     * @notice Deposit tokens to your account
     */
    function deposit(uint96 amount) external {
        require(amount != 0, ZeroAmount());
        _deposit(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Deposit tokens to another account
     */
    function depositFor(address account, uint96 amount) external {
        require(amount != 0, ZeroAmount());
        require(account != address(0), InvalidAddress());
        _deposit(msg.sender, account, amount);
    }

    /**
     * @notice Withdraw tokens from your account
     */
    function withdraw(uint96 amount) external {
        require(amount != 0, ZeroAmount());

        Account storage acc = accounts[msg.sender];
        require(!acc.suspended, AccountIsSuspended());
        require(acc.balance >= int96(amount), InsufficientBalance());

        acc.balance -= int96(amount);

        paymentToken.safeTransfer(msg.sender, amount);
        emit Withdrawal(msg.sender, amount);
    }

    // ============================================================================
    // Product Management
    // ============================================================================

    /**
     * @notice Register a new product/service
     */
    function registerProduct(string calldata name, address module) external onlyOwner returns (uint8 productId) {
        require(module != address(0), InvalidAddress());
        require(moduleToProductId[module] == 0, ModuleAlreadyRegistered());

        productId = nextProductId++;

        products[productId] = Product({
            name: name,
            module: module,
            revenueRecipient: address(0),
            unclaimedRevenue: 0,
            active: true
        });

        moduleToProductId[module] = productId;

        emit ProductRegistered(productId, name, module);
    }

    /**
     * @notice Set revenue recipient for a product
     */
    function setRevenueRecipient(uint8 productId, address recipient) external onlyOwner {
        require(productId != 0 && productId < nextProductId, UnknownProduct());
        products[productId].revenueRecipient = recipient;
    }

    /**
     * @notice Withdraw revenue for a product
     */
    function withdrawRevenue(uint8 productId) external {
        require(productId != 0 && productId < nextProductId, UnknownProduct());

        Product storage product = products[productId];
        require(product.revenueRecipient != address(0), NoRevenueRecipient());

        uint96 amount = product.unclaimedRevenue;
        require(amount != 0, ZeroAmount());

        product.unclaimedRevenue = 0;
        paymentToken.safeTransfer(product.revenueRecipient, amount);
    }

    /**
     * @notice Deactivate a product
     */
    function deactivateProduct(uint8 productId) external onlyOwner {
        products[productId].active = false;
    }

    // ============================================================================
    // Charging
    // ============================================================================

    /**
     * @notice Charge an account for a specific period
     * @dev Called by billing modules during settlement
     * @param account Account to charge
     * @param amount Amount to charge
     * @param period Period this charge belongs to
     * @return success Whether account had sufficient balance
     */
    function chargePeriod(address account, uint96 amount, uint40 period) external returns (bool success) {
        if (amount == 0) return true; // No-op for zero charges

        uint8 productId = moduleToProductId[msg.sender];
        require(productId != 0, UnknownProduct());

        // Periods must be in the past (already completed)
        require(period < getCurrentPeriod(), InvalidPeriod());

        Account storage acc = accounts[account];

        // Always charge the full amount (allow negative balance)
        acc.balance -= int96(amount);
        acc.totalSpent += amount;

        // Track revenue
        products[productId].unclaimedRevenue += amount;

        // Update suspension status
        bool wasSuspended = acc.suspended;
        acc.suspended = acc.balance < 0;

        // Emit suspension event if status changed
        if (!wasSuspended && acc.suspended) {
            emit AccountSuspended(account);
        }

        // Emit charge event with all relevant data
        emit PeriodCharge(
            account,
            productId,
            period,
            amount,
            acc.balance,
            acc.suspended
        );

        return acc.balance >= 0;
    }

    // ============================================================================
    // View Functions
    // ============================================================================

    /**
     * @notice Get account balance
     */
    function getBalance(address account) external view returns (int96) {
        return accounts[account].balance;
    }

    /**
     * @notice Check if account is suspended
     */
    function isAccountSuspended(address account) external view returns (bool) {
        return accounts[account].suspended;
    }

    /**
     * @notice Get account details
     */
    function getAccount(address account) external view returns (
        int96 balance,
        uint96 totalSpent,
        bool suspended
    ) {
        Account memory acc = accounts[account];
        return (acc.balance, acc.totalSpent, acc.suspended);
    }

    /**
     * @notice Get product details
     */
    function getProduct(uint8 productId) external view returns (
        string memory name,
        address module,
        address revenueRecipient,
        uint96 unclaimedRevenue,
        bool active
    ) {
        Product memory product = products[productId];
        return (product.name, product.module, product.revenueRecipient, product.unclaimedRevenue, product.active);
    }

    /**
     * @notice Get all charges for a specific period (settled and pending)
     * @param account The account to query
     * @param period The period to query
     * @return charges Array of charges
     */
    function getChargesForPeriod(address account, uint40 period) external view returns (ProductCharges[] memory charges) {
        uint8 productCount = nextProductId - 1;
        charges = new ProductCharges[](productCount);

        for (uint8 i = 0; i < productCount; i++) {
            uint8 id = i + 1;
            Product memory product = products[id];

            if (product.module != address(0) && product.active) {
                (uint96 amount, bool isSettled) = IBillingModule(product.module).getChargesForPeriod(account, period);
                charges[i] = ProductCharges({
                    productId: id,
                    productName: product.name,
                    amount: amount,
                    settled: isSettled
                });
            }
        }
    }

    // ============================================================================
    // Internal Functions
    // ============================================================================

    function _deposit(address from, address to, uint96 amount) internal {
        Account storage acc = accounts[to];

        acc.balance += int96(amount);

        // Auto-resume if balance becomes positive
        if (acc.suspended && acc.balance >= 0) {
            acc.suspended = false;
            emit AccountResumed(to);
        }

        paymentToken.safeTransferFrom(from, address(this), amount);
        emit Deposit(to, amount);
    }

    // ============================================================================
    // Storage Gap
    // ============================================================================

    uint256[50] private __gap;
}