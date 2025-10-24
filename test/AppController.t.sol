// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IAppController} from "../src/interfaces/IAppController.sol";
import {ComputeDeployer} from "./utils/ComputeDeployer.sol";
import {IApp} from "../src/interfaces/IApp.sol";
import {PermissionControllerMixin} from "@eigenlayer-contracts/src/contracts/mixins/PermissionControllerMixin.sol";
import {IReleaseManagerTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";

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

        // Try to upgrade as unauthorized user
        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
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

        // Try to start a non-existent app - should revert with InvalidAppStatus
        vm.prank(fakeAppAddress);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
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
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
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

        // Try to update metadata for a non-existent app - should revert
        vm.prank(fakeAppAddress);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
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
}
