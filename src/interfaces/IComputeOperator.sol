// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IStrategy} from "@eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

interface IComputeOperator {
    /// @notice Thrown when the caller is not the ComputeAVS
    error NotComputeAVS();

    /// @notice Thrown when the caller is not the AppController
    error NotAppController();

    /// @notice Thrown when the initial allocation has already been performed
    error InitialAllocationAlreadyComplete();

    /**
     * @notice Initializes the ComputeOperator
     * @param operatorMetadataURI The metadata URI for the operator
     */
    function initialize(string calldata operatorMetadataURI) external;

    /**
     * @notice Registers this operator for a specified operator set
     * @param operatorSetId The operator set ID within the app to register for
     * @dev This function can only be called by the AppController
     */
    function registerForOperatorSet(uint32 operatorSetId) external;

    /**
     * @notice Performs a one-time allocation of 10% of stake to operator set 0
     * @dev This function can only be called when the initial allocation has not yet been performed
     * @dev Allocates 10% (1e17 in WAD format) of the configured strategyToSlash to operator set 0
     */
    function performInitialAllocation() external;
}
