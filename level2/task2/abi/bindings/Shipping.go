// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ShippingMetaData contains all meta data concerning the Shipping contract.
var ShippingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"LogNewAlert\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Delivered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Shipped\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5060008060006101000a81548160ff021916908360028111156032576031603b565b5b0217905550606a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b61048f806100796000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806345f09ce914610046578063779822f7146100645780637e59301e1461006e575b600080fd5b61004e610078565b60405161005b9190610330565b60405180910390f35b61006c61009c565b005b6100766100fd565b005b606060008060009054906101000a900460ff1690506100968161015e565b91505090565b60026000806101000a81548160ff021916908360028111156100c1576100c0610352565b5b02179055507fca75c80642b93c980135b8355995abae3ab06033dc81f956b3824323b4903ef56040516100f3906103cd565b60405180910390a1565b60016000806101000a81548160ff0219169083600281111561012257610121610352565b5b02179055507fca75c80642b93c980135b8355995abae3ab06033dc81f956b3824323b4903ef560405161015490610439565b60405180910390a1565b606081600281111561017357610172610352565b5b6000600281111561018757610186610352565b5b036101c9576040518060400160405280600781526020017f50656e64696e6700000000000000000000000000000000000000000000000000815250905061029b565b8160028111156101dc576101db610352565b5b600160028111156101f0576101ef610352565b5b03610232576040518060400160405280600781526020017f5368697070656400000000000000000000000000000000000000000000000000815250905061029b565b81600281111561024557610244610352565b5b60028081111561025857610257610352565b5b0361029a576040518060400160405280600981526020017f44656c6976657265640000000000000000000000000000000000000000000000815250905061029b565b5b919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156102da5780820151818401526020810190506102bf565b60008484015250505050565b6000601f19601f8301169050919050565b6000610302826102a0565b61030c81856102ab565b935061031c8185602086016102bc565b610325816102e6565b840191505092915050565b6000602082019050818103600083015261034a81846102f7565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f596f7572207061636b6167652068617320617272697665640000000000000000600082015250565b60006103b76018836102ab565b91506103c282610381565b602082019050919050565b600060208201905081810360008301526103e6816103aa565b9050919050565b7f596f7572207061636b61676520686173206265656e2073686970706564000000600082015250565b6000610423601d836102ab565b915061042e826103ed565b602082019050919050565b6000602082019050818103600083015261045281610416565b905091905056fea2646970667358221220b45748c5e8dc8e0e9a31feb6835d5dbf3e7edafe2a4264c8fd486c3f6deea0a764736f6c634300081c0033",
}

// ShippingABI is the input ABI used to generate the binding from.
// Deprecated: Use ShippingMetaData.ABI instead.
var ShippingABI = ShippingMetaData.ABI

// ShippingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ShippingMetaData.Bin instead.
var ShippingBin = ShippingMetaData.Bin

// DeployShipping deploys a new Ethereum contract, binding an instance of Shipping to it.
func DeployShipping(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Shipping, error) {
	parsed, err := ShippingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ShippingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Shipping{ShippingCaller: ShippingCaller{contract: contract}, ShippingTransactor: ShippingTransactor{contract: contract}, ShippingFilterer: ShippingFilterer{contract: contract}}, nil
}

// Shipping is an auto generated Go binding around an Ethereum contract.
type Shipping struct {
	ShippingCaller     // Read-only binding to the contract
	ShippingTransactor // Write-only binding to the contract
	ShippingFilterer   // Log filterer for contract events
}

// ShippingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShippingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShippingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShippingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShippingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShippingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShippingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShippingSession struct {
	Contract     *Shipping         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShippingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShippingCallerSession struct {
	Contract *ShippingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ShippingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShippingTransactorSession struct {
	Contract     *ShippingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ShippingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShippingRaw struct {
	Contract *Shipping // Generic contract binding to access the raw methods on
}

// ShippingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShippingCallerRaw struct {
	Contract *ShippingCaller // Generic read-only contract binding to access the raw methods on
}

// ShippingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShippingTransactorRaw struct {
	Contract *ShippingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShipping creates a new instance of Shipping, bound to a specific deployed contract.
func NewShipping(address common.Address, backend bind.ContractBackend) (*Shipping, error) {
	contract, err := bindShipping(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Shipping{ShippingCaller: ShippingCaller{contract: contract}, ShippingTransactor: ShippingTransactor{contract: contract}, ShippingFilterer: ShippingFilterer{contract: contract}}, nil
}

// NewShippingCaller creates a new read-only instance of Shipping, bound to a specific deployed contract.
func NewShippingCaller(address common.Address, caller bind.ContractCaller) (*ShippingCaller, error) {
	contract, err := bindShipping(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShippingCaller{contract: contract}, nil
}

// NewShippingTransactor creates a new write-only instance of Shipping, bound to a specific deployed contract.
func NewShippingTransactor(address common.Address, transactor bind.ContractTransactor) (*ShippingTransactor, error) {
	contract, err := bindShipping(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShippingTransactor{contract: contract}, nil
}

// NewShippingFilterer creates a new log filterer instance of Shipping, bound to a specific deployed contract.
func NewShippingFilterer(address common.Address, filterer bind.ContractFilterer) (*ShippingFilterer, error) {
	contract, err := bindShipping(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShippingFilterer{contract: contract}, nil
}

// bindShipping binds a generic wrapper to an already deployed contract.
func bindShipping(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShippingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shipping *ShippingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shipping.Contract.ShippingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shipping *ShippingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shipping.Contract.ShippingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shipping *ShippingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shipping.Contract.ShippingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shipping *ShippingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shipping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shipping *ShippingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shipping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shipping *ShippingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shipping.Contract.contract.Transact(opts, method, params...)
}

// Status is a free data retrieval call binding the contract method 0x45f09ce9.
//
// Solidity: function Status() view returns(string)
func (_Shipping *ShippingCaller) Status(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Shipping.contract.Call(opts, &out, "Status")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x45f09ce9.
//
// Solidity: function Status() view returns(string)
func (_Shipping *ShippingSession) Status() (string, error) {
	return _Shipping.Contract.Status(&_Shipping.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x45f09ce9.
//
// Solidity: function Status() view returns(string)
func (_Shipping *ShippingCallerSession) Status() (string, error) {
	return _Shipping.Contract.Status(&_Shipping.CallOpts)
}

// Delivered is a paid mutator transaction binding the contract method 0x779822f7.
//
// Solidity: function Delivered() returns()
func (_Shipping *ShippingTransactor) Delivered(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shipping.contract.Transact(opts, "Delivered")
}

// Delivered is a paid mutator transaction binding the contract method 0x779822f7.
//
// Solidity: function Delivered() returns()
func (_Shipping *ShippingSession) Delivered() (*types.Transaction, error) {
	return _Shipping.Contract.Delivered(&_Shipping.TransactOpts)
}

// Delivered is a paid mutator transaction binding the contract method 0x779822f7.
//
// Solidity: function Delivered() returns()
func (_Shipping *ShippingTransactorSession) Delivered() (*types.Transaction, error) {
	return _Shipping.Contract.Delivered(&_Shipping.TransactOpts)
}

// Shipped is a paid mutator transaction binding the contract method 0x7e59301e.
//
// Solidity: function Shipped() returns()
func (_Shipping *ShippingTransactor) Shipped(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shipping.contract.Transact(opts, "Shipped")
}

// Shipped is a paid mutator transaction binding the contract method 0x7e59301e.
//
// Solidity: function Shipped() returns()
func (_Shipping *ShippingSession) Shipped() (*types.Transaction, error) {
	return _Shipping.Contract.Shipped(&_Shipping.TransactOpts)
}

// Shipped is a paid mutator transaction binding the contract method 0x7e59301e.
//
// Solidity: function Shipped() returns()
func (_Shipping *ShippingTransactorSession) Shipped() (*types.Transaction, error) {
	return _Shipping.Contract.Shipped(&_Shipping.TransactOpts)
}

// ShippingLogNewAlertIterator is returned from FilterLogNewAlert and is used to iterate over the raw logs and unpacked data for LogNewAlert events raised by the Shipping contract.
type ShippingLogNewAlertIterator struct {
	Event *ShippingLogNewAlert // Event containing the contract specifics and raw log

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
func (it *ShippingLogNewAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShippingLogNewAlert)
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
		it.Event = new(ShippingLogNewAlert)
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
func (it *ShippingLogNewAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShippingLogNewAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShippingLogNewAlert represents a LogNewAlert event raised by the Shipping contract.
type ShippingLogNewAlert struct {
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogNewAlert is a free log retrieval operation binding the contract event 0xca75c80642b93c980135b8355995abae3ab06033dc81f956b3824323b4903ef5.
//
// Solidity: event LogNewAlert(string description)
func (_Shipping *ShippingFilterer) FilterLogNewAlert(opts *bind.FilterOpts) (*ShippingLogNewAlertIterator, error) {

	logs, sub, err := _Shipping.contract.FilterLogs(opts, "LogNewAlert")
	if err != nil {
		return nil, err
	}
	return &ShippingLogNewAlertIterator{contract: _Shipping.contract, event: "LogNewAlert", logs: logs, sub: sub}, nil
}

// WatchLogNewAlert is a free log subscription operation binding the contract event 0xca75c80642b93c980135b8355995abae3ab06033dc81f956b3824323b4903ef5.
//
// Solidity: event LogNewAlert(string description)
func (_Shipping *ShippingFilterer) WatchLogNewAlert(opts *bind.WatchOpts, sink chan<- *ShippingLogNewAlert) (event.Subscription, error) {

	logs, sub, err := _Shipping.contract.WatchLogs(opts, "LogNewAlert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShippingLogNewAlert)
				if err := _Shipping.contract.UnpackLog(event, "LogNewAlert", log); err != nil {
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

// ParseLogNewAlert is a log parse operation binding the contract event 0xca75c80642b93c980135b8355995abae3ab06033dc81f956b3824323b4903ef5.
//
// Solidity: event LogNewAlert(string description)
func (_Shipping *ShippingFilterer) ParseLogNewAlert(log types.Log) (*ShippingLogNewAlert, error) {
	event := new(ShippingLogNewAlert)
	if err := _Shipping.contract.UnpackLog(event, "LogNewAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
