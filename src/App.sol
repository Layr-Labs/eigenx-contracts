// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {SemVerMixin} from "@eigenlayer-contracts/src/contracts/mixins/SemVerMixin.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {AppStorage} from "./storage/AppStorage.sol";
import {IApp} from "./interfaces/IApp.sol";

contract App is Initializable, SemVerMixin, AppStorage {
    constructor(string memory version, IPermissionController _permissionController)
        SemVerMixin(version)
        AppStorage(_permissionController)
    {
        _disableInitializers();
    }

    /// @inheritdoc IApp
    function initialize(address admin) external initializer {
        // Grant admin permissions to the owner
        permissionController.addPendingAdmin(address(this), admin);
    }
}
