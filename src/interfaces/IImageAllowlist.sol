// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/// @title IImageAllowlist
/// @notice Interface for managing CVM image allowlists using PCR measurements and TCB validation
interface IImageAllowlist {
    enum CVM {
        TDX,
        SEV_SNP
    }

    struct PCR {
        uint8 index;
        bytes32 value;
    }

    struct Image {
        PCR[] pcrs;
        string version;
        string description;
    }

    event ImageAdded(CVM indexed cvm, bytes32 indexed key, string version, string description);
    event ImageRemoved(CVM indexed cvm, bytes32 indexed key);
    event MinimumTCBUpdated(CVM indexed cvm, uint64 tcb);

    error NotSorted();
    error EmptyPCRs();
    error ImageNotFound();

    /**
     * @notice Initialize the ImageAllowlist contract
     * @param initialOwner The initial owner address
     */
    function initialize(address initialOwner) external;

    /**
     * @notice Adds an image to the allowlist
     * @param cvm The CVM type
     * @param image The image to add
     * @return key The key (hash of PCR measurements) for the added image
     */
    function addImage(CVM cvm, Image calldata image) external returns (bytes32 key);

    /**
     * @notice Removes an image from the allowlist
     * @param cvm The CVM type
     * @param key The key (hash of PCR measurements) to remove
     */
    function removeImage(CVM cvm, bytes32 key) external;

    /**
     * @notice Adds multiple images to the allowlist
     * @param cvm The CVM type
     * @param images_ Array of images to add
     * @return keys The keys (hashes of PCR measurements) for the added images
     */
    function addImages(CVM cvm, Image[] calldata images_) external returns (bytes32[] memory keys);

    /**
     * @notice Removes multiple images from the allowlist
     * @param cvm The CVM type
     * @param keys Array of keys (hashes of PCR measurements) to remove
     */
    function removeImages(CVM cvm, bytes32[] calldata keys) external;

    /**
     * @notice Sets the minimum TCB level for a CVM type
     * @param cvm The CVM type
     * @param tcb The minimum TCB level
     */
    function setMinimumTCB(CVM cvm, uint64 tcb) external;

    /**
     * @notice Checks if an image is allowed
     * @param cvm The CVM type
     * @param pcrs The PCR values identifying the image
     * @return Whether the image is allowed
     */
    function isImageAllowed(CVM cvm, PCR[] calldata pcrs) external view returns (bool);

    /**
     * @notice Checks if a TCB level meets the minimum requirement
     * @param cvm The CVM type
     * @param tcb The TCB level to check
     * @return Whether the TCB level is sufficient
     * @dev Returns true for any TCB if minimumTCB is unset (defaults to 0)
     */
    function isTCBValid(CVM cvm, uint64 tcb) external view returns (bool);

    /**
     * @notice Gets the allowed status for an image key
     * @param cvm The CVM type
     * @param key The key (hash of PCR measurements) to check
     * @return Whether the image is allowed
     */
    function images(CVM cvm, bytes32 key) external view returns (bool);

    /**
     * @notice Gets the minimum TCB level for a CVM type
     * @param cvm The CVM type
     * @return The minimum TCB level
     */
    function minimumTCB(CVM cvm) external view returns (uint64);
}
