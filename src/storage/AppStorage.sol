// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IApp} from "../interfaces/IApp.sol";

abstract contract AppStorage is IApp {
    /// CONSTANTS & IMMUTABLES
    /// @notice The EigenLayer PermissionController contract
    IPermissionController public immutable permissionController;

    constructor(IPermissionController _permissionController) {
        permissionController = _permissionController;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
