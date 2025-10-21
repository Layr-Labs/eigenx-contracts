// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IAppController} from "../src/interfaces/IAppController.sol";
import {ComputeDeployer} from "./utils/ComputeDeployer.sol";
import {ComputeAVSRegistrar} from "../src/ComputeAVSRegistrar.sol";
import {IComputeAVSRegistrar} from "../src/interfaces/IComputeAVSRegistrar.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IReleaseManagerTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {IAllowlistErrors} from "@eigenlayer-middleware/src/interfaces/IAllowlist.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract ComputeAVSRegistrarTest is ComputeDeployer {
    address public operator1 = makeAddr("operator1");
    address public operator2 = makeAddr("operator2");
    address public unauthorizedCaller = makeAddr("unauthorizedCaller");

    uint32 public constant TEST_OPERATOR_SET_ID = 1;

    function setUp() public {
        _deployContracts();

        // Set maximum active apps for the developer account used in tests
        vm.prank(admin);
        appController.setMaxActiveAppsPerUser(developer, 100);
    }

    // ========== Allowlist Management Tests ==========

    function test_addOperatorToAllowlist() public {
        OperatorSet memory operatorSet = OperatorSet({avs: address(computeAVSRegistrar), id: TEST_OPERATOR_SET_ID});

        // AppController (owner) adds operator to allowlist
        vm.prank(address(appController));
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator1);

        // Verify operator is allowed
        assertTrue(computeAVSRegistrar.isOperatorAllowed(operatorSet, operator1));
        assertFalse(computeAVSRegistrar.isOperatorAllowed(operatorSet, operator2));
    }

    function test_addOperatorToAllowlist_alreadyInList() public {
        OperatorSet memory operatorSet = OperatorSet({avs: address(computeAVSRegistrar), id: TEST_OPERATOR_SET_ID});

        // Add operator once
        vm.prank(address(appController));
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator1);

        // Try to add again - should revert
        vm.prank(address(appController));
        vm.expectRevert(IAllowlistErrors.OperatorAlreadyInAllowlist.selector);
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator1);
    }

    function test_removeOperatorFromAllowlist() public {
        OperatorSet memory operatorSet = OperatorSet({avs: address(computeAVSRegistrar), id: TEST_OPERATOR_SET_ID});

        // Add operator first
        vm.prank(address(appController));
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator1);
        assertTrue(computeAVSRegistrar.isOperatorAllowed(operatorSet, operator1));

        // Remove operator
        vm.prank(address(appController));
        computeAVSRegistrar.removeOperatorFromAllowlist(operatorSet, operator1);

        // Verify operator is no longer allowed
        assertFalse(computeAVSRegistrar.isOperatorAllowed(operatorSet, operator1));
    }

    function test_removeOperatorFromAllowlist_notInList() public {
        OperatorSet memory operatorSet = OperatorSet({avs: address(computeAVSRegistrar), id: TEST_OPERATOR_SET_ID});

        // Try to remove operator that was never added - should revert
        vm.prank(address(appController));
        vm.expectRevert(IAllowlistErrors.OperatorNotInAllowlist.selector);
        computeAVSRegistrar.removeOperatorFromAllowlist(operatorSet, operator1);
    }

    function test_getAllowedOperators() public {
        OperatorSet memory operatorSet = OperatorSet({avs: address(computeAVSRegistrar), id: TEST_OPERATOR_SET_ID});

        // Add multiple operators
        vm.startPrank(address(appController));
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator1);
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator2);
        vm.stopPrank();

        // Get all allowed operators
        address[] memory allowedOperators = computeAVSRegistrar.getAllowedOperators(operatorSet);

        assertEq(allowedOperators.length, 2);
        assertTrue(allowedOperators[0] == operator1 && allowedOperators[1] == operator2);
    }

    // ========== Permission Tests ==========

    function test_addOperatorToAllowlist_unauthorizedCaller() public {
        OperatorSet memory operatorSet = OperatorSet({avs: address(computeAVSRegistrar), id: TEST_OPERATOR_SET_ID});

        // Unauthorized caller tries to add operator - should revert
        vm.prank(unauthorizedCaller);
        vm.expectRevert("Ownable: caller is not the owner");
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator1);
    }

    function test_removeOperatorFromAllowlist_unauthorizedCaller() public {
        OperatorSet memory operatorSet = OperatorSet({avs: address(computeAVSRegistrar), id: TEST_OPERATOR_SET_ID});

        // First add an operator as authorized caller
        vm.prank(address(appController));
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator1);

        // Unauthorized caller tries to remove operator - should revert
        vm.prank(unauthorizedCaller);
        vm.expectRevert("Ownable: caller is not the owner");
        computeAVSRegistrar.removeOperatorFromAllowlist(operatorSet, operator1);
    }

    function test_multipleOperatorSets_differentAllowlists() public {
        // Create two operator sets
        vm.startPrank(address(appController));
        uint32 operatorSetId1 = computeAVSRegistrar.assignOperatorSetId();
        computeAVSRegistrar.createOperatorSet(operatorSetId1);

        uint32 operatorSetId2 = computeAVSRegistrar.assignOperatorSetId();
        computeAVSRegistrar.createOperatorSet(operatorSetId2);
        vm.stopPrank();

        OperatorSet memory operatorSet1 = OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId1});

        OperatorSet memory operatorSet2 = OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId2});

        // Add operator1 to set1, operator2 to set2
        vm.startPrank(address(appController));
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet1, operator1);
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet2, operator2);
        vm.stopPrank();

        // Verify allowlists are separate
        assertTrue(computeAVSRegistrar.isOperatorAllowed(operatorSet1, operator1));
        assertFalse(computeAVSRegistrar.isOperatorAllowed(operatorSet1, operator2));

        assertFalse(computeAVSRegistrar.isOperatorAllowed(operatorSet2, operator1));
        assertTrue(computeAVSRegistrar.isOperatorAllowed(operatorSet2, operator2));
    }

    // ========== Integration Tests ==========

    function test_appCreation_withAllowlist() public {
        // Create an app and verify its operator set is created
        vm.prank(developer);
        appController.createApp(
            keccak256("test_app"),
            IAppController.Release({
                rmsRelease: IReleaseManagerTypes.Release({
                    artifacts: new IReleaseManagerTypes.Artifact[](1), upgradeByTime: uint32(block.timestamp + 1 days)
                }),
                publicEnv: bytes(""),
                encryptedEnv: bytes("")
            })
        );

        // The app should get operator set ID 1 (0 is used for ComputeOperator)
        OperatorSet memory operatorSet1 = OperatorSet({avs: address(computeAVSRegistrar), id: 1});

        // Initially, no operators should be in the allowlist for this new operator set
        address[] memory allowedOperators = computeAVSRegistrar.getAllowedOperators(operatorSet1);
        assertEq(allowedOperators.length, 0);

        // AppController should be able to add operators to this set's allowlist
        vm.prank(address(appController));
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet1, operator1);

        assertTrue(computeAVSRegistrar.isOperatorAllowed(operatorSet1, operator1));
    }

    // ========== Access Control Tests ==========

    function test_assignOperatorSetId_onlyAppController() public {
        // Non-AppController tries to assign operator set ID - should revert
        vm.prank(unauthorizedCaller);
        vm.expectRevert(abi.encodeWithSelector(IComputeAVSRegistrar.NotAppController.selector));
        computeAVSRegistrar.assignOperatorSetId();
    }

    function test_createOperatorSet_onlyAppController() public {
        // Create a valid operator set ID first
        vm.prank(address(appController));
        uint32 operatorSetId = computeAVSRegistrar.assignOperatorSetId();

        // Non-AppController tries to create operator set - should revert
        vm.prank(unauthorizedCaller);
        vm.expectRevert(abi.encodeWithSelector(IComputeAVSRegistrar.NotAppController.selector));
        computeAVSRegistrar.createOperatorSet(operatorSetId);
    }

    function test_createOperatorSet_unassignedId() public {
        // Try to create operator set with unassigned ID - should revert
        vm.prank(address(appController));
        vm.expectRevert(abi.encodeWithSelector(IComputeAVSRegistrar.OperatorSetIdNotAssigned.selector));
        computeAVSRegistrar.createOperatorSet(999); // Using a high number that wasn't assigned
    }

    function test_createOperatorSet_success() public {
        // Assign and create operator set successfully
        vm.startPrank(address(appController));
        uint32 operatorSetId = computeAVSRegistrar.assignOperatorSetId();

        // Should not revert
        computeAVSRegistrar.createOperatorSet(operatorSetId);
        vm.stopPrank();

        // Verify the operator set was created by checking we can add operators to its allowlist
        OperatorSet memory operatorSet = OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId});

        vm.prank(address(appController));
        computeAVSRegistrar.addOperatorToAllowlist(operatorSet, operator1);
        assertTrue(computeAVSRegistrar.isOperatorAllowed(operatorSet, operator1));
    }

    // ========== AVSRegistrarAsIdentifier Tests ==========

    function test_supportsAVS() public view {
        // Should support its own address
        assertTrue(computeAVSRegistrar.supportsAVS(address(computeAVSRegistrar)));

        // Should not support other addresses
        assertFalse(computeAVSRegistrar.supportsAVS(address(appController)));
        assertFalse(computeAVSRegistrar.supportsAVS(operator1));
        assertFalse(computeAVSRegistrar.supportsAVS(address(0)));
    }

    // ========== State Management Tests ==========

    function test_nextOperatorSetId_increments() public {
        // Check initial state (operator set 0 was created during initialization)
        uint32 initialId = computeAVSRegistrar.nextOperatorSetId();
        assertEq(initialId, 1); // Should be 1 since 0 was used in initialization

        // Assign a new operator set ID
        vm.prank(address(appController));
        uint32 assignedId = computeAVSRegistrar.assignOperatorSetId();
        assertEq(assignedId, initialId);

        // Check that nextOperatorSetId incremented
        assertEq(computeAVSRegistrar.nextOperatorSetId(), initialId + 1);

        // Assign another one
        vm.prank(address(appController));
        uint32 secondAssignedId = computeAVSRegistrar.assignOperatorSetId();
        assertEq(secondAssignedId, initialId + 1);
        assertEq(computeAVSRegistrar.nextOperatorSetId(), initialId + 2);
    }

    // ========== Storage Tests ==========

    function test_immutableVariables() public view {
        // Verify immutable variables are set correctly
        assertTrue(address(computeAVSRegistrar.allocationManager()) != address(0));
        assertTrue(address(computeAVSRegistrar.keyRegistrar()) != address(0));
        assertEq(address(computeAVSRegistrar.appController()), address(appController));
    }
}
