// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import "./1-deployAppControllerImpl.s.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * Purpose: upgrade AppController proxy to new implementation and grant billing admin permissions
 */
contract UpgradeAndGrantPermissions is MultisigBuilder, DeployAppControllerImpl {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.computeOpsMultisig()) {
        Env.proxyAdmin()
            .upgrade(
                ITransparentUpgradeableProxy(address(Env.proxy.appController())), address(Env.impl.appController())
            );

        // Grant billing admin permission to set user limits
        Env.permissionController()
            .setAppointee(
                address(Env.proxy.appController()),
                Env.billingAdmin(),
                address(Env.proxy.appController()),
                AppController.setMaxActiveAppsPerUser.selector
            );

        // Grant billing admin permission to suspend accounts
        Env.permissionController()
            .setAppointee(
                address(Env.proxy.appController()),
                Env.billingAdmin(),
                address(Env.proxy.appController()),
                AppController.suspend.selector
            );
    }

    /// @dev Validate that computeOpsMultisig has admin permissions on AppController
    function _validatePermissions() internal view {
        assertTrue(
            Env.permissionController().isAdmin(address(Env.proxy.appController()), Env.computeOpsMultisig()),
            "computeOpsMultisig is not admin of AppController"
        );
        assertTrue(
            Env.permissionController()
                .canCall(
                    address(Env.proxy.appController()),
                    Env.billingAdmin(),
                    address(Env.proxy.appController()),
                    AppController.setMaxActiveAppsPerUser.selector
                ),
            "BillingAdmin cannot call setMaxActiveAppsPerUser"
        );
        assertTrue(
            Env.permissionController()
                .canCall(
                    address(Env.proxy.appController()),
                    Env.billingAdmin(),
                    address(Env.proxy.appController()),
                    AppController.suspend.selector
                ),
            "BillingAdmin cannot call suspend"
        );
    }

    function testScript() public virtual override {
        runAsEOA();

        execute();

        _validateNewAddresses({afterUpgrade: true});
        _validatePermissions();
    }
}
