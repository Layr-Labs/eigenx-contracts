// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Deploy} from "./1-deployContracts.s.sol";
import "../Env.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";

/**
 * Purpose:
 *      * enqueue a multisig transaction which;
 *             - upgrades EP
 *  This should be run via the operations multisig.
 */
contract AcceptAdminAndSetLimits is MultisigBuilder, Deploy {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.computeOpsMultisig()) {
        Env.permissionController().acceptAdmin(address(Env.proxy.appController()));
        Env.permissionController()
            .setAppointee(
                address(Env.proxy.appController()),
                Env.computeWhitelisterMultisig(),
                address(Env.proxy.appController()),
                AppController.setMaxGlobalActiveApps.selector
            );
        Env.permissionController()
            .setAppointee(
                address(Env.proxy.appController()),
                Env.computeWhitelisterMultisig(),
                address(Env.proxy.appController()),
                AppController.setMaxActiveAppsPerUser.selector
            );
        Env.proxy.appController().setMaxGlobalActiveApps(Env.MAX_GLOBAL_ACTIVE_APPS());
    }

    function testScript() public virtual override {
        runAsEOA();

        execute();

        _validateNewAddresses();
        _validateSetup();
    }

    function _validateSetup() internal view {
        assertTrue(
            Env.permissionController()
                .canCall(
                    address(Env.proxy.appController()),
                    Env.computeWhitelisterMultisig(),
                    address(Env.proxy.appController()),
                    AppController.setMaxGlobalActiveApps.selector
                ),
            "ComputeAdmin not admin"
        );
        assertTrue(
            Env.permissionController()
                .canCall(
                    address(Env.proxy.appController()),
                    Env.computeWhitelisterMultisig(),
                    address(Env.proxy.appController()),
                    AppController.setMaxActiveAppsPerUser.selector
                ),
            "ComputeAdmin not admin"
        );
        assertEq(
            Env.proxy.appController().maxGlobalActiveApps(),
            Env.MAX_GLOBAL_ACTIVE_APPS(),
            "Max global active apps mismatch"
        );
    }
}
