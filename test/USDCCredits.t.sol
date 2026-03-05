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
import {USDCCredits} from "../src/USDCCredits.sol";
import {IUSDCCredits} from "../src/interfaces/IUSDCCredits.sol";

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

contract USDCCreditsTest is Test {
    uint256 public constant MINIMUM_DEPOSIT = 5_000_000; // $5 USDC (6 decimals)
    uint256 public constant DEPOSIT_AMOUNT = 100_000_000; // $100 USDC

    address public owner = makeAddr("owner");
    address public depositor = makeAddr("depositor");
    address public recipient = makeAddr("recipient");
    address public treasuryAddr = makeAddr("treasury");
    address public unauthorizedCaller = makeAddr("unauthorizedCaller");

    MockUSDC public mockUSDC;
    MockToken public mockToken;
    USDCCredits public usdcCredits;
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
        USDCCredits impl = new USDCCredits({_usdc: IERC20(address(mockUSDC)), _treasury: treasuryAddr});

        // Upgrade and initialize
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(proxy)),
            address(impl),
            abi.encodeCall(USDCCredits.initialize, (owner, MINIMUM_DEPOSIT))
        );

        usdcCredits = USDCCredits(address(proxy));

        // Fund depositor with USDC
        mockUSDC.mint(depositor, 1_000_000_000); // $1000
    }

    // ==================== purchaseCredits() tests ====================

    function test_purchaseCredits() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcCredits), DEPOSIT_AMOUNT);

        vm.expectEmit(true, true, true, true);
        emit Deposit(depositor, depositor, DEPOSIT_AMOUNT);

        usdcCredits.purchaseCredits(DEPOSIT_AMOUNT);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), DEPOSIT_AMOUNT);
    }

    function test_purchaseCredits_belowMinimum() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcCredits), MINIMUM_DEPOSIT - 1);

        vm.expectRevert(IUSDCCredits.BelowMinimumDeposit.selector);
        usdcCredits.purchaseCredits(MINIMUM_DEPOSIT - 1);
        vm.stopPrank();
    }

    function test_purchaseCredits_exactMinimum() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcCredits), MINIMUM_DEPOSIT);

        usdcCredits.purchaseCredits(MINIMUM_DEPOSIT);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), MINIMUM_DEPOSIT);
    }

    function test_purchaseCredits_noApproval() public {
        vm.prank(depositor);
        vm.expectRevert("ERC20: insufficient allowance");
        usdcCredits.purchaseCredits(DEPOSIT_AMOUNT);
    }

    // ==================== purchaseCreditsFor() tests ====================

    function test_purchaseCreditsFor() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcCredits), DEPOSIT_AMOUNT);

        vm.expectEmit(true, true, true, true);
        emit Deposit(depositor, recipient, DEPOSIT_AMOUNT);

        usdcCredits.purchaseCreditsFor(DEPOSIT_AMOUNT, recipient);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), DEPOSIT_AMOUNT);
    }

    function test_purchaseCreditsFor_zeroAddress() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcCredits), DEPOSIT_AMOUNT);

        vm.expectRevert(IUSDCCredits.ZeroAddress.selector);
        usdcCredits.purchaseCreditsFor(DEPOSIT_AMOUNT, address(0));
        vm.stopPrank();
    }

    function test_purchaseCreditsFor_belowMinimum() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcCredits), MINIMUM_DEPOSIT - 1);

        vm.expectRevert(IUSDCCredits.BelowMinimumDeposit.selector);
        usdcCredits.purchaseCreditsFor(MINIMUM_DEPOSIT - 1, recipient);
        vm.stopPrank();
    }

    function test_purchaseCreditsFor_multipleDeposits() public {
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcCredits), DEPOSIT_AMOUNT * 3);

        usdcCredits.purchaseCreditsFor(DEPOSIT_AMOUNT, recipient);
        usdcCredits.purchaseCreditsFor(DEPOSIT_AMOUNT, recipient);
        usdcCredits.purchaseCreditsFor(DEPOSIT_AMOUNT, depositor);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), DEPOSIT_AMOUNT * 3);
    }

    // ==================== setMinimumDepositFor() tests ====================

    function test_setMinimumDepositFor() public {
        uint256 newMinimum = 10_000_000; // $10

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit MinimumDepositSet(MINIMUM_DEPOSIT, newMinimum);

        usdcCredits.setMinimumDepositFor(newMinimum);

        assertEq(usdcCredits.minimumDeposit(), newMinimum);
    }

    function test_setMinimumDepositFor_unauthorized() public {
        vm.prank(unauthorizedCaller);
        vm.expectRevert("Ownable: caller is not the owner");
        usdcCredits.setMinimumDepositFor(10_000_000);
    }

    function test_setMinimumDepositFor_toZero() public {
        vm.prank(owner);
        usdcCredits.setMinimumDepositFor(0);

        assertEq(usdcCredits.minimumDeposit(), 0);

        // Any amount should now work
        vm.startPrank(depositor);
        mockUSDC.approve(address(usdcCredits), 1);
        usdcCredits.purchaseCredits(1);
        vm.stopPrank();

        assertEq(mockUSDC.balanceOf(treasuryAddr), 1);
    }

    // ==================== sweep() tests ====================

    function test_sweep_usdc() public {
        // Simulate accidental direct transfer
        mockUSDC.mint(address(usdcCredits), 50_000_000);

        vm.prank(owner);
        usdcCredits.sweep(IERC20(address(mockUSDC)));

        assertEq(mockUSDC.balanceOf(address(usdcCredits)), 0);
        assertEq(mockUSDC.balanceOf(treasuryAddr), 50_000_000);
    }

    function test_sweep_otherToken() public {
        // Simulate accidental send of another token
        mockToken.mint(address(usdcCredits), 1_000_000);

        vm.prank(owner);
        usdcCredits.sweep(IERC20(address(mockToken)));

        assertEq(mockToken.balanceOf(address(usdcCredits)), 0);
        assertEq(mockToken.balanceOf(treasuryAddr), 1_000_000);
    }

    function test_sweep_zeroBalance() public {
        // Should not revert on zero balance
        vm.prank(owner);
        usdcCredits.sweep(IERC20(address(mockUSDC)));

        assertEq(mockUSDC.balanceOf(treasuryAddr), 0);
    }

    function test_sweep_unauthorized() public {
        mockUSDC.mint(address(usdcCredits), 50_000_000);

        vm.prank(unauthorizedCaller);
        vm.expectRevert("Ownable: caller is not the owner");
        usdcCredits.sweep(IERC20(address(mockUSDC)));
    }

    // ==================== Constructor/initializer validation tests ====================

    function test_constructor_zeroUsdc() public {
        vm.expectRevert(IUSDCCredits.ZeroAddress.selector);
        new USDCCredits({_usdc: IERC20(address(0)), _treasury: treasuryAddr});
    }

    function test_constructor_zeroTreasury() public {
        vm.expectRevert(IUSDCCredits.ZeroAddress.selector);
        new USDCCredits({_usdc: IERC20(address(mockUSDC)), _treasury: address(0)});
    }

    function test_initialize_zeroOwner() public {
        ProxyAdmin newProxyAdmin = new ProxyAdmin();
        EmptyContract emptyContract = new EmptyContract();
        TransparentUpgradeableProxy proxy =
            new TransparentUpgradeableProxy(address(emptyContract), address(newProxyAdmin), new bytes(0));

        USDCCredits impl = new USDCCredits({_usdc: IERC20(address(mockUSDC)), _treasury: treasuryAddr});

        // OwnableUpgradeable._transferOwnership does not revert on zero address,
        // but we verify the owner is correctly set in the normal path
        newProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(proxy)),
            address(impl),
            abi.encodeCall(USDCCredits.initialize, (address(0), MINIMUM_DEPOSIT))
        );
        // Owner should be zero — effectively locked, but not reverting
        assertEq(USDCCredits(address(proxy)).owner(), address(0));
    }

    // ==================== View function tests ====================

    function test_immutables() public view {
        assertEq(address(usdcCredits.usdc()), address(mockUSDC));
        assertEq(usdcCredits.treasury(), treasuryAddr);
        assertEq(usdcCredits.minimumDeposit(), MINIMUM_DEPOSIT);
        assertEq(usdcCredits.owner(), owner);
    }
}
