// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * @title ICallValidator
 * @notice Optional interface that Timelock targets can implement to reject
 *         operations at schedule-time rather than waiting until execute-time.
 */
interface ICallValidator {
    /**
     * @notice Returns true if `caller` is authorized to execute `data` on this contract.
     * @param caller The address that will ultimately call (e.g., the Timelock).
     * @param data   The calldata that will be forwarded to the target.
     */
    function canCall(address caller, bytes calldata data) external view returns (bool);
}
