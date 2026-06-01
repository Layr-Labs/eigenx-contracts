// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ComputeAVSRegistrar

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ComputeAVSRegistrarMetaData contains all meta data concerning the ComputeAVSRegistrar contract.
var ComputeAVSRegistrarMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_appController\",\"type\":\"address\",\"internalType\":\"contractIAppController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorToAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAppController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assignOperatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createOperatorSet\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedOperators\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorAllowed\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextOperatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorFromAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromAllowlist\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAppController\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyInAllowlist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotInAllowlist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetIdNotAssigned\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	ID:  "ComputeAVSRegistrar",
}

// ComputeAVSRegistrar is an auto generated Go binding around an Ethereum contract.
type ComputeAVSRegistrar struct {
	abi abi.ABI
}

// NewComputeAVSRegistrar creates a new instance of ComputeAVSRegistrar.
func NewComputeAVSRegistrar() *ComputeAVSRegistrar {
	parsed, err := ComputeAVSRegistrarMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ComputeAVSRegistrar{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ComputeAVSRegistrar) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(string _version, address _allocationManager, address _permissionController, address _keyRegistrar, address _appController) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackConstructor(_version string, _allocationManager common.Address, _permissionController common.Address, _keyRegistrar common.Address, _appController common.Address) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("", _version, _allocationManager, _permissionController, _keyRegistrar, _appController)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAddOperatorToAllowlist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1017873a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addOperatorToAllowlist((address,uint32) operatorSet, address operator) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackAddOperatorToAllowlist(operatorSet OperatorSet, operator common.Address) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("addOperatorToAllowlist", operatorSet, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddOperatorToAllowlist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1017873a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addOperatorToAllowlist((address,uint32) operatorSet, address operator) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackAddOperatorToAllowlist(operatorSet OperatorSet, operator common.Address) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("addOperatorToAllowlist", operatorSet, operator)
}

// PackAllocationManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca8aa7c7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allocationManager() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackAllocationManager() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("allocationManager")
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
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackAllocationManager() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("allocationManager")
}

// UnpackAllocationManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackAllocationManager(data []byte) (common.Address, error) {
	out, err := computeAVSRegistrar.abi.Unpack("allocationManager", data)
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
func (computeAVSRegistrar *ComputeAVSRegistrar) PackAppController() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("appController")
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
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackAppController() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("appController")
}

// UnpackAppController is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ede3ba7.
//
// Solidity: function appController() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackAppController(data []byte) (common.Address, error) {
	out, err := computeAVSRegistrar.abi.Unpack("appController", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAssignOperatorSetId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a10eef7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function assignOperatorSetId() returns(uint32 operatorSetId)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackAssignOperatorSetId() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("assignOperatorSetId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAssignOperatorSetId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a10eef7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function assignOperatorSetId() returns(uint32 operatorSetId)
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackAssignOperatorSetId() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("assignOperatorSetId")
}

// UnpackAssignOperatorSetId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a10eef7.
//
// Solidity: function assignOperatorSetId() returns(uint32 operatorSetId)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackAssignOperatorSetId(data []byte) (uint32, error) {
	out, err := computeAVSRegistrar.abi.Unpack("assignOperatorSetId", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackAvs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde1164bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function avs() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackAvs() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("avs")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAvs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde1164bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function avs() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackAvs() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("avs")
}

// UnpackAvs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackAvs(data []byte) (common.Address, error) {
	out, err := computeAVSRegistrar.abi.Unpack("avs", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCreateOperatorSet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74da4713.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createOperatorSet(uint32 operatorSetId) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackCreateOperatorSet(operatorSetId uint32) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("createOperatorSet", operatorSetId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateOperatorSet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74da4713.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createOperatorSet(uint32 operatorSetId) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackCreateOperatorSet(operatorSetId uint32) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("createOperatorSet", operatorSetId)
}

// PackDeregisterOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x303ca956.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deregisterOperator(address operator, address , uint32[] operatorSetIds) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackDeregisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("deregisterOperator", operator, arg1, operatorSetIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeregisterOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x303ca956.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deregisterOperator(address operator, address , uint32[] operatorSetIds) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackDeregisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("deregisterOperator", operator, arg1, operatorSetIds)
}

// PackGetAllowedOperators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fe94e16.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (computeAVSRegistrar *ComputeAVSRegistrar) PackGetAllowedOperators(operatorSet OperatorSet) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("getAllowedOperators", operatorSet)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAllowedOperators is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fe94e16.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackGetAllowedOperators(operatorSet OperatorSet) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("getAllowedOperators", operatorSet)
}

// UnpackGetAllowedOperators is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7fe94e16.
//
// Solidity: function getAllowedOperators((address,uint32) operatorSet) view returns(address[])
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackGetAllowedOperators(data []byte) ([]common.Address, error) {
	out, err := computeAVSRegistrar.abi.Unpack("getAllowedOperators", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf62d1888.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(string metadataURI) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackInitialize(metadataURI string) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("initialize", metadataURI)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf62d1888.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(string metadataURI) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackInitialize(metadataURI string) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("initialize", metadataURI)
}

// PackIsOperatorAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf91ff80c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackIsOperatorAllowed(operatorSet OperatorSet, operator common.Address) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("isOperatorAllowed", operatorSet, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperatorAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf91ff80c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackIsOperatorAllowed(operatorSet OperatorSet, operator common.Address) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("isOperatorAllowed", operatorSet, operator)
}

// UnpackIsOperatorAllowed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf91ff80c.
//
// Solidity: function isOperatorAllowed((address,uint32) operatorSet, address operator) view returns(bool)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackIsOperatorAllowed(data []byte) (bool, error) {
	out, err := computeAVSRegistrar.abi.Unpack("isOperatorAllowed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackKeyRegistrar is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ec45c7e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function keyRegistrar() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackKeyRegistrar() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("keyRegistrar")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackKeyRegistrar is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ec45c7e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function keyRegistrar() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackKeyRegistrar() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("keyRegistrar")
}

// UnpackKeyRegistrar is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackKeyRegistrar(data []byte) (common.Address, error) {
	out, err := computeAVSRegistrar.abi.Unpack("keyRegistrar", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackNextOperatorSetId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x833b3fa3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nextOperatorSetId() view returns(uint32)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackNextOperatorSetId() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("nextOperatorSetId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNextOperatorSetId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x833b3fa3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nextOperatorSetId() view returns(uint32)
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackNextOperatorSetId() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("nextOperatorSetId")
}

// UnpackNextOperatorSetId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x833b3fa3.
//
// Solidity: function nextOperatorSetId() view returns(uint32)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackNextOperatorSetId(data []byte) (uint32, error) {
	out, err := computeAVSRegistrar.abi.Unpack("nextOperatorSetId", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackOwner() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackOwner() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOwner(data []byte) (common.Address, error) {
	out, err := computeAVSRegistrar.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPermissionController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4657e26a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function permissionController() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackPermissionController() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("permissionController")
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
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackPermissionController() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("permissionController")
}

// UnpackPermissionController is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackPermissionController(data []byte) (common.Address, error) {
	out, err := computeAVSRegistrar.abi.Unpack("permissionController", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRegisterOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc63fd502.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerOperator(address operator, address , uint32[] operatorSetIds, bytes data) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackRegisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32, data []byte) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("registerOperator", operator, arg1, operatorSetIds, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc63fd502.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerOperator(address operator, address , uint32[] operatorSetIds, bytes data) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackRegisterOperator(operator common.Address, arg1 common.Address, operatorSetIds []uint32, data []byte) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("registerOperator", operator, arg1, operatorSetIds, data)
}

// PackRemoveOperatorFromAllowlist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a4d3d29.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeOperatorFromAllowlist((address,uint32) operatorSet, address operator) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackRemoveOperatorFromAllowlist(operatorSet OperatorSet, operator common.Address) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("removeOperatorFromAllowlist", operatorSet, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveOperatorFromAllowlist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a4d3d29.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeOperatorFromAllowlist((address,uint32) operatorSet, address operator) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackRemoveOperatorFromAllowlist(operatorSet OperatorSet, operator common.Address) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("removeOperatorFromAllowlist", operatorSet, operator)
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackRenounceOwnership() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceOwnership() returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackRenounceOwnership() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("renounceOwnership")
}

// PackSupportsAVS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb5265787.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackSupportsAVS(avs common.Address) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("supportsAVS", avs)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsAVS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb5265787.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackSupportsAVS(avs common.Address) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("supportsAVS", avs)
}

// UnpackSupportsAVS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackSupportsAVS(data []byte) (bool, error) {
	out, err := computeAVSRegistrar.abi.Unpack("supportsAVS", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := computeAVSRegistrar.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("transferOwnership", newOwner)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(string)
func (computeAVSRegistrar *ComputeAVSRegistrar) PackVersion() []byte {
	enc, err := computeAVSRegistrar.abi.Pack("version")
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
func (computeAVSRegistrar *ComputeAVSRegistrar) TryPackVersion() ([]byte, error) {
	return computeAVSRegistrar.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackVersion(data []byte) (string, error) {
	out, err := computeAVSRegistrar.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// ComputeAVSRegistrarInitialized represents a Initialized event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const ComputeAVSRegistrarInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ComputeAVSRegistrarInitialized) ContractEventName() string {
	return ComputeAVSRegistrarInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackInitializedEvent(log *types.Log) (*ComputeAVSRegistrarInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != computeAVSRegistrar.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ComputeAVSRegistrarInitialized)
	if len(log.Data) > 0 {
		if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range computeAVSRegistrar.abi.Events[event].Inputs {
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

// ComputeAVSRegistrarOperatorAddedToAllowlist represents a OperatorAddedToAllowlist event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorAddedToAllowlist struct {
	OperatorSet OperatorSet
	Operator    common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const ComputeAVSRegistrarOperatorAddedToAllowlistEventName = "OperatorAddedToAllowlist"

// ContractEventName returns the user-defined event name.
func (ComputeAVSRegistrarOperatorAddedToAllowlist) ContractEventName() string {
	return ComputeAVSRegistrarOperatorAddedToAllowlistEventName
}

// UnpackOperatorAddedToAllowlistEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperatorAddedToAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOperatorAddedToAllowlistEvent(log *types.Log) (*ComputeAVSRegistrarOperatorAddedToAllowlist, error) {
	event := "OperatorAddedToAllowlist"
	if len(log.Topics) == 0 || log.Topics[0] != computeAVSRegistrar.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ComputeAVSRegistrarOperatorAddedToAllowlist)
	if len(log.Data) > 0 {
		if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range computeAVSRegistrar.abi.Events[event].Inputs {
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

// ComputeAVSRegistrarOperatorDeregistered represents a OperatorDeregistered event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorDeregistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            *types.Log // Blockchain specific contextual infos
}

const ComputeAVSRegistrarOperatorDeregisteredEventName = "OperatorDeregistered"

// ContractEventName returns the user-defined event name.
func (ComputeAVSRegistrarOperatorDeregistered) ContractEventName() string {
	return ComputeAVSRegistrarOperatorDeregisteredEventName
}

// UnpackOperatorDeregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOperatorDeregisteredEvent(log *types.Log) (*ComputeAVSRegistrarOperatorDeregistered, error) {
	event := "OperatorDeregistered"
	if len(log.Topics) == 0 || log.Topics[0] != computeAVSRegistrar.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ComputeAVSRegistrarOperatorDeregistered)
	if len(log.Data) > 0 {
		if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range computeAVSRegistrar.abi.Events[event].Inputs {
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

// ComputeAVSRegistrarOperatorRegistered represents a OperatorRegistered event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorRegistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            *types.Log // Blockchain specific contextual infos
}

const ComputeAVSRegistrarOperatorRegisteredEventName = "OperatorRegistered"

// ContractEventName returns the user-defined event name.
func (ComputeAVSRegistrarOperatorRegistered) ContractEventName() string {
	return ComputeAVSRegistrarOperatorRegisteredEventName
}

// UnpackOperatorRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOperatorRegisteredEvent(log *types.Log) (*ComputeAVSRegistrarOperatorRegistered, error) {
	event := "OperatorRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != computeAVSRegistrar.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ComputeAVSRegistrarOperatorRegistered)
	if len(log.Data) > 0 {
		if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range computeAVSRegistrar.abi.Events[event].Inputs {
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

// ComputeAVSRegistrarOperatorRemovedFromAllowlist represents a OperatorRemovedFromAllowlist event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorRemovedFromAllowlist struct {
	OperatorSet OperatorSet
	Operator    common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const ComputeAVSRegistrarOperatorRemovedFromAllowlistEventName = "OperatorRemovedFromAllowlist"

// ContractEventName returns the user-defined event name.
func (ComputeAVSRegistrarOperatorRemovedFromAllowlist) ContractEventName() string {
	return ComputeAVSRegistrarOperatorRemovedFromAllowlistEventName
}

// UnpackOperatorRemovedFromAllowlistEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperatorRemovedFromAllowlist((address,uint32) indexed operatorSet, address indexed operator)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOperatorRemovedFromAllowlistEvent(log *types.Log) (*ComputeAVSRegistrarOperatorRemovedFromAllowlist, error) {
	event := "OperatorRemovedFromAllowlist"
	if len(log.Topics) == 0 || log.Topics[0] != computeAVSRegistrar.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ComputeAVSRegistrarOperatorRemovedFromAllowlist)
	if len(log.Data) > 0 {
		if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range computeAVSRegistrar.abi.Events[event].Inputs {
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

// ComputeAVSRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ComputeAVSRegistrarOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ComputeAVSRegistrarOwnershipTransferred) ContractEventName() string {
	return ComputeAVSRegistrarOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOwnershipTransferredEvent(log *types.Log) (*ComputeAVSRegistrarOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != computeAVSRegistrar.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ComputeAVSRegistrarOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range computeAVSRegistrar.abi.Events[event].Inputs {
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
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], computeAVSRegistrar.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return computeAVSRegistrar.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeAVSRegistrar.abi.Errors["KeyNotRegistered"].ID.Bytes()[:4]) {
		return computeAVSRegistrar.UnpackKeyNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeAVSRegistrar.abi.Errors["NotAllocationManager"].ID.Bytes()[:4]) {
		return computeAVSRegistrar.UnpackNotAllocationManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeAVSRegistrar.abi.Errors["NotAppController"].ID.Bytes()[:4]) {
		return computeAVSRegistrar.UnpackNotAppControllerError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeAVSRegistrar.abi.Errors["OperatorAlreadyInAllowlist"].ID.Bytes()[:4]) {
		return computeAVSRegistrar.UnpackOperatorAlreadyInAllowlistError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeAVSRegistrar.abi.Errors["OperatorNotInAllowlist"].ID.Bytes()[:4]) {
		return computeAVSRegistrar.UnpackOperatorNotInAllowlistError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeAVSRegistrar.abi.Errors["OperatorSetIdNotAssigned"].ID.Bytes()[:4]) {
		return computeAVSRegistrar.UnpackOperatorSetIdNotAssignedError(raw[4:])
	}
	if bytes.Equal(raw[:4], computeAVSRegistrar.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return computeAVSRegistrar.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ComputeAVSRegistrarInvalidShortString represents a InvalidShortString error raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ComputeAVSRegistrarInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackInvalidShortStringError(raw []byte) (*ComputeAVSRegistrarInvalidShortString, error) {
	out := new(ComputeAVSRegistrarInvalidShortString)
	if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeAVSRegistrarKeyNotRegistered represents a KeyNotRegistered error raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarKeyNotRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyNotRegistered()
func ComputeAVSRegistrarKeyNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x815589fb98e9c9e6451b30fae4d2e9b5e50d764ec5d0f94d6bce0bef87a8bc29")
}

// UnpackKeyNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyNotRegistered()
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackKeyNotRegisteredError(raw []byte) (*ComputeAVSRegistrarKeyNotRegistered, error) {
	out := new(ComputeAVSRegistrarKeyNotRegistered)
	if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, "KeyNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeAVSRegistrarNotAllocationManager represents a NotAllocationManager error raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarNotAllocationManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAllocationManager()
func ComputeAVSRegistrarNotAllocationManagerErrorID() common.Hash {
	return common.HexToHash("0xd797b1542aede3f6a97f85f80e31001b9023c8dcf6675c741cb029c3c79161b3")
}

// UnpackNotAllocationManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAllocationManager()
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackNotAllocationManagerError(raw []byte) (*ComputeAVSRegistrarNotAllocationManager, error) {
	out := new(ComputeAVSRegistrarNotAllocationManager)
	if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, "NotAllocationManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeAVSRegistrarNotAppController represents a NotAppController error raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarNotAppController struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAppController()
func ComputeAVSRegistrarNotAppControllerErrorID() common.Hash {
	return common.HexToHash("0x626e239362f9371458de20b06d3fd86b8e3c92ca75e136ef3540651430775634")
}

// UnpackNotAppControllerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAppController()
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackNotAppControllerError(raw []byte) (*ComputeAVSRegistrarNotAppController, error) {
	out := new(ComputeAVSRegistrarNotAppController)
	if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, "NotAppController", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeAVSRegistrarOperatorAlreadyInAllowlist represents a OperatorAlreadyInAllowlist error raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorAlreadyInAllowlist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperatorAlreadyInAllowlist()
func ComputeAVSRegistrarOperatorAlreadyInAllowlistErrorID() common.Hash {
	return common.HexToHash("0x86e0613f5e82ffbbd658f9ef285ec07634c8e5ca925e8cf2202b2e25b5690ef9")
}

// UnpackOperatorAlreadyInAllowlistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperatorAlreadyInAllowlist()
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOperatorAlreadyInAllowlistError(raw []byte) (*ComputeAVSRegistrarOperatorAlreadyInAllowlist, error) {
	out := new(ComputeAVSRegistrarOperatorAlreadyInAllowlist)
	if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, "OperatorAlreadyInAllowlist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeAVSRegistrarOperatorNotInAllowlist represents a OperatorNotInAllowlist error raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorNotInAllowlist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperatorNotInAllowlist()
func ComputeAVSRegistrarOperatorNotInAllowlistErrorID() common.Hash {
	return common.HexToHash("0x11134b84f4e655b7c11dadf49d587722b99b924aad00de453583a0c9adb7f6b5")
}

// UnpackOperatorNotInAllowlistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperatorNotInAllowlist()
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOperatorNotInAllowlistError(raw []byte) (*ComputeAVSRegistrarOperatorNotInAllowlist, error) {
	out := new(ComputeAVSRegistrarOperatorNotInAllowlist)
	if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, "OperatorNotInAllowlist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeAVSRegistrarOperatorSetIdNotAssigned represents a OperatorSetIdNotAssigned error raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarOperatorSetIdNotAssigned struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperatorSetIdNotAssigned()
func ComputeAVSRegistrarOperatorSetIdNotAssignedErrorID() common.Hash {
	return common.HexToHash("0x1c226c363506937c5c3bde8290259603a608404b25860cf987d075334bfe69af")
}

// UnpackOperatorSetIdNotAssignedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperatorSetIdNotAssigned()
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackOperatorSetIdNotAssignedError(raw []byte) (*ComputeAVSRegistrarOperatorSetIdNotAssigned, error) {
	out := new(ComputeAVSRegistrarOperatorSetIdNotAssigned)
	if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, "OperatorSetIdNotAssigned", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeAVSRegistrarStringTooLong represents a StringTooLong error raised by the ComputeAVSRegistrar contract.
type ComputeAVSRegistrarStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ComputeAVSRegistrarStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (computeAVSRegistrar *ComputeAVSRegistrar) UnpackStringTooLongError(raw []byte) (*ComputeAVSRegistrarStringTooLong, error) {
	out := new(ComputeAVSRegistrarStringTooLong)
	if err := computeAVSRegistrar.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}
