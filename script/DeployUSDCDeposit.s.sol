// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {
    TransparentUpgradeableProxy,
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {EmptyContract} from "@eigenlayer-contracts/src/test/mocks/EmptyContract.sol";
import {USDCDeposit} from "../src/USDCDeposit.sol";
import {IUSDCDeposit} from "../src/interfaces/IUSDCDeposit.sol";

contract DeployUSDCDeposit is Script {
    struct DeployParams {
        string version;
        IPermissionController permissionController;
        ProxyAdmin proxyAdmin;
        address initialOwner;
        IERC20 usdc;
        address treasury;
        uint256 minimumDeposit;
    }

    struct DeployedContracts {
        IUSDCDeposit usdcDeposit;
        USDCDeposit usdcDepositImpl;
    }

    function run(string memory environment) public {
        DeployParams memory params = parseDeployParams(environment);

        vm.startBroadcast();
        DeployedContracts memory deployed = deploy(params);
        _writeOutputToJson(environment, deployed);
        vm.stopBroadcast();
    }

    function deploy(DeployParams memory params) public returns (DeployedContracts memory) {
        require(bytes(params.version).length != 0, "Version must not be empty");
        require(address(params.permissionController) != address(0), "Permission controller must not be empty");
        require(address(params.proxyAdmin) != address(0), "Proxy admin must not be empty");
        require(address(params.initialOwner) != address(0), "Initial owner must not be empty");
        require(address(params.usdc) != address(0), "USDC address must not be empty");
        require(params.treasury != address(0), "Treasury address must not be empty");

        EmptyContract emptyContract = new EmptyContract();

        // Deploy proxy
        TransparentUpgradeableProxy proxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(params.proxyAdmin), new bytes(0));

        // Deploy implementation
        USDCDeposit impl = new USDCDeposit({
            _version: params.version,
            _permissionController: params.permissionController,
            _usdc: params.usdc,
            _treasury: params.treasury
        });

        // Upgrade and initialize
        params.proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(proxy)),
            address(impl),
            abi.encodeCall(USDCDeposit.initialize, (params.initialOwner, params.minimumDeposit))
        );

        // Accept admin role
        IPermissionController(address(params.permissionController)).acceptAdmin(address(proxy));

        console.log("USDCDeposit proxy:", address(proxy));
        console.log("USDCDeposit impl:", address(impl));

        return DeployedContracts({
            usdcDeposit: IUSDCDeposit(address(proxy)),
            usdcDepositImpl: impl
        });
    }

    function deployForTesting(DeployParams memory params) public returns (DeployedContracts memory deployment) {
        vm.startPrank(params.initialOwner);
        deployment = deploy(params);
        vm.stopPrank();
    }

    function parseDeployParams(string memory environment) public view returns (DeployParams memory) {
        string memory configPath = string.concat("script/deploys/", environment, "/usdc-deposit-config.json");
        string memory json = vm.readFile(configPath);

        return DeployParams({
            version: vm.parseJsonString(json, ".version"),
            permissionController: IPermissionController(vm.parseJsonAddress(json, ".permissionController")),
            proxyAdmin: ProxyAdmin(vm.parseJsonAddress(json, ".proxyAdmin")),
            initialOwner: vm.parseJsonAddress(json, ".initialOwner"),
            usdc: IERC20(vm.parseJsonAddress(json, ".usdc")),
            treasury: vm.parseJsonAddress(json, ".treasury"),
            minimumDeposit: vm.parseJsonUint(json, ".minimumDeposit")
        });
    }

    function _writeOutputToJson(string memory environment, DeployedContracts memory deployedContracts) internal {
        string memory addresses = "addresses";
        vm.serializeAddress(addresses, "usdcDeposit", address(deployedContracts.usdcDeposit));
        addresses = vm.serializeAddress(addresses, "usdcDepositImpl", address(deployedContracts.usdcDepositImpl));

        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        string memory outputDir = string.concat("script/deploys/", environment);
        vm.createDir(outputDir, true);

        string memory outputFile = string.concat(outputDir, "/usdc-deposit-output.json");
        vm.writeJson(finalJson, outputFile);
    }
}
