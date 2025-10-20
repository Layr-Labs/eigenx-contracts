// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {UsageBilling} from "../../../src/billing/UsageBilling.sol";

/**
 * @title MockUsageBilling
 * @notice Mock implementation of UsageBilling for testing
 */
contract MockUsageBilling is UsageBilling {
    mapping(address => uint256) public activeResourceCounts;
    mapping(address => uint40) public lastActivityTimestamps;

    constructor(address _billingCore) UsageBilling(_billingCore) {}

    function getActiveResourceCount(address account) external view override returns (uint256) {
        return activeResourceCounts[account];
    }

    function getLastActivityTimestamp(address account) external view override returns (uint40) {
        return lastActivityTimestamps[account];
    }

    // Test helpers
    function setActiveResourceCount(address account, uint256 count) external {
        activeResourceCounts[account] = count;
    }

    function setLastActivityTimestamp(address account, uint40 timestamp) external {
        lastActivityTimestamps[account] = timestamp;
    }

    function setUsageReporter(address reporter, bool authorized) external {
        _setUsageReporter(reporter, authorized);
    }
}
