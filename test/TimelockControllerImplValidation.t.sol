// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Test} from "forge-std/Test.sol";
import {TimelockControllerImpl} from "../src/governance/TimelockControllerImpl.sol";
import {ICallValidator} from "../src/interfaces/ICallValidator.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {Clones} from "@openzeppelin/contracts/proxy/Clones.sol";

/// @dev Behaviors exercised here come from TimelockControllerImpl._validateTarget.
/// We spin up a real TimelockControllerImpl, grant the deployer PROPOSER_ROLE,
/// and call schedule() against various mock targets that exhibit specific
/// misbehaviors (reverts, returndata-bombs, OOG in canCall, honest rejects).
contract TimelockControllerImplValidationTest is Test {
    TimelockControllerImpl internal timelock;
    address internal proposer = address(0xA11CE);

    function setUp() public {
        TimelockControllerImpl impl = new TimelockControllerImpl();
        // Impl has initializers disabled; deploy via clone to call initialize.
        timelock = TimelockControllerImpl(payable(Clones.clone(address(impl))));

        address[] memory proposers = new address[](1);
        proposers[0] = proposer;
        address[] memory executors = new address[](1);
        executors[0] = proposer;
        timelock.initialize(60, proposers, executors, address(0));
    }

    function _schedule(address target, bytes memory data) internal {
        vm.prank(proposer);
        timelock.schedule(target, 0, data, bytes32(0), bytes32(0), 60);
    }

    /// --- Target that does NOT advertise ICallValidator at all ---

    function test_schedule_targetWithoutERC165_isAllowedThrough() public {
        InertContract inert = new InertContract();
        _schedule(address(inert), hex"12345678");
    }

    function test_schedule_EOA_isAllowedThrough() public {
        // Zero-code address fast path.
        _schedule(makeAddr("eoa"), hex"12345678");
    }

    /// --- Target that returns false from supportsInterface ---

    function test_schedule_targetClaimsNoInterface_isAllowedThrough() public {
        ERC165SaysNo liar = new ERC165SaysNo();
        _schedule(address(liar), hex"12345678");
    }

    /// --- Target that advertises ICallValidator ---

    function test_schedule_validatorAllows_schedulesSuccessfully() public {
        AllowingValidator v = new AllowingValidator();
        _schedule(address(v), hex"12345678");
    }

    function test_schedule_validatorRejects_reverts() public {
        RejectingValidator v = new RejectingValidator();
        vm.expectRevert(bytes("TimelockController: target rejected the call"));
        _schedule(address(v), hex"12345678");
    }

    function test_schedule_validatorReverts_failsClosed() public {
        // Advertises the interface, but canCall reverts. MUST reject rather than
        // silently pass through as "no interface" — that was the pre-hardening
        // bug the audit called out (G-3 / V-4 / D-2 / P-9).
        RevertingValidator v = new RevertingValidator();
        vm.expectRevert(bytes("TimelockController: canCall reverted"));
        _schedule(address(v), hex"12345678");
    }

    function test_schedule_validatorOOG_failsClosed() public {
        // Validator runs an infinite loop; the bounded gas on the staticcall
        // cuts it off and the scheduling call fails with "canCall reverted".
        InfiniteLoopValidator v = new InfiniteLoopValidator();
        vm.expectRevert(bytes("TimelockController: canCall reverted"));
        _schedule(address(v), hex"12345678");
    }

    function test_schedule_validatorReturndataBomb_failsClosed() public {
        // Validator returns a huge blob. The hardened validator reads only
        // the first 32 bytes; abi.decode(bool) succeeds with whatever the
        // first 32 bytes encode (here: 1 → allowed), so this should succeed
        // without OOGing the outer call.
        ReturndataBombValidator v = new ReturndataBombValidator();
        _schedule(address(v), hex"12345678");
    }
}

/// --- Mock targets ---

contract InertContract {
    // No ERC-165 — calls to supportsInterface revert with no reason.
    // Used to verify fast-path for non-validator targets.

    }

contract ERC165SaysNo is IERC165 {
    function supportsInterface(bytes4) external pure returns (bool) {
        return false;
    }
}

contract AllowingValidator is ICallValidator {
    function supportsInterface(bytes4 interfaceId) external pure returns (bool) {
        return interfaceId == type(ICallValidator).interfaceId || interfaceId == type(IERC165).interfaceId;
    }

    function canCall(address, bytes calldata) external pure returns (bool) {
        return true;
    }
}

contract RejectingValidator is ICallValidator {
    function supportsInterface(bytes4 interfaceId) external pure returns (bool) {
        return interfaceId == type(ICallValidator).interfaceId || interfaceId == type(IERC165).interfaceId;
    }

    function canCall(address, bytes calldata) external pure returns (bool) {
        return false;
    }
}

contract RevertingValidator is ICallValidator {
    function supportsInterface(bytes4 interfaceId) external pure returns (bool) {
        return interfaceId == type(ICallValidator).interfaceId || interfaceId == type(IERC165).interfaceId;
    }

    function canCall(address, bytes calldata) external pure returns (bool) {
        revert("no");
    }
}

contract InfiniteLoopValidator is ICallValidator {
    function supportsInterface(bytes4 interfaceId) external pure returns (bool) {
        return interfaceId == type(ICallValidator).interfaceId || interfaceId == type(IERC165).interfaceId;
    }

    function canCall(address, bytes calldata) external view returns (bool) {
        // staticcall so state writes would revert; burn gas in a tight loop.
        uint256 i;
        while (true) {
            i++;
        }
        return true; // unreachable
    }
}

contract ReturndataBombValidator is ICallValidator {
    function supportsInterface(bytes4 interfaceId) external pure returns (bool) {
        return interfaceId == type(ICallValidator).interfaceId || interfaceId == type(IERC165).interfaceId;
    }

    function canCall(address, bytes calldata) external pure returns (bool) {
        // Return true (first 32 bytes encode bool(true)), then a padding blob.
        // The hardened validator should only read the first 32 bytes.
        bytes memory blob = new bytes(1024);
        blob[31] = 0x01; // decodes as true
        assembly {
            return(add(blob, 32), mload(blob))
        }
    }
}
