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
    IReleaseManager, IReleaseManagerTypes
} from "@eigenlayer-contracts/src/contracts/interfaces/IReleaseManager.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IComputeAVSRegistrar} from "./interfaces/IComputeAVSRegistrar.sol";
import {IComputeOperator} from "./interfaces/IComputeOperator.sol";
import {AppControllerStorage} from "./storage/AppControllerStorage.sol";
import {IAppController} from "./interfaces/IAppController.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import {IBeacon} from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import {ComputeBilling} from "./billing/ComputeBilling.sol";
import {IApp} from "./interfaces/IApp.sol";

contract AppController is Initializable, SignatureUtilsMixin, PermissionControllerMixin, AppControllerStorage, ComputeBilling {
    using EnumerableSet for EnumerableSet.AddressSet;

    /// MODIFIERS

    /// @notice Modifier to ensure app is not in the given status
    modifier appIsActive(IApp app) {
        require(_appConfigs[app].status != AppStatus.TERMINATED, InvalidAppStatus());
        require(_appConfigs[app].status != AppStatus.NONE, InvalidAppStatus());
        _;
    }

    /**
     * @param _version The version string to use for this contract's domain separator
     * @param _permissionController The PermissionController contract address
     * @param _computeAVSRegistrar The ComputeAVSRegistrar contract address
     * @param _computeOperator The ComputeOperator contract address
     * @param _appBeacon The beacon for creating App proxies
     * @param _billingCore The BillingCore contract address
     */
    constructor(
        string memory _version,
        IPermissionController _permissionController,
        IReleaseManager _releaseManager,
        IComputeAVSRegistrar _computeAVSRegistrar,
        IComputeOperator _computeOperator,
        IBeacon _appBeacon,
        address _billingCore
    )
        SignatureUtilsMixin(_version)
        PermissionControllerMixin(_permissionController)
        AppControllerStorage(_releaseManager, _computeOperator, _computeAVSRegistrar, _appBeacon)
        ComputeBilling(_billingCore)
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
        _userConfigs[user].maxActiveApps = limit;
        emit MaxActiveAppsSet(user, limit);
    }

    /// @inheritdoc IAppController
    function setMaxGlobalActiveApps(uint32 limit) external checkCanCall(address(this)) {
        maxGlobalActiveApps = limit;
        emit GlobalMaxActiveAppsSet(limit);
    }

    /// @inheritdoc IAppController
    function createApp(bytes32 salt, Release calldata release) external returns (IApp app) {
        return createApp(salt, release, 0, address(0));
    }

    /// @inheritdoc IAppController
    function createApp(bytes32 salt, Release calldata release, uint16 skuID, address account) public returns (IApp app) {
        UserConfig storage userConfig = _userConfigs[msg.sender];

        // Check global active app limit
        require(globalActiveAppCount < maxGlobalActiveApps, GlobalMaxActiveAppsExceeded());
        // Check user active app limit
        require(userConfig.activeAppCount < userConfig.maxActiveApps, MaxActiveAppsExceeded());

        // Increment active app counts
        globalActiveAppCount++;
        userConfig.activeAppCount++;

        // Assign an operator set ID to the app
        uint32 operatorSetId = computeAVSRegistrar.assignOperatorSetId();

        // Fork based on billing status
        if (skuID != 0) {
            // Billing-enabled app: Resource caps are checked in _startBilling, skip allowlist
            // Publish the initial release metadata URI
            releaseManager.publishMetadataURI(
                OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId}), "https://eigencloud.xyz"
            );
        } else {
            // Non-billing app: Use allowlist-based access control
            // Publish the initial release metadata URI
            releaseManager.publishMetadataURI(
                OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId}), "https://eigencloud.xyz"
            );

            // Create the operator set
            computeAVSRegistrar.createOperatorSet(operatorSetId);

            // Add the compute operator to the allowlist
            computeAVSRegistrar.addOperatorToAllowlist(
                OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId}), address(computeOperator)
            );

            // Register the compute operator for the operator set
            computeOperator.registerForOperatorSet(operatorSetId);
        }

        // Create app using BeaconProxy
        app = IApp(Create2.deploy(0, _calculateAppMixedSalt(msg.sender, salt), _calculateAppInitCode(msg.sender)));
        _appConfigs[app].operatorSetId = operatorSetId;
        _appConfigs[app].latestReleaseBlockNumber = 0;
        _appConfigs[app].creator = msg.sender;
        _appConfigs[app].skuID = skuID;
        _appConfigs[app].account = account;
        _allApps.add(address(app));

        emit AppCreated(msg.sender, app, operatorSetId);

        // Upgrade the app with the initial release
        _upgradeApp(app, release);
        _startApp(app);

        // Start billing for the app if billing is enabled (apps start in STARTED state)
        if (skuID != 0) {
            _startBilling(address(app), skuID, account);
        }
    }

    /// @inheritdoc IAppController
    function upgradeApp(IApp app, Release calldata release)
        external
        checkCanCall(address(app))
        appIsActive(app)
        returns (uint256)
    {
        return _upgradeApp(app, release);
    }

    /// @inheritdoc IAppController
    function startApp(IApp app) external checkCanCall(address(app)) {
        AppConfig storage config = _appConfigs[app];

        // If transitioning from STOPPED to STARTED and billing is enabled, update billing rate
        if (config.skuID != 0 && config.status == AppStatus.STOPPED) {
            _setRunningRate(address(app), config.account, config.skuID);
        }

        _startApp(app);
    }

    /// @inheritdoc IAppController
    function stopApp(IApp app) external checkCanCall(address(app)) {
        AppConfig storage config = _appConfigs[app];
        require(config.status == AppStatus.STARTED, InvalidAppStatus());
        config.status = AppStatus.STOPPED;

        // Update billing rate to stopped rate if billing is enabled
        if (config.skuID != 0) {
            _setStoppedRate(address(app), config.account, config.skuID);
        }

        emit AppStopped(app);
    }

    /// @inheritdoc IAppController
    function terminateApp(IApp app) external checkCanCall(address(app)) appIsActive(app) {
        _terminateApp(app);
    }

    /// @inheritdoc IAppController
    function terminateAppByAdmin(IApp app) external checkCanCall(address(this)) appIsActive(app) {
        _terminateApp(app);
        emit AppTerminatedByAdmin(app);
    }

    /// INTERNAL FUNCTIONS

    /**
     * @notice Upgrades an app to a new release by publishing it through the release manager
     * @param app The app instance to upgrade
     * @param release The new release data containing artifacts and metadata
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
        _appConfigs[app].status = AppStatus.STARTED;
        emit AppStarted(app);
    }

    /**
     * @notice Terminates an app and decrements active app counters
     * @param app The app instance to terminate
     */
    function _terminateApp(IApp app) internal {
        AppConfig storage config = _appConfigs[app];

        // Remove billing before terminating if billing is enabled
        if (config.skuID != 0) {
            // Determine if app is currently running or stopped
            bool isRunning = (config.status == AppStatus.STARTED);
            _removeBilling(address(app), config.account, config.skuID, isRunning);
        }

        config.status = AppStatus.TERMINATED;

        // Decrement active app counts to free up capacity
        globalActiveAppCount--;

        // Decrement the creator's active app count
        address appCreator = config.creator;
        _userConfigs[appCreator].activeAppCount--;

        emit AppTerminated(app);
    }

    /**
     * @notice Calculates a mixed salt for app deployment using deployer address and provided salt
     * @param deployer The address of the app deployer
     * @param salt The salt value
     */
    function _calculateAppMixedSalt(address deployer, bytes32 salt) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(deployer, salt));
    }

    /**
     * @notice Generates the initialization code for deploying a new app using beacon proxy pattern
     * @param deployer The address that will be set as the app owner during initialization
     */
    function _calculateAppInitCode(address deployer) internal view returns (bytes memory) {
        return abi.encodePacked(
            _BEACON_PROXY_BYTECODE, abi.encode(appBeacon, abi.encodeWithSelector(IApp.initialize.selector, deployer))
        );
    }

    /// ADMIN BILLING FUNCTIONS

    /**
     * @notice Enables billing for an existing app that was created without billing
     * @param app The app to enable billing for
     * @param skuID The SKU identifier for billing
     * @param account The billing account that will pay for this app
     * @dev Caller must be permissioned for the AppController
     * @dev App must be active (not terminated) and not already billed (skuID == 0)
     */
    function enableAppBilling(IApp app, uint16 skuID, address account) external checkCanCall(address(this)) appIsActive(app) {
        AppConfig storage config = _appConfigs[app];
        require(config.skuID == 0, "Billing already enabled");
        require(skuID != 0, "Invalid SKU");

        // Update config
        config.skuID = skuID;
        config.account = account;

        // Start billing based on current app state
        if (config.status == AppStatus.STARTED) {
            _startBilling(address(app), skuID, account);
        } else if (config.status == AppStatus.STOPPED) {
            // For stopped apps, we need to register them with the stopped rate
            // First "start" billing (which registers with running rate)
            _startBilling(address(app), skuID, account);
            // Then immediately switch to stopped rate
            _setStoppedRate(address(app), account, skuID);
        }
    }

    /**
     * @notice Sets the billing rates and resource requirements for a SKU
     * @param skuID The SKU identifier
     * @param name The human-readable name for the SKU (e.g., "n1-standard-2")
     * @param runningRate The cost per second when app is running
     * @param stoppedRate The cost per second when app is stopped
     * @param vcpus The number of virtual CPUs required
     * @dev Caller must be permissioned for the AppController
     */
    function setSKURate(uint16 skuID, string calldata name, uint96 runningRate, uint96 stoppedRate, uint16 vcpus) external checkCanCall(address(this)) {
        _setSKURate(skuID, name, runningRate, stoppedRate, vcpus);
    }

    /**
     * @notice Sets the global resource capacity limits
     * @param vcpuCap The maximum vCPUs allowed globally
     * @param vmInstanceCap The maximum VM instances allowed globally
     * @dev Caller must be permissioned for the AppController
     */
    function setResourceCap(uint16 vcpuCap, uint16 vmInstanceCap) external checkCanCall(address(this)) {
        _setResourceCap(vcpuCap, vmInstanceCap);
    }

    /**
     * @notice Changes the SKU for an app
     * @param app The app to change the SKU for
     * @param newSKUID The new SKU identifier
     * @dev Caller must be permissioned for the AppController
     * @dev App must be active (not terminated)
     */
    function changeAppSKU(IApp app, uint16 newSKUID) external checkCanCall(address(app)) appIsActive(app) {
        AppConfig storage config = _appConfigs[app];
        bool isRunning = (config.status == AppStatus.STARTED);

        _changeSKU(address(app), config.account, config.skuID, newSKUID, isRunning);

        // Update stored SKU
        config.skuID = newSKUID;
    }

    /**
     * @notice Changes the billing account for an app
     * @param app The app to change the billing account for
     * @param newAccount The new billing account address
     * @dev Caller must be permissioned for the AppController
     * @dev App must be active (not terminated)
     */
    function changeAppBillingAccount(IApp app, address newAccount) external checkCanCall(address(this)) appIsActive(app) {
        AppConfig storage config = _appConfigs[app];
        bool isRunning = (config.status == AppStatus.STARTED);

        _changeAccount(address(app), config.account, newAccount, config.skuID, isRunning);

        // Update stored billing account
        config.account = newAccount;
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
    function calculateAppId(address deployer, bytes32 salt) external view returns (IApp) {
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
            appConfigsMem[i] = _appConfigs[apps[i]];
        }
    }

    /// @inheritdoc IAppController
    function getAppsByDeveloper(address developer, uint256 offset, uint256 limit)
        external
        view
        returns (IApp[] memory apps, AppConfig[] memory appConfigsMem)
    {
        uint256 totalApps = _allApps.length();

        apps = new IApp[](limit);
        appConfigsMem = new AppConfig[](limit);
        uint256 skipped = 0;
        uint256 found = 0;
        for (uint256 i = 0; i < totalApps && found < limit; i++) {
            IApp app = IApp(_allApps.at(i));
            if (permissionController.isAdmin(address(app), developer)) {
                if (skipped < offset) {
                    skipped++;
                    continue;
                }

                apps[found] = app;
                appConfigsMem[found] = _appConfigs[app];
                found++;
            }
        }

        // Resize arrays to actual number found
        assembly {
            mstore(apps, found)
            mstore(appConfigsMem, found)
        }
    }

    /// @inheritdoc IAppController
    function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) external view returns (bytes32) {
        return _calculateSignableDigest(keccak256(abi.encode(API_PERMISSION_TYPEHASH, permission, expiry)));
    }
}
