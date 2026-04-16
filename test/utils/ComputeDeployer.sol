// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {Deploy} from "../../script/Deploy.s.sol";
import {Parser} from "../../script/Parser.s.sol";
import {IAppController} from "../../src/interfaces/IAppController.sol";
import {ComputeAVSRegistrar} from "../../src/ComputeAVSRegistrar.sol";
import {ComputeOperator} from "../../src/ComputeOperator.sol";
import {ImageAllowlist} from "../../src/ImageAllowlist.sol";
import {PermissionController} from "@eigenlayer-contracts/src/contracts/permissions/PermissionController.sol";
import {ReleaseManager} from "@eigenlayer-contracts/src/contracts/core/ReleaseManager.sol";
import {AllocationManager} from "@eigenlayer-contracts/src/contracts/core/AllocationManager.sol";
import {KeyRegistrar} from "@eigenlayer-contracts/src/contracts/permissions/KeyRegistrar.sol";
import {DelegationManager} from "@eigenlayer-contracts/src/contracts/core/DelegationManager.sol";
import {StrategyManager} from "@eigenlayer-contracts/src/contracts/core/StrategyManager.sol";
import {PauserRegistry} from "@eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";
import {EigenPodManager} from "@eigenlayer-contracts/src/contracts/pods/EigenPodManager.sol";
import {IStrategy} from "@eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {IETHPOSDeposit} from "@eigenlayer-contracts/src/contracts/interfaces/IETHPOSDeposit.sol";
import {IBeacon} from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {ISafeTimelockFactory} from "../../src/interfaces/ISafeTimelockFactory.sol";

contract ComputeDeployer is Test {
    address public developer = address(0x2);
    address public user = makeAddr("user");
    address public admin = makeAddr("admin");

    IAppController public appController;
    PermissionController public permissionController;
    ReleaseManager public releaseManager;
    AllocationManager public allocationManager;
    KeyRegistrar public keyRegistrar;
    DelegationManager public delegationManager;
    StrategyManager public strategyManager;
    PauserRegistry public pauserRegistry;
    EigenPodManager public eigenPodManager;
    ComputeAVSRegistrar public computeAVSRegistrar;
    ComputeOperator public computeOperator;
    ImageAllowlist public imageAllowlist;
    Deploy public deployer;

    string public constant VERSION = "1.0.0-test";

    function _deployContracts() internal {
        // Deploy EigenLayer contracts first
        _deployEigenLayerContracts();

        // Use the Deploy script with our EigenLayer contracts
        deployer = new Deploy();
        Parser.DeployParams memory params = Parser.DeployParams({
            version: VERSION,
            delegationManager: delegationManager,
            allocationManager: allocationManager,
            permissionController: permissionController,
            keyRegistrar: keyRegistrar,
            releaseManager: releaseManager,
            proxyAdmin: ProxyAdmin(address(0)), // Will be created by deploy script
            initialOwner: admin,
            operatorMetadataURI: "https://example.com/operator-metadata",
            avsMetadataURI: "https://example.com/avs-metadata",
            maxGlobalActiveApps: 100,
            adminMaxActiveApps: 100,
            safeSingleton: 0x41675C099F32341bf84BFc5382aF534df5C7461a,
            safeProxyFactory: 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67,
            safeFallbackHandler: 0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99,
            safeTimelockFactory: ISafeTimelockFactory(address(0)) // Deploy will create one
        });

        Parser.DeployedContracts memory deployed = deployer.deployForTesting(params);

        appController = deployed.appController;
        computeAVSRegistrar = ComputeAVSRegistrar(address(deployed.computeAVSRegistrar));
        computeOperator = ComputeOperator(address(deployed.computeOperator));
        imageAllowlist = ImageAllowlist(address(deployed.imageAllowlist));

        // Set up permission for AppController to call ReleaseManager.publishMetadataURI on behalf of ComputeAVSRegistrar
        vm.startPrank(address(computeAVSRegistrar));
        permissionController.setAppointee(
            address(computeAVSRegistrar), // account
            address(appController), // appointee
            address(releaseManager), // target (ReleaseManager)
            0x4840a67c // publishMetadataURI selector
        );
        vm.stopPrank();
    }

    function _deployEigenLayerContracts() internal {
        address[] memory pausers = new address[](1);
        pausers[0] = address(this);
        pauserRegistry = new PauserRegistry(pausers, address(this));

        // Predict addresses for EigenLayer circular dependencies
        uint256 nonce = vm.getNonce(address(this));
        address predictedDelegation = vm.computeCreateAddress(address(this), nonce + 5);

        // Deploy EigenLayer contracts
        permissionController = new PermissionController(VERSION);
        releaseManager = new ReleaseManager(permissionController, VERSION);
        allocationManager = new AllocationManager(
            DelegationManager(predictedDelegation),
            IStrategy(address(0)),
            pauserRegistry,
            permissionController,
            1 days,
            1 days,
            VERSION
        );
        strategyManager =
            new StrategyManager(allocationManager, DelegationManager(predictedDelegation), pauserRegistry, VERSION);
        eigenPodManager = new EigenPodManager(
            IETHPOSDeposit(address(0)),
            IBeacon(address(0)),
            DelegationManager(predictedDelegation),
            pauserRegistry,
            VERSION
        );
        delegationManager = new DelegationManager(
            strategyManager, eigenPodManager, allocationManager, pauserRegistry, permissionController, 1 days, VERSION
        );
        keyRegistrar = new KeyRegistrar(permissionController, allocationManager, VERSION);
    }
}
