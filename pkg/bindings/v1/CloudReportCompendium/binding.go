// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CloudReportCompendium

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

// CloudReportCompendiumStorageConstructorParams is an auto generated low-level Go binding around an user-defined struct.
type CloudReportCompendiumStorageConstructorParams struct {
	DelegationManager   common.Address
	AllocationManager   common.Address
	CertVerifierRouter  common.Address
	ComputeAVSRegistrar common.Address
	ComputeOperator     common.Address
	ReportAttester      common.Address
	ReferenceProjectId  string
	StrategyToSlash     common.Address
}

// ICloudReportCompendiumCloudReport is an auto generated low-level Go binding around an user-defined struct.
type ICloudReportCompendiumCloudReport struct {
	ProjectId     string
	FromTimestamp uint32
	ToTimestamp   uint32
	EigendaCert   []byte
}

// ICloudReportCompendiumCloudReportSubmission is an auto generated low-level Go binding around an user-defined struct.
type ICloudReportCompendiumCloudReportSubmission struct {
	SubmissionTimestamp uint32
	ToTimestamp         uint32
	EigendaCertHash     [20]byte
}

// CloudReportCompendiumMetaData contains all meta data concerning the CloudReportCompendium contract.
var CloudReportCompendiumMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCloudReportCompendiumStorage.ConstructorParams\",\"components\":[{\"name\":\"delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"certVerifierRouter\",\"type\":\"address\",\"internalType\":\"contractIEigenDACertVerifierRouter\"},{\"name\":\"computeAVSRegistrar\",\"type\":\"address\",\"internalType\":\"contractIComputeAVSRegistrar\"},{\"name\":\"computeOperator\",\"type\":\"address\",\"internalType\":\"contractIComputeOperator\"},{\"name\":\"reportAttester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceProjectId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"strategyToSlash\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CLOUD_REPORT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASHING_OPERATORSET_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateReportDigestHash\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structICloudReportCompendium.CloudReport\",\"components\":[{\"name\":\"projectId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fromTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"toTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eigendaCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"certVerifierRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenDACertVerifierRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeAVSRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIComputeAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIComputeOperator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_reportFreshnessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_slashAmountTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minSlashInterval\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastSlashTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestReportSubmission\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICloudReportCompendium.CloudReportSubmission\",\"components\":[{\"name\":\"submissionTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"toTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eigendaCertHash\",\"type\":\"bytes20\",\"internalType\":\"bytes20\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minSlashInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"referenceProjectIdHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reportAttester\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reportFreshnessThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMinSlashInterval\",\"inputs\":[{\"name\":\"_minSlashInterval\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReportFreshnessThreshold\",\"inputs\":[{\"name\":\"_reportFreshnessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlashAmountTokens\",\"inputs\":[{\"name\":\"_slashAmountTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"eigendaCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashAmountTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyToSlash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitReport\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structICloudReportCompendium.CloudReport\",\"components\":[{\"name\":\"projectId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fromTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"toTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eigendaCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinSlashIntervalSet\",\"inputs\":[{\"name\":\"minSlashInterval\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportFreshnessThresholdSet\",\"inputs\":[{\"name\":\"reportFreshnessThreshold\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportSubmitted\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICloudReportCompendium.CloudReport\",\"components\":[{\"name\":\"projectId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fromTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"toTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eigendaCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashAmountTokensSet\",\"inputs\":[{\"name\":\"slashAmountTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientAllocation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCertificate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProjectId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReportTimestamps\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoSlashRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportChainTimestampMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportFromFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportTooStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlashTooSoon\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// CloudReportCompendiumABI is the input ABI used to generate the binding from.
// Deprecated: Use CloudReportCompendiumMetaData.ABI instead.
var CloudReportCompendiumABI = CloudReportCompendiumMetaData.ABI

// CloudReportCompendium is an auto generated Go binding around an Ethereum contract.
type CloudReportCompendium struct {
	CloudReportCompendiumCaller     // Read-only binding to the contract
	CloudReportCompendiumTransactor // Write-only binding to the contract
	CloudReportCompendiumFilterer   // Log filterer for contract events
}

// CloudReportCompendiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type CloudReportCompendiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloudReportCompendiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CloudReportCompendiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloudReportCompendiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CloudReportCompendiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloudReportCompendiumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CloudReportCompendiumSession struct {
	Contract     *CloudReportCompendium // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CloudReportCompendiumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CloudReportCompendiumCallerSession struct {
	Contract *CloudReportCompendiumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CloudReportCompendiumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CloudReportCompendiumTransactorSession struct {
	Contract     *CloudReportCompendiumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CloudReportCompendiumRaw is an auto generated low-level Go binding around an Ethereum contract.
type CloudReportCompendiumRaw struct {
	Contract *CloudReportCompendium // Generic contract binding to access the raw methods on
}

// CloudReportCompendiumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CloudReportCompendiumCallerRaw struct {
	Contract *CloudReportCompendiumCaller // Generic read-only contract binding to access the raw methods on
}

// CloudReportCompendiumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CloudReportCompendiumTransactorRaw struct {
	Contract *CloudReportCompendiumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCloudReportCompendium creates a new instance of CloudReportCompendium, bound to a specific deployed contract.
func NewCloudReportCompendium(address common.Address, backend bind.ContractBackend) (*CloudReportCompendium, error) {
	contract, err := bindCloudReportCompendium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendium{CloudReportCompendiumCaller: CloudReportCompendiumCaller{contract: contract}, CloudReportCompendiumTransactor: CloudReportCompendiumTransactor{contract: contract}, CloudReportCompendiumFilterer: CloudReportCompendiumFilterer{contract: contract}}, nil
}

// NewCloudReportCompendiumCaller creates a new read-only instance of CloudReportCompendium, bound to a specific deployed contract.
func NewCloudReportCompendiumCaller(address common.Address, caller bind.ContractCaller) (*CloudReportCompendiumCaller, error) {
	contract, err := bindCloudReportCompendium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumCaller{contract: contract}, nil
}

// NewCloudReportCompendiumTransactor creates a new write-only instance of CloudReportCompendium, bound to a specific deployed contract.
func NewCloudReportCompendiumTransactor(address common.Address, transactor bind.ContractTransactor) (*CloudReportCompendiumTransactor, error) {
	contract, err := bindCloudReportCompendium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumTransactor{contract: contract}, nil
}

// NewCloudReportCompendiumFilterer creates a new log filterer instance of CloudReportCompendium, bound to a specific deployed contract.
func NewCloudReportCompendiumFilterer(address common.Address, filterer bind.ContractFilterer) (*CloudReportCompendiumFilterer, error) {
	contract, err := bindCloudReportCompendium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumFilterer{contract: contract}, nil
}

// bindCloudReportCompendium binds a generic wrapper to an already deployed contract.
func bindCloudReportCompendium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CloudReportCompendiumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CloudReportCompendium *CloudReportCompendiumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CloudReportCompendium.Contract.CloudReportCompendiumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CloudReportCompendium *CloudReportCompendiumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.CloudReportCompendiumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CloudReportCompendium *CloudReportCompendiumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.CloudReportCompendiumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CloudReportCompendium *CloudReportCompendiumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CloudReportCompendium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CloudReportCompendium *CloudReportCompendiumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CloudReportCompendium *CloudReportCompendiumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.contract.Transact(opts, method, params...)
}

// CLOUDREPORTTYPEHASH is a free data retrieval call binding the contract method 0x593ea37d.
//
// Solidity: function CLOUD_REPORT_TYPEHASH() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumCaller) CLOUDREPORTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "CLOUD_REPORT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CLOUDREPORTTYPEHASH is a free data retrieval call binding the contract method 0x593ea37d.
//
// Solidity: function CLOUD_REPORT_TYPEHASH() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumSession) CLOUDREPORTTYPEHASH() ([32]byte, error) {
	return _CloudReportCompendium.Contract.CLOUDREPORTTYPEHASH(&_CloudReportCompendium.CallOpts)
}

// CLOUDREPORTTYPEHASH is a free data retrieval call binding the contract method 0x593ea37d.
//
// Solidity: function CLOUD_REPORT_TYPEHASH() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) CLOUDREPORTTYPEHASH() ([32]byte, error) {
	return _CloudReportCompendium.Contract.CLOUDREPORTTYPEHASH(&_CloudReportCompendium.CallOpts)
}

// SLASHINGOPERATORSETID is a free data retrieval call binding the contract method 0x96b313c8.
//
// Solidity: function SLASHING_OPERATORSET_ID() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumCaller) SLASHINGOPERATORSETID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "SLASHING_OPERATORSET_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SLASHINGOPERATORSETID is a free data retrieval call binding the contract method 0x96b313c8.
//
// Solidity: function SLASHING_OPERATORSET_ID() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumSession) SLASHINGOPERATORSETID() (uint32, error) {
	return _CloudReportCompendium.Contract.SLASHINGOPERATORSETID(&_CloudReportCompendium.CallOpts)
}

// SLASHINGOPERATORSETID is a free data retrieval call binding the contract method 0x96b313c8.
//
// Solidity: function SLASHING_OPERATORSET_ID() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) SLASHINGOPERATORSETID() (uint32, error) {
	return _CloudReportCompendium.Contract.SLASHINGOPERATORSETID(&_CloudReportCompendium.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumSession) AllocationManager() (common.Address, error) {
	return _CloudReportCompendium.Contract.AllocationManager(&_CloudReportCompendium.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) AllocationManager() (common.Address, error) {
	return _CloudReportCompendium.Contract.AllocationManager(&_CloudReportCompendium.CallOpts)
}

// CalculateReportDigestHash is a free data retrieval call binding the contract method 0xee1d5079.
//
// Solidity: function calculateReportDigestHash((string,uint32,uint32,bytes) report) view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumCaller) CalculateReportDigestHash(opts *bind.CallOpts, report ICloudReportCompendiumCloudReport) ([32]byte, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "calculateReportDigestHash", report)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateReportDigestHash is a free data retrieval call binding the contract method 0xee1d5079.
//
// Solidity: function calculateReportDigestHash((string,uint32,uint32,bytes) report) view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumSession) CalculateReportDigestHash(report ICloudReportCompendiumCloudReport) ([32]byte, error) {
	return _CloudReportCompendium.Contract.CalculateReportDigestHash(&_CloudReportCompendium.CallOpts, report)
}

// CalculateReportDigestHash is a free data retrieval call binding the contract method 0xee1d5079.
//
// Solidity: function calculateReportDigestHash((string,uint32,uint32,bytes) report) view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) CalculateReportDigestHash(report ICloudReportCompendiumCloudReport) ([32]byte, error) {
	return _CloudReportCompendium.Contract.CalculateReportDigestHash(&_CloudReportCompendium.CallOpts, report)
}

// CertVerifierRouter is a free data retrieval call binding the contract method 0x10d36975.
//
// Solidity: function certVerifierRouter() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCaller) CertVerifierRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "certVerifierRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CertVerifierRouter is a free data retrieval call binding the contract method 0x10d36975.
//
// Solidity: function certVerifierRouter() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumSession) CertVerifierRouter() (common.Address, error) {
	return _CloudReportCompendium.Contract.CertVerifierRouter(&_CloudReportCompendium.CallOpts)
}

// CertVerifierRouter is a free data retrieval call binding the contract method 0x10d36975.
//
// Solidity: function certVerifierRouter() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) CertVerifierRouter() (common.Address, error) {
	return _CloudReportCompendium.Contract.CertVerifierRouter(&_CloudReportCompendium.CallOpts)
}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCaller) ComputeAVSRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "computeAVSRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumSession) ComputeAVSRegistrar() (common.Address, error) {
	return _CloudReportCompendium.Contract.ComputeAVSRegistrar(&_CloudReportCompendium.CallOpts)
}

// ComputeAVSRegistrar is a free data retrieval call binding the contract method 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) ComputeAVSRegistrar() (common.Address, error) {
	return _CloudReportCompendium.Contract.ComputeAVSRegistrar(&_CloudReportCompendium.CallOpts)
}

// ComputeOperator is a free data retrieval call binding the contract method 0x3fc4cbfd.
//
// Solidity: function computeOperator() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCaller) ComputeOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "computeOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeOperator is a free data retrieval call binding the contract method 0x3fc4cbfd.
//
// Solidity: function computeOperator() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumSession) ComputeOperator() (common.Address, error) {
	return _CloudReportCompendium.Contract.ComputeOperator(&_CloudReportCompendium.CallOpts)
}

// ComputeOperator is a free data retrieval call binding the contract method 0x3fc4cbfd.
//
// Solidity: function computeOperator() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) ComputeOperator() (common.Address, error) {
	return _CloudReportCompendium.Contract.ComputeOperator(&_CloudReportCompendium.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumSession) DelegationManager() (common.Address, error) {
	return _CloudReportCompendium.Contract.DelegationManager(&_CloudReportCompendium.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) DelegationManager() (common.Address, error) {
	return _CloudReportCompendium.Contract.DelegationManager(&_CloudReportCompendium.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumSession) DomainSeparator() ([32]byte, error) {
	return _CloudReportCompendium.Contract.DomainSeparator(&_CloudReportCompendium.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) DomainSeparator() ([32]byte, error) {
	return _CloudReportCompendium.Contract.DomainSeparator(&_CloudReportCompendium.CallOpts)
}

// LastSlashTimestamp is a free data retrieval call binding the contract method 0xc77f5062.
//
// Solidity: function lastSlashTimestamp() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumCaller) LastSlashTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "lastSlashTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastSlashTimestamp is a free data retrieval call binding the contract method 0xc77f5062.
//
// Solidity: function lastSlashTimestamp() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumSession) LastSlashTimestamp() (uint32, error) {
	return _CloudReportCompendium.Contract.LastSlashTimestamp(&_CloudReportCompendium.CallOpts)
}

// LastSlashTimestamp is a free data retrieval call binding the contract method 0xc77f5062.
//
// Solidity: function lastSlashTimestamp() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) LastSlashTimestamp() (uint32, error) {
	return _CloudReportCompendium.Contract.LastSlashTimestamp(&_CloudReportCompendium.CallOpts)
}

// LatestReportSubmission is a free data retrieval call binding the contract method 0x371517a4.
//
// Solidity: function latestReportSubmission() view returns((uint32,uint32,bytes20))
func (_CloudReportCompendium *CloudReportCompendiumCaller) LatestReportSubmission(opts *bind.CallOpts) (ICloudReportCompendiumCloudReportSubmission, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "latestReportSubmission")

	if err != nil {
		return *new(ICloudReportCompendiumCloudReportSubmission), err
	}

	out0 := *abi.ConvertType(out[0], new(ICloudReportCompendiumCloudReportSubmission)).(*ICloudReportCompendiumCloudReportSubmission)

	return out0, err

}

// LatestReportSubmission is a free data retrieval call binding the contract method 0x371517a4.
//
// Solidity: function latestReportSubmission() view returns((uint32,uint32,bytes20))
func (_CloudReportCompendium *CloudReportCompendiumSession) LatestReportSubmission() (ICloudReportCompendiumCloudReportSubmission, error) {
	return _CloudReportCompendium.Contract.LatestReportSubmission(&_CloudReportCompendium.CallOpts)
}

// LatestReportSubmission is a free data retrieval call binding the contract method 0x371517a4.
//
// Solidity: function latestReportSubmission() view returns((uint32,uint32,bytes20))
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) LatestReportSubmission() (ICloudReportCompendiumCloudReportSubmission, error) {
	return _CloudReportCompendium.Contract.LatestReportSubmission(&_CloudReportCompendium.CallOpts)
}

// MinSlashInterval is a free data retrieval call binding the contract method 0x18e38c93.
//
// Solidity: function minSlashInterval() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumCaller) MinSlashInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "minSlashInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinSlashInterval is a free data retrieval call binding the contract method 0x18e38c93.
//
// Solidity: function minSlashInterval() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumSession) MinSlashInterval() (uint32, error) {
	return _CloudReportCompendium.Contract.MinSlashInterval(&_CloudReportCompendium.CallOpts)
}

// MinSlashInterval is a free data retrieval call binding the contract method 0x18e38c93.
//
// Solidity: function minSlashInterval() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) MinSlashInterval() (uint32, error) {
	return _CloudReportCompendium.Contract.MinSlashInterval(&_CloudReportCompendium.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumSession) PermissionController() (common.Address, error) {
	return _CloudReportCompendium.Contract.PermissionController(&_CloudReportCompendium.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) PermissionController() (common.Address, error) {
	return _CloudReportCompendium.Contract.PermissionController(&_CloudReportCompendium.CallOpts)
}

// ReferenceProjectIdHash is a free data retrieval call binding the contract method 0xf2e9ef5d.
//
// Solidity: function referenceProjectIdHash() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumCaller) ReferenceProjectIdHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "referenceProjectIdHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReferenceProjectIdHash is a free data retrieval call binding the contract method 0xf2e9ef5d.
//
// Solidity: function referenceProjectIdHash() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumSession) ReferenceProjectIdHash() ([32]byte, error) {
	return _CloudReportCompendium.Contract.ReferenceProjectIdHash(&_CloudReportCompendium.CallOpts)
}

// ReferenceProjectIdHash is a free data retrieval call binding the contract method 0xf2e9ef5d.
//
// Solidity: function referenceProjectIdHash() view returns(bytes32)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) ReferenceProjectIdHash() ([32]byte, error) {
	return _CloudReportCompendium.Contract.ReferenceProjectIdHash(&_CloudReportCompendium.CallOpts)
}

// ReportAttester is a free data retrieval call binding the contract method 0x285c6e4f.
//
// Solidity: function reportAttester() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCaller) ReportAttester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "reportAttester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReportAttester is a free data retrieval call binding the contract method 0x285c6e4f.
//
// Solidity: function reportAttester() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumSession) ReportAttester() (common.Address, error) {
	return _CloudReportCompendium.Contract.ReportAttester(&_CloudReportCompendium.CallOpts)
}

// ReportAttester is a free data retrieval call binding the contract method 0x285c6e4f.
//
// Solidity: function reportAttester() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) ReportAttester() (common.Address, error) {
	return _CloudReportCompendium.Contract.ReportAttester(&_CloudReportCompendium.CallOpts)
}

// ReportFreshnessThreshold is a free data retrieval call binding the contract method 0x5bc35ef6.
//
// Solidity: function reportFreshnessThreshold() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumCaller) ReportFreshnessThreshold(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "reportFreshnessThreshold")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ReportFreshnessThreshold is a free data retrieval call binding the contract method 0x5bc35ef6.
//
// Solidity: function reportFreshnessThreshold() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumSession) ReportFreshnessThreshold() (uint32, error) {
	return _CloudReportCompendium.Contract.ReportFreshnessThreshold(&_CloudReportCompendium.CallOpts)
}

// ReportFreshnessThreshold is a free data retrieval call binding the contract method 0x5bc35ef6.
//
// Solidity: function reportFreshnessThreshold() view returns(uint32)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) ReportFreshnessThreshold() (uint32, error) {
	return _CloudReportCompendium.Contract.ReportFreshnessThreshold(&_CloudReportCompendium.CallOpts)
}

// SlashAmountTokens is a free data retrieval call binding the contract method 0x7d3ceb34.
//
// Solidity: function slashAmountTokens() view returns(uint256)
func (_CloudReportCompendium *CloudReportCompendiumCaller) SlashAmountTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "slashAmountTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashAmountTokens is a free data retrieval call binding the contract method 0x7d3ceb34.
//
// Solidity: function slashAmountTokens() view returns(uint256)
func (_CloudReportCompendium *CloudReportCompendiumSession) SlashAmountTokens() (*big.Int, error) {
	return _CloudReportCompendium.Contract.SlashAmountTokens(&_CloudReportCompendium.CallOpts)
}

// SlashAmountTokens is a free data retrieval call binding the contract method 0x7d3ceb34.
//
// Solidity: function slashAmountTokens() view returns(uint256)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) SlashAmountTokens() (*big.Int, error) {
	return _CloudReportCompendium.Contract.SlashAmountTokens(&_CloudReportCompendium.CallOpts)
}

// StrategyToSlash is a free data retrieval call binding the contract method 0x5ac080f1.
//
// Solidity: function strategyToSlash() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCaller) StrategyToSlash(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "strategyToSlash")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyToSlash is a free data retrieval call binding the contract method 0x5ac080f1.
//
// Solidity: function strategyToSlash() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumSession) StrategyToSlash() (common.Address, error) {
	return _CloudReportCompendium.Contract.StrategyToSlash(&_CloudReportCompendium.CallOpts)
}

// StrategyToSlash is a free data retrieval call binding the contract method 0x5ac080f1.
//
// Solidity: function strategyToSlash() view returns(address)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) StrategyToSlash() (common.Address, error) {
	return _CloudReportCompendium.Contract.StrategyToSlash(&_CloudReportCompendium.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CloudReportCompendium *CloudReportCompendiumCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CloudReportCompendium.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CloudReportCompendium *CloudReportCompendiumSession) Version() (string, error) {
	return _CloudReportCompendium.Contract.Version(&_CloudReportCompendium.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CloudReportCompendium *CloudReportCompendiumCallerSession) Version() (string, error) {
	return _CloudReportCompendium.Contract.Version(&_CloudReportCompendium.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d302011.
//
// Solidity: function initialize(uint32 _reportFreshnessThreshold, uint256 _slashAmountTokens, uint32 _minSlashInterval) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactor) Initialize(opts *bind.TransactOpts, _reportFreshnessThreshold uint32, _slashAmountTokens *big.Int, _minSlashInterval uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.contract.Transact(opts, "initialize", _reportFreshnessThreshold, _slashAmountTokens, _minSlashInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d302011.
//
// Solidity: function initialize(uint32 _reportFreshnessThreshold, uint256 _slashAmountTokens, uint32 _minSlashInterval) returns()
func (_CloudReportCompendium *CloudReportCompendiumSession) Initialize(_reportFreshnessThreshold uint32, _slashAmountTokens *big.Int, _minSlashInterval uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.Initialize(&_CloudReportCompendium.TransactOpts, _reportFreshnessThreshold, _slashAmountTokens, _minSlashInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d302011.
//
// Solidity: function initialize(uint32 _reportFreshnessThreshold, uint256 _slashAmountTokens, uint32 _minSlashInterval) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactorSession) Initialize(_reportFreshnessThreshold uint32, _slashAmountTokens *big.Int, _minSlashInterval uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.Initialize(&_CloudReportCompendium.TransactOpts, _reportFreshnessThreshold, _slashAmountTokens, _minSlashInterval)
}

// SetMinSlashInterval is a paid mutator transaction binding the contract method 0xe7a6e55a.
//
// Solidity: function setMinSlashInterval(uint32 _minSlashInterval) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactor) SetMinSlashInterval(opts *bind.TransactOpts, _minSlashInterval uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.contract.Transact(opts, "setMinSlashInterval", _minSlashInterval)
}

// SetMinSlashInterval is a paid mutator transaction binding the contract method 0xe7a6e55a.
//
// Solidity: function setMinSlashInterval(uint32 _minSlashInterval) returns()
func (_CloudReportCompendium *CloudReportCompendiumSession) SetMinSlashInterval(_minSlashInterval uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.SetMinSlashInterval(&_CloudReportCompendium.TransactOpts, _minSlashInterval)
}

// SetMinSlashInterval is a paid mutator transaction binding the contract method 0xe7a6e55a.
//
// Solidity: function setMinSlashInterval(uint32 _minSlashInterval) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactorSession) SetMinSlashInterval(_minSlashInterval uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.SetMinSlashInterval(&_CloudReportCompendium.TransactOpts, _minSlashInterval)
}

// SetReportFreshnessThreshold is a paid mutator transaction binding the contract method 0xe32c101d.
//
// Solidity: function setReportFreshnessThreshold(uint32 _reportFreshnessThreshold) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactor) SetReportFreshnessThreshold(opts *bind.TransactOpts, _reportFreshnessThreshold uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.contract.Transact(opts, "setReportFreshnessThreshold", _reportFreshnessThreshold)
}

// SetReportFreshnessThreshold is a paid mutator transaction binding the contract method 0xe32c101d.
//
// Solidity: function setReportFreshnessThreshold(uint32 _reportFreshnessThreshold) returns()
func (_CloudReportCompendium *CloudReportCompendiumSession) SetReportFreshnessThreshold(_reportFreshnessThreshold uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.SetReportFreshnessThreshold(&_CloudReportCompendium.TransactOpts, _reportFreshnessThreshold)
}

// SetReportFreshnessThreshold is a paid mutator transaction binding the contract method 0xe32c101d.
//
// Solidity: function setReportFreshnessThreshold(uint32 _reportFreshnessThreshold) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactorSession) SetReportFreshnessThreshold(_reportFreshnessThreshold uint32) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.SetReportFreshnessThreshold(&_CloudReportCompendium.TransactOpts, _reportFreshnessThreshold)
}

// SetSlashAmountTokens is a paid mutator transaction binding the contract method 0x2bef26c9.
//
// Solidity: function setSlashAmountTokens(uint256 _slashAmountTokens) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactor) SetSlashAmountTokens(opts *bind.TransactOpts, _slashAmountTokens *big.Int) (*types.Transaction, error) {
	return _CloudReportCompendium.contract.Transact(opts, "setSlashAmountTokens", _slashAmountTokens)
}

// SetSlashAmountTokens is a paid mutator transaction binding the contract method 0x2bef26c9.
//
// Solidity: function setSlashAmountTokens(uint256 _slashAmountTokens) returns()
func (_CloudReportCompendium *CloudReportCompendiumSession) SetSlashAmountTokens(_slashAmountTokens *big.Int) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.SetSlashAmountTokens(&_CloudReportCompendium.TransactOpts, _slashAmountTokens)
}

// SetSlashAmountTokens is a paid mutator transaction binding the contract method 0x2bef26c9.
//
// Solidity: function setSlashAmountTokens(uint256 _slashAmountTokens) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactorSession) SetSlashAmountTokens(_slashAmountTokens *big.Int) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.SetSlashAmountTokens(&_CloudReportCompendium.TransactOpts, _slashAmountTokens)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes eigendaCert) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactor) Slash(opts *bind.TransactOpts, eigendaCert []byte) (*types.Transaction, error) {
	return _CloudReportCompendium.contract.Transact(opts, "slash", eigendaCert)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes eigendaCert) returns()
func (_CloudReportCompendium *CloudReportCompendiumSession) Slash(eigendaCert []byte) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.Slash(&_CloudReportCompendium.TransactOpts, eigendaCert)
}

// Slash is a paid mutator transaction binding the contract method 0x5e47655f.
//
// Solidity: function slash(bytes eigendaCert) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactorSession) Slash(eigendaCert []byte) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.Slash(&_CloudReportCompendium.TransactOpts, eigendaCert)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x35ce650d.
//
// Solidity: function submitReport((string,uint32,uint32,bytes) report, bytes signature) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactor) SubmitReport(opts *bind.TransactOpts, report ICloudReportCompendiumCloudReport, signature []byte) (*types.Transaction, error) {
	return _CloudReportCompendium.contract.Transact(opts, "submitReport", report, signature)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x35ce650d.
//
// Solidity: function submitReport((string,uint32,uint32,bytes) report, bytes signature) returns()
func (_CloudReportCompendium *CloudReportCompendiumSession) SubmitReport(report ICloudReportCompendiumCloudReport, signature []byte) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.SubmitReport(&_CloudReportCompendium.TransactOpts, report, signature)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x35ce650d.
//
// Solidity: function submitReport((string,uint32,uint32,bytes) report, bytes signature) returns()
func (_CloudReportCompendium *CloudReportCompendiumTransactorSession) SubmitReport(report ICloudReportCompendiumCloudReport, signature []byte) (*types.Transaction, error) {
	return _CloudReportCompendium.Contract.SubmitReport(&_CloudReportCompendium.TransactOpts, report, signature)
}

// CloudReportCompendiumInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CloudReportCompendium contract.
type CloudReportCompendiumInitializedIterator struct {
	Event *CloudReportCompendiumInitialized // Event containing the contract specifics and raw log

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
func (it *CloudReportCompendiumInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudReportCompendiumInitialized)
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
		it.Event = new(CloudReportCompendiumInitialized)
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
func (it *CloudReportCompendiumInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudReportCompendiumInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudReportCompendiumInitialized represents a Initialized event raised by the CloudReportCompendium contract.
type CloudReportCompendiumInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) FilterInitialized(opts *bind.FilterOpts) (*CloudReportCompendiumInitializedIterator, error) {

	logs, sub, err := _CloudReportCompendium.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumInitializedIterator{contract: _CloudReportCompendium.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CloudReportCompendiumInitialized) (event.Subscription, error) {

	logs, sub, err := _CloudReportCompendium.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudReportCompendiumInitialized)
				if err := _CloudReportCompendium.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CloudReportCompendium *CloudReportCompendiumFilterer) ParseInitialized(log types.Log) (*CloudReportCompendiumInitialized, error) {
	event := new(CloudReportCompendiumInitialized)
	if err := _CloudReportCompendium.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudReportCompendiumMinSlashIntervalSetIterator is returned from FilterMinSlashIntervalSet and is used to iterate over the raw logs and unpacked data for MinSlashIntervalSet events raised by the CloudReportCompendium contract.
type CloudReportCompendiumMinSlashIntervalSetIterator struct {
	Event *CloudReportCompendiumMinSlashIntervalSet // Event containing the contract specifics and raw log

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
func (it *CloudReportCompendiumMinSlashIntervalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudReportCompendiumMinSlashIntervalSet)
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
		it.Event = new(CloudReportCompendiumMinSlashIntervalSet)
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
func (it *CloudReportCompendiumMinSlashIntervalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudReportCompendiumMinSlashIntervalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudReportCompendiumMinSlashIntervalSet represents a MinSlashIntervalSet event raised by the CloudReportCompendium contract.
type CloudReportCompendiumMinSlashIntervalSet struct {
	MinSlashInterval uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMinSlashIntervalSet is a free log retrieval operation binding the contract event 0xfe8b1886127fe004b6763d8d4bc454238f70a4bbbdb1afe92d527cb1dd7adf02.
//
// Solidity: event MinSlashIntervalSet(uint32 minSlashInterval)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) FilterMinSlashIntervalSet(opts *bind.FilterOpts) (*CloudReportCompendiumMinSlashIntervalSetIterator, error) {

	logs, sub, err := _CloudReportCompendium.contract.FilterLogs(opts, "MinSlashIntervalSet")
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumMinSlashIntervalSetIterator{contract: _CloudReportCompendium.contract, event: "MinSlashIntervalSet", logs: logs, sub: sub}, nil
}

// WatchMinSlashIntervalSet is a free log subscription operation binding the contract event 0xfe8b1886127fe004b6763d8d4bc454238f70a4bbbdb1afe92d527cb1dd7adf02.
//
// Solidity: event MinSlashIntervalSet(uint32 minSlashInterval)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) WatchMinSlashIntervalSet(opts *bind.WatchOpts, sink chan<- *CloudReportCompendiumMinSlashIntervalSet) (event.Subscription, error) {

	logs, sub, err := _CloudReportCompendium.contract.WatchLogs(opts, "MinSlashIntervalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudReportCompendiumMinSlashIntervalSet)
				if err := _CloudReportCompendium.contract.UnpackLog(event, "MinSlashIntervalSet", log); err != nil {
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

// ParseMinSlashIntervalSet is a log parse operation binding the contract event 0xfe8b1886127fe004b6763d8d4bc454238f70a4bbbdb1afe92d527cb1dd7adf02.
//
// Solidity: event MinSlashIntervalSet(uint32 minSlashInterval)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) ParseMinSlashIntervalSet(log types.Log) (*CloudReportCompendiumMinSlashIntervalSet, error) {
	event := new(CloudReportCompendiumMinSlashIntervalSet)
	if err := _CloudReportCompendium.contract.UnpackLog(event, "MinSlashIntervalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudReportCompendiumOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the CloudReportCompendium contract.
type CloudReportCompendiumOperatorSlashedIterator struct {
	Event *CloudReportCompendiumOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *CloudReportCompendiumOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudReportCompendiumOperatorSlashed)
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
		it.Event = new(CloudReportCompendiumOperatorSlashed)
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
func (it *CloudReportCompendiumOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudReportCompendiumOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudReportCompendiumOperatorSlashed represents a OperatorSlashed event raised by the CloudReportCompendium contract.
type CloudReportCompendiumOperatorSlashed struct {
	Operator      common.Address
	OperatorSetId uint32
	Reason        string
	SlashId       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xb192994ade0c76d20097267fa71ca8d1e9848cf6cd9c10c961dc607280c698ea.
//
// Solidity: event OperatorSlashed(address indexed operator, uint32 indexed operatorSetId, string reason, uint256 slashId)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) FilterOperatorSlashed(opts *bind.FilterOpts, operator []common.Address, operatorSetId []uint32) (*CloudReportCompendiumOperatorSlashedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorSetIdRule []interface{}
	for _, operatorSetIdItem := range operatorSetId {
		operatorSetIdRule = append(operatorSetIdRule, operatorSetIdItem)
	}

	logs, sub, err := _CloudReportCompendium.contract.FilterLogs(opts, "OperatorSlashed", operatorRule, operatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumOperatorSlashedIterator{contract: _CloudReportCompendium.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xb192994ade0c76d20097267fa71ca8d1e9848cf6cd9c10c961dc607280c698ea.
//
// Solidity: event OperatorSlashed(address indexed operator, uint32 indexed operatorSetId, string reason, uint256 slashId)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *CloudReportCompendiumOperatorSlashed, operator []common.Address, operatorSetId []uint32) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorSetIdRule []interface{}
	for _, operatorSetIdItem := range operatorSetId {
		operatorSetIdRule = append(operatorSetIdRule, operatorSetIdItem)
	}

	logs, sub, err := _CloudReportCompendium.contract.WatchLogs(opts, "OperatorSlashed", operatorRule, operatorSetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudReportCompendiumOperatorSlashed)
				if err := _CloudReportCompendium.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0xb192994ade0c76d20097267fa71ca8d1e9848cf6cd9c10c961dc607280c698ea.
//
// Solidity: event OperatorSlashed(address indexed operator, uint32 indexed operatorSetId, string reason, uint256 slashId)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) ParseOperatorSlashed(log types.Log) (*CloudReportCompendiumOperatorSlashed, error) {
	event := new(CloudReportCompendiumOperatorSlashed)
	if err := _CloudReportCompendium.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudReportCompendiumReportFreshnessThresholdSetIterator is returned from FilterReportFreshnessThresholdSet and is used to iterate over the raw logs and unpacked data for ReportFreshnessThresholdSet events raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportFreshnessThresholdSetIterator struct {
	Event *CloudReportCompendiumReportFreshnessThresholdSet // Event containing the contract specifics and raw log

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
func (it *CloudReportCompendiumReportFreshnessThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudReportCompendiumReportFreshnessThresholdSet)
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
		it.Event = new(CloudReportCompendiumReportFreshnessThresholdSet)
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
func (it *CloudReportCompendiumReportFreshnessThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudReportCompendiumReportFreshnessThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudReportCompendiumReportFreshnessThresholdSet represents a ReportFreshnessThresholdSet event raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportFreshnessThresholdSet struct {
	ReportFreshnessThreshold uint32
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterReportFreshnessThresholdSet is a free log retrieval operation binding the contract event 0xc18f06def690706eb30b65d846e570ddc356ce32f27daa0b926c7f36e0917947.
//
// Solidity: event ReportFreshnessThresholdSet(uint32 reportFreshnessThreshold)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) FilterReportFreshnessThresholdSet(opts *bind.FilterOpts) (*CloudReportCompendiumReportFreshnessThresholdSetIterator, error) {

	logs, sub, err := _CloudReportCompendium.contract.FilterLogs(opts, "ReportFreshnessThresholdSet")
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumReportFreshnessThresholdSetIterator{contract: _CloudReportCompendium.contract, event: "ReportFreshnessThresholdSet", logs: logs, sub: sub}, nil
}

// WatchReportFreshnessThresholdSet is a free log subscription operation binding the contract event 0xc18f06def690706eb30b65d846e570ddc356ce32f27daa0b926c7f36e0917947.
//
// Solidity: event ReportFreshnessThresholdSet(uint32 reportFreshnessThreshold)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) WatchReportFreshnessThresholdSet(opts *bind.WatchOpts, sink chan<- *CloudReportCompendiumReportFreshnessThresholdSet) (event.Subscription, error) {

	logs, sub, err := _CloudReportCompendium.contract.WatchLogs(opts, "ReportFreshnessThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudReportCompendiumReportFreshnessThresholdSet)
				if err := _CloudReportCompendium.contract.UnpackLog(event, "ReportFreshnessThresholdSet", log); err != nil {
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

// ParseReportFreshnessThresholdSet is a log parse operation binding the contract event 0xc18f06def690706eb30b65d846e570ddc356ce32f27daa0b926c7f36e0917947.
//
// Solidity: event ReportFreshnessThresholdSet(uint32 reportFreshnessThreshold)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) ParseReportFreshnessThresholdSet(log types.Log) (*CloudReportCompendiumReportFreshnessThresholdSet, error) {
	event := new(CloudReportCompendiumReportFreshnessThresholdSet)
	if err := _CloudReportCompendium.contract.UnpackLog(event, "ReportFreshnessThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudReportCompendiumReportSubmittedIterator is returned from FilterReportSubmitted and is used to iterate over the raw logs and unpacked data for ReportSubmitted events raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportSubmittedIterator struct {
	Event *CloudReportCompendiumReportSubmitted // Event containing the contract specifics and raw log

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
func (it *CloudReportCompendiumReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudReportCompendiumReportSubmitted)
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
		it.Event = new(CloudReportCompendiumReportSubmitted)
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
func (it *CloudReportCompendiumReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudReportCompendiumReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudReportCompendiumReportSubmitted represents a ReportSubmitted event raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportSubmitted struct {
	Report    ICloudReportCompendiumCloudReport
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReportSubmitted is a free log retrieval operation binding the contract event 0x120822bef9c8cbefdd48190a56b37379ededb9577ec38b15de29f5a89253167c.
//
// Solidity: event ReportSubmitted((string,uint32,uint32,bytes) report, uint256 timestamp)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) FilterReportSubmitted(opts *bind.FilterOpts) (*CloudReportCompendiumReportSubmittedIterator, error) {

	logs, sub, err := _CloudReportCompendium.contract.FilterLogs(opts, "ReportSubmitted")
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumReportSubmittedIterator{contract: _CloudReportCompendium.contract, event: "ReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchReportSubmitted is a free log subscription operation binding the contract event 0x120822bef9c8cbefdd48190a56b37379ededb9577ec38b15de29f5a89253167c.
//
// Solidity: event ReportSubmitted((string,uint32,uint32,bytes) report, uint256 timestamp)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) WatchReportSubmitted(opts *bind.WatchOpts, sink chan<- *CloudReportCompendiumReportSubmitted) (event.Subscription, error) {

	logs, sub, err := _CloudReportCompendium.contract.WatchLogs(opts, "ReportSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudReportCompendiumReportSubmitted)
				if err := _CloudReportCompendium.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
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

// ParseReportSubmitted is a log parse operation binding the contract event 0x120822bef9c8cbefdd48190a56b37379ededb9577ec38b15de29f5a89253167c.
//
// Solidity: event ReportSubmitted((string,uint32,uint32,bytes) report, uint256 timestamp)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) ParseReportSubmitted(log types.Log) (*CloudReportCompendiumReportSubmitted, error) {
	event := new(CloudReportCompendiumReportSubmitted)
	if err := _CloudReportCompendium.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudReportCompendiumSlashAmountTokensSetIterator is returned from FilterSlashAmountTokensSet and is used to iterate over the raw logs and unpacked data for SlashAmountTokensSet events raised by the CloudReportCompendium contract.
type CloudReportCompendiumSlashAmountTokensSetIterator struct {
	Event *CloudReportCompendiumSlashAmountTokensSet // Event containing the contract specifics and raw log

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
func (it *CloudReportCompendiumSlashAmountTokensSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudReportCompendiumSlashAmountTokensSet)
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
		it.Event = new(CloudReportCompendiumSlashAmountTokensSet)
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
func (it *CloudReportCompendiumSlashAmountTokensSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudReportCompendiumSlashAmountTokensSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudReportCompendiumSlashAmountTokensSet represents a SlashAmountTokensSet event raised by the CloudReportCompendium contract.
type CloudReportCompendiumSlashAmountTokensSet struct {
	SlashAmountTokens *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSlashAmountTokensSet is a free log retrieval operation binding the contract event 0xbc4559d0d67e98a959def403cabda2f31d878c322d0233e4280a910698b20a93.
//
// Solidity: event SlashAmountTokensSet(uint256 slashAmountTokens)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) FilterSlashAmountTokensSet(opts *bind.FilterOpts) (*CloudReportCompendiumSlashAmountTokensSetIterator, error) {

	logs, sub, err := _CloudReportCompendium.contract.FilterLogs(opts, "SlashAmountTokensSet")
	if err != nil {
		return nil, err
	}
	return &CloudReportCompendiumSlashAmountTokensSetIterator{contract: _CloudReportCompendium.contract, event: "SlashAmountTokensSet", logs: logs, sub: sub}, nil
}

// WatchSlashAmountTokensSet is a free log subscription operation binding the contract event 0xbc4559d0d67e98a959def403cabda2f31d878c322d0233e4280a910698b20a93.
//
// Solidity: event SlashAmountTokensSet(uint256 slashAmountTokens)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) WatchSlashAmountTokensSet(opts *bind.WatchOpts, sink chan<- *CloudReportCompendiumSlashAmountTokensSet) (event.Subscription, error) {

	logs, sub, err := _CloudReportCompendium.contract.WatchLogs(opts, "SlashAmountTokensSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudReportCompendiumSlashAmountTokensSet)
				if err := _CloudReportCompendium.contract.UnpackLog(event, "SlashAmountTokensSet", log); err != nil {
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

// ParseSlashAmountTokensSet is a log parse operation binding the contract event 0xbc4559d0d67e98a959def403cabda2f31d878c322d0233e4280a910698b20a93.
//
// Solidity: event SlashAmountTokensSet(uint256 slashAmountTokens)
func (_CloudReportCompendium *CloudReportCompendiumFilterer) ParseSlashAmountTokensSet(log types.Log) (*CloudReportCompendiumSlashAmountTokensSet, error) {
	event := new(CloudReportCompendiumSlashAmountTokensSet)
	if err := _CloudReportCompendium.contract.UnpackLog(event, "SlashAmountTokensSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
