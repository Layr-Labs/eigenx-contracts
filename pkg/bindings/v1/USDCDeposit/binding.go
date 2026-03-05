// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package USDCDeposit

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

// USDCDepositMetaData contains all meta data concerning the USDCDeposit contract.
var USDCDepositMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_usdc\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_treasury\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositFor\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minimumDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minimumDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumDeposit\",\"inputs\":[{\"name\":\"newMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumDepositSet\",\"inputs\":[{\"name\":\"oldMinimum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinimum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BelowMinimumDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// USDCDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use USDCDepositMetaData.ABI instead.
var USDCDepositABI = USDCDepositMetaData.ABI

// USDCDeposit is an auto generated Go binding around an Ethereum contract.
type USDCDeposit struct {
	USDCDepositCaller     // Read-only binding to the contract
	USDCDepositTransactor // Write-only binding to the contract
	USDCDepositFilterer   // Log filterer for contract events
}

// USDCDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDCDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDCDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDCDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDCDepositSession struct {
	Contract     *USDCDeposit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDCDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDCDepositCallerSession struct {
	Contract *USDCDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// USDCDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDCDepositTransactorSession struct {
	Contract     *USDCDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// USDCDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDCDepositRaw struct {
	Contract *USDCDeposit // Generic contract binding to access the raw methods on
}

// USDCDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDCDepositCallerRaw struct {
	Contract *USDCDepositCaller // Generic read-only contract binding to access the raw methods on
}

// USDCDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDCDepositTransactorRaw struct {
	Contract *USDCDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDCDeposit creates a new instance of USDCDeposit, bound to a specific deployed contract.
func NewUSDCDeposit(address common.Address, backend bind.ContractBackend) (*USDCDeposit, error) {
	contract, err := bindUSDCDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCDeposit{USDCDepositCaller: USDCDepositCaller{contract: contract}, USDCDepositTransactor: USDCDepositTransactor{contract: contract}, USDCDepositFilterer: USDCDepositFilterer{contract: contract}}, nil
}

// NewUSDCDepositCaller creates a new read-only instance of USDCDeposit, bound to a specific deployed contract.
func NewUSDCDepositCaller(address common.Address, caller bind.ContractCaller) (*USDCDepositCaller, error) {
	contract, err := bindUSDCDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCDepositCaller{contract: contract}, nil
}

// NewUSDCDepositTransactor creates a new write-only instance of USDCDeposit, bound to a specific deployed contract.
func NewUSDCDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*USDCDepositTransactor, error) {
	contract, err := bindUSDCDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCDepositTransactor{contract: contract}, nil
}

// NewUSDCDepositFilterer creates a new log filterer instance of USDCDeposit, bound to a specific deployed contract.
func NewUSDCDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCDepositFilterer, error) {
	contract, err := bindUSDCDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCDepositFilterer{contract: contract}, nil
}

// bindUSDCDeposit binds a generic wrapper to an already deployed contract.
func bindUSDCDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDCDepositMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDCDeposit *USDCDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCDeposit.Contract.USDCDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDCDeposit *USDCDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCDeposit.Contract.USDCDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDCDeposit *USDCDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCDeposit.Contract.USDCDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDCDeposit *USDCDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDCDeposit *USDCDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDCDeposit *USDCDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCDeposit.Contract.contract.Transact(opts, method, params...)
}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (_USDCDeposit *USDCDepositCaller) MinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCDeposit.contract.Call(opts, &out, "minimumDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (_USDCDeposit *USDCDepositSession) MinimumDeposit() (*big.Int, error) {
	return _USDCDeposit.Contract.MinimumDeposit(&_USDCDeposit.CallOpts)
}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (_USDCDeposit *USDCDepositCallerSession) MinimumDeposit() (*big.Int, error) {
	return _USDCDeposit.Contract.MinimumDeposit(&_USDCDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCDeposit *USDCDepositCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCDeposit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCDeposit *USDCDepositSession) Owner() (common.Address, error) {
	return _USDCDeposit.Contract.Owner(&_USDCDeposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCDeposit *USDCDepositCallerSession) Owner() (common.Address, error) {
	return _USDCDeposit.Contract.Owner(&_USDCDeposit.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_USDCDeposit *USDCDepositCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCDeposit.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_USDCDeposit *USDCDepositSession) Treasury() (common.Address, error) {
	return _USDCDeposit.Contract.Treasury(&_USDCDeposit.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_USDCDeposit *USDCDepositCallerSession) Treasury() (common.Address, error) {
	return _USDCDeposit.Contract.Treasury(&_USDCDeposit.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCDeposit *USDCDepositCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCDeposit.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCDeposit *USDCDepositSession) Usdc() (common.Address, error) {
	return _USDCDeposit.Contract.Usdc(&_USDCDeposit.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCDeposit *USDCDepositCallerSession) Usdc() (common.Address, error) {
	return _USDCDeposit.Contract.Usdc(&_USDCDeposit.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_USDCDeposit *USDCDepositTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_USDCDeposit *USDCDepositSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.Contract.Deposit(&_USDCDeposit.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_USDCDeposit *USDCDepositTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.Contract.Deposit(&_USDCDeposit.TransactOpts, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address account) returns()
func (_USDCDeposit *USDCDepositTransactor) DepositFor(opts *bind.TransactOpts, amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _USDCDeposit.contract.Transact(opts, "depositFor", amount, account)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address account) returns()
func (_USDCDeposit *USDCDepositSession) DepositFor(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _USDCDeposit.Contract.DepositFor(&_USDCDeposit.TransactOpts, amount, account)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address account) returns()
func (_USDCDeposit *USDCDepositTransactorSession) DepositFor(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _USDCDeposit.Contract.DepositFor(&_USDCDeposit.TransactOpts, amount, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (_USDCDeposit *USDCDepositTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _minimumDeposit *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.contract.Transact(opts, "initialize", initialOwner, _minimumDeposit)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (_USDCDeposit *USDCDepositSession) Initialize(initialOwner common.Address, _minimumDeposit *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.Contract.Initialize(&_USDCDeposit.TransactOpts, initialOwner, _minimumDeposit)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (_USDCDeposit *USDCDepositTransactorSession) Initialize(initialOwner common.Address, _minimumDeposit *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.Contract.Initialize(&_USDCDeposit.TransactOpts, initialOwner, _minimumDeposit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDCDeposit *USDCDepositTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCDeposit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDCDeposit *USDCDepositSession) RenounceOwnership() (*types.Transaction, error) {
	return _USDCDeposit.Contract.RenounceOwnership(&_USDCDeposit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDCDeposit *USDCDepositTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _USDCDeposit.Contract.RenounceOwnership(&_USDCDeposit.TransactOpts)
}

// SetMinimumDeposit is a paid mutator transaction binding the contract method 0xe78ec42e.
//
// Solidity: function setMinimumDeposit(uint256 newMinimum) returns()
func (_USDCDeposit *USDCDepositTransactor) SetMinimumDeposit(opts *bind.TransactOpts, newMinimum *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.contract.Transact(opts, "setMinimumDeposit", newMinimum)
}

// SetMinimumDeposit is a paid mutator transaction binding the contract method 0xe78ec42e.
//
// Solidity: function setMinimumDeposit(uint256 newMinimum) returns()
func (_USDCDeposit *USDCDepositSession) SetMinimumDeposit(newMinimum *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.Contract.SetMinimumDeposit(&_USDCDeposit.TransactOpts, newMinimum)
}

// SetMinimumDeposit is a paid mutator transaction binding the contract method 0xe78ec42e.
//
// Solidity: function setMinimumDeposit(uint256 newMinimum) returns()
func (_USDCDeposit *USDCDepositTransactorSession) SetMinimumDeposit(newMinimum *big.Int) (*types.Transaction, error) {
	return _USDCDeposit.Contract.SetMinimumDeposit(&_USDCDeposit.TransactOpts, newMinimum)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_USDCDeposit *USDCDepositTransactor) Sweep(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _USDCDeposit.contract.Transact(opts, "sweep", token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_USDCDeposit *USDCDepositSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _USDCDeposit.Contract.Sweep(&_USDCDeposit.TransactOpts, token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_USDCDeposit *USDCDepositTransactorSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _USDCDeposit.Contract.Sweep(&_USDCDeposit.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDCDeposit *USDCDepositTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _USDCDeposit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDCDeposit *USDCDepositSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDCDeposit.Contract.TransferOwnership(&_USDCDeposit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDCDeposit *USDCDepositTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDCDeposit.Contract.TransferOwnership(&_USDCDeposit.TransactOpts, newOwner)
}

// USDCDepositDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the USDCDeposit contract.
type USDCDepositDepositIterator struct {
	Event *USDCDepositDeposit // Event containing the contract specifics and raw log

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
func (it *USDCDepositDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCDepositDeposit)
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
		it.Event = new(USDCDepositDeposit)
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
func (it *USDCDepositDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCDepositDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCDepositDeposit represents a Deposit event raised by the USDCDeposit contract.
type USDCDepositDeposit struct {
	Depositor common.Address
	Account   common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed depositor, address indexed account, uint256 amount)
func (_USDCDeposit *USDCDepositFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address, account []common.Address) (*USDCDepositDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _USDCDeposit.contract.FilterLogs(opts, "Deposit", depositorRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &USDCDepositDepositIterator{contract: _USDCDeposit.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed depositor, address indexed account, uint256 amount)
func (_USDCDeposit *USDCDepositFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *USDCDepositDeposit, depositor []common.Address, account []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _USDCDeposit.contract.WatchLogs(opts, "Deposit", depositorRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCDepositDeposit)
				if err := _USDCDeposit.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed depositor, address indexed account, uint256 amount)
func (_USDCDeposit *USDCDepositFilterer) ParseDeposit(log types.Log) (*USDCDepositDeposit, error) {
	event := new(USDCDepositDeposit)
	if err := _USDCDeposit.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDCDepositInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the USDCDeposit contract.
type USDCDepositInitializedIterator struct {
	Event *USDCDepositInitialized // Event containing the contract specifics and raw log

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
func (it *USDCDepositInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCDepositInitialized)
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
		it.Event = new(USDCDepositInitialized)
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
func (it *USDCDepositInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCDepositInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCDepositInitialized represents a Initialized event raised by the USDCDeposit contract.
type USDCDepositInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_USDCDeposit *USDCDepositFilterer) FilterInitialized(opts *bind.FilterOpts) (*USDCDepositInitializedIterator, error) {

	logs, sub, err := _USDCDeposit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &USDCDepositInitializedIterator{contract: _USDCDeposit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_USDCDeposit *USDCDepositFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *USDCDepositInitialized) (event.Subscription, error) {

	logs, sub, err := _USDCDeposit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCDepositInitialized)
				if err := _USDCDeposit.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_USDCDeposit *USDCDepositFilterer) ParseInitialized(log types.Log) (*USDCDepositInitialized, error) {
	event := new(USDCDepositInitialized)
	if err := _USDCDeposit.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDCDepositMinimumDepositSetIterator is returned from FilterMinimumDepositSet and is used to iterate over the raw logs and unpacked data for MinimumDepositSet events raised by the USDCDeposit contract.
type USDCDepositMinimumDepositSetIterator struct {
	Event *USDCDepositMinimumDepositSet // Event containing the contract specifics and raw log

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
func (it *USDCDepositMinimumDepositSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCDepositMinimumDepositSet)
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
		it.Event = new(USDCDepositMinimumDepositSet)
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
func (it *USDCDepositMinimumDepositSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCDepositMinimumDepositSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCDepositMinimumDepositSet represents a MinimumDepositSet event raised by the USDCDeposit contract.
type USDCDepositMinimumDepositSet struct {
	OldMinimum *big.Int
	NewMinimum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinimumDepositSet is a free log retrieval operation binding the contract event 0x5da812c62f7dbf913260a291f7a1d45cb11ad9527d416f5affc1c3d348ea277e.
//
// Solidity: event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum)
func (_USDCDeposit *USDCDepositFilterer) FilterMinimumDepositSet(opts *bind.FilterOpts) (*USDCDepositMinimumDepositSetIterator, error) {

	logs, sub, err := _USDCDeposit.contract.FilterLogs(opts, "MinimumDepositSet")
	if err != nil {
		return nil, err
	}
	return &USDCDepositMinimumDepositSetIterator{contract: _USDCDeposit.contract, event: "MinimumDepositSet", logs: logs, sub: sub}, nil
}

// WatchMinimumDepositSet is a free log subscription operation binding the contract event 0x5da812c62f7dbf913260a291f7a1d45cb11ad9527d416f5affc1c3d348ea277e.
//
// Solidity: event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum)
func (_USDCDeposit *USDCDepositFilterer) WatchMinimumDepositSet(opts *bind.WatchOpts, sink chan<- *USDCDepositMinimumDepositSet) (event.Subscription, error) {

	logs, sub, err := _USDCDeposit.contract.WatchLogs(opts, "MinimumDepositSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCDepositMinimumDepositSet)
				if err := _USDCDeposit.contract.UnpackLog(event, "MinimumDepositSet", log); err != nil {
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

// ParseMinimumDepositSet is a log parse operation binding the contract event 0x5da812c62f7dbf913260a291f7a1d45cb11ad9527d416f5affc1c3d348ea277e.
//
// Solidity: event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum)
func (_USDCDeposit *USDCDepositFilterer) ParseMinimumDepositSet(log types.Log) (*USDCDepositMinimumDepositSet, error) {
	event := new(USDCDepositMinimumDepositSet)
	if err := _USDCDeposit.contract.UnpackLog(event, "MinimumDepositSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDCDepositOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the USDCDeposit contract.
type USDCDepositOwnershipTransferredIterator struct {
	Event *USDCDepositOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *USDCDepositOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCDepositOwnershipTransferred)
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
		it.Event = new(USDCDepositOwnershipTransferred)
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
func (it *USDCDepositOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCDepositOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCDepositOwnershipTransferred represents a OwnershipTransferred event raised by the USDCDeposit contract.
type USDCDepositOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDCDeposit *USDCDepositFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*USDCDepositOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USDCDeposit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &USDCDepositOwnershipTransferredIterator{contract: _USDCDeposit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDCDeposit *USDCDepositFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCDepositOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USDCDeposit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCDepositOwnershipTransferred)
				if err := _USDCDeposit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_USDCDeposit *USDCDepositFilterer) ParseOwnershipTransferred(log types.Log) (*USDCDepositOwnershipTransferred, error) {
	event := new(USDCDepositOwnershipTransferred)
	if err := _USDCDeposit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
