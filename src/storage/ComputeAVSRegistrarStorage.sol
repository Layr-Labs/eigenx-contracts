// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IAppController} from "../interfaces/IAppController.sol";
import {IComputeAVSRegistrar} from "../interfaces/IComputeAVSRegistrar.sol";

abstract contract ComputeAVSRegistrarStorage is IComputeAVSRegistrar {
    /// CONSTANTS & IMMUTABLES

    /// @dev The PermissionController contract
    IPermissionController public immutable permissionController;

    /// @dev The AppController contract
    IAppController public immutable appController;

    /// STORAGE

    /// @dev The next operator set ID to use for the next created operator set
    uint32 public nextOperatorSetId;

    /**
     * @param _permissionController The PermissionController contract
     * @param _appController The AppController contract
     */
    constructor(IPermissionController _permissionController, IAppController _appController) {
        permissionController = _permissionController;
        appController = _appController;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
