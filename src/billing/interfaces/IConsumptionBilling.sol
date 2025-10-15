// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

interface IConsumptionBilling {
    struct APIKey {
        uint128 pendingUnits;   // Units consumed this period
        bool isActive;
        address owner;          // Who controls the key
        address account;        // Billing account that pays for the key
    }

    // Errors
    error NotOracle();
    error KeyAlreadyExists(bytes32 keyId);
    error NotOwner();
    error KeyAlreadyInactive(bytes32 keyId);
    error InvalidAccount();
    error KeyDoesNotExist(bytes32 keyId);
    error KeyInactive(bytes32 keyId);
    error LengthMismatch();

    // Events
    event KeyCreated(bytes32 indexed keyId, address indexed owner);
    event UsageRecorded(bytes32 indexed keyId, uint128 units);
    event KeySettled(bytes32 indexed keyId, uint128 amount, uint128 unitsPaid);
    event KeyDeactivated(bytes32 indexed keyId);
    event AccountChanged(bytes32 indexed keyId, address indexed oldAccount, address indexed newAccount);
    event OracleChanged(address indexed oldOracle, address indexed newOracle);
}
