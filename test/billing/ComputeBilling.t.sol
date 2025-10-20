// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {ComputeBilling} from "../../src/billing/ComputeBilling.sol";
import {BillingCore} from "../../src/billing/BillingCore.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ERC20Mock} from "@eigenlayer-contracts/src/test/mocks/ERC20Mock.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {MockComputeBilling} from "./mocks/MockComputeBilling.sol";

contract ComputeBillingTest is Test {
    MockComputeBilling public computeBilling;
    BillingCore public billingCore;
    ERC20Mock public token;

    address public settler = makeAddr("settler");
    address public admin = makeAddr("admin");
    address public proxyAdmin = makeAddr("proxyAdmin");

    uint16 constant SKU_ID = 1;
    uint96 constant RUNNING_RATE = 1e15; // 0.001 tokens per second
    uint96 constant STOPPED_RATE = 1e14; // 0.0001 tokens per second
    uint16 constant VCPUS = 4;
    uint96 constant MIN_DEPOSIT = 1000e18;

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

        // Deploy MockComputeBilling
        computeBilling = new MockComputeBilling(address(billingCore));

        // Setup billing
        computeBilling.setSettler(settler, true);
        computeBilling.setResourceCap(1000, 1000);
        computeBilling.addSKU(SKU_ID, "Standard", RUNNING_RATE, STOPPED_RATE, VCPUS, MIN_DEPOSIT);

        // Register product with BillingCore
        vm.prank(admin);
        billingCore.registerProduct("Compute", address(computeBilling));
    }

    /**
     * @notice Test gas cost for settling 1 user
     */
    function test_settlePeriod_1user() public {
        address user1 = makeAddr("user1");
        address app1 = makeAddr("app1");

        // Setup user with sufficient balance
        vm.prank(user1);
        token.mint(user1, 10000e18);
        vm.prank(user1);
        token.approve(address(billingCore), type(uint256).max);
        vm.prank(user1);
        billingCore.deposit(5000e18);

        // Start billing for user
        computeBilling.setAppAccount(app1, user1);
        computeBilling.setAppSKU(app1, SKU_ID);
        computeBilling.setAppActive(app1, true);
        computeBilling.setAppRunning(app1, true);
        computeBilling.startBilling(app1, user1, SKU_ID);

        // Fast forward 35 days to move to next period (calendar month-based)
        vm.warp(block.timestamp + 35 days);

        // Settle for period 0
        vm.prank(settler);
        uint256 gasBefore = gasleft();
        computeBilling.settlePeriod(user1, 0);
        uint256 gasUsed = gasBefore - gasleft();

        console.log("Gas used for 1 user settlement:", gasUsed);
    }

    /**
     * @notice Test gas cost for settling 10 users
     */
    function test_settlePeriodBatch_10users() public {
        address[] memory users = new address[](10);
        address[] memory apps = new address[](10);

        // Setup 10 users
        for (uint256 i = 0; i < 10; i++) {
            users[i] = makeAddr(string(abi.encodePacked("user", vm.toString(i))));
            apps[i] = makeAddr(string(abi.encodePacked("app", vm.toString(i))));

            // Setup user with sufficient balance
            vm.prank(users[i]);
            token.mint(users[i], 10000e18);
            vm.prank(users[i]);
            token.approve(address(billingCore), type(uint256).max);
            vm.prank(users[i]);
            billingCore.deposit(5000e18);

            // Start billing for user
            computeBilling.setAppAccount(apps[i], users[i]);
            computeBilling.setAppSKU(apps[i], SKU_ID);
            computeBilling.setAppActive(apps[i], true);
            computeBilling.setAppRunning(apps[i], true);
            computeBilling.startBilling(apps[i], users[i], SKU_ID);
        }

        // Fast forward 35 days to move to next period (calendar month-based)
        vm.warp(block.timestamp + 35 days);

        // Settle for period 0
        vm.prank(settler);
        uint256 gasBefore = gasleft();
        computeBilling.settlePeriodBatch(users, 0);
        uint256 gasUsed = gasBefore - gasleft();

        console.log("Gas used for 10 users settlement:", gasUsed);
        console.log("Gas per user (10 users):", gasUsed / 10);
    }

    /**
     * @notice Test gas cost for settling 100 users
     */
    function test_settlePeriodBatch_100users() public {
        address[] memory users = new address[](100);
        address[] memory apps = new address[](100);

        // Setup 100 users
        for (uint256 i = 0; i < 100; i++) {
            users[i] = makeAddr(string(abi.encodePacked("user", vm.toString(i))));
            apps[i] = makeAddr(string(abi.encodePacked("app", vm.toString(i))));

            // Setup user with sufficient balance
            vm.prank(users[i]);
            token.mint(users[i], 10000e18);
            vm.prank(users[i]);
            token.approve(address(billingCore), type(uint256).max);
            vm.prank(users[i]);
            billingCore.deposit(5000e18);

            // Start billing for user
            computeBilling.setAppAccount(apps[i], users[i]);
            computeBilling.setAppSKU(apps[i], SKU_ID);
            computeBilling.setAppActive(apps[i], true);
            computeBilling.setAppRunning(apps[i], true);
            computeBilling.startBilling(apps[i], users[i], SKU_ID);
        }

        // Fast forward 35 days to move to next period (calendar month-based)
        vm.warp(block.timestamp + 35 days);

        // Settle for period 0
        vm.prank(settler);
        uint256 gasBefore = gasleft();
        computeBilling.settlePeriodBatch(users, 0);
        uint256 gasUsed = gasBefore - gasleft();

        console.log("Gas used for 100 users settlement:", gasUsed);
        console.log("Gas per user (100 users):", gasUsed / 100);
    }
}
