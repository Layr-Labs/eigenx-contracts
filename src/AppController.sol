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
import {ISablierFlow} from "@sablier/flow/src/interfaces/ISablierFlow.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {UD21x18} from "@prb/math/src/UD21x18.sol";
import {IApp} from "./interfaces/IApp.sol";

contract AppController is Initializable, SignatureUtilsMixin, PermissionControllerMixin, AppControllerStorage {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;

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
     * @param _flow The Sablier Flow contract for billing
     */
    constructor(
        string memory _version,
        IPermissionController _permissionController,
        IReleaseManager _releaseManager,
        IComputeAVSRegistrar _computeAVSRegistrar,
        IComputeOperator _computeOperator,
        IBeacon _appBeacon,
        ISablierFlow _flow
    )
        SignatureUtilsMixin(_version)
        PermissionControllerMixin(_permissionController)
        AppControllerStorage(_releaseManager, _computeOperator, _computeAVSRegistrar, _appBeacon, _flow)
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
    function setMaxGlobalVCPUs(uint32 limit) external checkCanCall(address(this)) {
        maxGlobalVCPUs = limit;
        emit MaxGlobalVCPUsSet(limit);
    }

    /// @inheritdoc IAppController
    function setSkuVCPUs(uint8 skuTier, uint32 vCPUs) external checkCanCall(address(this)) {
        skuVCPUs[skuTier] = vCPUs;
        emit SkuVCPUsSet(skuTier, vCPUs);
    }

    /// @inheritdoc IAppController
    function createApp(
        bytes32 salt,
        Release calldata release,
        uint8 skuTier,
        IERC20 billingToken,
        uint256 billingAccountId,
        uint128 initialDeposit
    ) external returns (IApp app) {
        // Check: SKU rates must be configured for this token and tier
        require(billingRates[billingToken][skuTier].running > 0, SkuRatesNotSet());

        // Check: SKU vCPUs must be configured
        uint32 appVCPUs = skuVCPUs[skuTier];
        require(appVCPUs > 0, SkuVCPUsNotSet());

        // Check: global vCPU limit
        require(globalVCPUCount + appVCPUs <= maxGlobalVCPUs, GlobalMaxVCPUsExceeded());

        // Increment global vCPU count
        globalVCPUCount += appVCPUs;

        // Assign an operator set ID to the app
        uint32 operatorSetId = computeAVSRegistrar.assignOperatorSetId();

        // Publish the initial release metadata URI
        releaseManager.publishMetadataURI(
            OperatorSet({avs: address(computeAVSRegistrar), id: operatorSetId}), "https://eigencloud.xyz"
        );

        // Create app using BeaconProxy
        app = IApp(Create2.deploy(0, _calculateAppMixedSalt(msg.sender, salt), _calculateAppInitCode(msg.sender)));
        _appConfigs[app].operatorSetId = operatorSetId;
        _appConfigs[app].latestReleaseBlockNumber = 0;
        _appConfigs[app].creator = msg.sender;
        _appConfigs[app].skuTier = skuTier;
        _appConfigs[app].billingToken = billingToken;
        _allApps.add(address(app));

        // Handle billing account creation or reuse
        uint256 accountId = _setupBillingAccount(app, skuTier, billingToken, billingAccountId);

        // Store billing info in app config
        _appConfigs[app].billingAccountId = accountId;
        _appConfigs[app].currentRate = billingRates[billingToken][skuTier].stopped;

        // Track app on this billing account
        _appsOnBillingAccount[accountId].add(address(app));

        // Handle optional initial deposit
        if (initialDeposit > 0) {
            address accountOwner = flow.getSender(accountId);
            address recipient = flow.getRecipient(accountId);
            _deposit(accountId, initialDeposit, billingToken, accountOwner, recipient);
        }

        emit AppCreated(msg.sender, app, operatorSetId);

        // Upgrade the app with the initial release
        _upgradeApp(app, release);
        _startApp(app);
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
        _startApp(app);
    }

    /// @inheritdoc IAppController
    function stopApp(IApp app) external checkCanCall(address(app)) {
        AppConfig storage config = _appConfigs[app];
        require(config.status == AppStatus.STARTED, InvalidAppStatus());
        config.status = AppStatus.STOPPED;

        // Switch billing rate from running to stopped
        uint128 newRate = billingRates[config.billingToken][config.skuTier].stopped;
        _adjustBillingRate(app, config.currentRate, newRate);
        config.currentRate = newRate;

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
        AppConfig storage config = _appConfigs[app];
        config.status = AppStatus.STARTED;

        // Switch billing rate from stopped to running
        uint128 newRate = billingRates[config.billingToken][config.skuTier].running;
        _adjustBillingRate(app, config.currentRate, newRate);
        config.currentRate = newRate;

        emit AppStarted(app);
    }

    /**
     * @notice Terminates an app and decrements active app counters
     * @param app The app instance to terminate
     */
    function _terminateApp(IApp app) internal {
        AppConfig storage config = _appConfigs[app];
        config.status = AppStatus.TERMINATED;

        uint256 billingAccountId = config.billingAccountId;

        // Remove app from billing account tracking
        _appsOnBillingAccount[billingAccountId].remove(address(app));

        // Adjust billing account rate to remove this app's contribution
        _adjustBillingRate(app, config.currentRate, 0);
        config.currentRate = 0;

        // If this was the last app on the billing account, void it to allow refund
        if (_appsOnBillingAccount[billingAccountId].length() == 0) {
            flow.void(billingAccountId);
        }

        // Free up vCPU capacity
        uint32 appVCPUs = skuVCPUs[config.skuTier];
        globalVCPUCount -= appVCPUs;

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

    /**
     * @notice Adjusts the billing account rate when switching between app states
     * @param app The app instance
     * @param oldRate The app's previous rate contribution
     * @param newRate The app's new rate contribution
     */
    function _adjustBillingRate(IApp app, uint128 oldRate, uint128 newRate) internal {
        uint256 billingAccountId = _appConfigs[app].billingAccountId;
        uint128 currentAccountRate = flow.getRatePerSecond(billingAccountId).unwrap();

        // Calculate new total account rate: current - old + new
        uint128 newAccountRateRaw = (currentAccountRate - oldRate) + newRate;
        UD21x18 newAccountRate = UD21x18.wrap(newAccountRateRaw);

        // Only adjust if account is not paused
        if (currentAccountRate > 0 || newAccountRateRaw > 0) {
            // If new rate is 0, pause the account
            if (newAccountRateRaw == 0) {
                flow.pause(billingAccountId);
            } else if (currentAccountRate == 0) {
                // If account is paused but new rate > 0, restart it
                flow.restart(billingAccountId, newAccountRate);
            } else {
                // Otherwise adjust the rate
                flow.adjustRatePerSecond(billingAccountId, newAccountRate);
            }
        }
    }

    /**
     * @notice Sets up billing account for an app - creates new or reuses existing
     * @param app The app instance
     * @param skuTier The SKU tier for billing rates
     * @param billingToken The token to use for billing
     * @param billingAccountId Billing account ID to use (0 to create new, non-zero to reuse existing)
     * @return accountId The billing account ID (either newly created or reused)
     */
    function _setupBillingAccount(IApp app, uint8 skuTier, IERC20 billingToken, uint256 billingAccountId)
        internal
        returns (uint256 accountId)
    {
        if (billingAccountId == 0) {
            // Create new billing account with stopped rate (app starts in STOPPED state)
            uint128 stoppedRate = billingRates[billingToken][skuTier].stopped;
            accountId = _createBillingAccount(msg.sender, billingToken, stoppedRate);
            emit BillingAccountCreated(app, accountId, billingToken, skuTier);
            return accountId;
        }

        // Reuse existing billing account
        // Check: billing account must exist and caller must have permission
        require(flow.isStream(billingAccountId), BillingAccountDoesNotExist());
        address accountOwner = flow.getSender(billingAccountId);
        require(_checkCanCall(accountOwner), NotBillingAccountOwner());

        // Check: billing token must match
        require(address(flow.getToken(billingAccountId)) == address(billingToken), BillingTokenMismatch());

        // Adjust the existing account's rate to account for new app
        uint128 currentAccountRate = flow.getRatePerSecond(billingAccountId).unwrap();
        uint128 newRateRaw = currentAccountRate + billingRates[billingToken][skuTier].stopped;
        UD21x18 newRate = UD21x18.wrap(newRateRaw);

        // Adjust rate based on account state
        if (currentAccountRate == 0 && newRateRaw > 0) {
            // Account is paused - restart it with new rate
            flow.restart(billingAccountId, newRate);
        } else if (currentAccountRate > 0) {
            // Account is active - adjust the rate
            flow.adjustRatePerSecond(billingAccountId, newRate);
        }

        emit AppAddedToBillingAccount(app, billingAccountId);

        return billingAccountId;
    }

    /// ADMIN FUNCTIONS

    /// @inheritdoc IAppController
    function setSkuRates(IERC20 token, uint8 skuTier, uint128 runningRate, uint128 stoppedRate)
        external
        checkCanCall(address(this))
    {
        billingRates[token][skuTier] = BillingRates({stopped: stoppedRate, running: runningRate});
        emit SkuRatesSet(token, skuTier, runningRate, stoppedRate);
    }

    /// @inheritdoc IAppController
    function setBillingRecipient(address recipient) external checkCanCall(address(this)) {
        billingRecipient = recipient;
        emit BillingRecipientSet(recipient);
    }

    /// USER BILLING FUNCTIONS

    /// @inheritdoc IAppController
    function createBillingAccount(IERC20 billingToken, uint8 skuTier, uint128 initialDeposit)
        external
        returns (uint256 billingAccountId)
    {
        // Check: SKU rates must be configured for this token and tier
        require(billingRates[billingToken][skuTier].running > 0, SkuRatesNotSet());

        // Create new billing account with rate = 0 (paused until app is added)
        billingAccountId = _createBillingAccount(msg.sender, billingToken, 0);

        // Handle optional initial deposit
        if (initialDeposit > 0) {
            _deposit(billingAccountId, initialDeposit, billingToken, msg.sender, billingRecipient);
        }

        emit BillingAccountCreated(IApp(address(0)), billingAccountId, billingToken, skuTier);
    }

    /// @inheritdoc IAppController
    function deposit(uint256 billingAccountId, uint128 amount) external {
        // Check: caller must have permission to manage the billing account
        address accountOwner = flow.getSender(billingAccountId);
        require(_checkCanCall(accountOwner), NotBillingAccountOwner());

        // Get the token and recipient for the deposit
        IERC20 token = flow.getToken(billingAccountId);
        address recipient = flow.getRecipient(billingAccountId);

        _deposit(billingAccountId, amount, token, accountOwner, recipient);
    }

    /**
     * @notice Internal function to handle token deposits to billing accounts
     * @param billingAccountId The billing account ID
     * @param amount The amount to deposit
     * @param token The token to deposit
     * @param accountOwner The account owner
     * @param recipient The account recipient
     */
    function _deposit(uint256 billingAccountId, uint128 amount, IERC20 token, address accountOwner, address recipient)
        internal
    {
        // Transfer tokens from caller to this contract
        token.safeTransferFrom(msg.sender, address(this), amount);

        // Approve SablierFlow to spend the tokens
        token.forceApprove(address(flow), amount);

        // Deposit to the account (SablierFlow will pull tokens from this contract)
        flow.deposit(billingAccountId, amount, accountOwner, recipient);
    }

    /**
     * @notice Internal helper to create a new billing account
     * @param sender The account sender (owner)
     * @param billingToken The token to use for billing
     * @param initialRate The initial rate per second
     * @return accountId The newly created billing account ID
     */
    function _createBillingAccount(address sender, IERC20 billingToken, uint128 initialRate)
        internal
        returns (uint256 accountId)
    {
        accountId = flow.create({
            sender: sender,
            recipient: billingRecipient,
            ratePerSecond: UD21x18.wrap(initialRate),
            token: billingToken,
            transferable: false
        });
    }

    /// @inheritdoc IAppController
    function refund(uint256 billingAccountId, uint128 amount) external {
        // Check: caller must have permission to manage the billing account
        address accountOwner = flow.getSender(billingAccountId);
        require(_checkCanCall(accountOwner), NotBillingAccountOwner());

        // Refund from the account
        flow.refund(billingAccountId, amount);
    }

    /// VIEW FUNCTIONS

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

    /// @inheritdoc IAppController
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
        )
    {
        AppConfig storage config = _appConfigs[app];
        billingAccountId = config.billingAccountId;
        skuTier = config.skuTier;
        token = config.billingToken;

        // Get live data from Sablier Flow
        balance = flow.getBalance(billingAccountId);
        ratePerSecond = flow.getRatePerSecond(billingAccountId).unwrap();

        // Get depletion time if account is not paused
        if (ratePerSecond > 0) {
            depletionTime = flow.depletionTimeOf(billingAccountId);
        }
    }

    /// @inheritdoc IAppController
    function getAppsOnBillingAccount(uint256 billingAccountId) external view returns (IApp[] memory apps) {
        uint256 length = _appsOnBillingAccount[billingAccountId].length();
        apps = new IApp[](length);
        for (uint256 i = 0; i < length; i++) {
            apps[i] = IApp(_appsOnBillingAccount[billingAccountId].at(i));
        }
    }
}
