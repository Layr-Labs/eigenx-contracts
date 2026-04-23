// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "./1-deployGovernanceContracts.s.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * @title UpgradeAppController (v1.5.0-governance phase 2)
 * @notice Multisig phase: point the AppController proxy at the new
 *         implementation deployed in phase 1. No initializer call — the
 *         impl's storage layout is append-only (new `bool timelocked` sits
 *         at byte 30 of an existing slot, previously zero on every app).
 *
 *         Once this upgrade lands, existing apps retain their prior state
 *         (creator, operatorSetId, status, billingType) and gain the
 *         timelocked flag defaulted to false. New governance features —
 *         transferOwnership, hardened sensitive-op gates, ICallValidator
 *         schedule-time checks — become active immediately.
 */
contract UpgradeAppController is MultisigBuilder, DeployGovernanceContracts {
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
