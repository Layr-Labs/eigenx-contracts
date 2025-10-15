// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IBillingCore} from "./interfaces/IBillingCore.sol";
import {IConsumptionBilling} from "./interfaces/IConsumptionBilling.sol";

/**
 * @title ConsumptionBilling
 * @notice Abstract base contract for consumption-based billing
 * @dev Resources (bytes32) accumulate usage units, billed at end of period.
 *      Extend this contract to add product-specific logic and access control.
 */
abstract contract ConsumptionBilling is IConsumptionBilling {
    IBillingCore public immutable billingCore;

    mapping(bytes32 => APIKey) public keys;
    mapping(address => bytes32[]) public accountKeys;  // billing account => list of key IDs

    address public oracle;  // Reports usage

    constructor(address _billingCore) {
        billingCore = IBillingCore(_billingCore);
    }

    /**
     * @dev Initializes the consumption billing state variables
     * @param _oracle Address of the oracle that reports usage
     */
    function __ConsumptionBilling_init(address _oracle) internal {
        oracle = _oracle;
    }

    modifier onlyOracle() {
        require(msg.sender == oracle, NotOracle());
        _;
    }

    // ============================================================================
    // Internal Admin Functions
    // ============================================================================

    function _setOracle(address _oracle) internal {
        address oldOracle = oracle;
        oracle = _oracle;
        emit OracleChanged(oldOracle, _oracle);
    }

    // ============================================================================
    // Internal Key Lifecycle Functions
    // ============================================================================

    function _createKey(bytes32 keyId, address owner) internal {
        require(!keys[keyId].isActive && keys[keyId].owner == address(0), KeyAlreadyExists(keyId));

        keys[keyId] = APIKey({
            pendingUnits: 0,
            isActive: true,
            owner: owner,
            account: owner
        });

        // Track key under billing account
        accountKeys[owner].push(keyId);

        emit KeyCreated(keyId, owner);
    }

    function deactivateKey(bytes32 keyId) external {
        APIKey storage key = keys[keyId];
        require(msg.sender == key.owner, NotOwner());
        require(key.isActive, KeyAlreadyInactive(keyId));

        key.isActive = false;
        emit KeyDeactivated(keyId);
    }

    function setAccount(bytes32 keyId, address newAccount) external {
        APIKey storage key = keys[keyId];
        require(msg.sender == key.owner, NotOwner());
        require(newAccount != address(0), InvalidAccount());
        require(key.owner != address(0), KeyDoesNotExist(keyId));

        address oldAccount = key.account;
        if (oldAccount == newAccount) return;

        // Move key from old account to new account
        _removeKeyFromAccount(oldAccount, keyId);
        accountKeys[newAccount].push(keyId);

        // Note: Should settle before changing account to bill old account correctly
        key.account = newAccount;
        emit AccountChanged(keyId, oldAccount, newAccount);
    }

    // ============================================================================
    // Usage Recording (Oracle)
    // ============================================================================

    function recordUsage(bytes32 keyId, uint128 units) external onlyOracle {
        APIKey storage key = keys[keyId];
        require(key.isActive, KeyInactive(keyId));

        key.pendingUnits += units;
        emit UsageRecorded(keyId, units);
    }

    function recordUsageBatch(
        bytes32[] calldata keyIds,
        uint128[] calldata units
    ) external onlyOracle {
        require(keyIds.length == units.length, LengthMismatch());

        for (uint256 i = 0; i < keyIds.length; i++) {
            APIKey storage key = keys[keyIds[i]];
            if (key.isActive) {
                key.pendingUnits += units[i];
                emit UsageRecorded(keyIds[i], units[i]);
            }
        }
    }

    // ============================================================================
    // Internal Settlement Functions
    // ============================================================================

    /**
     * @notice Settle a key's usage for a given amount
     * @dev Called by extending contract at end of billing period after calculating charges
     * @param keyId The API key to settle
     * @param amount The amount to charge (based on external pricing calculation)
     * @return unitsPaid The number of units that were settled
     */
    function _settleKey(bytes32 keyId, uint128 amount) internal returns (uint128 unitsPaid) {
        APIKey storage key = keys[keyId];
        unitsPaid = key.pendingUnits;

        if (unitsPaid == 0) return 0;

        key.pendingUnits = 0;

        // Charge the billing account
        billingCore.chargeAmount(key.account, amount);

        emit KeySettled(keyId, amount, unitsPaid);
    }

    /**
     * @notice Batch settle multiple keys
     * @dev Allows settling all keys for a billing period in one transaction
     */
    function _settleKeyBatch(
        bytes32[] calldata keyIds,
        uint128[] calldata amounts
    ) internal {
        require(keyIds.length == amounts.length, LengthMismatch());

        for (uint256 i = 0; i < keyIds.length; i++) {
            if (keys[keyIds[i]].pendingUnits > 0) {
                _settleKey(keyIds[i], amounts[i]);
            }
        }
    }

    // ============================================================================
    // View Functions
    // ============================================================================

    function getAccountKeys(address account) external view returns (bytes32[] memory) {
        return accountKeys[account];
    }

    function getAccountPendingUnits(address account) external view returns (uint128 totalUnits) {
        bytes32[] memory keyIds = accountKeys[account];
        for (uint256 i = 0; i < keyIds.length; i++) {
            totalUnits += keys[keyIds[i]].pendingUnits;
        }
    }

    function getAccountKeysDetailed(address account) external view returns (
        bytes32[] memory keyIds,
        uint128[] memory pendingUnits,
        bool[] memory activeStates,
        uint128 totalPending
    ) {
        bytes32[] memory accountKeyList = accountKeys[account];
        uint256 length = accountKeyList.length;

        keyIds = new bytes32[](length);
        pendingUnits = new uint128[](length);
        activeStates = new bool[](length);

        for (uint256 i = 0; i < length; i++) {
            bytes32 keyId = accountKeyList[i];
            APIKey memory key = keys[keyId];

            keyIds[i] = keyId;
            pendingUnits[i] = key.pendingUnits;
            activeStates[i] = key.isActive;
            totalPending += key.pendingUnits;
        }
    }

    // ============================================================================
    // Internal Functions
    // ============================================================================

    function _removeKeyFromAccount(address account, bytes32 keyId) internal {
        bytes32[] storage keyList = accountKeys[account];
        for (uint256 i = 0; i < keyList.length; i++) {
            if (keyList[i] == keyId) {
                keyList[i] = keyList[keyList.length - 1];
                keyList.pop();
                break;
            }
        }
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
