// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {ImageAllowlistStorage} from "./storage/ImageAllowlistStorage.sol";
import {IImageAllowlist} from "./interfaces/IImageAllowlist.sol";

/// @title ImageAllowlist
/// @notice Allowlist for CVM images using PCR measurements and minimum TCB versions
/// @dev Separate allowlists are maintained per CVM type (TDX and SEV-SNP)
contract ImageAllowlist is Initializable, OwnableUpgradeable, ImageAllowlistStorage {
    constructor() {
        _disableInitializers();
    }

    /// @inheritdoc IImageAllowlist
    function initialize(address initialOwner) external initializer {
        _transferOwnership(initialOwner);
    }

    /// @inheritdoc IImageAllowlist
    function addImage(CVM cvm, Image calldata image) external onlyOwner returns (bytes32) {
        return _addImage(cvm, image);
    }

    /// @inheritdoc IImageAllowlist
    function removeImage(CVM cvm, bytes32 key) external onlyOwner {
        _removeImage(cvm, key);
    }

    /// @inheritdoc IImageAllowlist
    function addImages(CVM cvm, Image[] calldata images_) external onlyOwner returns (bytes32[] memory keys) {
        keys = new bytes32[](images_.length);
        for (uint256 i = 0; i < images_.length; i++) {
            keys[i] = _addImage(cvm, images_[i]);
        }
    }

    /// @inheritdoc IImageAllowlist
    function removeImages(CVM cvm, bytes32[] calldata keys) external onlyOwner {
        for (uint256 i = 0; i < keys.length; i++) {
            _removeImage(cvm, keys[i]);
        }
    }

    /// @inheritdoc IImageAllowlist
    function setMinimumTCB(CVM cvm, uint64 tcb) external onlyOwner {
        minimumTCB[cvm] = tcb;
        emit MinimumTCBUpdated(cvm, tcb);
    }

    /// @inheritdoc IImageAllowlist
    function isImageAllowed(CVM cvm, PCR[] calldata pcrs) external view returns (bool) {
        return images[cvm][_hashPCRs(pcrs)];
    }

    /// @inheritdoc IImageAllowlist
    function isTCBValid(CVM cvm, uint64 tcb) external view returns (bool) {
        return tcb >= minimumTCB[cvm];
    }

    /**
     * @notice Adds an image to the allowlist for a CVM type
     * @param cvm The CVM type to add the image for
     * @param image The image to add
     * @return key The key (hash of PCR measurements) for the added image
     */
    function _addImage(CVM cvm, Image calldata image) private returns (bytes32 key) {
        require(image.pcrs.length != 0, EmptyPCRs());
        require(_isSorted(image.pcrs), NotSorted());

        key = _hashPCRs(image.pcrs);
        images[cvm][key] = true;
        emit ImageAdded(cvm, key, image.version, image.description);
    }

    /**
     * @notice Removes an image from the allowlist for a CVM type
     * @param cvm The CVM type to remove the image from
     * @param key The key (hash of PCR measurements) identifying the image
     */
    function _removeImage(CVM cvm, bytes32 key) private {
        require(images[cvm][key], ImageNotFound());
        images[cvm][key] = false;
        emit ImageRemoved(cvm, key);
    }

    /**
     * @notice Hashes an array of PCR values
     * @param pcrs The array of PCR values to hash
     * @return The keccak256 hash of the encoded PCR array
     */
    function _hashPCRs(PCR[] calldata pcrs) private pure returns (bytes32) {
        return keccak256(abi.encode(pcrs));
    }

    /**
     * @notice Checks if an array of PCR values is sorted by index
     * @param pcrs The array of PCR values to check
     * @return True if the array is sorted, false otherwise
     */
    function _isSorted(PCR[] calldata pcrs) private pure returns (bool) {
        for (uint256 i = 1; i < pcrs.length; i++) {
            if (pcrs[i].index <= pcrs[i - 1].index) return false;
        }
        return true;
    }
}
