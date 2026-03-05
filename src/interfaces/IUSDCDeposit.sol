// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IUSDCDeposit {
    /// @notice Thrown when deposit amount is below the minimum
    error BelowMinimumDeposit();

    /// @notice Thrown when a zero address is provided
    error ZeroAddress();

    /// @notice Emitted when a deposit is made
    event Deposit(address indexed depositor, address indexed account, uint256 amount);

    /// @notice Emitted when the minimum deposit amount is updated
    event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum);

    /**
     * @notice Initialize the USDCDeposit contract
     * @param initialOwner The owner address for the contract
     * @param _minimumDeposit The initial minimum deposit amount (in USDC's 6-decimal base units)
     */
    function initialize(address initialOwner, uint256 _minimumDeposit) external;

    /**
     * @notice Deposits USDC for msg.sender
     * @param amount The amount of USDC to deposit (in 6-decimal base units)
     * @dev Caller must have approved this contract to spend at least `amount` USDC
     */
    function deposit(uint256 amount) external;

    /**
     * @notice Deposits USDC on behalf of any address
     * @param amount The amount of USDC to deposit (in 6-decimal base units)
     * @param account The address to credit the deposit to
     * @dev Caller must have approved this contract to spend at least `amount` USDC
     * @dev Enables operators to fund agent accounts before the agent is deployed
     */
    function depositFor(uint256 amount, address account) external;

    /**
     * @notice Sets the minimum deposit amount
     * @param newMinimum The new minimum deposit amount (in USDC's 6-decimal base units)
     * @dev Caller must be the contract owner
     */
    function setMinimumDeposit(uint256 newMinimum) external;

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
     * @notice Returns the current minimum deposit amount
     */
    function minimumDeposit() external view returns (uint256);
}
