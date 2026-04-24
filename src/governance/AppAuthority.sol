// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {IAppAuthority} from "../interfaces/IAppAuthority.sol";
import {IApp} from "../interfaces/IApp.sol";

/**
 * @title AppAuthority
 * @notice Per-app ownership + RBAC state, extracted from AppController.
 *         AppController is the sole consumer — it authenticates callers of
 *         app-lifecycle operations and delegates the auth state here.
 *
 * @dev The contract enforces the Option-2 invariants by construction:
 *      - Only the owner may grant/revoke/transfer ADMIN.
 *      - The owner is always ADMIN on their scope, cannot renounce or
 *        self-revoke.
 *      - `transferScopeOwnership` is the only path that rotates the owner,
 *        and it adds the new owner + removes the previous owner from ADMIN
 *        atomically.
 *
 * @dev Why extract: the Timelock guarantee binds iff the admin set the
 *      critical-op gate trusts is a set the contract enforcing the gate
 *      controls. AppController delegates to this contract; this contract's
 *      mutation surface is minimal and fully enumerable. Moving the auth
 *      logic here makes the trust boundary explicit and shrinks the audit
 *      surface of AppController accordingly.
 */
contract AppAuthority is Initializable, IAppAuthority {
    using EnumerableSet for EnumerableSet.AddressSet;

    /// @notice The consumer contract authorized to call mutation methods that
    ///         pass through app-lifecycle events (scope initialization,
    ///         ownership transfer, admin migration).
    /// @dev Set at construction. Role-level grants/revokes/renounces are NOT
    ///      consumer-gated — users interact with them directly.
    address public immutable consumer;

    /// @notice Per-scope owner. Zero means uninitialized.
    mapping(IApp => address) internal _scopeOwner;

    /// @notice Per-scope, per-role set of holders.
    mapping(IApp => mapping(Role => EnumerableSet.AddressSet)) internal _roles;

    /**
     * @dev Reserved storage for future variables.
     *      See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;

    modifier onlyConsumer() {
        if (msg.sender != consumer) revert OnlyConsumer();
        _;
    }

    constructor(address _consumer) {
        if (_consumer == address(0)) revert ZeroAddress();
        consumer = _consumer;
        _disableInitializers();
    }

    /// @notice No-op initializer. All state is derived at construction
    ///         (consumer immutable) or via `initializeScope`. Exists so the
    ///         proxy can be deployed via upgradeAndCall.
    function initialize() external initializer {}

    /// @inheritdoc IAppAuthority
    function initializeScope(IApp scope, address owner) external onlyConsumer {
        if (owner == address(0)) revert ZeroAddress();
        if (_scopeOwner[scope] != address(0)) revert ScopeAlreadyInitialized();

        _scopeOwner[scope] = owner;
        _roles[scope][Role.ADMIN].add(owner);

        emit ScopeInitialized(scope, owner);
        emit RoleGranted(scope, Role.ADMIN, owner);
    }

    /// @inheritdoc IAppAuthority
    function transferScopeOwnership(IApp scope, address newOwner) external onlyConsumer {
        if (newOwner == address(0)) revert ZeroAddress();

        address previousOwner = _scopeOwner[scope];
        if (previousOwner == address(0)) revert ScopeNotInitialized();

        _scopeOwner[scope] = newOwner;

        // Add new owner first so the ADMIN set never empties between the
        // two writes (preserves the last-admin invariant during the swap).
        if (_roles[scope][Role.ADMIN].add(newOwner)) {
            emit RoleGranted(scope, Role.ADMIN, newOwner);
        }
        if (_roles[scope][Role.ADMIN].remove(previousOwner)) {
            emit RoleRevoked(scope, Role.ADMIN, previousOwner);
        }

        emit ScopeOwnershipTransferred(scope, previousOwner, newOwner);
    }

    /// @inheritdoc IAppAuthority
    function grantRole(IApp scope, Role role, address account) external {
        if (account == address(0)) revert ZeroAddress();

        if (role == Role.ADMIN) {
            if (msg.sender != _scopeOwner[scope]) revert NotScopeOwner();
        } else {
            if (!_roles[scope][Role.ADMIN].contains(msg.sender)) revert InvalidRole();
        }

        if (_roles[scope][role].add(account)) {
            emit RoleGranted(scope, role, account);
        }
    }

    /// @inheritdoc IAppAuthority
    function revokeRole(IApp scope, Role role, address account) external {
        if (role == Role.ADMIN) {
            if (msg.sender != _scopeOwner[scope]) revert NotScopeOwner();
            if (account == _scopeOwner[scope]) revert CannotRemoveOwnerAdmin();
        } else {
            if (!_roles[scope][Role.ADMIN].contains(msg.sender)) revert InvalidRole();
        }

        if (role == Role.ADMIN && _roles[scope][Role.ADMIN].length() == 1) {
            revert CannotRemoveLastAdmin();
        }

        if (_roles[scope][role].remove(account)) {
            emit RoleRevoked(scope, role, account);
        }
    }

    /// @inheritdoc IAppAuthority
    function renounceRole(IApp scope, Role role) external {
        if (role == Role.ADMIN) {
            if (msg.sender == _scopeOwner[scope]) revert CannotRemoveOwnerAdmin();
            if (_roles[scope][Role.ADMIN].length() == 1) revert CannotRemoveLastAdmin();
        }

        if (_roles[scope][role].remove(msg.sender)) {
            emit RoleRevoked(scope, role, msg.sender);
        }
    }

    /// @inheritdoc IAppAuthority
    function migrateAdmins(IApp[] calldata scopes, address[][] calldata admins) external onlyConsumer {
        require(scopes.length == admins.length, InvalidRole());
        for (uint256 i = 0; i < scopes.length; i++) {
            IApp scope = scopes[i];
            address[] calldata scopeAdmins = admins[i];
            for (uint256 j = 0; j < scopeAdmins.length; j++) {
                address adm = scopeAdmins[j];
                if (adm == address(0)) continue;
                if (_roles[scope][Role.ADMIN].add(adm)) {
                    emit RoleGranted(scope, Role.ADMIN, adm);
                }
            }
        }
    }

    /// @inheritdoc IAppAuthority
    function scopeOwner(IApp scope) external view returns (address) {
        return _scopeOwner[scope];
    }

    /// @inheritdoc IAppAuthority
    function isScopeOwner(IApp scope, address account) external view returns (bool) {
        return _scopeOwner[scope] == account;
    }

    /// @inheritdoc IAppAuthority
    function hasRole(IApp scope, Role role, address account) external view returns (bool) {
        return _roles[scope][role].contains(account);
    }

    /// @inheritdoc IAppAuthority
    function hasRoleOrAdmin(IApp scope, Role role, address account) external view returns (bool) {
        return _roles[scope][role].contains(account) || _roles[scope][Role.ADMIN].contains(account);
    }
}
