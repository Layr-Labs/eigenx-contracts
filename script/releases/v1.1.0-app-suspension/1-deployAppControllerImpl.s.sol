// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {AppController} from "../../../src/AppController.sol";

/**
 * Purpose: deploy new AppController implementation with suspension functionality
 */
contract DeployAppControllerImpl is EOADeployer {
    using Env for *;

    AppController public newAppControllerImpl;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy new AppController implementation
        newAppControllerImpl = new AppController({
            _version: "1.1.0",
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
        // make sure impl is non-zero
        assertTrue(address(Env.impl.appController()) != address(0), "AppController impl is zero");

        // validate proxy->impl addresses
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

        // Validate AppController configuration
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
        assertEq(appController.version(), "1.1.0", "AppController version mismatch");
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
