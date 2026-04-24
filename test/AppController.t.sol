// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IAppController} from "../src/interfaces/IAppController.sol";
import {IAppAuthority} from "../src/interfaces/IAppAuthority.sol";
import {AppController} from "../src/AppController.sol";
import {ComputeDeployer} from "./utils/ComputeDeployer.sol";
import {IApp} from "../src/interfaces/IApp.sol";
import {PermissionControllerMixin} from "@eigenlayer-contracts/src/contracts/mixins/PermissionControllerMixin.sol";
import {IReleaseManagerTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {ICallValidator} from "../src/interfaces/ICallValidator.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

contract AppControllerTest is ComputeDeployer {
    bytes32 public constant SALT = keccak256("test_salt");

    // Event definitions for testing
    event AppCreated(address indexed creator, IApp indexed app, uint32 operatorSetId);
    event AppTerminatedByAdmin(IApp indexed app);
    event AppSuspended(IApp indexed app);
    event AppSuspendedByAdmin(IApp indexed app);
    event AppMetadataURIUpdated(IApp indexed app, string metadataURI);

    function setUp() public {
        _deployContracts();
        // Set maximum active apps for the developer account used in tests
        vm.prank(admin);
        appController.setMaxActiveAppsPerUser(developer, 100);
    }

    function _setGlobalMaxActiveApps(uint32 limit) internal {
        vm.prank(admin);
        appController.setMaxGlobalActiveApps(limit);
    }

    function _setMaxActiveAppsPerUser(address _user, uint32 _limit) internal {
        vm.prank(admin);
        appController.setMaxActiveAppsPerUser(_user, _limit);
    }

    function test_initialize_createsOperatorSetZeroAndRegistersComputeOperator() public {
        // After deployment and initialization, the nextOperatorSetId should be 1
        // because operator set 0 was created and assigned during initialize()
        assertEq(computeAVSRegistrar.nextOperatorSetId(), 1, "nextOperatorSetId should be 1 after initialization");

        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        uint32 operatorSetId = appController.getAppOperatorSetId(app);
        assertEq(operatorSetId, 1, "First app should get operator set ID 1, since 0 was used in initialize()");

        // Verify that nextOperatorSetId incremented to 2 after creating the app
        assertEq(computeAVSRegistrar.nextOperatorSetId(), 2, "nextOperatorSetId should be 2 after creating one app");
    }

    function test_createApp() public {
        vm.startPrank(developer);

        // Expect the AppCreated event
        vm.expectEmit(true, true, true, true);
        emit AppCreated(developer, appController.calculateAppId(developer, SALT), 1);

        IApp app = appController.createApp(SALT, _assembleRelease());

        assertTrue(address(app) != address(0));

        assertEq(appController.getAppOperatorSetId(app), 1);
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STARTED));
        assertEq(appController.getAppLatestReleaseBlockNumber(app), block.number);

        // Verify the app was created at the predicted address
        assertEq(address(app), address(appController.calculateAppId(developer, SALT)));

        // Default billing account should be the creator
        assertEq(appController.getBillingAccount(app), developer);

        vm.stopPrank();
    }

    function test_createApp_differentSaltsCreateDifferentApps() public {
        vm.startPrank(developer);

        bytes32 salt1 = keccak256("salt1");
        bytes32 salt2 = keccak256("salt2");

        IApp app1 = appController.createApp(salt1, _assembleRelease());
        IApp app2 = appController.createApp(salt2, _assembleRelease());

        // Verify different salts create different app addresses
        assertTrue(address(app1) != address(app2));

        // Verify both apps are properly configured
        IAppController.AppStatus status1 = appController.getAppStatus(app1);
        IAppController.AppStatus status2 = appController.getAppStatus(app2);

        assertEq(uint256(status1), uint256(IAppController.AppStatus.STARTED));
        assertEq(uint256(status2), uint256(IAppController.AppStatus.STARTED));

        vm.stopPrank();
    }

    function test_createApp_differentDeployersCreateDifferentApps() public {
        address developer2 = address(0x3);

        // Set up maximum active apps for developer2
        _setMaxActiveAppsPerUser(developer2, 100);

        // Developer 1 creates an app
        vm.prank(developer);
        IApp app1 = appController.createApp(SALT, _assembleRelease());

        // Developer 2 creates an app with the same salt
        vm.prank(developer2);
        IApp app2 = appController.createApp(SALT, _assembleRelease());

        // Verify same salt but different deployers create different app addresses
        assertTrue(address(app1) != address(app2));

        // Verify both apps are properly configured
        IAppController.AppStatus status1 = appController.getAppStatus(app1);
        IAppController.AppStatus status2 = appController.getAppStatus(app2);

        assertEq(uint256(status1), uint256(IAppController.AppStatus.STARTED));
        assertEq(uint256(status2), uint256(IAppController.AppStatus.STARTED));

        // They should have different operator set IDs
        uint32 operatorSetId1 = appController.getAppOperatorSetId(app1);
        uint32 operatorSetId2 = appController.getAppOperatorSetId(app2);
        assertTrue(operatorSetId1 != operatorSetId2);
    }

    function test_createApp_revertsOnDuplicateSalt() public {
        vm.startPrank(developer);

        // Create first app
        appController.createApp(SALT, _assembleRelease());

        // Try to create another app with the same salt - should revert with Create2 error
        vm.expectRevert("Create2: Failed on deploy");
        appController.createApp(SALT, _assembleRelease());

        vm.stopPrank();
    }

    function test_createApp_operatorSetIdIncrementsAndAppsAreUnique() public {
        vm.startPrank(developer);

        bytes32[] memory salts = new bytes32[](5);
        IApp[] memory apps = new IApp[](5);

        // Create multiple apps from the same developer
        for (uint256 i = 0; i < 5; i++) {
            salts[i] = keccak256(abi.encodePacked("salt", i));
            apps[i] = appController.createApp(salts[i], _assembleRelease());
        }

        // Verify all apps are different and properly configured
        for (uint256 i = 0; i < 5; i++) {
            uint32 operatorSetId = appController.getAppOperatorSetId(apps[i]);
            assertEq(operatorSetId, uint32(i + 1));
            IAppController.AppStatus status = appController.getAppStatus(apps[i]);
            assertEq(uint256(status), uint256(IAppController.AppStatus.STARTED));

            // Verify no duplicates
            for (uint256 j = i + 1; j < 5; j++) {
                assertTrue(address(apps[i]) != address(apps[j]));
            }
        }

        vm.stopPrank();
    }

    // Limit enforcement tests
    function test_createApp_noGlobalLimitSet() public {
        // Reset global limit to 0 to test the case where no global limit is set
        _setGlobalMaxActiveApps(0);

        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(IAppController.GlobalMaxActiveAppsExceeded.selector));
        appController.createApp(keccak256("test"), _assembleRelease());
    }

    function test_createApp_noUserLimitSet() public {
        _setGlobalMaxActiveApps(100);
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(IAppController.MaxActiveAppsExceeded.selector));
        appController.createApp(keccak256("test"), _assembleRelease());
    }

    function test_createApp_userLimitReached() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 1);
        vm.prank(user);
        appController.createApp(keccak256("test1"), _assembleRelease());
        assertEq(appController.getActiveAppCount(user), 1);
        vm.expectRevert(abi.encodeWithSelector(IAppController.MaxActiveAppsExceeded.selector));
        appController.createApp(keccak256("test2"), _assembleRelease());
    }

    function test_createApp_globalLimitReached() public {
        _setGlobalMaxActiveApps(1);
        _setMaxActiveAppsPerUser(user, 100);
        vm.prank(user);
        appController.createApp(keccak256("test1"), _assembleRelease());
        assertEq(appController.globalActiveAppCount(), 1);
        vm.expectRevert(abi.encodeWithSelector(IAppController.GlobalMaxActiveAppsExceeded.selector));
        appController.createApp(keccak256("test2"), _assembleRelease());
    }

    // Admin and permission tests
    function test_setMaxGlobalActiveApps() public {
        vm.prank(admin);
        appController.setMaxGlobalActiveApps(100);
        assertEq(appController.maxGlobalActiveApps(), 100);
    }

    function test_setMaxGlobalActiveApps_notAdmin() public {
        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.setMaxGlobalActiveApps(100);
    }

    function test_setMaxActiveAppsPerUser() public {
        vm.prank(admin);
        appController.setMaxActiveAppsPerUser(user, 1);
        assertEq(appController.getMaxActiveAppsPerUser(user), 1);
    }

    function test_setMaxActiveAppsPerUser_notAdmin() public {
        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.setMaxActiveAppsPerUser(user, 1);
    }

    // ========== App Lifecycle Tests ==========

    function test_upgradeApp() public {
        // Create an app first
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("upgrade_test"), _assembleRelease());
        _acceptAppAdmin(app);

        // Prepare a release
        IReleaseManagerTypes.Artifact[] memory artifacts = new IReleaseManagerTypes.Artifact[](1);
        artifacts[0] =
            IReleaseManagerTypes.Artifact({digest: keccak256("test-digest"), registry: "ipfs://test-registry"});

        IReleaseManagerTypes.Release memory rmsRelease =
            IReleaseManagerTypes.Release({artifacts: artifacts, upgradeByTime: uint32(block.timestamp + 1 days)});

        IAppController.Release memory release =
            IAppController.Release({rmsRelease: rmsRelease, publicEnv: "", encryptedEnv: ""});

        // Upgrade the app (should also start it since it's in CREATED status)
        vm.prank(developer);
        appController.upgradeApp(app, release);

        // Verify app status changed to STARTED and release was published
        IAppController.AppStatus status = appController.getAppStatus(app);
        assertEq(uint256(status), uint256(IAppController.AppStatus.STARTED));
        uint32 latestReleaseBlockNumber = appController.getAppLatestReleaseBlockNumber(app);
        assertEq(latestReleaseBlockNumber, block.number);
    }

    function test_upgradeApp_notAuthorized() public {
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("auth_test"), _assembleRelease());

        IReleaseManagerTypes.Artifact[] memory artifacts = new IReleaseManagerTypes.Artifact[](1);
        artifacts[0] =
            IReleaseManagerTypes.Artifact({digest: keccak256("test-digest"), registry: "ipfs://test-registry"});

        IReleaseManagerTypes.Release memory rmsRelease =
            IReleaseManagerTypes.Release({artifacts: artifacts, upgradeByTime: uint32(block.timestamp + 1 days)});

        IAppController.Release memory release =
            IAppController.Release({rmsRelease: rmsRelease, publicEnv: "", encryptedEnv: ""});

        // Try to upgrade as unauthorized user. Critical ops are owner-gated
        // unconditionally now — not even an ADMIN other than the owner can
        // trigger an upgrade.
        vm.prank(user);
        vm.expectRevert(IAppController.NotCreator.selector);
        appController.upgradeApp(app, release);
    }

    function test_stopApp() public {
        // Create an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("stop_test"), _assembleRelease());
        _acceptAppAdmin(app);

        // Stop the app
        vm.prank(developer);
        appController.stopApp(app);

        // Verify status changed to STOPPED
        IAppController.AppStatus status = appController.getAppStatus(app);
        assertEq(uint256(status), uint256(IAppController.AppStatus.STOPPED));
    }

    function test_stopApp_alreadyStopped() public {
        // Create an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("stop_test"), _assembleRelease());
        _acceptAppAdmin(app);

        // Stop the app
        vm.prank(developer);
        appController.stopApp(app);

        // Stop it again
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.stopApp(app);
    }

    function test_startApp() public {
        // Create an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("start_test"), _assembleRelease());
        _acceptAppAdmin(app);

        // Stop the app first
        vm.prank(developer);
        appController.stopApp(app);

        // Now start it again
        vm.prank(developer);
        appController.startApp(app);

        // Verify status changed back to STARTED
        IAppController.AppStatus status = appController.getAppStatus(app);
        assertEq(uint256(status), uint256(IAppController.AppStatus.STARTED));
    }

    function test_startApp_alreadyStarted() public {
        // Create an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("start_test"), _assembleRelease());
        _acceptAppAdmin(app);

        // Verify status changed to STARTED
        IAppController.AppStatus statusBefore = appController.getAppStatus(app);
        assertEq(uint256(statusBefore), uint256(IAppController.AppStatus.STARTED));

        // Start the app
        vm.prank(developer);
        appController.startApp(app);

        // Verify status still STARTED
        IAppController.AppStatus statusAfter = appController.getAppStatus(app);
        assertEq(uint256(statusAfter), uint256(IAppController.AppStatus.STARTED));
    }

    function test_startApp_cannotStartTerminated() public {
        // Create and terminate an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("start_terminated"), _assembleRelease());
        _acceptAppAdmin(app);

        vm.prank(developer);
        appController.terminateApp(app);

        // Try to start a terminated app - should fail (new check in _startApp)
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.startApp(app);
    }

    function test_startApp_nonExistentAppReverts() public {
        // Use any arbitrary address that was NOT created through createApp()
        address fakeAppAddress = address(0xDEADBEEF);
        IApp fakeApp = IApp(fakeAppAddress);

        // Verify this app doesn't exist
        assertEq(uint256(appController.getAppStatus(fakeApp)), uint256(IAppController.AppStatus.NONE));

        // Try to start a non-existent app — now fails on the onlyAdmin
        // check first (the caller holds no team role on an address whose
        // team was never created). This is strictly stronger than the
        // pre-RBAC InvalidAppStatus revert.
        vm.prank(fakeAppAddress);
        vm.expectRevert(IAppController.InvalidTeamRole.selector);
        appController.startApp(fakeApp);

        // Verify status is still NONE
        assertEq(uint256(appController.getAppStatus(fakeApp)), uint256(IAppController.AppStatus.NONE));
    }

    function test_terminateApp() public {
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("terminate_test"), _assembleRelease());
        _acceptAppAdmin(app);

        // Terminate the app (can be done from any active status)
        vm.prank(developer);
        appController.terminateApp(app);

        // Verify status changed to TERMINATED
        IAppController.AppStatus status = appController.getAppStatus(app);
        assertEq(uint256(status), uint256(IAppController.AppStatus.TERMINATED));
    }

    function test_terminateApp_alreadyTerminated() public {
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("double_terminate"), _assembleRelease());
        _acceptAppAdmin(app);

        vm.prank(developer);
        appController.terminateApp(app);

        // Try to terminate again - should revert
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.terminateApp(app);
    }

    // ========== View Function Tests ==========

    function test_calculateAppId() public view {
        bytes32 salt = keccak256("calc_test");
        IApp expectedApp = appController.calculateAppId(developer, salt);

        // The calculated address should be deterministic
        IApp expectedApp2 = appController.calculateAppId(developer, salt);
        assertEq(address(expectedApp), address(expectedApp2));

        // Different salts should give different addresses
        IApp differentApp = appController.calculateAppId(developer, keccak256("different_salt"));
        assertTrue(address(expectedApp) != address(differentApp));
    }

    function test_getAppProperties() public {
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("config_test"), _assembleRelease());

        IAppController.AppStatus status = appController.getAppStatus(app);
        assertEq(uint256(status), uint256(IAppController.AppStatus.STARTED));

        uint32 operatorSetId = appController.getAppOperatorSetId(app);
        assertEq(operatorSetId, 1); // Should be 1 since 0 is used for ComputeOperator

        uint32 latestReleaseBlockNumber = appController.getAppLatestReleaseBlockNumber(app);
        assertEq(latestReleaseBlockNumber, block.number);
    }

    function test_getApps_offsetAndLimit() public {
        // Create 5 apps
        IApp[] memory createdApps = new IApp[](5);
        for (uint256 i = 0; i < 5; i++) {
            vm.prank(developer);
            createdApps[i] = appController.createApp(keccak256(abi.encodePacked("offset_test_", i)), _assembleRelease());
            _acceptAppAdmin(createdApps[i]);
        }

        // Test 1: Normal pagination - offset 0, limit 3
        (IApp[] memory apps1,) = appController.getApps(0, 3);
        assertEq(apps1.length, 3);
        assertEq(address(apps1[0]), address(createdApps[0]));
        assertEq(address(apps1[1]), address(createdApps[1]));
        assertEq(address(apps1[2]), address(createdApps[2]));

        // Test 2: Offset in middle - offset 2, limit 2
        (IApp[] memory apps2,) = appController.getApps(2, 2);
        assertEq(apps2.length, 2);
        assertEq(address(apps2[0]), address(createdApps[2]));
        assertEq(address(apps2[1]), address(createdApps[3]));

        // Test 3: Limit exceeds remaining items - offset 3, limit 10
        (IApp[] memory apps3,) = appController.getApps(3, 10);
        assertEq(apps3.length, 2); // Only 2 apps left (indices 3,4)
        assertEq(address(apps3[0]), address(createdApps[3]));
        assertEq(address(apps3[1]), address(createdApps[4]));

        // Test 4: Offset equals total apps - should return empty
        (IApp[] memory apps4,) = appController.getApps(5, 10);
        assertEq(apps4.length, 0);

        // Test 5: Offset exceeds total apps - should return empty
        (IApp[] memory apps5,) = appController.getApps(10, 5);
        assertEq(apps5.length, 0);

        // Test 6: Limit 0 - should return empty
        (IApp[] memory apps6,) = appController.getApps(1, 0);
        assertEq(apps6.length, 0);

        // Test 7: Last single item - offset 4, limit 1
        (IApp[] memory apps7,) = appController.getApps(4, 1);
        assertEq(apps7.length, 1);
        assertEq(address(apps7[0]), address(createdApps[4]));
    }

    function test_getAppsByDeveloper_offsetAndLimit() public {
        // Create 5 apps as developer for testing true pagination
        IApp[] memory createdApps = new IApp[](5);
        for (uint256 i = 0; i < 5; i++) {
            vm.prank(developer);
            createdApps[i] =
                appController.createApp(keccak256(abi.encodePacked("dev_offset_test_", i)), _assembleRelease());
            _acceptAppAdmin(createdApps[i]);
        }

        // Test 1: Developer pagination - offset 0, limit 3 (first 3 of developer's apps)
        (IApp[] memory apps1,) = appController.getAppsByDeveloper(developer, 0, 3);
        assertEq(apps1.length, 3);
        assertEq(address(apps1[0]), address(createdApps[0]));
        assertEq(address(apps1[1]), address(createdApps[1]));
        assertEq(address(apps1[2]), address(createdApps[2]));

        // Test 2: Developer pagination - offset 2, limit 2 (3rd and 4th of developer's apps)
        (IApp[] memory apps2,) = appController.getAppsByDeveloper(developer, 2, 2);
        assertEq(apps2.length, 2);
        assertEq(address(apps2[0]), address(createdApps[2]));
        assertEq(address(apps2[1]), address(createdApps[3]));

        // Test 3: Developer pagination - offset 3, limit 10 (last 2 of developer's apps)
        (IApp[] memory apps3,) = appController.getAppsByDeveloper(developer, 3, 10);
        assertEq(apps3.length, 2); // Only 2 apps left for this developer
        assertEq(address(apps3[0]), address(createdApps[3]));
        assertEq(address(apps3[1]), address(createdApps[4]));

        // Test 4: Developer pagination - offset equals developer's total apps
        (IApp[] memory apps4,) = appController.getAppsByDeveloper(developer, 5, 10);
        assertEq(apps4.length, 0);

        // Test 5: Developer pagination - offset exceeds developer's total apps
        (IApp[] memory apps5,) = appController.getAppsByDeveloper(developer, 10, 5);
        assertEq(apps5.length, 0);

        // Test 6: Unknown developer (should return empty)
        address otherDev = makeAddr("otherDeveloper");
        (IApp[] memory apps6,) = appController.getAppsByDeveloper(otherDev, 0, 10);
        assertEq(apps6.length, 0);

        // Test 7: Mixed ownership scenario - create app for another developer
        _setMaxActiveAppsPerUser(otherDev, 10);
        vm.prank(otherDev);
        IApp otherApp1 = appController.createApp(keccak256("mixed_test_1"), _assembleRelease());
        vm.prank(otherDev);
        permissionController.acceptAdmin(address(otherApp1));

        vm.prank(otherDev);
        IApp otherApp2 = appController.createApp(keccak256("mixed_test_2"), _assembleRelease());
        vm.prank(otherDev);
        permissionController.acceptAdmin(address(otherApp2));

        // Test true pagination: first developer should still get their 5 apps in order
        (IApp[] memory devApps,) = appController.getAppsByDeveloper(developer, 0, 10);
        assertEq(devApps.length, 5);
        for (uint256 i = 0; i < 5; i++) {
            assertEq(address(devApps[i]), address(createdApps[i]));
        }

        // Test true pagination: second developer should get their 2 apps in order
        (IApp[] memory otherApps,) = appController.getAppsByDeveloper(otherDev, 0, 10);
        assertEq(otherApps.length, 2);
        assertEq(address(otherApps[0]), address(otherApp1));
        assertEq(address(otherApps[1]), address(otherApp2));

        // Test pagination continuity: otherDev first app, then second app
        (IApp[] memory otherApps1,) = appController.getAppsByDeveloper(otherDev, 0, 1);
        assertEq(otherApps1.length, 1);
        assertEq(address(otherApps1[0]), address(otherApp1));

        (IApp[] memory otherApps2,) = appController.getAppsByDeveloper(otherDev, 1, 1);
        assertEq(otherApps2.length, 1);
        assertEq(address(otherApps2[0]), address(otherApp2));
    }

    function test_getAppsByCreator_offsetAndLimit() public {
        // Create 5 apps as developer for testing pagination
        IApp[] memory createdApps = new IApp[](5);
        for (uint256 i = 0; i < 5; i++) {
            vm.prank(developer);
            createdApps[i] =
                appController.createApp(keccak256(abi.encodePacked("creator_test_", i)), _assembleRelease());
        }

        // Test 1: Creator pagination - offset 0, limit 3 (first 3 of creator's apps)
        (IApp[] memory apps1,) = appController.getAppsByCreator(developer, 0, 3);
        assertEq(apps1.length, 3);
        assertEq(address(apps1[0]), address(createdApps[0]));
        assertEq(address(apps1[1]), address(createdApps[1]));
        assertEq(address(apps1[2]), address(createdApps[2]));

        // Test 2: Creator pagination - offset 2, limit 2 (3rd and 4th of creator's apps)
        (IApp[] memory apps2,) = appController.getAppsByCreator(developer, 2, 2);
        assertEq(apps2.length, 2);
        assertEq(address(apps2[0]), address(createdApps[2]));
        assertEq(address(apps2[1]), address(createdApps[3]));

        // Test 3: Creator pagination - offset 3, limit 10 (last 2 of creator's apps)
        (IApp[] memory apps3,) = appController.getAppsByCreator(developer, 3, 10);
        assertEq(apps3.length, 2); // Only 2 apps left for this creator
        assertEq(address(apps3[0]), address(createdApps[3]));
        assertEq(address(apps3[1]), address(createdApps[4]));

        // Test 4: Creator pagination - offset equals creator's total apps
        (IApp[] memory apps4,) = appController.getAppsByCreator(developer, 5, 10);
        assertEq(apps4.length, 0);

        // Test 5: Creator pagination - offset exceeds creator's total apps
        (IApp[] memory apps5,) = appController.getAppsByCreator(developer, 10, 5);
        assertEq(apps5.length, 0);

        // Test 6: Unknown creator (should return empty)
        address otherDev = makeAddr("otherCreator");
        (IApp[] memory apps6,) = appController.getAppsByCreator(otherDev, 0, 10);
        assertEq(apps6.length, 0);

        // Test 7: Mixed ownership scenario - create app for another creator
        _setMaxActiveAppsPerUser(otherDev, 10);
        vm.prank(otherDev);
        IApp otherApp1 = appController.createApp(keccak256("creator_mixed_test_1"), _assembleRelease());
        vm.prank(otherDev);
        IApp otherApp2 = appController.createApp(keccak256("creator_mixed_test_2"), _assembleRelease());

        // Test pagination: first creator should still get their 5 apps in order
        (IApp[] memory devApps,) = appController.getAppsByCreator(developer, 0, 10);
        assertEq(devApps.length, 5);
        for (uint256 i = 0; i < 5; i++) {
            assertEq(address(devApps[i]), address(createdApps[i]));
        }

        // Test pagination: second creator should get their 2 apps in order
        (IApp[] memory otherApps,) = appController.getAppsByCreator(otherDev, 0, 10);
        assertEq(otherApps.length, 2);
        assertEq(address(otherApps[0]), address(otherApp1));
        assertEq(address(otherApps[1]), address(otherApp2));
    }

    function test_getAppsByCreator_worksWithoutAdminRights() public {
        vm.prank(developer);
        IApp app1 = appController.createApp(keccak256("admin_test_1"), _assembleRelease());
        vm.prank(developer);
        IApp app2 = appController.createApp(keccak256("admin_test_2"), _assembleRelease());

        // Only accept admin on app1, leave app2 without accepting admin
        vm.prank(developer);
        permissionController.acceptAdmin(address(app1));

        // getAppsByCreator should return BOTH apps (filters by creator, not admin)
        (IApp[] memory creatorApps,) = appController.getAppsByCreator(developer, 0, 10);
        assertEq(creatorApps.length, 2);
        assertEq(address(creatorApps[0]), address(app1));
        assertEq(address(creatorApps[1]), address(app2));

        // getAppsByDeveloper should only return app1 (developer only accepted admin on app1)
        (IApp[] memory devApps,) = appController.getAppsByDeveloper(developer, 0, 10);
        assertEq(devApps.length, 1);
        assertEq(address(devApps[0]), address(app1));

        // Verify the creator field is set correctly for both apps
        assertEq(appController.getAppCreator(app1), developer);
        assertEq(appController.getAppCreator(app2), developer);

        // Verify developer is only admin of app1
        assertTrue(permissionController.isAdmin(address(app1), developer));
        assertFalse(permissionController.isAdmin(address(app2), developer));
    }

    // ========== Helper Functions ==========

    function _acceptAppAdmin(IApp app) internal {
        // Accept admin permissions on the app (developer should be pending admin after app creation)
        vm.prank(developer);
        permissionController.acceptAdmin(address(app));
    }

    function test_activeAppCounting_terminationFreesUpCapacity() public {
        // Set user to have only 2 active apps max
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 2);

        // User creates 2 apps (reaches limit)
        vm.prank(user);
        IApp app1 = appController.createApp(keccak256("capacity_test_1"), _assembleRelease());
        vm.prank(user);
        appController.createApp(keccak256("capacity_test_2"), _assembleRelease());

        // Verify counts
        assertEq(appController.getActiveAppCount(user), 2);
        assertEq(appController.globalActiveAppCount(), 2);

        // Try to create a third app - should fail
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(IAppController.MaxActiveAppsExceeded.selector));
        appController.createApp(keccak256("capacity_test_3"), _assembleRelease());

        // Accept admin for app1 so we can terminate it
        vm.prank(user);
        permissionController.acceptAdmin(address(app1));

        // Terminate app1 - should free up capacity
        vm.prank(user);
        appController.terminateApp(app1);

        // Verify counts decreased
        assertEq(appController.getActiveAppCount(user), 1);
        assertEq(appController.globalActiveAppCount(), 1);

        // Now user should be able to create a new app
        vm.prank(user);
        IApp app3 = appController.createApp(keccak256("capacity_test_3"), _assembleRelease());

        // Verify counts increased again
        assertEq(appController.getActiveAppCount(user), 2);
        assertEq(appController.globalActiveAppCount(), 2);

        // Verify app creator is tracked correctly
        assertEq(appController.getAppCreator(app3), user);
    }

    function test_terminateAppByAdmin_notAuthorized() public {
        // Create an app as developer
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("admin_terminate_unauthorized"), _assembleRelease());

        // Try to terminate as unauthorized user (not admin)
        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.terminateAppByAdmin(app);
    }

    function test_terminateAppByAdmin_authorized() public {
        // Create an app as developer
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("admin_terminate_authorized"), _assembleRelease());

        // Expect the AppTerminatedByAdmin event
        vm.expectEmit(true, false, false, false);
        emit AppTerminatedByAdmin(app);

        // Terminate app as admin (should succeed)
        vm.prank(admin);
        appController.terminateAppByAdmin(app);

        // Verify status changed to TERMINATED
        IAppController.AppStatus status = appController.getAppStatus(app);
        assertEq(uint256(status), uint256(IAppController.AppStatus.TERMINATED));

        // Verify active app count decreased
        assertEq(appController.getActiveAppCount(developer), 0);
        assertEq(appController.globalActiveAppCount(), 0);
    }

    // ========== Metadata URI Tests ==========

    function test_updateAppMetadataURI() public {
        // Create an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("metadata_test"), _assembleRelease());
        _acceptAppAdmin(app);

        string memory newMetadataURI = "https://example.com/metadata";

        // Expect the AppMetadataURIUpdated event
        vm.expectEmit(true, false, false, true);
        emit AppMetadataURIUpdated(app, newMetadataURI);

        // Update metadata URI as app admin
        vm.prank(developer);
        appController.updateAppMetadataURI(app, newMetadataURI);
    }

    function test_updateAppMetadataURI_notAuthorized() public {
        // Create an app as developer
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("metadata_unauthorized"), _assembleRelease());

        // Try to update metadata as unauthorized user
        vm.prank(user);
        vm.expectRevert(IAppController.InvalidTeamRole.selector);
        appController.updateAppMetadataURI(app, "https://example.com/metadata");
    }

    function test_updateAppMetadataURI_terminatedApp() public {
        // Create an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("metadata_terminated"), _assembleRelease());
        _acceptAppAdmin(app);

        // Terminate the app
        vm.prank(developer);
        appController.terminateApp(app);

        string memory newMetadataURI = "https://example.com/metadata";

        // Expect the AppMetadataURIUpdated event (should succeed even for terminated app)
        vm.expectEmit(true, false, false, true);
        emit AppMetadataURIUpdated(app, newMetadataURI);

        // Update metadata on terminated app - should succeed now
        vm.prank(developer);
        appController.updateAppMetadataURI(app, newMetadataURI);
    }

    function test_updateAppMetadataURI_nonExistentAppReverts() public {
        address fakeAppAddress = address(0xDEADBEEF);
        IApp fakeApp = IApp(fakeAppAddress);

        // Verify this app doesn't exist
        assertEq(uint256(appController.getAppStatus(fakeApp)), uint256(IAppController.AppStatus.NONE));

        // Try to update metadata for a non-existent app — caller holds no
        // DEVELOPER/ADMIN role on the fake team, so the outer role gate
        // rejects first. Strictly stronger than the pre-RBAC behavior.
        vm.prank(fakeAppAddress);
        vm.expectRevert(IAppController.InvalidTeamRole.selector);
        appController.updateAppMetadataURI(fakeApp, "https://example.com/fake-metadata");
    }

    function _assembleRelease() internal view returns (IAppController.Release memory) {
        IReleaseManagerTypes.Artifact[] memory artifacts = new IReleaseManagerTypes.Artifact[](1);
        artifacts[0] =
            IReleaseManagerTypes.Artifact({digest: keccak256("test-digest"), registry: "ipfs://test-registry"});

        IReleaseManagerTypes.Release memory rmsRelease =
            IReleaseManagerTypes.Release({artifacts: artifacts, upgradeByTime: uint32(block.timestamp + 1 days)});

        IAppController.Release memory release =
            IAppController.Release({rmsRelease: rmsRelease, publicEnv: "", encryptedEnv: ""});

        return release;
    }

    // ========== Suspension Tests ==========

    function test_suspend() public {
        // Setup
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);

        // User creates 3 apps
        vm.prank(user);
        IApp app1 = appController.createApp(keccak256("suspend_batch_1"), _assembleRelease());
        vm.prank(user);
        IApp app2 = appController.createApp(keccak256("suspend_batch_2"), _assembleRelease());
        vm.prank(user);
        IApp app3 = appController.createApp(keccak256("suspend_batch_3"), _assembleRelease());

        // Verify initial state
        assertEq(appController.getActiveAppCount(user), 3);
        assertEq(appController.getMaxActiveAppsPerUser(user), 10);

        // Admin suspends the account
        IApp[] memory apps = new IApp[](3);
        apps[0] = app1;
        apps[1] = app2;
        apps[2] = app3;

        vm.prank(admin);
        appController.suspend(user, apps);

        // Verify all apps are suspended
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app3)), uint256(IAppController.AppStatus.SUSPENDED));

        // Verify capacity is freed and max is zeroed
        assertEq(appController.getActiveAppCount(user), 0);
        assertEq(appController.getMaxActiveAppsPerUser(user), 0);
        assertEq(appController.globalActiveAppCount(), 0);
    }

    function test_suspend_notAuthorized() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);

        vm.prank(user);
        IApp app = appController.createApp(keccak256("suspend_auth"), _assembleRelease());

        IApp[] memory apps = new IApp[](1);
        apps[0] = app;

        // Try to suspend as non-admin, non-owner
        vm.prank(developer);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.suspend(user, apps);
    }

    function test_suspend_byAccountOwner() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);

        // User creates 2 apps
        vm.prank(user);
        IApp app1 = appController.createApp(keccak256("owner_suspend_1"), _assembleRelease());
        vm.prank(user);
        IApp app2 = appController.createApp(keccak256("owner_suspend_2"), _assembleRelease());

        // Verify initial state
        assertEq(appController.getActiveAppCount(user), 2);
        assertEq(appController.getMaxActiveAppsPerUser(user), 10);

        // User suspends their own account (no admin permission needed)
        IApp[] memory apps = new IApp[](2);
        apps[0] = app1;
        apps[1] = app2;

        vm.prank(user);
        appController.suspend(user, apps);

        // Verify suspension succeeded
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(appController.getActiveAppCount(user), 0);
        assertEq(appController.getMaxActiveAppsPerUser(user), 0);
    }

    function test_suspend_invalidOwnership() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);
        _setMaxActiveAppsPerUser(developer, 10);

        // User creates an app
        vm.prank(user);
        IApp app = appController.createApp(keccak256("suspend_ownership"), _assembleRelease());

        // Try to suspend it as developer's app (wrong owner)
        IApp[] memory apps = new IApp[](1);
        apps[0] = app;

        vm.prank(admin);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.suspend(developer, apps);
    }

    function test_suspend_accountHasActiveApps() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);

        // User creates 3 apps
        vm.prank(user);
        IApp app1 = appController.createApp(keccak256("active_1"), _assembleRelease());
        vm.prank(user);
        IApp app2 = appController.createApp(keccak256("active_2"), _assembleRelease());
        vm.prank(user);
        appController.createApp(keccak256("active_3"), _assembleRelease());

        // Try to suspend with only 2 apps provided (app3 still active)
        IApp[] memory apps = new IApp[](2);
        apps[0] = app1;
        apps[1] = app2;

        vm.prank(admin);
        vm.expectRevert(abi.encodeWithSelector(IAppController.AccountHasActiveApps.selector));
        appController.suspend(user, apps);
    }

    function test_suspend_mixedStates() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);

        // Create 4 apps (all start as STARTED)
        vm.prank(user);
        IApp app1 = appController.createApp(keccak256("mixed_1"), _assembleRelease());
        vm.prank(user);
        IApp app2 = appController.createApp(keccak256("mixed_2"), _assembleRelease());
        vm.prank(user);
        IApp app3 = appController.createApp(keccak256("mixed_3"), _assembleRelease());
        vm.prank(user);
        IApp app4 = appController.createApp(keccak256("mixed_4"), _assembleRelease());

        // Suspend all apps first
        IApp[] memory allApps = new IApp[](4);
        allApps[0] = app1;
        allApps[1] = app2;
        allApps[2] = app3;
        allApps[3] = app4;

        vm.prank(admin);
        appController.suspend(user, allApps);

        // Verify all suspended and counters zeroed
        assertEq(appController.getActiveAppCount(user), 0);
        assertEq(appController.getMaxActiveAppsPerUser(user), 0);

        // Restore capacity and restart some apps to create mixed states
        vm.prank(admin);
        appController.setMaxActiveAppsPerUser(user, 10);

        // Accept admin for apps we want to manipulate
        vm.prank(user);
        permissionController.acceptAdmin(address(app1));
        vm.prank(user);
        permissionController.acceptAdmin(address(app2));
        vm.prank(user);
        permissionController.acceptAdmin(address(app4));

        // Set up different states:
        // app1 -> STARTED
        vm.prank(user);
        appController.startApp(app1);

        // app2 -> STOPPED
        vm.prank(user);
        appController.startApp(app2);
        vm.prank(user);
        appController.stopApp(app2);

        // app3 -> SUSPENDED (leave it suspended)

        // app4 -> TERMINATED
        vm.prank(user);
        appController.startApp(app4);
        vm.prank(user);
        appController.terminateApp(app4);

        // Verify state before final suspend (app1 STARTED, app2 STOPPED, app3 SUSPENDED, app4 TERMINATED)
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.STARTED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.STOPPED));
        assertEq(uint256(appController.getAppStatus(app3)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app4)), uint256(IAppController.AppStatus.TERMINATED));
        assertEq(appController.getActiveAppCount(user), 2); // Only app1 and app2 are active

        // Suspend all apps (tests handling of mixed states)
        vm.prank(admin);
        appController.suspend(user, allApps);

        // Verify: app1 and app2 are suspended, app3 stays suspended, app4 stays terminated
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app3)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app4)), uint256(IAppController.AppStatus.TERMINATED));

        // Verify capacity is freed and max is zeroed
        assertEq(appController.getActiveAppCount(user), 0);
        assertEq(appController.getMaxActiveAppsPerUser(user), 0);
    }

    function test_suspend_emptyArray() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);

        // Suspend with empty array (should just zero out maxActiveApps)
        IApp[] memory apps = new IApp[](0);

        vm.prank(admin);
        appController.suspend(user, apps);

        // Verify max is zeroed and active count is 0
        assertEq(appController.getMaxActiveAppsPerUser(user), 0);
        assertEq(appController.getActiveAppCount(user), 0);
    }

    function test_suspend_thenResume() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);

        // User creates 2 apps
        vm.prank(user);
        IApp app1 = appController.createApp(keccak256("resume_1"), _assembleRelease());
        vm.prank(user);
        IApp app2 = appController.createApp(keccak256("resume_2"), _assembleRelease());

        // Verify initial state
        assertEq(appController.getActiveAppCount(user), 2);
        assertEq(appController.getMaxActiveAppsPerUser(user), 10);
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.STARTED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.STARTED));

        // Admin suspends the account
        IApp[] memory apps = new IApp[](2);
        apps[0] = app1;
        apps[1] = app2;

        vm.prank(admin);
        appController.suspend(user, apps);

        // Verify apps are suspended and counters are zeroed
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(appController.getActiveAppCount(user), 0);
        assertEq(appController.getMaxActiveAppsPerUser(user), 0);

        // Admin lifts ban by increasing maxActiveApps back to 2
        vm.prank(admin);
        appController.setMaxActiveAppsPerUser(user, 2);

        // Accept admin to start apps
        vm.prank(user);
        permissionController.acceptAdmin(address(app1));
        vm.prank(user);
        permissionController.acceptAdmin(address(app2));

        // User resumes apps
        vm.prank(user);
        appController.startApp(app1);
        vm.prank(user);
        appController.startApp(app2);

        // Verify apps are started and capacity is restored
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.STARTED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.STARTED));
        assertEq(appController.getActiveAppCount(user), 2);
        assertEq(appController.getMaxActiveAppsPerUser(user), 2);
    }

    // ========== Isolated Billing App Tests ==========

    function test_createAppWithIsolatedBilling() public {
        // Pre-compute the app address
        IApp expectedApp = appController.calculateAppId(developer, SALT);

        // Set quota for the app address (not the developer)
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());

        // Verify app was created at the expected address
        assertEq(address(app), address(expectedApp));

        // Verify billing is isolated
        assertEq(uint256(appController.getBillingType(app)), uint256(IAppController.BillingType.ISOLATED));

        // Verify billing account is the app address itself
        assertEq(appController.getBillingAccount(app), address(app));

        // Verify app is started
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STARTED));

        // Verify the app address's quota was used (not the developer's)
        assertEq(appController.getActiveAppCount(address(app)), 1);
        assertEq(appController.getActiveAppCount(developer), 0);
    }

    function test_createAppWithIsolatedBilling_noQuota_reverts() public {
        // Don't set any quota for the app address — should revert
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.MaxActiveAppsExceeded.selector));
        appController.createAppWithIsolatedBilling(SALT, _assembleRelease());
    }

    function test_createAppWithIsolatedBilling_terminateDecrementsAppQuota() public {
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());
        _acceptAppAdmin(app);

        // Verify app address has 1 active app
        assertEq(appController.getActiveAppCount(address(app)), 1);

        // Terminate the app
        vm.prank(developer);
        appController.terminateApp(app);

        // Verify the app address's count decremented (not the developer's)
        assertEq(appController.getActiveAppCount(address(app)), 0);
        assertEq(appController.getActiveAppCount(developer), 0);
    }

    function test_createAppWithIsolatedBilling_suspendByAppAddress() public {
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());

        // Admin suspends the app using the app address as the account
        IApp[] memory apps = new IApp[](1);
        apps[0] = app;

        vm.prank(admin);
        appController.suspend(address(app), apps);

        // Verify the app is suspended
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(appController.getActiveAppCount(address(app)), 0);
        assertEq(appController.getMaxActiveAppsPerUser(address(app)), 0);
    }

    function test_createAppWithIsolatedBilling_resumeFromSuspended() public {
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());
        _acceptAppAdmin(app);

        // Suspend via admin
        IApp[] memory apps = new IApp[](1);
        apps[0] = app;
        vm.prank(admin);
        appController.suspend(address(app), apps);

        // Verify suspended
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(appController.getActiveAppCount(address(app)), 0);

        // Restore quota for the app address
        _setMaxActiveAppsPerUser(address(app), 1);

        // Developer resumes the app
        vm.prank(developer);
        appController.startApp(app);

        // Verify resumed and app quota re-checked
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STARTED));
        assertEq(appController.getActiveAppCount(address(app)), 1);
    }

    function test_createAppWithIsolatedBilling_developerStillControls() public {
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());
        _acceptAppAdmin(app);

        // Developer can stop the app
        vm.prank(developer);
        appController.stopApp(app);
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STOPPED));

        // Developer can start the app
        vm.prank(developer);
        appController.startApp(app);
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STARTED));

        // Developer can upgrade the app
        vm.prank(developer);
        appController.upgradeApp(app, _assembleRelease());

        // Developer can terminate the app
        vm.prank(developer);
        appController.terminateApp(app);
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.TERMINATED));
    }

    function test_createApp_defaultBilling() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        // Existing createApp should have default billing
        assertEq(uint256(appController.getBillingType(app)), uint256(IAppController.BillingType.DEFAULT));
        assertEq(appController.getBillingAccount(app), developer);

        // Billing should go to the developer
        assertEq(appController.getActiveAppCount(developer), 1);
    }

    // ========== getAppsByBillingAccount Tests ==========

    function test_getAppsByBillingAccount_regularApps() public {
        // Create regular apps as developer
        vm.prank(developer);
        IApp app1 = appController.createApp(keccak256("billing_1"), _assembleRelease());
        vm.prank(developer);
        IApp app2 = appController.createApp(keccak256("billing_2"), _assembleRelease());

        // Query by developer address — should return both
        (IApp[] memory apps,) = appController.getAppsByBillingAccount(developer, 0, 10);
        assertEq(apps.length, 2);
        assertEq(address(apps[0]), address(app1));
        assertEq(address(apps[1]), address(app2));
    }

    function test_getAppsByBillingAccount_isolatedBillingApp() public {
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());

        // Query by app address — should return the isolated billing app
        (IApp[] memory appsByApp,) = appController.getAppsByBillingAccount(address(app), 0, 10);
        assertEq(appsByApp.length, 1);
        assertEq(address(appsByApp[0]), address(app));

        // Query by developer — should NOT return the isolated billing app
        (IApp[] memory appsByDev,) = appController.getAppsByBillingAccount(developer, 0, 10);
        assertEq(appsByDev.length, 0);
    }

    function test_getAppsByBillingAccount_mixed() public {
        // Create a regular app billed to developer
        vm.prank(developer);
        IApp regularApp = appController.createApp(keccak256("mixed_regular"), _assembleRelease());

        // Create an isolated billing app billed to itself
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);
        vm.prank(developer);
        IApp isolatedApp = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());

        // Verify billing accounts
        assertEq(appController.getBillingAccount(regularApp), developer);
        assertEq(appController.getBillingAccount(isolatedApp), address(isolatedApp));

        // Developer billing account returns only the regular app
        (IApp[] memory devApps,) = appController.getAppsByBillingAccount(developer, 0, 10);
        assertEq(devApps.length, 1);
        assertEq(address(devApps[0]), address(regularApp));

        // App address billing account returns only the isolated billing app
        (IApp[] memory isoApps,) = appController.getAppsByBillingAccount(address(isolatedApp), 0, 10);
        assertEq(isoApps.length, 1);
        assertEq(address(isoApps[0]), address(isolatedApp));
    }

    // ========== Isolated Billing Edge Cases ==========

    function test_suspendIsolatedBillingApp_withCreatorAccount_reverts() public {
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());

        // Trying to suspend using the creator (developer) as account should revert
        // because the isolated billing app's ownership check requires the app address, not the creator
        IApp[] memory apps = new IApp[](1);
        apps[0] = app;

        vm.prank(admin);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.suspend(developer, apps);
    }

    function test_getAppsByCreator_returnsIsolatedBillingApps() public {
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());

        // getAppsByCreator should still return the isolated billing app (creator is still set)
        (IApp[] memory apps,) = appController.getAppsByCreator(developer, 0, 10);
        assertEq(apps.length, 1);
        assertEq(address(apps[0]), address(app));

        // But getAppsByBillingAccount with developer should NOT return it
        (IApp[] memory billingApps,) = appController.getAppsByBillingAccount(developer, 0, 10);
        assertEq(billingApps.length, 0);
    }

    function test_terminateAppByAdmin_isolatedBilling_decrementsAppQuota() public {
        IApp expectedApp = appController.calculateAppId(developer, SALT);
        _setMaxActiveAppsPerUser(address(expectedApp), 1);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());

        // Verify app address has 1 active app
        assertEq(appController.getActiveAppCount(address(app)), 1);
        assertEq(appController.getActiveAppCount(developer), 0);

        // Admin terminates the app
        vm.prank(admin);
        appController.terminateAppByAdmin(app);

        // Verify the app address's count decremented (not the developer's)
        assertEq(appController.getActiveAppCount(address(app)), 0);
        assertEq(appController.getActiveAppCount(developer), 0);
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.TERMINATED));
    }

    // ========== Timelocked-gate regression tests ==========
    //
    // These tests pin down the runtime invariant that sensitive ops against a
    // Timelock-owned app must go through schedule → execute. The gate lives
    // in upgradeApp / terminateApp / terminateAppByAdmin and fires whenever
    // `_appConfigs[app].timelocked == true`, which is set at creation if
    // msg.sender is a factory-registered Timelock.
    //
    // We mock `isTimelock(developer)` to true so createApp flips the flag,
    // then assert that a PermissionController-permitted co-admin is blocked
    // from calling the sensitive op directly.

    function _mockIsTimelock(address account) internal {
        address factory = address(AppController(address(appController)).safeTimelockFactory());
        vm.mockCall(factory, abi.encodeWithSignature("isTimelock(address)", account), abi.encode(true));
    }

    function test_createApp_flagsTimelockedWhenCallerIsTimelock() public {
        _mockIsTimelock(developer);

        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        assertTrue(appController.getAppTimelocked(app), "Timelock-created app must have timelocked=true at creation");
    }

    function test_createApp_doesNotFlagTimelockedForNonTimelockCaller() public {
        // No mock: factory returns false for random addresses. Regression
        // guard that the fix doesn't over-apply the flag.
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        assertFalse(appController.getAppTimelocked(app), "EOA-created app must not be flagged timelocked");
    }

    function test_upgradeApp_timelockedBlocksNonOwnerEvenWithPermission() public {
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());
        require(appController.getAppTimelocked(app), "precondition: app must be timelocked");

        // Timelock (creator) grants ADMIN role to a co-admin. Under the RBAC
        // model, only ADMIN can even attempt upgradeApp; holding ADMIN is
        // the strongest authority the co-admin could plausibly have.
        address coAdmin = makeAddr("coAdmin");
        // Timelock ADMIN grant on a timelocked app requires msg.sender == creator;
        // developer IS the creator (via _mockIsTimelock), so this passes directly.
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, coAdmin);

        // Direct upgrade from the co-admin MUST revert with NotCreator —
        // the owner-gate blocks everyone except the current owner, regardless
        // of ADMIN membership.
        vm.prank(coAdmin);
        vm.expectRevert(IAppController.NotCreator.selector);
        appController.upgradeApp(app, _assembleRelease());

        // The Timelock itself (app.creator) can still call upgrade directly —
        // representing the path where a scheduled op is being executed.
        vm.prank(developer);
        appController.upgradeApp(app, _assembleRelease());
    }

    function test_terminateApp_timelockedBlocksNonOwner() public {
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address coAdmin = makeAddr("coAdmin");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, coAdmin);

        vm.prank(coAdmin);
        vm.expectRevert(IAppController.NotCreator.selector);
        appController.terminateApp(app);

        // Timelock (creator) can terminate.
        vm.prank(developer);
        appController.terminateApp(app);
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.TERMINATED));
    }

    function test_terminateAppByAdmin_refusesTimelockedApp() public {
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        // Even the protocol admin cannot terminate a timelocked app directly —
        // they must go through the Timelock to preserve the delay invariant.
        vm.prank(admin);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.terminateAppByAdmin(app);
    }

    // ========== transferOwnership ==========

    event AppOwnershipTransferred(IApp indexed app, address indexed previousOwner, address indexed newOwner);

    function test_transferOwnership_toEOA_clearsTimelocked() public {
        // Start with a Timelock-owned app.
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());
        vm.prank(developer);
        permissionController.acceptAdmin(address(app));
        assertTrue(appController.getAppTimelocked(app));

        address plainEOA = makeAddr("plainEOA");

        vm.expectEmit(true, true, true, true);
        emit AppOwnershipTransferred(app, developer, plainEOA);

        vm.prank(developer);
        appController.transferOwnership(app, plainEOA);

        assertFalse(appController.getAppTimelocked(app), "timelocked must clear when new owner is not a Timelock");
        assertEq(appController.getAppCreator(app), plainEOA);
    }

    function test_transferOwnership_toTimelock_setsTimelocked() public {
        // Non-timelocked starting state.
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());
        vm.prank(developer);
        permissionController.acceptAdmin(address(app));
        assertFalse(appController.getAppTimelocked(app));

        address newTimelock = makeAddr("newTimelock");
        _mockIsTimelock(newTimelock);

        vm.prank(developer);
        appController.transferOwnership(app, newTimelock);

        assertTrue(appController.getAppTimelocked(app), "timelocked must set when new owner is a Timelock");
        assertEq(appController.getAppCreator(app), newTimelock);
    }

    function test_transferOwnership_timelockedBlocksNonOwner() public {
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        // Timelock grants ADMIN to coAdmin. ADMIN is the strongest role a
        // co-admin could plausibly have; the creator-only gate must still
        // block them from moving the app out of governance.
        address coAdmin = makeAddr("coAdmin");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, coAdmin);

        address attacker = makeAddr("attacker");
        vm.prank(coAdmin);
        vm.expectRevert(IAppController.NotCreator.selector);
        appController.transferOwnership(app, attacker);

        // Timelock itself can still transfer.
        vm.prank(developer);
        appController.transferOwnership(app, attacker);
        assertEq(appController.getAppCreator(app), attacker);
    }

    function test_transferOwnership_revertsZeroAddress() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());
        // developer is auto-granted ADMIN at createApp time; no UAM step needed.

        vm.prank(developer);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.transferOwnership(app, address(0));
    }

    function test_transferOwnership_defaultBilling_movesActiveCounter() public {
        // Default billing: creator is the billing account. Counter must move
        // so future terminate/suspend on the new owner doesn't underflow.
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());
        vm.prank(developer);
        permissionController.acceptAdmin(address(app));

        address newOwner = makeAddr("newOwner");
        _setMaxActiveAppsPerUser(newOwner, 10);

        uint32 beforeDev = appController.getActiveAppCount(developer);
        uint32 beforeNew = appController.getActiveAppCount(newOwner);

        vm.prank(developer);
        appController.transferOwnership(app, newOwner);

        assertEq(appController.getActiveAppCount(developer), beforeDev - 1);
        assertEq(appController.getActiveAppCount(newOwner), beforeNew + 1);
    }

    function test_transferOwnership_isolatedBilling_doesNotMoveCounter() public {
        // ISOLATED billing apps bill the app address, so the creator's
        // active-app counter was never incremented and must not move on transfer.
        _setMaxActiveAppsPerUser(address(appController.calculateAppId(developer, SALT)), 10);

        vm.prank(developer);
        IApp app = appController.createAppWithIsolatedBilling(SALT, _assembleRelease());
        vm.prank(developer);
        permissionController.acceptAdmin(address(app));

        uint32 beforeApp = appController.getActiveAppCount(address(app));
        uint32 beforeDev = appController.getActiveAppCount(developer);

        address newOwner = makeAddr("newOwner");
        vm.prank(developer);
        appController.transferOwnership(app, newOwner);

        assertEq(appController.getActiveAppCount(address(app)), beforeApp);
        assertEq(appController.getActiveAppCount(developer), beforeDev);
        assertEq(appController.getActiveAppCount(newOwner), 0);
    }

    // ========== ICallValidator / canCall ==========

    function test_supportsInterface_advertisesCallValidator() public {
        AppController ac = AppController(address(appController));
        bytes4 callValidatorId = type(ICallValidator).interfaceId;
        bytes4 erc165Id = type(IERC165).interfaceId;

        assertTrue(ac.supportsInterface(callValidatorId), "AppController must advertise ICallValidator");
        assertTrue(ac.supportsInterface(erc165Id), "AppController must advertise IERC165");
        assertFalse(ac.supportsInterface(0xdeadbeef), "Unrelated interface must return false");
    }

    function test_canCall_returnsTrueForShortCalldata() public view {
        // Fallback: anything under 4 bytes can't be a meaningful call; defer.
        assertTrue(ICallValidator(address(appController)).canCall(address(this), ""));
        assertTrue(ICallValidator(address(appController)).canCall(address(this), hex"00"));
    }

    function test_canCall_rejectsNonTimelockCallerOnTimelockedUpgradeApp() public {
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        bytes memory callData = abi.encodeWithSelector(IAppController.upgradeApp.selector, app, _assembleRelease());
        address notOwner = makeAddr("notOwner");

        assertFalse(
            ICallValidator(address(appController)).canCall(notOwner, callData),
            "canCall must reject non-owner schedule of timelocked upgradeApp"
        );
        assertTrue(
            ICallValidator(address(appController)).canCall(developer, callData),
            "canCall must accept owner (Timelock itself) schedule of timelocked upgradeApp"
        );
    }

    function test_canCall_rejectsNonOwnerUpgradeAppEvenWhenNotTimelocked() public {
        // Owner-gated model: canCall rejects any non-owner schedule of a
        // critical op, regardless of timelocked state. Previously this was
        // only checked for timelocked apps; now owner-gating is unconditional
        // so schedule-time rejection matches runtime.
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());
        bytes memory callData = abi.encodeWithSelector(IAppController.upgradeApp.selector, app, _assembleRelease());

        assertFalse(
            ICallValidator(address(appController)).canCall(makeAddr("anyone"), callData),
            "canCall must reject non-owner schedule even for non-timelocked apps"
        );
        assertTrue(
            ICallValidator(address(appController)).canCall(developer, callData), "canCall must accept owner schedule"
        );
    }

    function test_canCall_rejectsTerminateAppByAdminOnTimelockedApp() public {
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        bytes memory callData = abi.encodeWithSelector(IAppController.terminateAppByAdmin.selector, app);
        // terminateAppByAdmin against a timelocked app is unconditionally doomed
        // regardless of caller — canCall reflects that.
        assertFalse(ICallValidator(address(appController)).canCall(admin, callData));
    }

    // ========== Team-role RBAC ==========

    function test_createApp_grantsAdminRoleToCreator() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        assertTrue(
            appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, developer), "creator must be seeded as ADMIN on create"
        );
    }

    function test_grantTeamRole_asAdmin() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address pauser = makeAddr("pauser");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.PAUSER, pauser);

        assertTrue(appAuthority.hasRole(app, IAppAuthority.Role.PAUSER, pauser));
    }

    function test_grantTeamRole_nonAdminCannotGrant() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address outsider = makeAddr("outsider");
        vm.prank(outsider);
        vm.expectRevert(IAppAuthority.InvalidRole.selector);
        appAuthority.grantRole(app, IAppAuthority.Role.PAUSER, outsider);
    }

    function test_grantTeamRole_timelockedOperationalRoleNotGated() public {
        // A-2 fix: PAUSER/DEVELOPER grants on timelocked apps are NOT
        // routed through the Timelock. Any existing ADMIN can grant — the
        // power delegated is operational only.
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        // Give coAdmin ADMIN via the Timelock (creator), then confirm
        // coAdmin can grant PAUSER freely (no creator-only gate).
        address coAdmin = makeAddr("coAdmin");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, coAdmin);

        address pauser = makeAddr("pauser");
        vm.prank(coAdmin);
        appAuthority.grantRole(app, IAppAuthority.Role.PAUSER, pauser);
        assertTrue(appAuthority.hasRole(app, IAppAuthority.Role.PAUSER, pauser));
    }

    function test_grantTeamRole_adminGrantGatedByCreator() public {
        // Owner-gated model: ADMIN grants are creator-only regardless of
        // timelocked/Safe/EOA ownership. Applies to timelocked apps (Timelock
        // is creator) and to non-timelocked apps (EOA/Safe is creator). A
        // co-ADMIN cannot mint another ADMIN.
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address coAdmin = makeAddr("coAdmin");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, coAdmin);

        // coAdmin has ADMIN but is NOT the owner: grant attempt must fail.
        address usurper = makeAddr("usurper");
        vm.prank(coAdmin);
        vm.expectRevert(IAppAuthority.NotScopeOwner.selector);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, usurper);

        // The Timelock (owner) can still grant.
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, usurper);
        assertTrue(appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, usurper));
    }

    function test_grantTeamRole_adminGrantGatedByCreator_nonTimelocked() public {
        // Safe/EOA-owned app: same owner-only rule applies. This is what
        // prevents a co-ADMIN from silently escalating to owner on Safe-owned
        // apps and bypassing the multisig.
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address coAdmin = makeAddr("coAdmin");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, coAdmin);

        address usurper = makeAddr("usurper");
        vm.prank(coAdmin);
        vm.expectRevert(IAppAuthority.NotScopeOwner.selector);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, usurper);
    }

    function test_revokeTeamRole_adminRevokeGatedByCreator() public {
        // Mirror of grant: revoking ADMIN is owner-only unconditionally.
        _mockIsTimelock(developer);
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address coAdmin = makeAddr("coAdmin");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.ADMIN, coAdmin);

        // coAdmin attempts to revoke the Timelock — MUST fail.
        vm.prank(coAdmin);
        vm.expectRevert(IAppAuthority.NotScopeOwner.selector);
        appAuthority.revokeRole(app, IAppAuthority.Role.ADMIN, developer);

        // The Timelock (creator) can revoke coAdmin.
        vm.prank(developer);
        appAuthority.revokeRole(app, IAppAuthority.Role.ADMIN, coAdmin);
        assertFalse(appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, coAdmin));
    }

    function test_revokeTeamRole_creatorCannotRevokeSelf() public {
        // The creator is the only address that can call revokeTeamRole(ADMIN),
        // but they cannot revoke their own ADMIN — the invariant
        // `creator ∈ ADMIN` must hold. transferOwnership is the only path that
        // rotates.
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        vm.prank(developer);
        vm.expectRevert(IAppAuthority.CannotRemoveOwnerAdmin.selector);
        appAuthority.revokeRole(app, IAppAuthority.Role.ADMIN, developer);
    }

    function test_renounceTeamRole_creatorCannotRenounceAdmin() public {
        // The creator cannot drop out of ADMIN via renounce — doing so would
        // leave the owner slot pointing at an address that isn't ADMIN
        // (though ADMIN is operational-only in this model, keeping the
        // invariant still matters for startApp and role-management clarity).
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        vm.prank(developer);
        vm.expectRevert(IAppAuthority.CannotRemoveOwnerAdmin.selector);
        appAuthority.renounceRole(app, IAppAuthority.Role.ADMIN);
    }

    function test_renounceTeamRole_operationalRoleOk() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address pauser = makeAddr("pauser");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.PAUSER, pauser);

        vm.prank(pauser);
        appAuthority.renounceRole(app, IAppAuthority.Role.PAUSER);
        assertFalse(appAuthority.hasRole(app, IAppAuthority.Role.PAUSER, pauser));
    }

    function test_stopApp_pauserCanCall() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address pauser = makeAddr("pauser");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.PAUSER, pauser);

        vm.prank(pauser);
        appController.stopApp(app);
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STOPPED));
    }

    function test_stopApp_developerCannotCall() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address dev = makeAddr("dev");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.DEVELOPER, dev);

        vm.prank(dev);
        vm.expectRevert(IAppController.InvalidTeamRole.selector);
        appController.stopApp(app);
    }

    function test_updateAppMetadataURI_developerCanCall() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address dev = makeAddr("dev");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.DEVELOPER, dev);

        vm.prank(dev);
        appController.updateAppMetadataURI(app, "ipfs://new");
    }

    function test_updateAppMetadataURI_pauserCannotCall() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address pauser = makeAddr("pauser");
        vm.prank(developer);
        appAuthority.grantRole(app, IAppAuthority.Role.PAUSER, pauser);

        vm.prank(pauser);
        vm.expectRevert(IAppController.InvalidTeamRole.selector);
        appController.updateAppMetadataURI(app, "ipfs://new");
    }

    function test_transferOwnership_grantsAdminToNewOwner() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address newOwner = makeAddr("newOwner");
        _setMaxActiveAppsPerUser(newOwner, 10);

        vm.prank(developer);
        appController.transferOwnership(app, newOwner);

        assertTrue(
            appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, newOwner),
            "new owner must be granted ADMIN automatically"
        );
    }

    function test_transferOwnership_removesPreviousOwnerFromAdmin() public {
        // Fix for audit A-3 / V-10: the previous owner must be removed from
        // the ADMIN set on transfer. Otherwise an old EOA (or an old
        // Timelock post-handoff) retains operational powers and, in the
        // non-timelocked case, can re-grab the app.
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address newOwner = makeAddr("newOwner");
        _setMaxActiveAppsPerUser(newOwner, 10);

        vm.prank(developer);
        appController.transferOwnership(app, newOwner);

        assertFalse(
            appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, developer),
            "previous owner must be removed from ADMIN on transfer"
        );
        assertTrue(appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, newOwner), "new owner must still be ADMIN");
    }

    function test_transferOwnership_previousOwnerLosesCriticalPower() public {
        // After transfer, the previous owner cannot perform critical ops
        // even though they used to. This is the direct user-visible
        // consequence of the owner-gate + ADMIN cleanup together.
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        address newOwner = makeAddr("newOwner");
        _setMaxActiveAppsPerUser(newOwner, 10);

        vm.prank(developer);
        appController.transferOwnership(app, newOwner);

        // Previous owner cannot upgrade, terminate, or transfer.
        vm.prank(developer);
        vm.expectRevert(IAppController.NotCreator.selector);
        appController.upgradeApp(app, _assembleRelease());

        vm.prank(developer);
        vm.expectRevert(IAppController.NotCreator.selector);
        appController.terminateApp(app);

        vm.prank(developer);
        vm.expectRevert(IAppController.NotCreator.selector);
        appController.transferOwnership(app, makeAddr("someoneElse"));
    }

    function test_migrateAdmins_seedsAdminFromPermissionController() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        // Set up two UAM admins who are NOT yet in the team-role set.
        address pcAdminA = makeAddr("pcAdminA");
        address pcAdminB = makeAddr("pcAdminB");
        vm.prank(developer);
        permissionController.acceptAdmin(address(app));
        vm.prank(developer);
        permissionController.addPendingAdmin(address(app), pcAdminA);
        vm.prank(pcAdminA);
        permissionController.acceptAdmin(address(app));
        vm.prank(developer);
        permissionController.addPendingAdmin(address(app), pcAdminB);
        vm.prank(pcAdminB);
        permissionController.acceptAdmin(address(app));

        assertFalse(appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, pcAdminA));
        assertFalse(appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, pcAdminB));

        IApp[] memory apps = new IApp[](1);
        apps[0] = app;
        vm.prank(admin);
        appController.migrateAppsToAppAuthority(apps);

        assertTrue(appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, pcAdminA));
        assertTrue(appAuthority.hasRole(app, IAppAuthority.Role.ADMIN, pcAdminB));
    }

    function test_migrateAdmins_callerMustBePlatformAdmin() public {
        vm.prank(developer);
        IApp app = appController.createApp(SALT, _assembleRelease());

        IApp[] memory apps = new IApp[](1);
        apps[0] = app;
        vm.prank(user);
        vm.expectRevert();
        appController.migrateAppsToAppAuthority(apps);
    }
}
