// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {IDelegationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {
    IAllocationManager,
    IAllocationManagerTypes
} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IStrategy} from "@eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {
    IEigenDACertVerifierRouter
} from "@eigenda-contracts/src/integrations/cert/interfaces/IEigenDACertVerifierRouter.sol";
import {IComputeAVSRegistrar} from "../interfaces/IComputeAVSRegistrar.sol";
import {IComputeOperator} from "../interfaces/IComputeOperator.sol";
import {ICloudReportCompendium} from "../interfaces/ICloudReportCompendium.sol";

abstract contract CloudReportCompendiumStorage is ICloudReportCompendium {
    /// STRUCTS
    struct ConstructorParams {
        IDelegationManager delegationManager;
        IAllocationManager allocationManager;
        IEigenDACertVerifierRouter certVerifierRouter;
        IComputeAVSRegistrar computeAVSRegistrar;
        IComputeOperator computeOperator;
        address reportAttester;
        uint32 operatorSetId;
        string referenceProjectId;
        IStrategy strategyToSlash;
    }

    /// CONSTANTS
    /// @inheritdoc ICloudReportCompendium
    bytes32 public constant CLOUD_REPORT_TYPEHASH =
        keccak256("CloudReport(string projectId,uint32 fromTimestamp,uint32 toTimestamp,bytes eigendaCert)");

    /// @notice EigenDA cert verification success status code
    uint8 internal constant EIGENDA_CERT_SUCCESS = 1;

    /// IMMUTABLES

    /// @inheritdoc ICloudReportCompendium
    IDelegationManager public immutable delegationManager;

    /// @inheritdoc ICloudReportCompendium
    IAllocationManager public immutable allocationManager;

    /// @inheritdoc ICloudReportCompendium
    IEigenDACertVerifierRouter public immutable certVerifierRouter;

    /// @inheritdoc ICloudReportCompendium
    IComputeAVSRegistrar public immutable computeAVSRegistrar;

    /// @inheritdoc ICloudReportCompendium
    IComputeOperator public immutable computeOperator;

    /// @inheritdoc ICloudReportCompendium
    address public immutable reportAttester;

    /// @inheritdoc ICloudReportCompendium
    uint32 public immutable operatorSetId;

    /// @inheritdoc ICloudReportCompendium
    bytes32 public immutable referenceProjectIdHash;

    /// @inheritdoc ICloudReportCompendium
    IStrategy public immutable strategyToSlash;

    /// STORAGE

    /// @inheritdoc ICloudReportCompendium
    uint32 public reportFreshnessThreshold;

    /// @inheritdoc ICloudReportCompendium
    uint256 public slashAmountTokens;

    /// @inheritdoc ICloudReportCompendium
    uint32 public minSlashInterval;

    /// @inheritdoc ICloudReportCompendium
    uint32 public lastSlashTimestamp;

    /// @notice The latest cloud report submission
    CloudReportSubmission internal _latestReportSubmission;

    constructor(ConstructorParams memory params) {
        delegationManager = params.delegationManager;
        allocationManager = params.allocationManager;
        certVerifierRouter = params.certVerifierRouter;
        computeAVSRegistrar = params.computeAVSRegistrar;
        computeOperator = params.computeOperator;
        reportAttester = params.reportAttester;
        operatorSetId = params.operatorSetId;
        referenceProjectIdHash = keccak256(abi.encodePacked(params.referenceProjectId));
        strategyToSlash = params.strategyToSlash;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[43] private __gap;
}
