// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package USDCCredits

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

// USDCCreditsMetaData contains all meta data concerning the USDCCredits contract.
var USDCCreditsMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_usdc\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_treasury\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minimumDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minimumDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"purchaseCredits\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"purchaseCreditsFor\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumDepositFor\",\"inputs\":[{\"name\":\"newMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumDepositSet\",\"inputs\":[{\"name\":\"oldMinimum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinimum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BelowMinimumDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	ID:  "USDCCredits",
}

// USDCCredits is an auto generated Go binding around an Ethereum contract.
type USDCCredits struct {
	abi abi.ABI
}

// NewUSDCCredits creates a new instance of USDCCredits.
func NewUSDCCredits() *USDCCredits {
	parsed, err := USDCCreditsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &USDCCredits{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *USDCCredits) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _usdc, address _treasury) returns()
func (uSDCCredits *USDCCredits) PackConstructor(_usdc common.Address, _treasury common.Address) []byte {
	enc, err := uSDCCredits.abi.Pack("", _usdc, _treasury)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcd6dc687.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address initialOwner, uint256 _minimumDeposit) returns()
func (uSDCCredits *USDCCredits) PackInitialize(initialOwner common.Address, minimumDeposit *big.Int) []byte {
	enc, err := uSDCCredits.abi.Pack("initialize", initialOwner, minimumDeposit)
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
func (uSDCCredits *USDCCredits) TryPackInitialize(initialOwner common.Address, minimumDeposit *big.Int) ([]byte, error) {
	return uSDCCredits.abi.Pack("initialize", initialOwner, minimumDeposit)
}

// PackMinimumDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x636bfbab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (uSDCCredits *USDCCredits) PackMinimumDeposit() []byte {
	enc, err := uSDCCredits.abi.Pack("minimumDeposit")
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
func (uSDCCredits *USDCCredits) TryPackMinimumDeposit() ([]byte, error) {
	return uSDCCredits.abi.Pack("minimumDeposit")
}

// UnpackMinimumDeposit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x636bfbab.
//
// Solidity: function minimumDeposit() view returns(uint256)
func (uSDCCredits *USDCCredits) UnpackMinimumDeposit(data []byte) (*big.Int, error) {
	out, err := uSDCCredits.abi.Unpack("minimumDeposit", data)
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
func (uSDCCredits *USDCCredits) PackOwner() []byte {
	enc, err := uSDCCredits.abi.Pack("owner")
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
func (uSDCCredits *USDCCredits) TryPackOwner() ([]byte, error) {
	return uSDCCredits.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (uSDCCredits *USDCCredits) UnpackOwner(data []byte) (common.Address, error) {
	out, err := uSDCCredits.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPurchaseCredits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbef101fb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function purchaseCredits(uint256 amount) returns()
func (uSDCCredits *USDCCredits) PackPurchaseCredits(amount *big.Int) []byte {
	enc, err := uSDCCredits.abi.Pack("purchaseCredits", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPurchaseCredits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbef101fb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function purchaseCredits(uint256 amount) returns()
func (uSDCCredits *USDCCredits) TryPackPurchaseCredits(amount *big.Int) ([]byte, error) {
	return uSDCCredits.abi.Pack("purchaseCredits", amount)
}

// PackPurchaseCreditsFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfeddbc3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function purchaseCreditsFor(uint256 amount, address account) returns()
func (uSDCCredits *USDCCredits) PackPurchaseCreditsFor(amount *big.Int, account common.Address) []byte {
	enc, err := uSDCCredits.abi.Pack("purchaseCreditsFor", amount, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPurchaseCreditsFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdfeddbc3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function purchaseCreditsFor(uint256 amount, address account) returns()
func (uSDCCredits *USDCCredits) TryPackPurchaseCreditsFor(amount *big.Int, account common.Address) ([]byte, error) {
	return uSDCCredits.abi.Pack("purchaseCreditsFor", amount, account)
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (uSDCCredits *USDCCredits) PackRenounceOwnership() []byte {
	enc, err := uSDCCredits.abi.Pack("renounceOwnership")
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
func (uSDCCredits *USDCCredits) TryPackRenounceOwnership() ([]byte, error) {
	return uSDCCredits.abi.Pack("renounceOwnership")
}

// PackSetMinimumDepositFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7110a9f6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinimumDepositFor(uint256 newMinimum) returns()
func (uSDCCredits *USDCCredits) PackSetMinimumDepositFor(newMinimum *big.Int) []byte {
	enc, err := uSDCCredits.abi.Pack("setMinimumDepositFor", newMinimum)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinimumDepositFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7110a9f6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinimumDepositFor(uint256 newMinimum) returns()
func (uSDCCredits *USDCCredits) TryPackSetMinimumDepositFor(newMinimum *big.Int) ([]byte, error) {
	return uSDCCredits.abi.Pack("setMinimumDepositFor", newMinimum)
}

// PackSweep is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01681a62.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweep(address token) returns()
func (uSDCCredits *USDCCredits) PackSweep(token common.Address) []byte {
	enc, err := uSDCCredits.abi.Pack("sweep", token)
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
func (uSDCCredits *USDCCredits) TryPackSweep(token common.Address) ([]byte, error) {
	return uSDCCredits.abi.Pack("sweep", token)
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (uSDCCredits *USDCCredits) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := uSDCCredits.abi.Pack("transferOwnership", newOwner)
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
func (uSDCCredits *USDCCredits) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return uSDCCredits.abi.Pack("transferOwnership", newOwner)
}

// PackTreasury is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61d027b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function treasury() view returns(address)
func (uSDCCredits *USDCCredits) PackTreasury() []byte {
	enc, err := uSDCCredits.abi.Pack("treasury")
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
func (uSDCCredits *USDCCredits) TryPackTreasury() ([]byte, error) {
	return uSDCCredits.abi.Pack("treasury")
}

// UnpackTreasury is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (uSDCCredits *USDCCredits) UnpackTreasury(data []byte) (common.Address, error) {
	out, err := uSDCCredits.abi.Unpack("treasury", data)
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
func (uSDCCredits *USDCCredits) PackUsdc() []byte {
	enc, err := uSDCCredits.abi.Pack("usdc")
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
func (uSDCCredits *USDCCredits) TryPackUsdc() ([]byte, error) {
	return uSDCCredits.abi.Pack("usdc")
}

// UnpackUsdc is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (uSDCCredits *USDCCredits) UnpackUsdc(data []byte) (common.Address, error) {
	out, err := uSDCCredits.abi.Unpack("usdc", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// USDCCreditsDeposit represents a Deposit event raised by the USDCCredits contract.
type USDCCreditsDeposit struct {
	Depositor common.Address
	Account   common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const USDCCreditsDepositEventName = "Deposit"

// ContractEventName returns the user-defined event name.
func (USDCCreditsDeposit) ContractEventName() string {
	return USDCCreditsDepositEventName
}

// UnpackDepositEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposit(address indexed depositor, address indexed account, uint256 amount)
func (uSDCCredits *USDCCredits) UnpackDepositEvent(log *types.Log) (*USDCCreditsDeposit, error) {
	event := "Deposit"
	if len(log.Topics) == 0 || log.Topics[0] != uSDCCredits.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCCreditsDeposit)
	if len(log.Data) > 0 {
		if err := uSDCCredits.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDCCredits.abi.Events[event].Inputs {
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

// USDCCreditsInitialized represents a Initialized event raised by the USDCCredits contract.
type USDCCreditsInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const USDCCreditsInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (USDCCreditsInitialized) ContractEventName() string {
	return USDCCreditsInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (uSDCCredits *USDCCredits) UnpackInitializedEvent(log *types.Log) (*USDCCreditsInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != uSDCCredits.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCCreditsInitialized)
	if len(log.Data) > 0 {
		if err := uSDCCredits.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDCCredits.abi.Events[event].Inputs {
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

// USDCCreditsMinimumDepositSet represents a MinimumDepositSet event raised by the USDCCredits contract.
type USDCCreditsMinimumDepositSet struct {
	OldMinimum *big.Int
	NewMinimum *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const USDCCreditsMinimumDepositSetEventName = "MinimumDepositSet"

// ContractEventName returns the user-defined event name.
func (USDCCreditsMinimumDepositSet) ContractEventName() string {
	return USDCCreditsMinimumDepositSetEventName
}

// UnpackMinimumDepositSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinimumDepositSet(uint256 oldMinimum, uint256 newMinimum)
func (uSDCCredits *USDCCredits) UnpackMinimumDepositSetEvent(log *types.Log) (*USDCCreditsMinimumDepositSet, error) {
	event := "MinimumDepositSet"
	if len(log.Topics) == 0 || log.Topics[0] != uSDCCredits.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCCreditsMinimumDepositSet)
	if len(log.Data) > 0 {
		if err := uSDCCredits.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDCCredits.abi.Events[event].Inputs {
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

// USDCCreditsOwnershipTransferred represents a OwnershipTransferred event raised by the USDCCredits contract.
type USDCCreditsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const USDCCreditsOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (USDCCreditsOwnershipTransferred) ContractEventName() string {
	return USDCCreditsOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (uSDCCredits *USDCCredits) UnpackOwnershipTransferredEvent(log *types.Log) (*USDCCreditsOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if len(log.Topics) == 0 || log.Topics[0] != uSDCCredits.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(USDCCreditsOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := uSDCCredits.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range uSDCCredits.abi.Events[event].Inputs {
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
func (uSDCCredits *USDCCredits) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], uSDCCredits.abi.Errors["BelowMinimumDeposit"].ID.Bytes()[:4]) {
		return uSDCCredits.UnpackBelowMinimumDepositError(raw[4:])
	}
	if bytes.Equal(raw[:4], uSDCCredits.abi.Errors["ZeroAddress"].ID.Bytes()[:4]) {
		return uSDCCredits.UnpackZeroAddressError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// USDCCreditsBelowMinimumDeposit represents a BelowMinimumDeposit error raised by the USDCCredits contract.
type USDCCreditsBelowMinimumDeposit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BelowMinimumDeposit()
func USDCCreditsBelowMinimumDepositErrorID() common.Hash {
	return common.HexToHash("0x5bbe86227276c69438659323dae1bd8779fe44e8baecbdf64f7c198d89e7773b")
}

// UnpackBelowMinimumDepositError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BelowMinimumDeposit()
func (uSDCCredits *USDCCredits) UnpackBelowMinimumDepositError(raw []byte) (*USDCCreditsBelowMinimumDeposit, error) {
	out := new(USDCCreditsBelowMinimumDeposit)
	if err := uSDCCredits.abi.UnpackIntoInterface(out, "BelowMinimumDeposit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// USDCCreditsZeroAddress represents a ZeroAddress error raised by the USDCCredits contract.
type USDCCreditsZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddress()
func USDCCreditsZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd92e233df2717d4a40030e20904abd27b68fcbeede117eaaccbbdac9618c8c73")
}

// UnpackZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddress()
func (uSDCCredits *USDCCredits) UnpackZeroAddressError(raw []byte) (*USDCCreditsZeroAddress, error) {
	out := new(USDCCreditsZeroAddress)
	if err := uSDCCredits.abi.UnpackIntoInterface(out, "ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}
