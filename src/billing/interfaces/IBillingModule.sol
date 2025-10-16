// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/**
 * @title IBillingModule
 * @notice Interface for billing modules to expose charges for a period
 * @dev Products can implement this interface to allow BillingCore to query charges
 */
interface IBillingModule {
    /**
     * @notice Get charges for an account in a specific period
     * @param account The account to query
     * @param period The period to query charges for
     * @return amount The total charges for this period (whether or not charged to BillingCore yet)
     */
    function getChargesForPeriod(address account, uint40 period) external view returns (uint96 amount);

    /**
     * @notice Get total outstanding charges that have not been charged to BillingCore yet
     * @param account The account to query
     * @return amount The total charges accumulated but not yet charged to BillingCore
     */
    function getOutstandingCharges(address account) external view returns (uint96 amount);
}
