// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ComputeOperator

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

// ComputeOperatorMetaData contains all meta data concerning the ComputeOperator contract.
var ComputeOperatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_appController\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_computeAVSRegistrar\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_strategyToSlash\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SLASHING_OPERATORSET_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeAVSRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"operatorMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"performInitialAllocation\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyToSlash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InitialAllocationAlreadyComplete\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAppController\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotComputeAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// ComputeOperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ComputeOperatorMetaData.ABI instead.
var ComputeOperatorABI = ComputeOperatorMetaData.ABI

// ComputeOperator is an auto generated Go binding around an Ethereum contract.
type ComputeOperator struct {
	ComputeOperatorCaller     // Read-only binding to the contract
	ComputeOperatorTransactor // Write-only binding to the contract
	ComputeOperatorFilterer   // Log filterer for contract events
}

// ComputeOperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComputeOperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeOperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComputeOperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeOperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComputeOperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeOperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComputeOperatorSession struct {
	Contract     *ComputeOperator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComputeOperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComputeOperatorCallerSession struct {
	Contract *ComputeOperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ComputeOperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComputeOperatorTransactorSession struct {
	Contract     *ComputeOperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ComputeOperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComputeOperatorRaw struct {
	Contract *ComputeOperator // Generic contract binding to access the raw methods on
}

// ComputeOperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComputeOperatorCallerRaw struct {
	Contract *ComputeOperatorCaller // Generic read-only contract binding to access the raw methods on
}

// ComputeOperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComputeOperatorTransactorRaw struct {
	Contract *ComputeOperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComputeOperator creates a new instance of ComputeOperator, bound to a specific deployed contract.
func NewComputeOperator(address common.Address, backend bind.ContractBackend) (*ComputeOperator, error) {
	contract, err := bindComputeOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComputeOperator{ComputeOperatorCaller: ComputeOperatorCaller{contract: contract}, ComputeOperatorTransactor: ComputeOperatorTransactor{contract: contract}, ComputeOperatorFilterer: ComputeOperatorFilterer{contract: contract}}, nil
}

// NewComputeOperatorCaller creates a new read-only instance of ComputeOperator, bound to a specific deployed contract.
func NewComputeOperatorCaller(address common.Address, caller bind.ContractCaller) (*ComputeOperatorCaller, error) {
	contract, err := bindComputeOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComputeOperatorCaller{contract: contract}, nil
}

// NewComputeOperatorTransactor creates a new write-only instance of ComputeOperator, bound to a specific deployed contract.
func NewComputeOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ComputeOperatorTransactor, error) {
	contract, err := bindComputeOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComputeOperatorTransactor{contract: contract}, nil
}

// NewComputeOperatorFilterer creates a new log filterer instance of ComputeOperator, bound to a specific deployed contract.
func NewComputeOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ComputeOperatorFilterer, error) {
	contract, err := bindComputeOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComputeOperatorFilterer{contract: contract}, nil
}

// bindComputeOperator binds a generic wrapper to an already deployed contract.
func bindComputeOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ComputeOperatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComputeOperator *ComputeOperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComputeOperator.Contract.ComputeOperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComputeOperator *ComputeOperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeOperator.Contract.ComputeOperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComputeOperator *ComputeOperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComputeOperator.Contract.ComputeOperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComputeOperator *ComputeOperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComputeOperator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComputeOperator *ComputeOperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeOperator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComputeOperator *ComputeOperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComputeOperator.Contract.contract.Transact(opts, method, params...)
}

// SLASHINGOPERATORSETID is a free data retrieval call binding the contract method 0x96b313c8.
//
// Solidity: function SLASHING_OPERATORSET_ID() view returns(uint32)
func (_ComputeOperator *ComputeOperatorCaller) SLASHINGOPERATORSETID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ComputeOperator.contract.Call(opts, &out, "SLASHING_OPERATORSET_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SLASHINGOPERATORSETID is a free data retrieval call binding the contract method 0x96b313c8.
//
// Solidity: function SLASHING_OPERATORSET_ID() view returns(uint32)
func (_ComputeOperator *ComputeOperatorSession) SLASHINGOPERATORSETID() (uint32, error) {
	return _ComputeOperator.Contract.SLASHINGOPERATORSETID(&_ComputeOperator.CallOpts)
}

// SLASHINGOPERATORSETID is a free data retrieval call binding the contract method 0x96b313c8.
//
// Solidity: function SLASHING_OPERATORSET_ID() view returns(uint32)
func (_ComputeOperator *ComputeOperatorCallerSession) SLASHINGOPERATORSETID() (uint32, error) {
	return _ComputeOperator.Contract.SLASHINGOPERATORSETID(&_ComputeOperator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ComputeOperator *ComputeOperatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeOperator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ComputeOperator *ComputeOperatorSession) AllocationManager() (common.Address, error) {
	return _ComputeOperator.Contract.AllocationManager(&_ComputeOperator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ComputeOperator *ComputeOperatorCallerSession) AllocationManager() (common.Address, error) {
	return _ComputeOperator.Contract.AllocationManager(&_ComputeOperator.CallOpts)
}

// AppController is a free data retrieval call binding the contract method 0x7ede3ba7.
//
// Solidity: function appController() view returns(address)
func (_ComputeOperator *ComputeOperatorCaller) AppController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeOperator.contract.Call(opts, &out, "appController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppController is a free data retrieval call binding the contract method 0x7ede3ba7.
//
// Solidity: function appController() view returns(address)
func (_ComputeOperator *ComputeOperatorSession) AppController() (common.Address, error) {
	return _ComputeOperator.Contract.AppController(&_ComputeOperator.CallOpts)
}

// AppController is a free data retrieval call binding the contract method 0x7ede3ba7.
//
// Solidity: function appController() view returns(address)
func (_ComputeOperator *ComputeOperatorCallerSession) AppController() (common.Address, error) {
	return _ComputeOperator.Contract.AppController(&_ComputeOperator.CallOpts)
}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_ComputeOperator *ComputeOperatorCaller) ComputeAVSRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeOperator.contract.Call(opts, &out, "computeAVSRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_ComputeOperator *ComputeOperatorSession) ComputeAVSRegistrar() (common.Address, error) {
	return _ComputeOperator.Contract.ComputeAVSRegistrar(&_ComputeOperator.CallOpts)
}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_ComputeOperator *ComputeOperatorCallerSession) ComputeAVSRegistrar() (common.Address, error) {
	return _ComputeOperator.Contract.ComputeAVSRegistrar(&_ComputeOperator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ComputeOperator *ComputeOperatorCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeOperator.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ComputeOperator *ComputeOperatorSession) DelegationManager() (common.Address, error) {
	return _ComputeOperator.Contract.DelegationManager(&_ComputeOperator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ComputeOperator *ComputeOperatorCallerSession) DelegationManager() (common.Address, error) {
	return _ComputeOperator.Contract.DelegationManager(&_ComputeOperator.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeOperator *ComputeOperatorCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeOperator.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeOperator *ComputeOperatorSession) PermissionController() (common.Address, error) {
	return _ComputeOperator.Contract.PermissionController(&_ComputeOperator.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeOperator *ComputeOperatorCallerSession) PermissionController() (common.Address, error) {
	return _ComputeOperator.Contract.PermissionController(&_ComputeOperator.CallOpts)
}

// StrategyToSlash is a free data retrieval call binding the contract method 0x5ac080f1.
//
// Solidity: function strategyToSlash() view returns(address)
func (_ComputeOperator *ComputeOperatorCaller) StrategyToSlash(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeOperator.contract.Call(opts, &out, "strategyToSlash")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyToSlash is a free data retrieval call binding the contract method 0x5ac080f1.
//
// Solidity: function strategyToSlash() view returns(address)
func (_ComputeOperator *ComputeOperatorSession) StrategyToSlash() (common.Address, error) {
	return _ComputeOperator.Contract.StrategyToSlash(&_ComputeOperator.CallOpts)
}

// StrategyToSlash is a free data retrieval call binding the contract method 0x5ac080f1.
//
// Solidity: function strategyToSlash() view returns(address)
func (_ComputeOperator *ComputeOperatorCallerSession) StrategyToSlash() (common.Address, error) {
	return _ComputeOperator.Contract.StrategyToSlash(&_ComputeOperator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeOperator *ComputeOperatorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ComputeOperator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeOperator *ComputeOperatorSession) Version() (string, error) {
	return _ComputeOperator.Contract.Version(&_ComputeOperator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeOperator *ComputeOperatorCallerSession) Version() (string, error) {
	return _ComputeOperator.Contract.Version(&_ComputeOperator.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string operatorMetadataURI) returns()
func (_ComputeOperator *ComputeOperatorTransactor) Initialize(opts *bind.TransactOpts, operatorMetadataURI string) (*types.Transaction, error) {
	return _ComputeOperator.contract.Transact(opts, "initialize", operatorMetadataURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string operatorMetadataURI) returns()
func (_ComputeOperator *ComputeOperatorSession) Initialize(operatorMetadataURI string) (*types.Transaction, error) {
	return _ComputeOperator.Contract.Initialize(&_ComputeOperator.TransactOpts, operatorMetadataURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string operatorMetadataURI) returns()
func (_ComputeOperator *ComputeOperatorTransactorSession) Initialize(operatorMetadataURI string) (*types.Transaction, error) {
	return _ComputeOperator.Contract.Initialize(&_ComputeOperator.TransactOpts, operatorMetadataURI)
}

// PerformInitialAllocation is a paid mutator transaction binding the contract method 0x33cd0d13.
//
// Solidity: function performInitialAllocation() returns()
func (_ComputeOperator *ComputeOperatorTransactor) PerformInitialAllocation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeOperator.contract.Transact(opts, "performInitialAllocation")
}

// PerformInitialAllocation is a paid mutator transaction binding the contract method 0x33cd0d13.
//
// Solidity: function performInitialAllocation() returns()
func (_ComputeOperator *ComputeOperatorSession) PerformInitialAllocation() (*types.Transaction, error) {
	return _ComputeOperator.Contract.PerformInitialAllocation(&_ComputeOperator.TransactOpts)
}

// PerformInitialAllocation is a paid mutator transaction binding the contract method 0x33cd0d13.
//
// Solidity: function performInitialAllocation() returns()
func (_ComputeOperator *ComputeOperatorTransactorSession) PerformInitialAllocation() (*types.Transaction, error) {
	return _ComputeOperator.Contract.PerformInitialAllocation(&_ComputeOperator.TransactOpts)
}

// RegisterForOperatorSet is a paid mutator transaction binding the contract method 0x4ca62978.
//
// Solidity: function registerForOperatorSet(uint32 operatorSetId) returns()
func (_ComputeOperator *ComputeOperatorTransactor) RegisterForOperatorSet(opts *bind.TransactOpts, operatorSetId uint32) (*types.Transaction, error) {
	return _ComputeOperator.contract.Transact(opts, "registerForOperatorSet", operatorSetId)
}

// RegisterForOperatorSet is a paid mutator transaction binding the contract method 0x4ca62978.
//
// Solidity: function registerForOperatorSet(uint32 operatorSetId) returns()
func (_ComputeOperator *ComputeOperatorSession) RegisterForOperatorSet(operatorSetId uint32) (*types.Transaction, error) {
	return _ComputeOperator.Contract.RegisterForOperatorSet(&_ComputeOperator.TransactOpts, operatorSetId)
}

// RegisterForOperatorSet is a paid mutator transaction binding the contract method 0x4ca62978.
//
// Solidity: function registerForOperatorSet(uint32 operatorSetId) returns()
func (_ComputeOperator *ComputeOperatorTransactorSession) RegisterForOperatorSet(operatorSetId uint32) (*types.Transaction, error) {
	return _ComputeOperator.Contract.RegisterForOperatorSet(&_ComputeOperator.TransactOpts, operatorSetId)
}

// ComputeOperatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ComputeOperator contract.
type ComputeOperatorInitializedIterator struct {
	Event *ComputeOperatorInitialized // Event containing the contract specifics and raw log

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
func (it *ComputeOperatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeOperatorInitialized)
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
		it.Event = new(ComputeOperatorInitialized)
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
func (it *ComputeOperatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeOperatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeOperatorInitialized represents a Initialized event raised by the ComputeOperator contract.
type ComputeOperatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ComputeOperator *ComputeOperatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ComputeOperatorInitializedIterator, error) {

	logs, sub, err := _ComputeOperator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ComputeOperatorInitializedIterator{contract: _ComputeOperator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ComputeOperator *ComputeOperatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ComputeOperatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ComputeOperator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeOperatorInitialized)
				if err := _ComputeOperator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ComputeOperator *ComputeOperatorFilterer) ParseInitialized(log types.Log) (*ComputeOperatorInitialized, error) {
	event := new(ComputeOperatorInitialized)
	if err := _ComputeOperator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
