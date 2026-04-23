# App Ownership Transfer (Future)

Technical specification for two-step app ownership transfer.

## Overview

Transfer app ownership between teams. Useful for acquisitions, team restructuring, or handing off projects.

## Interface

### Errors

```solidity
error NoPendingTransfer();
```

### Events

```solidity
event AppTransferInitiated(IApp indexed app, address indexed currentOwner, address indexed pendingOwner);
event AppTransferCancelled(IApp indexed app);
event AppTransferred(IApp indexed app, address indexed oldOwner, address indexed newOwner);
```

### Storage

Add to `AppConfig`:

```solidity
struct AppConfig {
    address owner;
    uint32 operatorSetId;
    uint32 latestReleaseBlockNumber;
    AppStatus status;
    address pendingOwner;  // <-- add this
}
```

### Functions

```solidity
/// @notice Initiates an app transfer to a new owner
/// @param app The app to transfer
/// @param newOwner The address of the new owner
/// @dev Only team admins can initiate transfers
/// @dev App must be active (STARTED or STOPPED)
/// @dev Can be called again to change the pending owner
function initiateAppTransfer(IApp app, address newOwner) external;

/// @notice Accepts a pending app transfer
/// @param app The app to accept transfer for
/// @dev Only the pending owner can accept
/// @dev The new owner must have available quota at time of acceptance
function acceptAppTransfer(IApp app) external;

/// @notice Cancels a pending app transfer
/// @param app The app to cancel transfer for
/// @dev Only team admins can cancel transfers
function cancelAppTransfer(IApp app) external;

/// @notice Gets the pending owner for an app transfer
/// @param app The app to check
/// @return The pending owner address, or address(0) if no pending transfer
function getPendingAppOwner(IApp app) external view returns (address);
```

## Implementation

### initiateAppTransfer

```solidity
function initiateAppTransfer(IApp app, address newOwner) external appIsActive(app) onlyAdmin(app) {
    address currentOwner = _appConfigs[app].owner;
    _appConfigs[app].pendingOwner = newOwner;
    emit AppTransferInitiated(app, currentOwner, newOwner);
}
```

### acceptAppTransfer

```solidity
function acceptAppTransfer(IApp app) external {
    AppConfig storage config = _appConfigs[app];
    address pendingOwner = config.pendingOwner;

    require(pendingOwner != address(0), NoPendingTransfer());
    require(msg.sender == pendingOwner, InvalidPermissions());

    // Check new owner has available quota
    UserConfig storage newOwnerConfig = _userConfigs[pendingOwner];
    require(newOwnerConfig.activeAppCount < newOwnerConfig.maxActiveApps, MaxActiveAppsExceeded());

    address oldOwner = config.owner;

    // Update app ownership
    config.owner = pendingOwner;
    config.pendingOwner = address(0);

    // Transfer quota: decrement old owner, increment new owner
    _userConfigs[oldOwner].activeAppCount--;
    newOwnerConfig.activeAppCount++;

    // Grant ADMIN_ROLE to new owner so they can manage the app
    _grantRole(_teamRole(pendingOwner, ADMIN_ROLE), pendingOwner);

    emit AppTransferred(app, oldOwner, pendingOwner);
}
```

### cancelAppTransfer

```solidity
function cancelAppTransfer(IApp app) external onlyAdmin(app) {
    require(_appConfigs[app].pendingOwner != address(0), NoPendingTransfer());
    _appConfigs[app].pendingOwner = address(0);
    emit AppTransferCancelled(app);
}
```

## Quota Behavior

- Quota is checked at **accept time**, not initiate time
- Old owner's `activeAppCount` decrements on accept
- New owner's `activeAppCount` increments on accept
- If new owner's quota fills between initiate and accept, accept reverts

## Access Control

- `initiateAppTransfer`: Requires `ADMIN_ROLE` on the app's owner team
- `acceptAppTransfer`: Only callable by the pending owner
- `cancelAppTransfer`: Requires `ADMIN_ROLE` on the app's owner team
- Old owner retains full control until transfer is accepted
- New owner automatically receives `ADMIN_ROLE` on their own team upon acceptance

## Edge Cases

- Calling `initiateAppTransfer` again overwrites the pending owner
- Old owner can still manage app (stop, upgrade, etc.) while transfer is pending
- Pending owner has no access until they accept
- Terminated apps cannot be transferred (`appIsActive` modifier)
