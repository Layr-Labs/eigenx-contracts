// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ComputeDeployer} from "./utils/ComputeDeployer.sol";
import {ComputeOperator} from "../src/ComputeOperator.sol";
import {IComputeOperator} from "../src/interfaces/IComputeOperator.sol";
import {IDelegationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IAllowlistErrors} from "@eigenlayer-middleware/src/interfaces/IAllowlist.sol";

contract ComputeOperatorTest is ComputeDeployer {
    address public unauthorizedCaller = makeAddr("unauthorizedCaller");
    address public newAdmin = makeAddr("newAdmin");
    string public constant METADATA_URI = "https://example.com/operator-metadata";

    function setUp() public {
        _deployContracts();

        // Set maximum active apps for developer
        vm.prank(admin);
        appController.setMaxActiveAppsPerUser(developer, 100);
    }

    // ========== Initialization Tests ==========

    function test_initialize_success() public view {
        // Verify the ComputeOperator was initialized correctly during setup
        assertTrue(delegationManager.isOperator(address(computeOperator)));
    }

    function test_initialize_alreadyInitialized() public {
        // ComputeOperator is already initialized in setUp via _deployContracts

        // Try to initialize again - should revert
        vm.expectRevert("Initializable: contract is already initialized");
        computeOperator.initialize(METADATA_URI);
    }

    // ========== Access Control Tests ==========

    function test_registerForOperatorSet_onlyAppController() public {
        // Unauthorized caller tries to register for operator set - should revert
        vm.prank(unauthorizedCaller);
        vm.expectRevert(abi.encodeWithSelector(IComputeOperator.NotAppController.selector));
        computeOperator.registerForOperatorSet(1);
    }

    function test_registerForOperatorSet_fromAppController() public {
        vm.startPrank(address(appController));
        uint32 operatorSetId = _createAndAllowlistOperatorSet(address(computeOperator));
        computeOperator.registerForOperatorSet(operatorSetId);
        vm.stopPrank();
    }

    // ========== Registration Tests ==========

    function test_registerForOperatorSet_multipleRegistrations() public {
        vm.startPrank(address(appController));
        uint32 operatorSetId1 = _createAndAllowlistOperatorSet(address(computeOperator));
        uint32 operatorSetId2 = _createAndAllowlistOperatorSet(address(computeOperator));

        computeOperator.registerForOperatorSet(operatorSetId1);
        computeOperator.registerForOperatorSet(operatorSetId2);
        vm.stopPrank();
    }

    function test_registerForOperatorSet_withoutAllowlist() public {
        // Create operator set but don't add to allowlist
        vm.startPrank(address(appController));
        uint32 operatorSetId = computeAVSRegistrar.assignOperatorSetId();
        computeAVSRegistrar.createOperatorSet(operatorSetId);

        // Try to register without being in allowlist - should revert
        vm.expectRevert(IAllowlistErrors.OperatorNotInAllowlist.selector);
        computeOperator.registerForOperatorSet(operatorSetId);
        vm.stopPrank();
    }

    // ========== Immutable Variables Tests ==========

    function test_immutableVariables() public view {
        // Verify all immutable variables are set correctly
        assertEq(address(computeOperator.delegationManager()), address(delegationManager));
        assertEq(address(computeOperator.allocationManager()), address(allocationManager));
        assertEq(address(computeOperator.permissionController()), address(permissionController));
        assertEq(computeOperator.appController(), address(appController));
        assertEq(computeOperator.computeAVSRegistrar(), address(computeAVSRegistrar));
    }

    // ========== Helper Functions ==========

    function _createAndAllowlistOperatorSet(address operator) internal returns (uint32) {
        uint32 operatorSetId = computeAVSRegistrar.assignOperatorSetId();
        computeAVSRegistrar.createOperatorSet(operatorSetId);
        computeAVSRegistrar.addOperatorToAllowlist(
            OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId}), operator
        );
        return operatorSetId;
    }
}
