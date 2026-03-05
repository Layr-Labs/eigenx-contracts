// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package USDCCredits

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

// USDCCreditsMetaData contains all meta data concerning the USDCCredits contract.
var USDCCreditsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_usdc\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_treasury\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minimumDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minimumDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"purchaseCredits\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"purchaseCreditsFor\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumDepositFor\",\"inputs\":[{\"name\":\"newMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumDepositSet\",\"inputs\":[{\"name\":\"oldMinimum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinimum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BelowMinimumDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// USDCCreditsABI is the input ABI used to generate the binding from.
// Deprecated: Use USDCCreditsMetaData.ABI instead.
var USDCCreditsABI = USDCCreditsMetaData.ABI

// USDCCredits is an auto generated Go binding around an Ethereum contract.
type USDCCredits struct {
	USDCCreditsCaller     // Read-only binding to the contract
	USDCCreditsTransactor // Write-only binding to the contract
	USDCCreditsFilterer   // Log filterer for contract events
}

// USDCCreditsCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDCCreditsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCCreditsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDCCreditsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCCreditsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDCCreditsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCCreditsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDCCreditsSession struct {
	Contract     *USDCCredits      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDCCreditsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDCCreditsCallerSession struct {
	Contract *USDCCreditsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// USDCCreditsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDCCreditsTransactorSession struct {
	Contract     *USDCCreditsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// USDCCreditsRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDCCreditsRaw struct {
	Contract *USDCCredits // Generic contract binding to access the raw methods on
}

// USDCCreditsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDCCreditsCallerRaw struct {
	Contract *USDCCreditsCaller // Generic read-only contract binding to access the raw methods on
}

// USDCCreditsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDCCreditsTransactorRaw struct {
	Contract *USDCCreditsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDCCredits creates a new instance of USDCCredits, bound to a specific deployed contract.
func NewUSDCCredits(address common.Address, backend bind.ContractBackend) (*USDCCredits, error) {
	contract, err := bindUSDCCredits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCCredits{USDCCreditsCaller: USDCCreditsCaller{contract: contract}, USDCCreditsTransactor: USDCCreditsTransactor{contract: contract}, USDCCreditsFilterer: USDCCreditsFilterer{contract: contract}}, nil
}

// NewUSDCCreditsCaller creates a new read-only instance of USDCCredits, bound to a specific deployed contract.
func NewUSDCCreditsCaller(address common.Address, caller bind.ContractCaller) (*USDCCreditsCaller, error) {
	contract, err := bindUSDCCredits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCCreditsCaller{contract: contract}, nil
}

// NewUSDCCreditsTransactor creates a new write-only instance of USDCCredits, bound to a specific deployed contract.
func NewUSDCCreditsTransactor(address common.Address, transactor bind.ContractTransactor) (*USDCCreditsTransactor, error) {
	contract, err := bindUSDCCredits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCCreditsTransactor{contract: contract}, nil
}

// NewUSDCCreditsFilterer creates a new log filterer instance of USDCCredits, bound to a specific deployed contract.
func NewUSDCCreditsFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCCreditsFilterer, error) {
	contract, err := bindUSDCCredits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCCreditsFilterer{contract: contract}, nil
}

// bindUSDCCredits binds a generic wrapper to an already deployed contract.
func bindUSDCCredits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDCCreditsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDCCredits *USDCCreditsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCCredits.Contract.USDCCreditsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDCCredits *USDCCreditsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCCredits.Contract.USDCCreditsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDCCredits *USDCCreditsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCCredits.Contract.USDCCreditsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDCCredits *USDCCreditsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCCredits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDCCredits *USDCCreditsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCCredits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDCCredits *USDCCreditsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCCredits.Contract.contract.Transact(opts, method, params...)
}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (_USDCCredits *USDCCreditsCaller) MinimumDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCCredits.contract.Call(opts, &out, "minimumDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (_USDCCredits *USDCCreditsSession) MinimumDeposit() (*big.Int, error) {
	return _USDCCredits.Contract.MinimumDeposit(&_USDCCredits.CallOpts)
}

// MinimumDeposit is a free data retrieval call binding the contract method 0x636bfbab.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (_USDCCredits *USDCCreditsCallerSession) MinimumDeposit() (*big.Int, error) {
	return _USDCCredits.Contract.MinimumDeposit(&_USDCCredits.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCCredits *USDCCreditsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCCredits.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCCredits *USDCCreditsSession) Owner() (common.Address, error) {
	return _USDCCredits.Contract.Owner(&_USDCCredits.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDCCredits *USDCCreditsCallerSession) Owner() (common.Address, error) {
	return _USDCCredits.Contract.Owner(&_USDCCredits.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_USDCCredits *USDCCreditsCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCCredits.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_USDCCredits *USDCCreditsSession) Treasury() (common.Address, error) {
	return _USDCCredits.Contract.Treasury(&_USDCCredits.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_USDCCredits *USDCCreditsCallerSession) Treasury() (common.Address, error) {
	return _USDCCredits.Contract.Treasury(&_USDCCredits.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCCredits *USDCCreditsCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCCredits.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCCredits *USDCCreditsSession) Usdc() (common.Address, error) {
	return _USDCCredits.Contract.Usdc(&_USDCCredits.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_USDCCredits *USDCCreditsCallerSession) Usdc() (common.Address, error) {
	return _USDCCredits.Contract.Usdc(&_USDCCredits.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (_USDCCredits *USDCCreditsTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _minimumDeposit *big.Int) (*types.Transaction, error) {
	return _USDCCredits.contract.Transact(opts, "initialize", initialOwner, _minimumDeposit)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (_USDCCredits *USDCCreditsSession) Initialize(initialOwner common.Address, _minimumDeposit *big.Int) (*types.Transaction, error) {
	return _USDCCredits.Contract.Initialize(&_USDCCredits.TransactOpts, initialOwner, _minimumDeposit)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (_USDCCredits *USDCCreditsTransactorSession) Initialize(initialOwner common.Address, _minimumDeposit *big.Int) (*types.Transaction, error) {
	return _USDCCredits.Contract.Initialize(&_USDCCredits.TransactOpts, initialOwner, _minimumDeposit)
}

// PurchaseCredits is a paid mutator transaction binding the contract method 0xbef101fb.
//
// Solidity: function purchaseCredits(uint256 amount) returns()
func (_USDCCredits *USDCCreditsTransactor) PurchaseCredits(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _USDCCredits.contract.Transact(opts, "purchaseCredits", amount)
}

// PurchaseCredits is a paid mutator transaction binding the contract method 0xbef101fb.
//
// Solidity: function purchaseCredits(uint256 amount) returns()
func (_USDCCredits *USDCCreditsSession) PurchaseCredits(amount *big.Int) (*types.Transaction, error) {
	return _USDCCredits.Contract.PurchaseCredits(&_USDCCredits.TransactOpts, amount)
}

// PurchaseCredits is a paid mutator transaction binding the contract method 0xbef101fb.
//
// Solidity: function purchaseCredits(uint256 amount) returns()
func (_USDCCredits *USDCCreditsTransactorSession) PurchaseCredits(amount *big.Int) (*types.Transaction, error) {
	return _USDCCredits.Contract.PurchaseCredits(&_USDCCredits.TransactOpts, amount)
}

// PurchaseCreditsFor is a paid mutator transaction binding the contract method 0xdfeddbc3.
//
// Solidity: function purchaseCreditsFor(uint256 amount, address account) returns()
func (_USDCCredits *USDCCreditsTransactor) PurchaseCreditsFor(opts *bind.TransactOpts, amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _USDCCredits.contract.Transact(opts, "purchaseCreditsFor", amount, account)
}

// PurchaseCreditsFor is a paid mutator transaction binding the contract method 0xdfeddbc3.
//
// Solidity: function purchaseCreditsFor(uint256 amount, address account) returns()
func (_USDCCredits *USDCCreditsSession) PurchaseCreditsFor(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _USDCCredits.Contract.PurchaseCreditsFor(&_USDCCredits.TransactOpts, amount, account)
}

// PurchaseCreditsFor is a paid mutator transaction binding the contract method 0xdfeddbc3.
//
// Solidity: function purchaseCreditsFor(uint256 amount, address account) returns()
func (_USDCCredits *USDCCreditsTransactorSession) PurchaseCreditsFor(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _USDCCredits.Contract.PurchaseCreditsFor(&_USDCCredits.TransactOpts, amount, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDCCredits *USDCCreditsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCCredits.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDCCredits *USDCCreditsSession) RenounceOwnership() (*types.Transaction, error) {
	return _USDCCredits.Contract.RenounceOwnership(&_USDCCredits.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDCCredits *USDCCreditsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _USDCCredits.Contract.RenounceOwnership(&_USDCCredits.TransactOpts)
}

// SetMinimumDepositFor is a paid mutator transaction binding the contract method 0x7110a9f6.
//
// Solidity: function setMinimumDepositFor(uint256 newMinimum) returns()
func (_USDCCredits *USDCCreditsTransactor) SetMinimumDepositFor(opts *bind.TransactOpts, newMinimum *big.Int) (*types.Transaction, error) {
	return _USDCCredits.contract.Transact(opts, "setMinimumDepositFor", newMinimum)
}

// SetMinimumDepositFor is a paid mutator transaction binding the contract method 0x7110a9f6.
//
// Solidity: function setMinimumDepositFor(uint256 newMinimum) returns()
func (_USDCCredits *USDCCreditsSession) SetMinimumDepositFor(newMinimum *big.Int) (*types.Transaction, error) {
	return _USDCCredits.Contract.SetMinimumDepositFor(&_USDCCredits.TransactOpts, newMinimum)
}

// SetMinimumDepositFor is a paid mutator transaction binding the contract method 0x7110a9f6.
//
// Solidity: function setMinimumDepositFor(uint256 newMinimum) returns()
func (_USDCCredits *USDCCreditsTransactorSession) SetMinimumDepositFor(newMinimum *big.Int) (*types.Transaction, error) {
	return _USDCCredits.Contract.SetMinimumDepositFor(&_USDCCredits.TransactOpts, newMinimum)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_USDCCredits *USDCCreditsTransactor) Sweep(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _USDCCredits.contract.Transact(opts, "sweep", token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_USDCCredits *USDCCreditsSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _USDCCredits.Contract.Sweep(&_USDCCredits.TransactOpts, token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address token) returns()
func (_USDCCredits *USDCCreditsTransactorSession) Sweep(token common.Address) (*types.Transaction, error) {
	return _USDCCredits.Contract.Sweep(&_USDCCredits.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDCCredits *USDCCreditsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _USDCCredits.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDCCredits *USDCCreditsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDCCredits.Contract.TransferOwnership(&_USDCCredits.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDCCredits *USDCCreditsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDCCredits.Contract.TransferOwnership(&_USDCCredits.TransactOpts, newOwner)
}

// USDCCreditsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the USDCCredits contract.
type USDCCreditsDepositIterator struct {
	Event *USDCCreditsDeposit // Event containing the contract specifics and raw log

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
func (it *USDCCreditsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCCreditsDeposit)
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
		it.Event = new(USDCCreditsDeposit)
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
func (it *USDCCreditsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCCreditsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCCreditsDeposit represents a Deposit event raised by the USDCCredits contract.
type USDCCreditsDeposit struct {
	Depositor common.Address
	Account   common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed depositor, address indexed account, uint256 amount)
func (_USDCCredits *USDCCreditsFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address, account []common.Address) (*USDCCreditsDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _USDCCredits.contract.FilterLogs(opts, "Deposit", depositorRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &USDCCreditsDepositIterator{contract: _USDCCredits.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed depositor, address indexed account, uint256 amount)
func (_USDCCredits *USDCCreditsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *USDCCreditsDeposit, depositor []common.Address, account []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _USDCCredits.contract.WatchLogs(opts, "Deposit", depositorRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCCreditsDeposit)
				if err := _USDCCredits.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_USDCCredits *USDCCreditsFilterer) ParseDeposit(log types.Log) (*USDCCreditsDeposit, error) {
	event := new(USDCCreditsDeposit)
	if err := _USDCCredits.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDCCreditsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the USDCCredits contract.
type USDCCreditsInitializedIterator struct {
	Event *USDCCreditsInitialized // Event containing the contract specifics and raw log

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
func (it *USDCCreditsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCCreditsInitialized)
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
		it.Event = new(USDCCreditsInitialized)
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
func (it *USDCCreditsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCCreditsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCCreditsInitialized represents a Initialized event raised by the USDCCredits contract.
type USDCCreditsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_USDCCredits *USDCCreditsFilterer) FilterInitialized(opts *bind.FilterOpts) (*USDCCreditsInitializedIterator, error) {

	logs, sub, err := _USDCCredits.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &USDCCreditsInitializedIterator{contract: _USDCCredits.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_USDCCredits *USDCCreditsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *USDCCreditsInitialized) (event.Subscription, error) {

	logs, sub, err := _USDCCredits.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCCreditsInitialized)
				if err := _USDCCredits.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_USDCCredits *USDCCreditsFilterer) ParseInitialized(log types.Log) (*USDCCreditsInitialized, error) {
	event := new(USDCCreditsInitialized)
	if err := _USDCCredits.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDCCreditsMinimumDepositSetIterator is returned from FilterMinimumDepositSet and is used to iterate over the raw logs and unpacked data for MinimumDepositSet events raised by the USDCCredits contract.
type USDCCreditsMinimumDepositSetIterator struct {
	Event *USDCCreditsMinimumDepositSet // Event containing the contract specifics and raw log

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
func (it *USDCCreditsMinimumDepositSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCCreditsMinimumDepositSet)
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
		it.Event = new(USDCCreditsMinimumDepositSet)
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
func (it *USDCCreditsMinimumDepositSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCCreditsMinimumDepositSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCCreditsMinimumDepositSet represents a MinimumDepositSet event raised by the USDCCredits contract.
type USDCCreditsMinimumDepositSet struct {
	OldMinimum *big.Int
	NewMinimum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinimumDepositSet is a free log retrieval operation binding the contract event 0x5da812c62f7dbf913260a291f7a1d45cb11ad9527d416f5affc1c3d348ea277e.
//
// Solidity: event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum)
func (_USDCCredits *USDCCreditsFilterer) FilterMinimumDepositSet(opts *bind.FilterOpts) (*USDCCreditsMinimumDepositSetIterator, error) {

	logs, sub, err := _USDCCredits.contract.FilterLogs(opts, "MinimumDepositSet")
	if err != nil {
		return nil, err
	}
	return &USDCCreditsMinimumDepositSetIterator{contract: _USDCCredits.contract, event: "MinimumDepositSet", logs: logs, sub: sub}, nil
}

// WatchMinimumDepositSet is a free log subscription operation binding the contract event 0x5da812c62f7dbf913260a291f7a1d45cb11ad9527d416f5affc1c3d348ea277e.
//
// Solidity: event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum)
func (_USDCCredits *USDCCreditsFilterer) WatchMinimumDepositSet(opts *bind.WatchOpts, sink chan<- *USDCCreditsMinimumDepositSet) (event.Subscription, error) {

	logs, sub, err := _USDCCredits.contract.WatchLogs(opts, "MinimumDepositSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCCreditsMinimumDepositSet)
				if err := _USDCCredits.contract.UnpackLog(event, "MinimumDepositSet", log); err != nil {
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
func (_USDCCredits *USDCCreditsFilterer) ParseMinimumDepositSet(log types.Log) (*USDCCreditsMinimumDepositSet, error) {
	event := new(USDCCreditsMinimumDepositSet)
	if err := _USDCCredits.contract.UnpackLog(event, "MinimumDepositSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDCCreditsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the USDCCredits contract.
type USDCCreditsOwnershipTransferredIterator struct {
	Event *USDCCreditsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *USDCCreditsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCCreditsOwnershipTransferred)
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
		it.Event = new(USDCCreditsOwnershipTransferred)
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
func (it *USDCCreditsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCCreditsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCCreditsOwnershipTransferred represents a OwnershipTransferred event raised by the USDCCredits contract.
type USDCCreditsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDCCredits *USDCCreditsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*USDCCreditsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USDCCredits.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &USDCCreditsOwnershipTransferredIterator{contract: _USDCCredits.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDCCredits *USDCCreditsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCCreditsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USDCCredits.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCCreditsOwnershipTransferred)
				if err := _USDCCredits.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_USDCCredits *USDCCreditsFilterer) ParseOwnershipTransferred(log types.Log) (*USDCCreditsOwnershipTransferred, error) {
	event := new(USDCCreditsOwnershipTransferred)
	if err := _USDCCredits.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
