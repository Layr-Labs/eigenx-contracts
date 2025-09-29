// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

interface IApp {
    /**
     * @notice Initializes the App
     * @param admin The UAM admin to be set for the app
     */
    function initialize(address admin) external;
}
