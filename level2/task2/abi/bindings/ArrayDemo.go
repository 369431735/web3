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

// ArrayDemoMetaData contains all meta data concerning the ArrayDemo contract.
var ArrayDemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"array\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArray\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getElement\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"i\",\"type\":\"int256\"}],\"name\":\"put\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506105178061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806338d941931461005c5780633a7d22bc1461008c578063be1c766b146100bc578063d504ea1d146100da578063f84c854a146100f8575b600080fd5b61007660048036038101906100719190610271565b610114565b60405161008391906102b7565b60405180910390f35b6100a660048036038101906100a19190610271565b610138565b6040516100b391906102b7565b60405180910390f35b6100c46101a6565b6040516100d191906102e1565b60405180910390f35b6100e26101b2565b6040516100ef91906103ba565b60405180910390f35b610112600480360381019061010d9190610408565b61020a565b005b6000818154811061012457600080fd5b906000526020600020016000915090505481565b600080805490508210610180576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017790610492565b60405180910390fd5b60008281548110610194576101936104b2565b5b90600052602060002001549050919050565b60008080549050905090565b6060600080548060200260200160405190810160405280929190818152602001828054801561020057602002820191906000526020600020905b8154815260200190600101908083116101ec575b5050505050905090565b600081908060018154018082558091505060019003906000526020600020016000909190919091505550565b600080fd5b6000819050919050565b61024e8161023b565b811461025957600080fd5b50565b60008135905061026b81610245565b92915050565b60006020828403121561028757610286610236565b5b60006102958482850161025c565b91505092915050565b6000819050919050565b6102b18161029e565b82525050565b60006020820190506102cc60008301846102a8565b92915050565b6102db8161023b565b82525050565b60006020820190506102f660008301846102d2565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6103318161029e565b82525050565b60006103438383610328565b60208301905092915050565b6000602082019050919050565b6000610367826102fc565b6103718185610307565b935061037c83610318565b8060005b838110156103ad5781516103948882610337565b975061039f8361034f565b925050600181019050610380565b5085935050505092915050565b600060208201905081810360008301526103d4818461035c565b905092915050565b6103e58161029e565b81146103f057600080fd5b50565b600081359050610402816103dc565b92915050565b60006020828403121561041e5761041d610236565b5b600061042c848285016103f3565b91505092915050565b600082825260208201905092915050565b7f496e646578206f7574206f6620626f756e647300000000000000000000000000600082015250565b600061047c601383610435565b915061048782610446565b602082019050919050565b600060208201905081810360008301526104ab8161046f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea26469706673582212204617867e42aa078a2ee94ddec5a1d82666f66be605055c1a569c8acbe7193b4a64736f6c634300081c0033",
}

// ArrayDemoABI is the input ABI used to generate the binding from.
// Deprecated: Use ArrayDemoMetaData.ABI instead.
var ArrayDemoABI = ArrayDemoMetaData.ABI

// ArrayDemoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArrayDemoMetaData.Bin instead.
var ArrayDemoBin = ArrayDemoMetaData.Bin

// DeployArrayDemo deploys a new Ethereum contract, binding an instance of ArrayDemo to it.
func DeployArrayDemo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArrayDemo, error) {
	parsed, err := ArrayDemoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArrayDemoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArrayDemo{ArrayDemoCaller: ArrayDemoCaller{contract: contract}, ArrayDemoTransactor: ArrayDemoTransactor{contract: contract}, ArrayDemoFilterer: ArrayDemoFilterer{contract: contract}}, nil
}

// ArrayDemo is an auto generated Go binding around an Ethereum contract.
type ArrayDemo struct {
	ArrayDemoCaller     // Read-only binding to the contract
	ArrayDemoTransactor // Write-only binding to the contract
	ArrayDemoFilterer   // Log filterer for contract events
}

// ArrayDemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArrayDemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayDemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArrayDemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayDemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArrayDemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayDemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArrayDemoSession struct {
	Contract     *ArrayDemo        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayDemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArrayDemoCallerSession struct {
	Contract *ArrayDemoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArrayDemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArrayDemoTransactorSession struct {
	Contract     *ArrayDemoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArrayDemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArrayDemoRaw struct {
	Contract *ArrayDemo // Generic contract binding to access the raw methods on
}

// ArrayDemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArrayDemoCallerRaw struct {
	Contract *ArrayDemoCaller // Generic read-only contract binding to access the raw methods on
}

// ArrayDemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArrayDemoTransactorRaw struct {
	Contract *ArrayDemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArrayDemo creates a new instance of ArrayDemo, bound to a specific deployed contract.
func NewArrayDemo(address common.Address, backend bind.ContractBackend) (*ArrayDemo, error) {
	contract, err := bindArrayDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArrayDemo{ArrayDemoCaller: ArrayDemoCaller{contract: contract}, ArrayDemoTransactor: ArrayDemoTransactor{contract: contract}, ArrayDemoFilterer: ArrayDemoFilterer{contract: contract}}, nil
}

// NewArrayDemoCaller creates a new read-only instance of ArrayDemo, bound to a specific deployed contract.
func NewArrayDemoCaller(address common.Address, caller bind.ContractCaller) (*ArrayDemoCaller, error) {
	contract, err := bindArrayDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayDemoCaller{contract: contract}, nil
}

// NewArrayDemoTransactor creates a new write-only instance of ArrayDemo, bound to a specific deployed contract.
func NewArrayDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*ArrayDemoTransactor, error) {
	contract, err := bindArrayDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayDemoTransactor{contract: contract}, nil
}

// NewArrayDemoFilterer creates a new log filterer instance of ArrayDemo, bound to a specific deployed contract.
func NewArrayDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*ArrayDemoFilterer, error) {
	contract, err := bindArrayDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArrayDemoFilterer{contract: contract}, nil
}

// bindArrayDemo binds a generic wrapper to an already deployed contract.
func bindArrayDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArrayDemoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArrayDemo *ArrayDemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArrayDemo.Contract.ArrayDemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArrayDemo *ArrayDemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayDemo.Contract.ArrayDemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArrayDemo *ArrayDemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArrayDemo.Contract.ArrayDemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArrayDemo *ArrayDemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArrayDemo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArrayDemo *ArrayDemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayDemo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArrayDemo *ArrayDemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArrayDemo.Contract.contract.Transact(opts, method, params...)
}

// Array is a free data retrieval call binding the contract method 0x38d94193.
//
// Solidity: function array(uint256 ) view returns(int256)
func (_ArrayDemo *ArrayDemoCaller) Array(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArrayDemo.contract.Call(opts, &out, "array", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Array is a free data retrieval call binding the contract method 0x38d94193.
//
// Solidity: function array(uint256 ) view returns(int256)
func (_ArrayDemo *ArrayDemoSession) Array(arg0 *big.Int) (*big.Int, error) {
	return _ArrayDemo.Contract.Array(&_ArrayDemo.CallOpts, arg0)
}

// Array is a free data retrieval call binding the contract method 0x38d94193.
//
// Solidity: function array(uint256 ) view returns(int256)
func (_ArrayDemo *ArrayDemoCallerSession) Array(arg0 *big.Int) (*big.Int, error) {
	return _ArrayDemo.Contract.Array(&_ArrayDemo.CallOpts, arg0)
}

// GetArray is a free data retrieval call binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() view returns(int256[])
func (_ArrayDemo *ArrayDemoCaller) GetArray(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _ArrayDemo.contract.Call(opts, &out, "getArray")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetArray is a free data retrieval call binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() view returns(int256[])
func (_ArrayDemo *ArrayDemoSession) GetArray() ([]*big.Int, error) {
	return _ArrayDemo.Contract.GetArray(&_ArrayDemo.CallOpts)
}

// GetArray is a free data retrieval call binding the contract method 0xd504ea1d.
//
// Solidity: function getArray() view returns(int256[])
func (_ArrayDemo *ArrayDemoCallerSession) GetArray() ([]*big.Int, error) {
	return _ArrayDemo.Contract.GetArray(&_ArrayDemo.CallOpts)
}

// GetElement is a free data retrieval call binding the contract method 0x3a7d22bc.
//
// Solidity: function getElement(uint256 index) view returns(int256)
func (_ArrayDemo *ArrayDemoCaller) GetElement(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArrayDemo.contract.Call(opts, &out, "getElement", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetElement is a free data retrieval call binding the contract method 0x3a7d22bc.
//
// Solidity: function getElement(uint256 index) view returns(int256)
func (_ArrayDemo *ArrayDemoSession) GetElement(index *big.Int) (*big.Int, error) {
	return _ArrayDemo.Contract.GetElement(&_ArrayDemo.CallOpts, index)
}

// GetElement is a free data retrieval call binding the contract method 0x3a7d22bc.
//
// Solidity: function getElement(uint256 index) view returns(int256)
func (_ArrayDemo *ArrayDemoCallerSession) GetElement(index *big.Int) (*big.Int, error) {
	return _ArrayDemo.Contract.GetElement(&_ArrayDemo.CallOpts, index)
}

// GetLength is a free data retrieval call binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() view returns(uint256)
func (_ArrayDemo *ArrayDemoCaller) GetLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArrayDemo.contract.Call(opts, &out, "getLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLength is a free data retrieval call binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() view returns(uint256)
func (_ArrayDemo *ArrayDemoSession) GetLength() (*big.Int, error) {
	return _ArrayDemo.Contract.GetLength(&_ArrayDemo.CallOpts)
}

// GetLength is a free data retrieval call binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() view returns(uint256)
func (_ArrayDemo *ArrayDemoCallerSession) GetLength() (*big.Int, error) {
	return _ArrayDemo.Contract.GetLength(&_ArrayDemo.CallOpts)
}

// Put is a paid mutator transaction binding the contract method 0xf84c854a.
//
// Solidity: function put(int256 i) returns()
func (_ArrayDemo *ArrayDemoTransactor) Put(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _ArrayDemo.contract.Transact(opts, "put", i)
}

// Put is a paid mutator transaction binding the contract method 0xf84c854a.
//
// Solidity: function put(int256 i) returns()
func (_ArrayDemo *ArrayDemoSession) Put(i *big.Int) (*types.Transaction, error) {
	return _ArrayDemo.Contract.Put(&_ArrayDemo.TransactOpts, i)
}

// Put is a paid mutator transaction binding the contract method 0xf84c854a.
//
// Solidity: function put(int256 i) returns()
func (_ArrayDemo *ArrayDemoTransactorSession) Put(i *big.Int) (*types.Transaction, error) {
	return _ArrayDemo.Contract.Put(&_ArrayDemo.TransactOpts, i)
}
