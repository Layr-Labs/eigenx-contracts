// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IApp} from "./IApp.sol";

/**
 * @title IAppAuthority
 * @notice Per-app RBAC + owner registry. Centralizes the auth surface that
 *         AppController previously owned inline — scope ownership (the
 *         "owner" concept formerly called `creator`) and team roles
 *         (ADMIN / PAUSER / DEVELOPER).
 *
 * @dev Design invariants:
 *      1. Every scope (app) has at most one owner. `scopeOwner[app] == 0`
 *         means the scope has not been initialized.
 *      2. The owner is always ADMIN on their scope. `transferOwnership` is
 *         the only path that rotates the owner; it adds the new owner to
 *         ADMIN and removes the previous owner atomically.
 *      3. ADMIN role mutations (grant/revoke) are owner-only. The owner
 *         cannot renounce or self-revoke ADMIN — `transferOwnership` is the
 *         only path out.
 *      4. Critical ops gated at the consumer layer (upgrade/transfer/
 *         terminate on AppController) use `isScopeOwner` — NOT
 *         `hasRole(ADMIN)`. ADMIN is operational-only in this model.
 *      5. `canCall` validates schedule-time that a Timelock operation
 *         targeting a consumer contract will pass runtime gates, so a
 *         doomed op doesn't burn a pending-op slot on the Timelock.
 */
interface IAppAuthority {
    /// @notice Thrown when a caller is not the current owner of the scope.
    error NotScopeOwner();

    /// @notice Thrown when a caller lacks the required role.
    error InvalidRole();

    /// @notice Thrown when the owner attempts to renounce or self-revoke ADMIN.
    error CannotRemoveOwnerAdmin();

    /// @notice Thrown when revoking/renouncing ADMIN would leave zero admins.
    error CannotRemoveLastAdmin();

    /// @notice Thrown when a scope already has an owner and an attempt is
    ///         made to initialize it.
    error ScopeAlreadyInitialized();

    /// @notice Thrown when an attempted operation requires the scope to be
    ///         initialized but it isn't.
    error ScopeNotInitialized();

    /// @notice Thrown when the caller is not the authorized consumer
    ///         (AppController) and a consumer-only method is called.
    error OnlyConsumer();

    /// @notice Thrown when a zero address is passed where non-zero is required.
    error ZeroAddress();

    /// @notice Emitted when a scope is first initialized with an owner.
    event ScopeInitialized(IApp indexed scope, address indexed owner);

    /// @notice Emitted when scope ownership is transferred.
    event ScopeOwnershipTransferred(IApp indexed scope, address indexed previousOwner, address indexed newOwner);

    /// @notice Emitted when a role is granted on a scope.
    event RoleGranted(IApp indexed scope, Role indexed role, address indexed account);

    /// @notice Emitted when a role is revoked on a scope (by owner, admin, or self-renounce).
    event RoleRevoked(IApp indexed scope, Role indexed role, address indexed account);

    /// @notice Per-app team roles. ADMIN is operational-only; it does NOT
    ///         convey power over critical ops. PAUSER and DEVELOPER are
    ///         delegable subsets of ADMIN's operational surface.
    enum Role {
        ADMIN,
        PAUSER,
        DEVELOPER
    }

    /**
     * @notice Initialize a scope's owner. May only be called by the consumer
     *         (AppController) once per scope, at app creation time. Grants
     *         the owner ADMIN atomically.
     * @dev Consumer-only; reverts `OnlyConsumer` from any other caller.
     */
    function initializeScope(IApp scope, address owner) external;

    /**
     * @notice Transfer ownership of a scope to a new address. Atomically
     *         adds `newOwner` to ADMIN and removes `previousOwner` from
     *         ADMIN. Consumer-only — the consumer contract is responsible
     *         for authenticating the caller.
     * @dev Consumer-only; reverts `OnlyConsumer` from any other caller.
     */
    function transferScopeOwnership(IApp scope, address newOwner) external;

    /**
     * @notice Grant a role to `account` on `scope`.
     * @dev For `role == ADMIN`: caller must be the scope's owner.
     * @dev For non-ADMIN: caller must hold ADMIN on the scope.
     */
    function grantRole(IApp scope, Role role, address account) external;

    /**
     * @notice Revoke a role from `account` on `scope`.
     * @dev For `role == ADMIN`: caller must be the scope's owner. The owner
     *      cannot revoke their own ADMIN. Revoking below the last-ADMIN
     *      floor reverts.
     * @dev For non-ADMIN: caller must hold ADMIN on the scope.
     */
    function revokeRole(IApp scope, Role role, address account) external;

    /**
     * @notice Renounce your own role on `scope`.
     * @dev The owner cannot renounce ADMIN. Non-owners may renounce ADMIN
     *      as long as the set doesn't empty.
     */
    function renounceRole(IApp scope, Role role) external;

    /**
     * @notice Migrate admins from an external source (e.g., legacy
     *         PermissionController) into the ADMIN role on each supplied
     *         scope. Consumer-only. Idempotent.
     * @dev For each (scope, admin) pair, grants ADMIN if not already present.
     *      Refuses to migrate into scopes that have already been initialized
     *      if the admin would conflict with the owner invariant — specifically,
     *      migrated admins are operational-only under this model, so the
     *      downgraded-from-critical-op semantics are already applied.
     */
    function migrateAdmins(IApp[] calldata scopes, address[][] calldata admins) external;

    /// @notice The current owner of a scope, or address(0) if uninitialized.
    function scopeOwner(IApp scope) external view returns (address);

    /// @notice Whether `account` is the current owner of `scope`.
    function isScopeOwner(IApp scope, address account) external view returns (bool);

    /// @notice Whether `account` holds `role` on `scope`.
    function hasRole(IApp scope, Role role, address account) external view returns (bool);

    /// @notice Whether `account` holds `role` on `scope`, OR holds ADMIN
    ///         (which is a superset for operational roles).
    function hasRoleOrAdmin(IApp scope, Role role, address account) external view returns (bool);

    /// @notice The consumer contract (AppController) that owns the
    ///         consumer-only mutation surface.
    function consumer() external view returns (address);
}
