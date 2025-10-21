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

contract Deploy is Parser {
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
        TransparentUpgradeableProxy computeAVSRegistrarProxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(params.proxyAdmin), new bytes(0));
        TransparentUpgradeableProxy computeOperatorProxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(params.proxyAdmin), new bytes(0));
        TransparentUpgradeableProxy appControllerProxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(params.proxyAdmin), new bytes(0));

        // Deploy App implementation and beacon
        App appImpl = new App(params.version, IPermissionController(params.permissionController));
        UpgradeableBeacon appBeacon = new UpgradeableBeacon(address(appImpl));

        // Deploy implementation contracts

        ComputeAVSRegistrar computeAVSRegistrarImpl = new ComputeAVSRegistrar({
            _version: params.version,
            _allocationManager: params.allocationManager,
            _permissionController: params.permissionController,
            _keyRegistrar: params.keyRegistrar,
            _appController: IAppController(address(appControllerProxy))
        });
        ComputeOperator computeOperatorImpl = new ComputeOperator({
            _version: params.version,
            _delegationManager: params.delegationManager,
            _allocationManager: params.allocationManager,
            _permissionController: params.permissionController,
            _appController: address(appControllerProxy),
            _computeAVSRegistrar: address(computeAVSRegistrarProxy)
        });
        AppController appControllerImpl = new AppController({
            _version: params.version,
            _permissionController: params.permissionController,
            _releaseManager: params.releaseManager,
            _computeAVSRegistrar: IComputeAVSRegistrar(address(computeAVSRegistrarProxy)),
            _computeOperator: IComputeOperator(address(computeOperatorProxy)),
            _appBeacon: appBeacon
        });

        // Upgrade proxies using ProxyAdmin
        params.proxyAdmin
            .upgradeAndCall(
                ITransparentUpgradeableProxy(address(computeAVSRegistrarProxy)),
                address(computeAVSRegistrarImpl),
                abi.encodeCall(ComputeAVSRegistrar.initialize, (params.avsMetadataURI))
            );
        params.proxyAdmin
            .upgradeAndCall(
                ITransparentUpgradeableProxy(address(computeOperatorProxy)),
                address(computeOperatorImpl),
                abi.encodeCall(ComputeOperator.initialize, (params.operatorMetadataURI))
            );
        params.proxyAdmin
            .upgradeAndCall(
                ITransparentUpgradeableProxy(address(appControllerProxy)),
                address(appControllerImpl),
                abi.encodeCall(AppController.initialize, (params.initialOwner))
            );

        // Accept admin role for AppController
        IPermissionController(address(params.permissionController)).acceptAdmin(address(appControllerProxy));

        // Set max global active apps and admin user limit
        IAppController(address(appControllerProxy)).setMaxGlobalActiveApps(params.maxGlobalActiveApps);
        IAppController(address(appControllerProxy))
            .setMaxActiveAppsPerUser(params.initialOwner, params.adminMaxActiveApps);

        console.log("Proxy Admin:", address(params.proxyAdmin));
        console.log("App Beacon:", address(appBeacon));
        console.log("AppController deployed at:", address(appControllerProxy));
        console.log("ComputeAVSRegistrar deployed at:", address(computeAVSRegistrarProxy));
        console.log("ComputeOperator deployed at:", address(computeOperatorProxy));

        return DeployedContracts({
            proxyAdmin: params.proxyAdmin,
            appBeacon: appBeacon,
            appImpl: appImpl,
            appController: IAppController(address(appControllerProxy)),
            appControllerImpl: appControllerImpl,
            computeAVSRegistrar: IComputeAVSRegistrar(address(computeAVSRegistrarProxy)),
            computeAVSRegistrarImpl: computeAVSRegistrarImpl,
            computeOperator: IComputeOperator(address(computeOperatorProxy)),
            computeOperatorImpl: computeOperatorImpl
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
        addresses =
            vm.serializeAddress(addresses, "computeOperatorImpl", address(deployedContracts.computeOperatorImpl));

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
