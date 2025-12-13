// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ISafeTimelockFactory} from "../interfaces/ISafeTimelockFactory.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @title SafeTimelockFactoryStorage
 * @notice Storage contract for SafeTimelockFactory
 */
abstract contract SafeTimelockFactoryStorage is ISafeTimelockFactory {
    using EnumerableSet for EnumerableSet.AddressSet;

    /// IMMUTABLES

    /// @notice The official Gnosis Safe singleton (master copy) address
    address public immutable safeSingleton;

    /// @notice The official Gnosis Safe proxy factory address
    address public immutable safeProxyFactory;

    /// @notice The default fallback handler for Safes
    address public immutable defaultFallbackHandler;

    /// @notice The TimelockController implementation for minimal proxies
    address public immutable timelockImplementation;

    /// STATE VARIABLES

    /// @notice Set of all Safes deployed by this factory
    EnumerableSet.AddressSet internal _safes;

    /// @notice Set of all Timelocks deployed by this factory
    EnumerableSet.AddressSet internal _timelocks;

    /// STORAGE GAP

    /// @notice Storage gap for future upgrades
    uint256[46] private __gap;

    /// CONSTRUCTOR

    /**
     * @param _safeSingleton The official Gnosis Safe singleton address
     * @param _safeProxyFactory The official Gnosis Safe proxy factory address
     * @param _defaultFallbackHandler The default fallback handler for Safes
     * @param _timelockImplementation The TimelockController implementation for minimal proxies
     */
    constructor(
        address _safeSingleton,
        address _safeProxyFactory,
        address _defaultFallbackHandler,
        address _timelockImplementation
    ) {
        safeSingleton = _safeSingleton;
        safeProxyFactory = _safeProxyFactory;
        defaultFallbackHandler = _defaultFallbackHandler;
        timelockImplementation = _timelockImplementation;
    }
}
