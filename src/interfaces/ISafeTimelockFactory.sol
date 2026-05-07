// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * @title ISafeTimelockFactory
 * @notice Factory for deploying verified Gnosis Safes and TimelockControllers
 * @dev Deployed entities can be verified as "official" via isSafe() and isTimelock()
 */
interface ISafeTimelockFactory {
    /// STRUCTS
    /// @notice Configuration for deploying a Gnosis Safe
    struct SafeConfig {
        address[] owners; // Addresses that can sign transactions
        uint256 threshold; // Number of required signatures
    }

    /// @notice Configuration for deploying a TimelockController
    struct TimelockConfig {
        uint256 minDelay; // Minimum delay in seconds before execution
        address[] proposers; // Addresses that can propose and cancel operations
        address[] executors; // Addresses that can execute ready operations
    }

    /// EVENTS

    /// @notice Emitted when a Gnosis Safe is deployed
    event SafeDeployed(
        address indexed deployer, address indexed safe, address[] owners, uint256 threshold, bytes32 salt
    );

    /// @notice Emitted when a TimelockController is deployed
    event TimelockDeployed(
        address indexed deployer,
        address indexed timelock,
        uint256 minDelay,
        address[] proposers,
        address[] executors,
        bytes32 salt
    );

    /// ERRORS

    /// @notice Thrown when proposers array is empty
    error NoProposers();

    /// @notice Thrown when executors array is empty
    error NoExecutors();

    /// @notice Thrown when a proposer address is zero
    error ZeroAddressProposer();

    /// @notice Thrown when an executor address is zero
    error ZeroAddressExecutor();

    /// EXTERNAL FUNCTIONS

    /**
     * @notice Deploys a new Gnosis Safe with deterministic address
     * @param config Safe configuration (owners, threshold)
     * @param salt User-provided salt for deterministic deployment
     * @return safe The deployed Safe address
     */
    function deploySafe(SafeConfig calldata config, bytes32 salt) external returns (address safe);

    /**
     * @notice Deploys a new TimelockController with deterministic address
     * @param config Timelock configuration (minDelay, proposers, executors)
     * @param salt User-provided salt for deterministic deployment
     * @return timelock The deployed TimelockController address
     */
    function deployTimelock(TimelockConfig calldata config, bytes32 salt) external returns (address timelock);

    /// VIEW FUNCTIONS

    /**
     * @notice Checks if an address is a Safe deployed by this factory
     * @param safe The address to check
     * @return True if the address is a factory-deployed Safe
     */
    function isSafe(address safe) external view returns (bool);

    /**
     * @notice Checks if an address is a Timelock deployed by this factory
     * @param timelock The address to check
     * @return True if the address is a factory-deployed Timelock
     */
    function isTimelock(address timelock) external view returns (bool);

    /**
     * @notice Pre-computes the address of a Safe deployment
     * @param deployer The address that will deploy
     * @param config Safe configuration
     * @param salt User-provided salt
     * @return The computed Safe address
     */
    function calculateSafeAddress(address deployer, SafeConfig calldata config, bytes32 salt)
        external
        view
        returns (address);

    /**
     * @notice Pre-computes the address of a Timelock deployment
     * @param deployer The address that will deploy
     * @param salt User-provided salt
     * @return The computed TimelockController address
     */
    function calculateTimelockAddress(address deployer, bytes32 salt) external view returns (address);

    /**
     * @notice Returns the official Gnosis Safe singleton address
     * @return The Safe singleton (master copy) address
     */
    function safeSingleton() external view returns (address);

    /**
     * @notice Returns the official Gnosis Safe proxy factory address
     * @return The SafeProxyFactory address
     */
    function safeProxyFactory() external view returns (address);

    /**
     * @notice Returns the TimelockController implementation address for minimal proxies
     * @return The TimelockControllerImpl address
     */
    function timelockImplementation() external view returns (address);

    /**
     * @notice Returns all Timelocks deployed by a given deployer
     * @param deployer The deployer address
     * @return Array of Timelock addresses
     */
    function getTimelocksByDeployer(address deployer) external view returns (address[] memory);

    /**
     * @notice Returns all Safes deployed by a given deployer
     * @param deployer The deployer address
     * @return Array of Safe addresses
     */
    function getSafesByDeployer(address deployer) external view returns (address[] memory);
}
