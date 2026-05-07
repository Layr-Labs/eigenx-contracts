// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "./1-deployGovernanceContracts.s.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {IApp} from "../../../src/interfaces/IApp.sol";
import {IAppController} from "../../../src/interfaces/IAppController.sol";

/**
 * @title UpgradeAppController
 * @notice Second step of the v1.5.0 release. Run by the ops multisig after
 *         the EOA-deployed impls from `1-deployGovernanceContracts.s.sol`
 *         are on-chain.
 *
 *         Performs two actions atomically in a single multisig transaction:
 *         1. `ProxyAdmin.upgrade` points AppController at the new impl. From
 *            this call on, AppController delegates auth to AppAuthority.
 *         2. `migrateAppsToAppAuthority` seeds AppAuthority state for every
 *            pre-existing app: initializes each scope from the cached
 *            `creator`, seeds ADMIN from the app's legacy PermissionController
 *            admins. Paginated in batches to stay under the block gas limit.
 *
 *         Atomicity matters: between steps 1 and 2, pre-existing apps have
 *         `appAuthority.scopeOwner(app) == address(0)`, so every owner-gated
 *         critical op would revert. Running both in one multisig tx closes
 *         that window entirely.
 *
 *         The storage layout is append-only vs. v1.4.0 plus the
 *         `pendingReleaseBlockNumber` field added in v1.4.1 — existing apps
 *         keep their prior state (creator, operatorSetId, status,
 *         billingType, latestReleaseBlockNumber) with no rewrites required.
 */
contract UpgradeAppController is MultisigBuilder, DeployGovernanceContracts {
    using Env for *;

    /// @notice Page size for the post-upgrade migration loop. Conservative
    ///         default that keeps the outer multisig tx well under block
    ///         gas limits even with dozens of admins per app.
    uint256 internal constant MIGRATION_PAGE_SIZE = 50;

    function _runAsMultisig() internal override prank(Env.computeOpsMultisig()) {
        // Step 1 — upgrade the AppController proxy to the new impl.
        Env.proxyAdmin()
            .upgrade(
                ITransparentUpgradeableProxy(address(Env.proxy.appController())), address(Env.impl.appController())
            );

        // Step 2 — seed AppAuthority state for every existing app.
        // `getApps` is an unchanged v1.4.0 view; it enumerates the full
        // `_allApps` set which the upgrade does not touch. Paginate to
        // avoid one giant calldata/gas spike.
        IAppController controller = IAppController(address(Env.proxy.appController()));
        uint256 offset = 0;
        while (true) {
            (IApp[] memory apps,) = controller.getApps(offset, MIGRATION_PAGE_SIZE);
            if (apps.length == 0) break;
            controller.migrateAppsToAppAuthority(apps);
            offset += apps.length;
            if (apps.length < MIGRATION_PAGE_SIZE) break;
        }
    }

    function testScript() public virtual override {
        runAsEOA();
        execute();
        _validateNewAddresses({afterUpgrade: true});
        _validateConstructors();
    }
}
