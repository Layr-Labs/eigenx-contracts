// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {SignatureUtilsMixin} from "@eigenlayer-contracts/src/contracts/mixins/SignatureUtilsMixin.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {PermissionControllerMixin} from "@eigenlayer-contracts/src/contracts/mixins/PermissionControllerMixin.sol";
import {
    IReleaseManager,
    IReleaseManagerTypes
} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IComputeAVSRegistrar} from "./interfaces/IComputeAVSRegistrar.sol";
import {IComputeOperator} from "./interfaces/IComputeOperator.sol";
import {AppControllerStorage} from "./storage/AppControllerStorage.sol";
import {IAppController} from "./interfaces/IAppController.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import {IBeacon} from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import {IApp} from "./interfaces/IApp.sol";
import {ISafeTimelockFactory} from "./interfaces/ISafeTimelockFactory.sol";

contract AppController is Initializable, SignatureUtilsMixin, PermissionControllerMixin, AppControllerStorage {
    using EnumerableSet for EnumerableSet.AddressSet;

    /// MODIFIERS

    /// @notice Modifier to ensure app exists
    modifier appExists(IApp app) {
        require(_exists(_appConfigs[app].status), InvalidAppStatus());
        _;
    }

    /// @notice Modifier to ensure app is in an active status
    modifier appIsActive(IApp app) {
        require(_isActive(_appConfigs[app].status), InvalidAppStatus());
        _;
    }

    /**
     * @param _version The version string to use for this contract's domain separator
     * @param _permissionController The PermissionController contract address
     * @param _computeAVSRegistrar The ComputeAVSRegistrar contract address
     * @param _computeOperator The ComputeOperator contract address
     * @param _appBeacon The beacon for creating App proxies
     */
    constructor(
        string memory _version,
        IPermissionController _permissionController,
        IReleaseManager _releaseManager,
        IComputeAVSRegistrar _computeAVSRegistrar,
        IComputeOperator _computeOperator,
        IBeacon _appBeacon,
        ISafeTimelockFactory _safeTimelockFactory
    )
        SignatureUtilsMixin(_version)
        PermissionControllerMixin(_permissionController)
        AppControllerStorage(_releaseManager, _computeOperator, _computeAVSRegistrar, _appBeacon, _safeTimelockFactory)
    {
        _disableInitializers();
    }

    /// @inheritdoc IAppController
    function initialize(address admin) external initializer {
        // Accept the ComputeAVSRegistrar as an admin
        permissionController.acceptAdmin(address(computeAVSRegistrar));

        // Add this contract itself as an admin
        permissionController.addPendingAdmin(address(this), address(this));
        permissionController.acceptAdmin(address(this));

        // Add the app controller as an admin
        permissionController.addPendingAdmin(address(this), admin);

        // Assign an initial operator set ID
        uint32 operatorSetId = computeAVSRegistrar.assignOperatorSetId();

        // Create the operator set
        computeAVSRegistrar.createOperatorSet(operatorSetId);

        // Add the compute operator to the allowlist
        computeAVSRegistrar.addOperatorToAllowlist(
            OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId}), address(computeOperator)
        );

        // Register the compute operator for the operator set
        computeOperator.registerForOperatorSet(operatorSetId);
    }

    /// @inheritdoc IAppController
    function setMaxActiveAppsPerUser(address user, uint32 limit) external checkCanCall(address(this)) {
        _setMaxActiveAppsPerUser(user, limit);
    }

    /// @inheritdoc IAppController
    function setMaxGlobalActiveApps(uint32 limit) external checkCanCall(address(this)) {
        maxGlobalActiveApps = limit;
        emit GlobalMaxActiveAppsSet(limit);
    }

    /// @inheritdoc IAppController
    function createApp(bytes32 salt, Release calldata release) external returns (IApp app) {
        _checkAndIncrementActiveApps(msg.sender);
        app = _deployApp(salt, release, BillingType.DEFAULT);
    }

    /// @inheritdoc IAppController
    function createAppWithIsolatedBilling(bytes32 salt, Release calldata release) external returns (IApp app) {
        _checkAndIncrementActiveApps(address(calculateAppId(msg.sender, salt)));
        app = _deployApp(salt, release, BillingType.ISOLATED);
    }

    /// @inheritdoc IAppController
    function upgradeApp(IApp app, Release calldata release)
        external
        checkCanCall(address(app))
        appIsActive(app)
        returns (uint256)
    {
        // When the app's owner is a factory-deployed Timelock, the sensitive
        // path must go through Timelock.schedule → execute so the delay window
        // actually applies. Any other PermissionController-permitted caller
        // (e.g. a co-admin granted via UAM) would bypass the queue otherwise.
        if (_appConfigs[app].timelocked) {
            require(msg.sender == _appConfigs[app].creator, InvalidPermissions());
        }
        return _upgradeApp(app, release);
    }

    /// @inheritdoc IAppController
    function transferOwnership(IApp app, address newOwner) external checkCanCall(address(app)) appExists(app) {
        require(newOwner != address(0), InvalidPermissions());

        AppConfigStorage storage config = _appConfigs[app];

        // When already timelocked, only the Timelock owner itself may move the
        // app. Otherwise any admin the Timelock had granted via UAM could
        // transfer out of its governance entirely, bypassing the queue delay.
        if (config.timelocked) {
            require(msg.sender == config.creator, InvalidPermissions());
        }

        address previousOwner = config.creator;
        config.creator = newOwner;

        // Flip the flag based on the new owner. Non-factory addresses (EOAs,
        // externally-deployed Safes, arbitrary contracts) clear it — they have
        // no schedule→execute semantics we can trust. Factory-deployed
        // Timelocks enable it.
        config.timelocked = address(safeTimelockFactory) != address(0) && safeTimelockFactory.isTimelock(newOwner);

        // ISOLATED billing apps bill the app address, not the creator, so
        // ownership transfer has no effect on billing accounting. DEFAULT
        // billing apps bill the creator, so we need to move the active-app
        // counter from the previous creator to the new one; otherwise a future
        // terminate/suspend would underflow the new owner's counter.
        if (config.billingType == BillingType.DEFAULT && _isActive(config.status)) {
            _userConfigs[previousOwner].activeAppCount--;
            _userConfigs[newOwner].activeAppCount++;
        }

        emit AppOwnershipTransferred(app, previousOwner, newOwner);
    }

    /// @inheritdoc IAppController
    function updateAppMetadataURI(IApp app, string calldata metadataURI)
        external
        checkCanCall(address(app))
        appExists(app)
    {
        emit AppMetadataURIUpdated(app, metadataURI);
    }

    /// @inheritdoc IAppController
    function startApp(IApp app) external checkCanCall(address(app)) appExists(app) {
        _startApp(app);
    }

    /// @inheritdoc IAppController
    function stopApp(IApp app) external checkCanCall(address(app)) {
        AppConfigStorage storage config = _appConfigs[app];
        require(config.status == AppStatus.STARTED, InvalidAppStatus());
        config.status = AppStatus.STOPPED;

        emit AppStopped(app);
    }

    /// @inheritdoc IAppController
    function terminateApp(IApp app) external checkCanCall(address(app)) appIsActive(app) {
        // When timelocked, only the Timelock owner itself (acting via
        // schedule → execute) may terminate. Termination is irreversible;
        // bypassing the queue here defeats the entire purpose of
        // transferring ownership to a Timelock.
        if (_appConfigs[app].timelocked) {
            require(msg.sender == _appConfigs[app].creator, InvalidPermissions());
        }
        _terminateApp(app);
    }

    /// @inheritdoc IAppController
    function terminateAppByAdmin(IApp app) external checkCanCall(address(this)) appIsActive(app) {
        // Protocol admin may not unilaterally terminate a Timelock-owned app;
        // doing so would bypass the delay the user specifically opted into.
        // If intervention is required, protocol admin should schedule the
        // termination through the Timelock itself.
        require(!_appConfigs[app].timelocked, InvalidPermissions());
        _terminateApp(app);
        emit AppTerminatedByAdmin(app);
    }

    /// @inheritdoc IAppController
    function suspend(address account, IApp[] calldata apps) external {
        // Allow account owner to self-suspend or AppController admin to enforce suspension
        require(msg.sender == account || _checkCanCall(address(this)), InvalidPermissions());

        // Suspend all provided apps, skipping apps that are already SUSPENDED, TERMINATED, or NONE
        for (uint256 i = 0; i < apps.length; i++) {
            IApp app = apps[i];
            AppConfigStorage memory config = _appConfigs[app];

            // Validate that the app is billed to this account
            address billingAccount = config.billingType == BillingType.ISOLATED ? address(app) : config.creator;
            require(billingAccount == account, InvalidAppStatus());

            // Only suspend if app is active (STARTED or STOPPED)
            if (_isActive(config.status)) {
                _suspendApp(app);
            }
        }

        // Verify all active apps were provided - prevents partial suspension
        require(_userConfigs[account].activeAppCount == 0, AccountHasActiveApps());

        // Zero-out the account's max active apps
        _setMaxActiveAppsPerUser(account, 0);
    }

    /// INTERNAL FUNCTIONS

    /**
     * @notice Deploys an app, configures it, publishes the initial release, and starts it
     * @param salt The salt for deterministic deployment
     * @param release The initial release to upgrade to
     * @param _billingType The billing type for the app (DEFAULT or ISOLATED)
     * @return app The deployed app instance
     */
    function _deployApp(bytes32 salt, Release calldata release, BillingType _billingType) internal returns (IApp app) {
        uint32 operatorSetId = computeAVSRegistrar.assignOperatorSetId();
        releaseManager.publishMetadataURI(
            OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId}), "https://eigencloud.xyz"
        );

        app = IApp(Create2.deploy(0, _calculateAppMixedSalt(msg.sender, salt), _calculateAppInitCode(msg.sender)));
        _appConfigs[app].operatorSetId = operatorSetId;
        _appConfigs[app].latestReleaseBlockNumber = 0;
        _appConfigs[app].creator = msg.sender;
        _appConfigs[app].billingType = _billingType;
        // If the creator is a factory-deployed Timelock, mark the app
        // timelocked at creation so sensitive-op gates fire immediately.
        // Not doing this here leaves a window where any co-admin could run
        // upgradeApp / terminateApp before ownership is "transferred" — and
        // in fact createApp never involves transferOwnership at all.
        // safeTimelockFactory may be address(0) on historical deployments
        // (pre-v1.5.0); in that case isTimelock is guaranteed to return false
        // via the interface contract, but we also defensively short-circuit.
        if (address(safeTimelockFactory) != address(0)) {
            _appConfigs[app].timelocked = safeTimelockFactory.isTimelock(msg.sender);
        }
        _allApps.add(address(app));

        emit AppCreated(msg.sender, app, operatorSetId);

        _upgradeApp(app, release);
        _startApp(app);
    }

    /**
     * @notice Checks if an app status is not NONE
     * @param status The app status to check
     * @return True if status is not NONE, false otherwise
     */
    function _exists(AppStatus status) internal pure returns (bool) {
        return status != AppStatus.NONE;
    }

    /**
     * @notice Checks if an app status is active
     * @param status The app status to check
     * @return True if status is STARTED or STOPPED, false otherwise
     */
    function _isActive(AppStatus status) internal pure returns (bool) {
        return status == AppStatus.STARTED || status == AppStatus.STOPPED;
    }

    /**
     * @notice Checks active app limits and increments counters for a user
     * @param user The user address to check and increment for
     */
    function _checkAndIncrementActiveApps(address user) internal {
        UserConfig storage userConfig = _userConfigs[user];

        // Check global active app limit
        require(globalActiveAppCount < maxGlobalActiveApps, GlobalMaxActiveAppsExceeded());
        // Check user active app limit
        require(userConfig.activeAppCount < userConfig.maxActiveApps, MaxActiveAppsExceeded());

        // Increment active app counts
        globalActiveAppCount++;
        userConfig.activeAppCount++;
    }

    /**
     * @notice Decrements global and billing account active app counters
     * @param app The app instance to decrement counters for
     */
    function _decrementActiveApps(IApp app) internal {
        // Decrement active app counts to free up capacity
        globalActiveAppCount--;

        // Decrement the billing account's active app count
        _userConfigs[getBillingAccount(app)].activeAppCount--;
    }

    /**
     * @notice Sets the maximum number of active apps allowed for a user
     * @param user The user address to set the limit for
     * @param limit The maximum number of active apps allowed
     */
    function _setMaxActiveAppsPerUser(address user, uint32 limit) internal {
        _userConfigs[user].maxActiveApps = limit;
        emit MaxActiveAppsSet(user, limit);
    }

    /**
     * @notice Upgrades an app to a new release by publishing it through the release manager
     * @param app The app instance to upgrade
     * @param release The new release data containing artifacts and metadata
     * @return releaseId The unique identifier assigned to the published release by the release manager
     */
    function _upgradeApp(IApp app, Release calldata release) internal returns (uint256 releaseId) {
        // Check that the release has exactly one artifact
        require(release.rmsRelease.artifacts.length == 1, MoreThanOneArtifact());
        releaseId = releaseManager.publishRelease(
            OperatorSet({avs: address(computeAVSRegistrar), id: _appConfigs[app].operatorSetId}), release.rmsRelease
        );
        _appConfigs[app].latestReleaseBlockNumber = uint32(block.number);

        emit AppUpgraded(app, releaseId, release);
    }

    /**
     * @notice Starts an app and marks it as started
     * @param app The app instance to start
     */
    function _startApp(IApp app) internal {
        AppConfigStorage storage config = _appConfigs[app];
        require(config.status != AppStatus.TERMINATED, InvalidAppStatus());

        // If resuming from suspended, re-check limits and increment active app counters
        if (config.status == AppStatus.SUSPENDED) {
            _checkAndIncrementActiveApps(getBillingAccount(app));
        }
        config.status = AppStatus.STARTED;
        emit AppStarted(app);
    }

    /**
     * @notice Terminates an app and decrements active app counters
     * @param app The app instance to terminate
     */
    function _terminateApp(IApp app) internal {
        _appConfigs[app].status = AppStatus.TERMINATED;
        _decrementActiveApps(app);
        emit AppTerminated(app);
    }

    /**
     * @notice Suspends an app and decrements active app counters
     * @param app The app instance to suspend
     */
    function _suspendApp(IApp app) internal {
        _appConfigs[app].status = AppStatus.SUSPENDED;
        _decrementActiveApps(app);
        emit AppSuspended(app);
    }

    /**
     * @notice Calculates a mixed salt for app deployment using deployer address and provided salt
     * @param deployer The address of the app deployer
     * @param salt The salt value
     * @return The keccak256 hash of deployer and salt, used for deterministic app address generation
     */
    function _calculateAppMixedSalt(address deployer, bytes32 salt) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(deployer, salt));
    }

    /**
     * @notice Generates the initialization code for deploying a new app using beacon proxy pattern
     * @param deployer The address that will be set as the app owner during initialization
     * @return The complete bytecode for deploying a BeaconProxy that initializes the app with the deployer
     */
    function _calculateAppInitCode(address deployer) internal view returns (bytes memory) {
        return abi.encodePacked(
            _BEACON_PROXY_BYTECODE, abi.encode(appBeacon, abi.encodeWithSelector(IApp.initialize.selector, deployer))
        );
    }

    /**
     * @notice Converts an AppConfigStorage to a public AppConfig
     */
    function _toAppConfig(AppConfigStorage storage s) private view returns (AppConfig memory) {
        return AppConfig({
            creator: s.creator,
            operatorSetId: s.operatorSetId,
            latestReleaseBlockNumber: s.latestReleaseBlockNumber,
            status: s.status
        });
    }

    /**
     * @notice Gets filtered apps based on a predicate function
     * @param predicate The predicate function to apply to each app
     * @param target The target address to filter by
     * @param offset The offset to start from
     * @param limit The maximum number of apps to return
     * @return apps The filtered apps
     * @return appConfigsMem The app configs for the filtered apps
     */
    function _getFilteredApps(
        function(IApp, address) view returns (bool) predicate,
        address target,
        uint256 offset,
        uint256 limit
    ) private view returns (IApp[] memory apps, AppConfig[] memory appConfigsMem) {
        uint256 totalApps = _allApps.length();

        apps = new IApp[](limit);
        appConfigsMem = new AppConfig[](limit);
        uint256 skipped = 0;
        uint256 found = 0;

        for (uint256 i = 0; i < totalApps && found < limit; i++) {
            IApp app = IApp(_allApps.at(i));
            if (predicate(app, target)) {
                if (skipped < offset) {
                    skipped++;
                    continue;
                }

                apps[found] = app;
                appConfigsMem[found] = _toAppConfig(_appConfigs[app]);
                found++;
            }
        }

        // Resize arrays to actual number found
        assembly {
            mstore(apps, found)
            mstore(appConfigsMem, found)
        }
    }

    /**
     * @notice Check if address is developer of app
     * @param app The app to check
     * @param developer The developer to check
     * @return True if the developer is the developer of the app
     */
    function _isDeveloper(IApp app, address developer) private view returns (bool) {
        return permissionController.isAdmin(address(app), developer);
    }

    /**
     * @notice Check if address is creator of app
     * @param app The app to check
     * @param creator The creator to check
     * @return True if the creator is the creator of the app
     */
    function _isCreator(IApp app, address creator) private view returns (bool) {
        return _appConfigs[app].creator == creator;
    }

    /**
     * @notice Check if an app is billed to the given account
     * @param app The app to check
     * @param account The billing account to match against
     * @return True if the app is billed to the account
     */
    function _isBilledTo(IApp app, address account) private view returns (bool) {
        return getBillingAccount(app) == account;
    }

    /// VIEW FUNCTIONS

    /// @inheritdoc IAppController
    function getMaxActiveAppsPerUser(address user) external view returns (uint32) {
        return _userConfigs[user].maxActiveApps;
    }

    /// @inheritdoc IAppController
    function getActiveAppCount(address user) external view returns (uint32) {
        return _userConfigs[user].activeAppCount;
    }

    /// @inheritdoc IAppController
    function calculateAppId(address deployer, bytes32 salt) public view returns (IApp) {
        return IApp(
            Create2.computeAddress(
                _calculateAppMixedSalt(deployer, salt),
                keccak256(_calculateAppInitCode(deployer)) //bytecode
            )
        );
    }

    /// @inheritdoc IAppController
    function getAppStatus(IApp app) external view returns (AppStatus) {
        return _appConfigs[app].status;
    }

    /// @inheritdoc IAppController
    function getAppCreator(IApp app) external view returns (address) {
        return _appConfigs[app].creator;
    }

    /**
     * @notice Resolves the billing account for an app
     * @param app The app instance to resolve billing for
     * @return The billing account address (app address if ISOLATED, creator if DEFAULT)
     */
    function getBillingAccount(IApp app) public view returns (address) {
        AppConfigStorage storage config = _appConfigs[app];
        return config.billingType == BillingType.ISOLATED ? address(app) : config.creator;
    }

    /// @inheritdoc IAppController
    function getBillingType(IApp app) external view returns (BillingType) {
        return _appConfigs[app].billingType;
    }

    /// @inheritdoc IAppController
    function getAppTimelocked(IApp app) external view returns (bool) {
        return _appConfigs[app].timelocked;
    }

    /// @inheritdoc IAppController
    function getAppOperatorSetId(IApp app) external view returns (uint32) {
        return _appConfigs[app].operatorSetId;
    }

    /// @inheritdoc IAppController
    function getAppLatestReleaseBlockNumber(IApp app) external view returns (uint32) {
        return _appConfigs[app].latestReleaseBlockNumber;
    }

    /// @inheritdoc IAppController
    function getApps(uint256 offset, uint256 limit)
        external
        view
        returns (IApp[] memory apps, AppConfig[] memory appConfigsMem)
    {
        uint256 totalApps = _allApps.length();
        if (offset >= totalApps) return (new IApp[](0), new AppConfig[](0));

        uint256 end = offset + limit > totalApps ? totalApps : offset + limit;
        uint256 rangeSize = end - offset;

        apps = new IApp[](rangeSize);
        appConfigsMem = new AppConfig[](rangeSize);
        for (uint256 i = 0; i < rangeSize; i++) {
            apps[i] = IApp(_allApps.at(offset + i));
            appConfigsMem[i] = _toAppConfig(_appConfigs[apps[i]]);
        }
    }

    /// @inheritdoc IAppController
    function getAppsByDeveloper(address developer, uint256 offset, uint256 limit)
        external
        view
        returns (IApp[] memory apps, AppConfig[] memory appConfigsMem)
    {
        return _getFilteredApps(_isDeveloper, developer, offset, limit);
    }

    /// @inheritdoc IAppController
    function getAppsByCreator(address creator, uint256 offset, uint256 limit)
        external
        view
        returns (IApp[] memory apps, AppConfig[] memory appConfigsMem)
    {
        return _getFilteredApps(_isCreator, creator, offset, limit);
    }

    /// @inheritdoc IAppController
    function getAppsByBillingAccount(address account, uint256 offset, uint256 limit)
        external
        view
        returns (IApp[] memory apps, AppConfig[] memory appConfigsMem)
    {
        return _getFilteredApps(_isBilledTo, account, offset, limit);
    }

    /// @inheritdoc IAppController
    function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) external view returns (bytes32) {
        return _calculateSignableDigest(keccak256(abi.encode(API_PERMISSION_TYPEHASH, permission, expiry)));
    }
}
