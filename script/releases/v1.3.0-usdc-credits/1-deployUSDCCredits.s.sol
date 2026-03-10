// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {USDCCredits} from "../../../src/USDCCredits.sol";

/**
 * Purpose: deploy USDCCredits contract for agent/contract wallet onboarding via USDC credit purchases
 */
contract DeployUSDCCredits is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy implementation
        USDCCredits impl = new USDCCredits({_usdc: IERC20(Env.USDC_TOKEN()), _treasury: Env.USDC_TREASURY()});

        // Deploy proxy with initialization
        // initialize sets computeOpsMultisig as owner and configures minimum purchase
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(impl),
            address(Env.proxyAdmin()),
            abi.encodeCall(USDCCredits.initialize, (Env.computeOpsMultisig(), Env.USDC_MINIMUM_PURCHASE()))
        );

        // Register in Env system
        deployImpl({name: type(USDCCredits).name, deployedTo: address(impl)});
        deployProxy({name: type(USDCCredits).name, deployedTo: address(proxy)});

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        runAsEOA();

        _validateNewAddresses();
        _validateConstructors();
        _validateState();
    }

    function _validateNewAddresses() internal view {
        // Validate impl is non-zero
        assertTrue(address(Env.impl.usdcCredits()) != address(0), "USDCCredits impl is zero");

        // Validate proxy is non-zero
        assertTrue(address(Env.proxy.usdcCredits()) != address(0), "USDCCredits proxy is zero");

        // Validate proxy->impl
        assertEq(
            _getProxyImpl(address(Env.proxy.usdcCredits())),
            address(Env.impl.usdcCredits()),
            "USDCCredits proxy->impl mismatch"
        );

        // Validate proxy admin
        assertEq(
            _getProxyAdmin(address(Env.proxy.usdcCredits())),
            address(Env.proxyAdmin()),
            "USDCCredits proxyAdmin mismatch"
        );
    }

    function _validateConstructors() internal view {
        USDCCredits usdcCredits = Env.impl.usdcCredits();

        // Validate immutables
        assertEq(address(usdcCredits.usdc()), Env.USDC_TOKEN(), "USDCCredits usdc mismatch");
        assertEq(usdcCredits.treasury(), Env.USDC_TREASURY(), "USDCCredits treasury mismatch");
    }

    function _validateState() internal view {
        USDCCredits usdcCredits = Env.proxy.usdcCredits();

        // Validate owner
        assertEq(usdcCredits.owner(), Env.computeOpsMultisig(), "USDCCredits owner mismatch");

        // Validate minimum purchase
        assertEq(usdcCredits.minimumPurchase(), Env.USDC_MINIMUM_PURCHASE(), "USDCCredits minimumPurchase mismatch");
    }

    function _getProxyImpl(address proxy) internal view returns (address) {
        return Env.proxyAdmin().getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    function _getProxyAdmin(address proxy) internal view returns (address) {
        return Env.proxyAdmin().getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }
}
