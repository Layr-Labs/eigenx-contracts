// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IDelegationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {
    IAllocationManager,
    IAllocationManagerTypes
} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {SemVerMixin} from "@eigenlayer-contracts/src/contracts/mixins/SemVerMixin.sol";
import {ComputeOperatorStorage} from "./storage/ComputeOperatorStorage.sol";
import {IComputeOperator} from "./interfaces/IComputeOperator.sol";

contract ComputeOperator is Initializable, SemVerMixin, ComputeOperatorStorage {
    /// @notice Ensures that the function is only callable by the app controller
    modifier onlyAppController() {
        require(msg.sender == appController, NotAppController());
        _;
    }

    /**
     * @param _version The version string to use for this contract's domain separator
     * @param _delegationManager The DelegationManager contract from EigenLayer
     * @param _allocationManager The AllocationManager contract from EigenLayer
     * @param _permissionController The PermissionController contract from EigenLayer
     * @param _appController The AppController contract
     * @param _computeAVSRegistrar The ComputeAVSRegistrar contract
     */
    constructor(
        string memory _version,
        IDelegationManager _delegationManager,
        IAllocationManager _allocationManager,
        IPermissionController _permissionController,
        address _appController,
        address _computeAVSRegistrar
    )
        SemVerMixin(_version)
        ComputeOperatorStorage(
            _delegationManager, _allocationManager, _permissionController, _appController, _computeAVSRegistrar
        )
    {
        _disableInitializers();
    }

    /// @inheritdoc IComputeOperator
    function initialize(string calldata operatorMetadataURI) external initializer {
        // Register as an EigenLayer operator
        delegationManager.registerAsOperator({
            initDelegationApprover: address(0), allocationDelay: 0, metadataURI: operatorMetadataURI
        });
    }

    /// @inheritdoc IComputeOperator
    function registerForOperatorSet(uint32 operatorSetId) external onlyAppController {
        uint32[] memory operatorSetIds = new uint32[](1);
        operatorSetIds[0] = operatorSetId;

        IAllocationManagerTypes.RegisterParams memory params = IAllocationManagerTypes.RegisterParams({
            avs: computeAVSRegistrar, operatorSetIds: operatorSetIds, data: new bytes(0)
        });

        allocationManager.registerForOperatorSets(address(this), params);
    }
}
