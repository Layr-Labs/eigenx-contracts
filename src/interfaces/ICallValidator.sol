// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @title ICallValidator
 * @notice Optional interface that Timelock targets can implement to reject
 *         operations at schedule-time rather than waiting until execute-time.
 *
 * @dev Targets MUST opt in via ERC-165 `supportsInterface` so that a reverting
 *      `canCall` is never silently interpreted as "not implemented":
 *
 *          supportsInterface(type(ICallValidator).interfaceId) == true
 *
 *      If a target advertises the interface and `canCall` reverts, the caller
 *      should fail the request — not fall through as if the interface was
 *      absent. See `TimelockControllerImpl._validateTarget` for the enforcement
 *      point.
 */
interface ICallValidator is IERC165 {
    /**
     * @notice Returns true if `caller` is authorized to execute `data` on this contract.
     * @param caller The address that will ultimately call (e.g., the Timelock).
     * @param data   The calldata that will be forwarded to the target.
     */
    function canCall(address caller, bytes calldata data) external view returns (bool);
}
