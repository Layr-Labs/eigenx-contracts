// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * Purpose: bootstrap a fresh chain (base) with a ProxyAdmin owned by the ops multisig.
 *
 * This leaves the env at version 1.2.0 so the shared v1.3.0-usdc-credits release
 * (from: 1.2.0) can run afterward to deploy USDCCredits against this ProxyAdmin.
 */
contract DeployProxyAdmin is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy ProxyAdmin (fresh chain, nothing exists yet)
        ProxyAdmin proxyAdmin = new ProxyAdmin();

        // Transfer ProxyAdmin ownership to ops multisig
        proxyAdmin.transferOwnership(Env.computeOpsMultisig());

        // Register in Zeus Env system
        deployContract({name: type(ProxyAdmin).name, deployedTo: address(proxyAdmin)});

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        runAsEOA();

        _validateNewAddresses();
        _validateState();
    }

    function _validateNewAddresses() internal view {
        assertTrue(address(Env.proxyAdmin()) != address(0), "ProxyAdmin is zero");
    }

    function _validateState() internal view {
        assertEq(Env.proxyAdmin().owner(), Env.computeOpsMultisig(), "ProxyAdmin owner mismatch");
    }
}
