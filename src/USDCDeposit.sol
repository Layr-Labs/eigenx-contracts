// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {USDCDepositStorage} from "./storage/USDCDepositStorage.sol";
import {IUSDCDeposit} from "./interfaces/IUSDCDeposit.sol";

contract USDCDeposit is Initializable, OwnableUpgradeable, USDCDepositStorage {
    using SafeERC20 for IERC20;

    /**
     * @param _usdc The USDC token contract address
     * @param _treasury The treasury address that receives all deposits
     */
    constructor(IERC20 _usdc, address _treasury) USDCDepositStorage(_usdc, _treasury) {
        _disableInitializers();
    }

    /// @inheritdoc IUSDCDeposit
    function initialize(address initialOwner, uint256 _minimumDeposit) external initializer {
        _transferOwnership(initialOwner);
        minimumDeposit = _minimumDeposit;
        emit MinimumDepositSet(0, _minimumDeposit);
    }

    /// @inheritdoc IUSDCDeposit
    function deposit(uint256 amount) external {
        depositFor(amount, msg.sender);
    }

    /// @inheritdoc IUSDCDeposit
    function depositFor(uint256 amount, address account) public {
        require(account != address(0), ZeroAddress());
        require(amount >= minimumDeposit, BelowMinimumDeposit());

        usdc.safeTransferFrom(msg.sender, treasury, amount);

        emit Deposit(msg.sender, account, amount);
    }

    /// @inheritdoc IUSDCDeposit
    function setMinimumDeposit(uint256 newMinimum) external onlyOwner {
        uint256 oldMinimum = minimumDeposit;
        minimumDeposit = newMinimum;
        emit MinimumDepositSet(oldMinimum, newMinimum);
    }

    /// @inheritdoc IUSDCDeposit
    function sweep(IERC20 token) external onlyOwner {
        uint256 balance = token.balanceOf(address(this));
        if (balance > 0) {
            token.safeTransfer(treasury, balance);
        }
    }
}
