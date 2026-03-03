// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {AppController} from "../../../src/AppController.sol";
import {SafeTimelockFactory} from "../../../src/factories/SafeTimelockFactory.sol";
import {TimelockControllerImpl} from "../../../src/governance/TimelockControllerImpl.sol";

/**
 * Purpose: deploy TimelockControllerImpl, SafeTimelockFactory (impl + proxy), and new AppController implementation
 */
contract DeployContracts is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // 1. Deploy TimelockControllerImpl (no constructor args)
        TimelockControllerImpl timelockControllerImpl = new TimelockControllerImpl();
        deployContract({name: type(TimelockControllerImpl).name, deployedTo: address(timelockControllerImpl)});

        // 2. Deploy SafeTimelockFactory implementation
        SafeTimelockFactory safeTimelockFactoryImpl = new SafeTimelockFactory({
            _safeSingleton: Env.safeSingleton(),
            _safeProxyFactory: Env.safeProxyFactory(),
            _defaultFallbackHandler: Env.defaultFallbackHandler(),
            _timelockImplementation: address(timelockControllerImpl)
        });
        deployImpl({name: type(SafeTimelockFactory).name, deployedTo: address(safeTimelockFactoryImpl)});

        // 3. Deploy SafeTimelockFactory proxy
        TransparentUpgradeableProxy safeTimelockFactoryProxy = new TransparentUpgradeableProxy(
            address(safeTimelockFactoryImpl),
            address(Env.proxyAdmin()),
            abi.encodeCall(SafeTimelockFactory.initialize, ())
        );
        deployProxy({name: type(SafeTimelockFactory).name, deployedTo: address(safeTimelockFactoryProxy)});

        // 4. Deploy new AppController implementation
        AppController newAppControllerImpl = new AppController({
            _version: Env.deployVersion(),
            _permissionController: Env.permissionController(),
            _releaseManager: Env.releaseManager(),
            _computeAVSRegistrar: Env.proxy.computeAVSRegistrar(),
            _computeOperator: Env.proxy.computeOperator(),
            _appBeacon: Env.beacon.appBeacon()
        });
        deployImpl({name: type(AppController).name, deployedTo: address(newAppControllerImpl)});

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        runAsEOA();

        _validateNewAddresses({afterUpgrade: false});
        _validateConstructors();
    }

    /// @dev Validate that all deployed addresses are non-zero and proxy mappings are correct
    function _validateNewAddresses(bool afterUpgrade) internal view {
        // TimelockControllerImpl
        assertTrue(address(Env.timelockControllerImpl()) != address(0), "TimelockControllerImpl is zero");

        // SafeTimelockFactory impl and proxy
        assertTrue(address(Env.impl.safeTimelockFactory()) != address(0), "SafeTimelockFactory impl is zero");
        assertTrue(address(Env.proxy.safeTimelockFactory()) != address(0), "SafeTimelockFactory proxy is zero");

        // AppController impl
        assertTrue(address(Env.impl.appController()) != address(0), "AppController impl is zero");

        // SafeTimelockFactory proxy -> impl mapping
        assertEq(
            _getProxyImpl(address(Env.proxy.safeTimelockFactory())),
            address(Env.impl.safeTimelockFactory()),
            "SafeTimelockFactory proxy->impl mismatch"
        );

        // SafeTimelockFactory proxy admin
        assertEq(
            _getProxyAdmin(address(Env.proxy.safeTimelockFactory())),
            address(Env.proxyAdmin()),
            "SafeTimelockFactory proxyAdmin mismatch"
        );

        // AppController proxy -> impl mapping (only after upgrade)
        if (afterUpgrade) {
            assertEq(
                _getProxyImpl(address(Env.proxy.appController())),
                address(Env.impl.appController()),
                "AppController proxy->impl mismatch"
            );
        }
    }

    /// @dev Validate constructor immutables on deployed contracts
    function _validateConstructors() internal view {
        // Validate SafeTimelockFactory impl immutables
        SafeTimelockFactory safeTimelockFactoryImpl = Env.impl.safeTimelockFactory();
        assertEq(
            safeTimelockFactoryImpl.safeSingleton(),
            Env.safeSingleton(),
            "SafeTimelockFactory safeSingleton mismatch"
        );
        assertEq(
            safeTimelockFactoryImpl.safeProxyFactory(),
            Env.safeProxyFactory(),
            "SafeTimelockFactory safeProxyFactory mismatch"
        );
        assertEq(
            safeTimelockFactoryImpl.timelockImplementation(),
            address(Env.timelockControllerImpl()),
            "SafeTimelockFactory timelockImplementation mismatch"
        );

        // Validate AppController impl constructor args
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
            address(appController.appBeacon()),
            address(Env.beacon.appBeacon()),
            "AppController appBeacon mismatch"
        );
    }

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }
}
