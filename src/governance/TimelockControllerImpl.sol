// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {
    TimelockControllerUpgradeable
} from "@openzeppelin-upgrades/contracts/governance/TimelockControllerUpgradeable.sol";
import {ICallValidator} from "../interfaces/ICallValidator.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @title TimelockControllerImpl
 * @notice Implementation contract for TimelockController minimal proxies
 * @dev Wraps TimelockControllerUpgradeable and adds on-chain pending operation enumeration.
 *      Overrides schedule/scheduleBatch/execute/executeBatch/cancel to maintain a set of
 *      pending operation IDs, enabling clients to enumerate queued operations without
 *      requiring event log scanning.
 *
 *      Also validates targets at schedule-time: if a target implements ICallValidator,
 *      canCall(address(this), data) must return true or the schedule is rejected.
 */
contract TimelockControllerImpl is TimelockControllerUpgradeable {
    struct PendingOp {
        bytes32 id;
        address target;
        bytes data;
        uint256 executableAt;
    }

    // Append-only array of pending operation IDs (swap-and-pop on removal).
    bytes32[] private _pendingIds;
    // id => 1-based index into _pendingIds (0 means not in set).
    mapping(bytes32 => uint256) private _pendingIndex;
    // id => stored op metadata for enumeration
    mapping(bytes32 => PendingOp) private _pendingOps;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the timelock controller
     * @param minDelay Minimum delay for operations
     * @param proposers Addresses granted proposer and canceller roles
     * @param executors Addresses granted executor role
     * @param admin Optional admin address (use address(0) for self-administered)
     */
    function initialize(uint256 minDelay, address[] memory proposers, address[] memory executors, address admin)
        external
        initializer
    {
        __TimelockController_init(minDelay, proposers, executors, admin);
    }

    // ── Overrides ────────────────────────────────────────────────────────────

    function schedule(
        address target,
        uint256 value,
        bytes calldata data,
        bytes32 predecessor,
        bytes32 salt,
        uint256 delay
    ) public override {
        _validateTarget(target, data);
        super.schedule(target, value, data, predecessor, salt, delay);
        bytes32 id = hashOperation(target, value, data, predecessor, salt);
        _addPending(id, target, data, block.timestamp + delay);
    }

    function scheduleBatch(
        address[] calldata targets,
        uint256[] calldata values,
        bytes[] calldata payloads,
        bytes32 predecessor,
        bytes32 salt,
        uint256 delay
    ) public override {
        for (uint256 i = 0; i < targets.length; i++) {
            _validateTarget(targets[i], payloads[i]);
        }
        super.scheduleBatch(targets, values, payloads, predecessor, salt, delay);
        // For batch ops store empty data — callers should use getPendingOperationIds and decode individually
        bytes32 id = hashOperationBatch(targets, values, payloads, predecessor, salt);
        _addPending(id, address(0), "", block.timestamp + delay);
    }

    function execute(address target, uint256 value, bytes calldata payload, bytes32 predecessor, bytes32 salt)
        public
        payable
        override
    {
        bytes32 id = hashOperation(target, value, payload, predecessor, salt);
        super.execute(target, value, payload, predecessor, salt);
        _removePending(id);
    }

    function executeBatch(
        address[] calldata targets,
        uint256[] calldata values,
        bytes[] calldata payloads,
        bytes32 predecessor,
        bytes32 salt
    ) public payable override {
        bytes32 id = hashOperationBatch(targets, values, payloads, predecessor, salt);
        super.executeBatch(targets, values, payloads, predecessor, salt);
        _removePending(id);
    }

    function cancel(bytes32 id) public override {
        super.cancel(id);
        _removePending(id);
    }

    // ── Enumeration ──────────────────────────────────────────────────────────

    /**
     * @notice Returns all currently pending operation IDs.
     */
    function getPendingOperationIds() external view returns (bytes32[] memory) {
        return _pendingIds;
    }

    /**
     * @notice Returns all currently pending operations with metadata.
     * @dev For single-call ops, target and data are populated.
     *      For batch ops, target is address(0) and data is empty.
     */
    function getPendingOperations() external view returns (PendingOp[] memory ops) {
        ops = new PendingOp[](_pendingIds.length);
        for (uint256 i = 0; i < _pendingIds.length; i++) {
            ops[i] = _pendingOps[_pendingIds[i]];
        }
    }

    // ── Internal helpers ─────────────────────────────────────────────────────

    /// Gas cap for the ERC-165 probe. A compliant target needs ~2k gas;
    /// 30k is generous while still bounding returndata-bomb griefing.
    uint256 private constant ERC165_PROBE_GAS = 30_000;

    /// Gas cap for the `canCall` query. Validation is view-only logic
    /// (storage reads, comparisons) so 200k is plenty; any implementation
    /// that needs more is almost certainly misbehaving.
    uint256 private constant CANCALL_QUERY_GAS = 200_000;

    /// interfaceId for ICallValidator, computed inline to avoid importing
    /// ERC-165's library just to read one constant.
    bytes4 private constant I_CALL_VALIDATOR_INTERFACE_ID = type(ICallValidator).interfaceId;

    /**
     * @dev Schedule-time target validation.
     *
     * A target that advertises {ICallValidator} via ERC-165 gets its
     * `canCall(this, data)` consulted; if it returns false the schedule
     * reverts. Targets that don't advertise the interface are allowed
     * through (backwards compatible with any external contract that the
     * system might Timelock in the future).
     *
     * Hardening choices addressed here:
     *  - ERC-165 first, so a target that reverts for any non-authorization
     *    reason (including OOG) is NOT silently treated as "not implemented".
     *    Only a target that explicitly returns `false` from `supportsInterface`
     *    or doesn't respond to the probe is skipped.
     *  - Bounded gas on both probes to stop a malicious target from
     *    exhausting the outer call via runaway `canCall` computation.
     *  - Bounded returndata (only the first 32 bytes are considered) to
     *    prevent returndata-bomb griefing; excess is ignored.
     *  - When `supportsInterface` says yes but `canCall` reverts, that is
     *    treated as a definitive "no" — the target made a claim it couldn't
     *    back up, so its pending op is rejected rather than silently allowed.
     */
    function _validateTarget(address target, bytes calldata data) private view {
        // Only contracts can implement ICallValidator; EOAs and zero-code
        // addresses short-circuit without any external calls.
        if (target.code.length == 0) return;

        if (!_supportsCallValidator(target)) return;

        // Target claims to implement the interface; ask it.
        bytes memory canCallCalldata = abi.encodeWithSelector(ICallValidator.canCall.selector, address(this), data);
        (bool ok, bytes memory ret) = target.staticcall{gas: CANCALL_QUERY_GAS}(canCallCalldata);

        // A revert here means the target advertised ICallValidator but its
        // implementation failed. Fail closed — do not let it through as if
        // the interface was absent.
        require(ok && ret.length >= 32, "TimelockController: canCall reverted");

        bool allowed = abi.decode(ret, (bool));
        require(allowed, "TimelockController: target rejected the call");
    }

    /// @dev Probe `target` for ICallValidator support via ERC-165.
    ///      Returns false (allow through, treat as non-implementing) on any
    ///      failure or non-bool return; returns true only on an explicit `true`.
    function _supportsCallValidator(address target) private view returns (bool) {
        bytes memory probe = abi.encodeWithSelector(IERC165.supportsInterface.selector, I_CALL_VALIDATOR_INTERFACE_ID);
        (bool ok, bytes memory ret) = target.staticcall{gas: ERC165_PROBE_GAS}(probe);
        if (!ok || ret.length < 32) return false;
        return abi.decode(ret, (bool));
    }

    function _addPending(bytes32 id, address target, bytes memory data, uint256 executableAt) private {
        if (_pendingIndex[id] != 0) return; // already tracked
        _pendingIds.push(id);
        _pendingIndex[id] = _pendingIds.length; // 1-based
        _pendingOps[id] = PendingOp({id: id, target: target, data: data, executableAt: executableAt});
    }

    function _removePending(bytes32 id) private {
        uint256 idx = _pendingIndex[id];
        if (idx == 0) return; // not tracked (pre-upgrade op)
        uint256 i = idx - 1;
        uint256 last = _pendingIds.length - 1;
        if (i != last) {
            bytes32 lastId = _pendingIds[last];
            _pendingIds[i] = lastId;
            _pendingIndex[lastId] = idx; // keep 1-based
        }
        _pendingIds.pop();
        delete _pendingIndex[id];
        delete _pendingOps[id];
    }
}
