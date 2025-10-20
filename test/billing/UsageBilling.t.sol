// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {BillingCore} from "../../src/billing/BillingCore.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ERC20Mock} from "@eigenlayer-contracts/src/test/mocks/ERC20Mock.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {MockUsageBilling} from "./mocks/MockUsageBilling.sol";

contract UsageBillingTest is Test {
    MockUsageBilling public usageBilling;
    BillingCore public billingCore;
    ERC20Mock public token;

    address public reporter = makeAddr("reporter");
    address public admin = makeAddr("admin");
    address public proxyAdmin = makeAddr("proxyAdmin");

    uint96 constant USAGE_AMOUNT = 100e18; // 100 tokens worth of usage per user

    function setUp() public {
        // Deploy mock token
        token = new ERC20Mock();

        // Deploy BillingCore implementation
        BillingCore billingCoreImpl = new BillingCore(address(token));

        // Deploy proxy
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(billingCoreImpl),
            proxyAdmin,
            abi.encodeWithSelector(BillingCore.initialize.selector, admin)
        );

        billingCore = BillingCore(address(proxy));

        // Deploy MockUsageBilling
        usageBilling = new MockUsageBilling(address(billingCore));

        // Setup billing
        usageBilling.setUsageReporter(reporter, true);

        // Register product with BillingCore
        vm.prank(admin);
        billingCore.registerProduct("Usage", address(usageBilling));
    }

    /**
     * @notice Test gas cost for settling 1 user
     */
    function test_settlePeriod_1user() public {
        address user1 = makeAddr("user1");

        // Setup user with sufficient balance
        vm.prank(user1);
        token.mint(user1, 10000e18);
        vm.prank(user1);
        token.approve(address(billingCore), type(uint256).max);
        vm.prank(user1);
        billingCore.deposit(5000e18);

        // Record usage for user in period 0
        vm.prank(reporter);
        usageBilling.recordUsageForPeriod(user1, USAGE_AMOUNT, 0);

        // Move to next period so we can settle period 0 (calendar month-based)
        vm.warp(block.timestamp + 35 days);

        // Settle for period 0
        vm.prank(reporter);
        uint256 gasBefore = gasleft();
        usageBilling.settlePeriod(user1, 0);
        uint256 gasUsed = gasBefore - gasleft();

        console.log("Gas used for 1 user settlement:", gasUsed);
    }

    /**
     * @notice Test gas cost for settling 10 users
     */
    function test_settlePeriodBatch_10users() public {
        address[] memory users = new address[](10);
        uint96[] memory amounts = new uint96[](10);

        // Setup 10 users
        for (uint256 i = 0; i < 10; i++) {
            users[i] = makeAddr(string(abi.encodePacked("user", vm.toString(i))));
            amounts[i] = USAGE_AMOUNT;

            // Setup user with sufficient balance
            vm.prank(users[i]);
            token.mint(users[i], 10000e18);
            vm.prank(users[i]);
            token.approve(address(billingCore), type(uint256).max);
            vm.prank(users[i]);
            billingCore.deposit(5000e18);
        }

        // Record usage for all users in period 0
        vm.prank(reporter);
        usageBilling.recordUsageBatch(users, amounts, 0);

        // Move to next period so we can settle period 0 (calendar month-based)
        vm.warp(block.timestamp + 35 days);

        // Settle for period 0
        vm.prank(reporter);
        uint256 gasBefore = gasleft();
        usageBilling.settlePeriodBatch(users, 0);
        uint256 gasUsed = gasBefore - gasleft();

        console.log("Gas used for 10 users settlement:", gasUsed);
        console.log("Gas per user (10 users):", gasUsed / 10);
    }

    /**
     * @notice Test gas cost for settling 100 users
     */
    function test_settlePeriodBatch_100users() public {
        address[] memory users = new address[](100);
        uint96[] memory amounts = new uint96[](100);

        // Setup 100 users
        for (uint256 i = 0; i < 100; i++) {
            users[i] = makeAddr(string(abi.encodePacked("user", vm.toString(i))));
            amounts[i] = USAGE_AMOUNT;

            // Setup user with sufficient balance
            vm.prank(users[i]);
            token.mint(users[i], 10000e18);
            vm.prank(users[i]);
            token.approve(address(billingCore), type(uint256).max);
            vm.prank(users[i]);
            billingCore.deposit(5000e18);
        }

        // Record usage for all users in period 0
        vm.prank(reporter);
        usageBilling.recordUsageBatch(users, amounts, 0);

        // Move to next period so we can settle period 0 (calendar month-based)
        vm.warp(block.timestamp + 35 days);

        // Settle for period 0
        vm.prank(reporter);
        uint256 gasBefore = gasleft();
        usageBilling.settlePeriodBatch(users, 0);
        uint256 gasUsed = gasBefore - gasleft();

        console.log("Gas used for 100 users settlement:", gasUsed);
        console.log("Gas per user (100 users):", gasUsed / 100);
    }
}
