// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IReleaseManagerTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IApp} from "./IApp.sol";

interface IAppController {
    /// @notice Billing rates per token and SKU tier
    struct BillingRates {
        uint128 stopped;
        uint128 running;
    }

    /// @notice Thrown when trying to create an app that already exists
    error AppAlreadyExists();

    /// @notice Thrown when trying to access an app that does not exist
    error AppDoesNotExist();

    /// @notice Thrown when an invalid release metadata URI is provided
    error InvalidReleaseMetadataURI();

    /// @notice Thrown when a RMS release has more than one artifact
    error MoreThanOneArtifact();

    /// @notice Thrown when trying to operate on an app in the given status
    error InvalidAppStatus();

    /// @notice Thrown when global maximum vCPU limit has been reached
    error GlobalMaxVCPUsExceeded();

    /// @notice Thrown when trying to use a billing account that doesn't exist
    error BillingAccountDoesNotExist();

    /// @notice Thrown when trying to use a billing account with wrong token
    error BillingTokenMismatch();

    /// @notice Thrown when caller is not the billing account owner
    error NotBillingAccountOwner();

    /// @notice Thrown when SKU rates are not set for a token
    error SkuRatesNotSet();

    /// @notice Thrown when SKU vCPUs are not configured
    error SkuVCPUsNotSet();

    /// @notice Emitted when a new app is successfully created
    event AppCreated(address indexed creator, IApp indexed app, uint32 operatorSetId);

    /// @notice Emitted when an app is upgraded
    event AppUpgraded(IApp indexed app, uint256 rmsReleaseId, Release release);

    /// @notice Emitted when an app is started
    event AppStarted(IApp indexed app);

    /// @notice Emitted when an app is stopped
    event AppStopped(IApp indexed app);

    /// @notice Emitted when an app is terminated
    event AppTerminated(IApp indexed app);

    /// @notice Emitted when an app is terminated by admin
    event AppTerminatedByAdmin(IApp indexed app);

    /// @notice Emitted when the global maximum vCPUs limit is set
    event MaxGlobalVCPUsSet(uint32 limit);

    /// @notice Emitted when SKU rates are set for a token and tier
    event SkuRatesSet(IERC20 indexed token, uint8 indexed skuTier, uint128 runningRate, uint128 stoppedRate);

    /// @notice Emitted when SKU vCPUs are configured
    event SkuVCPUsSet(uint8 indexed skuTier, uint32 vCPUs);

    /// @notice Emitted when the billing recipient is updated
    event BillingRecipientSet(address indexed billingRecipient);

    /// @notice Emitted when a billing account is created for an app
    event BillingAccountCreated(
        IApp indexed app, uint256 indexed billingAccountId, IERC20 indexed token, uint8 skuTier
    );

    /// @notice Emitted when an app is added to an existing billing account
    event AppAddedToBillingAccount(IApp indexed app, uint256 indexed billingAccountId);

    /**
     * @notice Enum for app status
     */
    enum AppStatus {
        NONE, // App has not been created yet
        STARTED, // App is has been started
        STOPPED, // App is has been stopped but can be restarted
        TERMINATED // App is permanently terminated

    }

    /**
     * @notice A struct containing a release and its environment
     * @param rmsRelease The release to publish
     * @param env The environment for the release
     * @param encryptedSecrets The encrypted secrets for the release
     */
    struct Release {
        IReleaseManagerTypes.Release rmsRelease;
        bytes publicEnv;
        bytes encryptedEnv;
    }

    /// @notice The controller's config for an app
    struct AppConfig {
        address creator;
        uint32 operatorSetId;
        uint32 latestReleaseBlockNumber;
        AppStatus status;
        uint8 skuTier;
        uint256 billingAccountId;
        IERC20 billingToken;
        uint128 currentRate;
    }

    /**
     * @notice Initialize the AppController contract
     * @param admin The admin address for the contract
     */
    function initialize(address admin) external;

    /**
     * @notice Sets the global maximum vCPUs limit
     * @param limit The maximum number of vCPUs globally
     * @dev Caller must be UAM permissioned for the AppController
     */
    function setMaxGlobalVCPUs(uint32 limit) external;

    /**
     * @notice Creates a new app instance with billing
     * @param salt The salt to use for the app
     * @param release The release to upgrade to
     * @param skuTier The SKU tier for billing rates
     * @param billingToken The token to use for billing payments
     * @param billingAccountId Billing account ID to use (0 to create new, non-zero to reuse existing)
     * @param initialDeposit Optional amount to deposit when creating/using billing account (0 to skip)
     * @return app The address of the newly created app
     */
    function createApp(
        bytes32 salt,
        Release calldata release,
        uint8 skuTier,
        IERC20 billingToken,
        uint256 billingAccountId,
        uint128 initialDeposit
    ) external returns (IApp app);

    /**
     * @notice Upgrades an app with a new release to the ReleaseManager
     * @param app The app to upgrade with the release
     * @param release The release to upgrade to
     * @return releaseId The unique identifier for the published release
     * @dev Caller must be UAM permissioned for the app
     * @dev The rms release must have exactly one artifact, with the digest being the docker
     * image digest and the registry being the docker registry it is stored at.
     * @dev The env must be a JSON marshalled bytes representing the public environment variables for the app.
     * @dev The encryptedSecrets must be a JSON Web Encryption objects of the encrypted secrets for the app encrypted with
     * the kms's public key.
     * @dev The app must not be AppStatus.TERMINATED
     */
    function upgradeApp(IApp app, Release calldata release) external returns (uint256);

    /**
     * @notice Starts an app, which starts the instance backing it
     * @param app The app to start
     * @dev Caller must be UAM permissioned for the app
     * @dev App must be AppStatus.STOPPED
     */
    function startApp(IApp app) external;

    /**
     * @notice Stops an app, which stops the instance backing it
     * @param app The app to stop
     * @dev Caller must be UAM permissioned for the app
     * @dev App must be AppStatus.STARTED
     */
    function stopApp(IApp app) external;

    /**
     * @notice Terminates an app permanently
     * @param app The app to terminate
     * @dev Caller must be UAM permissioned for the app
     * @dev Once terminated, no further write operations are allowed
     */
    function terminateApp(IApp app) external;

    /**
     * @notice Terminates an app permanently by admin
     * @param app The app to terminate
     * @dev Caller must be UAM permissioned for the AppController
     * @dev Once terminated, no further write operations are allowed
     */
    function terminateAppByAdmin(IApp app) external;

    /**
     * @notice Calculates the app id for a given deployer and salt
     * @param deployer The address of the deployer
     * @param salt The salt to use for the app
     * @return The app id
     */
    function calculateAppId(address deployer, bytes32 salt) external view returns (IApp);

    /**
     * @notice Gets the status for a given app
     * @param app The app to get the status for
     * @return The app status
     */
    function getAppStatus(IApp app) external view returns (AppStatus);

    /**
     * @notice Gets the creator of an app
     * @param app The app to check
     * @return The address of the app creator
     */
    function getAppCreator(IApp app) external view returns (address);

    /**
     * @notice Gets the operator set ID for a given app
     * @param app The app to get the operator set ID for
     * @return The operator set ID
     */
    function getAppOperatorSetId(IApp app) external view returns (uint32);

    /**
     * @notice Gets the latest release block number for a given app
     * @param app The app to get the latest release block number for
     * @return The latest release block number
     */
    function getAppLatestReleaseBlockNumber(IApp app) external view returns (uint32);

    /**
     * @notice Retrieves a paginated list of all apps and their configurations
     * @param offset The starting index for pagination (0-based)
     * @param limit The maximum number of apps to return in this page
     * @return apps An array of app contract instances
     * @return configs An array of corresponding app configurations
     */
    function getApps(uint256 offset, uint256 limit) external view returns (IApp[] memory, AppConfig[] memory);

    /**
     * @notice Retrieves a paginated list of apps where the specified developer is an admin and their configurations
     * @param developer The address of the developer
     * @param offset The starting index for pagination (0-based)
     * @param limit The maximum number of apps to return in this page
     * @return apps An array of app contract instances
     * @return configs An array of corresponding app configurations
     */
    function getAppsByDeveloper(address developer, uint256 offset, uint256 limit)
        external
        view
        returns (IApp[] memory, AppConfig[] memory);

    /**
     * @notice Calculates the digest hash for a given api permission
     * @param permission The permission to calculate the digest hash for
     * @param expiry The expiry of the api permission
     * @return The digest hash
     */
    function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) external view returns (bytes32);

    /**
     * @notice Sets the billing rates for a specific token and SKU tier
     * @param token The billing token
     * @param skuTier The SKU tier
     * @param runningRate The rate per second when app is STARTED
     * @param stoppedRate The rate per second when app is STOPPED
     * @dev Caller must be UAM permissioned for the AppController
     */
    function setSkuRates(IERC20 token, uint8 skuTier, uint128 runningRate, uint128 stoppedRate) external;

    /**
     * @notice Sets the address that receives billing payments
     * @param recipient The billing recipient address
     * @dev Caller must be UAM permissioned for the AppController
     */
    function setBillingRecipient(address recipient) external;

    /**
     * @notice Creates a new billing account and optionally deposits funds
     * @param billingToken The token to use for billing payments
     * @param skuTier The SKU tier for billing rates
     * @param initialDeposit Optional amount to deposit (0 to skip deposit)
     * @return billingAccountId The newly created billing account ID
     */
    function createBillingAccount(IERC20 billingToken, uint8 skuTier, uint128 initialDeposit)
        external
        returns (uint256 billingAccountId);

    /**
     * @notice Deposits funds into a billing account
     * @param billingAccountId The billing account ID
     * @param amount The amount to deposit
     */
    function deposit(uint256 billingAccountId, uint128 amount) external;

    /**
     * @notice Refunds excess funds from a billing account
     * @param billingAccountId The billing account ID
     * @param amount The amount to refund
     */
    function refund(uint256 billingAccountId, uint128 amount) external;

    /**
     * @notice Gets billing information for an app
     * @param app The app to query
     * @return billingAccountId The billing account ID
     * @return balance The current balance in the account
     * @return ratePerSecond The current rate per second
     * @return depletionTime When the account will run out of funds
     * @return skuTier The SKU tier
     * @return token The billing token
     */
    function getBillingInfo(IApp app)
        external
        view
        returns (
            uint256 billingAccountId,
            uint128 balance,
            uint128 ratePerSecond,
            uint256 depletionTime,
            uint8 skuTier,
            IERC20 token
        );

    /**
     * @notice Gets all apps using a specific billing account
     * @param billingAccountId The billing account ID
     * @return apps Array of apps sharing this billing account
     */
    function getAppsOnBillingAccount(uint256 billingAccountId) external view returns (IApp[] memory apps);

    /**
     * @notice Sets the vCPU allocation for a SKU tier
     * @param skuTier The SKU tier
     * @param vCPUs The number of vCPUs for this tier
     * @dev Caller must be UAM permissioned for the AppController
     */
    function setSkuVCPUs(uint8 skuTier, uint32 vCPUs) external;
}
