// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SafeTimelockFactory

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
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
var SafeTimelockFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_safeSingleton\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_safeProxyFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_safeFallbackHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timelockImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateSafeAddress\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structISafeTimelockFactory.SafeConfig\",\"components\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateTimelockAddress\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deploySafe\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structISafeTimelockFactory.SafeConfig\",\"components\":[{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"safe\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployTimelock\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structISafeTimelockFactory.TimelockConfig\",\"components\":[{\"name\":\"minDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"executors\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"timelock\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSafesByDeployer\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimelocksByDeployer\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isSafe\",\"inputs\":[{\"name\":\"safe\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTimelock\",\"inputs\":[{\"name\":\"timelock\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeFallbackHandler\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeProxyFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeSingleton\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timelockImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SafeDeployed\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"safe\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owners\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockDeployed\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timelock\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"minDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proposers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"executors\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"NoExecutors\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoProposers\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressExecutor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressProposer\",\"inputs\":[]}]",
}

// SafeTimelockFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeTimelockFactoryMetaData.ABI instead.
var SafeTimelockFactoryABI = SafeTimelockFactoryMetaData.ABI

// SafeTimelockFactory is an auto generated Go binding around an Ethereum contract.
type SafeTimelockFactory struct {
	SafeTimelockFactoryCaller     // Read-only binding to the contract
	SafeTimelockFactoryTransactor // Write-only binding to the contract
	SafeTimelockFactoryFilterer   // Log filterer for contract events
}

// SafeTimelockFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeTimelockFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTimelockFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeTimelockFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTimelockFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeTimelockFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTimelockFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeTimelockFactorySession struct {
	Contract     *SafeTimelockFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeTimelockFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeTimelockFactoryCallerSession struct {
	Contract *SafeTimelockFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SafeTimelockFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeTimelockFactoryTransactorSession struct {
	Contract     *SafeTimelockFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SafeTimelockFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeTimelockFactoryRaw struct {
	Contract *SafeTimelockFactory // Generic contract binding to access the raw methods on
}

// SafeTimelockFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeTimelockFactoryCallerRaw struct {
	Contract *SafeTimelockFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SafeTimelockFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeTimelockFactoryTransactorRaw struct {
	Contract *SafeTimelockFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeTimelockFactory creates a new instance of SafeTimelockFactory, bound to a specific deployed contract.
func NewSafeTimelockFactory(address common.Address, backend bind.ContractBackend) (*SafeTimelockFactory, error) {
	contract, err := bindSafeTimelockFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeTimelockFactory{SafeTimelockFactoryCaller: SafeTimelockFactoryCaller{contract: contract}, SafeTimelockFactoryTransactor: SafeTimelockFactoryTransactor{contract: contract}, SafeTimelockFactoryFilterer: SafeTimelockFactoryFilterer{contract: contract}}, nil
}

// NewSafeTimelockFactoryCaller creates a new read-only instance of SafeTimelockFactory, bound to a specific deployed contract.
func NewSafeTimelockFactoryCaller(address common.Address, caller bind.ContractCaller) (*SafeTimelockFactoryCaller, error) {
	contract, err := bindSafeTimelockFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeTimelockFactoryCaller{contract: contract}, nil
}

// NewSafeTimelockFactoryTransactor creates a new write-only instance of SafeTimelockFactory, bound to a specific deployed contract.
func NewSafeTimelockFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeTimelockFactoryTransactor, error) {
	contract, err := bindSafeTimelockFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeTimelockFactoryTransactor{contract: contract}, nil
}

// NewSafeTimelockFactoryFilterer creates a new log filterer instance of SafeTimelockFactory, bound to a specific deployed contract.
func NewSafeTimelockFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeTimelockFactoryFilterer, error) {
	contract, err := bindSafeTimelockFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeTimelockFactoryFilterer{contract: contract}, nil
}

// bindSafeTimelockFactory binds a generic wrapper to an already deployed contract.
func bindSafeTimelockFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeTimelockFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeTimelockFactory *SafeTimelockFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeTimelockFactory.Contract.SafeTimelockFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeTimelockFactory *SafeTimelockFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.SafeTimelockFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeTimelockFactory *SafeTimelockFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.SafeTimelockFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeTimelockFactory *SafeTimelockFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeTimelockFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeTimelockFactory *SafeTimelockFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeTimelockFactory *SafeTimelockFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.contract.Transact(opts, method, params...)
}

// CalculateSafeAddress is a free data retrieval call binding the contract method 0x68282cf3.
//
// Solidity: function calculateSafeAddress(address deployer, (address[],uint256) config, bytes32 salt) view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) CalculateSafeAddress(opts *bind.CallOpts, deployer common.Address, config ISafeTimelockFactorySafeConfig, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "calculateSafeAddress", deployer, config, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateSafeAddress is a free data retrieval call binding the contract method 0x68282cf3.
//
// Solidity: function calculateSafeAddress(address deployer, (address[],uint256) config, bytes32 salt) view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactorySession) CalculateSafeAddress(deployer common.Address, config ISafeTimelockFactorySafeConfig, salt [32]byte) (common.Address, error) {
	return _SafeTimelockFactory.Contract.CalculateSafeAddress(&_SafeTimelockFactory.CallOpts, deployer, config, salt)
}

// CalculateSafeAddress is a free data retrieval call binding the contract method 0x68282cf3.
//
// Solidity: function calculateSafeAddress(address deployer, (address[],uint256) config, bytes32 salt) view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) CalculateSafeAddress(deployer common.Address, config ISafeTimelockFactorySafeConfig, salt [32]byte) (common.Address, error) {
	return _SafeTimelockFactory.Contract.CalculateSafeAddress(&_SafeTimelockFactory.CallOpts, deployer, config, salt)
}

// CalculateTimelockAddress is a free data retrieval call binding the contract method 0x0e5ceba7.
//
// Solidity: function calculateTimelockAddress(address deployer, bytes32 salt) view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) CalculateTimelockAddress(opts *bind.CallOpts, deployer common.Address, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "calculateTimelockAddress", deployer, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateTimelockAddress is a free data retrieval call binding the contract method 0x0e5ceba7.
//
// Solidity: function calculateTimelockAddress(address deployer, bytes32 salt) view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactorySession) CalculateTimelockAddress(deployer common.Address, salt [32]byte) (common.Address, error) {
	return _SafeTimelockFactory.Contract.CalculateTimelockAddress(&_SafeTimelockFactory.CallOpts, deployer, salt)
}

// CalculateTimelockAddress is a free data retrieval call binding the contract method 0x0e5ceba7.
//
// Solidity: function calculateTimelockAddress(address deployer, bytes32 salt) view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) CalculateTimelockAddress(deployer common.Address, salt [32]byte) (common.Address, error) {
	return _SafeTimelockFactory.Contract.CalculateTimelockAddress(&_SafeTimelockFactory.CallOpts, deployer, salt)
}

// GetSafesByDeployer is a free data retrieval call binding the contract method 0x1eae53e8.
//
// Solidity: function getSafesByDeployer(address deployer) view returns(address[])
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) GetSafesByDeployer(opts *bind.CallOpts, deployer common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "getSafesByDeployer", deployer)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSafesByDeployer is a free data retrieval call binding the contract method 0x1eae53e8.
//
// Solidity: function getSafesByDeployer(address deployer) view returns(address[])
func (_SafeTimelockFactory *SafeTimelockFactorySession) GetSafesByDeployer(deployer common.Address) ([]common.Address, error) {
	return _SafeTimelockFactory.Contract.GetSafesByDeployer(&_SafeTimelockFactory.CallOpts, deployer)
}

// GetSafesByDeployer is a free data retrieval call binding the contract method 0x1eae53e8.
//
// Solidity: function getSafesByDeployer(address deployer) view returns(address[])
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) GetSafesByDeployer(deployer common.Address) ([]common.Address, error) {
	return _SafeTimelockFactory.Contract.GetSafesByDeployer(&_SafeTimelockFactory.CallOpts, deployer)
}

// GetTimelocksByDeployer is a free data retrieval call binding the contract method 0xb145953e.
//
// Solidity: function getTimelocksByDeployer(address deployer) view returns(address[])
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) GetTimelocksByDeployer(opts *bind.CallOpts, deployer common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "getTimelocksByDeployer", deployer)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTimelocksByDeployer is a free data retrieval call binding the contract method 0xb145953e.
//
// Solidity: function getTimelocksByDeployer(address deployer) view returns(address[])
func (_SafeTimelockFactory *SafeTimelockFactorySession) GetTimelocksByDeployer(deployer common.Address) ([]common.Address, error) {
	return _SafeTimelockFactory.Contract.GetTimelocksByDeployer(&_SafeTimelockFactory.CallOpts, deployer)
}

// GetTimelocksByDeployer is a free data retrieval call binding the contract method 0xb145953e.
//
// Solidity: function getTimelocksByDeployer(address deployer) view returns(address[])
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) GetTimelocksByDeployer(deployer common.Address) ([]common.Address, error) {
	return _SafeTimelockFactory.Contract.GetTimelocksByDeployer(&_SafeTimelockFactory.CallOpts, deployer)
}

// IsSafe is a free data retrieval call binding the contract method 0x769a28ac.
//
// Solidity: function isSafe(address safe) view returns(bool)
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) IsSafe(opts *bind.CallOpts, safe common.Address) (bool, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "isSafe", safe)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSafe is a free data retrieval call binding the contract method 0x769a28ac.
//
// Solidity: function isSafe(address safe) view returns(bool)
func (_SafeTimelockFactory *SafeTimelockFactorySession) IsSafe(safe common.Address) (bool, error) {
	return _SafeTimelockFactory.Contract.IsSafe(&_SafeTimelockFactory.CallOpts, safe)
}

// IsSafe is a free data retrieval call binding the contract method 0x769a28ac.
//
// Solidity: function isSafe(address safe) view returns(bool)
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) IsSafe(safe common.Address) (bool, error) {
	return _SafeTimelockFactory.Contract.IsSafe(&_SafeTimelockFactory.CallOpts, safe)
}

// IsTimelock is a free data retrieval call binding the contract method 0x9e93963c.
//
// Solidity: function isTimelock(address timelock) view returns(bool)
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) IsTimelock(opts *bind.CallOpts, timelock common.Address) (bool, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "isTimelock", timelock)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimelock is a free data retrieval call binding the contract method 0x9e93963c.
//
// Solidity: function isTimelock(address timelock) view returns(bool)
func (_SafeTimelockFactory *SafeTimelockFactorySession) IsTimelock(timelock common.Address) (bool, error) {
	return _SafeTimelockFactory.Contract.IsTimelock(&_SafeTimelockFactory.CallOpts, timelock)
}

// IsTimelock is a free data retrieval call binding the contract method 0x9e93963c.
//
// Solidity: function isTimelock(address timelock) view returns(bool)
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) IsTimelock(timelock common.Address) (bool, error) {
	return _SafeTimelockFactory.Contract.IsTimelock(&_SafeTimelockFactory.CallOpts, timelock)
}

// SafeFallbackHandler is a free data retrieval call binding the contract method 0xa03bad5a.
//
// Solidity: function safeFallbackHandler() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) SafeFallbackHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "safeFallbackHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeFallbackHandler is a free data retrieval call binding the contract method 0xa03bad5a.
//
// Solidity: function safeFallbackHandler() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactorySession) SafeFallbackHandler() (common.Address, error) {
	return _SafeTimelockFactory.Contract.SafeFallbackHandler(&_SafeTimelockFactory.CallOpts)
}

// SafeFallbackHandler is a free data retrieval call binding the contract method 0xa03bad5a.
//
// Solidity: function safeFallbackHandler() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) SafeFallbackHandler() (common.Address, error) {
	return _SafeTimelockFactory.Contract.SafeFallbackHandler(&_SafeTimelockFactory.CallOpts)
}

// SafeProxyFactory is a free data retrieval call binding the contract method 0x19964501.
//
// Solidity: function safeProxyFactory() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) SafeProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "safeProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeProxyFactory is a free data retrieval call binding the contract method 0x19964501.
//
// Solidity: function safeProxyFactory() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactorySession) SafeProxyFactory() (common.Address, error) {
	return _SafeTimelockFactory.Contract.SafeProxyFactory(&_SafeTimelockFactory.CallOpts)
}

// SafeProxyFactory is a free data retrieval call binding the contract method 0x19964501.
//
// Solidity: function safeProxyFactory() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) SafeProxyFactory() (common.Address, error) {
	return _SafeTimelockFactory.Contract.SafeProxyFactory(&_SafeTimelockFactory.CallOpts)
}

// SafeSingleton is a free data retrieval call binding the contract method 0xac7d146b.
//
// Solidity: function safeSingleton() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) SafeSingleton(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "safeSingleton")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeSingleton is a free data retrieval call binding the contract method 0xac7d146b.
//
// Solidity: function safeSingleton() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactorySession) SafeSingleton() (common.Address, error) {
	return _SafeTimelockFactory.Contract.SafeSingleton(&_SafeTimelockFactory.CallOpts)
}

// SafeSingleton is a free data retrieval call binding the contract method 0xac7d146b.
//
// Solidity: function safeSingleton() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) SafeSingleton() (common.Address, error) {
	return _SafeTimelockFactory.Contract.SafeSingleton(&_SafeTimelockFactory.CallOpts)
}

// TimelockImplementation is a free data retrieval call binding the contract method 0xc5e8f3e5.
//
// Solidity: function timelockImplementation() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCaller) TimelockImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SafeTimelockFactory.contract.Call(opts, &out, "timelockImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TimelockImplementation is a free data retrieval call binding the contract method 0xc5e8f3e5.
//
// Solidity: function timelockImplementation() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactorySession) TimelockImplementation() (common.Address, error) {
	return _SafeTimelockFactory.Contract.TimelockImplementation(&_SafeTimelockFactory.CallOpts)
}

// TimelockImplementation is a free data retrieval call binding the contract method 0xc5e8f3e5.
//
// Solidity: function timelockImplementation() view returns(address)
func (_SafeTimelockFactory *SafeTimelockFactoryCallerSession) TimelockImplementation() (common.Address, error) {
	return _SafeTimelockFactory.Contract.TimelockImplementation(&_SafeTimelockFactory.CallOpts)
}

// DeploySafe is a paid mutator transaction binding the contract method 0xd15a64ce.
//
// Solidity: function deploySafe((address[],uint256) config, bytes32 salt) returns(address safe)
func (_SafeTimelockFactory *SafeTimelockFactoryTransactor) DeploySafe(opts *bind.TransactOpts, config ISafeTimelockFactorySafeConfig, salt [32]byte) (*types.Transaction, error) {
	return _SafeTimelockFactory.contract.Transact(opts, "deploySafe", config, salt)
}

// DeploySafe is a paid mutator transaction binding the contract method 0xd15a64ce.
//
// Solidity: function deploySafe((address[],uint256) config, bytes32 salt) returns(address safe)
func (_SafeTimelockFactory *SafeTimelockFactorySession) DeploySafe(config ISafeTimelockFactorySafeConfig, salt [32]byte) (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.DeploySafe(&_SafeTimelockFactory.TransactOpts, config, salt)
}

// DeploySafe is a paid mutator transaction binding the contract method 0xd15a64ce.
//
// Solidity: function deploySafe((address[],uint256) config, bytes32 salt) returns(address safe)
func (_SafeTimelockFactory *SafeTimelockFactoryTransactorSession) DeploySafe(config ISafeTimelockFactorySafeConfig, salt [32]byte) (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.DeploySafe(&_SafeTimelockFactory.TransactOpts, config, salt)
}

// DeployTimelock is a paid mutator transaction binding the contract method 0xb7638897.
//
// Solidity: function deployTimelock((uint256,address[],address[]) config, bytes32 salt) returns(address timelock)
func (_SafeTimelockFactory *SafeTimelockFactoryTransactor) DeployTimelock(opts *bind.TransactOpts, config ISafeTimelockFactoryTimelockConfig, salt [32]byte) (*types.Transaction, error) {
	return _SafeTimelockFactory.contract.Transact(opts, "deployTimelock", config, salt)
}

// DeployTimelock is a paid mutator transaction binding the contract method 0xb7638897.
//
// Solidity: function deployTimelock((uint256,address[],address[]) config, bytes32 salt) returns(address timelock)
func (_SafeTimelockFactory *SafeTimelockFactorySession) DeployTimelock(config ISafeTimelockFactoryTimelockConfig, salt [32]byte) (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.DeployTimelock(&_SafeTimelockFactory.TransactOpts, config, salt)
}

// DeployTimelock is a paid mutator transaction binding the contract method 0xb7638897.
//
// Solidity: function deployTimelock((uint256,address[],address[]) config, bytes32 salt) returns(address timelock)
func (_SafeTimelockFactory *SafeTimelockFactoryTransactorSession) DeployTimelock(config ISafeTimelockFactoryTimelockConfig, salt [32]byte) (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.DeployTimelock(&_SafeTimelockFactory.TransactOpts, config, salt)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SafeTimelockFactory *SafeTimelockFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeTimelockFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SafeTimelockFactory *SafeTimelockFactorySession) Initialize() (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.Initialize(&_SafeTimelockFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SafeTimelockFactory *SafeTimelockFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _SafeTimelockFactory.Contract.Initialize(&_SafeTimelockFactory.TransactOpts)
}

// SafeTimelockFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryInitializedIterator struct {
	Event *SafeTimelockFactoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeTimelockFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeTimelockFactoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeTimelockFactoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeTimelockFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeTimelockFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeTimelockFactoryInitialized represents a Initialized event raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*SafeTimelockFactoryInitializedIterator, error) {

	logs, sub, err := _SafeTimelockFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SafeTimelockFactoryInitializedIterator{contract: _SafeTimelockFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SafeTimelockFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _SafeTimelockFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeTimelockFactoryInitialized)
				if err := _SafeTimelockFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) ParseInitialized(log types.Log) (*SafeTimelockFactoryInitialized, error) {
	event := new(SafeTimelockFactoryInitialized)
	if err := _SafeTimelockFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeTimelockFactorySafeDeployedIterator is returned from FilterSafeDeployed and is used to iterate over the raw logs and unpacked data for SafeDeployed events raised by the SafeTimelockFactory contract.
type SafeTimelockFactorySafeDeployedIterator struct {
	Event *SafeTimelockFactorySafeDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeTimelockFactorySafeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeTimelockFactorySafeDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeTimelockFactorySafeDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeTimelockFactorySafeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeTimelockFactorySafeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeTimelockFactorySafeDeployed represents a SafeDeployed event raised by the SafeTimelockFactory contract.
type SafeTimelockFactorySafeDeployed struct {
	Deployer  common.Address
	Safe      common.Address
	Owners    []common.Address
	Threshold *big.Int
	Salt      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSafeDeployed is a free log retrieval operation binding the contract event 0x1f1cf52f43299bdf7c10abfb65dc19bb454f2e61c2d5f2de2c1d6efbfd1e7a69.
//
// Solidity: event SafeDeployed(address indexed deployer, address indexed safe, address[] owners, uint256 threshold, bytes32 salt)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) FilterSafeDeployed(opts *bind.FilterOpts, deployer []common.Address, safe []common.Address) (*SafeTimelockFactorySafeDeployedIterator, error) {

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}
	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _SafeTimelockFactory.contract.FilterLogs(opts, "SafeDeployed", deployerRule, safeRule)
	if err != nil {
		return nil, err
	}
	return &SafeTimelockFactorySafeDeployedIterator{contract: _SafeTimelockFactory.contract, event: "SafeDeployed", logs: logs, sub: sub}, nil
}

// WatchSafeDeployed is a free log subscription operation binding the contract event 0x1f1cf52f43299bdf7c10abfb65dc19bb454f2e61c2d5f2de2c1d6efbfd1e7a69.
//
// Solidity: event SafeDeployed(address indexed deployer, address indexed safe, address[] owners, uint256 threshold, bytes32 salt)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) WatchSafeDeployed(opts *bind.WatchOpts, sink chan<- *SafeTimelockFactorySafeDeployed, deployer []common.Address, safe []common.Address) (event.Subscription, error) {

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}
	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _SafeTimelockFactory.contract.WatchLogs(opts, "SafeDeployed", deployerRule, safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeTimelockFactorySafeDeployed)
				if err := _SafeTimelockFactory.contract.UnpackLog(event, "SafeDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSafeDeployed is a log parse operation binding the contract event 0x1f1cf52f43299bdf7c10abfb65dc19bb454f2e61c2d5f2de2c1d6efbfd1e7a69.
//
// Solidity: event SafeDeployed(address indexed deployer, address indexed safe, address[] owners, uint256 threshold, bytes32 salt)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) ParseSafeDeployed(log types.Log) (*SafeTimelockFactorySafeDeployed, error) {
	event := new(SafeTimelockFactorySafeDeployed)
	if err := _SafeTimelockFactory.contract.UnpackLog(event, "SafeDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeTimelockFactoryTimelockDeployedIterator is returned from FilterTimelockDeployed and is used to iterate over the raw logs and unpacked data for TimelockDeployed events raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryTimelockDeployedIterator struct {
	Event *SafeTimelockFactoryTimelockDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeTimelockFactoryTimelockDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeTimelockFactoryTimelockDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeTimelockFactoryTimelockDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeTimelockFactoryTimelockDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeTimelockFactoryTimelockDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeTimelockFactoryTimelockDeployed represents a TimelockDeployed event raised by the SafeTimelockFactory contract.
type SafeTimelockFactoryTimelockDeployed struct {
	Deployer  common.Address
	Timelock  common.Address
	MinDelay  *big.Int
	Proposers []common.Address
	Executors []common.Address
	Salt      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockDeployed is a free log retrieval operation binding the contract event 0xfc86ac6f4767b3d388def1785cb7b0d7e671696f3261dc8942332f779b9cde52.
//
// Solidity: event TimelockDeployed(address indexed deployer, address indexed timelock, uint256 minDelay, address[] proposers, address[] executors, bytes32 salt)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) FilterTimelockDeployed(opts *bind.FilterOpts, deployer []common.Address, timelock []common.Address) (*SafeTimelockFactoryTimelockDeployedIterator, error) {

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}
	var timelockRule []interface{}
	for _, timelockItem := range timelock {
		timelockRule = append(timelockRule, timelockItem)
	}

	logs, sub, err := _SafeTimelockFactory.contract.FilterLogs(opts, "TimelockDeployed", deployerRule, timelockRule)
	if err != nil {
		return nil, err
	}
	return &SafeTimelockFactoryTimelockDeployedIterator{contract: _SafeTimelockFactory.contract, event: "TimelockDeployed", logs: logs, sub: sub}, nil
}

// WatchTimelockDeployed is a free log subscription operation binding the contract event 0xfc86ac6f4767b3d388def1785cb7b0d7e671696f3261dc8942332f779b9cde52.
//
// Solidity: event TimelockDeployed(address indexed deployer, address indexed timelock, uint256 minDelay, address[] proposers, address[] executors, bytes32 salt)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) WatchTimelockDeployed(opts *bind.WatchOpts, sink chan<- *SafeTimelockFactoryTimelockDeployed, deployer []common.Address, timelock []common.Address) (event.Subscription, error) {

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}
	var timelockRule []interface{}
	for _, timelockItem := range timelock {
		timelockRule = append(timelockRule, timelockItem)
	}

	logs, sub, err := _SafeTimelockFactory.contract.WatchLogs(opts, "TimelockDeployed", deployerRule, timelockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeTimelockFactoryTimelockDeployed)
				if err := _SafeTimelockFactory.contract.UnpackLog(event, "TimelockDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockDeployed is a log parse operation binding the contract event 0xfc86ac6f4767b3d388def1785cb7b0d7e671696f3261dc8942332f779b9cde52.
//
// Solidity: event TimelockDeployed(address indexed deployer, address indexed timelock, uint256 minDelay, address[] proposers, address[] executors, bytes32 salt)
func (_SafeTimelockFactory *SafeTimelockFactoryFilterer) ParseTimelockDeployed(log types.Log) (*SafeTimelockFactoryTimelockDeployed, error) {
	event := new(SafeTimelockFactoryTimelockDeployed)
	if err := _SafeTimelockFactory.contract.UnpackLog(event, "TimelockDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
