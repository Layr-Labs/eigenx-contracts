// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Test} from "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ImageAllowlist} from "../src/ImageAllowlist.sol";
import {IImageAllowlist} from "../src/interfaces/IImageAllowlist.sol";

contract ImageAllowlistTest is Test {
    ImageAllowlist public imageAllowlist;

    address public owner = makeAddr("owner");
    address public notOwner = makeAddr("notOwner");
    address public proxyAdmin = makeAddr("proxyAdmin");

    event ImageUpdated(IImageAllowlist.CVM indexed cvm, bytes32 indexed key, bool allowed);
    event MinimumTCBUpdated(IImageAllowlist.CVM indexed cvm, uint64 tcb);

    function setUp() public {
        ImageAllowlist implementation = new ImageAllowlist();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation), proxyAdmin, abi.encodeWithSelector(ImageAllowlist.initialize.selector, owner)
        );
        imageAllowlist = ImageAllowlist(address(proxy));
    }

    /// INITIALIZATION TESTS

    function test_initialize_setsOwner() public view {
        assertEq(imageAllowlist.owner(), owner);
    }

    function test_initialize_revertsIfCalledTwice() public {
        vm.expectRevert();
        imageAllowlist.initialize(owner);
    }

    /// SET IMAGE TESTS

    function test_setImage() public {
        IImageAllowlist.PCR[] memory pcrs = _createPCRs();
        bytes32 expectedKey = keccak256(abi.encode(pcrs));

        vm.expectEmit(true, true, true, true);
        emit ImageUpdated(IImageAllowlist.CVM.TDX, expectedKey, true);

        vm.prank(owner);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);

        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, pcrs));
        assertTrue(imageAllowlist.images(IImageAllowlist.CVM.TDX, expectedKey));
    }

    function test_setImage_canToggle() public {
        IImageAllowlist.PCR[] memory pcrs = _createPCRs();

        vm.startPrank(owner);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);
        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, pcrs));

        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, false);
        assertFalse(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, pcrs));
        vm.stopPrank();
    }

    function test_setImage_separatePerCVM() public {
        IImageAllowlist.PCR[] memory pcrs = _createPCRs();

        vm.startPrank(owner);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);
        imageAllowlist.setImage(IImageAllowlist.CVM.SEV_SNP, pcrs, false);
        vm.stopPrank();

        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, pcrs));
        assertFalse(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.SEV_SNP, pcrs));
    }

    function test_setImage_revertsIfNotOwner() public {
        vm.prank(notOwner);
        vm.expectRevert();
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, _createPCRs(), true);
    }

    function test_setImage_revertsIfEmptyPCRs() public {
        vm.prank(owner);
        vm.expectRevert(IImageAllowlist.EmptyPCRs.selector);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, new IImageAllowlist.PCR[](0), true);
    }

    function test_setImage_revertsIfNotSorted() public {
        IImageAllowlist.PCR[] memory pcrs = new IImageAllowlist.PCR[](2);
        pcrs[0] = IImageAllowlist.PCR({index: 2, value: bytes32(uint256(1))});
        pcrs[1] = IImageAllowlist.PCR({index: 1, value: bytes32(uint256(2))});

        vm.prank(owner);
        vm.expectRevert(IImageAllowlist.NotSorted.selector);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);

        // Duplicate indices also fail (not strictly increasing)
        pcrs[0].index = 1;
        vm.prank(owner);
        vm.expectRevert(IImageAllowlist.NotSorted.selector);
        imageAllowlist.setImage(IImageAllowlist.CVM.TDX, pcrs, true);
    }

    /// SET MINIMUM TCB TESTS

    function test_setMinimumTCB() public {
        vm.expectEmit(true, true, true, true);
        emit MinimumTCBUpdated(IImageAllowlist.CVM.TDX, 100);

        vm.prank(owner);
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.TDX, 100);

        assertEq(imageAllowlist.minimumTCB(IImageAllowlist.CVM.TDX), 100);
    }

    function test_setMinimumTCB_separatePerCVM() public {
        vm.startPrank(owner);
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.TDX, 100);
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.SEV_SNP, 200);
        vm.stopPrank();

        assertEq(imageAllowlist.minimumTCB(IImageAllowlist.CVM.TDX), 100);
        assertEq(imageAllowlist.minimumTCB(IImageAllowlist.CVM.SEV_SNP), 200);
    }

    function test_setMinimumTCB_revertsIfNotOwner() public {
        vm.prank(notOwner);
        vm.expectRevert();
        imageAllowlist.setMinimumTCB(IImageAllowlist.CVM.TDX, 100);
    }

    /// IS TCB VALID TESTS

    function test_isTCBValid() public {
        vm.prank(owner);
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

    /// HELPER FUNCTIONS

    function _createPCRs() internal pure returns (IImageAllowlist.PCR[] memory) {
        IImageAllowlist.PCR[] memory pcrs = new IImageAllowlist.PCR[](3);
        pcrs[0] = IImageAllowlist.PCR({index: 0, value: bytes32(uint256(1))});
        pcrs[1] = IImageAllowlist.PCR({index: 1, value: bytes32(uint256(2))});
        pcrs[2] = IImageAllowlist.PCR({index: 2, value: bytes32(uint256(3))});
        return pcrs;
    }
}
