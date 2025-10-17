// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/**
 * @title IBillingModule
 * @notice Interface implemented by all billing modules
 */
interface IBillingModule {
    /**
     * @notice Get charges for a specific period
     * @param account The account to query
     * @param period The period to query
     * @return amount The charges for this period
     * @return settled Whether the charges have been settled
     */
    function getChargesForPeriod(address account, uint40 period) external view returns (uint96 amount, bool settled);

    /**
     * @notice Get the count of active resources for an account
     * @param account The account to check
     * @return The number of active resources (e.g. VMs, API keys, etc)
     */
    function getActiveResourceCount(address account) external view returns (uint256);

    /**
     * @notice Get the timestamp of the last activity for an account
     * @param account The account to check
     * @return The timestamp of the last billing-related activity
     */
    function getLastActivityTimestamp(address account) external view returns (uint40);
}
