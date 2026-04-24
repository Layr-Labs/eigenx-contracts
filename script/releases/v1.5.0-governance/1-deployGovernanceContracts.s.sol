// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {AppController} from "../../../src/AppController.sol";
import {SafeTimelockFactory} from "../../../src/factories/SafeTimelockFactory.sol";
import {TimelockControllerImpl} from "../../../src/governance/TimelockControllerImpl.sol";
import {AppAuthority} from "../../../src/governance/AppAuthority.sol";
import {IAppAuthority} from "../../../src/interfaces/IAppAuthority.sol";

/**
 * @title DeployGovernanceContracts (v1.5.0-governance phase 1)
 * @notice EOA phase of the governance release.
 *
 * Deploys, in order:
 *   1. TimelockControllerImpl — clone-master for Timelocks created by the
 *      factory. Deployed directly (not behind a proxy); immutable.
 *   2. SafeTimelockFactory implementation — references the Timelock impl
 *      above and the canonical Gnosis Safe infrastructure for the current
 *      chain (pulled from zeus env).
 *   3. SafeTimelockFactory proxy — TransparentUpgradeableProxy pointing at
 *      the impl. ProxyAdmin is the existing protocol ProxyAdmin.
 *   4. New AppController implementation — wired to the factory proxy so
 *      the timelocked-flag detection (safeTimelockFactory.isTimelock) has
 *      a real target once the AppController proxy is upgraded in phase 2.
 *
 * The AppController proxy itself is NOT upgraded here; the multisig phase
 * does that in 2-upgradeAppController.s.sol.
 */
contract DeployGovernanceContracts is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // 1. TimelockControllerImpl — no constructor args, immutable.
        TimelockControllerImpl timelockImpl = new TimelockControllerImpl();
        deployContract({name: type(TimelockControllerImpl).name, deployedTo: address(timelockImpl)});

        // 2. SafeTimelockFactory implementation.
        SafeTimelockFactory safeTimelockFactoryImpl = new SafeTimelockFactory({
            _safeSingleton: Env.safeSingleton(),
            _safeProxyFactory: Env.safeProxyFactory(),
            _safeFallbackHandler: Env.safeFallbackHandler(),
            _timelockImplementation: address(timelockImpl)
        });
        deployImpl({name: type(SafeTimelockFactory).name, deployedTo: address(safeTimelockFactoryImpl)});

        // 3. SafeTimelockFactory proxy (TransparentUpgradeableProxy).
        TransparentUpgradeableProxy safeTimelockFactoryProxy = new TransparentUpgradeableProxy(
            address(safeTimelockFactoryImpl),
            address(Env.proxyAdmin()),
            abi.encodeCall(SafeTimelockFactory.initialize, ())
        );
        deployProxy({name: type(SafeTimelockFactory).name, deployedTo: address(safeTimelockFactoryProxy)});

        // 4. AppAuthority implementation — consumer-bound to the existing
        // AppController proxy. AppAuthority's `consumer` immutable is the
        // AppController proxy address; AppController at the proxy already
        // exists (it's the v1.4.0 AppController we're about to upgrade).
        AppAuthority appAuthorityImpl = new AppAuthority(address(Env.proxy.appController()));
        deployImpl({name: type(AppAuthority).name, deployedTo: address(appAuthorityImpl)});

        // 5. AppAuthority proxy.
        TransparentUpgradeableProxy appAuthorityProxy = new TransparentUpgradeableProxy(
            address(appAuthorityImpl), address(Env.proxyAdmin()), abi.encodeCall(AppAuthority.initialize, ())
        );
        deployProxy({name: type(AppAuthority).name, deployedTo: address(appAuthorityProxy)});

        // 6. New AppController implementation — wired to the AppAuthority
        //    proxy. AppController no longer references SafeTimelockFactory;
        //    the factory is still deployed (steps 2–3) because users /
        //    tooling can use it to deploy attested Safes and Timelocks, but
        //    AppController's correctness no longer depends on it.
        AppController newAppControllerImpl = new AppController({
            _version: Env.deployVersion(),
            _permissionController: Env.permissionController(),
            _releaseManager: Env.releaseManager(),
            _computeAVSRegistrar: Env.proxy.computeAVSRegistrar(),
            _computeOperator: Env.proxy.computeOperator(),
            _appBeacon: Env.beacon.appBeacon(),
            _appAuthority: IAppAuthority(address(appAuthorityProxy))
        });
        deployImpl({name: type(AppController).name, deployedTo: address(newAppControllerImpl)});

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        runAsEOA();

        _validateNewAddresses({afterUpgrade: false});
        _validateConstructors();
    }

    /// @dev Validate all phase-1 deployments are non-zero and wired correctly.
    function _validateNewAddresses(bool afterUpgrade) internal view {
        assertTrue(address(Env.timelockControllerImpl()) != address(0), "TimelockControllerImpl zero");
        assertTrue(address(Env.impl.safeTimelockFactory()) != address(0), "SafeTimelockFactory impl zero");
        assertTrue(address(Env.proxy.safeTimelockFactory()) != address(0), "SafeTimelockFactory proxy zero");
        assertTrue(address(Env.impl.appController()) != address(0), "AppController impl zero");

        // Proxy points at the fresh impl.
        assertEq(
            _getProxyImpl(address(Env.proxy.safeTimelockFactory())),
            address(Env.impl.safeTimelockFactory()),
            "SafeTimelockFactory proxy -> impl mismatch"
        );

        // Only after phase 2 does the AppController proxy point at the new impl.
        if (afterUpgrade) {
            assertEq(
                _getProxyImpl(address(Env.proxy.appController())),
                address(Env.impl.appController()),
                "AppController proxy -> impl mismatch"
            );
        }
    }

    /// @dev Cross-check immutables on the freshly deployed implementations.
    function _validateConstructors() internal view {
        SafeTimelockFactory factoryImpl = Env.impl.safeTimelockFactory();
        assertEq(factoryImpl.safeSingleton(), Env.safeSingleton(), "factory safeSingleton mismatch");
        assertEq(factoryImpl.safeProxyFactory(), Env.safeProxyFactory(), "factory safeProxyFactory mismatch");
        assertEq(factoryImpl.safeFallbackHandler(), Env.safeFallbackHandler(), "factory safeFallbackHandler mismatch");
        assertEq(
            factoryImpl.timelockImplementation(),
            address(Env.timelockControllerImpl()),
            "factory timelockImplementation mismatch"
        );

        AppController appImpl = Env.impl.appController();
        assertEq(appImpl.version(), Env.deployVersion(), "AppController version mismatch");
        assertEq(
            address(appImpl.permissionController()),
            address(Env.permissionController()),
            "permissionController mismatch"
        );
        assertEq(address(appImpl.releaseManager()), address(Env.releaseManager()), "releaseManager mismatch");
        assertEq(
            address(appImpl.computeAVSRegistrar()),
            address(Env.proxy.computeAVSRegistrar()),
            "computeAVSRegistrar mismatch"
        );
        assertEq(address(appImpl.computeOperator()), address(Env.proxy.computeOperator()), "computeOperator mismatch");
        assertEq(address(appImpl.appBeacon()), address(Env.beacon.appBeacon()), "appBeacon mismatch");
        assertEq(address(appImpl.appAuthority()), address(Env.proxy.appAuthority()), "appAuthority mismatch");

        AppAuthority authorityImpl = Env.impl.appAuthority();
        assertEq(authorityImpl.consumer(), address(Env.proxy.appController()), "AppAuthority consumer mismatch");
    }

    function _getProxyImpl(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }
}
