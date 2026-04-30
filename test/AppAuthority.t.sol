// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {AppAuthority} from "../src/governance/AppAuthority.sol";
import {IAppAuthority} from "../src/interfaces/IAppAuthority.sol";
import {IApp} from "../src/interfaces/IApp.sol";

contract AppAuthorityTest is Test {
    AppAuthority internal authority;

    address internal consumer = makeAddr("consumer");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");
    address internal carol = makeAddr("carol");

    IApp internal constant APP_A = IApp(address(0xA0A0));
    IApp internal constant APP_B = IApp(address(0xB0B0));

    function setUp() public {
        // Deploy the impl directly. Tests don't need the proxy layer —
        // they exercise pure authority logic. `_disableInitializers` in the
        // constructor means initialize() is unreachable, which is fine: the
        // initializer body is a no-op.
        authority = new AppAuthority(consumer);
    }

    // ========== scope initialization ==========

    function test_initializeScope_consumerOnly() public {
        vm.prank(alice);
        vm.expectRevert(IAppAuthority.OnlyConsumer.selector);
        authority.initializeScope(APP_A, alice);
    }

    function test_initializeScope_rejectsZeroOwner() public {
        vm.prank(consumer);
        vm.expectRevert(IAppAuthority.ZeroAddress.selector);
        authority.initializeScope(APP_A, address(0));
    }

    function test_initializeScope_seedsOwnerAndAdmin() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        assertEq(authority.scopeOwner(APP_A), alice);
        assertTrue(authority.isScopeOwner(APP_A, alice));
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, alice));
    }

    function test_initializeScope_rejectsReinit() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(consumer);
        vm.expectRevert(IAppAuthority.ScopeAlreadyInitialized.selector);
        authority.initializeScope(APP_A, bob);
    }

    // ========== ownership transfer ==========

    function test_transferScopeOwnership_consumerOnly() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(alice); // even the owner can't directly — goes via consumer
        vm.expectRevert(IAppAuthority.OnlyConsumer.selector);
        authority.transferScopeOwnership(APP_A, bob);
    }

    function test_transferScopeOwnership_rotatesAdmin() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(consumer);
        authority.transferScopeOwnership(APP_A, bob);

        assertEq(authority.scopeOwner(APP_A), bob);
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, bob));
        assertFalse(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, alice));
    }

    function test_transferScopeOwnership_rejectsUninitialized() public {
        vm.prank(consumer);
        vm.expectRevert(IAppAuthority.ScopeNotInitialized.selector);
        authority.transferScopeOwnership(APP_A, bob);
    }

    function test_transferScopeOwnership_rejectsZeroNewOwner() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(consumer);
        vm.expectRevert(IAppAuthority.ZeroAddress.selector);
        authority.transferScopeOwnership(APP_A, address(0));
    }

    // ========== grantRole ==========

    function test_grantRole_adminRequiresOwner() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        // alice (owner) can grant ADMIN
        vm.prank(alice);
        authority.grantRole(APP_A, IAppAuthority.Role.ADMIN, bob);
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, bob));

        // bob (ADMIN but not owner) cannot grant ADMIN
        vm.prank(bob);
        vm.expectRevert(IAppAuthority.NotScopeOwner.selector);
        authority.grantRole(APP_A, IAppAuthority.Role.ADMIN, carol);
    }

    function test_grantRole_pauserRequiresAdmin() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        // alice (owner, so ADMIN) can grant PAUSER
        vm.prank(alice);
        authority.grantRole(APP_A, IAppAuthority.Role.PAUSER, bob);
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.PAUSER, bob));

        // carol (not ADMIN) cannot grant PAUSER
        vm.prank(carol);
        vm.expectRevert(IAppAuthority.InvalidRole.selector);
        authority.grantRole(APP_A, IAppAuthority.Role.PAUSER, carol);

        // bob (granted ADMIN manually for this test) can grant PAUSER
        vm.prank(alice);
        authority.grantRole(APP_A, IAppAuthority.Role.ADMIN, bob);

        vm.prank(bob);
        authority.grantRole(APP_A, IAppAuthority.Role.PAUSER, carol);
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.PAUSER, carol));
    }

    function test_grantRole_rejectsZeroAccount() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(alice);
        vm.expectRevert(IAppAuthority.ZeroAddress.selector);
        authority.grantRole(APP_A, IAppAuthority.Role.ADMIN, address(0));
    }

    // ========== revokeRole ==========

    function test_revokeRole_adminRequiresOwner() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(alice);
        authority.grantRole(APP_A, IAppAuthority.Role.ADMIN, bob);

        // bob (ADMIN, not owner) cannot revoke alice
        vm.prank(bob);
        vm.expectRevert(IAppAuthority.NotScopeOwner.selector);
        authority.revokeRole(APP_A, IAppAuthority.Role.ADMIN, alice);

        // alice (owner) can revoke bob
        vm.prank(alice);
        authority.revokeRole(APP_A, IAppAuthority.Role.ADMIN, bob);
        assertFalse(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, bob));
    }

    function test_revokeRole_ownerCannotRevokeSelf() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(alice);
        vm.expectRevert(IAppAuthority.CannotRemoveOwnerAdmin.selector);
        authority.revokeRole(APP_A, IAppAuthority.Role.ADMIN, alice);
    }

    function test_revokeRole_cannotRemoveLastAdmin() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        // Cannot revoke alice when she's the last admin. But the owner-self
        // check fires first. Add bob then revoke alice's own role — still
        // blocked by CannotRemoveOwnerAdmin.
        vm.prank(alice);
        authority.grantRole(APP_A, IAppAuthority.Role.ADMIN, bob);

        // Remove bob: OK, alice remains.
        vm.prank(alice);
        authority.revokeRole(APP_A, IAppAuthority.Role.ADMIN, bob);

        // Can't revoke alice (owner-self block is stricter than last-admin).
        vm.prank(alice);
        vm.expectRevert(IAppAuthority.CannotRemoveOwnerAdmin.selector);
        authority.revokeRole(APP_A, IAppAuthority.Role.ADMIN, alice);
    }

    function test_revokeRole_pauserRequiresAdmin() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(alice);
        authority.grantRole(APP_A, IAppAuthority.Role.PAUSER, bob);

        // carol isn't admin — cannot revoke
        vm.prank(carol);
        vm.expectRevert(IAppAuthority.InvalidRole.selector);
        authority.revokeRole(APP_A, IAppAuthority.Role.PAUSER, bob);

        // alice (owner/admin) can revoke
        vm.prank(alice);
        authority.revokeRole(APP_A, IAppAuthority.Role.PAUSER, bob);
        assertFalse(authority.hasRole(APP_A, IAppAuthority.Role.PAUSER, bob));
    }

    // ========== renounceRole ==========

    function test_renounceRole_ownerCannotRenounceAdmin() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(alice);
        vm.expectRevert(IAppAuthority.CannotRemoveOwnerAdmin.selector);
        authority.renounceRole(APP_A, IAppAuthority.Role.ADMIN);
    }

    function test_renounceRole_coAdminCanRenounce() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(alice);
        authority.grantRole(APP_A, IAppAuthority.Role.ADMIN, bob);

        vm.prank(bob);
        authority.renounceRole(APP_A, IAppAuthority.Role.ADMIN);

        assertFalse(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, bob));
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, alice));
    }

    function test_renounceRole_operationalRoleOk() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        vm.prank(alice);
        authority.grantRole(APP_A, IAppAuthority.Role.PAUSER, bob);

        vm.prank(bob);
        authority.renounceRole(APP_A, IAppAuthority.Role.PAUSER);

        assertFalse(authority.hasRole(APP_A, IAppAuthority.Role.PAUSER, bob));
    }

    // ========== migrateAdmins ==========

    function test_migrateAdmins_consumerOnly() public {
        IApp[] memory scopes = new IApp[](1);
        address[][] memory admins = new address[][](1);
        scopes[0] = APP_A;
        admins[0] = new address[](0);

        vm.prank(alice);
        vm.expectRevert(IAppAuthority.OnlyConsumer.selector);
        authority.migrateAdmins(scopes, admins);
    }

    function test_migrateAdmins_seedsAdmins() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        IApp[] memory scopes = new IApp[](1);
        address[][] memory admins = new address[][](1);
        scopes[0] = APP_A;
        admins[0] = new address[](2);
        admins[0][0] = bob;
        admins[0][1] = carol;

        vm.prank(consumer);
        authority.migrateAdmins(scopes, admins);

        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, bob));
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, carol));
    }

    function test_migrateAdmins_idempotent() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        IApp[] memory scopes = new IApp[](1);
        address[][] memory admins = new address[][](1);
        scopes[0] = APP_A;
        admins[0] = new address[](1);
        admins[0][0] = bob;

        vm.prank(consumer);
        authority.migrateAdmins(scopes, admins);

        // Run again: safe.
        vm.prank(consumer);
        authority.migrateAdmins(scopes, admins);

        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, bob));
    }

    function test_migrateAdmins_skipsZeroAddress() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        IApp[] memory scopes = new IApp[](1);
        address[][] memory admins = new address[][](1);
        scopes[0] = APP_A;
        admins[0] = new address[](2);
        admins[0][0] = address(0);
        admins[0][1] = bob;

        vm.prank(consumer);
        authority.migrateAdmins(scopes, admins);

        assertFalse(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, address(0)));
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, bob));
    }

    // ========== hasRoleOrAdmin ==========

    function test_hasRoleOrAdmin_adminIsSupersetForOperational() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);

        // alice is ADMIN only. PAUSER query still returns true because ADMIN is
        // a superset for operational roles.
        assertTrue(authority.hasRoleOrAdmin(APP_A, IAppAuthority.Role.PAUSER, alice));
        assertTrue(authority.hasRoleOrAdmin(APP_A, IAppAuthority.Role.DEVELOPER, alice));

        // Direct role checks are strict.
        assertFalse(authority.hasRole(APP_A, IAppAuthority.Role.PAUSER, alice));
    }

    // ========== scope isolation ==========

    function test_scopesAreIsolated() public {
        vm.prank(consumer);
        authority.initializeScope(APP_A, alice);
        vm.prank(consumer);
        authority.initializeScope(APP_B, bob);

        // A-scope ADMIN on alice has no bearing on B-scope.
        assertTrue(authority.hasRole(APP_A, IAppAuthority.Role.ADMIN, alice));
        assertFalse(authority.hasRole(APP_B, IAppAuthority.Role.ADMIN, alice));

        // Owner gates are scope-local.
        vm.prank(alice);
        vm.expectRevert(IAppAuthority.NotScopeOwner.selector);
        authority.grantRole(APP_B, IAppAuthority.Role.ADMIN, carol);
    }
}
