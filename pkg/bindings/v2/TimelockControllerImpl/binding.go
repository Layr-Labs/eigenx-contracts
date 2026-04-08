// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TimelockControllerImpl

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

// TimelockControllerImplMetaData contains all meta data concerning the TimelockControllerImpl contract.
var TimelockControllerImplMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"CANCELLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXECUTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROPOSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TIMELOCK_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getMinDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestamp\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOperation\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hashOperationBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"minDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"executors\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperation\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationDone\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationPending\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperationReady\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"schedule\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"scheduleBatch\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallExecuted\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSalt\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallScheduled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"predecessor\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"delay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Cancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinDelayChange\",\"inputs\":[{\"name\":\"oldDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	ID:  "TimelockControllerImpl",
}

// TimelockControllerImpl is an auto generated Go binding around an Ethereum contract.
type TimelockControllerImpl struct {
	abi abi.ABI
}

// NewTimelockControllerImpl creates a new instance of TimelockControllerImpl.
func NewTimelockControllerImpl() *TimelockControllerImpl {
	parsed, err := TimelockControllerImplMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TimelockControllerImpl{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TimelockControllerImpl) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCANCELLERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb08e51c0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) PackCANCELLERROLE() []byte {
	enc, err := timelockControllerImpl.abi.Pack("CANCELLER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCANCELLERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb08e51c0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) TryPackCANCELLERROLE() ([]byte, error) {
	return timelockControllerImpl.abi.Pack("CANCELLER_ROLE")
}

// UnpackCANCELLERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) UnpackCANCELLERROLE(data []byte) ([32]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("CANCELLER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) PackDEFAULTADMINROLE() []byte {
	enc, err := timelockControllerImpl.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) TryPackDEFAULTADMINROLE() ([]byte, error) {
	return timelockControllerImpl.abi.Pack("DEFAULT_ADMIN_ROLE")
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackEXECUTORROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07bd0265.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) PackEXECUTORROLE() []byte {
	enc, err := timelockControllerImpl.abi.Pack("EXECUTOR_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEXECUTORROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07bd0265.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) TryPackEXECUTORROLE() ([]byte, error) {
	return timelockControllerImpl.abi.Pack("EXECUTOR_ROLE")
}

// UnpackEXECUTORROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) UnpackEXECUTORROLE(data []byte) ([32]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("EXECUTOR_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackPROPOSERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f61f4f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) PackPROPOSERROLE() []byte {
	enc, err := timelockControllerImpl.abi.Pack("PROPOSER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPROPOSERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f61f4f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) TryPackPROPOSERROLE() ([]byte, error) {
	return timelockControllerImpl.abi.Pack("PROPOSER_ROLE")
}

// UnpackPROPOSERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) UnpackPROPOSERROLE(data []byte) ([32]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("PROPOSER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackTIMELOCKADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d3cf6fc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) PackTIMELOCKADMINROLE() []byte {
	enc, err := timelockControllerImpl.abi.Pack("TIMELOCK_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTIMELOCKADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d3cf6fc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) TryPackTIMELOCKADMINROLE() ([]byte, error) {
	return timelockControllerImpl.abi.Pack("TIMELOCK_ADMIN_ROLE")
}

// UnpackTIMELOCKADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0d3cf6fc.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) UnpackTIMELOCKADMINROLE(data []byte) ([32]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("TIMELOCK_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d252f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancel(bytes32 id) returns()
func (timelockControllerImpl *TimelockControllerImpl) PackCancel(id [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("cancel", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d252f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancel(bytes32 id) returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackCancel(id [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("cancel", id)
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x134008d3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (timelockControllerImpl *TimelockControllerImpl) PackExecute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("execute", target, value, payload, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x134008d3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackExecute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("execute", target, value, payload, predecessor, salt)
}

// PackExecuteBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe38335e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (timelockControllerImpl *TimelockControllerImpl) PackExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("executeBatch", targets, values, payloads, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe38335e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("executeBatch", targets, values, payloads, predecessor, salt)
}

// PackGetMinDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf27a0c92.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinDelay() view returns(uint256)
func (timelockControllerImpl *TimelockControllerImpl) PackGetMinDelay() []byte {
	enc, err := timelockControllerImpl.abi.Pack("getMinDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf27a0c92.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinDelay() view returns(uint256)
func (timelockControllerImpl *TimelockControllerImpl) TryPackGetMinDelay() ([]byte, error) {
	return timelockControllerImpl.abi.Pack("getMinDelay")
}

// UnpackGetMinDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (timelockControllerImpl *TimelockControllerImpl) UnpackGetMinDelay(data []byte) (*big.Int, error) {
	out, err := timelockControllerImpl.abi.Unpack("getMinDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) TryPackGetRoleAdmin(role [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("getRoleAdmin", role)
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd45c4435.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (timelockControllerImpl *TimelockControllerImpl) PackGetTimestamp(id [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("getTimestamp", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd45c4435.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (timelockControllerImpl *TimelockControllerImpl) TryPackGetTimestamp(id [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("getTimestamp", id)
}

// UnpackGetTimestamp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (timelockControllerImpl *TimelockControllerImpl) UnpackGetTimestamp(data []byte) (*big.Int, error) {
	out, err := timelockControllerImpl.abi.Unpack("getTimestamp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (timelockControllerImpl *TimelockControllerImpl) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := timelockControllerImpl.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackGrantRole(role [32]byte, account common.Address) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("grantRole", role, account)
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := timelockControllerImpl.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) TryPackHasRole(role [32]byte, account common.Address) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("hasRole", role, account)
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) UnpackHasRole(data []byte) (bool, error) {
	out, err := timelockControllerImpl.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHashOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8065657f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) PackHashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("hashOperation", target, value, data, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8065657f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) TryPackHashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("hashOperation", target, value, data, predecessor, salt)
}

// UnpackHashOperation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) UnpackHashOperation(data []byte) ([32]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("hashOperation", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackHashOperationBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1c5f427.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) PackHashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("hashOperationBatch", targets, values, payloads, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashOperationBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1c5f427.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) TryPackHashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("hashOperationBatch", targets, values, payloads, predecessor, salt)
}

// UnpackHashOperationBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timelockControllerImpl *TimelockControllerImpl) UnpackHashOperationBatch(data []byte) ([32]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("hashOperationBatch", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4c4c7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (timelockControllerImpl *TimelockControllerImpl) PackInitialize(minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) []byte {
	enc, err := timelockControllerImpl.abi.Pack("initialize", minDelay, proposers, executors, admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4c4c7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackInitialize(minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("initialize", minDelay, proposers, executors, admin)
}

// PackIsOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31d50750.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) PackIsOperation(id [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("isOperation", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31d50750.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) TryPackIsOperation(id [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("isOperation", id)
}

// UnpackIsOperation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) UnpackIsOperation(data []byte) (bool, error) {
	out, err := timelockControllerImpl.abi.Unpack("isOperation", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOperationDone is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab0f529.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) PackIsOperationDone(id [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("isOperationDone", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperationDone is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab0f529.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) TryPackIsOperationDone(id [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("isOperationDone", id)
}

// UnpackIsOperationDone is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) UnpackIsOperationDone(data []byte) (bool, error) {
	out, err := timelockControllerImpl.abi.Unpack("isOperationDone", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOperationPending is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x584b153e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) PackIsOperationPending(id [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("isOperationPending", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperationPending is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x584b153e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) TryPackIsOperationPending(id [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("isOperationPending", id)
}

// UnpackIsOperationPending is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) UnpackIsOperationPending(data []byte) (bool, error) {
	out, err := timelockControllerImpl.abi.Unpack("isOperationPending", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOperationReady is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13bc9f20.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) PackIsOperationReady(id [32]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("isOperationReady", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperationReady is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13bc9f20.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) TryPackIsOperationReady(id [32]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("isOperationReady", id)
}

// UnpackIsOperationReady is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) UnpackIsOperationReady(data []byte) (bool, error) {
	out, err := timelockControllerImpl.abi.Unpack("isOperationReady", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) TryPackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (timelockControllerImpl *TimelockControllerImpl) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := timelockControllerImpl.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (timelockControllerImpl *TimelockControllerImpl) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := timelockControllerImpl.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackRenounceRole(role [32]byte, account common.Address) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("renounceRole", role, account)
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (timelockControllerImpl *TimelockControllerImpl) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := timelockControllerImpl.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackRevokeRole(role [32]byte, account common.Address) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("revokeRole", role, account)
}

// PackSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01d5062a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (timelockControllerImpl *TimelockControllerImpl) PackSchedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) []byte {
	enc, err := timelockControllerImpl.abi.Pack("schedule", target, value, data, predecessor, salt, delay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01d5062a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackSchedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("schedule", target, value, data, predecessor, salt, delay)
}

// PackScheduleBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f2a0bb0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (timelockControllerImpl *TimelockControllerImpl) PackScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) []byte {
	enc, err := timelockControllerImpl.abi.Pack("scheduleBatch", targets, values, payloads, predecessor, salt, delay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackScheduleBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f2a0bb0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("scheduleBatch", targets, values, payloads, predecessor, salt, delay)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := timelockControllerImpl.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (timelockControllerImpl *TimelockControllerImpl) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := timelockControllerImpl.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackUpdateDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64d62353.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (timelockControllerImpl *TimelockControllerImpl) PackUpdateDelay(newDelay *big.Int) []byte {
	enc, err := timelockControllerImpl.abi.Pack("updateDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64d62353.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (timelockControllerImpl *TimelockControllerImpl) TryPackUpdateDelay(newDelay *big.Int) ([]byte, error) {
	return timelockControllerImpl.abi.Pack("updateDelay", newDelay)
}

// TimelockControllerImplCallExecuted represents a CallExecuted event raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplCallExecutedEventName = "CallExecuted"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplCallExecuted) ContractEventName() string {
	return TimelockControllerImplCallExecutedEventName
}

// UnpackCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (timelockControllerImpl *TimelockControllerImpl) UnpackCallExecutedEvent(log *types.Log) (*TimelockControllerImplCallExecuted, error) {
	event := "CallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplCallExecuted)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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

// TimelockControllerImplCallSalt represents a CallSalt event raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallSalt struct {
	Id   [32]byte
	Salt [32]byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplCallSaltEventName = "CallSalt"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplCallSalt) ContractEventName() string {
	return TimelockControllerImplCallSaltEventName
}

// UnpackCallSaltEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (timelockControllerImpl *TimelockControllerImpl) UnpackCallSaltEvent(log *types.Log) (*TimelockControllerImplCallSalt, error) {
	event := "CallSalt"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplCallSalt)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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

// TimelockControllerImplCallScheduled represents a CallScheduled event raised by the TimelockControllerImpl contract.
type TimelockControllerImplCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplCallScheduledEventName = "CallScheduled"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplCallScheduled) ContractEventName() string {
	return TimelockControllerImplCallScheduledEventName
}

// UnpackCallScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (timelockControllerImpl *TimelockControllerImpl) UnpackCallScheduledEvent(log *types.Log) (*TimelockControllerImplCallScheduled, error) {
	event := "CallScheduled"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplCallScheduled)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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

// TimelockControllerImplCancelled represents a Cancelled event raised by the TimelockControllerImpl contract.
type TimelockControllerImplCancelled struct {
	Id  [32]byte
	Raw *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplCancelledEventName = "Cancelled"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplCancelled) ContractEventName() string {
	return TimelockControllerImplCancelledEventName
}

// UnpackCancelledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (timelockControllerImpl *TimelockControllerImpl) UnpackCancelledEvent(log *types.Log) (*TimelockControllerImplCancelled, error) {
	event := "Cancelled"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplCancelled)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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

// TimelockControllerImplInitialized represents a Initialized event raised by the TimelockControllerImpl contract.
type TimelockControllerImplInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplInitialized) ContractEventName() string {
	return TimelockControllerImplInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (timelockControllerImpl *TimelockControllerImpl) UnpackInitializedEvent(log *types.Log) (*TimelockControllerImplInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplInitialized)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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

// TimelockControllerImplMinDelayChange represents a MinDelayChange event raised by the TimelockControllerImpl contract.
type TimelockControllerImplMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplMinDelayChangeEventName = "MinDelayChange"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplMinDelayChange) ContractEventName() string {
	return TimelockControllerImplMinDelayChangeEventName
}

// UnpackMinDelayChangeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (timelockControllerImpl *TimelockControllerImpl) UnpackMinDelayChangeEvent(log *types.Log) (*TimelockControllerImplMinDelayChange, error) {
	event := "MinDelayChange"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplMinDelayChange)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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

// TimelockControllerImplRoleAdminChanged represents a RoleAdminChanged event raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplRoleAdminChanged) ContractEventName() string {
	return TimelockControllerImplRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (timelockControllerImpl *TimelockControllerImpl) UnpackRoleAdminChangedEvent(log *types.Log) (*TimelockControllerImplRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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

// TimelockControllerImplRoleGranted represents a RoleGranted event raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplRoleGranted) ContractEventName() string {
	return TimelockControllerImplRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (timelockControllerImpl *TimelockControllerImpl) UnpackRoleGrantedEvent(log *types.Log) (*TimelockControllerImplRoleGranted, error) {
	event := "RoleGranted"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplRoleGranted)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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

// TimelockControllerImplRoleRevoked represents a RoleRevoked event raised by the TimelockControllerImpl contract.
type TimelockControllerImplRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TimelockControllerImplRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (TimelockControllerImplRoleRevoked) ContractEventName() string {
	return TimelockControllerImplRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (timelockControllerImpl *TimelockControllerImpl) UnpackRoleRevokedEvent(log *types.Log) (*TimelockControllerImplRoleRevoked, error) {
	event := "RoleRevoked"
	if len(log.Topics) == 0 || log.Topics[0] != timelockControllerImpl.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockControllerImplRoleRevoked)
	if len(log.Data) > 0 {
		if err := timelockControllerImpl.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelockControllerImpl.abi.Events[event].Inputs {
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
