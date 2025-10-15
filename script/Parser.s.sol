// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
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

contract Parser is Script {
    struct DeployParams {
        string version;
        IDelegationManager delegationManager;
        IAllocationManager allocationManager;
        IPermissionController permissionController;
        IKeyRegistrar keyRegistrar;
        IReleaseManager releaseManager;
        ProxyAdmin proxyAdmin;
        address initialOwner;
        string operatorMetadataURI;
        string avsMetadataURI;
        uint32 maxGlobalActiveApps;
        uint32 adminMaxActiveApps;
        address billingCore;
    }

    struct DeployedContracts {
        ProxyAdmin proxyAdmin;
        UpgradeableBeacon appBeacon;
        App appImpl;
        IAppController appController;
        AppController appControllerImpl;
        IComputeAVSRegistrar computeAVSRegistrar;
        ComputeAVSRegistrar computeAVSRegistrarImpl;
        IComputeOperator computeOperator;
        ComputeOperator computeOperatorImpl;
    }

    function parseDeployParams(string memory environment) public view returns (DeployParams memory) {
        string memory configPath = string.concat("script/deploys/", environment, "/config.json");
        string memory json = vm.readFile(configPath);

        // Parse billingCore, default to address(0) if not present
        address billingCoreAddr;
        try vm.parseJsonAddress(json, ".billingCore") returns (address addr) {
            billingCoreAddr = addr;
        } catch {
            billingCoreAddr = address(0);
        }

        DeployParams memory params = DeployParams({
            version: vm.parseJsonString(json, ".version"),
            delegationManager: IDelegationManager(vm.parseJsonAddress(json, ".delegationManager")),
            allocationManager: IAllocationManager(vm.parseJsonAddress(json, ".allocationManager")),
            permissionController: IPermissionController(vm.parseJsonAddress(json, ".permissionController")),
            keyRegistrar: IKeyRegistrar(vm.parseJsonAddress(json, ".keyRegistrar")),
            releaseManager: IReleaseManager(vm.parseJsonAddress(json, ".releaseManager")),
            proxyAdmin: ProxyAdmin(vm.parseJsonAddress(json, ".proxyAdmin")),
            initialOwner: vm.parseJsonAddress(json, ".initialOwner"),
            operatorMetadataURI: vm.parseJsonString(json, ".operatorMetadataURI"),
            avsMetadataURI: vm.parseJsonString(json, ".avsMetadataURI"),
            maxGlobalActiveApps: uint32(vm.parseJsonUint(json, ".maxGlobalActiveApps")),
            adminMaxActiveApps: uint32(vm.parseJsonUint(json, ".adminMaxActiveApps")),
            billingCore: billingCoreAddr
        });

        return params;
    }

    function parseDeployedContracts(string memory environment) public view returns (DeployedContracts memory) {
        string memory outputPath = string.concat("script/deploys/", environment, "/output.json");
        string memory json = vm.readFile(outputPath);

        DeployedContracts memory deployedContracts = DeployedContracts({
            proxyAdmin: ProxyAdmin(vm.parseJsonAddress(json, ".addresses.proxyAdmin")),
            appBeacon: UpgradeableBeacon(vm.parseJsonAddress(json, ".addresses.appBeacon")),
            appImpl: App(vm.parseJsonAddress(json, ".addresses.appImpl")),
            appController: IAppController(vm.parseJsonAddress(json, ".addresses.appController")),
            appControllerImpl: AppController(vm.parseJsonAddress(json, ".addresses.appControllerImpl")),
            computeAVSRegistrar: IComputeAVSRegistrar(vm.parseJsonAddress(json, ".addresses.computeAVSRegistrar")),
            computeAVSRegistrarImpl: ComputeAVSRegistrar(vm.parseJsonAddress(json, ".addresses.computeAVSRegistrarImpl")),
            computeOperator: IComputeOperator(vm.parseJsonAddress(json, ".addresses.computeOperator")),
            computeOperatorImpl: ComputeOperator(vm.parseJsonAddress(json, ".addresses.computeOperatorImpl"))
        });

        return deployedContracts;
    }
}
