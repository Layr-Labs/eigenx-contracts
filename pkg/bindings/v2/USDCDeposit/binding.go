// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package USDCDeposit

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

// USDCDepositMetaData contains all meta data concerning the USDCDeposit contract.
var USDCDepositMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_usdc\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_treasury\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositFor\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minimumDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minimumDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumDeposit\",\"inputs\":[{\"name\":\"newMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumDepositSet\",\"inputs\":[{\"name\":\"oldMinimum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinimum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BelowMinimumDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	ID:  "USDCDeposit",
}

// USDCDeposit is an auto generated Go binding around an Ethereum contract.
type USDCDeposit struct {
	abi abi.ABI
}

// NewUSDCDeposit creates a new instance of USDCDeposit.
func NewUSDCDeposit() *USDCDeposit {
	parsed, err := USDCDepositMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &USDCDeposit{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *USDCDeposit) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _usdc, address _treasury) returns()
func (uSDCDeposit *USDCDeposit) PackConstructor(_usdc common.Address, _treasury common.Address) []byte {
	enc, err := uSDCDeposit.abi.Pack("", _usdc, _treasury)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb6b55f25.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deposit(uint256 amount) returns()
func (uSDCDeposit *USDCDeposit) PackDeposit(amount *big.Int) []byte {
	enc, err := uSDCDeposit.abi.Pack("deposit", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb6b55f25.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deposit(uint256 amount) returns()
func (uSDCDeposit *USDCDeposit) TryPackDeposit(amount *big.Int) ([]byte, error) {
	return uSDCDeposit.abi.Pack("deposit", amount)
}

// PackDepositFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36efd16f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function depositFor(uint256 amount, address account) returns()
func (uSDCDeposit *USDCDeposit) PackDepositFor(amount *big.Int, account common.Address) []byte {
	enc, err := uSDCDeposit.abi.Pack("depositFor", amount, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDepositFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36efd16f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function depositFor(uint256 amount, address account) returns()
func (uSDCDeposit *USDCDeposit) TryPackDepositFor(amount *big.Int, account common.Address) ([]byte, error) {
	return uSDCDeposit.abi.Pack("depositFor", amount, account)
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcd6dc687.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (uSDCDeposit *USDCDeposit) PackInitialize(initialOwner common.Address, minimumDeposit *big.Int) []byte {
	enc, err := uSDCDeposit.abi.Pack("initialize", initialOwner, minimumDeposit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcd6dc687.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (uSDCDeposit *USDCDeposit) TryPackInitialize(initialOwner common.Address, minimumDeposit *big.Int) ([]byte, error) {
	return uSDCDeposit.abi.Pack("initialize", initialOwner, minimumDeposit)
}

// PackMinimumDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x636bfbab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (uSDCDeposit *USDCDeposit) PackMinimumDeposit() []byte {
	enc, err := uSDCDeposit.abi.Pack("minimumDeposit")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinimumDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x636bfbab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (uSDCDeposit *USDCDeposit) TryPackMinimumDeposit() ([]byte, error) {
	return uSDCDeposit.abi.Pack("minimumDeposit")
}

// UnpackMinimumDeposit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x636bfbab.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (uSDCDeposit *USDCDeposit) UnpackMinimumDeposit(data []byte) (*big.Int, error) {
	out, err := uSDCDeposit.abi.Unpack("minimumDeposit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (uSDCDeposit *USDCDeposit) PackOwner() []byte {
	enc, err := uSDCDeposit.abi.Pack("owner")
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
func (uSDCDeposit *USDCDeposit) TryPackOwner() ([]byte, error) {
	return uSDCDeposit.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (uSDCDeposit *USDCDeposit) UnpackOwner(data []byte) (common.Address, error) {
	out, err := uSDCDeposit.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (uSDCDeposit *USDCDeposit) PackRenounceOwnership() []byte {
	enc, err := uSDCDeposit.abi.Pack("renounceOwnership")
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
func (uSDCDeposit *USDCDeposit) TryPackRenounceOwnership() ([]byte, error) {
	return uSDCDeposit.abi.Pack("renounceOwnership")
}

// PackSetMinimumDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe78ec42e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinimumDeposit(uint256 newMinimum) returns()
func (uSDCDeposit *USDCDeposit) PackSetMinimumDeposit(newMinimum *big.Int) []byte {
	enc, err := uSDCDeposit.abi.Pack("setMinimumDeposit", newMinimum)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinimumDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe78ec42e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinimumDeposit(uint256 newMinimum) returns()
func (uSDCDeposit *USDCDeposit) TryPackSetMinimumDeposit(newMinimum *big.Int) ([]byte, error) {
	return uSDCDeposit.abi.Pack("setMinimumDeposit", newMinimum)
}

// PackSweep is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01681a62.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweep(address token) returns()
func (uSDCDeposit *USDCDeposit) PackSweep(token common.Address) []byte {
	enc, err := uSDCDeposit.abi.Pack("sweep", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSweep is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01681a62.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sweep(address token) returns()
func (uSDCDeposit *USDCDeposit) TryPackSweep(token common.Address) ([]byte, error) {
	return uSDCDeposit.abi.Pack("sweep", token)
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (uSDCDeposit *USDCDeposit) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := uSDCDeposit.abi.Pack("transferOwnership", newOwner)
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
func (uSDCDeposit *USDCDeposit) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return uSDCDeposit.abi.Pack("transferOwnership", newOwner)
}

// PackTreasury is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61d027b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function treasury() view returns(address)
func (uSDCDeposit *USDCDeposit) PackTreasury() []byte {
	enc, err := uSDCDeposit.abi.Pack("treasury")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTreasury is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61d027b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function treasury() view returns(address)
func (uSDCDeposit *USDCDeposit) TryPackTreasury() ([]byte, error) {
	return uSDCDeposit.abi.Pack("treasury")
}

// UnpackTreasury is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (uSDCDeposit *USDCDeposit) UnpackTreasury(data []byte) (common.Address, error) {
	out, err := uSDCDeposit.abi.Unpack("treasury", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUsdc is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e413bee.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function usdc() view returns(address)
func (uSDCDeposit *USDCDeposit) PackUsdc() []byte {
	enc, err := uSDCDeposit.abi.Pack("usdc")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUsdc is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e413bee.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function usdc() view returns(address)
func (uSDCDeposit *USDCDeposit) TryPackUsdc() ([]byte, error) {
	return uSDCDeposit.abi.Pack("usdc")
}

// UnpackUsdc is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (uSDCDeposit *USDCDeposit) UnpackUsdc(data []byte) (common.Address, error) {
	out, err := uSDCDeposit.abi.Unpack("usdc", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// USDCDepositDeposit represents a Deposit event raised by the USDCDeposit contract.
type USDCDepositDeposit struct {
	Depositor common.Address
	Account   common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const USDCDepositDepositEventName = "Deposit"

// ContractEventName returns the user-defined event name.
func (USDCDepositDeposit) ContractEventName() string {
	return USDCDepositDepositEventName
}

// UnpackDepositEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposit(address indexed depositor, address indexed account, uint256 amount)
func (uSDCDeposit *USDCDeposit) UnpackDepositEvent(log *types.Log) (*USDCDepositDeposit, error) {
	event := "Deposit"
	if len(log.Topics) == 0 || log.Topics[0] != uSDCDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCDepositDeposit)
	if len(log.Data) > 0 {
		if err := uSDCDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDCDeposit.abi.Events[event].Inputs {
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

// USDCDepositInitialized represents a Initialized event raised by the USDCDeposit contract.
type USDCDepositInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const USDCDepositInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (USDCDepositInitialized) ContractEventName() string {
	return USDCDepositInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (uSDCDeposit *USDCDeposit) UnpackInitializedEvent(log *types.Log) (*USDCDepositInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != uSDCDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCDepositInitialized)
	if len(log.Data) > 0 {
		if err := uSDCDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDCDeposit.abi.Events[event].Inputs {
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

// USDCDepositMinimumDepositSet represents a MinimumDepositSet event raised by the USDCDeposit contract.
type USDCDepositMinimumDepositSet struct {
	OldMinimum *big.Int
	NewMinimum *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const USDCDepositMinimumDepositSetEventName = "MinimumDepositSet"

// ContractEventName returns the user-defined event name.
func (USDCDepositMinimumDepositSet) ContractEventName() string {
	return USDCDepositMinimumDepositSetEventName
}

// UnpackMinimumDepositSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum)
func (uSDCDeposit *USDCDeposit) UnpackMinimumDepositSetEvent(log *types.Log) (*USDCDepositMinimumDepositSet, error) {
	event := "MinimumDepositSet"
	if len(log.Topics) == 0 || log.Topics[0] != uSDCDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCDepositMinimumDepositSet)
	if len(log.Data) > 0 {
		if err := uSDCDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDCDeposit.abi.Events[event].Inputs {
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

// USDCDepositOwnershipTransferred represents a OwnershipTransferred event raised by the USDCDeposit contract.
type USDCDepositOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const USDCDepositOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (USDCDepositOwnershipTransferred) ContractEventName() string {
	return USDCDepositOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (uSDCDeposit *USDCDeposit) UnpackOwnershipTransferredEvent(log *types.Log) (*USDCDepositOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != uSDCDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCDepositOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := uSDCDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDCDeposit.abi.Events[event].Inputs {
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
func (uSDCDeposit *USDCDeposit) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], uSDCDeposit.abi.Errors["BelowMinimumDeposit"].ID.Bytes()[:4]) {
		return uSDCDeposit.UnpackBelowMinimumDepositError(raw[4:])
	}
	if bytes.Equal(raw[:4], uSDCDeposit.abi.Errors["ZeroAddress"].ID.Bytes()[:4]) {
		return uSDCDeposit.UnpackZeroAddressError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// USDCDepositBelowMinimumDeposit represents a BelowMinimumDeposit error raised by the USDCDeposit contract.
type USDCDepositBelowMinimumDeposit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BelowMinimumDeposit()
func USDCDepositBelowMinimumDepositErrorID() common.Hash {
	return common.HexToHash("0x5bbe86227276c69438659323dae1bd8779fe44e8baecbdf64f7c198d89e7773b")
}

// UnpackBelowMinimumDepositError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BelowMinimumDeposit()
func (uSDCDeposit *USDCDeposit) UnpackBelowMinimumDepositError(raw []byte) (*USDCDepositBelowMinimumDeposit, error) {
	out := new(USDCDepositBelowMinimumDeposit)
	if err := uSDCDeposit.abi.UnpackIntoInterface(out, "BelowMinimumDeposit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// USDCDepositZeroAddress represents a ZeroAddress error raised by the USDCDeposit contract.
type USDCDepositZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddress()
func USDCDepositZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd92e233df2717d4a40030e20904abd27b68fcbeede117eaaccbbdac9618c8c73")
}

// UnpackZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddress()
func (uSDCDeposit *USDCDeposit) UnpackZeroAddressError(raw []byte) (*USDCDepositZeroAddress, error) {
	out := new(USDCDepositZeroAddress)
	if err := uSDCDeposit.abi.UnpackIntoInterface(out, "ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}
