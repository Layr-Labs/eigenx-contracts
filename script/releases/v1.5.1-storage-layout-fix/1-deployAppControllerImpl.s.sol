// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {AppController} from "../../../src/AppController.sol";

/**
 * Purpose: deploy new AppController implementation that restores the AppConfigStorage layout.
 *
 * The v1.5.0 implementation inherited a struct field (`pendingReleaseBlockNumber`) that had been
 * inserted into the middle of `AppConfigStorage`, before `status`. That pushed the struct from 30
 * bytes (one packed slot) to 34 bytes (two slots), shifting `status` and `billingType` into a
 * fresh, never-written slot. As a result every pre-existing app read back with `status == NONE`
 * ("uninitialized") and `billingType == DEFAULT` after the v1.5.0 upgrade.
 *
 * v1.5.1 appends `pendingReleaseBlockNumber` to the end of `AppConfigStorage` so the original five
 * fields keep their packed offsets in slot 0, restoring the correct status for every app that was
 * not mutated under the broken v1.5.0 layout.
 */
contract DeployAppControllerImpl is EOADeployer {
    using Env for *;

    AppController public newAppControllerImpl;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy new AppController implementation
        newAppControllerImpl = new AppController({
            _version: Env.deployVersion(),
            _permissionController: Env.permissionController(),
            _releaseManager: Env.releaseManager(),
            _computeAVSRegistrar: Env.proxy.computeAVSRegistrar(),
            _computeOperator: Env.proxy.computeOperator(),
            _appBeacon: Env.beacon.appBeacon()
        });

        // Register new implementation in Env system
        deployImpl({name: type(AppController).name, deployedTo: address(newAppControllerImpl)});

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        runAsEOA();

        _validateNewAddresses({afterUpgrade: false});
        _validateConstructors();
    }

    /// @dev Validate the new `Env.impl` addresses are non-zero
    function _validateNewAddresses(bool afterUpgrade) internal view {
        assertTrue(address(Env.impl.appController()) != address(0), "AppController impl is zero");

        if (afterUpgrade) {
            assertEq(
                _getProxyImpl(address(Env.proxy.appController())),
                address(Env.impl.appController()),
                "AppController proxy->impl mismatch"
            );
        }
    }

    /// @dev Validate that the storage of the newly deployed implementation
    function _validateConstructors() internal view {
        AppController appController = Env.impl.appController();

        assertEq(
            address(appController.permissionController()),
            address(Env.permissionController()),
            "AppController permissionController mismatch"
        );
        assertEq(
            address(appController.releaseManager()),
            address(Env.releaseManager()),
            "AppController releaseManager mismatch"
        );
        assertEq(appController.version(), Env.deployVersion(), "AppController version mismatch");
        assertEq(
            address(appController.computeAVSRegistrar()),
            address(Env.proxy.computeAVSRegistrar()),
            "AppController computeAVSRegistrar mismatch"
        );
        assertEq(
            address(appController.computeOperator()),
            address(Env.proxy.computeOperator()),
            "AppController computeOperator mismatch"
        );
        assertEq(
            address(appController.appBeacon()), address(Env.beacon.appBeacon()), "AppController appBeacon mismatch"
        );
    }

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }
}
