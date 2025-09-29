// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Vm.sol";
import "zeus-templates/utils/ZEnvHelpers.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {DelegationManager} from "@eigenlayer-contracts/src/contracts/core/DelegationManager.sol";
import {AllocationManager} from "@eigenlayer-contracts/src/contracts/core/AllocationManager.sol";
import {PermissionController} from "@eigenlayer-contracts/src/contracts/permissions/PermissionController.sol";
import {KeyRegistrar} from "@eigenlayer-contracts/src/contracts/permissions/KeyRegistrar.sol";
import {ReleaseManager} from "@eigenlayer-contracts/src/contracts/core/ReleaseManager.sol";

import {App} from "../../src/App.sol";
import {AppController} from "../../src/AppController.sol";
import {ComputeAVSRegistrar} from "../../src/ComputeAVSRegistrar.sol";
import {ComputeOperator} from "../../src/ComputeOperator.sol";

library Env {
    using ZEnvHelpers for *;

    /// Dummy types and variables to facilitate syntax, e.g: `Env.proxy.appController()`
    enum DeployedProxy {
        A
    }
    enum DeployedBeacon {
        A
    }
    enum DeployedImpl {
        A
    }
    enum DeployedInstance {
        A
    }

    DeployedProxy internal constant proxy = DeployedProxy.A;
    DeployedBeacon internal constant beacon = DeployedBeacon.A;
    DeployedImpl internal constant impl = DeployedImpl.A;
    DeployedInstance internal constant instance = DeployedInstance.A;

    /**
     * env
     */
    function env() internal view returns (string memory) {
        return _string("ZEUS_ENV");
    }

    function envVersion() internal view returns (string memory) {
        return _string("ZEUS_ENV_VERSION");
    }

    function deployVersion() internal view returns (string memory) {
        return _string("ZEUS_DEPLOY_TO_VERSION");
    }

    /**
     * eigenlayer contracts
     */
    function delegationManager() internal view returns (DelegationManager) {
        return DelegationManager(_envAddress(type(DelegationManager).name));
    }

    function allocationManager() internal view returns (AllocationManager) {
        return AllocationManager(_envAddress(type(AllocationManager).name));
    }

    function permissionController() internal view returns (PermissionController) {
        return PermissionController(_envAddress(type(PermissionController).name));
    }

    function keyRegistrar() internal view returns (KeyRegistrar) {
        return KeyRegistrar(_envAddress(type(KeyRegistrar).name));
    }

    function releaseManager() internal view returns (ReleaseManager) {
        return ReleaseManager(_envAddress(type(ReleaseManager).name));
    }

    /**
     * compute contracts
     */
    function app(DeployedProxy) internal view returns (App) {
        return App(_deployedProxy(type(App).name));
    }

    function appController(DeployedProxy) internal view returns (AppController) {
        return AppController(_deployedProxy(type(AppController).name));
    }

    function computeAVSRegistrar(DeployedProxy) internal view returns (ComputeAVSRegistrar) {
        return ComputeAVSRegistrar(_deployedProxy(type(ComputeAVSRegistrar).name));
    }

    function computeOperator(DeployedProxy) internal view returns (ComputeOperator) {
        return ComputeOperator(_deployedProxy(type(ComputeOperator).name));
    }

    function appBeacon(DeployedBeacon) internal view returns (UpgradeableBeacon) {
        return UpgradeableBeacon(_deployedBeacon(type(App).name));
    }

    /**
     * implementation contracts
     */
    function app(DeployedImpl) internal view returns (App) {
        return App(_deployedImpl(type(App).name));
    }

    function appController(DeployedImpl) internal view returns (AppController) {
        return AppController(_deployedImpl(type(AppController).name));
    }

    function computeAVSRegistrar(DeployedImpl) internal view returns (ComputeAVSRegistrar) {
        return ComputeAVSRegistrar(_deployedImpl(type(ComputeAVSRegistrar).name));
    }

    function computeOperator(DeployedImpl) internal view returns (ComputeOperator) {
        return ComputeOperator(_deployedImpl(type(ComputeOperator).name));
    }

    /**
     * governance contracts
     */
    function computeWhitelisterMultisig() internal view returns (address) {
        return _envAddress("computeWhitelisterMultisig");
    }

    function computeOpsMultisig() internal view returns (address) {
        return _envAddress("computeOpsMultisig");
    }

    function communityMultisig() internal view returns (address) {
        return _envAddress("communityMultisig");
    }

    function executorMultisig() internal view returns (address) {
        return _envAddress("executorMultisig");
    }

    function proxyAdmin() internal view returns (ProxyAdmin) {
        return ProxyAdmin(_deployedContract(type(ProxyAdmin).name));
    }

    /**
     * helpers
     */
    function multiSendCallOnly() internal view returns (address) {
        return _envAddress("MultiSendCallOnly");
    }

    /**
     * compute config parameters
     */
    function MAX_GLOBAL_ACTIVE_APPS() internal view returns (uint32) {
        return _envU32("MAX_GLOBAL_ACTIVE_APPS");
    }

    function OPERATOR_METADATA_URI() internal view returns (string memory) {
        return _envString("OPERATOR_METADATA_URI");
    }

    function AVS_METADATA_URI() internal view returns (string memory) {
        return _envString("AVS_METADATA_URI");
    }

    /**
     * Helpers
     */
    function _deployedContract(string memory name) private view returns (address) {
        return ZEnvHelpers.state().deployedContract(name);
    }

    function _deployedInstance(string memory name, uint256 idx) private view returns (address) {
        return ZEnvHelpers.state().deployedInstance(name, idx);
    }

    function _deployedInstanceCount(string memory name) private view returns (uint256) {
        return ZEnvHelpers.state().deployedInstanceCount(name);
    }

    function _deployedProxy(string memory name) private view returns (address) {
        return ZEnvHelpers.state().deployedProxy(name);
    }

    function _deployedBeacon(string memory name) private view returns (address) {
        return ZEnvHelpers.state().deployedBeacon(name);
    }

    function _deployedImpl(string memory name) private view returns (address) {
        return ZEnvHelpers.state().deployedImpl(name);
    }

    function _envAddress(string memory key) private view returns (address) {
        return ZEnvHelpers.state().envAddress(key);
    }

    function _envU256(string memory key) private view returns (uint256) {
        return ZEnvHelpers.state().envU256(key);
    }

    function _envU64(string memory key) private view returns (uint64) {
        return ZEnvHelpers.state().envU64(key);
    }

    function _envU32(string memory key) private view returns (uint32) {
        return ZEnvHelpers.state().envU32(key);
    }

    function _envU16(string memory key) private view returns (uint16) {
        return ZEnvHelpers.state().envU16(key);
    }

    function _envBool(string memory key) private view returns (bool) {
        return ZEnvHelpers.state().envBool(key);
    }

    function _envString(string memory key) private view returns (string memory) {
        return ZEnvHelpers.state().envString(key);
    }

    address internal constant VM_ADDRESS = address(uint160(uint256(keccak256("hevm cheat code"))));
    Vm internal constant vm = Vm(VM_ADDRESS);

    function _string(string memory key) private view returns (string memory) {
        return vm.envString(key);
    }

    function _strEq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}
