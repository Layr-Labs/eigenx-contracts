// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {ImageAllowlist} from "../../../src/ImageAllowlist.sol";

/**
 * Purpose: deploy ImageAllowlist contract for CVM image validation
 */
contract DeployImageAllowlist is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy implementation
        ImageAllowlist impl = new ImageAllowlist();

        // Deploy proxy with initialization
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(impl),
            address(Env.proxyAdmin()),
            abi.encodeCall(ImageAllowlist.initialize, (Env.computeOpsMultisig()))
        );

        // Register in Env system
        deployImpl({name: type(ImageAllowlist).name, deployedTo: address(impl)});
        deployProxy({name: type(ImageAllowlist).name, deployedTo: address(proxy)});

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        runAsEOA();

        _validateNewAddresses();
        _validateState();
    }

    function _validateNewAddresses() internal view {
        // Validate impl is non-zero
        assertTrue(address(Env.impl.imageAllowlist()) != address(0), "ImageAllowlist impl is zero");

        // Validate proxy is non-zero
        assertTrue(address(Env.proxy.imageAllowlist()) != address(0), "ImageAllowlist proxy is zero");

        // Validate proxy->impl
        assertEq(
            _getProxyImpl(address(Env.proxy.imageAllowlist())),
            address(Env.impl.imageAllowlist()),
            "ImageAllowlist proxy->impl mismatch"
        );

        // Validate proxy admin
        assertEq(
            _getProxyAdmin(address(Env.proxy.imageAllowlist())),
            address(Env.proxyAdmin()),
            "ImageAllowlist proxyAdmin mismatch"
        );
    }

    function _validateState() internal view {
        ImageAllowlist imageAllowlist = Env.proxy.imageAllowlist();

        // Validate owner
        assertEq(imageAllowlist.owner(), Env.computeOpsMultisig(), "ImageAllowlist owner mismatch");
    }

    function _getProxyImpl(address proxy) internal view returns (address) {
        return Env.proxyAdmin().getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    function _getProxyAdmin(address proxy) internal view returns (address) {
        return Env.proxyAdmin().getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }
}
