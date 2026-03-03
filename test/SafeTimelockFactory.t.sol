// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {ISafeProxyFactory} from "../src/interfaces/ISafeProxyFactory.sol";
import {ISafeTimelockFactory} from "../src/interfaces/ISafeTimelockFactory.sol";
import {SafeTimelockFactory} from "../src/factories/SafeTimelockFactory.sol";
import {TimelockControllerImpl} from "../src/governance/TimelockControllerImpl.sol";

/**
 * @dev Minimal mock for the Gnosis Safe proxy factory.
 *
 * proxyCreationCode() returns a 3-byte init code that deploys an empty contract:
 *   3d 3d f3  →  RETURNDATASIZE RETURNDATASIZE RETURN  →  returns 0 bytes → empty contract
 *
 * createProxyWithNonce mirrors the real Safe factory's create2 formula so that
 * SafeTimelockFactory.calculateSafeAddress() produces an identical address.
 */
contract MockSafeProxyFactory is ISafeProxyFactory {
    bytes internal constant PROXY_CODE = hex"3d3df3";

    function proxyCreationCode() external pure returns (bytes memory) {
        return PROXY_CODE;
    }

    function createProxyWithNonce(address _singleton, bytes memory initializer, uint256 saltNonce)
        external
        returns (address proxy)
    {
        bytes memory deploymentData = abi.encodePacked(PROXY_CODE, uint256(uint160(_singleton)));
        bytes32 salt = keccak256(abi.encodePacked(keccak256(initializer), saltNonce));
        assembly {
            proxy := create2(0, add(deploymentData, 0x20), mload(deploymentData), salt)
        }
        require(proxy != address(0), "MockSafeProxyFactory: create2 failed");
    }
}

contract SafeTimelockFactoryTest is Test {
    SafeTimelockFactory public factory;
    TimelockControllerImpl public timelockImpl;
    MockSafeProxyFactory public mockSafeProxyFactory;

    address public proposer = makeAddr("proposer");
    address public executor = makeAddr("executor");

    bytes32 public constant SALT = keccak256("test_salt");

    event SafeDeployed(
        address indexed deployer, address indexed safe, address[] owners, uint256 threshold, bytes32 salt
    );
    event TimelockDeployed(
        address indexed deployer,
        address indexed timelock,
        uint256 minDelay,
        address[] proposers,
        address[] executors,
        bytes32 salt
    );

    function setUp() public {
        timelockImpl = new TimelockControllerImpl();
        mockSafeProxyFactory = new MockSafeProxyFactory();

        SafeTimelockFactory factoryImpl = new SafeTimelockFactory({
            _safeSingleton: address(0),
            _safeProxyFactory: address(mockSafeProxyFactory),
            _defaultFallbackHandler: address(0),
            _timelockImplementation: address(timelockImpl)
        });

        ProxyAdmin proxyAdmin = new ProxyAdmin();
        factory = SafeTimelockFactory(
            address(
                new TransparentUpgradeableProxy(
                    address(factoryImpl),
                    address(proxyAdmin),
                    abi.encodeCall(SafeTimelockFactory.initialize, ())
                )
            )
        );
    }

    // ========== Timelock deployment tests ==========

    function test_deployTimelock() public {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;
        address[] memory executors = new address[](1);
        executors[0] = executor;

        ISafeTimelockFactory.TimelockConfig memory config =
            ISafeTimelockFactory.TimelockConfig({minDelay: 1 days, proposers: proposers, executors: executors});

        vm.expectEmit(true, false, false, true);
        emit TimelockDeployed(address(this), address(0), 1 days, proposers, executors, SALT);

        address timelock = factory.deployTimelock(config, SALT);

        assertTrue(timelock != address(0), "timelock is zero");
        assertTrue(factory.isTimelock(timelock), "isTimelock should be true");
        assertEq(TimelockControllerImpl(payable(timelock)).getMinDelay(), 1 days, "minDelay mismatch");
    }

    function test_deployTimelock_timelockInitialized() public {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;
        address[] memory executors = new address[](1);
        executors[0] = executor;

        ISafeTimelockFactory.TimelockConfig memory config =
            ISafeTimelockFactory.TimelockConfig({minDelay: 2 days, proposers: proposers, executors: executors});

        address timelock = factory.deployTimelock(config, SALT);
        TimelockControllerImpl tc = TimelockControllerImpl(payable(timelock));

        bytes32 PROPOSER_ROLE = tc.PROPOSER_ROLE();
        bytes32 EXECUTOR_ROLE = tc.EXECUTOR_ROLE();
        bytes32 CANCELLER_ROLE = tc.CANCELLER_ROLE();

        assertTrue(tc.hasRole(PROPOSER_ROLE, proposer), "proposer missing PROPOSER_ROLE");
        assertTrue(tc.hasRole(CANCELLER_ROLE, proposer), "proposer missing CANCELLER_ROLE");
        assertTrue(tc.hasRole(EXECUTOR_ROLE, executor), "executor missing EXECUTOR_ROLE");
    }

    function test_deployTimelock_deterministicAddress() public {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;
        address[] memory executors = new address[](1);
        executors[0] = executor;

        ISafeTimelockFactory.TimelockConfig memory config =
            ISafeTimelockFactory.TimelockConfig({minDelay: 1 days, proposers: proposers, executors: executors});

        address predicted = factory.calculateTimelockAddress(address(this), SALT);
        address deployed = factory.deployTimelock(config, SALT);

        assertEq(deployed, predicted, "deployed address does not match predicted");
    }

    function test_deployTimelock_differentSalts() public {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;
        address[] memory executors = new address[](1);
        executors[0] = executor;

        ISafeTimelockFactory.TimelockConfig memory config =
            ISafeTimelockFactory.TimelockConfig({minDelay: 1 days, proposers: proposers, executors: executors});

        bytes32 salt1 = keccak256("salt1");
        bytes32 salt2 = keccak256("salt2");

        address timelock1 = factory.deployTimelock(config, salt1);
        address timelock2 = factory.deployTimelock(config, salt2);

        assertTrue(timelock1 != timelock2, "different salts should produce different addresses");
        assertTrue(factory.isTimelock(timelock1), "timelock1 not registered");
        assertTrue(factory.isTimelock(timelock2), "timelock2 not registered");
    }

    function test_deployTimelock_reverts_noProposers() public {
        address[] memory proposers = new address[](0);
        address[] memory executors = new address[](1);
        executors[0] = executor;

        ISafeTimelockFactory.TimelockConfig memory config =
            ISafeTimelockFactory.TimelockConfig({minDelay: 1 days, proposers: proposers, executors: executors});

        vm.expectRevert(ISafeTimelockFactory.NoProposers.selector);
        factory.deployTimelock(config, SALT);
    }

    function test_deployTimelock_reverts_noExecutors() public {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;
        address[] memory executors = new address[](0);

        ISafeTimelockFactory.TimelockConfig memory config =
            ISafeTimelockFactory.TimelockConfig({minDelay: 1 days, proposers: proposers, executors: executors});

        vm.expectRevert(ISafeTimelockFactory.NoExecutors.selector);
        factory.deployTimelock(config, SALT);
    }

    function test_deployTimelock_reverts_zeroAddressProposer() public {
        address[] memory proposers = new address[](1);
        proposers[0] = address(0);
        address[] memory executors = new address[](1);
        executors[0] = executor;

        ISafeTimelockFactory.TimelockConfig memory config =
            ISafeTimelockFactory.TimelockConfig({minDelay: 1 days, proposers: proposers, executors: executors});

        vm.expectRevert(ISafeTimelockFactory.ZeroAddressProposer.selector);
        factory.deployTimelock(config, SALT);
    }

    function test_deployTimelock_reverts_zeroAddressExecutor() public {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;
        address[] memory executors = new address[](1);
        executors[0] = address(0);

        ISafeTimelockFactory.TimelockConfig memory config =
            ISafeTimelockFactory.TimelockConfig({minDelay: 1 days, proposers: proposers, executors: executors});

        vm.expectRevert(ISafeTimelockFactory.ZeroAddressExecutor.selector);
        factory.deployTimelock(config, SALT);
    }

    // ========== Safe deployment tests ==========

    function test_deploySafe() public {
        address[] memory owners = new address[](2);
        owners[0] = makeAddr("owner1");
        owners[1] = makeAddr("owner2");

        ISafeTimelockFactory.SafeConfig memory config =
            ISafeTimelockFactory.SafeConfig({owners: owners, threshold: 2});

        vm.expectEmit(true, false, false, true);
        emit SafeDeployed(address(this), address(0), owners, 2, SALT);

        address safe = factory.deploySafe(config, SALT);

        assertTrue(safe != address(0), "safe is zero");
        assertTrue(factory.isSafe(safe), "isSafe should be true");
    }

    function test_deploySafe_deterministicAddress() public {
        address[] memory owners = new address[](1);
        owners[0] = makeAddr("owner");

        ISafeTimelockFactory.SafeConfig memory config =
            ISafeTimelockFactory.SafeConfig({owners: owners, threshold: 1});

        address predicted = factory.calculateSafeAddress(address(this), config, SALT);
        address deployed = factory.deploySafe(config, SALT);

        assertEq(deployed, predicted, "deployed Safe address does not match predicted");
    }

    // ========== Registry tests ==========

    function test_isTimelock_false_for_random_address() public view {
        assertFalse(factory.isTimelock(address(0xDEAD)), "random address should not be a timelock");
    }

    function test_isSafe_false_for_random_address() public view {
        assertFalse(factory.isSafe(address(0xBEEF)), "random address should not be a safe");
    }
}
