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

    /// @notice Thrown when a caller lacks the required per-app team role.
    error InvalidTeamRole();

    /// @notice Thrown when revoking/renouncing ADMIN would leave the team with zero admins.
    error CannotRevokeLastAdmin();

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

    /// @notice Emitted when app ownership is transferred to a new address
    event AppOwnershipTransferred(IApp indexed app, address indexed previousOwner, address indexed newOwner);

    /// @notice Emitted when a team role is granted on an app.
    event TeamRoleGranted(IApp indexed app, TeamRole indexed role, address indexed account);

    /// @notice Emitted when a team role is revoked on an app (by an admin or via renounce).
    event TeamRoleRevoked(IApp indexed app, TeamRole indexed role, address indexed account);

    /// @notice Enum for app status
    enum AppStatus {
        NONE, // App has not been created yet
        STARTED, // App has been started
        STOPPED, // App has been stopped but can be restarted
        TERMINATED, // App is permanently terminated
        SUSPENDED // App is suspended and can be started again, but does not have reserved capacity
    }

    /// @notice Billing type for an app
    enum BillingType {
        DEFAULT, // Billed to the creator's account
        ISOLATED // Billed to the app's own address
    }

    /// @notice Per-app team roles.
    /// @dev ADMIN is the authoritative role for CRITICAL app-level ops
    ///      (upgradeApp / transferOwnership / terminateApp) and for managing
    ///      team membership. PAUSER and DEVELOPER are OPERATIONAL:
    ///
    ///        - PAUSER    → may call stopApp
    ///        - DEVELOPER → may call updateAppMetadataURI
    ///        - ADMIN     → may call everything operational, plus all
    ///                      critical ops, plus grant/revoke team roles,
    ///                      plus transferOwnership.
    ///
    ///      Roles live in `AppController` storage, NOT in
    ///      PermissionController. That's intentional — the Timelock
    ///      guarantees only bind if the admin set the gates trust is a
    ///      set this contract controls. PermissionController admins are
    ///      still used for platform-admin-level functions on this
    ///      contract itself (setMaxActiveAppsPerUser, setMaxGlobalActiveApps,
    ///      migrateAdmins, terminateAppByAdmin, suspend).
    enum TeamRole {
        ADMIN,
        PAUSER,
        DEVELOPER
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

    /// @notice The controller's config for an app (public-facing, ABI-stable)
    struct AppConfig {
        address creator;
        uint32 operatorSetId;
        uint32 latestReleaseBlockNumber;
        AppStatus status;
    }

    /// @notice Internal storage config for an app, extends AppConfig with additional fields
    struct AppConfigStorage {
        // Slot layout (all 32 bytes packed):
        //   bytes 0-19:  address creator            (20 bytes)
        //   bytes 20-23: uint32  operatorSetId      ( 4 bytes)
        //   bytes 24-27: uint32  latestReleaseBlockNumber ( 4 bytes)
        //   byte  28:    AppStatus status           ( 1 byte)
        //   byte  29:    BillingType billingType    ( 1 byte)   ← present on v1.4.0 chain state
        //   byte  30:    bool timelocked            ( 1 byte)   ← new in v1.5.0; safely zero on all v1.4.0 apps
        //   byte  31:    (unused)
        address creator;
        uint32 operatorSetId;
        uint32 latestReleaseBlockNumber;
        AppStatus status;
        BillingType billingType;
        // true = owner is a factory Timelock; sensitive ops must go through
        // Timelock.schedule → execute. Must NOT be placed at byte 29 — that
        // byte already holds `billingType` on existing deployed contracts.
        bool timelocked;
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
     * @notice Creates a new app with isolated billing, where costs are billed to the app's own address
     * @param salt The salt to use for the app
     * @param release The release to upgrade to
     * @return app The address of the newly created app
     * @dev The app address is pre-computed and must have quota set via setMaxActiveAppsPerUser before calling
     */
    function createAppWithIsolatedBilling(bytes32 salt, Release calldata release) external returns (IApp app);

    /**
     * @notice Upgrades an app with a new release to the ReleaseManager. Critical op.
     * @param app The app to upgrade with the release
     * @param release The release to upgrade to
     * @return releaseId The unique identifier for the published release
     * @dev Caller must hold ADMIN on the app's team.
     * @dev When the app is timelocked, caller must additionally be the creator
     *      (the Timelock itself) — in practice, this forces the call to come
     *      from a scheduled → executed Timelock operation.
     * @dev The rms release must have exactly one artifact, with the digest being the docker
     * image digest and the registry being the docker registry it is stored at.
     * @dev The env must be a JSON marshalled bytes representing the public environment variables for the app.
     * @dev The encryptedSecrets must be a JSON Web Encryption objects of the encrypted secrets for the app encrypted with
     * the kms's public key.
     * @dev The app must not be AppStatus.TERMINATED
     */
    function upgradeApp(IApp app, Release calldata release) external returns (uint256);

    /**
     * @notice Transfers app ownership to a new address. Critical op.
     * @param app The app to transfer ownership of
     * @param newOwner The new owner address
     * @dev Caller must hold ADMIN on the app's team.
     * @dev When `newOwner` is a factory-deployed Timelock the app's `timelocked`
     *      flag is flipped to true; otherwise it's cleared.
     * @dev When the app is already timelocked, only the Timelock itself (via
     *      schedule → execute) may transfer. The new owner is automatically
     *      granted ADMIN on the new team so they can govern going forward.
     */
    function transferOwnership(IApp app, address newOwner) external;

    /**
     * @notice Updates the metadata URI for an app. Operational.
     * @param app The app to update the metadata URI for
     * @param metadataURI The new metadata URI
     * @dev Permitted to ADMIN or DEVELOPER.
     */
    function updateAppMetadataURI(IApp app, string calldata metadataURI) external;

    /**
     * @notice Starts an app, which starts the instance backing it.
     * @param app The app to start
     * @dev Permitted to ADMIN only — starting commits capacity; treated as
     *      privileged. App must be STOPPED.
     */
    function startApp(IApp app) external;

    /**
     * @notice Stops an app, which stops the instance backing it. Operational.
     * @param app The app to stop
     * @dev Permitted to ADMIN or PAUSER.
     * @dev App must be AppStatus.STARTED.
     */
    function stopApp(IApp app) external;

    /**
     * @notice Terminates an app permanently. Critical op.
     * @param app The app to terminate
     * @dev Caller must hold ADMIN on the app's team.
     * @dev When the app is timelocked, caller must additionally be the
     *      creator (Timelock).
     * @dev Once terminated, no further write operations are allowed.
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
     * @notice Grant a team role to `account` on `app`.
     * @dev Caller must be ADMIN on the app.
     * @dev When the team is timelocked AND `role == ADMIN`, the caller must
     *      additionally be the creator (the Timelock itself) — which in
     *      practice means going through schedule → execute. Grants of PAUSER
     *      / DEVELOPER on a timelocked app do NOT require going through the
     *      Timelock; those are operational powers the admin can revoke at
     *      any time. Fixes audit finding A-2.
     */
    function grantTeamRole(IApp app, TeamRole role, address account) external;

    /**
     * @notice Revoke a team role from `account` on `app`.
     * @dev Caller must be ADMIN on the app.
     * @dev When the team is timelocked AND `role == ADMIN`, the caller must
     *      additionally be the creator. Revoking any ADMIN below the
     *      last-admin floor reverts. Operational role revocations on a
     *      timelocked app are NOT timelock-gated. Fixes audit finding A-3.
     */
    function revokeTeamRole(IApp app, TeamRole role, address account) external;

    /**
     * @notice Renounce your own team role on `app`.
     * @dev Renouncing ADMIN below the last-admin floor reverts.
     */
    function renounceTeamRole(IApp app, TeamRole role) external;

    /**
     * @notice Returns true iff `account` holds `role` on `app`.
     */
    function hasTeamRole(IApp app, TeamRole role, address account) external view returns (bool);

    /**
     * @notice Migrate pre-v1.5.0 apps from PermissionController-based auth to
     *         the new per-app team RBAC. For each app in `apps`, every admin
     *         currently registered in PermissionController is granted the
     *         ADMIN role in AppController's own storage. No-op for apps that
     *         already have an ADMIN.
     * @dev Caller must be UAM permissioned for the AppController itself
     *      (platform admin). Intended to be called once per app after the
     *      v1.5.0 upgrade; safe to call again (idempotent per-(app, admin)).
     */
    function migrateAdmins(IApp[] calldata apps) external;

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
     * @notice Gets the billing account for an app
     * @param app The app to check
     * @return The billing account address (app address if ISOLATED, creator if DEFAULT)
     */
    function getBillingAccount(IApp app) external view returns (address);

    /**
     * @notice Gets the billing type for an app
     * @param app The app to check
     * @return The billing type (DEFAULT or ISOLATED)
     */
    function getBillingType(IApp app) external view returns (BillingType);

    /**
     * @notice Returns whether the app's creator is a factory Timelock.
     * @param app The app to check
     * @return True iff sensitive ops (upgrade/terminate) must go through
     *         Timelock.schedule → execute — i.e. direct calls by any non-owner
     *         are rejected regardless of PermissionController grants.
     */
    function getAppTimelocked(IApp app) external view returns (bool);

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
     * @notice Retrieves a paginated list of apps billed to the specified account and their configurations
     * @param account The billing account address (creator for regular apps, app address for isolated billing apps)
     * @param offset The starting index for pagination (0-based)
     * @param limit The maximum number of apps to return in this page
     * @return apps An array of app contract instances
     * @return configs An array of corresponding app configurations
     */
    function getAppsByBillingAccount(address account, uint256 offset, uint256 limit)
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
