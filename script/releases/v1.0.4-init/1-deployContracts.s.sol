// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import {IDelegationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IReleaseManager} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {EmptyContract} from "@eigenlayer-contracts/src/test/mocks/EmptyContract.sol";

import {App} from "../../../src/App.sol";
import {AppController} from "../../../src/AppController.sol";
import {ComputeAVSRegistrar} from "../../../src/ComputeAVSRegistrar.sol";
import {ComputeOperator} from "../../../src/ComputeOperator.sol";
import {IAppController} from "../../../src/interfaces/IAppController.sol";
import {IComputeAVSRegistrar} from "../../../src/interfaces/IComputeAVSRegistrar.sol";
import {IComputeOperator} from "../../../src/interfaces/IComputeOperator.sol";

/**
 * Purpose: use an EOA to deploy all compute contracts.
 */
contract Deploy is EOADeployer {
    using Env for *;

    struct DeployedContracts {
        ProxyAdmin proxyAdmin;
        UpgradeableBeacon appBeacon;
        App appImpl;
        AppController appController;
        AppController appControllerImpl;
        ComputeAVSRegistrar computeAVSRegistrar;
        ComputeAVSRegistrar computeAVSRegistrarImpl;
        ComputeOperator computeOperator;
        ComputeOperator computeOperatorImpl;
    }

    function _deploy() internal returns (DeployedContracts memory) {
        // Deploy proxy admin
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        EmptyContract emptyContract = new EmptyContract();

        // Deploy proxies
        TransparentUpgradeableProxy computeAVSRegistrarProxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), new bytes(0));
        TransparentUpgradeableProxy computeOperatorProxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), new bytes(0));
        TransparentUpgradeableProxy appControllerProxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), new bytes(0));

        // Deploy App implementation and beacon
        App appImpl = new App(Env.deployVersion(), Env.permissionController());
        UpgradeableBeacon appBeacon = new UpgradeableBeacon(address(appImpl));

        // Deploy implementation contracts
        ComputeAVSRegistrar computeAVSRegistrarImpl = new ComputeAVSRegistrar({
            _version: Env.deployVersion(),
            _allocationManager: Env.allocationManager(),
            _permissionController: Env.permissionController(),
            _keyRegistrar: Env.keyRegistrar(),
            _appController: IAppController(address(appControllerProxy))
        });

        ComputeOperator computeOperatorImpl = new ComputeOperator({
            _version: Env.deployVersion(),
            _delegationManager: Env.delegationManager(),
            _allocationManager: Env.allocationManager(),
            _permissionController: Env.permissionController(),
            _appController: address(appControllerProxy),
            _computeAVSRegistrar: address(computeAVSRegistrarProxy)
        });

        AppController appControllerImpl = new AppController({
            _version: Env.deployVersion(),
            _permissionController: Env.permissionController(),
            _releaseManager: Env.releaseManager(),
            _computeAVSRegistrar: IComputeAVSRegistrar(address(computeAVSRegistrarProxy)),
            _computeOperator: IComputeOperator(address(computeOperatorProxy)),
            _appBeacon: appBeacon,
            _billingCore: address(0) // Billing not yet enabled
        });

        // Upgrade proxies using ProxyAdmin
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(computeAVSRegistrarProxy)),
            address(computeAVSRegistrarImpl),
            abi.encodeCall(ComputeAVSRegistrar.initialize, (Env.AVS_METADATA_URI()))
        );

        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(computeOperatorProxy)),
            address(computeOperatorImpl),
            abi.encodeCall(ComputeOperator.initialize, (Env.OPERATOR_METADATA_URI()))
        );

        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(appControllerProxy)),
            address(appControllerImpl),
            abi.encodeCall(AppController.initialize, (Env.computeOpsMultisig()))
        );

        proxyAdmin.transferOwnership(Env.computeOpsMultisig());
        appBeacon.transferOwnership(Env.computeOpsMultisig());

        return DeployedContracts({
            proxyAdmin: proxyAdmin,
            appBeacon: appBeacon,
            appImpl: appImpl,
            appController: AppController(address(appControllerProxy)),
            appControllerImpl: appControllerImpl,
            computeAVSRegistrar: ComputeAVSRegistrar(address(computeAVSRegistrarProxy)),
            computeAVSRegistrarImpl: computeAVSRegistrarImpl,
            computeOperator: ComputeOperator(address(computeOperatorProxy)),
            computeOperatorImpl: computeOperatorImpl
        });
    }

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        vm.startBroadcast();

        DeployedContracts memory deployed = _deploy();

        // Register implementations in Env system
        deployImpl({
            name: type(AppController).name,
            deployedTo: address(deployed.appControllerImpl)
        });

        deployImpl({
            name: type(ComputeAVSRegistrar).name,
            deployedTo: address(deployed.computeAVSRegistrarImpl)
        });

        deployImpl({
            name: type(ComputeOperator).name,
            deployedTo: address(deployed.computeOperatorImpl)
        });

        // Register proxies in Env system
        deployProxy({
            name: type(AppController).name,
            deployedTo: address(deployed.appController)
        });

        deployProxy({
            name: type(ComputeAVSRegistrar).name,
            deployedTo: address(deployed.computeAVSRegistrar)
        });

        deployProxy({
            name: type(ComputeOperator).name,
            deployedTo: address(deployed.computeOperator)
        });

        // Register beacon
        deployImpl({
            name: type(App).name,
            deployedTo: address(deployed.appImpl)
        });

        deployBeacon({
            name: type(App).name,
            deployedTo: address(deployed.appBeacon)
        });

        // Register proxy admin
        deployContract({
            name: type(ProxyAdmin).name,
            deployedTo: address(deployed.proxyAdmin)
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        runAsEOA();

        _validateNewAddresses();
        _validateConstructors();
    }

    /// @dev Validated the new `Env.impl` addresses are non-zero
    function _validateNewAddresses() internal view {
        // make sure impls are non-zero
        assertTrue(address(Env.impl.app()) != address(0), "App impl is zero");
        assertTrue(address(Env.impl.appController()) != address(0), "AppController impl is zero");
        assertTrue(address(Env.impl.computeAVSRegistrar()) != address(0), "ComputeAVSRegistrar impl is zero");
        assertTrue(address(Env.impl.computeOperator()) != address(0), "ComputeOperator impl is zero");

        // make sure proxies are non-zero
        assertTrue(address(Env.proxy.appController()) != address(0), "AppController proxy is zero");
        assertTrue(address(Env.proxy.computeAVSRegistrar()) != address(0), "ComputeAVSRegistrar proxy is zero");
        assertTrue(address(Env.proxy.computeOperator()) != address(0), "ComputeOperator proxy is zero");

        // make sure beacon is non-zero
        assertTrue(address(Env.beacon.appBeacon()) != address(0), "AppBeacon is zero");

        // validate proxy->impl addresses
        assertEq(
            _getProxyImpl(address(Env.proxy.appController())),
            address(Env.impl.appController()),
            "AppController proxy->impl mismatch"
        );
        assertEq(
            _getProxyImpl(address(Env.proxy.computeAVSRegistrar())),
            address(Env.impl.computeAVSRegistrar()),
            "ComputeAVSRegistrar proxy->impl mismatch"
        );
        assertEq(
            _getProxyImpl(address(Env.proxy.computeOperator())),
            address(Env.impl.computeOperator()),
            "ComputeOperator proxy->impl mismatch"
        );
        assertEq(Env.beacon.appBeacon().implementation(), address(Env.impl.app()), "AppBeacon proxy->impl mismatch");

        // validate proxy admins
        assertEq(
            _getProxyAdmin(address(Env.proxy.appController())),
            address(Env.proxyAdmin()),
            "AppController proxyAdmin mismatch"
        );
        assertEq(
            _getProxyAdmin(address(Env.proxy.computeAVSRegistrar())),
            address(Env.proxyAdmin()),
            "ComputeAVSRegistrar proxyAdmin mismatch"
        );
        assertEq(
            _getProxyAdmin(address(Env.proxy.computeOperator())),
            address(Env.proxyAdmin()),
            "ComputeOperator proxyAdmin mismatch"
        );
        assertEq(Env.beacon.appBeacon().owner(), address(Env.computeOpsMultisig()), "AppBeacon proxyAdmin mismatch");

        assertEq(Env.proxyAdmin().owner(), Env.computeOpsMultisig(), "ProxyAdmin owner mismatch");
    }

    /// @dev Validate that the storage of the newly deployed and upgraded contracts
    function _validateConstructors() internal view {
        AppController appController = Env.impl.appController();
        ComputeOperator computeOperator = Env.impl.computeOperator();
        ComputeAVSRegistrar computeAVSRegistrar = Env.impl.computeAVSRegistrar();
        App app = Env.impl.app();

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

        // Validate ComputeOperator configuration
        assertEq(
            address(computeOperator.delegationManager()),
            address(Env.delegationManager()),
            "ComputeOperator delegationManager mismatch"
        );
        assertEq(
            address(computeOperator.allocationManager()),
            address(Env.allocationManager()),
            "ComputeOperator allocationManager mismatch"
        );
        assertEq(computeOperator.version(), Env.deployVersion(), "ComputeOperator version mismatch");
        assertEq(
            address(computeOperator.computeAVSRegistrar()),
            address(Env.proxy.computeAVSRegistrar()),
            "ComputeOperator computeAVSRegistrar mismatch"
        );
        assertEq(
            address(computeOperator.appController()),
            address(Env.proxy.appController()),
            "ComputeOperator appController mismatch"
        );

        // Validate ComputeAVSRegistrar configuration
        assertEq(
            address(computeAVSRegistrar.allocationManager()),
            address(Env.allocationManager()),
            "ComputeAVSRegistrar allocationManager mismatch"
        );
        assertEq(
            address(computeAVSRegistrar.keyRegistrar()),
            address(Env.keyRegistrar()),
            "ComputeAVSRegistrar keyRegistrar mismatch"
        );
        assertEq(computeAVSRegistrar.version(), Env.deployVersion(), "ComputeAVSRegistrar version mismatch");
        assertEq(
            address(computeAVSRegistrar.appController()),
            address(Env.proxy.appController()),
            "ComputeAVSRegistrar appController mismatch"
        );
        // Validate App configuration
        assertEq(app.version(), Env.deployVersion(), "App version mismatch");
    }

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }

    function _strEq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}
