// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

interface IImageAllowlist {
    enum CVM {
        TDX,
        SEV_SNP
    }

    struct PCR {
        uint8 index;
        bytes32 value;
    }

    event ImageUpdated(CVM indexed cvm, bytes32 indexed key, bool allowed);
    event MinimumTCBUpdated(CVM indexed cvm, uint64 tcb);

    error NotSorted();
    error EmptyPCRs();

    /**
     * @notice Initialize the ImageAllowlist contract
     * @param initialOwner The initial owner address
     */
    function initialize(address initialOwner) external;

    /**
     * @notice Sets an image's allowlist status
     * @param cvm The CVM type
     * @param pcrs The PCR values identifying the image
     * @param isAllowed_ Whether the image should be allowed
     */
    function setImage(CVM cvm, PCR[] calldata pcrs, bool isAllowed_) external;

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
     * @notice Gets the allowed status for an image hash
     * @param cvm The CVM type
     * @param key The image hash to check
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
