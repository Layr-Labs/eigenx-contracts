// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ImageAllowlist

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

// IImageAllowlistImage is an auto generated low-level Go binding around an user-defined struct.
type IImageAllowlistImage struct {
	Pcrs        []IImageAllowlistPCR
	Version     string
	Description string
}

// IImageAllowlistPCR is an auto generated low-level Go binding around an user-defined struct.
type IImageAllowlistPCR struct {
	Index uint8
	Value [32]byte
}

// ImageAllowlistMetaData contains all meta data concerning the ImageAllowlist contract.
var ImageAllowlistMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addImage\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"image\",\"type\":\"tuple\",\"internalType\":\"structIImageAllowlist.Image\",\"components\":[{\"name\":\"pcrs\",\"type\":\"tuple[]\",\"internalType\":\"structIImageAllowlist.PCR[]\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addImages\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"images_\",\"type\":\"tuple[]\",\"internalType\":\"structIImageAllowlist.Image[]\",\"components\":[{\"name\":\"pcrs\",\"type\":\"tuple[]\",\"internalType\":\"structIImageAllowlist.PCR[]\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"keys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"images\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isImageAllowed\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"pcrs\",\"type\":\"tuple[]\",\"internalType\":\"structIImageAllowlist.PCR[]\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTCBValid\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"tcb\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumTCB\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeImage\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeImages\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"keys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumTCB\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"tcb\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ImageAdded\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageRemoved\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumTCBUpdated\",\"inputs\":[{\"name\":\"cvm\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIImageAllowlist.CVM\"},{\"name\":\"tcb\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyPCRs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ImageNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]}]",
	ID:  "ImageAllowlist",
}

// ImageAllowlist is an auto generated Go binding around an Ethereum contract.
type ImageAllowlist struct {
	abi abi.ABI
}

// NewImageAllowlist creates a new instance of ImageAllowlist.
func NewImageAllowlist() *ImageAllowlist {
	parsed, err := ImageAllowlistMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ImageAllowlist{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ImageAllowlist) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddImage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xffc2585d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addImage(uint8 cvm, ((uint8,bytes32)[],string,string) image) returns(bytes32)
func (imageAllowlist *ImageAllowlist) PackAddImage(cvm uint8, image IImageAllowlistImage) []byte {
	enc, err := imageAllowlist.abi.Pack("addImage", cvm, image)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddImage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xffc2585d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addImage(uint8 cvm, ((uint8,bytes32)[],string,string) image) returns(bytes32)
func (imageAllowlist *ImageAllowlist) TryPackAddImage(cvm uint8, image IImageAllowlistImage) ([]byte, error) {
	return imageAllowlist.abi.Pack("addImage", cvm, image)
}

// UnpackAddImage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xffc2585d.
//
// Solidity: function addImage(uint8 cvm, ((uint8,bytes32)[],string,string) image) returns(bytes32)
func (imageAllowlist *ImageAllowlist) UnpackAddImage(data []byte) ([32]byte, error) {
	out, err := imageAllowlist.abi.Unpack("addImage", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackAddImages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x188675b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addImages(uint8 cvm, ((uint8,bytes32)[],string,string)[] images_) returns(bytes32[] keys)
func (imageAllowlist *ImageAllowlist) PackAddImages(cvm uint8, images []IImageAllowlistImage) []byte {
	enc, err := imageAllowlist.abi.Pack("addImages", cvm, images)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddImages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x188675b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addImages(uint8 cvm, ((uint8,bytes32)[],string,string)[] images_) returns(bytes32[] keys)
func (imageAllowlist *ImageAllowlist) TryPackAddImages(cvm uint8, images []IImageAllowlistImage) ([]byte, error) {
	return imageAllowlist.abi.Pack("addImages", cvm, images)
}

// UnpackAddImages is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x188675b3.
//
// Solidity: function addImages(uint8 cvm, ((uint8,bytes32)[],string,string)[] images_) returns(bytes32[] keys)
func (imageAllowlist *ImageAllowlist) UnpackAddImages(data []byte) ([][32]byte, error) {
	out, err := imageAllowlist.abi.Unpack("addImages", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackImages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a4f02de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function images(uint8 , bytes32 ) view returns(bool)
func (imageAllowlist *ImageAllowlist) PackImages(arg0 uint8, arg1 [32]byte) []byte {
	enc, err := imageAllowlist.abi.Pack("images", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackImages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a4f02de.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function images(uint8 , bytes32 ) view returns(bool)
func (imageAllowlist *ImageAllowlist) TryPackImages(arg0 uint8, arg1 [32]byte) ([]byte, error) {
	return imageAllowlist.abi.Pack("images", arg0, arg1)
}

// UnpackImages is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4a4f02de.
//
// Solidity: function images(uint8 , bytes32 ) view returns(bool)
func (imageAllowlist *ImageAllowlist) UnpackImages(data []byte) (bool, error) {
	out, err := imageAllowlist.abi.Unpack("images", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address initialOwner) returns()
func (imageAllowlist *ImageAllowlist) PackInitialize(initialOwner common.Address) []byte {
	enc, err := imageAllowlist.abi.Pack("initialize", initialOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address initialOwner) returns()
func (imageAllowlist *ImageAllowlist) TryPackInitialize(initialOwner common.Address) ([]byte, error) {
	return imageAllowlist.abi.Pack("initialize", initialOwner)
}

// PackIsImageAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d5a1daf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isImageAllowed(uint8 cvm, (uint8,bytes32)[] pcrs) view returns(bool)
func (imageAllowlist *ImageAllowlist) PackIsImageAllowed(cvm uint8, pcrs []IImageAllowlistPCR) []byte {
	enc, err := imageAllowlist.abi.Pack("isImageAllowed", cvm, pcrs)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsImageAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d5a1daf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isImageAllowed(uint8 cvm, (uint8,bytes32)[] pcrs) view returns(bool)
func (imageAllowlist *ImageAllowlist) TryPackIsImageAllowed(cvm uint8, pcrs []IImageAllowlistPCR) ([]byte, error) {
	return imageAllowlist.abi.Pack("isImageAllowed", cvm, pcrs)
}

// UnpackIsImageAllowed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d5a1daf.
//
// Solidity: function isImageAllowed(uint8 cvm, (uint8,bytes32)[] pcrs) view returns(bool)
func (imageAllowlist *ImageAllowlist) UnpackIsImageAllowed(data []byte) (bool, error) {
	out, err := imageAllowlist.abi.Unpack("isImageAllowed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsTCBValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43e3dd05.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isTCBValid(uint8 cvm, uint64 tcb) view returns(bool)
func (imageAllowlist *ImageAllowlist) PackIsTCBValid(cvm uint8, tcb uint64) []byte {
	enc, err := imageAllowlist.abi.Pack("isTCBValid", cvm, tcb)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsTCBValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43e3dd05.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isTCBValid(uint8 cvm, uint64 tcb) view returns(bool)
func (imageAllowlist *ImageAllowlist) TryPackIsTCBValid(cvm uint8, tcb uint64) ([]byte, error) {
	return imageAllowlist.abi.Pack("isTCBValid", cvm, tcb)
}

// UnpackIsTCBValid is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x43e3dd05.
//
// Solidity: function isTCBValid(uint8 cvm, uint64 tcb) view returns(bool)
func (imageAllowlist *ImageAllowlist) UnpackIsTCBValid(data []byte) (bool, error) {
	out, err := imageAllowlist.abi.Unpack("isTCBValid", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMinimumTCB is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4ec8d4fa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minimumTCB(uint8 ) view returns(uint64)
func (imageAllowlist *ImageAllowlist) PackMinimumTCB(arg0 uint8) []byte {
	enc, err := imageAllowlist.abi.Pack("minimumTCB", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinimumTCB is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4ec8d4fa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minimumTCB(uint8 ) view returns(uint64)
func (imageAllowlist *ImageAllowlist) TryPackMinimumTCB(arg0 uint8) ([]byte, error) {
	return imageAllowlist.abi.Pack("minimumTCB", arg0)
}

// UnpackMinimumTCB is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4ec8d4fa.
//
// Solidity: function minimumTCB(uint8 ) view returns(uint64)
func (imageAllowlist *ImageAllowlist) UnpackMinimumTCB(data []byte) (uint64, error) {
	out, err := imageAllowlist.abi.Unpack("minimumTCB", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (imageAllowlist *ImageAllowlist) PackOwner() []byte {
	enc, err := imageAllowlist.abi.Pack("owner")
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
func (imageAllowlist *ImageAllowlist) TryPackOwner() ([]byte, error) {
	return imageAllowlist.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (imageAllowlist *ImageAllowlist) UnpackOwner(data []byte) (common.Address, error) {
	out, err := imageAllowlist.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRemoveImage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x036257e2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeImage(uint8 cvm, bytes32 key) returns()
func (imageAllowlist *ImageAllowlist) PackRemoveImage(cvm uint8, key [32]byte) []byte {
	enc, err := imageAllowlist.abi.Pack("removeImage", cvm, key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveImage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x036257e2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeImage(uint8 cvm, bytes32 key) returns()
func (imageAllowlist *ImageAllowlist) TryPackRemoveImage(cvm uint8, key [32]byte) ([]byte, error) {
	return imageAllowlist.abi.Pack("removeImage", cvm, key)
}

// PackRemoveImages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd79799b5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeImages(uint8 cvm, bytes32[] keys) returns()
func (imageAllowlist *ImageAllowlist) PackRemoveImages(cvm uint8, keys [][32]byte) []byte {
	enc, err := imageAllowlist.abi.Pack("removeImages", cvm, keys)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveImages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd79799b5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeImages(uint8 cvm, bytes32[] keys) returns()
func (imageAllowlist *ImageAllowlist) TryPackRemoveImages(cvm uint8, keys [][32]byte) ([]byte, error) {
	return imageAllowlist.abi.Pack("removeImages", cvm, keys)
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (imageAllowlist *ImageAllowlist) PackRenounceOwnership() []byte {
	enc, err := imageAllowlist.abi.Pack("renounceOwnership")
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
func (imageAllowlist *ImageAllowlist) TryPackRenounceOwnership() ([]byte, error) {
	return imageAllowlist.abi.Pack("renounceOwnership")
}

// PackSetMinimumTCB is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa54c8287.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinimumTCB(uint8 cvm, uint64 tcb) returns()
func (imageAllowlist *ImageAllowlist) PackSetMinimumTCB(cvm uint8, tcb uint64) []byte {
	enc, err := imageAllowlist.abi.Pack("setMinimumTCB", cvm, tcb)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinimumTCB is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa54c8287.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinimumTCB(uint8 cvm, uint64 tcb) returns()
func (imageAllowlist *ImageAllowlist) TryPackSetMinimumTCB(cvm uint8, tcb uint64) ([]byte, error) {
	return imageAllowlist.abi.Pack("setMinimumTCB", cvm, tcb)
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (imageAllowlist *ImageAllowlist) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := imageAllowlist.abi.Pack("transferOwnership", newOwner)
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
func (imageAllowlist *ImageAllowlist) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return imageAllowlist.abi.Pack("transferOwnership", newOwner)
}

// ImageAllowlistImageAdded represents a ImageAdded event raised by the ImageAllowlist contract.
type ImageAllowlistImageAdded struct {
	Cvm         uint8
	Key         [32]byte
	Version     string
	Description string
	Raw         *types.Log // Blockchain specific contextual infos
}

const ImageAllowlistImageAddedEventName = "ImageAdded"

// ContractEventName returns the user-defined event name.
func (ImageAllowlistImageAdded) ContractEventName() string {
	return ImageAllowlistImageAddedEventName
}

// UnpackImageAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ImageAdded(uint8 indexed cvm, bytes32 indexed key, string version, string description)
func (imageAllowlist *ImageAllowlist) UnpackImageAddedEvent(log *types.Log) (*ImageAllowlistImageAdded, error) {
	event := "ImageAdded"
	if len(log.Topics) == 0 || log.Topics[0] != imageAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ImageAllowlistImageAdded)
	if len(log.Data) > 0 {
		if err := imageAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range imageAllowlist.abi.Events[event].Inputs {
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

// ImageAllowlistImageRemoved represents a ImageRemoved event raised by the ImageAllowlist contract.
type ImageAllowlistImageRemoved struct {
	Cvm uint8
	Key [32]byte
	Raw *types.Log // Blockchain specific contextual infos
}

const ImageAllowlistImageRemovedEventName = "ImageRemoved"

// ContractEventName returns the user-defined event name.
func (ImageAllowlistImageRemoved) ContractEventName() string {
	return ImageAllowlistImageRemovedEventName
}

// UnpackImageRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ImageRemoved(uint8 indexed cvm, bytes32 indexed key)
func (imageAllowlist *ImageAllowlist) UnpackImageRemovedEvent(log *types.Log) (*ImageAllowlistImageRemoved, error) {
	event := "ImageRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != imageAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ImageAllowlistImageRemoved)
	if len(log.Data) > 0 {
		if err := imageAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range imageAllowlist.abi.Events[event].Inputs {
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

// ImageAllowlistInitialized represents a Initialized event raised by the ImageAllowlist contract.
type ImageAllowlistInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const ImageAllowlistInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ImageAllowlistInitialized) ContractEventName() string {
	return ImageAllowlistInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (imageAllowlist *ImageAllowlist) UnpackInitializedEvent(log *types.Log) (*ImageAllowlistInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != imageAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ImageAllowlistInitialized)
	if len(log.Data) > 0 {
		if err := imageAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range imageAllowlist.abi.Events[event].Inputs {
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

// ImageAllowlistMinimumTCBUpdated represents a MinimumTCBUpdated event raised by the ImageAllowlist contract.
type ImageAllowlistMinimumTCBUpdated struct {
	Cvm uint8
	Tcb uint64
	Raw *types.Log // Blockchain specific contextual infos
}

const ImageAllowlistMinimumTCBUpdatedEventName = "MinimumTCBUpdated"

// ContractEventName returns the user-defined event name.
func (ImageAllowlistMinimumTCBUpdated) ContractEventName() string {
	return ImageAllowlistMinimumTCBUpdatedEventName
}

// UnpackMinimumTCBUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinimumTCBUpdated(uint8 indexed cvm, uint64 tcb)
func (imageAllowlist *ImageAllowlist) UnpackMinimumTCBUpdatedEvent(log *types.Log) (*ImageAllowlistMinimumTCBUpdated, error) {
	event := "MinimumTCBUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != imageAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ImageAllowlistMinimumTCBUpdated)
	if len(log.Data) > 0 {
		if err := imageAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range imageAllowlist.abi.Events[event].Inputs {
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

// ImageAllowlistOwnershipTransferred represents a OwnershipTransferred event raised by the ImageAllowlist contract.
type ImageAllowlistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const ImageAllowlistOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ImageAllowlistOwnershipTransferred) ContractEventName() string {
	return ImageAllowlistOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (imageAllowlist *ImageAllowlist) UnpackOwnershipTransferredEvent(log *types.Log) (*ImageAllowlistOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != imageAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ImageAllowlistOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := imageAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range imageAllowlist.abi.Events[event].Inputs {
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
func (imageAllowlist *ImageAllowlist) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], imageAllowlist.abi.Errors["EmptyPCRs"].ID.Bytes()[:4]) {
		return imageAllowlist.UnpackEmptyPCRsError(raw[4:])
	}
	if bytes.Equal(raw[:4], imageAllowlist.abi.Errors["ImageNotFound"].ID.Bytes()[:4]) {
		return imageAllowlist.UnpackImageNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], imageAllowlist.abi.Errors["NotSorted"].ID.Bytes()[:4]) {
		return imageAllowlist.UnpackNotSortedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ImageAllowlistEmptyPCRs represents a EmptyPCRs error raised by the ImageAllowlist contract.
type ImageAllowlistEmptyPCRs struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyPCRs()
func ImageAllowlistEmptyPCRsErrorID() common.Hash {
	return common.HexToHash("0x6e1433191c7eef11cd1d10ffec63e7b15699dd22294e03d2fc8ce718f52db688")
}

// UnpackEmptyPCRsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyPCRs()
func (imageAllowlist *ImageAllowlist) UnpackEmptyPCRsError(raw []byte) (*ImageAllowlistEmptyPCRs, error) {
	out := new(ImageAllowlistEmptyPCRs)
	if err := imageAllowlist.abi.UnpackIntoInterface(out, "EmptyPCRs", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ImageAllowlistImageNotFound represents a ImageNotFound error raised by the ImageAllowlist contract.
type ImageAllowlistImageNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ImageNotFound()
func ImageAllowlistImageNotFoundErrorID() common.Hash {
	return common.HexToHash("0x264a7c138f2c67fcefd73820af910dccfa14df264cddda11b7049873eda2046d")
}

// UnpackImageNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ImageNotFound()
func (imageAllowlist *ImageAllowlist) UnpackImageNotFoundError(raw []byte) (*ImageAllowlistImageNotFound, error) {
	out := new(ImageAllowlistImageNotFound)
	if err := imageAllowlist.abi.UnpackIntoInterface(out, "ImageNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ImageAllowlistNotSorted represents a NotSorted error raised by the ImageAllowlist contract.
type ImageAllowlistNotSorted struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotSorted()
func ImageAllowlistNotSortedErrorID() common.Hash {
	return common.HexToHash("0xba50f911bcee6515ba6b7c6c51e77a66c0832d63ca24d4b5e29951fad43ddf1e")
}

// UnpackNotSortedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotSorted()
func (imageAllowlist *ImageAllowlist) UnpackNotSortedError(raw []byte) (*ImageAllowlistNotSorted, error) {
	out := new(ImageAllowlistNotSorted)
	if err := imageAllowlist.abi.UnpackIntoInterface(out, "NotSorted", raw); err != nil {
		return nil, err
	}
	return out, nil
}
