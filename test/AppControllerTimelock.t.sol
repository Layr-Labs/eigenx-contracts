// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IAppController} from "../src/interfaces/IAppController.sol";
import {AppController} from "../src/AppController.sol";
import {ComputeDeployer} from "./utils/ComputeDeployer.sol";
import {IApp} from "../src/interfaces/IApp.sol";
import {PermissionControllerMixin} from "@eigenlayer-contracts/src/contracts/mixins/PermissionControllerMixin.sol";
import {IReleaseManagerTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";

contract AppControllerTimelockTest is ComputeDeployer {
    // Event definitions for testing
    event UpgradeTimelockSet(IApp indexed app, uint40 delay, bool locked);
    event UpgradeQueued(IApp indexed app, bytes32 indexed releaseHash, uint40 readyAt);
    event UpgradeCancelled(IApp indexed app, bytes32 indexed releaseHash);
    event AppUpgraded(IApp indexed app, uint256 rmsReleaseId, IAppController.Release release);

    IApp public app;
    uint40 public constant TIMELOCK_DELAY = 2 days;

    function setUp() public {
        _deployContracts();
        // Set maximum active apps for the developer account used in tests
        vm.prank(admin);
        appController.setMaxActiveAppsPerUser(developer, 100);

        // Create and accept admin on a test app
        vm.prank(developer);
        app = appController.createApp(keccak256("timelock_test"), _assembleRelease());
        vm.prank(developer);
        permissionController.acceptAdmin(address(app));
    }

    // ========== setUpgradeTimelock Tests ==========

    function test_setUpgradeTimelock_setsDelayAndEmitsEvent() public {
        vm.expectEmit(true, false, false, true);
        emit UpgradeTimelockSet(app, TIMELOCK_DELAY, false);

        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.UpgradeTimelockConfig memory config = appController.getUpgradeTimelockConfig(app);
        assertEq(config.delay, TIMELOCK_DELAY);
        assertFalse(config.locked);
    }

    function test_setUpgradeTimelock_lockPreventsDelayReduction() public {
        // Set and lock the timelock
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, true);

        IAppController.UpgradeTimelockConfig memory config = appController.getUpgradeTimelockConfig(app);
        assertTrue(config.locked);

        // Try to reduce delay - should revert
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.TimelockDelayLocked.selector));
        appController.setUpgradeTimelock(app, 1 days, false);
    }

    function test_setUpgradeTimelock_lockAllowsDelayIncrease() public {
        // Set and lock the timelock
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, true);

        // Increase delay - should succeed
        vm.prank(developer);
        appController.setUpgradeTimelock(app, 5 days, false);

        IAppController.UpgradeTimelockConfig memory config = appController.getUpgradeTimelockConfig(app);
        assertEq(config.delay, 5 days);
        assertTrue(config.locked); // Lock persists even when lock=false is passed
    }

    function test_setUpgradeTimelock_lockIsIrrevocable() public {
        // Lock the timelock
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, true);

        // Try to set same delay with lock=false - lock persists
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.UpgradeTimelockConfig memory config = appController.getUpgradeTimelockConfig(app);
        assertTrue(config.locked);
    }

    function test_setUpgradeTimelock_canDisableWhenNotLocked() public {
        // Set a delay without locking
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        // Reduce to 0 (disable) - should succeed since not locked
        vm.prank(developer);
        appController.setUpgradeTimelock(app, 0, false);

        IAppController.UpgradeTimelockConfig memory config = appController.getUpgradeTimelockConfig(app);
        assertEq(config.delay, 0);
    }

    function test_setUpgradeTimelock_notAuthorized() public {
        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);
    }

    function test_setUpgradeTimelock_nonExistentAppReverts() public {
        IApp fakeApp = IApp(address(0xDEADBEEF));
        vm.prank(address(0xDEADBEEF));
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.setUpgradeTimelock(fakeApp, TIMELOCK_DELAY, false);
    }

    // ========== upgradeApp with Timelock Tests ==========

    function test_upgradeApp_withTimelockConfigured_reverts() public {
        // Enable timelock
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        // Direct upgradeApp should revert
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.UpgradeMustBeQueued.selector));
        appController.upgradeApp(app, _assembleRelease());
    }

    function test_upgradeApp_withoutTimelock_succeeds() public {
        // No timelock configured - direct upgrade should work
        vm.prank(developer);
        appController.upgradeApp(app, _assembleRelease());

        assertEq(appController.getAppLatestReleaseBlockNumber(app), block.number);
    }

    // ========== queueUpgradeApp Tests ==========

    function test_queueUpgradeApp_queuesAndEmitsEvent() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();
        bytes32 expectedHash = keccak256(abi.encode(release));
        uint40 expectedReadyAt = uint40(block.timestamp) + TIMELOCK_DELAY;

        vm.expectEmit(true, true, false, true);
        emit UpgradeQueued(app, expectedHash, expectedReadyAt);

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        assertEq(readyAt, expectedReadyAt);

        IAppController.PendingUpgrade memory pending = appController.getPendingUpgrade(app);
        assertEq(pending.releaseHash, expectedHash);
        assertEq(pending.readyAt, expectedReadyAt);
    }

    function test_queueUpgradeApp_noTimelockConfigured_reverts() public {
        // No timelock set - should revert
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.TimelockNotConfigured.selector));
        appController.queueUpgradeApp(app, _assembleRelease());
    }

    function test_queueUpgradeApp_alreadyQueued_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        vm.prank(developer);
        appController.queueUpgradeApp(app, _assembleRelease());

        // Second queue should revert
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.UpgradeAlreadyQueued.selector));
        appController.queueUpgradeApp(app, _assembleRelease());
    }

    function test_queueUpgradeApp_notActive_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        // Terminate the app
        vm.prank(developer);
        appController.terminateApp(app);

        // Queue should revert since app is terminated
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.queueUpgradeApp(app, _assembleRelease());
    }

    function test_queueUpgradeApp_notAuthorized_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.queueUpgradeApp(app, _assembleRelease());
    }

    // ========== executeQueuedUpgrade Tests ==========

    function test_executeQueuedUpgrade_afterDelay_succeeds() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Warp past the delay
        vm.warp(readyAt);

        vm.prank(developer);
        uint256 releaseId = appController.executeQueuedUpgrade(app, release);

        // Verify upgrade occurred
        assertGt(releaseId, 0);
        assertEq(appController.getAppLatestReleaseBlockNumber(app), block.number);

        // Verify pending upgrade was cleared
        IAppController.PendingUpgrade memory pending = appController.getPendingUpgrade(app);
        assertEq(pending.readyAt, 0);
        assertEq(pending.releaseHash, bytes32(0));
    }

    function test_executeQueuedUpgrade_beforeDelay_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Warp to just before readyAt
        vm.warp(readyAt - 1);

        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.TimelockNotExpired.selector));
        appController.executeQueuedUpgrade(app, release);
    }

    function test_executeQueuedUpgrade_afterExpiry_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Warp past expiry window (readyAt + 7 days)
        vm.warp(uint256(readyAt) + AppController(address(appController)).UPGRADE_EXPIRY_WINDOW());

        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.TimelockExpired.selector));
        appController.executeQueuedUpgrade(app, release);
    }

    function test_executeQueuedUpgrade_hashMismatch_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        vm.warp(readyAt);

        // Try to execute with a different release
        IAppController.Release memory differentRelease = _assembleReleaseWithDigest(keccak256("different"));

        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.ReleaseHashMismatch.selector));
        appController.executeQueuedUpgrade(app, differentRelease);
    }

    function test_executeQueuedUpgrade_noPending_reverts() public {
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.NoPendingUpgrade.selector));
        appController.executeQueuedUpgrade(app, _assembleRelease());
    }

    function test_executeQueuedUpgrade_appTerminated_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Terminate the app by admin
        vm.prank(admin);
        appController.terminateAppByAdmin(app);

        vm.warp(readyAt);

        // Execute should fail because app is terminated
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.executeQueuedUpgrade(app, release);
    }

    function test_executeQueuedUpgrade_appSuspended_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Suspend the account
        IApp[] memory apps = new IApp[](1);
        apps[0] = app;
        vm.prank(admin);
        appController.suspend(developer, apps);

        vm.warp(readyAt);

        // Execute should fail because app is suspended
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.executeQueuedUpgrade(app, release);
    }

    function test_executeQueuedUpgrade_notAuthorized_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        vm.warp(readyAt);

        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.executeQueuedUpgrade(app, release);
    }

    // ========== cancelQueuedUpgrade Tests ==========

    function test_cancelQueuedUpgrade_succeeds() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();
        bytes32 expectedHash = keccak256(abi.encode(release));

        vm.prank(developer);
        appController.queueUpgradeApp(app, release);

        vm.expectEmit(true, true, false, true);
        emit UpgradeCancelled(app, expectedHash);

        vm.prank(developer);
        appController.cancelQueuedUpgrade(app);

        // Verify pending upgrade was cleared
        IAppController.PendingUpgrade memory pending = appController.getPendingUpgrade(app);
        assertEq(pending.readyAt, 0);
    }

    function test_cancelQueuedUpgrade_noPending_reverts() public {
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.NoPendingUpgrade.selector));
        appController.cancelQueuedUpgrade(app);
    }

    function test_cancelQueuedUpgrade_notAuthorized_reverts() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        vm.prank(developer);
        appController.queueUpgradeApp(app, _assembleRelease());

        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.cancelQueuedUpgrade(app);
    }

    function test_cancelQueuedUpgrade_thenRequeue() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        // Queue, cancel, then re-queue
        vm.prank(developer);
        appController.queueUpgradeApp(app, _assembleRelease());

        vm.prank(developer);
        appController.cancelQueuedUpgrade(app);

        // Re-queue should succeed
        IAppController.Release memory newRelease = _assembleReleaseWithDigest(keccak256("new-digest"));

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, newRelease);

        IAppController.PendingUpgrade memory pending = appController.getPendingUpgrade(app);
        assertEq(pending.readyAt, readyAt);
        assertEq(pending.releaseHash, keccak256(abi.encode(newRelease)));
    }

    // ========== Lifecycle Integration Tests ==========

    function test_lifecycle_fullTimelockFlow() public {
        // 1. Enable timelock
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, true);

        // 2. Direct upgrade fails
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.UpgradeMustBeQueued.selector));
        appController.upgradeApp(app, _assembleRelease());

        // 3. Queue upgrade
        IAppController.Release memory release = _assembleRelease();
        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // 4. Premature execute fails
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.TimelockNotExpired.selector));
        appController.executeQueuedUpgrade(app, release);

        // 5. Execute after delay succeeds
        vm.warp(readyAt);
        vm.prank(developer);
        appController.executeQueuedUpgrade(app, release);

        assertEq(appController.getAppLatestReleaseBlockNumber(app), block.number);
    }

    function test_lifecycle_createAppUnaffectedByTimelock() public {
        // Create a new app (the initial release in createApp bypasses the timelock)
        vm.prank(developer);
        IApp newApp = appController.createApp(keccak256("new_app"), _assembleRelease());

        // Verify it was created and started successfully
        assertEq(uint256(appController.getAppStatus(newApp)), uint256(IAppController.AppStatus.STARTED));
    }

    function test_lifecycle_stoppedAppCanQueueAndExecute() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        // Stop the app
        vm.prank(developer);
        appController.stopApp(app);

        // Queue upgrade on stopped app
        IAppController.Release memory release = _assembleRelease();
        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Execute after delay
        vm.warp(readyAt);
        vm.prank(developer);
        appController.executeQueuedUpgrade(app, release);

        // App should still be stopped (upgrade doesn't change status)
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STOPPED));
    }

    function test_lifecycle_cancelAfterDelayPassed() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Warp past delay - cancel should still work
        vm.warp(readyAt + 1 days);

        vm.prank(developer);
        appController.cancelQueuedUpgrade(app);

        IAppController.PendingUpgrade memory pending = appController.getPendingUpgrade(app);
        assertEq(pending.readyAt, 0);
    }

    function test_lifecycle_delayChangeDoesNotAffectQueuedAction() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Increase the delay AFTER queuing
        vm.prank(developer);
        appController.setUpgradeTimelock(app, 30 days, false);

        // The already-queued action should still execute at original readyAt
        vm.warp(readyAt);

        vm.prank(developer);
        appController.executeQueuedUpgrade(app, release);

        assertEq(appController.getAppLatestReleaseBlockNumber(app), block.number);
    }

    function test_lifecycle_terminatedAppActionNotExecutable() public {
        vm.prank(developer);
        appController.setUpgradeTimelock(app, TIMELOCK_DELAY, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        uint40 readyAt = appController.queueUpgradeApp(app, release);

        // Admin terminates the app
        vm.prank(admin);
        appController.terminateAppByAdmin(app);

        vm.warp(readyAt);

        // Execute fails (app is terminated)
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.executeQueuedUpgrade(app, release);

        // But cancel still works (doesn't require appIsActive)
        vm.prank(developer);
        appController.cancelQueuedUpgrade(app);
    }

    // ========== View Function Tests ==========

    function test_getUpgradeTimelockConfig_defaultIsZero() public view {
        IAppController.UpgradeTimelockConfig memory config = appController.getUpgradeTimelockConfig(app);
        assertEq(config.delay, 0);
        assertFalse(config.locked);
    }

    function test_getPendingUpgrade_defaultIsZero() public view {
        IAppController.PendingUpgrade memory pending = appController.getPendingUpgrade(app);
        assertEq(pending.readyAt, 0);
        assertEq(pending.releaseHash, bytes32(0));
    }

    // ========== Security Edge Cases ==========

    /// @notice Verify that locking at delay=0 is prevented (would permanently brick timelock)
    function test_setUpgradeTimelock_cannotLockZeroDelay() public {
        vm.prank(developer);
        vm.expectRevert(IAppController.CannotLockZeroDelay.selector);
        appController.setUpgradeTimelock(app, 0, true);
    }

    /// @notice Verify that re-executing after a successful execute reverts with NoPendingUpgrade
    function test_executeQueuedUpgrade_doubleExecute_reverts() public {
        uint40 delay = 2 days;
        vm.prank(developer);
        appController.setUpgradeTimelock(app, delay, false);

        IAppController.Release memory release = _assembleRelease();

        vm.prank(developer);
        appController.queueUpgradeApp(app, release);

        vm.warp(block.timestamp + delay);

        // First execution succeeds
        vm.prank(developer);
        appController.executeQueuedUpgrade(app, release);

        // Second execution reverts
        vm.prank(developer);
        vm.expectRevert(IAppController.NoPendingUpgrade.selector);
        appController.executeQueuedUpgrade(app, release);
    }

    /// @notice Verify that two different apps have independent timelock configs
    function test_setUpgradeTimelock_independentPerApp() public {
        // Create a second app
        IAppController.Release memory release = _assembleRelease();
        vm.prank(developer);
        IApp app2 = appController.createApp(bytes32(uint256(2)), release);

        // Accept admin for app2
        vm.prank(developer);
        permissionController.acceptAdmin(address(app2));

        // Set different delays for each app
        vm.prank(developer);
        appController.setUpgradeTimelock(app, 1 days, false);

        vm.prank(developer);
        appController.setUpgradeTimelock(app2, 7 days, true);

        // Verify they are independent
        IAppController.UpgradeTimelockConfig memory config1 = appController.getUpgradeTimelockConfig(app);
        IAppController.UpgradeTimelockConfig memory config2 = appController.getUpgradeTimelockConfig(app2);

        assertEq(config1.delay, 1 days);
        assertFalse(config1.locked);
        assertEq(config2.delay, 7 days);
        assertTrue(config2.locked);
    }

    // ========== Helper Functions ==========

    function _assembleRelease() internal view returns (IAppController.Release memory) {
        return _assembleReleaseWithDigest(keccak256("test-digest"));
    }

    function _assembleReleaseWithDigest(bytes32 digest) internal view returns (IAppController.Release memory) {
        IReleaseManagerTypes.Artifact[] memory artifacts = new IReleaseManagerTypes.Artifact[](1);
        artifacts[0] = IReleaseManagerTypes.Artifact({digest: digest, registry: "ipfs://test-registry"});

        // Use a far-future upgradeByTime to survive vm.warp in timelock tests
        IReleaseManagerTypes.Release memory rmsRelease =
            IReleaseManagerTypes.Release({artifacts: artifacts, upgradeByTime: uint32(block.timestamp + 30 days)});

        return IAppController.Release({rmsRelease: rmsRelease, publicEnv: "", encryptedEnv: ""});
    }
}
