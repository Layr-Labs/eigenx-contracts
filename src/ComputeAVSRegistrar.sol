// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {
    IAllocationManager,
    IAllocationManagerTypes
} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar, IKeyRegistrarTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {SemVerMixin} from "@eigenlayer-contracts/src/contracts/mixins/SemVerMixin.sol";
import {AVSRegistrar} from "@eigenlayer-middleware/src/middlewareV2/registrar/AVSRegistrar.sol";
import {IAVSRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IAVSRegistrar.sol";
import {Allowlist} from "@eigenlayer-middleware/src/middlewareV2/registrar/modules/Allowlist.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IStrategy} from "@eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {IAppController} from "./interfaces/IAppController.sol";
import {ComputeAVSRegistrarStorage} from "./storage/ComputeAVSRegistrarStorage.sol";
import {IComputeAVSRegistrar} from "./interfaces/IComputeAVSRegistrar.sol";

contract ComputeAVSRegistrar is SemVerMixin, AVSRegistrar, Allowlist, ComputeAVSRegistrarStorage {
    /// @notice Ensures that the function is only callable by the app controller
    modifier onlyAppController() {
        require(msg.sender == address(appController), NotAppController());
        _;
    }

    /**
     * @param _version The version string to use for this contract's domain separator
     * @param _allocationManager The AllocationManager contract from EigenLayer
     * @param _permissionController The PermissionController contract from EigenLayer
     * @param _keyRegistrar The KeyRegistrar contract from EigenLayer
     * @param _appController The AppController contract
     */
    constructor(
        string memory _version,
        IAllocationManager _allocationManager,
        IPermissionController _permissionController,
        IKeyRegistrar _keyRegistrar,
        IAppController _appController
    )
        SemVerMixin(_version)
        AVSRegistrar(_allocationManager, _keyRegistrar)
        ComputeAVSRegistrarStorage(_permissionController, _appController)
    {
        _disableInitializers();
    }

    /// @inheritdoc IComputeAVSRegistrar
    function initialize(string memory metadataURI) public initializer {
        // Set up allowlist management
        __AVSRegistrar_init(address(this));
        __Allowlist_init(address(appController));

        // Set the metadataURI and the registrar for the AVS to this registrar contract
        allocationManager.updateAVSMetadataURI(address(this), metadataURI);
        allocationManager.setAVSRegistrar(address(this), this);

        // add this contract itself as an admin
        permissionController.addPendingAdmin(address(this), address(this));
        permissionController.acceptAdmin(address(this));

        // add the app controller as an admin
        permissionController.addPendingAdmin(address(this), address(appController));
    }

    /// @inheritdoc IComputeAVSRegistrar
    function assignOperatorSetId() external onlyAppController returns (uint32 operatorSetId) {
        operatorSetId = nextOperatorSetId++;
    }

    /// @inheritdoc IComputeAVSRegistrar
    function createOperatorSet(uint32 operatorSetId) external onlyAppController {
        // Ensure we don't create operator sets beyond what has been assigned
        require(operatorSetId < nextOperatorSetId, OperatorSetIdNotAssigned());

        // Create the operator set with no strategies (can be configured later)
        IAllocationManagerTypes.CreateSetParams[] memory params = new IAllocationManagerTypes.CreateSetParams[](1);
        params[0] =
            IAllocationManagerTypes.CreateSetParams({operatorSetId: operatorSetId, strategies: new IStrategy[](0)});
        allocationManager.createOperatorSets(address(this), params);
    }

    /// @notice Before registering operator, check if the operator is in the allowlist
    function _beforeRegisterOperator(address operator, uint32[] calldata operatorSetIds, bytes calldata data)
        internal
        override
    {
        super._beforeRegisterOperator(operator, operatorSetIds, data);

        for (uint32 i; i < operatorSetIds.length; ++i) {
            require(
                isOperatorAllowed(OperatorSet({avs: address(this), id: operatorSetIds[i]}), operator),
                OperatorNotInAllowlist()
            );
        }
    }

    /// @notice Override to skip key registrar checks
    function registerOperator(
        address operator,
        address, /* avs */
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external virtual override(AVSRegistrar, IAVSRegistrar) onlyAllocationManager {
        _beforeRegisterOperator(operator, operatorSetIds, data);

        _afterRegisterOperator(operator, operatorSetIds, data);

        emit OperatorRegistered(operator, operatorSetIds);
    }
}
