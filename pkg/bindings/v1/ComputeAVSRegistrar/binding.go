// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ComputeAVSRegistrar

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ComputeAVSRegistrarMetaData contains all meta data concerning the ComputeAVSRegistrar contract.
var ComputeAVSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_appController\",\"type\":\"address\",\"internalType\":\"contractIAppController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorToAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAppController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assignOperatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedOperators\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorAllowed\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextOperatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorFromAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAppController\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyInAllowlist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotInAllowlist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetIdNotAssigned\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// ComputeAVSRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use ComputeAVSRegistrarMetaData.ABI instead.
var ComputeAVSRegistrarABI = ComputeAVSRegistrarMetaData.ABI

// ComputeAVSRegistrar is an auto generated Go binding around an Ethereum contract.
type ComputeAVSRegistrar struct {
	ComputeAVSRegistrarCaller     // Read-only binding to the contract
	ComputeAVSRegistrarTransactor // Write-only binding to the contract
	ComputeAVSRegistrarFilterer   // Log filterer for contract events
}

// ComputeAVSRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComputeAVSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeAVSRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComputeAVSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeAVSRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComputeAVSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputeAVSRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComputeAVSRegistrarSession struct {
	Contract     *ComputeAVSRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ComputeAVSRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComputeAVSRegistrarCallerSession struct {
	Contract *ComputeAVSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ComputeAVSRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComputeAVSRegistrarTransactorSession struct {
	Contract     *ComputeAVSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ComputeAVSRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComputeAVSRegistrarRaw struct {
	Contract *ComputeAVSRegistrar // Generic contract binding to access the raw methods on
}

// ComputeAVSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComputeAVSRegistrarCallerRaw struct {
	Contract *ComputeAVSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// ComputeAVSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComputeAVSRegistrarTransactorRaw struct {
	Contract *ComputeAVSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComputeAVSRegistrar creates a new instance of ComputeAVSRegistrar, bound to a specific deployed contract.
func NewComputeAVSRegistrar(address common.Address, backend bind.ContractBackend) (*ComputeAVSRegistrar, error) {
	contract, err := bindComputeAVSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrar{ComputeAVSRegistrarCaller: ComputeAVSRegistrarCaller{contract: contract}, ComputeAVSRegistrarTransactor: ComputeAVSRegistrarTransactor{contract: contract}, ComputeAVSRegistrarFilterer: ComputeAVSRegistrarFilterer{contract: contract}}, nil
}

// NewComputeAVSRegistrarCaller creates a new read-only instance of ComputeAVSRegistrar, bound to a specific deployed contract.
func NewComputeAVSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*ComputeAVSRegistrarCaller, error) {
	contract, err := bindComputeAVSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarCaller{contract: contract}, nil
}

// NewComputeAVSRegistrarTransactor creates a new write-only instance of ComputeAVSRegistrar, bound to a specific deployed contract.
func NewComputeAVSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*ComputeAVSRegistrarTransactor, error) {
	contract, err := bindComputeAVSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarTransactor{contract: contract}, nil
}

// NewComputeAVSRegistrarFilterer creates a new log filterer instance of ComputeAVSRegistrar, bound to a specific deployed contract.
func NewComputeAVSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*ComputeAVSRegistrarFilterer, error) {
	contract, err := bindComputeAVSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarFilterer{contract: contract}, nil
}

// bindComputeAVSRegistrar binds a generic wrapper to an already deployed contract.
func bindComputeAVSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ComputeAVSRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComputeAVSRegistrar *ComputeAVSRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComputeAVSRegistrar.Contract.ComputeAVSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComputeAVSRegistrar *ComputeAVSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.ComputeAVSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComputeAVSRegistrar *ComputeAVSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.ComputeAVSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComputeAVSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) AllocationManager() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.AllocationManager(&_ComputeAVSRegistrar.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) AllocationManager() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.AllocationManager(&_ComputeAVSRegistrar.CallOpts)
}

// AppController is a free data retrieval call binding the contract method 0x7ede3ba7.
//
// Solidity: function appController() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) AppController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "appController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppController is a free data retrieval call binding the contract method 0x7ede3ba7.
//
// Solidity: function appController() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) AppController() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.AppController(&_ComputeAVSRegistrar.CallOpts)
}

// AppController is a free data retrieval call binding the contract method 0x7ede3ba7.
//
// Solidity: function appController() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) AppController() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.AppController(&_ComputeAVSRegistrar.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) Avs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "avs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) Avs() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.Avs(&_ComputeAVSRegistrar.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) Avs() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.Avs(&_ComputeAVSRegistrar.CallOpts)
}

// GetAllowedOperators is a free data retrieval call binding the contract method 0x7fe94e16.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) GetAllowedOperators(opts *bind.CallOpts, operatorSet OperatorSet) ([]common.Address, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "getAllowedOperators", operatorSet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedOperators is a free data retrieval call binding the contract method 0x7fe94e16.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) GetAllowedOperators(operatorSet OperatorSet) ([]common.Address, error) {
	return _ComputeAVSRegistrar.Contract.GetAllowedOperators(&_ComputeAVSRegistrar.CallOpts, operatorSet)
}

// GetAllowedOperators is a free data retrieval call binding the contract method 0x7fe94e16.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) GetAllowedOperators(operatorSet OperatorSet) ([]common.Address, error) {
	return _ComputeAVSRegistrar.Contract.GetAllowedOperators(&_ComputeAVSRegistrar.CallOpts, operatorSet)
}

// IsOperatorAllowed is a free data retrieval call binding the contract method 0xf91ff80c.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) IsOperatorAllowed(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "isOperatorAllowed", operatorSet, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorAllowed is a free data retrieval call binding the contract method 0xf91ff80c.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) IsOperatorAllowed(operatorSet OperatorSet, operator common.Address) (bool, error) {
	return _ComputeAVSRegistrar.Contract.IsOperatorAllowed(&_ComputeAVSRegistrar.CallOpts, operatorSet, operator)
}

// IsOperatorAllowed is a free data retrieval call binding the contract method 0xf91ff80c.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) IsOperatorAllowed(operatorSet OperatorSet, operator common.Address) (bool, error) {
	return _ComputeAVSRegistrar.Contract.IsOperatorAllowed(&_ComputeAVSRegistrar.CallOpts, operatorSet, operator)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) KeyRegistrar() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.KeyRegistrar(&_ComputeAVSRegistrar.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) KeyRegistrar() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.KeyRegistrar(&_ComputeAVSRegistrar.CallOpts)
}

// NextOperatorSetId is a free data retrieval call binding the contract method 0x833b3fa3.
//
// Solidity: function nextOperatorSetId() view returns(uint32)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) NextOperatorSetId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "nextOperatorSetId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextOperatorSetId is a free data retrieval call binding the contract method 0x833b3fa3.
//
// Solidity: function nextOperatorSetId() view returns(uint32)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) NextOperatorSetId() (uint32, error) {
	return _ComputeAVSRegistrar.Contract.NextOperatorSetId(&_ComputeAVSRegistrar.CallOpts)
}

// NextOperatorSetId is a free data retrieval call binding the contract method 0x833b3fa3.
//
// Solidity: function nextOperatorSetId() view returns(uint32)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) NextOperatorSetId() (uint32, error) {
	return _ComputeAVSRegistrar.Contract.NextOperatorSetId(&_ComputeAVSRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) Owner() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.Owner(&_ComputeAVSRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) Owner() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.Owner(&_ComputeAVSRegistrar.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) PermissionController() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.PermissionController(&_ComputeAVSRegistrar.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) PermissionController() (common.Address, error) {
	return _ComputeAVSRegistrar.Contract.PermissionController(&_ComputeAVSRegistrar.CallOpts)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) SupportsAVS(opts *bind.CallOpts, _avs common.Address) (bool, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "supportsAVS", _avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _ComputeAVSRegistrar.Contract.SupportsAVS(&_ComputeAVSRegistrar.CallOpts, _avs)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _ComputeAVSRegistrar.Contract.SupportsAVS(&_ComputeAVSRegistrar.CallOpts, _avs)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ComputeAVSRegistrar.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) Version() (string, error) {
	return _ComputeAVSRegistrar.Contract.Version(&_ComputeAVSRegistrar.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarCallerSession) Version() (string, error) {
	return _ComputeAVSRegistrar.Contract.Version(&_ComputeAVSRegistrar.CallOpts)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x1017873a.
//
// Solidity: function addOperatorToAllowlist((address,uint32) operatorSet, address operator) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) AddOperatorToAllowlist(opts *bind.TransactOpts, operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "addOperatorToAllowlist", operatorSet, operator)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x1017873a.
//
// Solidity: function addOperatorToAllowlist((address,uint32) operatorSet, address operator) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) AddOperatorToAllowlist(operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.AddOperatorToAllowlist(&_ComputeAVSRegistrar.TransactOpts, operatorSet, operator)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x1017873a.
//
// Solidity: function addOperatorToAllowlist((address,uint32) operatorSet, address operator) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) AddOperatorToAllowlist(operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.AddOperatorToAllowlist(&_ComputeAVSRegistrar.TransactOpts, operatorSet, operator)
}

// AssignOperatorSetId is a paid mutator transaction binding the contract method 0x3a10eef7.
//
// Solidity: function assignOperatorSetId() returns(uint32 operatorSetId)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) AssignOperatorSetId(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "assignOperatorSetId")
}

// AssignOperatorSetId is a paid mutator transaction binding the contract method 0x3a10eef7.
//
// Solidity: function assignOperatorSetId() returns(uint32 operatorSetId)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) AssignOperatorSetId() (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.AssignOperatorSetId(&_ComputeAVSRegistrar.TransactOpts)
}

// AssignOperatorSetId is a paid mutator transaction binding the contract method 0x3a10eef7.
//
// Solidity: function assignOperatorSetId() returns(uint32 operatorSetId)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) AssignOperatorSetId() (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.AssignOperatorSetId(&_ComputeAVSRegistrar.TransactOpts)
}

// CreateOperatorSet is a paid mutator transaction binding the contract method 0x74da4713.
//
// Solidity: function createOperatorSet(uint32 operatorSetId) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) CreateOperatorSet(opts *bind.TransactOpts, operatorSetId uint32) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "createOperatorSet", operatorSetId)
}

// CreateOperatorSet is a paid mutator transaction binding the contract method 0x74da4713.
//
// Solidity: function createOperatorSet(uint32 operatorSetId) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) CreateOperatorSet(operatorSetId uint32) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.CreateOperatorSet(&_ComputeAVSRegistrar.TransactOpts, operatorSetId)
}

// CreateOperatorSet is a paid mutator transaction binding the contract method 0x74da4713.
//
// Solidity: function createOperatorSet(uint32 operatorSetId) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) CreateOperatorSet(operatorSetId uint32) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.CreateOperatorSet(&_ComputeAVSRegistrar.TransactOpts, operatorSetId)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address , uint32[] operatorSetIds) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "deregisterOperator", operator, arg1, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address , uint32[] operatorSetIds) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) DeregisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.DeregisterOperator(&_ComputeAVSRegistrar.TransactOpts, operator, arg1, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address , uint32[] operatorSetIds) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) DeregisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.DeregisterOperator(&_ComputeAVSRegistrar.TransactOpts, operator, arg1, operatorSetIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string metadataURI) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) Initialize(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "initialize", metadataURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string metadataURI) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) Initialize(metadataURI string) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.Initialize(&_ComputeAVSRegistrar.TransactOpts, metadataURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string metadataURI) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) Initialize(metadataURI string) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.Initialize(&_ComputeAVSRegistrar.TransactOpts, metadataURI)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address , uint32[] operatorSetIds, bytes data) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "registerOperator", operator, arg1, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address , uint32[] operatorSetIds, bytes data) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) RegisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.RegisterOperator(&_ComputeAVSRegistrar.TransactOpts, operator, arg1, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address , uint32[] operatorSetIds, bytes data) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) RegisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.RegisterOperator(&_ComputeAVSRegistrar.TransactOpts, operator, arg1, operatorSetIds, data)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x0a4d3d29.
//
// Solidity: function removeOperatorFromAllowlist((address,uint32) operatorSet, address operator) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) RemoveOperatorFromAllowlist(opts *bind.TransactOpts, operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "removeOperatorFromAllowlist", operatorSet, operator)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x0a4d3d29.
//
// Solidity: function removeOperatorFromAllowlist((address,uint32) operatorSet, address operator) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) RemoveOperatorFromAllowlist(operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.RemoveOperatorFromAllowlist(&_ComputeAVSRegistrar.TransactOpts, operatorSet, operator)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x0a4d3d29.
//
// Solidity: function removeOperatorFromAllowlist((address,uint32) operatorSet, address operator) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) RemoveOperatorFromAllowlist(operatorSet OperatorSet, operator common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.RemoveOperatorFromAllowlist(&_ComputeAVSRegistrar.TransactOpts, operatorSet, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.RenounceOwnership(&_ComputeAVSRegistrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.RenounceOwnership(&_ComputeAVSRegistrar.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.TransferOwnership(&_ComputeAVSRegistrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ComputeAVSRegistrar *ComputeAVSRegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ComputeAVSRegistrar.Contract.TransferOwnership(&_ComputeAVSRegistrar.TransactOpts, newOwner)
}

// ComputeAVSRegistrarInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarInitializedIterator struct {
	Event *ComputeAVSRegistrarInitialized // Event containing the contract specifics and raw log

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
func (it *ComputeAVSRegistrarInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeAVSRegistrarInitialized)
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
		it.Event = new(ComputeAVSRegistrarInitialized)
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
func (it *ComputeAVSRegistrarInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeAVSRegistrarInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeAVSRegistrarInitialized represents a Initialized event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) FilterInitialized(opts *bind.FilterOpts) (*ComputeAVSRegistrarInitializedIterator, error) {

	logs, sub, err := _ComputeAVSRegistrar.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarInitializedIterator{contract: _ComputeAVSRegistrar.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ComputeAVSRegistrarInitialized) (event.Subscription, error) {

	logs, sub, err := _ComputeAVSRegistrar.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeAVSRegistrarInitialized)
				if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) ParseInitialized(log types.Log) (*ComputeAVSRegistrarInitialized, error) {
	event := new(ComputeAVSRegistrarInitialized)
	if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComputeAVSRegistrarOperatorAddedToAllowlistIterator is returned from FilterOperatorAddedToAllowlist and is used to iterate over the raw logs and unpacked data for OperatorAddedToAllowlist events raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorAddedToAllowlistIterator struct {
	Event *ComputeAVSRegistrarOperatorAddedToAllowlist // Event containing the contract specifics and raw log

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
func (it *ComputeAVSRegistrarOperatorAddedToAllowlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeAVSRegistrarOperatorAddedToAllowlist)
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
		it.Event = new(ComputeAVSRegistrarOperatorAddedToAllowlist)
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
func (it *ComputeAVSRegistrarOperatorAddedToAllowlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeAVSRegistrarOperatorAddedToAllowlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeAVSRegistrarOperatorAddedToAllowlist represents a OperatorAddedToAllowlist event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorAddedToAllowlist struct {
	OperatorSet OperatorSet
	Operator    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToAllowlist is a free log retrieval operation binding the contract event 0xfe795219771c42bdbb61ef308cc2b33e1e35b35a3364499b99b2ec2287f20c8c.
//
// Solidity: event OperatorAddedToAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) FilterOperatorAddedToAllowlist(opts *bind.FilterOpts, operatorSet []OperatorSet, operator []common.Address) (*ComputeAVSRegistrarOperatorAddedToAllowlistIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.FilterLogs(opts, "OperatorAddedToAllowlist", operatorSetRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarOperatorAddedToAllowlistIterator{contract: _ComputeAVSRegistrar.contract, event: "OperatorAddedToAllowlist", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToAllowlist is a free log subscription operation binding the contract event 0xfe795219771c42bdbb61ef308cc2b33e1e35b35a3364499b99b2ec2287f20c8c.
//
// Solidity: event OperatorAddedToAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) WatchOperatorAddedToAllowlist(opts *bind.WatchOpts, sink chan<- *ComputeAVSRegistrarOperatorAddedToAllowlist, operatorSet []OperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.WatchLogs(opts, "OperatorAddedToAllowlist", operatorSetRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeAVSRegistrarOperatorAddedToAllowlist)
				if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OperatorAddedToAllowlist", log); err != nil {
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

// ParseOperatorAddedToAllowlist is a log parse operation binding the contract event 0xfe795219771c42bdbb61ef308cc2b33e1e35b35a3364499b99b2ec2287f20c8c.
//
// Solidity: event OperatorAddedToAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) ParseOperatorAddedToAllowlist(log types.Log) (*ComputeAVSRegistrarOperatorAddedToAllowlist, error) {
	event := new(ComputeAVSRegistrarOperatorAddedToAllowlist)
	if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OperatorAddedToAllowlist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComputeAVSRegistrarOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorDeregisteredIterator struct {
	Event *ComputeAVSRegistrarOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ComputeAVSRegistrarOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeAVSRegistrarOperatorDeregistered)
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
		it.Event = new(ComputeAVSRegistrarOperatorDeregistered)
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
func (it *ComputeAVSRegistrarOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeAVSRegistrarOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeAVSRegistrarOperatorDeregistered represents a OperatorDeregistered event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorDeregistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*ComputeAVSRegistrarOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarOperatorDeregisteredIterator{contract: _ComputeAVSRegistrar.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ComputeAVSRegistrarOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeAVSRegistrarOperatorDeregistered)
				if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) ParseOperatorDeregistered(log types.Log) (*ComputeAVSRegistrarOperatorDeregistered, error) {
	event := new(ComputeAVSRegistrarOperatorDeregistered)
	if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComputeAVSRegistrarOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorRegisteredIterator struct {
	Event *ComputeAVSRegistrarOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ComputeAVSRegistrarOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeAVSRegistrarOperatorRegistered)
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
		it.Event = new(ComputeAVSRegistrarOperatorRegistered)
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
func (it *ComputeAVSRegistrarOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeAVSRegistrarOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeAVSRegistrarOperatorRegistered represents a OperatorRegistered event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorRegistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ComputeAVSRegistrarOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarOperatorRegisteredIterator{contract: _ComputeAVSRegistrar.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ComputeAVSRegistrarOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeAVSRegistrarOperatorRegistered)
				if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) ParseOperatorRegistered(log types.Log) (*ComputeAVSRegistrarOperatorRegistered, error) {
	event := new(ComputeAVSRegistrarOperatorRegistered)
	if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComputeAVSRegistrarOperatorRemovedFromAllowlistIterator is returned from FilterOperatorRemovedFromAllowlist and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromAllowlist events raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorRemovedFromAllowlistIterator struct {
	Event *ComputeAVSRegistrarOperatorRemovedFromAllowlist // Event containing the contract specifics and raw log

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
func (it *ComputeAVSRegistrarOperatorRemovedFromAllowlistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeAVSRegistrarOperatorRemovedFromAllowlist)
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
		it.Event = new(ComputeAVSRegistrarOperatorRemovedFromAllowlist)
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
func (it *ComputeAVSRegistrarOperatorRemovedFromAllowlistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeAVSRegistrarOperatorRemovedFromAllowlistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeAVSRegistrarOperatorRemovedFromAllowlist represents a OperatorRemovedFromAllowlist event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorRemovedFromAllowlist struct {
	OperatorSet OperatorSet
	Operator    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromAllowlist is a free log retrieval operation binding the contract event 0x533bf6e1348e64eb9448930dece3436586c031d36722adbc7ccb479809128806.
//
// Solidity: event OperatorRemovedFromAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) FilterOperatorRemovedFromAllowlist(opts *bind.FilterOpts, operatorSet []OperatorSet, operator []common.Address) (*ComputeAVSRegistrarOperatorRemovedFromAllowlistIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.FilterLogs(opts, "OperatorRemovedFromAllowlist", operatorSetRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarOperatorRemovedFromAllowlistIterator{contract: _ComputeAVSRegistrar.contract, event: "OperatorRemovedFromAllowlist", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromAllowlist is a free log subscription operation binding the contract event 0x533bf6e1348e64eb9448930dece3436586c031d36722adbc7ccb479809128806.
//
// Solidity: event OperatorRemovedFromAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) WatchOperatorRemovedFromAllowlist(opts *bind.WatchOpts, sink chan<- *ComputeAVSRegistrarOperatorRemovedFromAllowlist, operatorSet []OperatorSet, operator []common.Address) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.WatchLogs(opts, "OperatorRemovedFromAllowlist", operatorSetRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeAVSRegistrarOperatorRemovedFromAllowlist)
				if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OperatorRemovedFromAllowlist", log); err != nil {
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

// ParseOperatorRemovedFromAllowlist is a log parse operation binding the contract event 0x533bf6e1348e64eb9448930dece3436586c031d36722adbc7ccb479809128806.
//
// Solidity: event OperatorRemovedFromAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) ParseOperatorRemovedFromAllowlist(log types.Log) (*ComputeAVSRegistrarOperatorRemovedFromAllowlist, error) {
	event := new(ComputeAVSRegistrarOperatorRemovedFromAllowlist)
	if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OperatorRemovedFromAllowlist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComputeAVSRegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOwnershipTransferredIterator struct {
	Event *ComputeAVSRegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ComputeAVSRegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputeAVSRegistrarOwnershipTransferred)
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
		it.Event = new(ComputeAVSRegistrarOwnershipTransferred)
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
func (it *ComputeAVSRegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputeAVSRegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputeAVSRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ComputeAVSRegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ComputeAVSRegistrarOwnershipTransferredIterator{contract: _ComputeAVSRegistrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ComputeAVSRegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ComputeAVSRegistrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputeAVSRegistrarOwnershipTransferred)
				if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ComputeAVSRegistrar *ComputeAVSRegistrarFilterer) ParseOwnershipTransferred(log types.Log) (*ComputeAVSRegistrarOwnershipTransferred, error) {
	event := new(ComputeAVSRegistrarOwnershipTransferred)
	if err := _ComputeAVSRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
