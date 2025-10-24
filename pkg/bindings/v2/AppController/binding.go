// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AppController

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

// IAppControllerAppConfig is an auto generated low-level Go binding around an user-defined struct.
type IAppControllerAppConfig struct {
	Creator                  common.Address
	OperatorSetId            uint32
	LatestReleaseBlockNumber uint32
	Status                   uint8
}

// IAppControllerRelease is an auto generated low-level Go binding around an user-defined struct.
type IAppControllerRelease struct {
	RmsRelease   IReleaseManagerTypesRelease
	PublicEnv    []byte
	EncryptedEnv []byte
}

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

// AppControllerMetaData contains all meta data concerning the AppController contract.
var AppControllerMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_releaseManager\",\"type\":\"address\",\"internalType\":\"contractIReleaseManager\"},{\"name\":\"_computeAVSRegistrar\",\"type\":\"address\",\"internalType\":\"contractIComputeAVSRegistrar\"},{\"name\":\"_computeOperator\",\"type\":\"address\",\"internalType\":\"contractIComputeOperator\"},{\"name\":\"_appBeacon\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"API_PERMISSION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateApiPermissionDigestHash\",\"inputs\":[{\"name\":\"permission\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateAppId\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeAVSRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIComputeAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIComputeOperator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createApp\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIAppController.Release\",\"components\":[{\"name\":\"rmsRelease\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"publicEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encryptedEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveAppCount\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppCreator\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppLatestReleaseBlockNumber\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppOperatorSetId\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppStatus\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.AppStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApps\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"},{\"name\":\"appConfigsMem\",\"type\":\"tuple[]\",\"internalType\":\"structIAppController.AppConfig[]\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"latestReleaseBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.AppStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppsByCreator\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"},{\"name\":\"appConfigsMem\",\"type\":\"tuple[]\",\"internalType\":\"structIAppController.AppConfig[]\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"latestReleaseBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.AppStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppsByDeveloper\",\"inputs\":[{\"name\":\"developer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"},{\"name\":\"appConfigsMem\",\"type\":\"tuple[]\",\"internalType\":\"structIAppController.AppConfig[]\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"latestReleaseBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIAppController.AppStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxActiveAppsPerUser\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalActiveAppCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxGlobalActiveApps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIReleaseManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMaxActiveAppsPerUser\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGlobalActiveApps\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"suspend\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"apps\",\"type\":\"address[]\",\"internalType\":\"contractIApp[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"terminateApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"terminateAppByAdmin\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAppMetadataURI\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeApp\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"internalType\":\"contractIApp\"},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIAppController.Release\",\"components\":[{\"name\":\"rmsRelease\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"publicEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encryptedEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AppCreated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppMetadataURIUpdated\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppStarted\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppStopped\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppSuspended\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppTerminated\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppTerminatedByAdmin\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppUpgraded\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIApp\"},{\"name\":\"rmsReleaseId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAppController.Release\",\"components\":[{\"name\":\"rmsRelease\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"publicEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"encryptedEnv\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalMaxActiveAppsSet\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxActiveAppsSet\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccountHasActiveApps\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GlobalMaxActiveAppsExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAppStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReleaseMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxActiveAppsExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MoreThanOneArtifact\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	ID:  "AppController",
}

// AppController is an auto generated Go binding around an Ethereum contract.
type AppController struct {
	abi abi.ABI
}

// NewAppController creates a new instance of AppController.
func NewAppController() *AppController {
	parsed, err := AppControllerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AppController{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AppController) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(string _version, address _permissionController, address _releaseManager, address _computeAVSRegistrar, address _computeOperator, address _appBeacon) returns()
func (appController *AppController) PackConstructor(_version string, _permissionController common.Address, _releaseManager common.Address, _computeAVSRegistrar common.Address, _computeOperator common.Address, _appBeacon common.Address) []byte {
	enc, err := appController.abi.Pack("", _version, _permissionController, _releaseManager, _computeAVSRegistrar, _computeOperator, _appBeacon)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAPIPERMISSIONTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a7f922c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function API_PERMISSION_TYPEHASH() view returns(bytes32)
func (appController *AppController) PackAPIPERMISSIONTYPEHASH() []byte {
	enc, err := appController.abi.Pack("API_PERMISSION_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAPIPERMISSIONTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a7f922c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function API_PERMISSION_TYPEHASH() view returns(bytes32)
func (appController *AppController) TryPackAPIPERMISSIONTYPEHASH() ([]byte, error) {
	return appController.abi.Pack("API_PERMISSION_TYPEHASH")
}

// UnpackAPIPERMISSIONTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a7f922c.
//
// Solidity: function API_PERMISSION_TYPEHASH() view returns(bytes32)
func (appController *AppController) UnpackAPIPERMISSIONTYPEHASH(data []byte) ([32]byte, error) {
	out, err := appController.abi.Unpack("API_PERMISSION_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackAppBeacon is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a52d0b5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function appBeacon() view returns(address)
func (appController *AppController) PackAppBeacon() []byte {
	enc, err := appController.abi.Pack("appBeacon")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAppBeacon is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a52d0b5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function appBeacon() view returns(address)
func (appController *AppController) TryPackAppBeacon() ([]byte, error) {
	return appController.abi.Pack("appBeacon")
}

// UnpackAppBeacon is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a52d0b5.
//
// Solidity: function appBeacon() view returns(address)
func (appController *AppController) UnpackAppBeacon(data []byte) (common.Address, error) {
	out, err := appController.abi.Unpack("appBeacon", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCalculateApiPermissionDigestHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9690b7e7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) view returns(bytes32)
func (appController *AppController) PackCalculateApiPermissionDigestHash(permission [4]byte, expiry *big.Int) []byte {
	enc, err := appController.abi.Pack("calculateApiPermissionDigestHash", permission, expiry)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateApiPermissionDigestHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9690b7e7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) view returns(bytes32)
func (appController *AppController) TryPackCalculateApiPermissionDigestHash(permission [4]byte, expiry *big.Int) ([]byte, error) {
	return appController.abi.Pack("calculateApiPermissionDigestHash", permission, expiry)
}

// UnpackCalculateApiPermissionDigestHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9690b7e7.
//
// Solidity: function calculateApiPermissionDigestHash(bytes4 permission, uint256 expiry) view returns(bytes32)
func (appController *AppController) UnpackCalculateApiPermissionDigestHash(data []byte) ([32]byte, error) {
	out, err := appController.abi.Unpack("calculateApiPermissionDigestHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackCalculateAppId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36b18874.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateAppId(address deployer, bytes32 salt) view returns(address)
func (appController *AppController) PackCalculateAppId(deployer common.Address, salt [32]byte) []byte {
	enc, err := appController.abi.Pack("calculateAppId", deployer, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateAppId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36b18874.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateAppId(address deployer, bytes32 salt) view returns(address)
func (appController *AppController) TryPackCalculateAppId(deployer common.Address, salt [32]byte) ([]byte, error) {
	return appController.abi.Pack("calculateAppId", deployer, salt)
}

// UnpackCalculateAppId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x36b18874.
//
// Solidity: function calculateAppId(address deployer, bytes32 salt) view returns(address)
func (appController *AppController) UnpackCalculateAppId(data []byte) (common.Address, error) {
	out, err := appController.abi.Unpack("calculateAppId", data)
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
func (appController *AppController) PackComputeAVSRegistrar() []byte {
	enc, err := appController.abi.Pack("computeAVSRegistrar")
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
func (appController *AppController) TryPackComputeAVSRegistrar() ([]byte, error) {
	return appController.abi.Pack("computeAVSRegistrar")
}

// UnpackComputeAVSRegistrar is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (appController *AppController) UnpackComputeAVSRegistrar(data []byte) (common.Address, error) {
	out, err := appController.abi.Unpack("computeAVSRegistrar", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackComputeOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3fc4cbfd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function computeOperator() view returns(address)
func (appController *AppController) PackComputeOperator() []byte {
	enc, err := appController.abi.Pack("computeOperator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackComputeOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3fc4cbfd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function computeOperator() view returns(address)
func (appController *AppController) TryPackComputeOperator() ([]byte, error) {
	return appController.abi.Pack("computeOperator")
}

// UnpackComputeOperator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3fc4cbfd.
//
// Solidity: function computeOperator() view returns(address)
func (appController *AppController) UnpackComputeOperator(data []byte) (common.Address, error) {
	out, err := appController.abi.Unpack("computeOperator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCreateApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa60daa8f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createApp(bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (appController *AppController) PackCreateApp(salt [32]byte, release IAppControllerRelease) []byte {
	enc, err := appController.abi.Pack("createApp", salt, release)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa60daa8f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createApp(bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (appController *AppController) TryPackCreateApp(salt [32]byte, release IAppControllerRelease) ([]byte, error) {
	return appController.abi.Pack("createApp", salt, release)
}

// UnpackCreateApp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa60daa8f.
//
// Solidity: function createApp(bytes32 salt, (((bytes32,string)[],uint32),bytes,bytes) release) returns(address app)
func (appController *AppController) UnpackCreateApp(data []byte) (common.Address, error) {
	out, err := appController.abi.Unpack("createApp", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDomainSeparator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf698da25.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (appController *AppController) PackDomainSeparator() []byte {
	enc, err := appController.abi.Pack("domainSeparator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDomainSeparator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf698da25.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (appController *AppController) TryPackDomainSeparator() ([]byte, error) {
	return appController.abi.Pack("domainSeparator")
}

// UnpackDomainSeparator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (appController *AppController) UnpackDomainSeparator(data []byte) ([32]byte, error) {
	out, err := appController.abi.Unpack("domainSeparator", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetActiveAppCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c2199fb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getActiveAppCount(address user) view returns(uint32)
func (appController *AppController) PackGetActiveAppCount(user common.Address) []byte {
	enc, err := appController.abi.Pack("getActiveAppCount", user)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetActiveAppCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c2199fb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getActiveAppCount(address user) view returns(uint32)
func (appController *AppController) TryPackGetActiveAppCount(user common.Address) ([]byte, error) {
	return appController.abi.Pack("getActiveAppCount", user)
}

// UnpackGetActiveAppCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0c2199fb.
//
// Solidity: function getActiveAppCount(address user) view returns(uint32)
func (appController *AppController) UnpackGetActiveAppCount(data []byte) (uint32, error) {
	out, err := appController.abi.Unpack("getActiveAppCount", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetAppCreator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67962d48.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAppCreator(address app) view returns(address)
func (appController *AppController) PackGetAppCreator(app common.Address) []byte {
	enc, err := appController.abi.Pack("getAppCreator", app)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAppCreator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67962d48.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAppCreator(address app) view returns(address)
func (appController *AppController) TryPackGetAppCreator(app common.Address) ([]byte, error) {
	return appController.abi.Pack("getAppCreator", app)
}

// UnpackGetAppCreator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x67962d48.
//
// Solidity: function getAppCreator(address app) view returns(address)
func (appController *AppController) UnpackGetAppCreator(data []byte) (common.Address, error) {
	out, err := appController.abi.Unpack("getAppCreator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAppLatestReleaseBlockNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ffbdce6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAppLatestReleaseBlockNumber(address app) view returns(uint32)
func (appController *AppController) PackGetAppLatestReleaseBlockNumber(app common.Address) []byte {
	enc, err := appController.abi.Pack("getAppLatestReleaseBlockNumber", app)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAppLatestReleaseBlockNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ffbdce6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAppLatestReleaseBlockNumber(address app) view returns(uint32)
func (appController *AppController) TryPackGetAppLatestReleaseBlockNumber(app common.Address) ([]byte, error) {
	return appController.abi.Pack("getAppLatestReleaseBlockNumber", app)
}

// UnpackGetAppLatestReleaseBlockNumber is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ffbdce6.
//
// Solidity: function getAppLatestReleaseBlockNumber(address app) view returns(uint32)
func (appController *AppController) UnpackGetAppLatestReleaseBlockNumber(data []byte) (uint32, error) {
	out, err := appController.abi.Unpack("getAppLatestReleaseBlockNumber", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetAppOperatorSetId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6eb2099f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAppOperatorSetId(address app) view returns(uint32)
func (appController *AppController) PackGetAppOperatorSetId(app common.Address) []byte {
	enc, err := appController.abi.Pack("getAppOperatorSetId", app)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAppOperatorSetId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6eb2099f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAppOperatorSetId(address app) view returns(uint32)
func (appController *AppController) TryPackGetAppOperatorSetId(app common.Address) ([]byte, error) {
	return appController.abi.Pack("getAppOperatorSetId", app)
}

// UnpackGetAppOperatorSetId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6eb2099f.
//
// Solidity: function getAppOperatorSetId(address app) view returns(uint32)
func (appController *AppController) UnpackGetAppOperatorSetId(data []byte) (uint32, error) {
	out, err := appController.abi.Unpack("getAppOperatorSetId", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetAppStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5aae178.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAppStatus(address app) view returns(uint8)
func (appController *AppController) PackGetAppStatus(app common.Address) []byte {
	enc, err := appController.abi.Pack("getAppStatus", app)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAppStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5aae178.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAppStatus(address app) view returns(uint8)
func (appController *AppController) TryPackGetAppStatus(app common.Address) ([]byte, error) {
	return appController.abi.Pack("getAppStatus", app)
}

// UnpackGetAppStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd5aae178.
//
// Solidity: function getAppStatus(address app) view returns(uint8)
func (appController *AppController) UnpackGetAppStatus(data []byte) (uint8, error) {
	out, err := appController.abi.Unpack("getAppStatus", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackGetApps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa37e7e44.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getApps(uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) PackGetApps(offset *big.Int, limit *big.Int) []byte {
	enc, err := appController.abi.Pack("getApps", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetApps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa37e7e44.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getApps(uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) TryPackGetApps(offset *big.Int, limit *big.Int) ([]byte, error) {
	return appController.abi.Pack("getApps", offset, limit)
}

// GetAppsOutput serves as a container for the return parameters of contract
// method GetApps.
type GetAppsOutput struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}

// UnpackGetApps is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa37e7e44.
//
// Solidity: function getApps(uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) UnpackGetApps(data []byte) (GetAppsOutput, error) {
	out, err := appController.abi.Unpack("getApps", data)
	outstruct := new(GetAppsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Apps = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AppConfigsMem = *abi.ConvertType(out[1], new([]IAppControllerAppConfig)).(*[]IAppControllerAppConfig)
	return *outstruct, nil
}

// PackGetAppsByCreator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8099ef2e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAppsByCreator(address creator, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) PackGetAppsByCreator(creator common.Address, offset *big.Int, limit *big.Int) []byte {
	enc, err := appController.abi.Pack("getAppsByCreator", creator, offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAppsByCreator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8099ef2e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAppsByCreator(address creator, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) TryPackGetAppsByCreator(creator common.Address, offset *big.Int, limit *big.Int) ([]byte, error) {
	return appController.abi.Pack("getAppsByCreator", creator, offset, limit)
}

// GetAppsByCreatorOutput serves as a container for the return parameters of contract
// method GetAppsByCreator.
type GetAppsByCreatorOutput struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}

// UnpackGetAppsByCreator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8099ef2e.
//
// Solidity: function getAppsByCreator(address creator, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) UnpackGetAppsByCreator(data []byte) (GetAppsByCreatorOutput, error) {
	out, err := appController.abi.Unpack("getAppsByCreator", data)
	outstruct := new(GetAppsByCreatorOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Apps = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AppConfigsMem = *abi.ConvertType(out[1], new([]IAppControllerAppConfig)).(*[]IAppControllerAppConfig)
	return *outstruct, nil
}

// PackGetAppsByDeveloper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf36618ac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAppsByDeveloper(address developer, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) PackGetAppsByDeveloper(developer common.Address, offset *big.Int, limit *big.Int) []byte {
	enc, err := appController.abi.Pack("getAppsByDeveloper", developer, offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAppsByDeveloper is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf36618ac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAppsByDeveloper(address developer, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) TryPackGetAppsByDeveloper(developer common.Address, offset *big.Int, limit *big.Int) ([]byte, error) {
	return appController.abi.Pack("getAppsByDeveloper", developer, offset, limit)
}

// GetAppsByDeveloperOutput serves as a container for the return parameters of contract
// method GetAppsByDeveloper.
type GetAppsByDeveloperOutput struct {
	Apps          []common.Address
	AppConfigsMem []IAppControllerAppConfig
}

// UnpackGetAppsByDeveloper is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf36618ac.
//
// Solidity: function getAppsByDeveloper(address developer, uint256 offset, uint256 limit) view returns(address[] apps, (address,uint32,uint32,uint8)[] appConfigsMem)
func (appController *AppController) UnpackGetAppsByDeveloper(data []byte) (GetAppsByDeveloperOutput, error) {
	out, err := appController.abi.Unpack("getAppsByDeveloper", data)
	outstruct := new(GetAppsByDeveloperOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Apps = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AppConfigsMem = *abi.ConvertType(out[1], new([]IAppControllerAppConfig)).(*[]IAppControllerAppConfig)
	return *outstruct, nil
}

// PackGetMaxActiveAppsPerUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40f70a9e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMaxActiveAppsPerUser(address user) view returns(uint32)
func (appController *AppController) PackGetMaxActiveAppsPerUser(user common.Address) []byte {
	enc, err := appController.abi.Pack("getMaxActiveAppsPerUser", user)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMaxActiveAppsPerUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40f70a9e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMaxActiveAppsPerUser(address user) view returns(uint32)
func (appController *AppController) TryPackGetMaxActiveAppsPerUser(user common.Address) ([]byte, error) {
	return appController.abi.Pack("getMaxActiveAppsPerUser", user)
}

// UnpackGetMaxActiveAppsPerUser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40f70a9e.
//
// Solidity: function getMaxActiveAppsPerUser(address user) view returns(uint32)
func (appController *AppController) UnpackGetMaxActiveAppsPerUser(data []byte) (uint32, error) {
	out, err := appController.abi.Unpack("getMaxActiveAppsPerUser", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGlobalActiveAppCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8aa2bd3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function globalActiveAppCount() view returns(uint32)
func (appController *AppController) PackGlobalActiveAppCount() []byte {
	enc, err := appController.abi.Pack("globalActiveAppCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGlobalActiveAppCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8aa2bd3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function globalActiveAppCount() view returns(uint32)
func (appController *AppController) TryPackGlobalActiveAppCount() ([]byte, error) {
	return appController.abi.Pack("globalActiveAppCount")
}

// UnpackGlobalActiveAppCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa8aa2bd3.
//
// Solidity: function globalActiveAppCount() view returns(uint32)
func (appController *AppController) UnpackGlobalActiveAppCount(data []byte) (uint32, error) {
	out, err := appController.abi.Unpack("globalActiveAppCount", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address admin) returns()
func (appController *AppController) PackInitialize(admin common.Address) []byte {
	enc, err := appController.abi.Pack("initialize", admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address admin) returns()
func (appController *AppController) TryPackInitialize(admin common.Address) ([]byte, error) {
	return appController.abi.Pack("initialize", admin)
}

// PackMaxGlobalActiveApps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa46530a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function maxGlobalActiveApps() view returns(uint32)
func (appController *AppController) PackMaxGlobalActiveApps() []byte {
	enc, err := appController.abi.Pack("maxGlobalActiveApps")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMaxGlobalActiveApps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa46530a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function maxGlobalActiveApps() view returns(uint32)
func (appController *AppController) TryPackMaxGlobalActiveApps() ([]byte, error) {
	return appController.abi.Pack("maxGlobalActiveApps")
}

// UnpackMaxGlobalActiveApps is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa46530a2.
//
// Solidity: function maxGlobalActiveApps() view returns(uint32)
func (appController *AppController) UnpackMaxGlobalActiveApps(data []byte) (uint32, error) {
	out, err := appController.abi.Unpack("maxGlobalActiveApps", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackPermissionController is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4657e26a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function permissionController() view returns(address)
func (appController *AppController) PackPermissionController() []byte {
	enc, err := appController.abi.Pack("permissionController")
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
func (appController *AppController) TryPackPermissionController() ([]byte, error) {
	return appController.abi.Pack("permissionController")
}

// UnpackPermissionController is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (appController *AppController) UnpackPermissionController(data []byte) (common.Address, error) {
	out, err := appController.abi.Unpack("permissionController", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackReleaseManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d78573e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function releaseManager() view returns(address)
func (appController *AppController) PackReleaseManager() []byte {
	enc, err := appController.abi.Pack("releaseManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReleaseManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d78573e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function releaseManager() view returns(address)
func (appController *AppController) TryPackReleaseManager() ([]byte, error) {
	return appController.abi.Pack("releaseManager")
}

// UnpackReleaseManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0d78573e.
//
// Solidity: function releaseManager() view returns(address)
func (appController *AppController) UnpackReleaseManager(data []byte) (common.Address, error) {
	out, err := appController.abi.Unpack("releaseManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetMaxActiveAppsPerUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd49fec2b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMaxActiveAppsPerUser(address user, uint32 limit) returns()
func (appController *AppController) PackSetMaxActiveAppsPerUser(user common.Address, limit uint32) []byte {
	enc, err := appController.abi.Pack("setMaxActiveAppsPerUser", user, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMaxActiveAppsPerUser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd49fec2b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMaxActiveAppsPerUser(address user, uint32 limit) returns()
func (appController *AppController) TryPackSetMaxActiveAppsPerUser(user common.Address, limit uint32) ([]byte, error) {
	return appController.abi.Pack("setMaxActiveAppsPerUser", user, limit)
}

// PackSetMaxGlobalActiveApps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb438c141.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMaxGlobalActiveApps(uint32 limit) returns()
func (appController *AppController) PackSetMaxGlobalActiveApps(limit uint32) []byte {
	enc, err := appController.abi.Pack("setMaxGlobalActiveApps", limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMaxGlobalActiveApps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb438c141.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMaxGlobalActiveApps(uint32 limit) returns()
func (appController *AppController) TryPackSetMaxGlobalActiveApps(limit uint32) ([]byte, error) {
	return appController.abi.Pack("setMaxGlobalActiveApps", limit)
}

// PackStartApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4dd8ef81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function startApp(address app) returns()
func (appController *AppController) PackStartApp(app common.Address) []byte {
	enc, err := appController.abi.Pack("startApp", app)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStartApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4dd8ef81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function startApp(address app) returns()
func (appController *AppController) TryPackStartApp(app common.Address) ([]byte, error) {
	return appController.abi.Pack("startApp", app)
}

// PackStopApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe70b81b1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function stopApp(address app) returns()
func (appController *AppController) PackStopApp(app common.Address) []byte {
	enc, err := appController.abi.Pack("stopApp", app)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStopApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe70b81b1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function stopApp(address app) returns()
func (appController *AppController) TryPackStopApp(app common.Address) ([]byte, error) {
	return appController.abi.Pack("stopApp", app)
}

// PackSuspend is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb1e6ff7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function suspend(address account, address[] apps) returns()
func (appController *AppController) PackSuspend(account common.Address, apps []common.Address) []byte {
	enc, err := appController.abi.Pack("suspend", account, apps)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSuspend is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb1e6ff7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function suspend(address account, address[] apps) returns()
func (appController *AppController) TryPackSuspend(account common.Address, apps []common.Address) ([]byte, error) {
	return appController.abi.Pack("suspend", account, apps)
}

// PackTerminateApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x90e2196b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function terminateApp(address app) returns()
func (appController *AppController) PackTerminateApp(app common.Address) []byte {
	enc, err := appController.abi.Pack("terminateApp", app)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTerminateApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x90e2196b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function terminateApp(address app) returns()
func (appController *AppController) TryPackTerminateApp(app common.Address) ([]byte, error) {
	return appController.abi.Pack("terminateApp", app)
}

// PackTerminateAppByAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd66f7771.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function terminateAppByAdmin(address app) returns()
func (appController *AppController) PackTerminateAppByAdmin(app common.Address) []byte {
	enc, err := appController.abi.Pack("terminateAppByAdmin", app)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTerminateAppByAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd66f7771.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function terminateAppByAdmin(address app) returns()
func (appController *AppController) TryPackTerminateAppByAdmin(app common.Address) ([]byte, error) {
	return appController.abi.Pack("terminateAppByAdmin", app)
}

// PackUpdateAppMetadataURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65aa9a65.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateAppMetadataURI(address app, string metadataURI) returns()
func (appController *AppController) PackUpdateAppMetadataURI(app common.Address, metadataURI string) []byte {
	enc, err := appController.abi.Pack("updateAppMetadataURI", app, metadataURI)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateAppMetadataURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65aa9a65.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateAppMetadataURI(address app, string metadataURI) returns()
func (appController *AppController) TryPackUpdateAppMetadataURI(app common.Address, metadataURI string) ([]byte, error) {
	return appController.abi.Pack("updateAppMetadataURI", app, metadataURI)
}

// PackUpgradeApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd80a956b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeApp(address app, (((bytes32,string)[],uint32),bytes,bytes) release) returns(uint256)
func (appController *AppController) PackUpgradeApp(app common.Address, release IAppControllerRelease) []byte {
	enc, err := appController.abi.Pack("upgradeApp", app, release)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpgradeApp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd80a956b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function upgradeApp(address app, (((bytes32,string)[],uint32),bytes,bytes) release) returns(uint256)
func (appController *AppController) TryPackUpgradeApp(app common.Address, release IAppControllerRelease) ([]byte, error) {
	return appController.abi.Pack("upgradeApp", app, release)
}

// UnpackUpgradeApp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd80a956b.
//
// Solidity: function upgradeApp(address app, (((bytes32,string)[],uint32),bytes,bytes) release) returns(uint256)
func (appController *AppController) UnpackUpgradeApp(data []byte) (*big.Int, error) {
	out, err := appController.abi.Unpack("upgradeApp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(string)
func (appController *AppController) PackVersion() []byte {
	enc, err := appController.abi.Pack("version")
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
func (appController *AppController) TryPackVersion() ([]byte, error) {
	return appController.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (appController *AppController) UnpackVersion(data []byte) (string, error) {
	out, err := appController.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// AppControllerAppCreated represents a AppCreated event raised by the AppController contract.
type AppControllerAppCreated struct {
	Creator       common.Address
	App           common.Address
	OperatorSetId uint32
	Raw           *types.Log // Blockchain specific contextual infos
}

const AppControllerAppCreatedEventName = "AppCreated"

// ContractEventName returns the user-defined event name.
func (AppControllerAppCreated) ContractEventName() string {
	return AppControllerAppCreatedEventName
}

// UnpackAppCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppCreated(address indexed creator, address indexed app, uint32 operatorSetId)
func (appController *AppController) UnpackAppCreatedEvent(log *types.Log) (*AppControllerAppCreated, error) {
	event := "AppCreated"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerAppCreated)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerAppMetadataURIUpdated represents a AppMetadataURIUpdated event raised by the AppController contract.
type AppControllerAppMetadataURIUpdated struct {
	App         common.Address
	MetadataURI string
	Raw         *types.Log // Blockchain specific contextual infos
}

const AppControllerAppMetadataURIUpdatedEventName = "AppMetadataURIUpdated"

// ContractEventName returns the user-defined event name.
func (AppControllerAppMetadataURIUpdated) ContractEventName() string {
	return AppControllerAppMetadataURIUpdatedEventName
}

// UnpackAppMetadataURIUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppMetadataURIUpdated(address indexed app, string metadataURI)
func (appController *AppController) UnpackAppMetadataURIUpdatedEvent(log *types.Log) (*AppControllerAppMetadataURIUpdated, error) {
	event := "AppMetadataURIUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerAppMetadataURIUpdated)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerAppStarted represents a AppStarted event raised by the AppController contract.
type AppControllerAppStarted struct {
	App common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const AppControllerAppStartedEventName = "AppStarted"

// ContractEventName returns the user-defined event name.
func (AppControllerAppStarted) ContractEventName() string {
	return AppControllerAppStartedEventName
}

// UnpackAppStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppStarted(address indexed app)
func (appController *AppController) UnpackAppStartedEvent(log *types.Log) (*AppControllerAppStarted, error) {
	event := "AppStarted"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerAppStarted)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerAppStopped represents a AppStopped event raised by the AppController contract.
type AppControllerAppStopped struct {
	App common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const AppControllerAppStoppedEventName = "AppStopped"

// ContractEventName returns the user-defined event name.
func (AppControllerAppStopped) ContractEventName() string {
	return AppControllerAppStoppedEventName
}

// UnpackAppStoppedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppStopped(address indexed app)
func (appController *AppController) UnpackAppStoppedEvent(log *types.Log) (*AppControllerAppStopped, error) {
	event := "AppStopped"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerAppStopped)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerAppSuspended represents a AppSuspended event raised by the AppController contract.
type AppControllerAppSuspended struct {
	App common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const AppControllerAppSuspendedEventName = "AppSuspended"

// ContractEventName returns the user-defined event name.
func (AppControllerAppSuspended) ContractEventName() string {
	return AppControllerAppSuspendedEventName
}

// UnpackAppSuspendedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppSuspended(address indexed app)
func (appController *AppController) UnpackAppSuspendedEvent(log *types.Log) (*AppControllerAppSuspended, error) {
	event := "AppSuspended"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerAppSuspended)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerAppTerminated represents a AppTerminated event raised by the AppController contract.
type AppControllerAppTerminated struct {
	App common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const AppControllerAppTerminatedEventName = "AppTerminated"

// ContractEventName returns the user-defined event name.
func (AppControllerAppTerminated) ContractEventName() string {
	return AppControllerAppTerminatedEventName
}

// UnpackAppTerminatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppTerminated(address indexed app)
func (appController *AppController) UnpackAppTerminatedEvent(log *types.Log) (*AppControllerAppTerminated, error) {
	event := "AppTerminated"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerAppTerminated)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerAppTerminatedByAdmin represents a AppTerminatedByAdmin event raised by the AppController contract.
type AppControllerAppTerminatedByAdmin struct {
	App common.Address
	Raw *types.Log // Blockchain specific contextual infos
}

const AppControllerAppTerminatedByAdminEventName = "AppTerminatedByAdmin"

// ContractEventName returns the user-defined event name.
func (AppControllerAppTerminatedByAdmin) ContractEventName() string {
	return AppControllerAppTerminatedByAdminEventName
}

// UnpackAppTerminatedByAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppTerminatedByAdmin(address indexed app)
func (appController *AppController) UnpackAppTerminatedByAdminEvent(log *types.Log) (*AppControllerAppTerminatedByAdmin, error) {
	event := "AppTerminatedByAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerAppTerminatedByAdmin)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerAppUpgraded represents a AppUpgraded event raised by the AppController contract.
type AppControllerAppUpgraded struct {
	App          common.Address
	RmsReleaseId *big.Int
	Release      IAppControllerRelease
	Raw          *types.Log // Blockchain specific contextual infos
}

const AppControllerAppUpgradedEventName = "AppUpgraded"

// ContractEventName returns the user-defined event name.
func (AppControllerAppUpgraded) ContractEventName() string {
	return AppControllerAppUpgradedEventName
}

// UnpackAppUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AppUpgraded(address indexed app, uint256 rmsReleaseId, (((bytes32,string)[],uint32),bytes,bytes) release)
func (appController *AppController) UnpackAppUpgradedEvent(log *types.Log) (*AppControllerAppUpgraded, error) {
	event := "AppUpgraded"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerAppUpgraded)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerGlobalMaxActiveAppsSet represents a GlobalMaxActiveAppsSet event raised by the AppController contract.
type AppControllerGlobalMaxActiveAppsSet struct {
	Limit uint32
	Raw   *types.Log // Blockchain specific contextual infos
}

const AppControllerGlobalMaxActiveAppsSetEventName = "GlobalMaxActiveAppsSet"

// ContractEventName returns the user-defined event name.
func (AppControllerGlobalMaxActiveAppsSet) ContractEventName() string {
	return AppControllerGlobalMaxActiveAppsSetEventName
}

// UnpackGlobalMaxActiveAppsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GlobalMaxActiveAppsSet(uint32 limit)
func (appController *AppController) UnpackGlobalMaxActiveAppsSetEvent(log *types.Log) (*AppControllerGlobalMaxActiveAppsSet, error) {
	event := "GlobalMaxActiveAppsSet"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerGlobalMaxActiveAppsSet)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerInitialized represents a Initialized event raised by the AppController contract.
type AppControllerInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const AppControllerInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (AppControllerInitialized) ContractEventName() string {
	return AppControllerInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (appController *AppController) UnpackInitializedEvent(log *types.Log) (*AppControllerInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerInitialized)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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

// AppControllerMaxActiveAppsSet represents a MaxActiveAppsSet event raised by the AppController contract.
type AppControllerMaxActiveAppsSet struct {
	User  common.Address
	Limit uint32
	Raw   *types.Log // Blockchain specific contextual infos
}

const AppControllerMaxActiveAppsSetEventName = "MaxActiveAppsSet"

// ContractEventName returns the user-defined event name.
func (AppControllerMaxActiveAppsSet) ContractEventName() string {
	return AppControllerMaxActiveAppsSetEventName
}

// UnpackMaxActiveAppsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MaxActiveAppsSet(address indexed user, uint32 limit)
func (appController *AppController) UnpackMaxActiveAppsSetEvent(log *types.Log) (*AppControllerMaxActiveAppsSet, error) {
	event := "MaxActiveAppsSet"
	if len(log.Topics) == 0 || log.Topics[0] != appController.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AppControllerMaxActiveAppsSet)
	if len(log.Data) > 0 {
		if err := appController.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range appController.abi.Events[event].Inputs {
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
func (appController *AppController) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], appController.abi.Errors["AccountHasActiveApps"].ID.Bytes()[:4]) {
		return appController.UnpackAccountHasActiveAppsError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["AppAlreadyExists"].ID.Bytes()[:4]) {
		return appController.UnpackAppAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["AppDoesNotExist"].ID.Bytes()[:4]) {
		return appController.UnpackAppDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["GlobalMaxActiveAppsExceeded"].ID.Bytes()[:4]) {
		return appController.UnpackGlobalMaxActiveAppsExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["InvalidAppStatus"].ID.Bytes()[:4]) {
		return appController.UnpackInvalidAppStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["InvalidPermissions"].ID.Bytes()[:4]) {
		return appController.UnpackInvalidPermissionsError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["InvalidReleaseMetadataURI"].ID.Bytes()[:4]) {
		return appController.UnpackInvalidReleaseMetadataURIError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return appController.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["InvalidSignature"].ID.Bytes()[:4]) {
		return appController.UnpackInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["MaxActiveAppsExceeded"].ID.Bytes()[:4]) {
		return appController.UnpackMaxActiveAppsExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["MoreThanOneArtifact"].ID.Bytes()[:4]) {
		return appController.UnpackMoreThanOneArtifactError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["SignatureExpired"].ID.Bytes()[:4]) {
		return appController.UnpackSignatureExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], appController.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return appController.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// AppControllerAccountHasActiveApps represents a AccountHasActiveApps error raised by the AppController contract.
type AppControllerAccountHasActiveApps struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountHasActiveApps()
func AppControllerAccountHasActiveAppsErrorID() common.Hash {
	return common.HexToHash("0x0b9a8c798c922a29c1afc3a04fbb78a2b7721cbd0fc220aa2a1c7fc6aef42a23")
}

// UnpackAccountHasActiveAppsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountHasActiveApps()
func (appController *AppController) UnpackAccountHasActiveAppsError(raw []byte) (*AppControllerAccountHasActiveApps, error) {
	out := new(AppControllerAccountHasActiveApps)
	if err := appController.abi.UnpackIntoInterface(out, "AccountHasActiveApps", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerAppAlreadyExists represents a AppAlreadyExists error raised by the AppController contract.
type AppControllerAppAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AppAlreadyExists()
func AppControllerAppAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x6ccdc2fd81154bfcb848e706d73f7f2d24c1e992861d3f3a304ac05d0eb90563")
}

// UnpackAppAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AppAlreadyExists()
func (appController *AppController) UnpackAppAlreadyExistsError(raw []byte) (*AppControllerAppAlreadyExists, error) {
	out := new(AppControllerAppAlreadyExists)
	if err := appController.abi.UnpackIntoInterface(out, "AppAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerAppDoesNotExist represents a AppDoesNotExist error raised by the AppController contract.
type AppControllerAppDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AppDoesNotExist()
func AppControllerAppDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x54e0d7d224f5265326b69fdaeee66d5554717f6e421c012376d762458ec3145b")
}

// UnpackAppDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AppDoesNotExist()
func (appController *AppController) UnpackAppDoesNotExistError(raw []byte) (*AppControllerAppDoesNotExist, error) {
	out := new(AppControllerAppDoesNotExist)
	if err := appController.abi.UnpackIntoInterface(out, "AppDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerGlobalMaxActiveAppsExceeded represents a GlobalMaxActiveAppsExceeded error raised by the AppController contract.
type AppControllerGlobalMaxActiveAppsExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GlobalMaxActiveAppsExceeded()
func AppControllerGlobalMaxActiveAppsExceededErrorID() common.Hash {
	return common.HexToHash("0x42ca568f1a2d48ff9983e89b4d077f83c78d686b8400118b0286116476903411")
}

// UnpackGlobalMaxActiveAppsExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GlobalMaxActiveAppsExceeded()
func (appController *AppController) UnpackGlobalMaxActiveAppsExceededError(raw []byte) (*AppControllerGlobalMaxActiveAppsExceeded, error) {
	out := new(AppControllerGlobalMaxActiveAppsExceeded)
	if err := appController.abi.UnpackIntoInterface(out, "GlobalMaxActiveAppsExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerInvalidAppStatus represents a InvalidAppStatus error raised by the AppController contract.
type AppControllerInvalidAppStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAppStatus()
func AppControllerInvalidAppStatusErrorID() common.Hash {
	return common.HexToHash("0x609749e2bd5e8ccbdc8b98f99f091d349dfc73ed28105ec353346da311adf9a7")
}

// UnpackInvalidAppStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAppStatus()
func (appController *AppController) UnpackInvalidAppStatusError(raw []byte) (*AppControllerInvalidAppStatus, error) {
	out := new(AppControllerInvalidAppStatus)
	if err := appController.abi.UnpackIntoInterface(out, "InvalidAppStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerInvalidPermissions represents a InvalidPermissions error raised by the AppController contract.
type AppControllerInvalidPermissions struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPermissions()
func AppControllerInvalidPermissionsErrorID() common.Hash {
	return common.HexToHash("0x932d94f726428388537b641940dd88f9f37f70be827ee507792b87e4d26875f9")
}

// UnpackInvalidPermissionsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPermissions()
func (appController *AppController) UnpackInvalidPermissionsError(raw []byte) (*AppControllerInvalidPermissions, error) {
	out := new(AppControllerInvalidPermissions)
	if err := appController.abi.UnpackIntoInterface(out, "InvalidPermissions", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerInvalidReleaseMetadataURI represents a InvalidReleaseMetadataURI error raised by the AppController contract.
type AppControllerInvalidReleaseMetadataURI struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidReleaseMetadataURI()
func AppControllerInvalidReleaseMetadataURIErrorID() common.Hash {
	return common.HexToHash("0xc0c0e766cc08276c8450f40f5280d1528ad98813952a6db0dd95fa9d12f0e9bd")
}

// UnpackInvalidReleaseMetadataURIError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidReleaseMetadataURI()
func (appController *AppController) UnpackInvalidReleaseMetadataURIError(raw []byte) (*AppControllerInvalidReleaseMetadataURI, error) {
	out := new(AppControllerInvalidReleaseMetadataURI)
	if err := appController.abi.UnpackIntoInterface(out, "InvalidReleaseMetadataURI", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerInvalidShortString represents a InvalidShortString error raised by the AppController contract.
type AppControllerInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func AppControllerInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (appController *AppController) UnpackInvalidShortStringError(raw []byte) (*AppControllerInvalidShortString, error) {
	out := new(AppControllerInvalidShortString)
	if err := appController.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerInvalidSignature represents a InvalidSignature error raised by the AppController contract.
type AppControllerInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSignature()
func AppControllerInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0x8baa579fce362245063d36f11747a89dd489c54795634fc673cc0e0db51fedc5")
}

// UnpackInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSignature()
func (appController *AppController) UnpackInvalidSignatureError(raw []byte) (*AppControllerInvalidSignature, error) {
	out := new(AppControllerInvalidSignature)
	if err := appController.abi.UnpackIntoInterface(out, "InvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerMaxActiveAppsExceeded represents a MaxActiveAppsExceeded error raised by the AppController contract.
type AppControllerMaxActiveAppsExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MaxActiveAppsExceeded()
func AppControllerMaxActiveAppsExceededErrorID() common.Hash {
	return common.HexToHash("0x82d7ade6fc17104b977e5162fdfb656bbca3a73606c59522debb5d22bb5a2727")
}

// UnpackMaxActiveAppsExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MaxActiveAppsExceeded()
func (appController *AppController) UnpackMaxActiveAppsExceededError(raw []byte) (*AppControllerMaxActiveAppsExceeded, error) {
	out := new(AppControllerMaxActiveAppsExceeded)
	if err := appController.abi.UnpackIntoInterface(out, "MaxActiveAppsExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerMoreThanOneArtifact represents a MoreThanOneArtifact error raised by the AppController contract.
type AppControllerMoreThanOneArtifact struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MoreThanOneArtifact()
func AppControllerMoreThanOneArtifactErrorID() common.Hash {
	return common.HexToHash("0x98c82cefcadc45d2ae5a92b736a54ac6669c4ac1957e68744f1f5ac615cf62d9")
}

// UnpackMoreThanOneArtifactError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MoreThanOneArtifact()
func (appController *AppController) UnpackMoreThanOneArtifactError(raw []byte) (*AppControllerMoreThanOneArtifact, error) {
	out := new(AppControllerMoreThanOneArtifact)
	if err := appController.abi.UnpackIntoInterface(out, "MoreThanOneArtifact", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerSignatureExpired represents a SignatureExpired error raised by the AppController contract.
type AppControllerSignatureExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SignatureExpired()
func AppControllerSignatureExpiredErrorID() common.Hash {
	return common.HexToHash("0x0819bdcd8da3b255f8b8bf8497982cf672847d37ad445f8f2edca874c1fcb6a3")
}

// UnpackSignatureExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SignatureExpired()
func (appController *AppController) UnpackSignatureExpiredError(raw []byte) (*AppControllerSignatureExpired, error) {
	out := new(AppControllerSignatureExpired)
	if err := appController.abi.UnpackIntoInterface(out, "SignatureExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AppControllerStringTooLong represents a StringTooLong error raised by the AppController contract.
type AppControllerStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func AppControllerStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (appController *AppController) UnpackStringTooLongError(raw []byte) (*AppControllerStringTooLong, error) {
	out := new(AppControllerStringTooLong)
	if err := appController.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}
