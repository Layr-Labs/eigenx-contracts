// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {Parser} from "./Parser.s.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {
    TransparentUpgradeableProxy,
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {IDelegationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IReleaseManager} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {EmptyContract} from "@eigenlayer-contracts/src/test/mocks/EmptyContract.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {App} from "../src/App.sol";
import {IAppController} from "../src/interfaces/IAppController.sol";
import {IComputeAVSRegistrar} from "../src/interfaces/IComputeAVSRegistrar.sol";
import {IComputeOperator} from "../src/interfaces/IComputeOperator.sol";
import {AppController} from "../src/AppController.sol";
import {ComputeAVSRegistrar} from "../src/ComputeAVSRegistrar.sol";
import {ComputeOperator} from "../src/ComputeOperator.sol";
import {ImageAllowlist} from "../src/ImageAllowlist.sol";
import {IImageAllowlist} from "../src/interfaces/IImageAllowlist.sol";
import {SafeTimelockFactory} from "../src/factories/SafeTimelockFactory.sol";
import {TimelockControllerImpl} from "../src/governance/TimelockControllerImpl.sol";
import {IAppAuthority} from "../src/interfaces/IAppAuthority.sol";
import {AppAuthority} from "../src/governance/AppAuthority.sol";

contract Deploy is Parser {
    struct Proxies {
        TransparentUpgradeableProxy computeAVSRegistrar;
        TransparentUpgradeableProxy computeOperator;
        TransparentUpgradeableProxy appController;
        TransparentUpgradeableProxy imageAllowlist;
    }

    struct Implementations {
        App app;
        ComputeAVSRegistrar computeAVSRegistrar;
        ComputeOperator computeOperator;
        AppController appController;
        ImageAllowlist imageAllowlist;
        AppAuthority appAuthority;
    }

    function run(string memory environment) public {
        run(environment, parseDeployParams(environment));
    }

    function run(string memory environment, DeployParams memory params) public {
        vm.startBroadcast();

        _writeOutputToJson(environment, deploy(params));

        vm.stopBroadcast();
    }

    function deploy(DeployParams memory params) public returns (DeployedContracts memory) {
        require(bytes(params.version).length != 0, "Version must not be empty");
        require(address(params.delegationManager) != address(0), "Delegation manager must not be empty");
        require(address(params.allocationManager) != address(0), "Allocation manager must not be empty");
        require(address(params.permissionController) != address(0), "Permission controller must not be empty");
        require(address(params.keyRegistrar) != address(0), "Key registrar must not be empty");
        require(address(params.releaseManager) != address(0), "Release manager must not be empty");
        require(address(params.initialOwner) != address(0), "Initial owner must not be empty");
        require(bytes(params.operatorMetadataURI).length != 0, "Operator metadata URI must not be empty");
        require(bytes(params.avsMetadataURI).length != 0, "AVS metadata URI must not be empty");

        // Deploy proxy admin if not provided
        if (address(params.proxyAdmin) == address(0)) {
            params.proxyAdmin = new ProxyAdmin();
        }
        EmptyContract emptyContract = new EmptyContract();

        // Deploy proxies
        Proxies memory proxies = Proxies({
            computeAVSRegistrar: new TransparentUpgradeableProxy(
                address(emptyContract), address(params.proxyAdmin), new bytes(0)
            ),
            computeOperator: new TransparentUpgradeableProxy(
                address(emptyContract), address(params.proxyAdmin), new bytes(0)
            ),
            appController: new TransparentUpgradeableProxy(
                address(emptyContract), address(params.proxyAdmin), new bytes(0)
            ),
            imageAllowlist: new TransparentUpgradeableProxy(
                address(emptyContract), address(params.proxyAdmin), new bytes(0)
            )
        });

        // Deploy App beacon
        UpgradeableBeacon appBeacon =
            new UpgradeableBeacon(address(new App(params.version, IPermissionController(params.permissionController))));

        // Deploy SafeTimelockFactory so tests / local deploys can exercise
        // the deployTimelock / deploySafe paths. AppController no longer
        // depends on this factory — governance is whatever the app's owner
        // contract is. Safe infrastructure addresses are zero here; local
        // deploys don't exercise deploySafe. Production releases wire real
        // Safe addresses in the v1.5.0 release script.
        TimelockControllerImpl timelockImpl = new TimelockControllerImpl();
        SafeTimelockFactory safeTimelockFactoryImpl = new SafeTimelockFactory({
            _safeSingleton: address(0),
            _safeProxyFactory: address(0),
            _safeFallbackHandler: address(0),
            _timelockImplementation: address(timelockImpl)
        });
        new TransparentUpgradeableProxy(
            address(safeTimelockFactoryImpl),
            address(params.proxyAdmin),
            abi.encodeCall(SafeTimelockFactory.initialize, ())
        );

        // Deploy AppAuthority (owns per-app RBAC state consumed by
        // AppController). The consumer (AppController proxy) is already
        // known — it's the proxy we just created above. The impl holds
        // the consumer immutable; the proxy just fronts it.
        AppAuthority appAuthorityImpl = new AppAuthority(address(proxies.appController));
        TransparentUpgradeableProxy appAuthorityProxy = new TransparentUpgradeableProxy(
            address(appAuthorityImpl), address(params.proxyAdmin), abi.encodeCall(AppAuthority.initialize, ())
        );

        // Deploy implementation contracts
        Implementations memory impls = Implementations({
            app: App(appBeacon.implementation()),
            computeAVSRegistrar: new ComputeAVSRegistrar({
                _version: params.version,
                _allocationManager: params.allocationManager,
                _permissionController: params.permissionController,
                _keyRegistrar: params.keyRegistrar,
                _appController: IAppController(address(proxies.appController))
            }),
            computeOperator: new ComputeOperator({
                _version: params.version,
                _delegationManager: params.delegationManager,
                _allocationManager: params.allocationManager,
                _permissionController: params.permissionController,
                _appController: address(proxies.appController),
                _computeAVSRegistrar: address(proxies.computeAVSRegistrar)
            }),
            appController: new AppController({
                _version: params.version,
                _permissionController: params.permissionController,
                _releaseManager: params.releaseManager,
                _computeAVSRegistrar: IComputeAVSRegistrar(address(proxies.computeAVSRegistrar)),
                _computeOperator: IComputeOperator(address(proxies.computeOperator)),
                _appBeacon: appBeacon,
                _appAuthority: IAppAuthority(address(appAuthorityProxy))
            }),
            imageAllowlist: new ImageAllowlist(),
            appAuthority: appAuthorityImpl
        });

        // Upgrade proxies using ProxyAdmin
        params.proxyAdmin
            .upgradeAndCall(
                ITransparentUpgradeableProxy(address(proxies.computeAVSRegistrar)),
                address(impls.computeAVSRegistrar),
                abi.encodeCall(ComputeAVSRegistrar.initialize, (params.avsMetadataURI))
            );
        params.proxyAdmin
            .upgradeAndCall(
                ITransparentUpgradeableProxy(address(proxies.computeOperator)),
                address(impls.computeOperator),
                abi.encodeCall(ComputeOperator.initialize, (params.operatorMetadataURI))
            );
        params.proxyAdmin
                .upgradeAndCall(
                ITransparentUpgradeableProxy(address(proxies.appController)),
                address(impls.appController),
                abi.encodeCall(AppController.initialize, (params.initialOwner))
            );
        params.proxyAdmin
            .upgradeAndCall(
                ITransparentUpgradeableProxy(address(proxies.imageAllowlist)),
                address(impls.imageAllowlist),
                abi.encodeCall(ImageAllowlist.initialize, (params.initialOwner))
            );

        // Accept admin role for AppController
        IPermissionController(address(params.permissionController)).acceptAdmin(address(proxies.appController));

        // Set max global active apps and admin user limit
        IAppController(address(proxies.appController)).setMaxGlobalActiveApps(params.maxGlobalActiveApps);
        IAppController(address(proxies.appController))
            .setMaxActiveAppsPerUser(params.initialOwner, params.adminMaxActiveApps);

        console.log("Proxy Admin:", address(params.proxyAdmin));
        console.log("App Beacon:", address(appBeacon));
        console.log("AppController deployed at:", address(proxies.appController));
        console.log("ComputeAVSRegistrar deployed at:", address(proxies.computeAVSRegistrar));
        console.log("ComputeOperator deployed at:", address(proxies.computeOperator));
        console.log("ImageAllowlist deployed at:", address(proxies.imageAllowlist));

        return DeployedContracts({
            proxyAdmin: params.proxyAdmin,
            appBeacon: appBeacon,
            appImpl: impls.app,
            appController: IAppController(address(proxies.appController)),
            appControllerImpl: impls.appController,
            computeAVSRegistrar: IComputeAVSRegistrar(address(proxies.computeAVSRegistrar)),
            computeAVSRegistrarImpl: impls.computeAVSRegistrar,
            computeOperator: IComputeOperator(address(proxies.computeOperator)),
            computeOperatorImpl: impls.computeOperator,
            imageAllowlist: IImageAllowlist(address(proxies.imageAllowlist)),
            imageAllowlistImpl: impls.imageAllowlist,
            appAuthority: IAppAuthority(address(appAuthorityProxy)),
            appAuthorityImpl: impls.appAuthority
        });
    }

    function deployForTesting(DeployParams memory params) public returns (DeployedContracts memory deployment) {
        vm.startPrank(params.initialOwner);
        deployment = deploy(params);
        vm.stopPrank();
    }

    function _writeOutputToJson(string memory environment, DeployedContracts memory deployedContracts) internal {
        // Add the addresses object
        string memory addresses = "addresses";
        vm.serializeAddress(addresses, "proxyAdmin", address(deployedContracts.proxyAdmin));
        vm.serializeAddress(addresses, "appBeacon", address(deployedContracts.appBeacon));
        vm.serializeAddress(addresses, "appImpl", address(deployedContracts.appImpl));
        vm.serializeAddress(addresses, "appController", address(deployedContracts.appController));
        vm.serializeAddress(addresses, "appControllerImpl", address(deployedContracts.appControllerImpl));
        vm.serializeAddress(addresses, "computeAVSRegistrar", address(deployedContracts.computeAVSRegistrar));
        vm.serializeAddress(addresses, "computeAVSRegistrarImpl", address(deployedContracts.computeAVSRegistrarImpl));
        vm.serializeAddress(addresses, "computeOperator", address(deployedContracts.computeOperator));
        vm.serializeAddress(addresses, "computeOperatorImpl", address(deployedContracts.computeOperatorImpl));
        vm.serializeAddress(addresses, "imageAllowlist", address(deployedContracts.imageAllowlist));
        vm.serializeAddress(addresses, "imageAllowlistImpl", address(deployedContracts.imageAllowlistImpl));
        vm.serializeAddress(addresses, "appAuthority", address(deployedContracts.appAuthority));
        addresses = vm.serializeAddress(addresses, "appAuthorityImpl", address(deployedContracts.appAuthorityImpl));

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat("script/deploys/", environment, "/output.json");

        // Create directory if it doesn't exist
        string memory outputDir = string.concat("script/deploys/", environment);
        vm.createDir(outputDir, true);

        vm.writeJson(finalJson, outputFile);
    }
}
