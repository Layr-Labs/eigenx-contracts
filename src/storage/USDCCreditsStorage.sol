// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IUSDCCredits} from "../interfaces/IUSDCCredits.sol";

abstract contract USDCCreditsStorage is IUSDCCredits {
    /// CONSTANTS & IMMUTABLES
    /// @notice The USDC token contract
    IERC20 public immutable usdc;

    /// @notice The treasury address that receives all deposits
    address public immutable treasury;

    /// @notice The minimum deposit amount (in USDC's 6-decimal base units)
    uint256 public minimumDeposit;

    constructor(IERC20 _usdc, address _treasury) {
        require(address(_usdc) != address(0), ZeroAddress());
        require(_treasury != address(0), ZeroAddress());
        usdc = _usdc;
        treasury = _treasury;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
