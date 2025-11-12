// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CloudReportCompendium

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

// CloudReportCompendiumStorageConstructorParams is an auto generated low-level Go binding around an user-defined struct.
type CloudReportCompendiumStorageConstructorParams struct {
	DelegationManager   common.Address
	AllocationManager   common.Address
	CertVerifierRouter  common.Address
	ComputeAVSRegistrar common.Address
	ComputeOperator     common.Address
	ReportAttester      common.Address
	OperatorSetId       uint32
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
	FromTimestamp       uint32
	ToTimestamp         uint32
	EigendaCertHash     [20]byte
}

// CloudReportCompendiumMetaData contains all meta data concerning the CloudReportCompendium contract.
var CloudReportCompendiumMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structCloudReportCompendiumStorage.ConstructorParams\",\"components\":[{\"name\":\"delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"certVerifierRouter\",\"type\":\"address\",\"internalType\":\"contractIEigenDACertVerifierRouter\"},{\"name\":\"computeAVSRegistrar\",\"type\":\"address\",\"internalType\":\"contractIComputeAVSRegistrar\"},{\"name\":\"computeOperator\",\"type\":\"address\",\"internalType\":\"contractIComputeOperator\"},{\"name\":\"reportAttester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"referenceProjectId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"strategyToSlash\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CLOUD_REPORT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateReportDigestHash\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structICloudReportCompendium.CloudReport\",\"components\":[{\"name\":\"projectId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fromTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"toTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eigendaCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"certVerifierRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenDACertVerifierRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeAVSRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIComputeAVSRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIComputeOperator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_reportFreshnessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_slashAmountTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestReportSubmission\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICloudReportCompendium.CloudReportSubmission\",\"components\":[{\"name\":\"submissionTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"toTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eigendaCertHash\",\"type\":\"bytes20\",\"internalType\":\"bytes20\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"referenceProjectIdHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reportAttester\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reportFreshnessThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setReportFreshnessThreshold\",\"inputs\":[{\"name\":\"_reportFreshnessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlashAmountTokens\",\"inputs\":[{\"name\":\"_slashAmountTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slash\",\"inputs\":[{\"name\":\"eigendaCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashAmountTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyToSlash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitReport\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structICloudReportCompendium.CloudReport\",\"components\":[{\"name\":\"projectId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fromTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"toTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eigendaCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportFreshnessThresholdSet\",\"inputs\":[{\"name\":\"reportFreshnessThreshold\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportSubmitted\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICloudReportCompendium.CloudReport\",\"components\":[{\"name\":\"projectId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fromTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"toTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eigendaCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashAmountTokensSet\",\"inputs\":[{\"name\":\"slashAmountTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientAllocation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCertificate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProjectId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReportTimestamps\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoSlashRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportChainTimestampMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportFromFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReportTooStale\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	ID:  "CloudReportCompendium",
}

// CloudReportCompendium is an auto generated Go binding around an Ethereum contract.
type CloudReportCompendium struct {
	abi abi.ABI
}

// NewCloudReportCompendium creates a new instance of CloudReportCompendium.
func NewCloudReportCompendium() *CloudReportCompendium {
	parsed, err := CloudReportCompendiumMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CloudReportCompendium{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CloudReportCompendium) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(string _version, address _permissionController, (address,address,address,address,address,address,uint32,string,address) params) returns()
func (cloudReportCompendium *CloudReportCompendium) PackConstructor(_version string, _permissionController common.Address, params CloudReportCompendiumStorageConstructorParams) []byte {
	enc, err := cloudReportCompendium.abi.Pack("", _version, _permissionController, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCLOUDREPORTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x593ea37d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function CLOUD_REPORT_TYPEHASH() view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) PackCLOUDREPORTTYPEHASH() []byte {
	enc, err := cloudReportCompendium.abi.Pack("CLOUD_REPORT_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCLOUDREPORTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x593ea37d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function CLOUD_REPORT_TYPEHASH() view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) TryPackCLOUDREPORTTYPEHASH() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("CLOUD_REPORT_TYPEHASH")
}

// UnpackCLOUDREPORTTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x593ea37d.
//
// Solidity: function CLOUD_REPORT_TYPEHASH() view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) UnpackCLOUDREPORTTYPEHASH(data []byte) ([32]byte, error) {
	out, err := cloudReportCompendium.abi.Unpack("CLOUD_REPORT_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackAllocationManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca8aa7c7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allocationManager() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) PackAllocationManager() []byte {
	enc, err := cloudReportCompendium.abi.Pack("allocationManager")
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
func (cloudReportCompendium *CloudReportCompendium) TryPackAllocationManager() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("allocationManager")
}

// UnpackAllocationManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) UnpackAllocationManager(data []byte) (common.Address, error) {
	out, err := cloudReportCompendium.abi.Unpack("allocationManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCalculateReportDigestHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xee1d5079.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateReportDigestHash((string,uint32,uint32,bytes) report) view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) PackCalculateReportDigestHash(report ICloudReportCompendiumCloudReport) []byte {
	enc, err := cloudReportCompendium.abi.Pack("calculateReportDigestHash", report)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateReportDigestHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xee1d5079.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateReportDigestHash((string,uint32,uint32,bytes) report) view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) TryPackCalculateReportDigestHash(report ICloudReportCompendiumCloudReport) ([]byte, error) {
	return cloudReportCompendium.abi.Pack("calculateReportDigestHash", report)
}

// UnpackCalculateReportDigestHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xee1d5079.
//
// Solidity: function calculateReportDigestHash((string,uint32,uint32,bytes) report) view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) UnpackCalculateReportDigestHash(data []byte) ([32]byte, error) {
	out, err := cloudReportCompendium.abi.Unpack("calculateReportDigestHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackCertVerifierRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10d36975.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function certVerifierRouter() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) PackCertVerifierRouter() []byte {
	enc, err := cloudReportCompendium.abi.Pack("certVerifierRouter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCertVerifierRouter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10d36975.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function certVerifierRouter() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) TryPackCertVerifierRouter() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("certVerifierRouter")
}

// UnpackCertVerifierRouter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x10d36975.
//
// Solidity: function certVerifierRouter() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) UnpackCertVerifierRouter(data []byte) (common.Address, error) {
	out, err := cloudReportCompendium.abi.Unpack("certVerifierRouter", data)
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
func (cloudReportCompendium *CloudReportCompendium) PackComputeAVSRegistrar() []byte {
	enc, err := cloudReportCompendium.abi.Pack("computeAVSRegistrar")
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
func (cloudReportCompendium *CloudReportCompendium) TryPackComputeAVSRegistrar() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("computeAVSRegistrar")
}

// UnpackComputeAVSRegistrar is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xef6d92c6.
//
// Solidity: function computeAVSRegistrar() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) UnpackComputeAVSRegistrar(data []byte) (common.Address, error) {
	out, err := cloudReportCompendium.abi.Unpack("computeAVSRegistrar", data)
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
func (cloudReportCompendium *CloudReportCompendium) PackComputeOperator() []byte {
	enc, err := cloudReportCompendium.abi.Pack("computeOperator")
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
func (cloudReportCompendium *CloudReportCompendium) TryPackComputeOperator() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("computeOperator")
}

// UnpackComputeOperator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3fc4cbfd.
//
// Solidity: function computeOperator() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) UnpackComputeOperator(data []byte) (common.Address, error) {
	out, err := cloudReportCompendium.abi.Unpack("computeOperator", data)
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
func (cloudReportCompendium *CloudReportCompendium) PackDelegationManager() []byte {
	enc, err := cloudReportCompendium.abi.Pack("delegationManager")
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
func (cloudReportCompendium *CloudReportCompendium) TryPackDelegationManager() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("delegationManager")
}

// UnpackDelegationManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) UnpackDelegationManager(data []byte) (common.Address, error) {
	out, err := cloudReportCompendium.abi.Unpack("delegationManager", data)
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
func (cloudReportCompendium *CloudReportCompendium) PackDomainSeparator() []byte {
	enc, err := cloudReportCompendium.abi.Pack("domainSeparator")
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
func (cloudReportCompendium *CloudReportCompendium) TryPackDomainSeparator() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("domainSeparator")
}

// UnpackDomainSeparator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) UnpackDomainSeparator(data []byte) ([32]byte, error) {
	out, err := cloudReportCompendium.abi.Unpack("domainSeparator", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x282558d9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(uint32 _reportFreshnessThreshold, uint256 _slashAmountTokens) returns()
func (cloudReportCompendium *CloudReportCompendium) PackInitialize(reportFreshnessThreshold uint32, slashAmountTokens *big.Int) []byte {
	enc, err := cloudReportCompendium.abi.Pack("initialize", reportFreshnessThreshold, slashAmountTokens)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x282558d9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(uint32 _reportFreshnessThreshold, uint256 _slashAmountTokens) returns()
func (cloudReportCompendium *CloudReportCompendium) TryPackInitialize(reportFreshnessThreshold uint32, slashAmountTokens *big.Int) ([]byte, error) {
	return cloudReportCompendium.abi.Pack("initialize", reportFreshnessThreshold, slashAmountTokens)
}

// PackLatestReportSubmission is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x371517a4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function latestReportSubmission() view returns((uint32,uint32,uint32,bytes20))
func (cloudReportCompendium *CloudReportCompendium) PackLatestReportSubmission() []byte {
	enc, err := cloudReportCompendium.abi.Pack("latestReportSubmission")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLatestReportSubmission is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x371517a4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function latestReportSubmission() view returns((uint32,uint32,uint32,bytes20))
func (cloudReportCompendium *CloudReportCompendium) TryPackLatestReportSubmission() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("latestReportSubmission")
}

// UnpackLatestReportSubmission is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x371517a4.
//
// Solidity: function latestReportSubmission() view returns((uint32,uint32,uint32,bytes20))
func (cloudReportCompendium *CloudReportCompendium) UnpackLatestReportSubmission(data []byte) (ICloudReportCompendiumCloudReportSubmission, error) {
	out, err := cloudReportCompendium.abi.Unpack("latestReportSubmission", data)
	if err != nil {
		return *new(ICloudReportCompendiumCloudReportSubmission), err
	}
	out0 := *abi.ConvertType(out[0], new(ICloudReportCompendiumCloudReportSubmission)).(*ICloudReportCompendiumCloudReportSubmission)
	return out0, nil
}

// PackOperatorSetId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1ebfc37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function operatorSetId() view returns(uint32)
func (cloudReportCompendium *CloudReportCompendium) PackOperatorSetId() []byte {
	enc, err := cloudReportCompendium.abi.Pack("operatorSetId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOperatorSetId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1ebfc37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function operatorSetId() view returns(uint32)
func (cloudReportCompendium *CloudReportCompendium) TryPackOperatorSetId() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("operatorSetId")
}

// UnpackOperatorSetId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe1ebfc37.
//
// Solidity: function operatorSetId() view returns(uint32)
func (cloudReportCompendium *CloudReportCompendium) UnpackOperatorSetId(data []byte) (uint32, error) {
	out, err := cloudReportCompendium.abi.Unpack("operatorSetId", data)
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
func (cloudReportCompendium *CloudReportCompendium) PackPermissionController() []byte {
	enc, err := cloudReportCompendium.abi.Pack("permissionController")
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
func (cloudReportCompendium *CloudReportCompendium) TryPackPermissionController() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("permissionController")
}

// UnpackPermissionController is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) UnpackPermissionController(data []byte) (common.Address, error) {
	out, err := cloudReportCompendium.abi.Unpack("permissionController", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackReferenceProjectIdHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2e9ef5d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function referenceProjectIdHash() view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) PackReferenceProjectIdHash() []byte {
	enc, err := cloudReportCompendium.abi.Pack("referenceProjectIdHash")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReferenceProjectIdHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2e9ef5d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function referenceProjectIdHash() view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) TryPackReferenceProjectIdHash() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("referenceProjectIdHash")
}

// UnpackReferenceProjectIdHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2e9ef5d.
//
// Solidity: function referenceProjectIdHash() view returns(bytes32)
func (cloudReportCompendium *CloudReportCompendium) UnpackReferenceProjectIdHash(data []byte) ([32]byte, error) {
	out, err := cloudReportCompendium.abi.Unpack("referenceProjectIdHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackReportAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x285c6e4f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reportAttester() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) PackReportAttester() []byte {
	enc, err := cloudReportCompendium.abi.Pack("reportAttester")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReportAttester is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x285c6e4f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reportAttester() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) TryPackReportAttester() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("reportAttester")
}

// UnpackReportAttester is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x285c6e4f.
//
// Solidity: function reportAttester() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) UnpackReportAttester(data []byte) (common.Address, error) {
	out, err := cloudReportCompendium.abi.Unpack("reportAttester", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackReportFreshnessThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5bc35ef6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reportFreshnessThreshold() view returns(uint32)
func (cloudReportCompendium *CloudReportCompendium) PackReportFreshnessThreshold() []byte {
	enc, err := cloudReportCompendium.abi.Pack("reportFreshnessThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReportFreshnessThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5bc35ef6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reportFreshnessThreshold() view returns(uint32)
func (cloudReportCompendium *CloudReportCompendium) TryPackReportFreshnessThreshold() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("reportFreshnessThreshold")
}

// UnpackReportFreshnessThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5bc35ef6.
//
// Solidity: function reportFreshnessThreshold() view returns(uint32)
func (cloudReportCompendium *CloudReportCompendium) UnpackReportFreshnessThreshold(data []byte) (uint32, error) {
	out, err := cloudReportCompendium.abi.Unpack("reportFreshnessThreshold", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackSetReportFreshnessThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe32c101d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setReportFreshnessThreshold(uint32 _reportFreshnessThreshold) returns()
func (cloudReportCompendium *CloudReportCompendium) PackSetReportFreshnessThreshold(reportFreshnessThreshold uint32) []byte {
	enc, err := cloudReportCompendium.abi.Pack("setReportFreshnessThreshold", reportFreshnessThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetReportFreshnessThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe32c101d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setReportFreshnessThreshold(uint32 _reportFreshnessThreshold) returns()
func (cloudReportCompendium *CloudReportCompendium) TryPackSetReportFreshnessThreshold(reportFreshnessThreshold uint32) ([]byte, error) {
	return cloudReportCompendium.abi.Pack("setReportFreshnessThreshold", reportFreshnessThreshold)
}

// PackSetSlashAmountTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2bef26c9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSlashAmountTokens(uint256 _slashAmountTokens) returns()
func (cloudReportCompendium *CloudReportCompendium) PackSetSlashAmountTokens(slashAmountTokens *big.Int) []byte {
	enc, err := cloudReportCompendium.abi.Pack("setSlashAmountTokens", slashAmountTokens)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSlashAmountTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2bef26c9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSlashAmountTokens(uint256 _slashAmountTokens) returns()
func (cloudReportCompendium *CloudReportCompendium) TryPackSetSlashAmountTokens(slashAmountTokens *big.Int) ([]byte, error) {
	return cloudReportCompendium.abi.Pack("setSlashAmountTokens", slashAmountTokens)
}

// PackSlash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e47655f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function slash(bytes eigendaCert) returns()
func (cloudReportCompendium *CloudReportCompendium) PackSlash(eigendaCert []byte) []byte {
	enc, err := cloudReportCompendium.abi.Pack("slash", eigendaCert)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSlash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e47655f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function slash(bytes eigendaCert) returns()
func (cloudReportCompendium *CloudReportCompendium) TryPackSlash(eigendaCert []byte) ([]byte, error) {
	return cloudReportCompendium.abi.Pack("slash", eigendaCert)
}

// PackSlashAmountTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d3ceb34.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function slashAmountTokens() view returns(uint256)
func (cloudReportCompendium *CloudReportCompendium) PackSlashAmountTokens() []byte {
	enc, err := cloudReportCompendium.abi.Pack("slashAmountTokens")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSlashAmountTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d3ceb34.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function slashAmountTokens() view returns(uint256)
func (cloudReportCompendium *CloudReportCompendium) TryPackSlashAmountTokens() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("slashAmountTokens")
}

// UnpackSlashAmountTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7d3ceb34.
//
// Solidity: function slashAmountTokens() view returns(uint256)
func (cloudReportCompendium *CloudReportCompendium) UnpackSlashAmountTokens(data []byte) (*big.Int, error) {
	out, err := cloudReportCompendium.abi.Unpack("slashAmountTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackStrategyToSlash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ac080f1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function strategyToSlash() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) PackStrategyToSlash() []byte {
	enc, err := cloudReportCompendium.abi.Pack("strategyToSlash")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStrategyToSlash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ac080f1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function strategyToSlash() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) TryPackStrategyToSlash() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("strategyToSlash")
}

// UnpackStrategyToSlash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5ac080f1.
//
// Solidity: function strategyToSlash() view returns(address)
func (cloudReportCompendium *CloudReportCompendium) UnpackStrategyToSlash(data []byte) (common.Address, error) {
	out, err := cloudReportCompendium.abi.Unpack("strategyToSlash", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSubmitReport is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35ce650d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitReport((string,uint32,uint32,bytes) report, bytes signature) returns()
func (cloudReportCompendium *CloudReportCompendium) PackSubmitReport(report ICloudReportCompendiumCloudReport, signature []byte) []byte {
	enc, err := cloudReportCompendium.abi.Pack("submitReport", report, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitReport is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35ce650d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitReport((string,uint32,uint32,bytes) report, bytes signature) returns()
func (cloudReportCompendium *CloudReportCompendium) TryPackSubmitReport(report ICloudReportCompendiumCloudReport, signature []byte) ([]byte, error) {
	return cloudReportCompendium.abi.Pack("submitReport", report, signature)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(string)
func (cloudReportCompendium *CloudReportCompendium) PackVersion() []byte {
	enc, err := cloudReportCompendium.abi.Pack("version")
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
func (cloudReportCompendium *CloudReportCompendium) TryPackVersion() ([]byte, error) {
	return cloudReportCompendium.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (cloudReportCompendium *CloudReportCompendium) UnpackVersion(data []byte) (string, error) {
	out, err := cloudReportCompendium.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// CloudReportCompendiumInitialized represents a Initialized event raised by the CloudReportCompendium contract.
type CloudReportCompendiumInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const CloudReportCompendiumInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (CloudReportCompendiumInitialized) ContractEventName() string {
	return CloudReportCompendiumInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (cloudReportCompendium *CloudReportCompendium) UnpackInitializedEvent(log *types.Log) (*CloudReportCompendiumInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != cloudReportCompendium.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CloudReportCompendiumInitialized)
	if len(log.Data) > 0 {
		if err := cloudReportCompendium.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cloudReportCompendium.abi.Events[event].Inputs {
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

// CloudReportCompendiumOperatorSlashed represents a OperatorSlashed event raised by the CloudReportCompendium contract.
type CloudReportCompendiumOperatorSlashed struct {
	Operator      common.Address
	OperatorSetId uint32
	Reason        string
	SlashId       *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const CloudReportCompendiumOperatorSlashedEventName = "OperatorSlashed"

// ContractEventName returns the user-defined event name.
func (CloudReportCompendiumOperatorSlashed) ContractEventName() string {
	return CloudReportCompendiumOperatorSlashedEventName
}

// UnpackOperatorSlashedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperatorSlashed(address indexed operator, uint32 indexed operatorSetId, string reason, uint256 slashId)
func (cloudReportCompendium *CloudReportCompendium) UnpackOperatorSlashedEvent(log *types.Log) (*CloudReportCompendiumOperatorSlashed, error) {
	event := "OperatorSlashed"
	if len(log.Topics) == 0 || log.Topics[0] != cloudReportCompendium.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CloudReportCompendiumOperatorSlashed)
	if len(log.Data) > 0 {
		if err := cloudReportCompendium.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cloudReportCompendium.abi.Events[event].Inputs {
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

// CloudReportCompendiumReportFreshnessThresholdSet represents a ReportFreshnessThresholdSet event raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportFreshnessThresholdSet struct {
	ReportFreshnessThreshold uint32
	Raw                      *types.Log // Blockchain specific contextual infos
}

const CloudReportCompendiumReportFreshnessThresholdSetEventName = "ReportFreshnessThresholdSet"

// ContractEventName returns the user-defined event name.
func (CloudReportCompendiumReportFreshnessThresholdSet) ContractEventName() string {
	return CloudReportCompendiumReportFreshnessThresholdSetEventName
}

// UnpackReportFreshnessThresholdSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ReportFreshnessThresholdSet(uint32 reportFreshnessThreshold)
func (cloudReportCompendium *CloudReportCompendium) UnpackReportFreshnessThresholdSetEvent(log *types.Log) (*CloudReportCompendiumReportFreshnessThresholdSet, error) {
	event := "ReportFreshnessThresholdSet"
	if len(log.Topics) == 0 || log.Topics[0] != cloudReportCompendium.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CloudReportCompendiumReportFreshnessThresholdSet)
	if len(log.Data) > 0 {
		if err := cloudReportCompendium.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cloudReportCompendium.abi.Events[event].Inputs {
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

// CloudReportCompendiumReportSubmitted represents a ReportSubmitted event raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportSubmitted struct {
	Report    ICloudReportCompendiumCloudReport
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const CloudReportCompendiumReportSubmittedEventName = "ReportSubmitted"

// ContractEventName returns the user-defined event name.
func (CloudReportCompendiumReportSubmitted) ContractEventName() string {
	return CloudReportCompendiumReportSubmittedEventName
}

// UnpackReportSubmittedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ReportSubmitted((string,uint32,uint32,bytes) report, uint256 timestamp)
func (cloudReportCompendium *CloudReportCompendium) UnpackReportSubmittedEvent(log *types.Log) (*CloudReportCompendiumReportSubmitted, error) {
	event := "ReportSubmitted"
	if len(log.Topics) == 0 || log.Topics[0] != cloudReportCompendium.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CloudReportCompendiumReportSubmitted)
	if len(log.Data) > 0 {
		if err := cloudReportCompendium.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cloudReportCompendium.abi.Events[event].Inputs {
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

// CloudReportCompendiumSlashAmountTokensSet represents a SlashAmountTokensSet event raised by the CloudReportCompendium contract.
type CloudReportCompendiumSlashAmountTokensSet struct {
	SlashAmountTokens *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const CloudReportCompendiumSlashAmountTokensSetEventName = "SlashAmountTokensSet"

// ContractEventName returns the user-defined event name.
func (CloudReportCompendiumSlashAmountTokensSet) ContractEventName() string {
	return CloudReportCompendiumSlashAmountTokensSetEventName
}

// UnpackSlashAmountTokensSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SlashAmountTokensSet(uint256 slashAmountTokens)
func (cloudReportCompendium *CloudReportCompendium) UnpackSlashAmountTokensSetEvent(log *types.Log) (*CloudReportCompendiumSlashAmountTokensSet, error) {
	event := "SlashAmountTokensSet"
	if len(log.Topics) == 0 || log.Topics[0] != cloudReportCompendium.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CloudReportCompendiumSlashAmountTokensSet)
	if len(log.Data) > 0 {
		if err := cloudReportCompendium.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cloudReportCompendium.abi.Events[event].Inputs {
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
func (cloudReportCompendium *CloudReportCompendium) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["InsufficientAllocation"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackInsufficientAllocationError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["InvalidCertificate"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackInvalidCertificateError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["InvalidPermissions"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackInvalidPermissionsError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["InvalidProjectId"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackInvalidProjectIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["InvalidReportTimestamps"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackInvalidReportTimestampsError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["InvalidSignature"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["NoSlashRequired"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackNoSlashRequiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["ReportChainTimestampMismatch"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackReportChainTimestampMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["ReportFromFuture"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackReportFromFutureError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["ReportTooStale"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackReportTooStaleError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["SignatureExpired"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackSignatureExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], cloudReportCompendium.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return cloudReportCompendium.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// CloudReportCompendiumInsufficientAllocation represents a InsufficientAllocation error raised by the CloudReportCompendium contract.
type CloudReportCompendiumInsufficientAllocation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientAllocation()
func CloudReportCompendiumInsufficientAllocationErrorID() common.Hash {
	return common.HexToHash("0xdf2d4774c7cfdf4b27ad6553bec24b489a3b3201e3a9418e8f4926fe00431ffd")
}

// UnpackInsufficientAllocationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientAllocation()
func (cloudReportCompendium *CloudReportCompendium) UnpackInsufficientAllocationError(raw []byte) (*CloudReportCompendiumInsufficientAllocation, error) {
	out := new(CloudReportCompendiumInsufficientAllocation)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "InsufficientAllocation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumInvalidCertificate represents a InvalidCertificate error raised by the CloudReportCompendium contract.
type CloudReportCompendiumInvalidCertificate struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCertificate()
func CloudReportCompendiumInvalidCertificateErrorID() common.Hash {
	return common.HexToHash("0x1d39f94682eb2e7a1e48043ed42e27dcfea3c7e8a7c7f80ec1c2f7dae3a1c97a")
}

// UnpackInvalidCertificateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCertificate()
func (cloudReportCompendium *CloudReportCompendium) UnpackInvalidCertificateError(raw []byte) (*CloudReportCompendiumInvalidCertificate, error) {
	out := new(CloudReportCompendiumInvalidCertificate)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "InvalidCertificate", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumInvalidPermissions represents a InvalidPermissions error raised by the CloudReportCompendium contract.
type CloudReportCompendiumInvalidPermissions struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPermissions()
func CloudReportCompendiumInvalidPermissionsErrorID() common.Hash {
	return common.HexToHash("0x932d94f726428388537b641940dd88f9f37f70be827ee507792b87e4d26875f9")
}

// UnpackInvalidPermissionsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPermissions()
func (cloudReportCompendium *CloudReportCompendium) UnpackInvalidPermissionsError(raw []byte) (*CloudReportCompendiumInvalidPermissions, error) {
	out := new(CloudReportCompendiumInvalidPermissions)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "InvalidPermissions", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumInvalidProjectId represents a InvalidProjectId error raised by the CloudReportCompendium contract.
type CloudReportCompendiumInvalidProjectId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidProjectId()
func CloudReportCompendiumInvalidProjectIdErrorID() common.Hash {
	return common.HexToHash("0xabcd0b5c8ca7446a866b21429aaa322a60b28a925757755a9452b3da8b152cc1")
}

// UnpackInvalidProjectIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidProjectId()
func (cloudReportCompendium *CloudReportCompendium) UnpackInvalidProjectIdError(raw []byte) (*CloudReportCompendiumInvalidProjectId, error) {
	out := new(CloudReportCompendiumInvalidProjectId)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "InvalidProjectId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumInvalidReportTimestamps represents a InvalidReportTimestamps error raised by the CloudReportCompendium contract.
type CloudReportCompendiumInvalidReportTimestamps struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidReportTimestamps()
func CloudReportCompendiumInvalidReportTimestampsErrorID() common.Hash {
	return common.HexToHash("0x4f6dc836c400201281f1a7561a36f0aa5a0a999af82692bfb6601ed86a660d3e")
}

// UnpackInvalidReportTimestampsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidReportTimestamps()
func (cloudReportCompendium *CloudReportCompendium) UnpackInvalidReportTimestampsError(raw []byte) (*CloudReportCompendiumInvalidReportTimestamps, error) {
	out := new(CloudReportCompendiumInvalidReportTimestamps)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "InvalidReportTimestamps", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumInvalidShortString represents a InvalidShortString error raised by the CloudReportCompendium contract.
type CloudReportCompendiumInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func CloudReportCompendiumInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (cloudReportCompendium *CloudReportCompendium) UnpackInvalidShortStringError(raw []byte) (*CloudReportCompendiumInvalidShortString, error) {
	out := new(CloudReportCompendiumInvalidShortString)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumInvalidSignature represents a InvalidSignature error raised by the CloudReportCompendium contract.
type CloudReportCompendiumInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSignature()
func CloudReportCompendiumInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0x8baa579fce362245063d36f11747a89dd489c54795634fc673cc0e0db51fedc5")
}

// UnpackInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSignature()
func (cloudReportCompendium *CloudReportCompendium) UnpackInvalidSignatureError(raw []byte) (*CloudReportCompendiumInvalidSignature, error) {
	out := new(CloudReportCompendiumInvalidSignature)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "InvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumNoSlashRequired represents a NoSlashRequired error raised by the CloudReportCompendium contract.
type CloudReportCompendiumNoSlashRequired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoSlashRequired()
func CloudReportCompendiumNoSlashRequiredErrorID() common.Hash {
	return common.HexToHash("0xb4f5fa5a011726121c0d298940aea06d451fc4a04cbbcf01e77b128ada7b8a87")
}

// UnpackNoSlashRequiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoSlashRequired()
func (cloudReportCompendium *CloudReportCompendium) UnpackNoSlashRequiredError(raw []byte) (*CloudReportCompendiumNoSlashRequired, error) {
	out := new(CloudReportCompendiumNoSlashRequired)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "NoSlashRequired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumReportChainTimestampMismatch represents a ReportChainTimestampMismatch error raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportChainTimestampMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReportChainTimestampMismatch()
func CloudReportCompendiumReportChainTimestampMismatchErrorID() common.Hash {
	return common.HexToHash("0xb19f39e95dd1674ddd44f1d3550568ccc4cc4b2275a8ef5f8f33eaf9e49f5b9b")
}

// UnpackReportChainTimestampMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReportChainTimestampMismatch()
func (cloudReportCompendium *CloudReportCompendium) UnpackReportChainTimestampMismatchError(raw []byte) (*CloudReportCompendiumReportChainTimestampMismatch, error) {
	out := new(CloudReportCompendiumReportChainTimestampMismatch)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "ReportChainTimestampMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumReportFromFuture represents a ReportFromFuture error raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportFromFuture struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReportFromFuture()
func CloudReportCompendiumReportFromFutureErrorID() common.Hash {
	return common.HexToHash("0xc1208954973cfa6377fe7ddc754262a69c9c624ae1f4fd2cfd2e5050c939ee15")
}

// UnpackReportFromFutureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReportFromFuture()
func (cloudReportCompendium *CloudReportCompendium) UnpackReportFromFutureError(raw []byte) (*CloudReportCompendiumReportFromFuture, error) {
	out := new(CloudReportCompendiumReportFromFuture)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "ReportFromFuture", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumReportTooStale represents a ReportTooStale error raised by the CloudReportCompendium contract.
type CloudReportCompendiumReportTooStale struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReportTooStale()
func CloudReportCompendiumReportTooStaleErrorID() common.Hash {
	return common.HexToHash("0xc262d324b1569da9488224a8c2d9e0cf753ee43616a13e7b2bf9baea921906c1")
}

// UnpackReportTooStaleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReportTooStale()
func (cloudReportCompendium *CloudReportCompendium) UnpackReportTooStaleError(raw []byte) (*CloudReportCompendiumReportTooStale, error) {
	out := new(CloudReportCompendiumReportTooStale)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "ReportTooStale", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumSignatureExpired represents a SignatureExpired error raised by the CloudReportCompendium contract.
type CloudReportCompendiumSignatureExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SignatureExpired()
func CloudReportCompendiumSignatureExpiredErrorID() common.Hash {
	return common.HexToHash("0x0819bdcd8da3b255f8b8bf8497982cf672847d37ad445f8f2edca874c1fcb6a3")
}

// UnpackSignatureExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SignatureExpired()
func (cloudReportCompendium *CloudReportCompendium) UnpackSignatureExpiredError(raw []byte) (*CloudReportCompendiumSignatureExpired, error) {
	out := new(CloudReportCompendiumSignatureExpired)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "SignatureExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudReportCompendiumStringTooLong represents a StringTooLong error raised by the CloudReportCompendium contract.
type CloudReportCompendiumStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func CloudReportCompendiumStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (cloudReportCompendium *CloudReportCompendium) UnpackStringTooLongError(raw []byte) (*CloudReportCompendiumStringTooLong, error) {
	out := new(CloudReportCompendiumStringTooLong)
	if err := cloudReportCompendium.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}
