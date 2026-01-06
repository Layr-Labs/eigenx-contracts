// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ImageAllowlist

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

// IImageAllowlistPCR is an auto generated low-level Go binding around an user-defined struct.
type IImageAllowlistPCR struct {
	Index uint8
	Value [32]byte
}

// ImageAllowlistMetaData contains all meta data concerning the ImageAllowlist contract.
var ImageAllowlistMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"images\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isImageAllowed\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"pcrs\",\"type\":\"tuple[]\",\"internalType\":\"structIImageAllowlist.PCR[]\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTCBValid\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"tcb\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumTCB\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setImage\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"pcrs\",\"type\":\"tuple[]\",\"internalType\":\"structIImageAllowlist.PCR[]\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"isAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumTCB\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"tcb\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ImageUpdated\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumTCBUpdated\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"tcb\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyPCRs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]}]",
}

// ImageAllowlistABI is the input ABI used to generate the binding from.
// Deprecated: Use ImageAllowlistMetaData.ABI instead.
var ImageAllowlistABI = ImageAllowlistMetaData.ABI

// ImageAllowlist is an auto generated Go binding around an Ethereum contract.
type ImageAllowlist struct {
	ImageAllowlistCaller     // Read-only binding to the contract
	ImageAllowlistTransactor // Write-only binding to the contract
	ImageAllowlistFilterer   // Log filterer for contract events
}

// ImageAllowlistCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImageAllowlistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImageAllowlistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImageAllowlistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImageAllowlistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImageAllowlistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImageAllowlistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImageAllowlistSession struct {
	Contract     *ImageAllowlist   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImageAllowlistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImageAllowlistCallerSession struct {
	Contract *ImageAllowlistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ImageAllowlistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImageAllowlistTransactorSession struct {
	Contract     *ImageAllowlistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ImageAllowlistRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImageAllowlistRaw struct {
	Contract *ImageAllowlist // Generic contract binding to access the raw methods on
}

// ImageAllowlistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImageAllowlistCallerRaw struct {
	Contract *ImageAllowlistCaller // Generic read-only contract binding to access the raw methods on
}

// ImageAllowlistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImageAllowlistTransactorRaw struct {
	Contract *ImageAllowlistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImageAllowlist creates a new instance of ImageAllowlist, bound to a specific deployed contract.
func NewImageAllowlist(address common.Address, backend bind.ContractBackend) (*ImageAllowlist, error) {
	contract, err := bindImageAllowlist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ImageAllowlist{ImageAllowlistCaller: ImageAllowlistCaller{contract: contract}, ImageAllowlistTransactor: ImageAllowlistTransactor{contract: contract}, ImageAllowlistFilterer: ImageAllowlistFilterer{contract: contract}}, nil
}

// NewImageAllowlistCaller creates a new read-only instance of ImageAllowlist, bound to a specific deployed contract.
func NewImageAllowlistCaller(address common.Address, caller bind.ContractCaller) (*ImageAllowlistCaller, error) {
	contract, err := bindImageAllowlist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImageAllowlistCaller{contract: contract}, nil
}

// NewImageAllowlistTransactor creates a new write-only instance of ImageAllowlist, bound to a specific deployed contract.
func NewImageAllowlistTransactor(address common.Address, transactor bind.ContractTransactor) (*ImageAllowlistTransactor, error) {
	contract, err := bindImageAllowlist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImageAllowlistTransactor{contract: contract}, nil
}

// NewImageAllowlistFilterer creates a new log filterer instance of ImageAllowlist, bound to a specific deployed contract.
func NewImageAllowlistFilterer(address common.Address, filterer bind.ContractFilterer) (*ImageAllowlistFilterer, error) {
	contract, err := bindImageAllowlist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImageAllowlistFilterer{contract: contract}, nil
}

// bindImageAllowlist binds a generic wrapper to an already deployed contract.
func bindImageAllowlist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ImageAllowlistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImageAllowlist *ImageAllowlistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImageAllowlist.Contract.ImageAllowlistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImageAllowlist *ImageAllowlistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.ImageAllowlistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImageAllowlist *ImageAllowlistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.ImageAllowlistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImageAllowlist *ImageAllowlistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImageAllowlist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImageAllowlist *ImageAllowlistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImageAllowlist *ImageAllowlistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.contract.Transact(opts, method, params...)
}

// Images is a free data retrieval call binding the contract method 0x4a4f02de.
//
// Solidity: function images(uint8 , bytes32 ) view returns(bool)
func (_ImageAllowlist *ImageAllowlistCaller) Images(opts *bind.CallOpts, arg0 uint8, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _ImageAllowlist.contract.Call(opts, &out, "images", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Images is a free data retrieval call binding the contract method 0x4a4f02de.
//
// Solidity: function images(uint8 , bytes32 ) view returns(bool)
func (_ImageAllowlist *ImageAllowlistSession) Images(arg0 uint8, arg1 [32]byte) (bool, error) {
	return _ImageAllowlist.Contract.Images(&_ImageAllowlist.CallOpts, arg0, arg1)
}

// Images is a free data retrieval call binding the contract method 0x4a4f02de.
//
// Solidity: function images(uint8 , bytes32 ) view returns(bool)
func (_ImageAllowlist *ImageAllowlistCallerSession) Images(arg0 uint8, arg1 [32]byte) (bool, error) {
	return _ImageAllowlist.Contract.Images(&_ImageAllowlist.CallOpts, arg0, arg1)
}

// IsImageAllowed is a free data retrieval call binding the contract method 0x2d5a1daf.
//
// Solidity: function isImageAllowed(uint8 cvm, (uint8,bytes32)[] pcrs) view returns(bool)
func (_ImageAllowlist *ImageAllowlistCaller) IsImageAllowed(opts *bind.CallOpts, cvm uint8, pcrs []IImageAllowlistPCR) (bool, error) {
	var out []interface{}
	err := _ImageAllowlist.contract.Call(opts, &out, "isImageAllowed", cvm, pcrs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsImageAllowed is a free data retrieval call binding the contract method 0x2d5a1daf.
//
// Solidity: function isImageAllowed(uint8 cvm, (uint8,bytes32)[] pcrs) view returns(bool)
func (_ImageAllowlist *ImageAllowlistSession) IsImageAllowed(cvm uint8, pcrs []IImageAllowlistPCR) (bool, error) {
	return _ImageAllowlist.Contract.IsImageAllowed(&_ImageAllowlist.CallOpts, cvm, pcrs)
}

// IsImageAllowed is a free data retrieval call binding the contract method 0x2d5a1daf.
//
// Solidity: function isImageAllowed(uint8 cvm, (uint8,bytes32)[] pcrs) view returns(bool)
func (_ImageAllowlist *ImageAllowlistCallerSession) IsImageAllowed(cvm uint8, pcrs []IImageAllowlistPCR) (bool, error) {
	return _ImageAllowlist.Contract.IsImageAllowed(&_ImageAllowlist.CallOpts, cvm, pcrs)
}

// IsTCBValid is a free data retrieval call binding the contract method 0x43e3dd05.
//
// Solidity: function isTCBValid(uint8 cvm, uint64 tcb) view returns(bool)
func (_ImageAllowlist *ImageAllowlistCaller) IsTCBValid(opts *bind.CallOpts, cvm uint8, tcb uint64) (bool, error) {
	var out []interface{}
	err := _ImageAllowlist.contract.Call(opts, &out, "isTCBValid", cvm, tcb)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTCBValid is a free data retrieval call binding the contract method 0x43e3dd05.
//
// Solidity: function isTCBValid(uint8 cvm, uint64 tcb) view returns(bool)
func (_ImageAllowlist *ImageAllowlistSession) IsTCBValid(cvm uint8, tcb uint64) (bool, error) {
	return _ImageAllowlist.Contract.IsTCBValid(&_ImageAllowlist.CallOpts, cvm, tcb)
}

// IsTCBValid is a free data retrieval call binding the contract method 0x43e3dd05.
//
// Solidity: function isTCBValid(uint8 cvm, uint64 tcb) view returns(bool)
func (_ImageAllowlist *ImageAllowlistCallerSession) IsTCBValid(cvm uint8, tcb uint64) (bool, error) {
	return _ImageAllowlist.Contract.IsTCBValid(&_ImageAllowlist.CallOpts, cvm, tcb)
}

// MinimumTCB is a free data retrieval call binding the contract method 0x4ec8d4fa.
//
// Solidity: function minimumTCB(uint8 ) view returns(uint64)
func (_ImageAllowlist *ImageAllowlistCaller) MinimumTCB(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _ImageAllowlist.contract.Call(opts, &out, "minimumTCB", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumTCB is a free data retrieval call binding the contract method 0x4ec8d4fa.
//
// Solidity: function minimumTCB(uint8 ) view returns(uint64)
func (_ImageAllowlist *ImageAllowlistSession) MinimumTCB(arg0 uint8) (uint64, error) {
	return _ImageAllowlist.Contract.MinimumTCB(&_ImageAllowlist.CallOpts, arg0)
}

// MinimumTCB is a free data retrieval call binding the contract method 0x4ec8d4fa.
//
// Solidity: function minimumTCB(uint8 ) view returns(uint64)
func (_ImageAllowlist *ImageAllowlistCallerSession) MinimumTCB(arg0 uint8) (uint64, error) {
	return _ImageAllowlist.Contract.MinimumTCB(&_ImageAllowlist.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ImageAllowlist *ImageAllowlistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ImageAllowlist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ImageAllowlist *ImageAllowlistSession) Owner() (common.Address, error) {
	return _ImageAllowlist.Contract.Owner(&_ImageAllowlist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ImageAllowlist *ImageAllowlistCallerSession) Owner() (common.Address, error) {
	return _ImageAllowlist.Contract.Owner(&_ImageAllowlist.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ImageAllowlist *ImageAllowlistTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ImageAllowlist.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ImageAllowlist *ImageAllowlistSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.Initialize(&_ImageAllowlist.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ImageAllowlist *ImageAllowlistTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.Initialize(&_ImageAllowlist.TransactOpts, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ImageAllowlist *ImageAllowlistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImageAllowlist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ImageAllowlist *ImageAllowlistSession) RenounceOwnership() (*types.Transaction, error) {
	return _ImageAllowlist.Contract.RenounceOwnership(&_ImageAllowlist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ImageAllowlist *ImageAllowlistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ImageAllowlist.Contract.RenounceOwnership(&_ImageAllowlist.TransactOpts)
}

// SetImage is a paid mutator transaction binding the contract method 0xb85bbdca.
//
// Solidity: function setImage(uint8 cvm, (uint8,bytes32)[] pcrs, bool isAllowed) returns()
func (_ImageAllowlist *ImageAllowlistTransactor) SetImage(opts *bind.TransactOpts, cvm uint8, pcrs []IImageAllowlistPCR, isAllowed bool) (*types.Transaction, error) {
	return _ImageAllowlist.contract.Transact(opts, "setImage", cvm, pcrs, isAllowed)
}

// SetImage is a paid mutator transaction binding the contract method 0xb85bbdca.
//
// Solidity: function setImage(uint8 cvm, (uint8,bytes32)[] pcrs, bool isAllowed) returns()
func (_ImageAllowlist *ImageAllowlistSession) SetImage(cvm uint8, pcrs []IImageAllowlistPCR, isAllowed bool) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.SetImage(&_ImageAllowlist.TransactOpts, cvm, pcrs, isAllowed)
}

// SetImage is a paid mutator transaction binding the contract method 0xb85bbdca.
//
// Solidity: function setImage(uint8 cvm, (uint8,bytes32)[] pcrs, bool isAllowed) returns()
func (_ImageAllowlist *ImageAllowlistTransactorSession) SetImage(cvm uint8, pcrs []IImageAllowlistPCR, isAllowed bool) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.SetImage(&_ImageAllowlist.TransactOpts, cvm, pcrs, isAllowed)
}

// SetMinimumTCB is a paid mutator transaction binding the contract method 0xa54c8287.
//
// Solidity: function setMinimumTCB(uint8 cvm, uint64 tcb) returns()
func (_ImageAllowlist *ImageAllowlistTransactor) SetMinimumTCB(opts *bind.TransactOpts, cvm uint8, tcb uint64) (*types.Transaction, error) {
	return _ImageAllowlist.contract.Transact(opts, "setMinimumTCB", cvm, tcb)
}

// SetMinimumTCB is a paid mutator transaction binding the contract method 0xa54c8287.
//
// Solidity: function setMinimumTCB(uint8 cvm, uint64 tcb) returns()
func (_ImageAllowlist *ImageAllowlistSession) SetMinimumTCB(cvm uint8, tcb uint64) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.SetMinimumTCB(&_ImageAllowlist.TransactOpts, cvm, tcb)
}

// SetMinimumTCB is a paid mutator transaction binding the contract method 0xa54c8287.
//
// Solidity: function setMinimumTCB(uint8 cvm, uint64 tcb) returns()
func (_ImageAllowlist *ImageAllowlistTransactorSession) SetMinimumTCB(cvm uint8, tcb uint64) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.SetMinimumTCB(&_ImageAllowlist.TransactOpts, cvm, tcb)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ImageAllowlist *ImageAllowlistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ImageAllowlist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ImageAllowlist *ImageAllowlistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.TransferOwnership(&_ImageAllowlist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ImageAllowlist *ImageAllowlistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ImageAllowlist.Contract.TransferOwnership(&_ImageAllowlist.TransactOpts, newOwner)
}

// ImageAllowlistImageUpdatedIterator is returned from FilterImageUpdated and is used to iterate over the raw logs and unpacked data for ImageUpdated events raised by the ImageAllowlist contract.
type ImageAllowlistImageUpdatedIterator struct {
	Event *ImageAllowlistImageUpdated // Event containing the contract specifics and raw log

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
func (it *ImageAllowlistImageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImageAllowlistImageUpdated)
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
		it.Event = new(ImageAllowlistImageUpdated)
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
func (it *ImageAllowlistImageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImageAllowlistImageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImageAllowlistImageUpdated represents a ImageUpdated event raised by the ImageAllowlist contract.
type ImageAllowlistImageUpdated struct {
	Cvm     uint8
	Key     [32]byte
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterImageUpdated is a free log retrieval operation binding the contract event 0xf13729199ed190a7f2ddb596e495afbc3e8fcebf4035edc8e97a3d47bf9b91d5.
//
// Solidity: event ImageUpdated(uint8 indexed cvm, bytes32 indexed key, bool allowed)
func (_ImageAllowlist *ImageAllowlistFilterer) FilterImageUpdated(opts *bind.FilterOpts, cvm []uint8, key [][32]byte) (*ImageAllowlistImageUpdatedIterator, error) {

	var cvmRule []interface{}
	for _, cvmItem := range cvm {
		cvmRule = append(cvmRule, cvmItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ImageAllowlist.contract.FilterLogs(opts, "ImageUpdated", cvmRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &ImageAllowlistImageUpdatedIterator{contract: _ImageAllowlist.contract, event: "ImageUpdated", logs: logs, sub: sub}, nil
}

// WatchImageUpdated is a free log subscription operation binding the contract event 0xf13729199ed190a7f2ddb596e495afbc3e8fcebf4035edc8e97a3d47bf9b91d5.
//
// Solidity: event ImageUpdated(uint8 indexed cvm, bytes32 indexed key, bool allowed)
func (_ImageAllowlist *ImageAllowlistFilterer) WatchImageUpdated(opts *bind.WatchOpts, sink chan<- *ImageAllowlistImageUpdated, cvm []uint8, key [][32]byte) (event.Subscription, error) {

	var cvmRule []interface{}
	for _, cvmItem := range cvm {
		cvmRule = append(cvmRule, cvmItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ImageAllowlist.contract.WatchLogs(opts, "ImageUpdated", cvmRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImageAllowlistImageUpdated)
				if err := _ImageAllowlist.contract.UnpackLog(event, "ImageUpdated", log); err != nil {
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

// ParseImageUpdated is a log parse operation binding the contract event 0xf13729199ed190a7f2ddb596e495afbc3e8fcebf4035edc8e97a3d47bf9b91d5.
//
// Solidity: event ImageUpdated(uint8 indexed cvm, bytes32 indexed key, bool allowed)
func (_ImageAllowlist *ImageAllowlistFilterer) ParseImageUpdated(log types.Log) (*ImageAllowlistImageUpdated, error) {
	event := new(ImageAllowlistImageUpdated)
	if err := _ImageAllowlist.contract.UnpackLog(event, "ImageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImageAllowlistInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ImageAllowlist contract.
type ImageAllowlistInitializedIterator struct {
	Event *ImageAllowlistInitialized // Event containing the contract specifics and raw log

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
func (it *ImageAllowlistInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImageAllowlistInitialized)
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
		it.Event = new(ImageAllowlistInitialized)
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
func (it *ImageAllowlistInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImageAllowlistInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImageAllowlistInitialized represents a Initialized event raised by the ImageAllowlist contract.
type ImageAllowlistInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ImageAllowlist *ImageAllowlistFilterer) FilterInitialized(opts *bind.FilterOpts) (*ImageAllowlistInitializedIterator, error) {

	logs, sub, err := _ImageAllowlist.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ImageAllowlistInitializedIterator{contract: _ImageAllowlist.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ImageAllowlist *ImageAllowlistFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ImageAllowlistInitialized) (event.Subscription, error) {

	logs, sub, err := _ImageAllowlist.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImageAllowlistInitialized)
				if err := _ImageAllowlist.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ImageAllowlist *ImageAllowlistFilterer) ParseInitialized(log types.Log) (*ImageAllowlistInitialized, error) {
	event := new(ImageAllowlistInitialized)
	if err := _ImageAllowlist.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImageAllowlistMinimumTCBUpdatedIterator is returned from FilterMinimumTCBUpdated and is used to iterate over the raw logs and unpacked data for MinimumTCBUpdated events raised by the ImageAllowlist contract.
type ImageAllowlistMinimumTCBUpdatedIterator struct {
	Event *ImageAllowlistMinimumTCBUpdated // Event containing the contract specifics and raw log

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
func (it *ImageAllowlistMinimumTCBUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImageAllowlistMinimumTCBUpdated)
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
		it.Event = new(ImageAllowlistMinimumTCBUpdated)
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
func (it *ImageAllowlistMinimumTCBUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImageAllowlistMinimumTCBUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImageAllowlistMinimumTCBUpdated represents a MinimumTCBUpdated event raised by the ImageAllowlist contract.
type ImageAllowlistMinimumTCBUpdated struct {
	Cvm uint8
	Tcb uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinimumTCBUpdated is a free log retrieval operation binding the contract event 0xaf8be91f34cee8255c2e461ce2588f8ffa3e5c01d2f0ec91b0f4a010bce9a392.
//
// Solidity: event MinimumTCBUpdated(uint8 indexed cvm, uint64 tcb)
func (_ImageAllowlist *ImageAllowlistFilterer) FilterMinimumTCBUpdated(opts *bind.FilterOpts, cvm []uint8) (*ImageAllowlistMinimumTCBUpdatedIterator, error) {

	var cvmRule []interface{}
	for _, cvmItem := range cvm {
		cvmRule = append(cvmRule, cvmItem)
	}

	logs, sub, err := _ImageAllowlist.contract.FilterLogs(opts, "MinimumTCBUpdated", cvmRule)
	if err != nil {
		return nil, err
	}
	return &ImageAllowlistMinimumTCBUpdatedIterator{contract: _ImageAllowlist.contract, event: "MinimumTCBUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumTCBUpdated is a free log subscription operation binding the contract event 0xaf8be91f34cee8255c2e461ce2588f8ffa3e5c01d2f0ec91b0f4a010bce9a392.
//
// Solidity: event MinimumTCBUpdated(uint8 indexed cvm, uint64 tcb)
func (_ImageAllowlist *ImageAllowlistFilterer) WatchMinimumTCBUpdated(opts *bind.WatchOpts, sink chan<- *ImageAllowlistMinimumTCBUpdated, cvm []uint8) (event.Subscription, error) {

	var cvmRule []interface{}
	for _, cvmItem := range cvm {
		cvmRule = append(cvmRule, cvmItem)
	}

	logs, sub, err := _ImageAllowlist.contract.WatchLogs(opts, "MinimumTCBUpdated", cvmRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImageAllowlistMinimumTCBUpdated)
				if err := _ImageAllowlist.contract.UnpackLog(event, "MinimumTCBUpdated", log); err != nil {
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

// ParseMinimumTCBUpdated is a log parse operation binding the contract event 0xaf8be91f34cee8255c2e461ce2588f8ffa3e5c01d2f0ec91b0f4a010bce9a392.
//
// Solidity: event MinimumTCBUpdated(uint8 indexed cvm, uint64 tcb)
func (_ImageAllowlist *ImageAllowlistFilterer) ParseMinimumTCBUpdated(log types.Log) (*ImageAllowlistMinimumTCBUpdated, error) {
	event := new(ImageAllowlistMinimumTCBUpdated)
	if err := _ImageAllowlist.contract.UnpackLog(event, "MinimumTCBUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImageAllowlistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ImageAllowlist contract.
type ImageAllowlistOwnershipTransferredIterator struct {
	Event *ImageAllowlistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ImageAllowlistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImageAllowlistOwnershipTransferred)
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
		it.Event = new(ImageAllowlistOwnershipTransferred)
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
func (it *ImageAllowlistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImageAllowlistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImageAllowlistOwnershipTransferred represents a OwnershipTransferred event raised by the ImageAllowlist contract.
type ImageAllowlistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ImageAllowlist *ImageAllowlistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ImageAllowlistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ImageAllowlist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ImageAllowlistOwnershipTransferredIterator{contract: _ImageAllowlist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ImageAllowlist *ImageAllowlistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ImageAllowlistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ImageAllowlist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImageAllowlistOwnershipTransferred)
				if err := _ImageAllowlist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ImageAllowlist *ImageAllowlistFilterer) ParseOwnershipTransferred(log types.Log) (*ImageAllowlistOwnershipTransferred, error) {
	event := new(ImageAllowlistOwnershipTransferred)
	if err := _ImageAllowlist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
