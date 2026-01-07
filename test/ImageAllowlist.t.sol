// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {ComputeDeployer} from "./utils/ComputeDeployer.sol";
import {IImageAllowlist} from "../src/interfaces/IImageAllowlist.sol";

contract ImageAllowlistTest is ComputeDeployer {
    event ImageAdded(IImageAllowlist.CVM indexed cvm, bytes32 indexed key, string version, string description);
    event ImageRemoved(IImageAllowlist.CVM indexed cvm, bytes32 indexed key);
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

    function test_addImage() public {
        IImageAllowlist.Image memory image =
            IImageAllowlist.Image({pcrs: _createPCRs(), version: "v1.0.0", description: "test-image"});
        bytes32 expectedKey = keccak256(abi.encode(image.pcrs));

        vm.expectEmit(true, true, true, true);
        emit ImageAdded(IImageAllowlist.CVM.TDX, expectedKey, "v1.0.0", "test-image");

        vm.prank(admin);
        imageAllowlist.addImage(IImageAllowlist.CVM.TDX, image);

        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, image.pcrs));
        assertTrue(imageAllowlist.images(IImageAllowlist.CVM.TDX, expectedKey));
    }

    function test_removeImage() public {
        IImageAllowlist.Image memory image =
            IImageAllowlist.Image({pcrs: _createPCRs(), version: "v1.0.0", description: "test-image"});
        bytes32 expectedKey = keccak256(abi.encode(image.pcrs));

        vm.startPrank(admin);
        imageAllowlist.addImage(IImageAllowlist.CVM.TDX, image);
        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, image.pcrs));

        vm.expectEmit(true, true, true, true);
        emit ImageRemoved(IImageAllowlist.CVM.TDX, expectedKey);

        imageAllowlist.removeImage(IImageAllowlist.CVM.TDX, expectedKey);
        assertFalse(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, image.pcrs));
        vm.stopPrank();
    }

    function test_addImage_separatePerCVM() public {
        IImageAllowlist.Image memory image =
            IImageAllowlist.Image({pcrs: _createPCRs(), version: "v1.0.0", description: "tdx-image"});

        vm.startPrank(admin);
        imageAllowlist.addImage(IImageAllowlist.CVM.TDX, image);
        vm.stopPrank();

        assertTrue(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.TDX, image.pcrs));
        assertFalse(imageAllowlist.isImageAllowed(IImageAllowlist.CVM.SEV_SNP, image.pcrs));
    }

    function test_addImage_revertsIfNotOwner() public {
        IImageAllowlist.Image memory image =
            IImageAllowlist.Image({pcrs: _createPCRs(), version: "v1.0.0", description: "test-image"});

        vm.prank(user);
        vm.expectRevert();
        imageAllowlist.addImage(IImageAllowlist.CVM.TDX, image);
    }

    function test_removeImage_revertsIfNotOwner() public {
        vm.prank(user);
        vm.expectRevert();
        imageAllowlist.removeImage(IImageAllowlist.CVM.TDX, bytes32(uint256(1)));
    }

    function test_removeImage_revertsIfNotFound() public {
        vm.prank(admin);
        vm.expectRevert(IImageAllowlist.ImageNotFound.selector);
        imageAllowlist.removeImage(IImageAllowlist.CVM.TDX, bytes32(uint256(1)));
    }

    function test_addImage_revertsIfEmptyPCRs() public {
        IImageAllowlist.Image memory image =
            IImageAllowlist.Image({pcrs: new IImageAllowlist.PCR[](0), version: "v1.0.0", description: "test-image"});

        vm.prank(admin);
        vm.expectRevert(IImageAllowlist.EmptyPCRs.selector);
        imageAllowlist.addImage(IImageAllowlist.CVM.TDX, image);
    }

    function test_addImage_revertsIfNotSorted() public {
        IImageAllowlist.PCR[] memory pcrs = new IImageAllowlist.PCR[](2);
        pcrs[0] = IImageAllowlist.PCR({index: 2, value: bytes32(uint256(1))});
        pcrs[1] = IImageAllowlist.PCR({index: 1, value: bytes32(uint256(2))});

        IImageAllowlist.Image memory image =
            IImageAllowlist.Image({pcrs: pcrs, version: "v1.0.0", description: "test-image"});

        vm.prank(admin);
        vm.expectRevert(IImageAllowlist.NotSorted.selector);
        imageAllowlist.addImage(IImageAllowlist.CVM.TDX, image);

        // Duplicate indices also fail (not strictly increasing)
        pcrs[0].index = 1;
        vm.prank(admin);
        vm.expectRevert(IImageAllowlist.NotSorted.selector);
        imageAllowlist.addImage(IImageAllowlist.CVM.TDX, image);
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

    function test_addImages() public {
        IImageAllowlist.Image[] memory inputs = new IImageAllowlist.Image[](2);
        inputs[0] = IImageAllowlist.Image({pcrs: _createPCRs(), version: "v1.0.0", description: "image-1"});
        inputs[1] = IImageAllowlist.Image({pcrs: _createPCRs2(), version: "v2.0.0", description: "image-2"});

        bytes32 key1 = keccak256(abi.encode(inputs[0].pcrs));
        bytes32 key2 = keccak256(abi.encode(inputs[1].pcrs));

        vm.prank(admin);
        imageAllowlist.addImages(IImageAllowlist.CVM.TDX, inputs);

        assertTrue(imageAllowlist.images(IImageAllowlist.CVM.TDX, key1));
        assertTrue(imageAllowlist.images(IImageAllowlist.CVM.TDX, key2));
    }

    function test_removeImages() public {
        IImageAllowlist.Image[] memory inputs = new IImageAllowlist.Image[](2);
        inputs[0] = IImageAllowlist.Image({pcrs: _createPCRs(), version: "v1.0.0", description: "image-1"});
        inputs[1] = IImageAllowlist.Image({pcrs: _createPCRs2(), version: "v2.0.0", description: "image-2"});

        bytes32 key1 = keccak256(abi.encode(inputs[0].pcrs));
        bytes32 key2 = keccak256(abi.encode(inputs[1].pcrs));

        vm.startPrank(admin);
        imageAllowlist.addImages(IImageAllowlist.CVM.TDX, inputs);

        assertTrue(imageAllowlist.images(IImageAllowlist.CVM.TDX, key1));
        assertTrue(imageAllowlist.images(IImageAllowlist.CVM.TDX, key2));

        bytes32[] memory keys = new bytes32[](2);
        keys[0] = key1;
        keys[1] = key2;

        imageAllowlist.removeImages(IImageAllowlist.CVM.TDX, keys);

        assertFalse(imageAllowlist.images(IImageAllowlist.CVM.TDX, key1));
        assertFalse(imageAllowlist.images(IImageAllowlist.CVM.TDX, key2));
        vm.stopPrank();
    }

    function _createPCRs() internal pure returns (IImageAllowlist.PCR[] memory) {
        IImageAllowlist.PCR[] memory pcrs = new IImageAllowlist.PCR[](3);
        pcrs[0] = IImageAllowlist.PCR({index: 0, value: bytes32(uint256(1))});
        pcrs[1] = IImageAllowlist.PCR({index: 1, value: bytes32(uint256(2))});
        pcrs[2] = IImageAllowlist.PCR({index: 2, value: bytes32(uint256(3))});
        return pcrs;
    }

    function _createPCRs2() internal pure returns (IImageAllowlist.PCR[] memory) {
        IImageAllowlist.PCR[] memory pcrs = new IImageAllowlist.PCR[](2);
        pcrs[0] = IImageAllowlist.PCR({index: 0, value: bytes32(uint256(10))});
        pcrs[1] = IImageAllowlist.PCR({index: 1, value: bytes32(uint256(20))});
        return pcrs;
    }
}
