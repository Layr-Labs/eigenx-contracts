// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {
    TimelockControllerUpgradeable
} from "@openzeppelin-upgrades/contracts/governance/TimelockControllerUpgradeable.sol";

/**
 * @title TimelockControllerImpl
 * @notice Implementation contract for TimelockController minimal proxies
 * @dev Wraps TimelockControllerUpgradeable to expose a public initialize function
 */
contract TimelockControllerImpl is TimelockControllerUpgradeable {
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the timelock controller
     * @param minDelay Minimum delay for operations
     * @param proposers Addresses granted proposer and canceller roles
     * @param executors Addresses granted executor role
     * @param admin Optional admin address (use address(0) for self-administered)
     */
    function initialize(uint256 minDelay, address[] memory proposers, address[] memory executors, address admin)
        external
        initializer
    {
        __TimelockController_init(minDelay, proposers, executors, admin);
    }
}
