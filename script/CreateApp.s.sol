// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {stdJson} from "forge-std/StdJson.sol";
import {Script, console} from "forge-std/Script.sol";
import {Parser} from "./Parser.s.sol";
import {IReleaseManagerTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IAppController} from "../src/interfaces/IAppController.sol";
import {IApp} from "../src/interfaces/IApp.sol";

contract CreateApp is Parser {
    using stdJson for string;

    function run(string memory environment) public {
        DeployParams memory deployParams = parseDeployParams(environment);
        DeployedContracts memory deployedContracts = parseDeployedContracts(environment);

        vm.startBroadcast();
        console.log("[CreateApp] Using deployer address:", msg.sender);

        console.log("[CreateApp] Creating App via AppController contract:", address(deployedContracts.appController));
        // Deploy a new app via the app controller

        IReleaseManagerTypes.Artifact[] memory artifacts = new IReleaseManagerTypes.Artifact[](1);
        artifacts[0] = IReleaseManagerTypes.Artifact({
            digest: bytes32(0xa4d1b5030e72b2f8f96b7a21975f34f5e73787f1205bad5bc4c768c87b322192),
            registry: "index.docker.io/cavaneigen/demo"
        });

        IReleaseManagerTypes.Release memory rmsRelease =
            IReleaseManagerTypes.Release({artifacts: artifacts, upgradeByTime: uint32(block.timestamp + 1 days)});

        IAppController.Release memory release = IAppController.Release({
            rmsRelease: rmsRelease,
            publicEnv: hex"7b7d",
            encryptedEnv: hex"65794a68624763694f694a535530457454304646554330794e5459694c434a6c626d4d694f694a424d6a553252304e4e496e302e424b304f613244356d58476532363657556d445764456977517554674337695863366141486f413476314243392d7242786e6f43306d5a6e56454571383350616e6f6f525476512d6e632d364b473642527a52797069507056586b54676d365a386f6a6b4d46555337456848656e73306f2d5145417858577563625f5564694f4f5149664b4b5465492d5359564f62306662552d482d4b485a625357513257784672466d556f587a4f43384c4b7a5a55736c534a344e58776c4843687a61485867766c6f6c616549413241656c497745485a4b66596a69304267746d51424a335f6756566b6656305f507a71745f77665173733167435742785f44484a6f737a58476c514a6a63596a6f53375261725850494539364b4e766f365578615251754a72705269362d744c7666546933667a6b6f584268302d7a657452516c394f4c644e4541725a43416652645f415a53467563696d6231736e7874314b796e4b457935345347464b584278746133654d58446a47795a764438793035755251452d78646e73672d596e37674f624f71504d5278656972764c7a726b744a4f746b594c495a414a444c6a53494c4a62325045677852474c4e50765a4737764a62627534366b783839356d71695a703369714e37505f6a624c484569426f75624e66526b5154355130345a3868435231594a794a6b5030516b535a4152656d4970393235796755687735374351593745676e6e69693951664442427736734738594266386449615a744d503836673850564539614c7236653444493653586f69653537684b79337168785f3965384a7970754c793038724b4c5f66556d524d4b6f377a354164486e66364b733962375752666534584f4b6b614f546f507255574f47626e4b4d467154734c4b6657706d4b7972396f344669325176384c4539466b553231715f374b426e625f30632e696f426655376b6c7935335a6b754f6e2e3637512e564839664b73575468774f4f6b416931754642772d41"
        });

        IApp app = deployedContracts.appController.createApp(bytes32(vm.randomUint()), release);
        console.log("[CreateApp] App deployment completed successfully! App address:", address(app));

        deployParams.permissionController.acceptAdmin(address(app));
        console.log("[CreateApp] Admin role accepted for App:", address(app));

        vm.stopBroadcast();

        // Write deployment info to output file
        _writeOutputToJson(environment, app);
    }

    function _writeOutputToJson(string memory environment, IApp app) internal {
        // Add the addresses object
        string memory values = "values";
        values = vm.serializeAddress(values, "app", address(app));

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "values", values);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat("script/deploys/", environment, "/create_app_output.json");

        // Create directory if it doesn't exist
        string memory outputDir = string.concat("script/deploys/", environment);
        vm.createDir(outputDir, true);

        vm.writeJson(finalJson, outputFile);
    }
}
