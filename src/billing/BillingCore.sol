// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IBillingCore} from "./interfaces/IBillingCore.sol";

contract BillingCore is IBillingCore, Ownable {
    using SafeERC20 for IERC20;

    IERC20 public immutable token;

    mapping(uint8 productID => Product product) public products;
    mapping(address implementation => uint8 productID) public productIds;
    uint8 private _nextProductId = 1;

    mapping(address => Account) public accounts;
    mapping(address account => uint128 debt) public accountDebt;
    mapping(address account => bool suspended) public accountSuspended;
    
    constructor(address _token) Ownable() {
        token = IERC20(_token);
    }
    
    // Admin functions
    function authorizeProduct(address implementation) external onlyOwner returns (uint8) {
        require(productIds[implementation] == 0, ProductAlreadyAuthorized(implementation));

        uint8 id = _nextProductId++;

        products[id] = Product({
            implementation: implementation,
            earned: 0
        });
        productIds[implementation] = id;

        emit ProductAuthorized(id, implementation);
        return id;
    }

    // User functions
    function deposit(uint128 amount) external {
        accounts[msg.sender].balance += amount;
        token.safeTransferFrom(msg.sender, address(this), amount);
        emit Deposited(msg.sender, amount);
    }

    function depositFor(address user, uint128 amount) external {
        accounts[user].balance += amount;
        token.safeTransferFrom(msg.sender, address(this), amount);
        emit Deposited(user, amount);
    }

    // Batch operations
    function depositBatch(address[] calldata users, uint128[] calldata amounts) external {
        require(users.length == amounts.length, LengthMismatch());

        uint256 total = 0;
        for (uint256 i = 0; i < users.length; i++) {
            accounts[users[i]].balance += amounts[i];
            total += amounts[i];
            emit Deposited(users[i], amounts[i]);
        }

        token.safeTransferFrom(msg.sender, address(this), total);
    }

    // Product functions - Streaming pattern
    function adjustAccountRate(address account, uint128 delta, bool increase) external {
        uint8 productId = productIds[msg.sender];
        require(productId != 0, ProductNotAuthorized());

        // Settle before adjusting rate
        settleAccount(account);

        Account storage acc = accounts[account];

        // Initialize account if needed
        if (acc.lastSettled == 0) {
            acc.lastSettled = uint40(block.timestamp);
        }

        if (increase) {
            acc.totalBurnRate += delta;
        } else {
            acc.totalBurnRate -= delta;
        }

        emit AccountRateAdjusted(account, acc.totalBurnRate);
    }

    // Product functions - Consumption pattern
    function chargeAmount(address account, uint128 amount) external returns (bool success) {
        uint8 productId = productIds[msg.sender];
        require(productId != 0, ProductNotAuthorized());

        if (amount == 0) return true;

        Account storage acc = accounts[account];

        // Settle any streaming charges first
        if (acc.totalBurnRate > 0) {
            settleAccount(account);
        }

        if (acc.balance >= amount) {
            // Sufficient balance
            acc.balance -= amount;
            products[productId].earned += amount;
            emit AccountCharged(account, amount);
            return true;
        } else {
            // Insufficient balance - go into debt
            uint128 paid = acc.balance;
            uint128 debt = amount - paid;

            acc.balance = 0;
            accountDebt[account] += debt;
            products[productId].earned += paid;

            if (!accountSuspended[account]) {
                accountSuspended[account] = true;
                emit AccountSuspended(account);
            }

            if (paid > 0) {
                emit AccountCharged(account, paid);
            }
            emit AccountDebtIncurred(account, debt);
            return false;
        }
    }

    // Settle account's streaming charges
    function settleAccount(address account) public returns (uint128 charged) {
        Account storage acc = accounts[account];

        if (acc.totalBurnRate == 0 || acc.lastSettled == 0) return 0;

        uint256 elapsed = block.timestamp - acc.lastSettled;
        if (elapsed == 0) return 0;

        uint256 totalOwed = acc.totalBurnRate * elapsed;
        charged = uint128(totalOwed);

        if (acc.balance >= charged) {
            // Sufficient balance
            acc.balance -= charged;
            acc.lastSettled = uint40(block.timestamp);
            // Note: We don't track which product earned what for streaming
            // Products would need to implement their own tracking if needed
            emit AccountCharged(account, charged);
        } else {
            // Insufficient - go into debt
            uint128 paid = acc.balance;
            uint128 debt = charged - paid;

            acc.balance = 0;
            accountDebt[account] += debt;
            acc.lastSettled = uint40(block.timestamp);

            if (!accountSuspended[account]) {
                accountSuspended[account] = true;
                emit AccountSuspended(account);
            }

            if (paid > 0) {
                emit AccountCharged(account, paid);
            }
            emit AccountDebtIncurred(account, debt);
        }

        return charged;
    }

    // Pay off debt for an account
    function payDebt(address account, uint128 amount) external {
        uint128 currentDebt = accountDebt[account];
        require(currentDebt > 0, NoDebt(account));

        uint128 payment = amount > currentDebt ? currentDebt : amount;

        accountDebt[account] -= payment;
        token.safeTransferFrom(msg.sender, address(this), payment);

        if (accountDebt[account] == 0 && accountSuspended[account]) {
            accountSuspended[account] = false;
            emit AccountResumed(account);
        }
    }

    // View functions
    function checkAccountSolvency(address account) external view returns (
        bool solvent,
        uint256 totalOwed,
        uint256 balance
    ) {
        Account memory acc = accounts[account];
        balance = acc.balance;

        // Calculate pending charges since last settlement
        uint256 elapsed = block.timestamp - acc.lastSettled;
        uint256 pendingCharges = acc.totalBurnRate * elapsed;

        totalOwed = accountDebt[account] + pendingCharges;
        solvent = balance >= totalOwed;
    }

    // Product withdrawal
    function withdrawEarnings(uint128 amount) external {
        uint8 productId = productIds[msg.sender];
        require(productId != 0, ProductNotAuthorized());

        Product storage product = products[productId];
        require(product.earned >= amount, InsufficientEarnings(product.earned, amount));

        product.earned -= amount;
        token.safeTransfer(msg.sender, amount);
        emit ProductWithdrawal(productId, amount);
    }
}
