// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

// Minimal interface for the Gnosis Safe singleton. Hand-rolled to avoid
// pulling the full safe-global/safe-contracts dependency for two function
// selectors.
//
// Signature source: Safe singleton `setup(address[],uint256,address,bytes,
// address,address,uint256,address)`. Stable across Safe v1.3.0 / v1.3.0-l2 /
// v1.4.1 / v1.4.1-l2 — the versions currently deployed on the chains we
// target (mainnet, Sepolia, OP-stack L2s). If Safe ships a breaking change
// in a future major version, update this file and pin the deployed
// singleton in zeus env accordingly.
interface ISafe {
    function setup(
        address[] calldata _owners,
        uint256 _threshold,
        address to,
        bytes calldata data,
        address fallbackHandler,
        address paymentToken,
        uint256 payment,
        address payable paymentReceiver
    ) external;
}
