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

    function test_suspendApp() public {
        // Create an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("suspend_test"), _assembleRelease());
        _acceptAppAdmin(app);

        // Verify app is started
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STARTED));

        // Expect the AppSuspended event
        vm.expectEmit(true, false, false, false);
        emit AppSuspended(app);

        // Suspend the app
        vm.prank(developer);
        appController.suspendApp(app);

        // Verify status changed to SUSPENDED
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.SUSPENDED));
    }

    function test_suspendApp_notAuthorized() public {
        // Create an app as developer
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("suspend_unauthorized"), _assembleRelease());

        // Try to suspend as unauthorized user
        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.suspendApp(app);
    }

    function test_suspendApp_alreadySuspended() public {
        // Create and suspend an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("suspend_already"), _assembleRelease());
        _acceptAppAdmin(app);

        vm.prank(developer);
        appController.suspendApp(app);

        // Try to suspend again - should revert (appIsActive modifier checks for SUSPENDED)
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.suspendApp(app);
    }

    function test_suspendApp_decreasesCapacity() public {
        // Setup limits
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 10);

        // User creates an app
        vm.prank(user);
        IApp app = appController.createApp(keccak256("capacity_suspend"), _assembleRelease());

        // Verify counts increased
        assertEq(appController.getActiveAppCount(user), 1);
        assertEq(appController.globalActiveAppCount(), 1);

        // Accept admin to suspend
        vm.prank(user);
        permissionController.acceptAdmin(address(app));

        // Suspend the app
        vm.prank(user);
        appController.suspendApp(app);

        // Verify counts decreased
        assertEq(appController.getActiveAppCount(user), 0);
        assertEq(appController.globalActiveAppCount(), 0);
    }

    function test_suspendAppByAdmin() public {
        // Create an app as developer
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("suspend_admin"), _assembleRelease());

        // Expect the AppSuspendedByAdmin event
        vm.expectEmit(true, false, false, false);
        emit AppSuspendedByAdmin(app);

        // Suspend app as admin
        vm.prank(admin);
        appController.suspendAppByAdmin(app);

        // Verify status changed to SUSPENDED
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.SUSPENDED));

        // Verify capacity was freed
        assertEq(appController.getActiveAppCount(developer), 0);
        assertEq(appController.globalActiveAppCount(), 0);
    }

    function test_suspendAppByAdmin_notAuthorized() public {
        // Create an app as developer
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("suspend_admin_unauthorized"), _assembleRelease());

        // Try to suspend as unauthorized user (not admin)
        vm.prank(user);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        appController.suspendAppByAdmin(app);
    }

    function test_suspendApp_freesUpCapacityForNewApps() public {
        // Set user to have only 1 active app max
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(user, 1);

        // User creates 1 app (reaches limit)
        vm.prank(user);
        IApp app1 = appController.createApp(keccak256("capacity_free_1"), _assembleRelease());

        // Try to create a second app - should fail
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(IAppController.MaxActiveAppsExceeded.selector));
        appController.createApp(keccak256("capacity_free_2"), _assembleRelease());

        // Accept admin and suspend app1
        vm.prank(user);
        permissionController.acceptAdmin(address(app1));
        vm.prank(user);
        appController.suspendApp(app1);

        // Now user should be able to create a new app
        vm.prank(user);
        IApp app2 = appController.createApp(keccak256("capacity_free_2"), _assembleRelease());

        // Verify counts
        assertEq(appController.getActiveAppCount(user), 1);
        assertEq(appController.globalActiveAppCount(), 1);

        // Verify app1 is suspended and app2 is started
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.STARTED));
    }

    function test_startApp_fromSuspended() public {
        // Create and suspend an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("start_suspended"), _assembleRelease());
        _acceptAppAdmin(app);

        vm.prank(developer);
        appController.suspendApp(app);

        // Verify app is suspended and capacity is freed
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.SUSPENDED));
        assertEq(appController.getActiveAppCount(developer), 0);

        // Start the app again
        vm.prank(developer);
        appController.startApp(app);

        // Verify app is started and capacity is allocated
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STARTED));
        assertEq(appController.getActiveAppCount(developer), 1);
        assertEq(appController.globalActiveAppCount(), 1);
    }

    function test_startApp_fromSuspended_delegatedUser() public {
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(developer, 10);

        // Developer creates an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("delegated_suspend"), _assembleRelease());

        address user2 = makeAddr("user2");

        // Developer accepts admin and grants user2 permission via appointee
        vm.prank(developer);
        permissionController.acceptAdmin(address(app));
        vm.prank(developer);
        permissionController.setAppointee(address(app), user2, address(appController), IAppController.startApp.selector);

        // Developer suspends the app
        vm.prank(developer);
        appController.suspendApp(app);

        // Verify counts after suspend
        assertEq(appController.getActiveAppCount(developer), 0, "Developer count should be 0 after suspend");
        assertEq(appController.getActiveAppCount(user2), 0, "User2 count should be 0");

        // User2 (delegated) starts the app
        vm.prank(user2);
        appController.startApp(app);

        // Developer's count should increase (they're the creator), NOT user2's
        assertEq(appController.getActiveAppCount(developer), 1, "Developer (creator) count should be 1");
        assertEq(appController.getActiveAppCount(user2), 0, "User2 (delegated) count should still be 0");
        assertEq(appController.globalActiveAppCount(), 1, "Global count should be 1");
        assertEq(uint256(appController.getAppStatus(app)), uint256(IAppController.AppStatus.STARTED));
    }

    function test_startApp_fromSuspended_globalLimitReached() public {
        // Setup limits
        _setGlobalMaxActiveApps(1);
        _setMaxActiveAppsPerUser(developer, 10);

        // Create and suspend an app
        vm.prank(developer);
        IApp app1 = appController.createApp(keccak256("global_limit_1"), _assembleRelease());
        _acceptAppAdmin(app1);
        vm.prank(developer);
        appController.suspendApp(app1);

        // Create another app to fill global capacity
        vm.prank(developer);
        appController.createApp(keccak256("global_limit_2"), _assembleRelease());

        // Try to start the suspended app - should fail due to global limit
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.GlobalMaxActiveAppsExceeded.selector));
        appController.startApp(app1);
    }

    function test_startApp_fromSuspended_userLimitReached() public {
        // Setup limits
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(developer, 1);

        // Create and suspend an app
        vm.prank(developer);
        IApp app1 = appController.createApp(keccak256("user_limit_1"), _assembleRelease());
        _acceptAppAdmin(app1);
        vm.prank(developer);
        appController.suspendApp(app1);

        // Create another app to fill user capacity
        vm.prank(developer);
        appController.createApp(keccak256("user_limit_2"), _assembleRelease());

        // Try to start the suspended app - should fail due to user limit
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.MaxActiveAppsExceeded.selector));
        appController.startApp(app1);
    }

    function test_stopApp_rejectsSuspended() public {
        // Create and suspend an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("stop_suspended"), _assembleRelease());
        _acceptAppAdmin(app);

        vm.prank(developer);
        appController.suspendApp(app);

        // Try to stop a suspended app - should fail (appIsActive modifier)
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.stopApp(app);
    }

    function test_upgradeApp_rejectsSuspended() public {
        // Create and suspend an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("upgrade_suspended"), _assembleRelease());
        _acceptAppAdmin(app);

        vm.prank(developer);
        appController.suspendApp(app);

        // Try to upgrade a suspended app - should fail (appIsActive modifier)
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.upgradeApp(app, _assembleRelease());
    }

    function test_terminateApp_rejectsSuspended() public {
        // Create and suspend an app
        vm.prank(developer);
        IApp app = appController.createApp(keccak256("terminate_suspended"), _assembleRelease());
        _acceptAppAdmin(app);

        vm.prank(developer);
        appController.suspendApp(app);

        // Try to terminate a suspended app - should fail (appIsActive modifier)
        vm.prank(developer);
        vm.expectRevert(abi.encodeWithSelector(IAppController.InvalidAppStatus.selector));
        appController.terminateApp(app);
    }

    function test_suspendAndRestartCycle() public {
        // Setup limits
        _setGlobalMaxActiveApps(100);
        _setMaxActiveAppsPerUser(developer, 2);

        // Create two apps
        vm.prank(developer);
        IApp app1 = appController.createApp(keccak256("cycle_1"), _assembleRelease());
        _acceptAppAdmin(app1);

        vm.prank(developer);
        IApp app2 = appController.createApp(keccak256("cycle_2"), _assembleRelease());
        _acceptAppAdmin(app2);

        // Verify both are active
        assertEq(appController.getActiveAppCount(developer), 2);

        // Suspend both apps
        vm.prank(developer);
        appController.suspendApp(app1);
        vm.prank(developer);
        appController.suspendApp(app2);

        // Verify capacity is freed
        assertEq(appController.getActiveAppCount(developer), 0);
        assertEq(appController.globalActiveAppCount(), 0);

        // Restart both apps
        vm.prank(developer);
        appController.startApp(app1);
        vm.prank(developer);
        appController.startApp(app2);

        // Verify capacity is allocated again
        assertEq(appController.getActiveAppCount(developer), 2);
        assertEq(appController.globalActiveAppCount(), 2);
        assertEq(uint256(appController.getAppStatus(app1)), uint256(IAppController.AppStatus.STARTED));
        assertEq(uint256(appController.getAppStatus(app2)), uint256(IAppController.AppStatus.STARTED));
    }
}
