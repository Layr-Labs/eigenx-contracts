// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IAVSRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IAVSRegistrar.sol";
import {IAllowlist} from "@eigenlayer-middleware/src/interfaces/IAllowlist.sol";

interface IComputeAVSRegistrar is IAVSRegistrar, IAllowlist {
    /// @notice Thrown when the caller is not the AppController
    error NotAppController();

    /// @notice Thrown when trying to create an operator set with an unassigned ID
    error OperatorSetIdNotAssigned();

    /**
     * @notice Initializes the ComputeAVSRegistrar
     * @param metadataURI The metadata URI for the AVS
     */
    function initialize(string calldata metadataURI) external;

    /**
     * @notice Assigns a unique operator set ID
     * @return operatorSetId The ID of the assigned operator set
     */
    function assignOperatorSetId() external returns (uint32);

    /**
     * @notice Creates and configures an operator set
     * @param operatorSetId The ID of the operator set to create
     */
    function createOperatorSet(uint32 operatorSetId) external;
}
