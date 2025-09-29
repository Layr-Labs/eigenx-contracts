// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IPermissionController

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

// IPermissionControllerMetaData contains all meta data concerning the IPermissionController contract.
var IPermissionControllerMetaData = bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canCall\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointeePermissions\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAppointees\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPendingAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pendingAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminAdded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdminAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotHaveZeroAdmins\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]}]",
	ID:  "IPermissionController",
}

// IPermissionController is an auto generated Go binding around an Ethereum contract.
type IPermissionController struct {
	abi abi.ABI
}

// NewIPermissionController creates a new instance of IPermissionController.
func NewIPermissionController() *IPermissionController {
	parsed, err := IPermissionControllerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IPermissionController{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IPermissionController) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAcceptAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x628806ef.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptAdmin(address account) returns()
func (iPermissionController *IPermissionController) PackAcceptAdmin(account common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("acceptAdmin", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x628806ef.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function acceptAdmin(address account) returns()
func (iPermissionController *IPermissionController) TryPackAcceptAdmin(account common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("acceptAdmin", account)
}

// PackAddPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb5a4e87.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (iPermissionController *IPermissionController) PackAddPendingAdmin(account common.Address, admin common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("addPendingAdmin", account, admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb5a4e87.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (iPermissionController *IPermissionController) TryPackAddPendingAdmin(account common.Address, admin common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("addPendingAdmin", account, admin)
}

// PackCanCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdf595cb8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) returns(bool)
func (iPermissionController *IPermissionController) PackCanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) []byte {
	enc, err := iPermissionController.abi.Pack("canCall", account, caller, target, selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCanCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdf595cb8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) returns(bool)
func (iPermissionController *IPermissionController) TryPackCanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) ([]byte, error) {
	return iPermissionController.abi.Pack("canCall", account, caller, target, selector)
}

// UnpackCanCall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) returns(bool)
func (iPermissionController *IPermissionController) UnpackCanCall(data []byte) (bool, error) {
	out, err := iPermissionController.abi.Unpack("canCall", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad5f2210.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (iPermissionController *IPermissionController) PackGetAdmins(account common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("getAdmins", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad5f2210.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (iPermissionController *IPermissionController) TryPackGetAdmins(account common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("getAdmins", account)
}

// UnpackGetAdmins is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (iPermissionController *IPermissionController) UnpackGetAdmins(data []byte) ([]common.Address, error) {
	out, err := iPermissionController.abi.Unpack("getAdmins", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetAppointeePermissions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x882a3b38.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAppointeePermissions(address account, address appointee) returns(address[], bytes4[])
func (iPermissionController *IPermissionController) PackGetAppointeePermissions(account common.Address, appointee common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("getAppointeePermissions", account, appointee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAppointeePermissions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x882a3b38.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAppointeePermissions(address account, address appointee) returns(address[], bytes4[])
func (iPermissionController *IPermissionController) TryPackGetAppointeePermissions(account common.Address, appointee common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("getAppointeePermissions", account, appointee)
}

// GetAppointeePermissionsOutput serves as a container for the return parameters of contract
// method GetAppointeePermissions.
type GetAppointeePermissionsOutput struct {
	Arg0 []common.Address
	Arg1 [][4]byte
}

// UnpackGetAppointeePermissions is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) returns(address[], bytes4[])
func (iPermissionController *IPermissionController) UnpackGetAppointeePermissions(data []byte) (GetAppointeePermissionsOutput, error) {
	out, err := iPermissionController.abi.Unpack("getAppointeePermissions", data)
	outstruct := new(GetAppointeePermissionsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([][4]byte)).(*[][4]byte)
	return *outstruct, nil
}

// PackGetAppointees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfddbdefd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) returns(address[])
func (iPermissionController *IPermissionController) PackGetAppointees(account common.Address, target common.Address, selector [4]byte) []byte {
	enc, err := iPermissionController.abi.Pack("getAppointees", account, target, selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAppointees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfddbdefd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) returns(address[])
func (iPermissionController *IPermissionController) TryPackGetAppointees(account common.Address, target common.Address, selector [4]byte) ([]byte, error) {
	return iPermissionController.abi.Pack("getAppointees", account, target, selector)
}

// UnpackGetAppointees is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) returns(address[])
func (iPermissionController *IPermissionController) UnpackGetAppointees(data []byte) ([]common.Address, error) {
	out, err := iPermissionController.abi.Unpack("getAppointees", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetPendingAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6bddfa1f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (iPermissionController *IPermissionController) PackGetPendingAdmins(account common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("getPendingAdmins", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPendingAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6bddfa1f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (iPermissionController *IPermissionController) TryPackGetPendingAdmins(account common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("getPendingAdmins", account)
}

// UnpackGetPendingAdmins is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (iPermissionController *IPermissionController) UnpackGetPendingAdmins(data []byte) ([]common.Address, error) {
	out, err := iPermissionController.abi.Unpack("getPendingAdmins", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackIsAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91006745.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (iPermissionController *IPermissionController) PackIsAdmin(account common.Address, caller common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("isAdmin", account, caller)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91006745.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (iPermissionController *IPermissionController) TryPackIsAdmin(account common.Address, caller common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("isAdmin", account, caller)
}

// UnpackIsAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (iPermissionController *IPermissionController) UnpackIsAdmin(data []byte) (bool, error) {
	out, err := iPermissionController.abi.Unpack("isAdmin", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad8aca77.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (iPermissionController *IPermissionController) PackIsPendingAdmin(account common.Address, pendingAdmin common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("isPendingAdmin", account, pendingAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad8aca77.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (iPermissionController *IPermissionController) TryPackIsPendingAdmin(account common.Address, pendingAdmin common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("isPendingAdmin", account, pendingAdmin)
}

// UnpackIsPendingAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (iPermissionController *IPermissionController) UnpackIsPendingAdmin(data []byte) (bool, error) {
	out, err := iPermissionController.abi.Unpack("isPendingAdmin", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x268959e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (iPermissionController *IPermissionController) PackRemoveAdmin(account common.Address, admin common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("removeAdmin", account, admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x268959e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (iPermissionController *IPermissionController) TryPackRemoveAdmin(account common.Address, admin common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("removeAdmin", account, admin)
}

// PackRemoveAppointee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06641201.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (iPermissionController *IPermissionController) PackRemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) []byte {
	enc, err := iPermissionController.abi.Pack("removeAppointee", account, appointee, target, selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveAppointee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06641201.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (iPermissionController *IPermissionController) TryPackRemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) ([]byte, error) {
	return iPermissionController.abi.Pack("removeAppointee", account, appointee, target, selector)
}

// PackRemovePendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f906cf9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (iPermissionController *IPermissionController) PackRemovePendingAdmin(account common.Address, admin common.Address) []byte {
	enc, err := iPermissionController.abi.Pack("removePendingAdmin", account, admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemovePendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f906cf9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (iPermissionController *IPermissionController) TryPackRemovePendingAdmin(account common.Address, admin common.Address) ([]byte, error) {
	return iPermissionController.abi.Pack("removePendingAdmin", account, admin)
}

// PackSetAppointee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x950d806e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (iPermissionController *IPermissionController) PackSetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) []byte {
	enc, err := iPermissionController.abi.Pack("setAppointee", account, appointee, target, selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAppointee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x950d806e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (iPermissionController *IPermissionController) TryPackSetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) ([]byte, error) {
	return iPermissionController.abi.Pack("setAppointee", account, appointee, target, selector)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(string)
func (iPermissionController *IPermissionController) PackVersion() []byte {
	enc, err := iPermissionController.abi.Pack("version")
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
func (iPermissionController *IPermissionController) TryPackVersion() ([]byte, error) {
	return iPermissionController.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (iPermissionController *IPermissionController) UnpackVersion(data []byte) (string, error) {
	out, err := iPermissionController.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// IPermissionControllerAdminRemoved represents a AdminRemoved event raised by the IPermissionController contract.
type IPermissionControllerAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const IPermissionControllerAdminRemovedEventName = "AdminRemoved"

// ContractEventName returns the user-defined event name.
func (IPermissionControllerAdminRemoved) ContractEventName() string {
	return IPermissionControllerAdminRemovedEventName
}

// UnpackAdminRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (iPermissionController *IPermissionController) UnpackAdminRemovedEvent(log *types.Log) (*IPermissionControllerAdminRemoved, error) {
	event := "AdminRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != iPermissionController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IPermissionControllerAdminRemoved)
	if len(log.Data) > 0 {
		if err := iPermissionController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iPermissionController.abi.Events[event].Inputs {
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

// IPermissionControllerAdminSet represents a AdminSet event raised by the IPermissionController contract.
type IPermissionControllerAdminSet struct {
	Account common.Address
	Admin   common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const IPermissionControllerAdminSetEventName = "AdminSet"

// ContractEventName returns the user-defined event name.
func (IPermissionControllerAdminSet) ContractEventName() string {
	return IPermissionControllerAdminSetEventName
}

// UnpackAdminSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (iPermissionController *IPermissionController) UnpackAdminSetEvent(log *types.Log) (*IPermissionControllerAdminSet, error) {
	event := "AdminSet"
	if len(log.Topics) == 0 || log.Topics[0] != iPermissionController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IPermissionControllerAdminSet)
	if len(log.Data) > 0 {
		if err := iPermissionController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iPermissionController.abi.Events[event].Inputs {
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

// IPermissionControllerAppointeeRemoved represents a AppointeeRemoved event raised by the IPermissionController contract.
type IPermissionControllerAppointeeRemoved struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IPermissionControllerAppointeeRemovedEventName = "AppointeeRemoved"

// ContractEventName returns the user-defined event name.
func (IPermissionControllerAppointeeRemoved) ContractEventName() string {
	return IPermissionControllerAppointeeRemovedEventName
}

// UnpackAppointeeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (iPermissionController *IPermissionController) UnpackAppointeeRemovedEvent(log *types.Log) (*IPermissionControllerAppointeeRemoved, error) {
	event := "AppointeeRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != iPermissionController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IPermissionControllerAppointeeRemoved)
	if len(log.Data) > 0 {
		if err := iPermissionController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iPermissionController.abi.Events[event].Inputs {
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

// IPermissionControllerAppointeeSet represents a AppointeeSet event raised by the IPermissionController contract.
type IPermissionControllerAppointeeSet struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IPermissionControllerAppointeeSetEventName = "AppointeeSet"

// ContractEventName returns the user-defined event name.
func (IPermissionControllerAppointeeSet) ContractEventName() string {
	return IPermissionControllerAppointeeSetEventName
}

// UnpackAppointeeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (iPermissionController *IPermissionController) UnpackAppointeeSetEvent(log *types.Log) (*IPermissionControllerAppointeeSet, error) {
	event := "AppointeeSet"
	if len(log.Topics) == 0 || log.Topics[0] != iPermissionController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IPermissionControllerAppointeeSet)
	if len(log.Data) > 0 {
		if err := iPermissionController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iPermissionController.abi.Events[event].Inputs {
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

// IPermissionControllerPendingAdminAdded represents a PendingAdminAdded event raised by the IPermissionController contract.
type IPermissionControllerPendingAdminAdded struct {
	Account common.Address
	Admin   common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const IPermissionControllerPendingAdminAddedEventName = "PendingAdminAdded"

// ContractEventName returns the user-defined event name.
func (IPermissionControllerPendingAdminAdded) ContractEventName() string {
	return IPermissionControllerPendingAdminAddedEventName
}

// UnpackPendingAdminAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (iPermissionController *IPermissionController) UnpackPendingAdminAddedEvent(log *types.Log) (*IPermissionControllerPendingAdminAdded, error) {
	event := "PendingAdminAdded"
	if len(log.Topics) == 0 || log.Topics[0] != iPermissionController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IPermissionControllerPendingAdminAdded)
	if len(log.Data) > 0 {
		if err := iPermissionController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iPermissionController.abi.Events[event].Inputs {
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

// IPermissionControllerPendingAdminRemoved represents a PendingAdminRemoved event raised by the IPermissionController contract.
type IPermissionControllerPendingAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const IPermissionControllerPendingAdminRemovedEventName = "PendingAdminRemoved"

// ContractEventName returns the user-defined event name.
func (IPermissionControllerPendingAdminRemoved) ContractEventName() string {
	return IPermissionControllerPendingAdminRemovedEventName
}

// UnpackPendingAdminRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (iPermissionController *IPermissionController) UnpackPendingAdminRemovedEvent(log *types.Log) (*IPermissionControllerPendingAdminRemoved, error) {
	event := "PendingAdminRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != iPermissionController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IPermissionControllerPendingAdminRemoved)
	if len(log.Data) > 0 {
		if err := iPermissionController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iPermissionController.abi.Events[event].Inputs {
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
func (iPermissionController *IPermissionController) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iPermissionController.abi.Errors["AdminAlreadyPending"].ID.Bytes()[:4]) {
		return iPermissionController.UnpackAdminAlreadyPendingError(raw[4:])
	}
	if bytes.Equal(raw[:4], iPermissionController.abi.Errors["AdminAlreadySet"].ID.Bytes()[:4]) {
		return iPermissionController.UnpackAdminAlreadySetError(raw[4:])
	}
	if bytes.Equal(raw[:4], iPermissionController.abi.Errors["AdminNotPending"].ID.Bytes()[:4]) {
		return iPermissionController.UnpackAdminNotPendingError(raw[4:])
	}
	if bytes.Equal(raw[:4], iPermissionController.abi.Errors["AdminNotSet"].ID.Bytes()[:4]) {
		return iPermissionController.UnpackAdminNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], iPermissionController.abi.Errors["AppointeeAlreadySet"].ID.Bytes()[:4]) {
		return iPermissionController.UnpackAppointeeAlreadySetError(raw[4:])
	}
	if bytes.Equal(raw[:4], iPermissionController.abi.Errors["AppointeeNotSet"].ID.Bytes()[:4]) {
		return iPermissionController.UnpackAppointeeNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], iPermissionController.abi.Errors["CannotHaveZeroAdmins"].ID.Bytes()[:4]) {
		return iPermissionController.UnpackCannotHaveZeroAdminsError(raw[4:])
	}
	if bytes.Equal(raw[:4], iPermissionController.abi.Errors["NotAdmin"].ID.Bytes()[:4]) {
		return iPermissionController.UnpackNotAdminError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IPermissionControllerAdminAlreadyPending represents a AdminAlreadyPending error raised by the IPermissionController contract.
type IPermissionControllerAdminAlreadyPending struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AdminAlreadyPending()
func IPermissionControllerAdminAlreadyPendingErrorID() common.Hash {
	return common.HexToHash("0x3357dbc6b31bbe597b579c6a53a828b3ee69167bc490b2a726a39b45457af13a")
}

// UnpackAdminAlreadyPendingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AdminAlreadyPending()
func (iPermissionController *IPermissionController) UnpackAdminAlreadyPendingError(raw []byte) (*IPermissionControllerAdminAlreadyPending, error) {
	out := new(IPermissionControllerAdminAlreadyPending)
	if err := iPermissionController.abi.UnpackIntoInterface(out, "AdminAlreadyPending", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IPermissionControllerAdminAlreadySet represents a AdminAlreadySet error raised by the IPermissionController contract.
type IPermissionControllerAdminAlreadySet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AdminAlreadySet()
func IPermissionControllerAdminAlreadySetErrorID() common.Hash {
	return common.HexToHash("0x980b072844ef81673ff8cd372ad952b9831c82ce17bed844e6bc4d65df50a980")
}

// UnpackAdminAlreadySetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AdminAlreadySet()
func (iPermissionController *IPermissionController) UnpackAdminAlreadySetError(raw []byte) (*IPermissionControllerAdminAlreadySet, error) {
	out := new(IPermissionControllerAdminAlreadySet)
	if err := iPermissionController.abi.UnpackIntoInterface(out, "AdminAlreadySet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IPermissionControllerAdminNotPending represents a AdminNotPending error raised by the IPermissionController contract.
type IPermissionControllerAdminNotPending struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AdminNotPending()
func IPermissionControllerAdminNotPendingErrorID() common.Hash {
	return common.HexToHash("0xbed8295f73beeda47547aba177a870bd5261d82560db8a05af7925270ade0e6b")
}

// UnpackAdminNotPendingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AdminNotPending()
func (iPermissionController *IPermissionController) UnpackAdminNotPendingError(raw []byte) (*IPermissionControllerAdminNotPending, error) {
	out := new(IPermissionControllerAdminNotPending)
	if err := iPermissionController.abi.UnpackIntoInterface(out, "AdminNotPending", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IPermissionControllerAdminNotSet represents a AdminNotSet error raised by the IPermissionController contract.
type IPermissionControllerAdminNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AdminNotSet()
func IPermissionControllerAdminNotSetErrorID() common.Hash {
	return common.HexToHash("0xe2db0360c04c5107e6e3a721a69465af91c2ddfc4f8a90890486d09ecc156677")
}

// UnpackAdminNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AdminNotSet()
func (iPermissionController *IPermissionController) UnpackAdminNotSetError(raw []byte) (*IPermissionControllerAdminNotSet, error) {
	out := new(IPermissionControllerAdminNotSet)
	if err := iPermissionController.abi.UnpackIntoInterface(out, "AdminNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IPermissionControllerAppointeeAlreadySet represents a AppointeeAlreadySet error raised by the IPermissionController contract.
type IPermissionControllerAppointeeAlreadySet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AppointeeAlreadySet()
func IPermissionControllerAppointeeAlreadySetErrorID() common.Hash {
	return common.HexToHash("0xad8efeb78c08fed6ec420ffbb4031c124421d45151a37efeb96e06263e45e2d9")
}

// UnpackAppointeeAlreadySetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AppointeeAlreadySet()
func (iPermissionController *IPermissionController) UnpackAppointeeAlreadySetError(raw []byte) (*IPermissionControllerAppointeeAlreadySet, error) {
	out := new(IPermissionControllerAppointeeAlreadySet)
	if err := iPermissionController.abi.UnpackIntoInterface(out, "AppointeeAlreadySet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IPermissionControllerAppointeeNotSet represents a AppointeeNotSet error raised by the IPermissionController contract.
type IPermissionControllerAppointeeNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AppointeeNotSet()
func IPermissionControllerAppointeeNotSetErrorID() common.Hash {
	return common.HexToHash("0x262118cd24dcda43b20b66d1394c4a7d791ed5c85b8dbb23ee660e566376ead3")
}

// UnpackAppointeeNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AppointeeNotSet()
func (iPermissionController *IPermissionController) UnpackAppointeeNotSetError(raw []byte) (*IPermissionControllerAppointeeNotSet, error) {
	out := new(IPermissionControllerAppointeeNotSet)
	if err := iPermissionController.abi.UnpackIntoInterface(out, "AppointeeNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IPermissionControllerCannotHaveZeroAdmins represents a CannotHaveZeroAdmins error raised by the IPermissionController contract.
type IPermissionControllerCannotHaveZeroAdmins struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotHaveZeroAdmins()
func IPermissionControllerCannotHaveZeroAdminsErrorID() common.Hash {
	return common.HexToHash("0x867449581171d16992e1a2ca7c0a8fa3688f21698f6e40691a445d38f409dc1f")
}

// UnpackCannotHaveZeroAdminsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotHaveZeroAdmins()
func (iPermissionController *IPermissionController) UnpackCannotHaveZeroAdminsError(raw []byte) (*IPermissionControllerCannotHaveZeroAdmins, error) {
	out := new(IPermissionControllerCannotHaveZeroAdmins)
	if err := iPermissionController.abi.UnpackIntoInterface(out, "CannotHaveZeroAdmins", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IPermissionControllerNotAdmin represents a NotAdmin error raised by the IPermissionController contract.
type IPermissionControllerNotAdmin struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAdmin()
func IPermissionControllerNotAdminErrorID() common.Hash {
	return common.HexToHash("0x7bfa4b9fb0cd3687c1f539f384b3f3f258f2c9aa9186353d0815413b508ed97d")
}

// UnpackNotAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAdmin()
func (iPermissionController *IPermissionController) UnpackNotAdminError(raw []byte) (*IPermissionControllerNotAdmin, error) {
	out := new(IPermissionControllerNotAdmin)
	if err := iPermissionController.abi.UnpackIntoInterface(out, "NotAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}
