// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "./1-deployGovernanceContracts.s.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * @title UpgradeAppController
 * @notice Second step of the v1.5.0 release. Run by the ops multisig after
 *         the EOA-deployed impls from `1-deployGovernanceContracts.s.sol`
 *         are on-chain. Points the AppController proxy at the new impl via
 *         `ProxyAdmin.upgrade`. No initializer call — the new impl's
 *         storage layout is append-only vs. v1.4.0, so every existing app
 *         keeps its prior state (creator, operatorSetId, status, billingType)
 *         with no rewrites required.
 *
 *         Post-upgrade, AppController delegates auth to AppAuthority (also
 *         deployed in phase 1). Critical ops are owner-gated; operational
 *         roles (PAUSER, DEVELOPER) can be granted via AppAuthority. A
 *         follow-up call to `migrateAppsToAppAuthority` seeds AppAuthority
 *         state for existing apps — without it, auth checks on pre-upgrade
 *         apps revert because AppAuthority has no owner recorded.
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
