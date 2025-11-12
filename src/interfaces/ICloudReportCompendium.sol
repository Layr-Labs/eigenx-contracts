// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {
    IEigenDACertVerifierRouter
} from "@eigenda-contracts/src/integrations/cert/interfaces/IEigenDACertVerifierRouter.sol";
import {IStrategy} from "@eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

interface ICloudReportCompendium {
    /// ERRORS
    /// @notice Thrown when the report's project ID does not match the reference project ID
    error InvalidProjectId();

    /// @notice Thrown when the report is too old (exceeds freshness threshold)
    error ReportTooStale();

    /// @notice Thrown when the EigenDA certificate is invalid or hash does not match
    error InvalidCertificate();

    /// @notice Thrown when checks pass and no slash is required
    error NoSlashRequired();

    /// @notice Thrown when there is insufficient allocation to slash
    error InsufficientAllocation();

    /// @notice Thrown when report's fromTimestamp doesn't match previous report's toTimestamp
    error ReportChainTimestampMismatch();

    /// @notice Thrown when report timestamps are invalid (fromTimestamp >= toTimestamp)
    error InvalidReportTimestamps();

    /// @notice Thrown when report toTimestamp is in the future
    error ReportFromFuture();

    /// EVENTS

    /// @notice Emitted when a cloud report is submitted
    event ReportSubmitted(CloudReport report, uint256 timestamp);

    /// @notice Emitted when an operator is slashed
    event OperatorSlashed(address indexed operator, uint32 indexed operatorSetId, string reason, uint256 slashId);

    /// @notice Emitted when the report freshness threshold is updated
    event ReportFreshnessThresholdSet(uint32 reportFreshnessThreshold);

    /// @notice Emitted when the slash amount in tokens is updated
    event SlashAmountTokensSet(uint256 slashAmountTokens);

    /// STRUCTS

    struct CloudReportSubmission {
        uint32 submissionTimestamp;
        uint32 fromTimestamp;
        uint32 toTimestamp;
        bytes20 eigendaCertHash; // using bytes20 for packing
    }

    struct CloudReport {
        string projectId;
        uint32 fromTimestamp;
        uint32 toTimestamp;
        bytes eigendaCert;
    }

    /// VIEW FUNCTIONS

    /// @notice The EIP-712 typehash for the CloudReport struct.
    function CLOUD_REPORT_TYPEHASH() external view returns (bytes32);

    /// @notice The address of the EigenDA Cert Verifier Router contract.
    function certVerifierRouter() external view returns (IEigenDACertVerifierRouter);

    /// @notice The address that must sign cloud reports.
    function reportAttester() external view returns (address);

    /// @notice The operator set ID this compendium is tracking.
    function operatorSetId() external view returns (uint32);

    /// @notice Hash of the reference project ID that reports must match.
    function referenceProjectIdHash() external view returns (bytes32);

    /// @notice The strategy to slash when an operator fails checks.
    function strategyToSlash() external view returns (IStrategy);

    /// @notice The number of seconds within which a report must be submitted.
    function reportFreshnessThreshold() external view returns (uint32);

    /// @notice The absolute amount in tokens to slash.
    function slashAmountTokens() external view returns (uint256);

    /// @notice The latest cloud report submission.
    function latestReportSubmission() external view returns (CloudReportSubmission memory);

    /**
     * @notice Calculates the digest hash for a cloud report.
     * @param report The cloud report.
     * @return The digest hash.
     */
    function calculateReportDigestHash(CloudReport calldata report) external view returns (bytes32);

    /**
     * @notice Initializes the CloudReportCompendium contract
     * @param _reportFreshnessThreshold The number of blocks within which a report must be submitted
     * @param _slashAmountTokens The absolute amount in tokens to slash
     */
    function initialize(uint32 _reportFreshnessThreshold, uint256 _slashAmountTokens) external;

    /**
     * @notice Sets the report freshness threshold.
     * @param _reportFreshnessThreshold The new threshold in seconds.
     * @dev Caller must be UAM permissioned
     */
    function setReportFreshnessThreshold(uint32 _reportFreshnessThreshold) external;

    /**
     * @notice Sets the slash amount in tokens.
     * @param _slashAmountTokens The new slash amount.
     * @dev Caller must be UAM permissioned
     */
    function setSlashAmountTokens(uint256 _slashAmountTokens) external;

    /**
     * @notice Submits a cloud report.
     * @param report The cloud report to submit.
     * @param signature The signature of the report by the report attester.
     * @dev Caller must be UAM permissioned
     */
    function submitReport(CloudReport calldata report, bytes calldata signature) external;

    /**
     * @notice Slashes the operator for the operatorSet by running the requisite checks and seeing if they fail.
     * The checks are:
     *    1. Check that the latest report is fresh enough (within REPORT_FRESHNESS_THRESHOLD of the current timestamp)
     *    2. Check that the EigenDA cert is valid
     * If either of these checks fail, the operator is slashed for SLASH_AMOUNT of their stake.
     * @param eigendaCert The EigenDA certificate to verify (preimage of stored hash).
     */
    function slash(bytes calldata eigendaCert) external;
}
