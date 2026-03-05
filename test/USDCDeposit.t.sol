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
    uint256 public constant MINIMUM_DEPOSIT = 5_000_000; // $5 USDC (6 decimals)
    uint256 public constant DEPOSIT_AMOUNT = 100_000_000; // $100 USDC

    address public owner = makeAddr("owner");
    address public depositor = makeAddr("depositor");
    address public recipient = makeAddr("recipient");
    address public treasuryAddr = makeAddr("treasury");
    address public unauthorizedCaller = makeAddr("unauthorizedCaller");

    MockUSDC public mockUSDC;
    MockToken public mockToken;
    USDCDeposit public usdcDeposit;
    ProxyAdmin public proxyAdmin;

    event Deposit(address indexed depositor, address indexed account, uint256 amount);
    event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum);

    function setUp() public {
        // Deploy mock tokens
        mockUSDC = new MockUSDC();
        mockToken = new MockToken();

        // Deploy proxy infrastructure
        proxyAdmin = new ProxyAdmin();
        EmptyContract emptyContract = new EmptyContract();

        TransparentUpgradeableProxy proxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), new bytes(0));

        // Deploy implementation
        USDCDeposit impl = new USDCDeposit({_usdc: IERC20(address(mockUSDC)), _treasury: treasuryAddr});

        // Upgrade and initialize
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(proxy)),
            address(impl),
            abi.encodeCall(USDCDeposit.initialize, (owner, MINIMUM_DEPOSIT))
        );

        usdcDeposit = USDCDeposit(address(proxy));

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

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit MinimumDepositSet(MINIMUM_DEPOSIT, newMinimum);

        usdcDeposit.setMinimumDeposit(newMinimum);

        assertEq(usdcDeposit.minimumDeposit(), newMinimum);
    }

    function test_setMinimumDeposit_unauthorized() public {
        vm.prank(unauthorizedCaller);
        vm.expectRevert("Ownable: caller is not the owner");
        usdcDeposit.setMinimumDeposit(10_000_000);
    }

    function test_setMinimumDeposit_toZero() public {
        vm.prank(owner);
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

        vm.prank(owner);
        usdcDeposit.sweep(IERC20(address(mockUSDC)));

        assertEq(mockUSDC.balanceOf(address(usdcDeposit)), 0);
        assertEq(mockUSDC.balanceOf(treasuryAddr), 50_000_000);
    }

    function test_sweep_otherToken() public {
        // Simulate accidental send of another token
        mockToken.mint(address(usdcDeposit), 1_000_000);

        vm.prank(owner);
        usdcDeposit.sweep(IERC20(address(mockToken)));

        assertEq(mockToken.balanceOf(address(usdcDeposit)), 0);
        assertEq(mockToken.balanceOf(treasuryAddr), 1_000_000);
    }

    function test_sweep_zeroBalance() public {
        // Should not revert on zero balance
        vm.prank(owner);
        usdcDeposit.sweep(IERC20(address(mockUSDC)));

        assertEq(mockUSDC.balanceOf(treasuryAddr), 0);
    }

    function test_sweep_unauthorized() public {
        mockUSDC.mint(address(usdcDeposit), 50_000_000);

        vm.prank(unauthorizedCaller);
        vm.expectRevert("Ownable: caller is not the owner");
        usdcDeposit.sweep(IERC20(address(mockUSDC)));
    }

    // ==================== Constructor/initializer validation tests ====================

    function test_constructor_zeroUsdc() public {
        vm.expectRevert(IUSDCDeposit.ZeroAddress.selector);
        new USDCDeposit({_usdc: IERC20(address(0)), _treasury: treasuryAddr});
    }

    function test_constructor_zeroTreasury() public {
        vm.expectRevert(IUSDCDeposit.ZeroAddress.selector);
        new USDCDeposit({_usdc: IERC20(address(mockUSDC)), _treasury: address(0)});
    }

    function test_initialize_zeroOwner() public {
        ProxyAdmin newProxyAdmin = new ProxyAdmin();
        EmptyContract emptyContract = new EmptyContract();
        TransparentUpgradeableProxy proxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(newProxyAdmin), new bytes(0));

        USDCDeposit impl = new USDCDeposit({_usdc: IERC20(address(mockUSDC)), _treasury: treasuryAddr});

        // OwnableUpgradeable._transferOwnership does not revert on zero address,
        // but we verify the owner is correctly set in the normal path
        newProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(proxy)),
            address(impl),
            abi.encodeCall(USDCDeposit.initialize, (address(0), MINIMUM_DEPOSIT))
        );
        // Owner should be zero — effectively locked, but not reverting
        assertEq(USDCDeposit(address(proxy)).owner(), address(0));
    }

    // ==================== View function tests ====================

    function test_immutables() public view {
        assertEq(address(usdcDeposit.usdc()), address(mockUSDC));
        assertEq(usdcDeposit.treasury(), treasuryAddr);
        assertEq(usdcDeposit.minimumDeposit(), MINIMUM_DEPOSIT);
        assertEq(usdcDeposit.owner(), owner);
    }
}
