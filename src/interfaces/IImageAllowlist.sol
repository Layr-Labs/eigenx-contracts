// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/// @title IImageAllowlist
/// @notice Interface for managing image allowlists using PCR measurements and TCB validation
interface IImageAllowlist {
    enum Platform {
        INTEL_TDX,
        AMD_SEV_SNP,
        GOOGLE_VTPM
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

    event ImageAdded(Platform indexed platform, bytes32 indexed key, string version, string description);
    event ImageRemoved(Platform indexed platform, bytes32 indexed key);
    event MinimumTCBUpdated(Platform indexed platform, uint64 tcb);

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
     * @param platform The Platform type
     * @param image The image to add
     * @return key The key (hash of PCR measurements) for the added image
     */
    function addImage(Platform platform, Image calldata image) external returns (bytes32 key);

    /**
     * @notice Removes an image from the allowlist
     * @param platform The Platform type
     * @param key The key (hash of PCR measurements) to remove
     */
    function removeImage(Platform platform, bytes32 key) external;

    /**
     * @notice Adds multiple images to the allowlist
     * @param platform The Platform type
     * @param images_ Array of images to add
     * @return keys The keys (hashes of PCR measurements) for the added images
     */
    function addImages(Platform platform, Image[] calldata images_) external returns (bytes32[] memory keys);

    /**
     * @notice Removes multiple images from the allowlist
     * @param platform The Platform type
     * @param keys Array of keys (hashes of PCR measurements) to remove
     */
    function removeImages(Platform platform, bytes32[] calldata keys) external;

    /**
     * @notice Sets the minimum TCB level for a Platform type
     * @param platform The Platform type
     * @param tcb The minimum TCB level
     */
    function setMinimumTCB(Platform platform, uint64 tcb) external;

    /**
     * @notice Checks if an image is allowed
     * @param platform The Platform type
     * @param pcrs The PCR values identifying the image
     * @return Whether the image is allowed
     */
    function isImageAllowed(Platform platform, PCR[] calldata pcrs) external view returns (bool);

    /**
     * @notice Checks if a TCB level meets the minimum requirement
     * @param platform The Platform type
     * @param tcb The TCB level to check
     * @return Whether the TCB level is sufficient
     * @dev Returns true for any TCB if minimumTCB is unset (defaults to 0)
     */
    function isTCBValid(Platform platform, uint64 tcb) external view returns (bool);

    /**
     * @notice Gets the allowed status for an image key
     * @param platform The Platform type
     * @param key The key (hash of PCR measurements) to check
     * @return Whether the image is allowed
     */
    function images(Platform platform, bytes32 key) external view returns (bool);

    /**
     * @notice Gets the minimum TCB level for a Platform type
     * @param platform The Platform type
     * @return The minimum TCB level
     */
    function minimumTCB(Platform platform) external view returns (uint64);
}
