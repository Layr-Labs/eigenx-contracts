// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {USDCDeposit} from "../../../src/USDCDeposit.sol";

/**
 * Purpose: deploy USDCDeposit contract for agent/contract wallet onboarding via USDC deposits
 */
contract DeployUSDCDeposit is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy implementation
        USDCDeposit impl = new USDCDeposit({_usdc: IERC20(Env.USDC_TOKEN()), _treasury: Env.USDC_TREASURY()});

        // Deploy proxy with initialization
        // initialize sets computeOpsMultisig as owner and configures minimum deposit
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(impl),
            address(Env.proxyAdmin()),
            abi.encodeCall(USDCDeposit.initialize, (Env.computeOpsMultisig(), Env.USDC_MINIMUM_DEPOSIT()))
        );

        // Register in Env system
        deployImpl({name: type(USDCDeposit).name, deployedTo: address(impl)});
        deployProxy({name: type(USDCDeposit).name, deployedTo: address(proxy)});

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
        assertTrue(address(Env.impl.usdcDeposit()) != address(0), "USDCDeposit impl is zero");

        // Validate proxy is non-zero
        assertTrue(address(Env.proxy.usdcDeposit()) != address(0), "USDCDeposit proxy is zero");

        // Validate proxy->impl
        assertEq(
            _getProxyImpl(address(Env.proxy.usdcDeposit())),
            address(Env.impl.usdcDeposit()),
            "USDCDeposit proxy->impl mismatch"
        );

        // Validate proxy admin
        assertEq(
            _getProxyAdmin(address(Env.proxy.usdcDeposit())),
            address(Env.proxyAdmin()),
            "USDCDeposit proxyAdmin mismatch"
        );
    }

    function _validateConstructors() internal view {
        USDCDeposit usdcDeposit = Env.impl.usdcDeposit();

        // Validate immutables
        assertEq(address(usdcDeposit.usdc()), Env.USDC_TOKEN(), "USDCDeposit usdc mismatch");
        assertEq(usdcDeposit.treasury(), Env.USDC_TREASURY(), "USDCDeposit treasury mismatch");
    }

    function _validateState() internal view {
        USDCDeposit usdcDeposit = Env.proxy.usdcDeposit();

        // Validate owner
        assertEq(usdcDeposit.owner(), Env.computeOpsMultisig(), "USDCDeposit owner mismatch");

        // Validate minimum deposit
        assertEq(usdcDeposit.minimumDeposit(), Env.USDC_MINIMUM_DEPOSIT(), "USDCDeposit minimumDeposit mismatch");
    }

    function _getProxyImpl(address proxy) internal view returns (address) {
        return Env.proxyAdmin().getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    function _getProxyAdmin(address proxy) internal view returns (address) {
        return Env.proxyAdmin().getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }
}
