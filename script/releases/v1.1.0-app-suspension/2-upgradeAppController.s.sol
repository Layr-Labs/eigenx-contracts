// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import "./1-deployAppControllerImpl.s.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * Purpose: upgrade AppController proxy to new implementation via ProxyAdmin
 */
contract UpgradeAppController is MultisigBuilder, DeployAppControllerImpl {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal virtual override prank(Env.computeOpsMultisig()) {
        Env.proxyAdmin()
            .upgrade(
                ITransparentUpgradeableProxy(address(Env.proxy.appController())), address(Env.impl.appController())
            );
    }

    function testScript() public virtual override {
        runAsEOA();

        execute();

        _validateNewAddresses({afterUpgrade: true});
    }
}
