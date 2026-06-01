// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ComputeOperator

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

// ComputeOperatorMetaData contains all meta data concerning the ComputeOperator contract.
var ComputeOperatorMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_appController\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_computeAVSRegistrar\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeAVSRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"operatorMetadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerForOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAppController\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotComputeAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	ID:  "ComputeOperator",
}

// ComputeOperator is an auto generated Go binding around an Ethereum contract.
type ComputeOperator struct {
	abi abi.ABI
}

// NewComputeOperator creates a new instance of ComputeOperator.
func NewComputeOperator() *ComputeOperator {
	parsed, err := ComputeOperatorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ComputeOperator{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ComputeOperator) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(string _version, address _delegationManager, address _allocationManager, address _permissionController, address _appController, address _computeAVSRegistrar) returns()
func (computeOperator *ComputeOperator) PackConstructor(_version string, _delegationManager common.Address, _allocationManager common.Address, _permissionController common.Address, _appController common.Address, _computeAVSRegistrar common.Address) []byte {
	enc, err := computeOperator.abi.Pack("", _version, _delegationManager, _allocationManager, _permissionController, _appController, _computeAVSRegistrar)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllocationManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca8aa7c7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allocationManager() view returns(address)
func (computeOperator *ComputeOperator) PackAllocationManager() []byte {
	enc, err := computeOperator.abi.Pack("allocationManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllocationManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca8aa7c7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allocationManager() view returns(address)
func (computeOperator *ComputeOperator) TryPackAllocationManager() ([]byte, error) {
	return computeOperator.abi.Pack("allocationManager")
}

// UnpackAllocationManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (computeOperator *ComputeOperator) UnpackAllocationManager(data []byte) (common.Address, error) {
	out, err := computeOperator.abi.Unpack("allocationManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAppController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ede3ba7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function appController() view returns(address)
func (computeOperator *ComputeOperator) PackAppController() []byte {
	enc, err := computeOperator.abi.Pack("appController")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAppController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ede3ba7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function appController() view returns(address)
func (computeOperator *ComputeOperator) TryPackAppController() ([]byte, error) {
	return computeOperator.abi.Pack("appController")
}

// UnpackAppController is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ede3ba7.
//
// Solidity: function appController() view returns(address)
func (computeOperator *ComputeOperator) UnpackAppController(data []byte) (common.Address, error) {
	out, err := computeOperator.abi.Unpack("appController", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackComputeAVSRegistrar is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef6d92c6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (computeOperator *ComputeOperator) PackComputeAVSRegistrar() []byte {
	enc, err := computeOperator.abi.Pack("computeAVSRegistrar")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackComputeAVSRegistrar is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef6d92c6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (computeOperator *ComputeOperator) TryPackComputeAVSRegistrar() ([]byte, error) {
	return computeOperator.abi.Pack("computeAVSRegistrar")
}

// UnpackComputeAVSRegistrar is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (computeOperator *ComputeOperator) UnpackComputeAVSRegistrar(data []byte) (common.Address, error) {
	out, err := computeOperator.abi.Unpack("computeAVSRegistrar", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDelegationManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea4d3c9b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function delegationManager() view returns(address)
func (computeOperator *ComputeOperator) PackDelegationManager() []byte {
	enc, err := computeOperator.abi.Pack("delegationManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDelegationManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea4d3c9b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function delegationManager() view returns(address)
func (computeOperator *ComputeOperator) TryPackDelegationManager() ([]byte, error) {
	return computeOperator.abi.Pack("delegationManager")
}

// UnpackDelegationManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (computeOperator *ComputeOperator) UnpackDelegationManager(data []byte) (common.Address, error) {
	out, err := computeOperator.abi.Unpack("delegationManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf62d1888.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(string operatorMetadataURI) returns()
func (computeOperator *ComputeOperator) PackInitialize(operatorMetadataURI string) []byte {
	enc, err := computeOperator.abi.Pack("initialize", operatorMetadataURI)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf62d1888.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(string operatorMetadataURI) returns()
func (computeOperator *ComputeOperator) TryPackInitialize(operatorMetadataURI string) ([]byte, error) {
	return computeOperator.abi.Pack("initialize", operatorMetadataURI)
}

// PackPermissionController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4657e26a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function permissionController() view returns(address)
func (computeOperator *ComputeOperator) PackPermissionController() []byte {
	enc, err := computeOperator.abi.Pack("permissionController")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPermissionController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4657e26a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function permissionController() view returns(address)
func (computeOperator *ComputeOperator) TryPackPermissionController() ([]byte, error) {
	return computeOperator.abi.Pack("permissionController")
}

// UnpackPermissionController is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (computeOperator *ComputeOperator) UnpackPermissionController(data []byte) (common.Address, error) {
	out, err := computeOperator.abi.Unpack("permissionController", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRegisterForOperatorSet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4ca62978.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerForOperatorSet(uint32 operatorSetId) returns()
func (computeOperator *ComputeOperator) PackRegisterForOperatorSet(operatorSetId uint32) []byte {
	enc, err := computeOperator.abi.Pack("registerForOperatorSet", operatorSetId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterForOperatorSet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4ca62978.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerForOperatorSet(uint32 operatorSetId) returns()
func (computeOperator *ComputeOperator) TryPackRegisterForOperatorSet(operatorSetId uint32) ([]byte, error) {
	return computeOperator.abi.Pack("registerForOperatorSet", operatorSetId)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(string)
func (computeOperator *ComputeOperator) PackVersion() []byte {
	enc, err := computeOperator.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function version() view returns(string)
func (computeOperator *ComputeOperator) TryPackVersion() ([]byte, error) {
	return computeOperator.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (computeOperator *ComputeOperator) UnpackVersion(data []byte) (string, error) {
	out, err := computeOperator.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// ComputeOperatorInitialized represents a Initialized event raised by the ComputeOperator contract.
type ComputeOperatorInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const ComputeOperatorInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ComputeOperatorInitialized) ContractEventName() string {
	return ComputeOperatorInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (computeOperator *ComputeOperator) UnpackInitializedEvent(log *types.Log) (*ComputeOperatorInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != computeOperator.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ComputeOperatorInitialized)
	if len(log.Data) > 0 {
		if err := computeOperator.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range computeOperator.abi.Events[event].Inputs {
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
func (computeOperator *ComputeOperator) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], computeOperator.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return computeOperator.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeOperator.abi.Errors["NotAppController"].ID.Bytes()[:4]) {
		return computeOperator.UnpackNotAppControllerError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeOperator.abi.Errors["NotComputeAVS"].ID.Bytes()[:4]) {
		return computeOperator.UnpackNotComputeAVSError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeOperator.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return computeOperator.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ComputeOperatorInvalidShortString represents a InvalidShortString error raised by the ComputeOperator contract.
type ComputeOperatorInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ComputeOperatorInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (computeOperator *ComputeOperator) UnpackInvalidShortStringError(raw []byte) (*ComputeOperatorInvalidShortString, error) {
	out := new(ComputeOperatorInvalidShortString)
	if err := computeOperator.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeOperatorNotAppController represents a NotAppController error raised by the ComputeOperator contract.
type ComputeOperatorNotAppController struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAppController()
func ComputeOperatorNotAppControllerErrorID() common.Hash {
	return common.HexToHash("0x626e239362f9371458de20b06d3fd86b8e3c92ca75e136ef3540651430775634")
}

// UnpackNotAppControllerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAppController()
func (computeOperator *ComputeOperator) UnpackNotAppControllerError(raw []byte) (*ComputeOperatorNotAppController, error) {
	out := new(ComputeOperatorNotAppController)
	if err := computeOperator.abi.UnpackIntoInterface(out, "NotAppController", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeOperatorNotComputeAVS represents a NotComputeAVS error raised by the ComputeOperator contract.
type ComputeOperatorNotComputeAVS struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotComputeAVS()
func ComputeOperatorNotComputeAVSErrorID() common.Hash {
	return common.HexToHash("0x9eb9d2bf028a73ca7659c61046247ceb6984dfc6241d8403d6d05e387081ab96")
}

// UnpackNotComputeAVSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotComputeAVS()
func (computeOperator *ComputeOperator) UnpackNotComputeAVSError(raw []byte) (*ComputeOperatorNotComputeAVS, error) {
	out := new(ComputeOperatorNotComputeAVS)
	if err := computeOperator.abi.UnpackIntoInterface(out, "NotComputeAVS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeOperatorStringTooLong represents a StringTooLong error raised by the ComputeOperator contract.
type ComputeOperatorStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ComputeOperatorStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (computeOperator *ComputeOperator) UnpackStringTooLongError(raw []byte) (*ComputeOperatorStringTooLong, error) {
	out := new(ComputeOperatorStringTooLong)
	if err := computeOperator.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}
