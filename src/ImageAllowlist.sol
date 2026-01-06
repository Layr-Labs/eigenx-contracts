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
    function setImage(CVM cvm, PCR[] calldata pcrs, bool isAllowed) external onlyOwner {
        require(pcrs.length != 0, EmptyPCRs());
        require(_isSorted(pcrs), NotSorted());

        bytes32 key = _hashPCRs(pcrs);
        images[cvm][key] = isAllowed;
        emit ImageUpdated(cvm, key, isAllowed);
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
