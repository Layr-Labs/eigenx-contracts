// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IReleaseManagerTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {IApp} from "./IApp.sol";

interface IAppController {
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

    /// @notice Thrown when an address has exceeded its maximum active app limit
    error MaxActiveAppsExceeded();

    /// @notice Thrown when global maximum active app limit has been reached
    error GlobalMaxActiveAppsExceeded();

    /// @notice Thrown when trying to suspend an account that still has active apps
    error AccountHasActiveApps();

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

    /// @notice Emitted when an app is suspended
    event AppSuspended(IApp indexed app);

    /// @notice Emitted when the maximum active apps limit is set for an address
    event MaxActiveAppsSet(address indexed user, uint32 limit);

    /// @notice Emitted when the global maximum active apps limit is set
    event GlobalMaxActiveAppsSet(uint32 limit);

    /// @notice Emitted when an app's metadata URI is updated
    event AppMetadataURIUpdated(IApp indexed app, string metadataURI);

    /**
     * @notice Enum for app status
     */
    enum AppStatus {
        NONE, // App has not been created yet
        STARTED, // App has been started
        STOPPED, // App has been stopped but can be restarted
        TERMINATED, // App is permanently terminated
        SUSPENDED // App is suspended and can be started again, but does not have reserved capacity
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
    }

    /// @notice User configuration and state
    struct UserConfig {
        uint32 maxActiveApps;
        uint32 activeAppCount;
    }

    /**
     * @notice Initialize the AppController contract
     * @param admin The admin address for the contract
     */
    function initialize(address admin) external;

    /**
     * @notice Sets the maximum active apps limit for a specific address
     * @param user The address to set the limit for
     * @param limit The maximum number of active apps the address can have
     * @dev Caller must be UAM permissioned for the AppController
     */
    function setMaxActiveAppsPerUser(address user, uint32 limit) external;

    /**
     * @notice Sets the global maximum active apps limit
     * @param limit The maximum number of active apps globally
     * @dev Caller must be UAM permissioned for the AppController
     */
    function setMaxGlobalActiveApps(uint32 limit) external;

    /**
     * @notice Creates a new app instance
     * @param salt The salt to use for the app
     * @param release The release to upgrade to
     * @return app The address of the newly created app
     */
    function createApp(bytes32 salt, Release calldata release) external returns (IApp app);

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
     * @notice Updates the metadata URI for an app
     * @param app The app to update the metadata URI for
     * @param metadataURI The new metadata URI
     * @dev Caller must be UAM permissioned for the app
     */
    function updateAppMetadataURI(IApp app, string calldata metadataURI) external;

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
     * @notice Suspends all active apps for an account and sets their max active apps to 0
     * @param account The account to suspend
     * @param apps The apps to suspend (must all be created by account and include ALL active apps)
     * @dev Caller must be the account owner or UAM permissioned for the AppController
     * @dev All active apps (STARTED or STOPPED) must be provided, otherwise reverts with AccountHasActiveApps
     * @dev Apps already suspended or terminated are silently skipped
     */
    function suspend(address account, IApp[] calldata apps) external;

    /**
     * @notice Gets the maximum global active apps limit
     * @return The maximum number of active apps globally
     */
    function maxGlobalActiveApps() external view returns (uint32);

    /**
     * @notice Gets the current global active app count
     * @return The current number of active apps globally
     */
    function globalActiveAppCount() external view returns (uint32);

    /**
     * @notice Gets the maximum active apps for a user
     * @param user The address to check
     * @return The maximum number of active apps for the user
     */
    function getMaxActiveAppsPerUser(address user) external view returns (uint32);

    /**
     * @notice Gets the current active app count for a user
     * @param user The address to check
     * @return The current number of active apps for the user
     */
    function getActiveAppCount(address user) external view returns (uint32);

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
     * @notice Retrieves a paginated list of apps created by the specified address and their configurations
     * @param creator The address of the creator
     * @param offset The starting index for pagination (0-based)
     * @param limit The maximum number of apps to return in this page
     * @return apps An array of app contract instances
     * @return configs An array of corresponding app configurations
     */
    function getAppsByCreator(address creator, uint256 offset, uint256 limit)
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
}
