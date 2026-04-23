// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {Clones} from "@openzeppelin/contracts/proxy/Clones.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {SafeTimelockFactoryStorage} from "../storage/SafeTimelockFactoryStorage.sol";
import {ISafeTimelockFactory} from "../interfaces/ISafeTimelockFactory.sol";
import {ISafe} from "../interfaces/ISafe.sol";
import {ISafeProxyFactory} from "../interfaces/ISafeProxyFactory.sol";
import {TimelockControllerImpl} from "../governance/TimelockControllerImpl.sol";

/**
 * @title SafeTimelockFactory
 * @notice Factory for deploying verified Gnosis Safes and TimelockControllers
 * @dev Deployed entities can be verified as "official" via isSafe() and isTimelock()
 */
contract SafeTimelockFactory is Initializable, SafeTimelockFactoryStorage {
    using EnumerableSet for EnumerableSet.AddressSet;

    constructor(
        address _safeSingleton,
        address _safeProxyFactory,
        address _safeFallbackHandler,
        address _timelockImplementation
    ) SafeTimelockFactoryStorage(_safeSingleton, _safeProxyFactory, _safeFallbackHandler, _timelockImplementation) {
        _disableInitializers();
    }

    function initialize() external initializer {}

    /// EXTERNAL FUNCTIONS

    /// @inheritdoc ISafeTimelockFactory
    function deploySafe(SafeConfig calldata config, bytes32 salt) external returns (address safe) {
        safe = _deploySafe(config, salt);
        _safes.add(safe);
        _safesByDeployer[msg.sender].add(safe);
        emit SafeDeployed(msg.sender, safe, config.owners, config.threshold, salt);
    }

    /// @inheritdoc ISafeTimelockFactory
    function deployTimelock(TimelockConfig calldata config, bytes32 salt) external returns (address timelock) {
        _validateTimelockConfig(config);
        timelock = _deployTimelock(config, salt);
        _timelocks.add(timelock);
        _timelocksByDeployer[msg.sender].add(timelock);
        emit TimelockDeployed(msg.sender, timelock, config.minDelay, config.proposers, config.executors, salt);
    }

    /// VIEW FUNCTIONS

    /// @inheritdoc ISafeTimelockFactory
    function isSafe(address safe) external view returns (bool) {
        return _safes.contains(safe);
    }

    /// @inheritdoc ISafeTimelockFactory
    function isTimelock(address timelock) external view returns (bool) {
        return _timelocks.contains(timelock);
    }

    /// @inheritdoc ISafeTimelockFactory
    function getTimelocksByDeployer(address deployer) external view returns (address[] memory) {
        return _timelocksByDeployer[deployer].values();
    }

    /// @inheritdoc ISafeTimelockFactory
    function getSafesByDeployer(address deployer) external view returns (address[] memory) {
        return _safesByDeployer[deployer].values();
    }

    /// @inheritdoc ISafeTimelockFactory
    function calculateSafeAddress(address deployer, SafeConfig calldata config, bytes32 salt)
        external
        view
        returns (address)
    {
        bytes memory initializer = _encodeSafeInitializer(config);
        uint256 saltNonce = uint256(_deriveSalt(deployer, salt));
        bytes32 creationSalt = keccak256(abi.encodePacked(keccak256(initializer), saltNonce));
        bytes memory proxyCreationCode = ISafeProxyFactory(safeProxyFactory).proxyCreationCode();
        bytes memory deploymentData = abi.encodePacked(proxyCreationCode, uint256(uint160(safeSingleton)));
        return Create2.computeAddress(creationSalt, keccak256(deploymentData), safeProxyFactory);
    }

    /// @inheritdoc ISafeTimelockFactory
    function calculateTimelockAddress(address deployer, bytes32 salt) external view returns (address) {
        bytes32 mixedSalt = _deriveSalt(deployer, salt);
        return Clones.predictDeterministicAddress(timelockImplementation, mixedSalt);
    }

    /// INTERNAL FUNCTIONS

    function _validateTimelockConfig(TimelockConfig calldata config) internal pure {
        require(config.proposers.length > 0, NoProposers());
        require(config.executors.length > 0, NoExecutors());
        for (uint256 i = 0; i < config.proposers.length; i++) {
            require(config.proposers[i] != address(0), ZeroAddressProposer());
        }
        for (uint256 i = 0; i < config.executors.length; i++) {
            require(config.executors[i] != address(0), ZeroAddressExecutor());
        }
    }

    function _deploySafe(SafeConfig calldata config, bytes32 salt) internal returns (address safe) {
        bytes memory initializer = _encodeSafeInitializer(config);
        uint256 saltNonce = uint256(_deriveSalt(msg.sender, salt));
        safe = ISafeProxyFactory(safeProxyFactory).createProxyWithNonce(safeSingleton, initializer, saltNonce);
    }

    function _encodeSafeInitializer(SafeConfig calldata config) internal view returns (bytes memory) {
        return abi.encodeWithSelector(
            ISafe.setup.selector,
            config.owners,
            config.threshold,
            address(0),
            "",
            safeFallbackHandler,
            address(0),
            0,
            payable(address(0))
        );
    }

    function _deployTimelock(TimelockConfig calldata config, bytes32 salt) internal returns (address timelock) {
        bytes32 mixedSalt = _deriveSalt(msg.sender, salt);
        timelock = Clones.cloneDeterministic(timelockImplementation, mixedSalt);

        // forgefmt: disable-next-item
        TimelockControllerImpl(payable(timelock)).initialize(config.minDelay, config.proposers, config.executors, address(0));
    }

    function _deriveSalt(address deployer, bytes32 salt) internal view returns (bytes32) {
        return keccak256(abi.encodePacked(address(this), deployer, salt));
    }
}
