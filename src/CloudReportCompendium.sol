// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {PermissionControllerMixin} from "@eigenlayer-contracts/src/contracts/mixins/PermissionControllerMixin.sol";
import {SignatureUtilsMixin} from "@eigenlayer-contracts/src/contracts/mixins/SignatureUtilsMixin.sol";
import {
    IAllocationManager,
    IAllocationManagerTypes
} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IDelegationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IStrategy} from "@eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {
    IEigenDACertVerifierRouter
} from "@eigenda-contracts/src/integrations/cert/interfaces/IEigenDACertVerifierRouter.sol";
import {IComputeAVSRegistrar} from "./interfaces/IComputeAVSRegistrar.sol";
import {IComputeOperator} from "./interfaces/IComputeOperator.sol";
import {CloudReportCompendiumStorage} from "./storage/CloudReportCompendiumStorage.sol";
import {ICloudReportCompendium} from "./interfaces/ICloudReportCompendium.sol";

contract CloudReportCompendium is
    Initializable,
    SignatureUtilsMixin,
    PermissionControllerMixin,
    CloudReportCompendiumStorage
{
    using ECDSA for bytes32;

    /**
     * @param _version The version string to use for this contract's domain separator
     * @param _permissionController The PermissionController contract address
     * @param params The constructor parameters struct
     */
    constructor(
        string memory _version,
        IPermissionController _permissionController,
        CloudReportCompendiumStorage.ConstructorParams memory params
    )
        SignatureUtilsMixin(_version)
        PermissionControllerMixin(_permissionController)
        CloudReportCompendiumStorage(params)
    {
        _disableInitializers();
    }

    /// @inheritdoc ICloudReportCompendium
    function initialize(uint32 _reportFreshnessThreshold, uint256 _slashAmountTokens) external initializer {
        _setReportFreshnessThreshold(_reportFreshnessThreshold);
        _setSlashAmountTokens(_slashAmountTokens);
    }

    /// @inheritdoc ICloudReportCompendium
    function setReportFreshnessThreshold(uint32 _reportFreshnessThreshold) external checkCanCall(address(this)) {
        _setReportFreshnessThreshold(_reportFreshnessThreshold);
    }

    /// @inheritdoc ICloudReportCompendium
    function setSlashAmountTokens(uint256 _slashAmountTokens) external checkCanCall(address(this)) {
        _setSlashAmountTokens(_slashAmountTokens);
    }

    /// @inheritdoc ICloudReportCompendium
    function submitReport(CloudReport calldata report, bytes calldata signature) external checkCanCall(address(this)) {
        // Verify project ID matches
        require(keccak256(abi.encodePacked(report.projectId)) == referenceProjectIdHash, InvalidProjectId());
        // Verify timestamps
        require(
            _latestReportSubmission.toTimestamp == 0 || report.fromTimestamp == _latestReportSubmission.toTimestamp,
            ReportChainTimestampMismatch()
        );
        require(report.fromTimestamp < report.toTimestamp, InvalidReportTimestamps());
        require(report.toTimestamp <= block.timestamp, ReportFromFuture());

        // Verify signature
        bytes32 reportHash = _hashCloudReport(report);
        bytes32 digest = _calculateSignableDigest(reportHash);
        address signer = digest.recover(signature);
        require(signer == reportAttester, InvalidSignature());

        // Store the report
        _latestReportSubmission = CloudReportSubmission({
            submissionTimestamp: uint32(block.timestamp),
            fromTimestamp: report.fromTimestamp,
            toTimestamp: report.toTimestamp,
            eigendaCertHash: bytes20(keccak256(report.eigendaCert))
        });

        emit ReportSubmitted(report, block.timestamp);
    }

    /// @inheritdoc ICloudReportCompendium
    function slash(bytes calldata eigendaCert) external {
        // Check eigenDaCert consistency
        require(bytes20(keccak256(eigendaCert)) == _latestReportSubmission.eigendaCertHash, InvalidCertificate());

        string memory reason;
        if (block.timestamp - _latestReportSubmission.submissionTimestamp >= reportFreshnessThreshold) {
            // check report freshness
            reason = "Report too stale";
        } else if (certVerifierRouter.checkDACert(eigendaCert) != EIGENDA_CERT_SUCCESS) {
            // check eigenda cert verification
            reason = "Invalid EigenDA certificate";
        } else {
            revert NoSlashRequired();
        }

        // Perform slashing
        _performSlash(reason);
    }

    /// INTERNAL FUNCTIONS

    function _setReportFreshnessThreshold(uint32 _reportFreshnessThreshold) internal {
        reportFreshnessThreshold = _reportFreshnessThreshold;
        emit ReportFreshnessThresholdSet(_reportFreshnessThreshold);
    }

    function _setSlashAmountTokens(uint256 _slashAmountTokens) internal {
        slashAmountTokens = _slashAmountTokens;
        emit SlashAmountTokensSet(_slashAmountTokens);
    }

    /**
     * @notice Performs the slashing operation with calculated proportions
     * @param reason The reason for slashing
     */
    function _performSlash(string memory reason) internal {
        address operator = address(computeOperator);

        // Calculate the proportion to slash
        uint256 wadToSlash = _calculateSlashProportion(operator);

        // Create single-element arrays
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyToSlash;

        uint256[] memory wadsToSlash = new uint256[](1);
        wadsToSlash[0] = wadToSlash;

        // Execute slash
        (uint256 slashId,) = allocationManager.slashOperator(
            address(computeAVSRegistrar),
            IAllocationManagerTypes.SlashingParams({
                operator: operator,
                operatorSetId: operatorSetId,
                strategies: strategies,
                wadsToSlash: wadsToSlash,
                description: reason
            })
        );

        emit OperatorSlashed(operator, operatorSetId, reason, slashId);
    }

    /**
     * @notice Calculates the proportion to slash based on allocated shares
     * @param operator The operator address
     * @return The proportion to slash (in WAD, capped at 1e18)
     */
    function _calculateSlashProportion(address operator) internal view returns (uint256) {
        // TODO: Implement proportion calculation
        // Should convert slashAmountTokens to proportion of allocated stake
        return 0;
    }

    /**
     * @notice Hashes a CloudReport struct using EIP-712
     * @param report The cloud report to hash
     * @return The hash of the cloud report
     */
    function _hashCloudReport(CloudReport calldata report) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(
                CLOUD_REPORT_TYPEHASH,
                keccak256(abi.encodePacked(report.projectId)),
                report.fromTimestamp,
                report.toTimestamp,
                keccak256(report.eigendaCert)
            )
        );
    }

    /// VIEW FUNCTIONS

    /// @inheritdoc ICloudReportCompendium
    function latestReportSubmission() external view returns (CloudReportSubmission memory) {
        return _latestReportSubmission;
    }

    /// @inheritdoc ICloudReportCompendium
    function calculateReportDigestHash(CloudReport calldata report) external view returns (bytes32) {
        return _calculateSignableDigest(_hashCloudReport(report));
    }
}
