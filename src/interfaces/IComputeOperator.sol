// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

interface IComputeOperator {
    /// @notice Thrown when the caller is not the ComputeAVS
    error NotComputeAVS();

    /// @notice Thrown when the caller is not the AppController
    error NotAppController();

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
}
