// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TimelockControllerImpl

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

// TimelockControllerImplPendingOp is an auto generated low-level Go binding around an user-defined struct.
type TimelockControllerImplPendingOp struct {
	Id           [32]byte
	Target       common.Address
	Data         []byte
	ExecutableAt *big.Int
}

// TimelockControllerImplMetaData contains all meta data concerning the TimelockControllerImpl contract.
var TimelockControllerImplMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"CANCELLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXECUTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROPOSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TIMELOCK_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getMinDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingOperationIds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingOperations\",\"inputs\":[],\"outputs\":[{\"name\":\"ops\",\"type\":\"tuple[]\",\"internalType\":\"structTimelockControllerImpl.PendingOp[]\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executableAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestamp\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOperation\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hashOperationBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"minDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"executors\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperation\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationDone\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationPending\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationReady\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"schedule\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"scheduleBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallExecuted\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSalt\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallScheduled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Cancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinDelayChange\",\"inputs\":[{\"name\":\"oldDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// TimelockControllerImplABI is the input ABI used to generate the binding from.
// Deprecated: Use TimelockControllerImplMetaData.ABI instead.
var TimelockControllerImplABI = TimelockControllerImplMetaData.ABI

// TimelockControllerImpl is an auto generated Go binding around an Ethereum contract.
type TimelockControllerImpl struct {
	TimelockControllerImplCaller     // Read-only binding to the contract
	TimelockControllerImplTransactor // Write-only binding to the contract
	TimelockControllerImplFilterer   // Log filterer for contract events
}

// TimelockControllerImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimelockControllerImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockControllerImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimelockControllerImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockControllerImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimelockControllerImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockControllerImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimelockControllerImplSession struct {
	Contract     *TimelockControllerImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TimelockControllerImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimelockControllerImplCallerSession struct {
	Contract *TimelockControllerImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TimelockControllerImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimelockControllerImplTransactorSession struct {
	Contract     *TimelockControllerImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TimelockControllerImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimelockControllerImplRaw struct {
	Contract *TimelockControllerImpl // Generic contract binding to access the raw methods on
}

// TimelockControllerImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimelockControllerImplCallerRaw struct {
	Contract *TimelockControllerImplCaller // Generic read-only contract binding to access the raw methods on
}

// TimelockControllerImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimelockControllerImplTransactorRaw struct {
	Contract *TimelockControllerImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimelockControllerImpl creates a new instance of TimelockControllerImpl, bound to a specific deployed contract.
func NewTimelockControllerImpl(address common.Address, backend bind.ContractBackend) (*TimelockControllerImpl, error) {
	contract, err := bindTimelockControllerImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImpl{TimelockControllerImplCaller: TimelockControllerImplCaller{contract: contract}, TimelockControllerImplTransactor: TimelockControllerImplTransactor{contract: contract}, TimelockControllerImplFilterer: TimelockControllerImplFilterer{contract: contract}}, nil
}

// NewTimelockControllerImplCaller creates a new read-only instance of TimelockControllerImpl, bound to a specific deployed contract.
func NewTimelockControllerImplCaller(address common.Address, caller bind.ContractCaller) (*TimelockControllerImplCaller, error) {
	contract, err := bindTimelockControllerImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplCaller{contract: contract}, nil
}

// NewTimelockControllerImplTransactor creates a new write-only instance of TimelockControllerImpl, bound to a specific deployed contract.
func NewTimelockControllerImplTransactor(address common.Address, transactor bind.ContractTransactor) (*TimelockControllerImplTransactor, error) {
	contract, err := bindTimelockControllerImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplTransactor{contract: contract}, nil
}

// NewTimelockControllerImplFilterer creates a new log filterer instance of TimelockControllerImpl, bound to a specific deployed contract.
func NewTimelockControllerImplFilterer(address common.Address, filterer bind.ContractFilterer) (*TimelockControllerImplFilterer, error) {
	contract, err := bindTimelockControllerImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplFilterer{contract: contract}, nil
}

// bindTimelockControllerImpl binds a generic wrapper to an already deployed contract.
func bindTimelockControllerImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimelockControllerImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimelockControllerImpl *TimelockControllerImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimelockControllerImpl.Contract.TimelockControllerImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimelockControllerImpl *TimelockControllerImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.TimelockControllerImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimelockControllerImpl *TimelockControllerImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.TimelockControllerImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimelockControllerImpl *TimelockControllerImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimelockControllerImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimelockControllerImpl *TimelockControllerImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimelockControllerImpl *TimelockControllerImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.contract.Transact(opts, method, params...)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCaller) CANCELLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "CANCELLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplSession) CANCELLERROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.CANCELLERROLE(&_TimelockControllerImpl.CallOpts)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) CANCELLERROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.CANCELLERROLE(&_TimelockControllerImpl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.DEFAULTADMINROLE(&_TimelockControllerImpl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.DEFAULTADMINROLE(&_TimelockControllerImpl.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplSession) EXECUTORROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.EXECUTORROLE(&_TimelockControllerImpl.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.EXECUTORROLE(&_TimelockControllerImpl.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCaller) PROPOSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "PROPOSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplSession) PROPOSERROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.PROPOSERROLE(&_TimelockControllerImpl.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) PROPOSERROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.PROPOSERROLE(&_TimelockControllerImpl.CallOpts)
}

// TIMELOCKADMINROLE is a free data retrieval call binding the contract method 0x0d3cf6fc.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCaller) TIMELOCKADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "TIMELOCK_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TIMELOCKADMINROLE is a free data retrieval call binding the contract method 0x0d3cf6fc.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplSession) TIMELOCKADMINROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.TIMELOCKADMINROLE(&_TimelockControllerImpl.CallOpts)
}

// TIMELOCKADMINROLE is a free data retrieval call binding the contract method 0x0d3cf6fc.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) TIMELOCKADMINROLE() ([32]byte, error) {
	return _TimelockControllerImpl.Contract.TIMELOCKADMINROLE(&_TimelockControllerImpl.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_TimelockControllerImpl *TimelockControllerImplCaller) GetMinDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "getMinDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_TimelockControllerImpl *TimelockControllerImplSession) GetMinDelay() (*big.Int, error) {
	return _TimelockControllerImpl.Contract.GetMinDelay(&_TimelockControllerImpl.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) GetMinDelay() (*big.Int, error) {
	return _TimelockControllerImpl.Contract.GetMinDelay(&_TimelockControllerImpl.CallOpts)
}

// GetPendingOperationIds is a free data retrieval call binding the contract method 0x7b146198.
//
// Solidity: function getPendingOperationIds() view returns(bytes32[])
func (_TimelockControllerImpl *TimelockControllerImplCaller) GetPendingOperationIds(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "getPendingOperationIds")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPendingOperationIds is a free data retrieval call binding the contract method 0x7b146198.
//
// Solidity: function getPendingOperationIds() view returns(bytes32[])
func (_TimelockControllerImpl *TimelockControllerImplSession) GetPendingOperationIds() ([][32]byte, error) {
	return _TimelockControllerImpl.Contract.GetPendingOperationIds(&_TimelockControllerImpl.CallOpts)
}

// GetPendingOperationIds is a free data retrieval call binding the contract method 0x7b146198.
//
// Solidity: function getPendingOperationIds() view returns(bytes32[])
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) GetPendingOperationIds() ([][32]byte, error) {
	return _TimelockControllerImpl.Contract.GetPendingOperationIds(&_TimelockControllerImpl.CallOpts)
}

// GetPendingOperations is a free data retrieval call binding the contract method 0x64440ce5.
//
// Solidity: function getPendingOperations() view returns((bytes32,address,bytes,uint256)[] ops)
func (_TimelockControllerImpl *TimelockControllerImplCaller) GetPendingOperations(opts *bind.CallOpts) ([]TimelockControllerImplPendingOp, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "getPendingOperations")

	if err != nil {
		return *new([]TimelockControllerImplPendingOp), err
	}

	out0 := *abi.ConvertType(out[0], new([]TimelockControllerImplPendingOp)).(*[]TimelockControllerImplPendingOp)

	return out0, err

}

// GetPendingOperations is a free data retrieval call binding the contract method 0x64440ce5.
//
// Solidity: function getPendingOperations() view returns((bytes32,address,bytes,uint256)[] ops)
func (_TimelockControllerImpl *TimelockControllerImplSession) GetPendingOperations() ([]TimelockControllerImplPendingOp, error) {
	return _TimelockControllerImpl.Contract.GetPendingOperations(&_TimelockControllerImpl.CallOpts)
}

// GetPendingOperations is a free data retrieval call binding the contract method 0x64440ce5.
//
// Solidity: function getPendingOperations() view returns((bytes32,address,bytes,uint256)[] ops)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) GetPendingOperations() ([]TimelockControllerImplPendingOp, error) {
	return _TimelockControllerImpl.Contract.GetPendingOperations(&_TimelockControllerImpl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TimelockControllerImpl.Contract.GetRoleAdmin(&_TimelockControllerImpl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TimelockControllerImpl.Contract.GetRoleAdmin(&_TimelockControllerImpl.CallOpts, role)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_TimelockControllerImpl *TimelockControllerImplCaller) GetTimestamp(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "getTimestamp", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_TimelockControllerImpl *TimelockControllerImplSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _TimelockControllerImpl.Contract.GetTimestamp(&_TimelockControllerImpl.CallOpts, id)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _TimelockControllerImpl.Contract.GetTimestamp(&_TimelockControllerImpl.CallOpts, id)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TimelockControllerImpl.Contract.HasRole(&_TimelockControllerImpl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TimelockControllerImpl.Contract.HasRole(&_TimelockControllerImpl.CallOpts, role, account)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCaller) HashOperation(opts *bind.CallOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "hashOperation", target, value, data, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _TimelockControllerImpl.Contract.HashOperation(&_TimelockControllerImpl.CallOpts, target, value, data, predecessor, salt)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _TimelockControllerImpl.Contract.HashOperation(&_TimelockControllerImpl.CallOpts, target, value, data, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCaller) HashOperationBatch(opts *bind.CallOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "hashOperationBatch", targets, values, payloads, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _TimelockControllerImpl.Contract.HashOperationBatch(&_TimelockControllerImpl.CallOpts, targets, values, payloads, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _TimelockControllerImpl.Contract.HashOperationBatch(&_TimelockControllerImpl.CallOpts, targets, values, payloads, predecessor, salt)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCaller) IsOperation(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "isOperation", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplSession) IsOperation(id [32]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.IsOperation(&_TimelockControllerImpl.CallOpts, id)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) IsOperation(id [32]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.IsOperation(&_TimelockControllerImpl.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCaller) IsOperationDone(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "isOperationDone", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplSession) IsOperationDone(id [32]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.IsOperationDone(&_TimelockControllerImpl.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) IsOperationDone(id [32]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.IsOperationDone(&_TimelockControllerImpl.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCaller) IsOperationPending(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "isOperationPending", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplSession) IsOperationPending(id [32]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.IsOperationPending(&_TimelockControllerImpl.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) IsOperationPending(id [32]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.IsOperationPending(&_TimelockControllerImpl.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCaller) IsOperationReady(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "isOperationReady", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplSession) IsOperationReady(id [32]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.IsOperationReady(&_TimelockControllerImpl.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) IsOperationReady(id [32]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.IsOperationReady(&_TimelockControllerImpl.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TimelockControllerImpl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.SupportsInterface(&_TimelockControllerImpl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TimelockControllerImpl *TimelockControllerImplCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TimelockControllerImpl.Contract.SupportsInterface(&_TimelockControllerImpl.CallOpts, interfaceId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) Cancel(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Cancel(&_TimelockControllerImpl.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Cancel(&_TimelockControllerImpl.TransactOpts, id)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "execute", target, value, payload, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) Execute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Execute(&_TimelockControllerImpl.TransactOpts, target, value, payload, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) Execute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Execute(&_TimelockControllerImpl.TransactOpts, target, value, payload, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) ExecuteBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "executeBatch", targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.ExecuteBatch(&_TimelockControllerImpl.TransactOpts, targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.ExecuteBatch(&_TimelockControllerImpl.TransactOpts, targets, values, payloads, predecessor, salt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.GrantRole(&_TimelockControllerImpl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.GrantRole(&_TimelockControllerImpl.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4c4c7b3.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) Initialize(opts *bind.TransactOpts, minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "initialize", minDelay, proposers, executors, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4c4c7b3.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) Initialize(minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Initialize(&_TimelockControllerImpl.TransactOpts, minDelay, proposers, executors, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4c4c7b3.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) Initialize(minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Initialize(&_TimelockControllerImpl.TransactOpts, minDelay, proposers, executors, admin)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.OnERC1155BatchReceived(&_TimelockControllerImpl.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.OnERC1155BatchReceived(&_TimelockControllerImpl.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.OnERC1155Received(&_TimelockControllerImpl.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.OnERC1155Received(&_TimelockControllerImpl.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.OnERC721Received(&_TimelockControllerImpl.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.OnERC721Received(&_TimelockControllerImpl.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.RenounceRole(&_TimelockControllerImpl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.RenounceRole(&_TimelockControllerImpl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.RevokeRole(&_TimelockControllerImpl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.RevokeRole(&_TimelockControllerImpl.TransactOpts, role, account)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) Schedule(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "schedule", target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Schedule(&_TimelockControllerImpl.TransactOpts, target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Schedule(&_TimelockControllerImpl.TransactOpts, target, value, data, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) ScheduleBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "scheduleBatch", targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.ScheduleBatch(&_TimelockControllerImpl.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.ScheduleBatch(&_TimelockControllerImpl.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) UpdateDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.Transact(opts, "updateDelay", newDelay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) UpdateDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.UpdateDelay(&_TimelockControllerImpl.TransactOpts, newDelay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) UpdateDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.UpdateDelay(&_TimelockControllerImpl.TransactOpts, newDelay)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimelockControllerImpl.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimelockControllerImpl *TimelockControllerImplSession) Receive() (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Receive(&_TimelockControllerImpl.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimelockControllerImpl *TimelockControllerImplTransactorSession) Receive() (*types.Transaction, error) {
	return _TimelockControllerImpl.Contract.Receive(&_TimelockControllerImpl.TransactOpts)
}

// TimelockControllerImplCallExecutedIterator is returned from FilterCallExecuted and is used to iterate over the raw logs and unpacked data for CallExecuted events raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallExecutedIterator struct {
	Event *TimelockControllerImplCallExecuted // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplCallExecuted)
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
		it.Event = new(TimelockControllerImplCallExecuted)
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
func (it *TimelockControllerImplCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplCallExecuted represents a CallExecuted event raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallExecuted is a free log retrieval operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterCallExecuted(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*TimelockControllerImplCallExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplCallExecutedIterator{contract: _TimelockControllerImpl.contract, event: "CallExecuted", logs: logs, sub: sub}, nil
}

// WatchCallExecuted is a free log subscription operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplCallExecuted, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplCallExecuted)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "CallExecuted", log); err != nil {
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

// ParseCallExecuted is a log parse operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseCallExecuted(log types.Log) (*TimelockControllerImplCallExecuted, error) {
	event := new(TimelockControllerImplCallExecuted)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "CallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockControllerImplCallSaltIterator is returned from FilterCallSalt and is used to iterate over the raw logs and unpacked data for CallSalt events raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallSaltIterator struct {
	Event *TimelockControllerImplCallSalt // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplCallSaltIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplCallSalt)
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
		it.Event = new(TimelockControllerImplCallSalt)
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
func (it *TimelockControllerImplCallSaltIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplCallSaltIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplCallSalt represents a CallSalt event raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallSalt struct {
	Id   [32]byte
	Salt [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCallSalt is a free log retrieval operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterCallSalt(opts *bind.FilterOpts, id [][32]byte) (*TimelockControllerImplCallSaltIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "CallSalt", idRule)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplCallSaltIterator{contract: _TimelockControllerImpl.contract, event: "CallSalt", logs: logs, sub: sub}, nil
}

// WatchCallSalt is a free log subscription operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchCallSalt(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplCallSalt, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "CallSalt", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplCallSalt)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "CallSalt", log); err != nil {
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

// ParseCallSalt is a log parse operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseCallSalt(log types.Log) (*TimelockControllerImplCallSalt, error) {
	event := new(TimelockControllerImplCallSalt)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "CallSalt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockControllerImplCallScheduledIterator is returned from FilterCallScheduled and is used to iterate over the raw logs and unpacked data for CallScheduled events raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallScheduledIterator struct {
	Event *TimelockControllerImplCallScheduled // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplCallScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplCallScheduled)
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
		it.Event = new(TimelockControllerImplCallScheduled)
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
func (it *TimelockControllerImplCallScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplCallScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplCallScheduled represents a CallScheduled event raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCallScheduled is a free log retrieval operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterCallScheduled(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*TimelockControllerImplCallScheduledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplCallScheduledIterator{contract: _TimelockControllerImpl.contract, event: "CallScheduled", logs: logs, sub: sub}, nil
}

// WatchCallScheduled is a free log subscription operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchCallScheduled(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplCallScheduled, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplCallScheduled)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "CallScheduled", log); err != nil {
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

// ParseCallScheduled is a log parse operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseCallScheduled(log types.Log) (*TimelockControllerImplCallScheduled, error) {
	event := new(TimelockControllerImplCallScheduled)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "CallScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockControllerImplCancelledIterator is returned from FilterCancelled and is used to iterate over the raw logs and unpacked data for Cancelled events raised by the TimelockControllerImpl contract.
type TimelockControllerImplCancelledIterator struct {
	Event *TimelockControllerImplCancelled // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplCancelled)
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
		it.Event = new(TimelockControllerImplCancelled)
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
func (it *TimelockControllerImplCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplCancelled represents a Cancelled event raised by the TimelockControllerImpl contract.
type TimelockControllerImplCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelled is a free log retrieval operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterCancelled(opts *bind.FilterOpts, id [][32]byte) (*TimelockControllerImplCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplCancelledIterator{contract: _TimelockControllerImpl.contract, event: "Cancelled", logs: logs, sub: sub}, nil
}

// WatchCancelled is a free log subscription operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchCancelled(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplCancelled)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "Cancelled", log); err != nil {
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

// ParseCancelled is a log parse operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseCancelled(log types.Log) (*TimelockControllerImplCancelled, error) {
	event := new(TimelockControllerImplCancelled)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "Cancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockControllerImplInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TimelockControllerImpl contract.
type TimelockControllerImplInitializedIterator struct {
	Event *TimelockControllerImplInitialized // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplInitialized)
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
		it.Event = new(TimelockControllerImplInitialized)
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
func (it *TimelockControllerImplInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplInitialized represents a Initialized event raised by the TimelockControllerImpl contract.
type TimelockControllerImplInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterInitialized(opts *bind.FilterOpts) (*TimelockControllerImplInitializedIterator, error) {

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplInitializedIterator{contract: _TimelockControllerImpl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplInitialized) (event.Subscription, error) {

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplInitialized)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseInitialized(log types.Log) (*TimelockControllerImplInitialized, error) {
	event := new(TimelockControllerImplInitialized)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockControllerImplMinDelayChangeIterator is returned from FilterMinDelayChange and is used to iterate over the raw logs and unpacked data for MinDelayChange events raised by the TimelockControllerImpl contract.
type TimelockControllerImplMinDelayChangeIterator struct {
	Event *TimelockControllerImplMinDelayChange // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplMinDelayChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplMinDelayChange)
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
		it.Event = new(TimelockControllerImplMinDelayChange)
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
func (it *TimelockControllerImplMinDelayChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplMinDelayChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplMinDelayChange represents a MinDelayChange event raised by the TimelockControllerImpl contract.
type TimelockControllerImplMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinDelayChange is a free log retrieval operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterMinDelayChange(opts *bind.FilterOpts) (*TimelockControllerImplMinDelayChangeIterator, error) {

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplMinDelayChangeIterator{contract: _TimelockControllerImpl.contract, event: "MinDelayChange", logs: logs, sub: sub}, nil
}

// WatchMinDelayChange is a free log subscription operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchMinDelayChange(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplMinDelayChange) (event.Subscription, error) {

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplMinDelayChange)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
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

// ParseMinDelayChange is a log parse operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseMinDelayChange(log types.Log) (*TimelockControllerImplMinDelayChange, error) {
	event := new(TimelockControllerImplMinDelayChange)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockControllerImplRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleAdminChangedIterator struct {
	Event *TimelockControllerImplRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplRoleAdminChanged)
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
		it.Event = new(TimelockControllerImplRoleAdminChanged)
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
func (it *TimelockControllerImplRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplRoleAdminChanged represents a RoleAdminChanged event raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TimelockControllerImplRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplRoleAdminChangedIterator{contract: _TimelockControllerImpl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplRoleAdminChanged)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseRoleAdminChanged(log types.Log) (*TimelockControllerImplRoleAdminChanged, error) {
	event := new(TimelockControllerImplRoleAdminChanged)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockControllerImplRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleGrantedIterator struct {
	Event *TimelockControllerImplRoleGranted // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplRoleGranted)
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
		it.Event = new(TimelockControllerImplRoleGranted)
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
func (it *TimelockControllerImplRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplRoleGranted represents a RoleGranted event raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TimelockControllerImplRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplRoleGrantedIterator{contract: _TimelockControllerImpl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplRoleGranted)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseRoleGranted(log types.Log) (*TimelockControllerImplRoleGranted, error) {
	event := new(TimelockControllerImplRoleGranted)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockControllerImplRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleRevokedIterator struct {
	Event *TimelockControllerImplRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TimelockControllerImplRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockControllerImplRoleRevoked)
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
		it.Event = new(TimelockControllerImplRoleRevoked)
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
func (it *TimelockControllerImplRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockControllerImplRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockControllerImplRoleRevoked represents a RoleRevoked event raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TimelockControllerImplRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TimelockControllerImplRoleRevokedIterator{contract: _TimelockControllerImpl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TimelockControllerImplRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TimelockControllerImpl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockControllerImplRoleRevoked)
				if err := _TimelockControllerImpl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimelockControllerImpl *TimelockControllerImplFilterer) ParseRoleRevoked(log types.Log) (*TimelockControllerImplRoleRevoked, error) {
	event := new(TimelockControllerImplRoleRevoked)
	if err := _TimelockControllerImpl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
