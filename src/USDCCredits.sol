// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {USDCCreditsStorage} from "./storage/USDCCreditsStorage.sol";
import {IUSDCCredits} from "./interfaces/IUSDCCredits.sol";

contract USDCCredits is Initializable, OwnableUpgradeable, USDCCreditsStorage {
    using SafeERC20 for IERC20;

    /**
     * @param _usdc The USDC token contract address
     * @param _treasury The treasury address that receives all deposits
     */
    constructor(IERC20 _usdc, address _treasury) USDCCreditsStorage(_usdc, _treasury) {
        _disableInitializers();
    }

    /// @inheritdoc IUSDCCredits
    function initialize(address initialOwner, uint256 _minimumDeposit) external initializer {
        _transferOwnership(initialOwner);
        minimumDeposit = _minimumDeposit;
        emit MinimumDepositSet(0, _minimumDeposit);
    }

    /// @inheritdoc IUSDCCredits
    function purchaseCredits(uint256 amount) external {
        purchaseCreditsFor(amount, msg.sender);
    }

    /// @inheritdoc IUSDCCredits
    function purchaseCreditsFor(uint256 amount, address account) public {
        require(account != address(0), ZeroAddress());
        require(amount >= minimumDeposit, BelowMinimumDeposit());

        usdc.safeTransferFrom(msg.sender, treasury, amount);

        emit Deposit(msg.sender, account, amount);
    }

    /// @inheritdoc IUSDCCredits
    function setMinimumDepositFor(uint256 newMinimum) external onlyOwner {
        uint256 oldMinimum = minimumDeposit;
        minimumDeposit = newMinimum;
        emit MinimumDepositSet(oldMinimum, newMinimum);
    }

    /// @inheritdoc IUSDCCredits
    function sweep(IERC20 token) external onlyOwner {
        uint256 balance = token.balanceOf(address(this));
        if (balance > 0) {
            token.safeTransfer(treasury, balance);
        }
    }
}
