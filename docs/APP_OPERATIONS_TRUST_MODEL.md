# App Operations & Trust Model

> **Note:** This branch includes a POC implementation of the changes described in this spec.

## Overview

This spec describes a new operational model for Eigen Compute apps, focusing on team operations and trust guarantees. The goals are:

1. Make the operational structure legible for teams building on Eigen Compute
2. Enable teams to make verifiable commitments to their users (e.g., upgrade delays that can't be bypassed)

## Design Principles

1. **Role-based, not function-based** - Simpler to reason about than per-function permissions
2. **Team-scoped** - Roles are granted per-team (owner address), not per-app
3. **Legibility** - Trust configurations must be displayable on all interfaces
4. **Factory-deployed primitives** - Timelocks and multisigs deployed via factories can be verified as legitimate

## Team Model

An app owner's address IS the team identity. This is similar to GitHub or Vercel teams, but the team is built around the owner address rather than a separate entity. This keeps things simple - no team entity to create, straightforward billing tied to the owner, and works naturally for the common case of a single EOA or multisig. Ownership can be transferred to a new address if the team structure changes (see Future Additions).

Any role (including the team owner) can be:
- An EOA (solo dev or operator)
- A multisig (shared control)
- A multisig behind a timelock (for enforced delays on sensitive operations)

This also enables a cleaner UX: manage your team once, and all apps under that team inherit the same permissions - similar to how Vercel or GitHub organization access works. App-level permissions would require managing access separately for each app, and wouldn't align with billing which is at the owner level.

## Roles

Roles are scoped to the team via `keccak256(abi.encode(teamAddress, TeamRole))`.

| Role | Capabilities |
|------|-------------|
| `ADMIN` | Full access: create apps, upgrade, start, stop, terminate, manage roles |
| `PAUSER` | Stop apps (emergency response) |
| `DEVELOPER` | Update metadata, view logs, submit builds |

Admins implicitly have all other roles. The last admin cannot be revoked or renounced (prevents accidental lockout).

## SafeTimelockFactory

Factory for deploying verified Gnosis Safes and OZ TimelockControllers. Deployed contracts can be verified via `isSafe()` and `isTimelock()`.

### Why a factory?

- Factory deployment guarantees unmodified OZ TimelockController (no bypasses)
- Enables UIs to display accurate delay information with confidence
- Supports atomic Safe + Timelock setup via EIP-7702: predict Safe address, deploy both, and configure Safe as proposer in a single transaction

### Safe Deployment

```solidity
struct SafeConfig {
    address[] owners;
    uint256 threshold;
}

deploySafe(config, salt) → deterministic address
```

### Timelock Deployment

```solidity
struct TimelockConfig {
    uint256 minDelay;
    address[] proposers;
    address[] executors;
}

deployTimelock(config, salt) → deterministic address
```

## Typical Configurations

| Setup | Admin Entity | Delay |
|-------|-------------|-------|
| Dev/testing | EOA | 0 |
| Production | Multisig behind timelock | >0 |

For emergency response, grant `PAUSER` to a lower-threshold multisig or EOA that can stop apps without going through the timelock.

## CLI Integration

### Identity (`ecloud auth`)

**Login to existing identity:**

```
ecloud auth login

> Select identity type:
  1. EOA (private key)
  2. Gnosis Safe (enter address)
```

**Create new identity:**

```
ecloud auth new

> What would you like to create?
  1. EOA (new private key)
  2. Gnosis Safe

# For Safe:
> Enter owner addresses: 0x1111..., 0x2222..., 0x3333...
> Enter threshold (e.g., 2 of 3): 2

> Add timelock delay? (y/N) y
> Enter minimum delay (e.g., "24h", "7d"): 24h

> Deploying Safe + Timelock via factory...
```

When a timelock is configured, the CLI deploys via SafeTimelockFactory with your identity as both proposer and executor.

**App actions with Safe:**

When logged in as a Safe, app actions propose a transaction and link to the Safe UI:

```
ecloud compute app upgrade

> Transaction proposed to Safe.
> View and sign at: https://app.safe.global/transactions/queue?safe=eth:0x1234...
```

**Viewing trust configuration:**

```
ecloud compute app info <app-id>

> App: my-app
> Owner: 0x1234... (3/5 Safe behind 24h timelock)
> Status: STARTED
>
> Team Roles:
>   ADMIN:     0x1234... (3/5 Safe behind 24h timelock)
>              0xabcd... (EOA)
>   PAUSER:    0x5678... (2/3 Safe, no delay)
>   DEVELOPER: 0x9999... (EOA)
```

### Team Management (`ecloud compute team`)

Manage roles for your team:

```
ecloud compute team grant <address> [role]
ecloud compute team revoke <address> [role]
ecloud compute team list
```

**Grant (interactive):**

```
ecloud compute team grant 0x1234...

> Select role to grant:
  1. ADMIN - Full access: create apps, upgrade, start, stop, terminate, manage roles
  2. PAUSER - Stop apps (emergency response)
  3. DEVELOPER - Update metadata, view logs, submit builds

> Role granted successfully.
```

**Revoke (interactive):**

```
ecloud compute team revoke 0x1234...

> 0x1234... has: PAUSER, DEVELOPER
> Select role to revoke:
  1. PAUSER
  2. DEVELOPER

> Role revoked successfully.
```

## Migration

`migrateAdmins(apps[])` is a one-time state migration performed by Eigen admins to bridge existing PermissionController admins to AccessControl.

## Future Additions

- **Timelock Operations CLI** (`ecloud compute tx list`, `ecloud compute tx execute`) - View pending operations and execute when ready. Requires Safe Transaction Service integration and event indexing.
- **[App Ownership Transfer](future/APP_OWNERSHIP_TRANSFER.md)** - Two-step transfer pattern to move apps between teams (useful for acquisitions, team restructuring)
