// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SafeTimelockFactory

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// ISafeTimelockFactorySafeConfig is an auto generated low-level Go binding around an user-defined struct.
type ISafeTimelockFactorySafeConfig struct {
	Owners    []common.Address
	Threshold *big.Int
}

// ISafeTimelockFactoryTimelockConfig is an auto generated low-level Go binding around an user-defined struct.
type ISafeTimelockFactoryTimelockConfig struct {
	MinDelay  *big.Int
	Proposers []common.Address
	Executors []common.Address
}

// SafeTimelockFactoryMetaData contains all meta data concerning the SafeTimelockFactory contract.
var SafeTimelockFactoryMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_safeSingleton\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_safeProxyFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_safeFallbackHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timelockImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateSafeAddress\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structISafeTimelockFactory.SafeConfig\",\"components\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateTimelockAddress\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deploySafe\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structISafeTimelockFactory.SafeConfig\",\"components\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"safe\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTimelock\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structISafeTimelockFactory.TimelockConfig\",\"components\":[{\"name\":\"minDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"executors\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"timelock\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isSafe\",\"inputs\":[{\"name\":\"safe\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTimelock\",\"inputs\":[{\"name\":\"timelock\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeFallbackHandler\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeProxyFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeSingleton\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timelockImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SafeDeployed\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"safe\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owners\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockDeployed\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timelock\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"minDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proposers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"executors\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"NoExecutors\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoProposers\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressExecutor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressProposer\",\"inputs\":[]}]",
	ID:  "SafeTimelockFactory",
}

// SafeTimelockFactory is an auto generated Go binding around an Ethereum contract.
type SafeTimelockFactory struct {
	abi abi.ABI
}

// NewSafeTimelockFactory creates a new instance of SafeTimelockFactory.
func NewSafeTimelockFactory() *SafeTimelockFactory {
	parsed, err := SafeTimelockFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &SafeTimelockFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *SafeTimelockFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _safeSingleton, address _safeProxyFactory, address _safeFallbackHandler, address _timelockImplementation) returns()
func (safeTimelockFactory *SafeTimelockFactory) PackConstructor(_safeSingleton common.Address, _safeProxyFactory common.Address, _safeFallbackHandler common.Address, _timelockImplementation common.Address) []byte {
	enc, err := safeTimelockFactory.abi.Pack("", _safeSingleton, _safeProxyFactory, _safeFallbackHandler, _timelockImplementation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCalculateSafeAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x68282cf3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateSafeAddress(address deployer, (address[],uint256) config, bytes32 salt) view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) PackCalculateSafeAddress(deployer common.Address, config ISafeTimelockFactorySafeConfig, salt [32]byte) []byte {
	enc, err := safeTimelockFactory.abi.Pack("calculateSafeAddress", deployer, config, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateSafeAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x68282cf3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateSafeAddress(address deployer, (address[],uint256) config, bytes32 salt) view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) TryPackCalculateSafeAddress(deployer common.Address, config ISafeTimelockFactorySafeConfig, salt [32]byte) ([]byte, error) {
	return safeTimelockFactory.abi.Pack("calculateSafeAddress", deployer, config, salt)
}

// UnpackCalculateSafeAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x68282cf3.
//
// Solidity: function calculateSafeAddress(address deployer, (address[],uint256) config, bytes32 salt) view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) UnpackCalculateSafeAddress(data []byte) (common.Address, error) {
	out, err := safeTimelockFactory.abi.Unpack("calculateSafeAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCalculateTimelockAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e5ceba7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateTimelockAddress(address deployer, bytes32 salt) view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) PackCalculateTimelockAddress(deployer common.Address, salt [32]byte) []byte {
	enc, err := safeTimelockFactory.abi.Pack("calculateTimelockAddress", deployer, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateTimelockAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e5ceba7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateTimelockAddress(address deployer, bytes32 salt) view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) TryPackCalculateTimelockAddress(deployer common.Address, salt [32]byte) ([]byte, error) {
	return safeTimelockFactory.abi.Pack("calculateTimelockAddress", deployer, salt)
}

// UnpackCalculateTimelockAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0e5ceba7.
//
// Solidity: function calculateTimelockAddress(address deployer, bytes32 salt) view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) UnpackCalculateTimelockAddress(data []byte) (common.Address, error) {
	out, err := safeTimelockFactory.abi.Unpack("calculateTimelockAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDeploySafe is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd15a64ce.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deploySafe((address[],uint256) config, bytes32 salt) returns(address safe)
func (safeTimelockFactory *SafeTimelockFactory) PackDeploySafe(config ISafeTimelockFactorySafeConfig, salt [32]byte) []byte {
	enc, err := safeTimelockFactory.abi.Pack("deploySafe", config, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeploySafe is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd15a64ce.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deploySafe((address[],uint256) config, bytes32 salt) returns(address safe)
func (safeTimelockFactory *SafeTimelockFactory) TryPackDeploySafe(config ISafeTimelockFactorySafeConfig, salt [32]byte) ([]byte, error) {
	return safeTimelockFactory.abi.Pack("deploySafe", config, salt)
}

// UnpackDeploySafe is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd15a64ce.
//
// Solidity: function deploySafe((address[],uint256) config, bytes32 salt) returns(address safe)
func (safeTimelockFactory *SafeTimelockFactory) UnpackDeploySafe(data []byte) (common.Address, error) {
	out, err := safeTimelockFactory.abi.Unpack("deploySafe", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDeployTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7638897.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deployTimelock((uint256,address[],address[]) config, bytes32 salt) returns(address timelock)
func (safeTimelockFactory *SafeTimelockFactory) PackDeployTimelock(config ISafeTimelockFactoryTimelockConfig, salt [32]byte) []byte {
	enc, err := safeTimelockFactory.abi.Pack("deployTimelock", config, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeployTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7638897.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deployTimelock((uint256,address[],address[]) config, bytes32 salt) returns(address timelock)
func (safeTimelockFactory *SafeTimelockFactory) TryPackDeployTimelock(config ISafeTimelockFactoryTimelockConfig, salt [32]byte) ([]byte, error) {
	return safeTimelockFactory.abi.Pack("deployTimelock", config, salt)
}

// UnpackDeployTimelock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb7638897.
//
// Solidity: function deployTimelock((uint256,address[],address[]) config, bytes32 salt) returns(address timelock)
func (safeTimelockFactory *SafeTimelockFactory) UnpackDeployTimelock(data []byte) (common.Address, error) {
	out, err := safeTimelockFactory.abi.Unpack("deployTimelock", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize() returns()
func (safeTimelockFactory *SafeTimelockFactory) PackInitialize() []byte {
	enc, err := safeTimelockFactory.abi.Pack("initialize")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize() returns()
func (safeTimelockFactory *SafeTimelockFactory) TryPackInitialize() ([]byte, error) {
	return safeTimelockFactory.abi.Pack("initialize")
}

// PackIsSafe is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x769a28ac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isSafe(address safe) view returns(bool)
func (safeTimelockFactory *SafeTimelockFactory) PackIsSafe(safe common.Address) []byte {
	enc, err := safeTimelockFactory.abi.Pack("isSafe", safe)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsSafe is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x769a28ac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isSafe(address safe) view returns(bool)
func (safeTimelockFactory *SafeTimelockFactory) TryPackIsSafe(safe common.Address) ([]byte, error) {
	return safeTimelockFactory.abi.Pack("isSafe", safe)
}

// UnpackIsSafe is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x769a28ac.
//
// Solidity: function isSafe(address safe) view returns(bool)
func (safeTimelockFactory *SafeTimelockFactory) UnpackIsSafe(data []byte) (bool, error) {
	out, err := safeTimelockFactory.abi.Unpack("isSafe", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9e93963c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isTimelock(address timelock) view returns(bool)
func (safeTimelockFactory *SafeTimelockFactory) PackIsTimelock(timelock common.Address) []byte {
	enc, err := safeTimelockFactory.abi.Pack("isTimelock", timelock)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9e93963c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isTimelock(address timelock) view returns(bool)
func (safeTimelockFactory *SafeTimelockFactory) TryPackIsTimelock(timelock common.Address) ([]byte, error) {
	return safeTimelockFactory.abi.Pack("isTimelock", timelock)
}

// UnpackIsTimelock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9e93963c.
//
// Solidity: function isTimelock(address timelock) view returns(bool)
func (safeTimelockFactory *SafeTimelockFactory) UnpackIsTimelock(data []byte) (bool, error) {
	out, err := safeTimelockFactory.abi.Unpack("isTimelock", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSafeFallbackHandler is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa03bad5a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeFallbackHandler() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) PackSafeFallbackHandler() []byte {
	enc, err := safeTimelockFactory.abi.Pack("safeFallbackHandler")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeFallbackHandler is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa03bad5a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeFallbackHandler() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) TryPackSafeFallbackHandler() ([]byte, error) {
	return safeTimelockFactory.abi.Pack("safeFallbackHandler")
}

// UnpackSafeFallbackHandler is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa03bad5a.
//
// Solidity: function safeFallbackHandler() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) UnpackSafeFallbackHandler(data []byte) (common.Address, error) {
	out, err := safeTimelockFactory.abi.Unpack("safeFallbackHandler", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSafeProxyFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x19964501.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeProxyFactory() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) PackSafeProxyFactory() []byte {
	enc, err := safeTimelockFactory.abi.Pack("safeProxyFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeProxyFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x19964501.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeProxyFactory() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) TryPackSafeProxyFactory() ([]byte, error) {
	return safeTimelockFactory.abi.Pack("safeProxyFactory")
}

// UnpackSafeProxyFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x19964501.
//
// Solidity: function safeProxyFactory() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) UnpackSafeProxyFactory(data []byte) (common.Address, error) {
	out, err := safeTimelockFactory.abi.Unpack("safeProxyFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSafeSingleton is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac7d146b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeSingleton() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) PackSafeSingleton() []byte {
	enc, err := safeTimelockFactory.abi.Pack("safeSingleton")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeSingleton is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac7d146b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeSingleton() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) TryPackSafeSingleton() ([]byte, error) {
	return safeTimelockFactory.abi.Pack("safeSingleton")
}

// UnpackSafeSingleton is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac7d146b.
//
// Solidity: function safeSingleton() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) UnpackSafeSingleton(data []byte) (common.Address, error) {
	out, err := safeTimelockFactory.abi.Unpack("safeSingleton", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTimelockImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5e8f3e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelockImplementation() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) PackTimelockImplementation() []byte {
	enc, err := safeTimelockFactory.abi.Pack("timelockImplementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimelockImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5e8f3e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timelockImplementation() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) TryPackTimelockImplementation() ([]byte, error) {
	return safeTimelockFactory.abi.Pack("timelockImplementation")
}

// UnpackTimelockImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc5e8f3e5.
//
// Solidity: function timelockImplementation() view returns(address)
func (safeTimelockFactory *SafeTimelockFactory) UnpackTimelockImplementation(data []byte) (common.Address, error) {
	out, err := safeTimelockFactory.abi.Unpack("timelockImplementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// SafeTimelockFactoryInitialized represents a Initialized event raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const SafeTimelockFactoryInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (SafeTimelockFactoryInitialized) ContractEventName() string {
	return SafeTimelockFactoryInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (safeTimelockFactory *SafeTimelockFactory) UnpackInitializedEvent(log *types.Log) (*SafeTimelockFactoryInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != safeTimelockFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SafeTimelockFactoryInitialized)
	if len(log.Data) > 0 {
		if err := safeTimelockFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range safeTimelockFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SafeTimelockFactorySafeDeployed represents a SafeDeployed event raised by the SafeTimelockFactory contract.
type SafeTimelockFactorySafeDeployed struct {
	Deployer  common.Address
	Safe      common.Address
	Owners    []common.Address
	Threshold *big.Int
	Salt      [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const SafeTimelockFactorySafeDeployedEventName = "SafeDeployed"

// ContractEventName returns the user-defined event name.
func (SafeTimelockFactorySafeDeployed) ContractEventName() string {
	return SafeTimelockFactorySafeDeployedEventName
}

// UnpackSafeDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SafeDeployed(address indexed deployer, address indexed safe, address[] owners, uint256 threshold, bytes32 salt)
func (safeTimelockFactory *SafeTimelockFactory) UnpackSafeDeployedEvent(log *types.Log) (*SafeTimelockFactorySafeDeployed, error) {
	event := "SafeDeployed"
	if len(log.Topics) == 0 || log.Topics[0] != safeTimelockFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SafeTimelockFactorySafeDeployed)
	if len(log.Data) > 0 {
		if err := safeTimelockFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range safeTimelockFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// SafeTimelockFactoryTimelockDeployed represents a TimelockDeployed event raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryTimelockDeployed struct {
	Deployer  common.Address
	Timelock  common.Address
	MinDelay  *big.Int
	Proposers []common.Address
	Executors []common.Address
	Salt      [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const SafeTimelockFactoryTimelockDeployedEventName = "TimelockDeployed"

// ContractEventName returns the user-defined event name.
func (SafeTimelockFactoryTimelockDeployed) ContractEventName() string {
	return SafeTimelockFactoryTimelockDeployedEventName
}

// UnpackTimelockDeployedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockDeployed(address indexed deployer, address indexed timelock, uint256 minDelay, address[] proposers, address[] executors, bytes32 salt)
func (safeTimelockFactory *SafeTimelockFactory) UnpackTimelockDeployedEvent(log *types.Log) (*SafeTimelockFactoryTimelockDeployed, error) {
	event := "TimelockDeployed"
	if len(log.Topics) == 0 || log.Topics[0] != safeTimelockFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SafeTimelockFactoryTimelockDeployed)
	if len(log.Data) > 0 {
		if err := safeTimelockFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range safeTimelockFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (safeTimelockFactory *SafeTimelockFactory) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], safeTimelockFactory.abi.Errors["NoExecutors"].ID.Bytes()[:4]) {
		return safeTimelockFactory.UnpackNoExecutorsError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeTimelockFactory.abi.Errors["NoProposers"].ID.Bytes()[:4]) {
		return safeTimelockFactory.UnpackNoProposersError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeTimelockFactory.abi.Errors["ZeroAddressExecutor"].ID.Bytes()[:4]) {
		return safeTimelockFactory.UnpackZeroAddressExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeTimelockFactory.abi.Errors["ZeroAddressProposer"].ID.Bytes()[:4]) {
		return safeTimelockFactory.UnpackZeroAddressProposerError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// SafeTimelockFactoryNoExecutors represents a NoExecutors error raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryNoExecutors struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoExecutors()
func SafeTimelockFactoryNoExecutorsErrorID() common.Hash {
	return common.HexToHash("0xebad7a3917673636faa3bf87473642d74ec4970df627b81fb966aa4512d4ff9e")
}

// UnpackNoExecutorsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoExecutors()
func (safeTimelockFactory *SafeTimelockFactory) UnpackNoExecutorsError(raw []byte) (*SafeTimelockFactoryNoExecutors, error) {
	out := new(SafeTimelockFactoryNoExecutors)
	if err := safeTimelockFactory.abi.UnpackIntoInterface(out, "NoExecutors", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeTimelockFactoryNoProposers represents a NoProposers error raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryNoProposers struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoProposers()
func SafeTimelockFactoryNoProposersErrorID() common.Hash {
	return common.HexToHash("0x80ffbcc335a5e874626ebd60b239200cd2bbff73591ae555058a6f678f2f6e01")
}

// UnpackNoProposersError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoProposers()
func (safeTimelockFactory *SafeTimelockFactory) UnpackNoProposersError(raw []byte) (*SafeTimelockFactoryNoProposers, error) {
	out := new(SafeTimelockFactoryNoProposers)
	if err := safeTimelockFactory.abi.UnpackIntoInterface(out, "NoProposers", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeTimelockFactoryZeroAddressExecutor represents a ZeroAddressExecutor error raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryZeroAddressExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddressExecutor()
func SafeTimelockFactoryZeroAddressExecutorErrorID() common.Hash {
	return common.HexToHash("0xfd1d5901a90fb747931b09fbcdd82984a0ab6a26f1a8a75e416ce9251d56e708")
}

// UnpackZeroAddressExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddressExecutor()
func (safeTimelockFactory *SafeTimelockFactory) UnpackZeroAddressExecutorError(raw []byte) (*SafeTimelockFactoryZeroAddressExecutor, error) {
	out := new(SafeTimelockFactoryZeroAddressExecutor)
	if err := safeTimelockFactory.abi.UnpackIntoInterface(out, "ZeroAddressExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeTimelockFactoryZeroAddressProposer represents a ZeroAddressProposer error raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryZeroAddressProposer struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddressProposer()
func SafeTimelockFactoryZeroAddressProposerErrorID() common.Hash {
	return common.HexToHash("0x765e2e8cd4c853dae1bb4c199ca3822bdfd4b513e6602f7132b09a8a5e701714")
}

// UnpackZeroAddressProposerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddressProposer()
func (safeTimelockFactory *SafeTimelockFactory) UnpackZeroAddressProposerError(raw []byte) (*SafeTimelockFactoryZeroAddressProposer, error) {
	out := new(SafeTimelockFactoryZeroAddressProposer)
	if err := safeTimelockFactory.abi.UnpackIntoInterface(out, "ZeroAddressProposer", raw); err != nil {
		return nil, err
	}
	return out, nil
}
