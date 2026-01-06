// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {ComputeDeployer} from "./utils/ComputeDeployer.sol";
import {IImageAllowlist} from "../src/interfaces/IImageAllowlist.sol";

contract ImageAllowlistTest is ComputeDeployer {
    event ImageUpdated(IImageAllowlist.CVM indexed cvm, bytes32 indexed key, bool allowed);
    event MinimumTCBUpdated(IImageAllowlist.CVM indexed cvm, uint64 tcb);

    function setUp() public {
        _deployContracts();
    }

    function test_initialize_setsOwner() public view {
        assertEq(imageAllowlist.owner(), admin);
    }

    function test_initialize_revertsIfCalledTwice() public {
        vm.expectRevert();
        imageAllowlist.initialize(admin);
    }

    function test_setImage() public {
        IImageAllowlist.PCR[] memory pcrs = _createPCRs();
        bytes32 expectedKey = keccak256(abi.encode(pcrs));

        vm.expectEmit(true, true, true, true);
        emit ImageUpdated(IImageAllowlist.CVM.TDX, expectedKey, true);

        vm.prank(admin);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);

        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, pcrs));
        assertTrue(imageAllowlist.images(IImageAllowlist.CVM.TDX, expectedKey));
    }

    function test_setImage_canToggle() public {
        IImageAllowlist.PCR[] memory pcrs = _createPCRs();

        vm.startPrank(admin);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);
        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, pcrs));

        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, false);
        assertFalse(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, pcrs));
        vm.stopPrank();
    }

    function test_setImage_separatePerCVM() public {
        IImageAllowlist.PCR[] memory pcrs = _createPCRs();

        vm.startPrank(admin);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);
        imageAllowlist.setImage(IImageAllowlist.CVM.SEV_SNP, pcrs, false);
        vm.stopPrank();

        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, pcrs));
        assertFalse(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.SEV_SNP, pcrs));
    }

    function test_setImage_revertsIfNotOwner() public {
        vm.prank(user);
        vm.expectRevert();
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, _createPCRs(), true);
    }

    function test_setImage_revertsIfEmptyPCRs() public {
        vm.prank(admin);
        vm.expectRevert(IImageAllowlist.EmptyPCRs.selector);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, new IImageAllowlist.PCR[](0), true);
    }

    function test_setImage_revertsIfNotSorted() public {
        IImageAllowlist.PCR[] memory pcrs = new IImageAllowlist.PCR[](2);
        pcrs[0] = IImageAllowlist.PCR({index: 2, value: bytes32(uint256(1))});
        pcrs[1] = IImageAllowlist.PCR({index: 1, value: bytes32(uint256(2))});

        vm.prank(admin);
        vm.expectRevert(IImageAllowlist.NotSorted.selector);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);

        // Duplicate indices also fail (not strictly increasing)
        pcrs[0].index = 1;
        vm.prank(admin);
        vm.expectRevert(IImageAllowlist.NotSorted.selector);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);
    }

    function test_setMinimumTCB() public {
        vm.expectEmit(true, true, true, true);
        emit MinimumTCBUpdated(IImageAllowlist.CVM.TDX, 100);

        vm.prank(admin);
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.TDX, 100);

        assertEq(imageAllowlist.minimumTCB(IImageAllowlist.CVM.TDX), 100);
    }

    function test_setMinimumTCB_separatePerCVM() public {
        vm.startPrank(admin);
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.TDX, 100);
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.SEV_SNP, 200);
        vm.stopPrank();

        assertEq(imageAllowlist.minimumTCB(IImageAllowlist.CVM.TDX), 100);
        assertEq(imageAllowlist.minimumTCB(IImageAllowlist.CVM.SEV_SNP), 200);
    }

    function test_setMinimumTCB_revertsIfNotOwner() public {
        vm.prank(user);
        vm.expectRevert();
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.TDX, 100);
    }

    function test_isTCBValid() public {
        vm.prank(admin);
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.TDX, 100);

        assertTrue(imageAllowlist.isTCBValid(IImageAllowlist.CVM.TDX, 101));
        assertTrue(imageAllowlist.isTCBValid(IImageAllowlist.CVM.TDX, 100));
        assertFalse(imageAllowlist.isTCBValid(IImageAllowlist.CVM.TDX, 99));
    }

    function test_isTCBValid_defaultsToZero() public view {
        // minimumTCB defaults to 0, so any TCB >= 0 is valid
        assertTrue(imageAllowlist.isTCBValid(IImageAllowlist.CVM.TDX, 0));
        assertTrue(imageAllowlist.isTCBValid(IImageAllowlist.CVM.TDX, 1));
    }

    function _createPCRs() internal pure returns (IImageAllowlist.PCR[] memory) {
        IImageAllowlist.PCR[] memory pcrs = new IImageAllowlist.PCR[](3);
        pcrs[0] = IImageAllowlist.PCR({index: 0, value: bytes32(uint256(1))});
        pcrs[1] = IImageAllowlist.PCR({index: 1, value: bytes32(uint256(2))});
        pcrs[2] = IImageAllowlist.PCR({index: 2, value: bytes32(uint256(3))});
        return pcrs;
    }
}
