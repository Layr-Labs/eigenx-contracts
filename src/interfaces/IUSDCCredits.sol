// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IUSDCCredits {
    /// @notice Thrown when purchase amount is below the minimum
    error BelowMinimumPurchase();

    /// @notice Thrown when a zero address is provided
    error ZeroAddress();

    /// @notice Emitted when credits are purchased
    event CreditsPurchased(address indexed purchaser, address indexed account, uint256 amount);

    /// @notice Emitted when the minimum purchase amount is updated
    event MinimumPurchaseSet(uint256 oldMinimum, uint256 newMinimum);

    /**
     * @notice Initialize the USDCCredits contract
     * @param initialOwner The owner address for the contract
     * @param _minimumPurchase The initial minimum purchase amount (in USDC's 6-decimal base units)
     */
    function initialize(address initialOwner, uint256 _minimumPurchase) external;

    /**
     * @notice Purchases USDC credits for msg.sender
     * @param amount The amount of USDC to spend (in 6-decimal base units)
     * @dev Caller must have approved this contract to spend at least `amount` USDC
     */
    function purchaseCredits(uint256 amount) external;

    /**
     * @notice Purchases USDC credits on behalf of any address
     * @param amount The amount of USDC to spend (in 6-decimal base units)
     * @param account The address to credit the purchase to
     * @dev Caller must have approved this contract to spend at least `amount` USDC
     * @dev Enables operators to fund agent accounts before the agent is deployed
     */
    function purchaseCreditsFor(uint256 amount, address account) external;

    /**
     * @notice Sets the minimum purchase amount
     * @param newMinimum The new minimum purchase amount (in USDC's 6-decimal base units)
     * @dev Caller must be the contract owner
     */
    function setMinimumPurchaseFor(uint256 newMinimum) external;

    /**
     * @notice Recovers tokens accidentally sent directly to the contract
     * @param token The ERC20 token to sweep
     * @dev Transfers the contract's full balance of the token to the treasury
     * @dev Caller must be the contract owner
     */
    function sweep(IERC20 token) external;

    /**
     * @notice Returns the USDC token address
     */
    function usdc() external view returns (IERC20);

    /**
     * @notice Returns the treasury address
     */
    function treasury() external view returns (address);

    /**
     * @notice Returns the current minimum purchase amount
     */
    function minimumPurchase() external view returns (uint256);
}
