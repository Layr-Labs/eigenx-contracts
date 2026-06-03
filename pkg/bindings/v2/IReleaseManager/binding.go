// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IReleaseManager

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

// IReleaseManagerTypesArtifact is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerTypesArtifact struct {
	Digest   [32]byte
	Registry string
}

// IReleaseManagerTypesRelease is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerTypesRelease struct {
	Artifacts     []IReleaseManagerTypesArtifact
	UpgradeByTime uint32
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IReleaseManagerMetaData contains all meta data concerning the IReleaseManager contract.
var IReleaseManagerMetaData = bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUpgradeByTime\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalReleases\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MetadataURIPublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUpgradeByTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustPublishMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoReleases\",\"inputs\":[]}]",
	ID:  "IReleaseManager",
}

// IReleaseManager is an auto generated Go binding around an Ethereum contract.
type IReleaseManager struct {
	abi abi.ABI
}

// NewIReleaseManager creates a new instance of IReleaseManager.
func NewIReleaseManager() *IReleaseManager {
	parsed, err := IReleaseManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IReleaseManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IReleaseManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetLatestRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd30eeb88.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (iReleaseManager *IReleaseManager) PackGetLatestRelease(operatorSet OperatorSet) []byte {
	enc, err := iReleaseManager.abi.Pack("getLatestRelease", operatorSet)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLatestRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd30eeb88.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (iReleaseManager *IReleaseManager) TryPackGetLatestRelease(operatorSet OperatorSet) ([]byte, error) {
	return iReleaseManager.abi.Pack("getLatestRelease", operatorSet)
}

// GetLatestReleaseOutput serves as a container for the return parameters of contract
// method GetLatestRelease.
type GetLatestReleaseOutput struct {
	Arg0 *big.Int
	Arg1 IReleaseManagerTypesRelease
}

// UnpackGetLatestRelease is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (iReleaseManager *IReleaseManager) UnpackGetLatestRelease(data []byte) (GetLatestReleaseOutput, error) {
	out, err := iReleaseManager.abi.Unpack("getLatestRelease", data)
	outstruct := new(GetLatestReleaseOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = *abi.ConvertType(out[1], new(IReleaseManagerTypesRelease)).(*IReleaseManagerTypesRelease)
	return *outstruct, nil
}

// PackGetLatestUpgradeByTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9e0ed68.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (iReleaseManager *IReleaseManager) PackGetLatestUpgradeByTime(operatorSet OperatorSet) []byte {
	enc, err := iReleaseManager.abi.Pack("getLatestUpgradeByTime", operatorSet)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLatestUpgradeByTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9e0ed68.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (iReleaseManager *IReleaseManager) TryPackGetLatestUpgradeByTime(operatorSet OperatorSet) ([]byte, error) {
	return iReleaseManager.abi.Pack("getLatestUpgradeByTime", operatorSet)
}

// UnpackGetLatestUpgradeByTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (iReleaseManager *IReleaseManager) UnpackGetLatestUpgradeByTime(data []byte) (uint32, error) {
	out, err := iReleaseManager.abi.Unpack("getLatestUpgradeByTime", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetMetadataURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb053b56d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (iReleaseManager *IReleaseManager) PackGetMetadataURI(operatorSet OperatorSet) []byte {
	enc, err := iReleaseManager.abi.Pack("getMetadataURI", operatorSet)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMetadataURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb053b56d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (iReleaseManager *IReleaseManager) TryPackGetMetadataURI(operatorSet OperatorSet) ([]byte, error) {
	return iReleaseManager.abi.Pack("getMetadataURI", operatorSet)
}

// UnpackGetMetadataURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (iReleaseManager *IReleaseManager) UnpackGetMetadataURI(data []byte) (string, error) {
	out, err := iReleaseManager.abi.Unpack("getMetadataURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGetRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3acab5fc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (iReleaseManager *IReleaseManager) PackGetRelease(operatorSet OperatorSet, releaseId *big.Int) []byte {
	enc, err := iReleaseManager.abi.Pack("getRelease", operatorSet, releaseId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3acab5fc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (iReleaseManager *IReleaseManager) TryPackGetRelease(operatorSet OperatorSet, releaseId *big.Int) ([]byte, error) {
	return iReleaseManager.abi.Pack("getRelease", operatorSet, releaseId)
}

// UnpackGetRelease is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (iReleaseManager *IReleaseManager) UnpackGetRelease(data []byte) (IReleaseManagerTypesRelease, error) {
	out, err := iReleaseManager.abi.Unpack("getRelease", data)
	if err != nil {
		return *new(IReleaseManagerTypesRelease), err
	}
	out0 := *abi.ConvertType(out[0], new(IReleaseManagerTypesRelease)).(*IReleaseManagerTypesRelease)
	return out0, nil
}

// PackGetTotalReleases is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66f409f7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (iReleaseManager *IReleaseManager) PackGetTotalReleases(operatorSet OperatorSet) []byte {
	enc, err := iReleaseManager.abi.Pack("getTotalReleases", operatorSet)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTotalReleases is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66f409f7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (iReleaseManager *IReleaseManager) TryPackGetTotalReleases(operatorSet OperatorSet) ([]byte, error) {
	return iReleaseManager.abi.Pack("getTotalReleases", operatorSet)
}

// UnpackGetTotalReleases is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (iReleaseManager *IReleaseManager) UnpackGetTotalReleases(data []byte) (*big.Int, error) {
	out, err := iReleaseManager.abi.Unpack("getTotalReleases", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackIsValidRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x517e4068.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (iReleaseManager *IReleaseManager) PackIsValidRelease(operatorSet OperatorSet, releaseId *big.Int) []byte {
	enc, err := iReleaseManager.abi.Pack("isValidRelease", operatorSet, releaseId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsValidRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x517e4068.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (iReleaseManager *IReleaseManager) TryPackIsValidRelease(operatorSet OperatorSet, releaseId *big.Int) ([]byte, error) {
	return iReleaseManager.abi.Pack("isValidRelease", operatorSet, releaseId)
}

// UnpackIsValidRelease is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (iReleaseManager *IReleaseManager) UnpackIsValidRelease(data []byte) (bool, error) {
	out, err := iReleaseManager.abi.Unpack("isValidRelease", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPublishMetadataURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4840a67c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (iReleaseManager *IReleaseManager) PackPublishMetadataURI(operatorSet OperatorSet, metadataURI string) []byte {
	enc, err := iReleaseManager.abi.Pack("publishMetadataURI", operatorSet, metadataURI)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPublishMetadataURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4840a67c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (iReleaseManager *IReleaseManager) TryPackPublishMetadataURI(operatorSet OperatorSet, metadataURI string) ([]byte, error) {
	return iReleaseManager.abi.Pack("publishMetadataURI", operatorSet, metadataURI)
}

// PackPublishRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c09ea82.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (iReleaseManager *IReleaseManager) PackPublishRelease(operatorSet OperatorSet, release IReleaseManagerTypesRelease) []byte {
	enc, err := iReleaseManager.abi.Pack("publishRelease", operatorSet, release)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPublishRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7c09ea82.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (iReleaseManager *IReleaseManager) TryPackPublishRelease(operatorSet OperatorSet, release IReleaseManagerTypesRelease) ([]byte, error) {
	return iReleaseManager.abi.Pack("publishRelease", operatorSet, release)
}

// UnpackPublishRelease is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (iReleaseManager *IReleaseManager) UnpackPublishRelease(data []byte) (*big.Int, error) {
	out, err := iReleaseManager.abi.Unpack("publishRelease", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// IReleaseManagerMetadataURIPublished represents a MetadataURIPublished event raised by the IReleaseManager contract.
type IReleaseManagerMetadataURIPublished struct {
	OperatorSet OperatorSet
	MetadataURI string
	Raw         *types.Log // Blockchain specific contextual infos
}

const IReleaseManagerMetadataURIPublishedEventName = "MetadataURIPublished"

// ContractEventName returns the user-defined event name.
func (IReleaseManagerMetadataURIPublished) ContractEventName() string {
	return IReleaseManagerMetadataURIPublishedEventName
}

// UnpackMetadataURIPublishedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (iReleaseManager *IReleaseManager) UnpackMetadataURIPublishedEvent(log *types.Log) (*IReleaseManagerMetadataURIPublished, error) {
	event := "MetadataURIPublished"
	if len(log.Topics) == 0 || log.Topics[0] != iReleaseManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IReleaseManagerMetadataURIPublished)
	if len(log.Data) > 0 {
		if err := iReleaseManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iReleaseManager.abi.Events[event].Inputs {
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

// IReleaseManagerReleasePublished represents a ReleasePublished event raised by the IReleaseManager contract.
type IReleaseManagerReleasePublished struct {
	OperatorSet OperatorSet
	ReleaseId   *big.Int
	Release     IReleaseManagerTypesRelease
	Raw         *types.Log // Blockchain specific contextual infos
}

const IReleaseManagerReleasePublishedEventName = "ReleasePublished"

// ContractEventName returns the user-defined event name.
func (IReleaseManagerReleasePublished) ContractEventName() string {
	return IReleaseManagerReleasePublishedEventName
}

// UnpackReleasePublishedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (iReleaseManager *IReleaseManager) UnpackReleasePublishedEvent(log *types.Log) (*IReleaseManagerReleasePublished, error) {
	event := "ReleasePublished"
	if len(log.Topics) == 0 || log.Topics[0] != iReleaseManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IReleaseManagerReleasePublished)
	if len(log.Data) > 0 {
		if err := iReleaseManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iReleaseManager.abi.Events[event].Inputs {
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
func (iReleaseManager *IReleaseManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iReleaseManager.abi.Errors["InvalidMetadataURI"].ID.Bytes()[:4]) {
		return iReleaseManager.UnpackInvalidMetadataURIError(raw[4:])
	}
	if bytes.Equal(raw[:4], iReleaseManager.abi.Errors["InvalidUpgradeByTime"].ID.Bytes()[:4]) {
		return iReleaseManager.UnpackInvalidUpgradeByTimeError(raw[4:])
	}
	if bytes.Equal(raw[:4], iReleaseManager.abi.Errors["MustPublishMetadataURI"].ID.Bytes()[:4]) {
		return iReleaseManager.UnpackMustPublishMetadataURIError(raw[4:])
	}
	if bytes.Equal(raw[:4], iReleaseManager.abi.Errors["NoReleases"].ID.Bytes()[:4]) {
		return iReleaseManager.UnpackNoReleasesError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IReleaseManagerInvalidMetadataURI represents a InvalidMetadataURI error raised by the IReleaseManager contract.
type IReleaseManagerInvalidMetadataURI struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMetadataURI()
func IReleaseManagerInvalidMetadataURIErrorID() common.Hash {
	return common.HexToHash("0xeec403f0a426cb998978041335218af13d12b0aefa43d1377f21848e4146382e")
}

// UnpackInvalidMetadataURIError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMetadataURI()
func (iReleaseManager *IReleaseManager) UnpackInvalidMetadataURIError(raw []byte) (*IReleaseManagerInvalidMetadataURI, error) {
	out := new(IReleaseManagerInvalidMetadataURI)
	if err := iReleaseManager.abi.UnpackIntoInterface(out, "InvalidMetadataURI", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IReleaseManagerInvalidUpgradeByTime represents a InvalidUpgradeByTime error raised by the IReleaseManager contract.
type IReleaseManagerInvalidUpgradeByTime struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidUpgradeByTime()
func IReleaseManagerInvalidUpgradeByTimeErrorID() common.Hash {
	return common.HexToHash("0x325ec75ff7a7da5bd1beb24ba891a4bd24420ccdba286388fd030f817e7a3600")
}

// UnpackInvalidUpgradeByTimeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidUpgradeByTime()
func (iReleaseManager *IReleaseManager) UnpackInvalidUpgradeByTimeError(raw []byte) (*IReleaseManagerInvalidUpgradeByTime, error) {
	out := new(IReleaseManagerInvalidUpgradeByTime)
	if err := iReleaseManager.abi.UnpackIntoInterface(out, "InvalidUpgradeByTime", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IReleaseManagerMustPublishMetadataURI represents a MustPublishMetadataURI error raised by the IReleaseManager contract.
type IReleaseManagerMustPublishMetadataURI struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MustPublishMetadataURI()
func IReleaseManagerMustPublishMetadataURIErrorID() common.Hash {
	return common.HexToHash("0x827cdcaee63f67112b340e74841c3d365632c66e6ab2b33ce151cc589663d848")
}

// UnpackMustPublishMetadataURIError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MustPublishMetadataURI()
func (iReleaseManager *IReleaseManager) UnpackMustPublishMetadataURIError(raw []byte) (*IReleaseManagerMustPublishMetadataURI, error) {
	out := new(IReleaseManagerMustPublishMetadataURI)
	if err := iReleaseManager.abi.UnpackIntoInterface(out, "MustPublishMetadataURI", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IReleaseManagerNoReleases represents a NoReleases error raised by the IReleaseManager contract.
type IReleaseManagerNoReleases struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoReleases()
func IReleaseManagerNoReleasesErrorID() common.Hash {
	return common.HexToHash("0xf4634142e74f4b03a7665765c8f78c5fad09d0c98a67acc37107853e6824dafe")
}

// UnpackNoReleasesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoReleases()
func (iReleaseManager *IReleaseManager) UnpackNoReleasesError(raw []byte) (*IReleaseManagerNoReleases, error) {
	out := new(IReleaseManagerNoReleases)
	if err := iReleaseManager.abi.UnpackIntoInterface(out, "NoReleases", raw); err != nil {
		return nil, err
	}
	return out, nil
}
