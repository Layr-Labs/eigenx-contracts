// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {
    TransparentUpgradeableProxy,
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {PermissionController} from "@eigenlayer-contracts/src/contracts/permissions/PermissionController.sol";
import {IPermissionController} from "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";
import {PermissionControllerMixin} from "@eigenlayer-contracts/src/contracts/mixins/PermissionControllerMixin.sol";
import {EmptyContract} from "@eigenlayer-contracts/src/test/mocks/EmptyContract.sol";
import {USDCDeposit} from "../src/USDCDeposit.sol";
import {IUSDCDeposit} from "../src/interfaces/IUSDCDeposit.sol";

/// @notice Mock ERC20 token for testing
contract MockUSDC is ERC20 {
    constructor() ERC20("USD Coin", "USDC") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function decimals() public pure override returns (uint8) {
        return 6;
    }
}

/// @notice Mock ERC20 for testing sweep of non-USDC tokens
contract MockToken is ERC20 {
    constructor() ERC20("Mock Token", "MOCK") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract USDCDepositTest is Test {
    string public constant VERSION = "1.0.0-test";
    uint256 public constant MINIMUM_DEPOSIT = 5_000_000; // $5 USDC (6 decimals)
    uint256 public constant DEPOSIT_AMOUNT = 100_000_000; // $100 USDC

    address public admin = makeAddr("admin");
    address public depositor = makeAddr("depositor");
    address public recipient = makeAddr("recipient");
    address public treasuryAddr = makeAddr("treasury");
    address public unauthorizedCaller = makeAddr("unauthorizedCaller");

    PermissionController public permissionController;
    MockUSDC public mockUSDC;
    MockToken public mockToken;
    USDCDeposit public usdcDeposit;
    ProxyAdmin public proxyAdmin;

    event Deposit(address indexed depositor, address indexed account, uint256 amount);
    event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum);

    function setUp() public {
        // Deploy PermissionController
        permissionController = new PermissionController(VERSION);

        // Deploy mock tokens
        mockUSDC = new MockUSDC();
        mockToken = new MockToken();

        // Deploy proxy infrastructure
        proxyAdmin = new ProxyAdmin();
        EmptyContract emptyContract = new EmptyContract();

        TransparentUpgradeableProxy proxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), new bytes(0));

        // Deploy implementation
        USDCDeposit impl = new USDCDeposit({
            _version: VERSION,
            _permissionController: IPermissionController(address(permissionController)),
            _usdc: IERC20(address(mockUSDC)),
            _treasury: treasuryAddr
        });

        // Upgrade and initialize
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(proxy)),
            address(impl),
            abi.encodeCall(USDCDeposit.initialize, (admin, MINIMUM_DEPOSIT))
        );

        usdcDeposit = USDCDeposit(address(proxy));

        // Accept admin role
        vm.prank(admin);
        permissionController.acceptAdmin(address(usdcDeposit));

        // Fund depositor with USDC
        mockUSDC.mint(depositor, 1_000_000_000); // $1000
    }

    // ==================== deposit() tests ====================

    function test_deposit() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcDeposit), DEPOSIT_AMOUNT);

        vm.expectEmit(true, true, true, true);
        emit Deposit(depositor, depositor, DEPOSIT_AMOUNT);

        usdcDeposit.deposit(DEPOSIT_AMOUNT);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), DEPOSIT_AMOUNT);
    }

    function test_deposit_belowMinimum() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcDeposit), MINIMUM_DEPOSIT - 1);

        vm.expectRevert(IUSDCDeposit.BelowMinimumDeposit.selector);
        usdcDeposit.deposit(MINIMUM_DEPOSIT - 1);
        vm.stopPrank();
    }

    function test_deposit_exactMinimum() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcDeposit), MINIMUM_DEPOSIT);

        usdcDeposit.deposit(MINIMUM_DEPOSIT);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), MINIMUM_DEPOSIT);
    }

    function test_deposit_noApproval() public {
        vm.prank(depositor);
        vm.expectRevert("ERC20: insufficient allowance");
        usdcDeposit.deposit(DEPOSIT_AMOUNT);
    }

    // ==================== depositFor() tests ====================

    function test_depositFor() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcDeposit), DEPOSIT_AMOUNT);

        vm.expectEmit(true, true, true, true);
        emit Deposit(depositor, recipient, DEPOSIT_AMOUNT);

        usdcDeposit.depositFor(DEPOSIT_AMOUNT, recipient);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), DEPOSIT_AMOUNT);
    }

    function test_depositFor_zeroAddress() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcDeposit), DEPOSIT_AMOUNT);

        vm.expectRevert(IUSDCDeposit.ZeroAddress.selector);
        usdcDeposit.depositFor(DEPOSIT_AMOUNT, address(0));
        vm.stopPrank();
    }

    function test_depositFor_belowMinimum() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcDeposit), MINIMUM_DEPOSIT - 1);

        vm.expectRevert(IUSDCDeposit.BelowMinimumDeposit.selector);
        usdcDeposit.depositFor(MINIMUM_DEPOSIT - 1, recipient);
        vm.stopPrank();
    }

    function test_depositFor_multipleDeposits() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcDeposit), DEPOSIT_AMOUNT * 3);

        usdcDeposit.depositFor(DEPOSIT_AMOUNT, recipient);
        usdcDeposit.depositFor(DEPOSIT_AMOUNT, recipient);
        usdcDeposit.depositFor(DEPOSIT_AMOUNT, depositor);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), DEPOSIT_AMOUNT * 3);
    }

    // ==================== setMinimumDeposit() tests ====================

    function test_setMinimumDeposit() public {
        uint256 newMinimum = 10_000_000; // $10

        vm.prank(admin);
        vm.expectEmit(true, true, true, true);
        emit MinimumDepositSet(MINIMUM_DEPOSIT, newMinimum);

        usdcDeposit.setMinimumDeposit(newMinimum);

        assertEq(usdcDeposit.minimumDeposit(), newMinimum);
    }

    function test_setMinimumDeposit_unauthorized() public {
        vm.prank(unauthorizedCaller);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        usdcDeposit.setMinimumDeposit(10_000_000);
    }

    function test_setMinimumDeposit_toZero() public {
        vm.prank(admin);
        usdcDeposit.setMinimumDeposit(0);

        assertEq(usdcDeposit.minimumDeposit(), 0);

        // Any amount should now work
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcDeposit), 1);
        usdcDeposit.deposit(1);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), 1);
    }

    // ==================== sweep() tests ====================

    function test_sweep_usdc() public {
        // Simulate accidental direct transfer
        mockUSDC.mint(address(usdcDeposit), 50_000_000);

        vm.prank(admin);
        usdcDeposit.sweep(IERC20(address(mockUSDC)));

        assertEq(mockUSDC.balanceOf(address(usdcDeposit)), 0);
        assertEq(mockUSDC.balanceOf(treasuryAddr), 50_000_000);
    }

    function test_sweep_otherToken() public {
        // Simulate accidental send of another token
        mockToken.mint(address(usdcDeposit), 1_000_000);

        vm.prank(admin);
        usdcDeposit.sweep(IERC20(address(mockToken)));

        assertEq(mockToken.balanceOf(address(usdcDeposit)), 0);
        assertEq(mockToken.balanceOf(treasuryAddr), 1_000_000);
    }

    function test_sweep_zeroBalance() public {
        // Should not revert on zero balance
        vm.prank(admin);
        usdcDeposit.sweep(IERC20(address(mockUSDC)));

        assertEq(mockUSDC.balanceOf(treasuryAddr), 0);
    }

    function test_sweep_unauthorized() public {
        mockUSDC.mint(address(usdcDeposit), 50_000_000);

        vm.prank(unauthorizedCaller);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        usdcDeposit.sweep(IERC20(address(mockUSDC)));
    }

    // ==================== View function tests ====================

    function test_immutables() public view {
        assertEq(address(usdcDeposit.usdc()), address(mockUSDC));
        assertEq(usdcDeposit.treasury(), treasuryAddr);
        assertEq(usdcDeposit.minimumDeposit(), MINIMUM_DEPOSIT);
    }
}
