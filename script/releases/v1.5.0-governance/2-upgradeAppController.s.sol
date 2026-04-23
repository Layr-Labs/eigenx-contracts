// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "./1-deployContracts.s.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * Purpose: upgrade AppController proxy to the new implementation with team RBAC and governance support
 */
contract UpgradeAppController is MultisigBuilder, DeployContracts {
    using Env for *;

    function _runAsMultisig() internal override prank(Env.computeOpsMultisig()) {
        Env.proxyAdmin()
            .upgrade(
                ITransparentUpgradeableProxy(address(Env.proxy.appController())), address(Env.impl.appController())
            );
    }

    function testScript() public virtual override {
        runAsEOA();

        execute();

        _validateNewAddresses({afterUpgrade: true});
        _validateConstructors();
    }
}
