// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {IImageAllowlist} from "../interfaces/IImageAllowlist.sol";

abstract contract ImageAllowlistStorage is IImageAllowlist {
    /// @inheritdoc IImageAllowlist
    mapping(CVM => mapping(bytes32 => bool)) public images;

    /// @inheritdoc IImageAllowlist
    mapping(CVM => uint64) public minimumTCB;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
