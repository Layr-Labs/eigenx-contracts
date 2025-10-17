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
}
