// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcchannel

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChannelCounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type ChannelCounterpartyData struct {
	PortId    string
	ChannelId string
}

// ChannelData is an auto generated low-level Go binding around an user-defined struct.
type ChannelData struct {
	State          uint8
	Ordering       uint8
	Counterparty   ChannelCounterpartyData
	ConnectionHops []string
	Version        string
}

// IBCChannelMsgChannelOpenInit is an auto generated low-level Go binding around an user-defined struct.
type IBCChannelMsgChannelOpenInit struct {
	ChannelId string
	PortId    string
	Channel   ChannelData
}

// IbcchannelABI is the input ABI used to generate the binding from.
const IbcchannelABI = "[{\"inputs\":[{\"internalType\":\"contractProvableStore\",\"name\":\"store\",\"type\":\"address\"},{\"internalType\":\"contractIBCClient\",\"name\":\"ibcclient_\",\"type\":\"address\"},{\"internalType\":\"contractIBCConnection\",\"name\":\"ibcconnection_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"}],\"internalType\":\"structIBCChannel.MsgChannelOpenInit\",\"name\":\"msg_\",\"type\":\"tuple\"}],\"name\":\"channelOpenInit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Ibcchannel is an auto generated Go binding around an Ethereum contract.
type Ibcchannel struct {
	IbcchannelCaller     // Read-only binding to the contract
	IbcchannelTransactor // Write-only binding to the contract
	IbcchannelFilterer   // Log filterer for contract events
}

// IbcchannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcchannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcchannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcchannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcchannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcchannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcchannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcchannelSession struct {
	Contract     *Ibcchannel       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcchannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcchannelCallerSession struct {
	Contract *IbcchannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IbcchannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcchannelTransactorSession struct {
	Contract     *IbcchannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IbcchannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcchannelRaw struct {
	Contract *Ibcchannel // Generic contract binding to access the raw methods on
}

// IbcchannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcchannelCallerRaw struct {
	Contract *IbcchannelCaller // Generic read-only contract binding to access the raw methods on
}

// IbcchannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcchannelTransactorRaw struct {
	Contract *IbcchannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcchannel creates a new instance of Ibcchannel, bound to a specific deployed contract.
func NewIbcchannel(address common.Address, backend bind.ContractBackend) (*Ibcchannel, error) {
	contract, err := bindIbcchannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibcchannel{IbcchannelCaller: IbcchannelCaller{contract: contract}, IbcchannelTransactor: IbcchannelTransactor{contract: contract}, IbcchannelFilterer: IbcchannelFilterer{contract: contract}}, nil
}

// NewIbcchannelCaller creates a new read-only instance of Ibcchannel, bound to a specific deployed contract.
func NewIbcchannelCaller(address common.Address, caller bind.ContractCaller) (*IbcchannelCaller, error) {
	contract, err := bindIbcchannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcchannelCaller{contract: contract}, nil
}

// NewIbcchannelTransactor creates a new write-only instance of Ibcchannel, bound to a specific deployed contract.
func NewIbcchannelTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcchannelTransactor, error) {
	contract, err := bindIbcchannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcchannelTransactor{contract: contract}, nil
}

// NewIbcchannelFilterer creates a new log filterer instance of Ibcchannel, bound to a specific deployed contract.
func NewIbcchannelFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcchannelFilterer, error) {
	contract, err := bindIbcchannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcchannelFilterer{contract: contract}, nil
}

// bindIbcchannel binds a generic wrapper to an already deployed contract.
func bindIbcchannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IbcchannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcchannel *IbcchannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcchannel.Contract.IbcchannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcchannel *IbcchannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcchannel.Contract.IbcchannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcchannel *IbcchannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcchannel.Contract.IbcchannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibcchannel *IbcchannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibcchannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibcchannel *IbcchannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibcchannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibcchannel *IbcchannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibcchannel.Contract.contract.Transact(opts, method, params...)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcchannel *IbcchannelTransactor) ChannelOpenInit(opts *bind.TransactOpts, msg_ IBCChannelMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcchannel.contract.Transact(opts, "channelOpenInit", msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcchannel *IbcchannelSession) ChannelOpenInit(msg_ IBCChannelMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenInit(&_Ibcchannel.TransactOpts, msg_)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0xd4fd4a05.
//
// Solidity: function channelOpenInit((string,string,(uint8,uint8,(string,string),string[],string)) msg_) returns(string)
func (_Ibcchannel *IbcchannelTransactorSession) ChannelOpenInit(msg_ IBCChannelMsgChannelOpenInit) (*types.Transaction, error) {
	return _Ibcchannel.Contract.ChannelOpenInit(&_Ibcchannel.TransactOpts, msg_)
}
