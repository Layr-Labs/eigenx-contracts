// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

// Minimal interface for Gnosis SafeProxyFactory. Hand-rolled to avoid
// depending on safe-global/safe-contracts for two function selectors.
//
// Signature source: SafeProxyFactory v1.3.0 through v1.4.1 — stable across
// the factory versions currently deployed on our target chains.
// `createProxyWithNonce` is what we call; `proxyCreationCode` is used to
// precompute deterministic Safe addresses via CREATE2.
interface ISafeProxyFactory {
    function createProxyWithNonce(address _singleton, bytes memory initializer, uint256 saltNonce)
        external
        returns (address proxy);

    function proxyCreationCode() external pure returns (bytes memory);
}
