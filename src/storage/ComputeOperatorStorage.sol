// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IDelegationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IComputeOperator} from "../interfaces/IComputeOperator.sol";

abstract contract ComputeOperatorStorage is IComputeOperator {
    /// CONSTANTS & IMMUTABLES
    /// @notice The EigenLayer DelegationManager contract
    IDelegationManager public immutable delegationManager;

    /// @notice The EigenLayer AllocationManager contract
    IAllocationManager public immutable allocationManager;

    /// @notice The EigenLayer PermissionController contract
    IPermissionController public immutable permissionController;

    /// @notice The AppController contract
    address public immutable appController;

    /// @notice The ComputeAVSRegistrar contract
    address public immutable computeAVSRegistrar;

    constructor(
        IDelegationManager _delegationManager,
        IAllocationManager _allocationManager,
        IPermissionController _permissionController,
        address _appController,
        address _computeAVSRegistrar
    ) {
        delegationManager = _delegationManager;
        allocationManager = _allocationManager;
        permissionController = _permissionController;
        appController = _appController;
        computeAVSRegistrar = _computeAVSRegistrar;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
