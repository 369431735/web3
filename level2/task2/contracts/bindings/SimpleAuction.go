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

// SimpleAuctionMetaData contains all meta data concerning the SimpleAuction contract.
var SimpleAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_biddingTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auctionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b35380380610b3583398181016040528101906100329190610124565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550814261007e9190610193565b60018190555050506101c7565b600080fd5b6000819050919050565b6100a381610090565b81146100ae57600080fd5b50565b6000815190506100c08161009a565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100f1826100c6565b9050919050565b610101816100e6565b811461010c57600080fd5b50565b60008151905061011e816100f8565b92915050565b6000806040838503121561013b5761013a61008b565b5b6000610149858286016100b1565b925050602061015a8582860161010f565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061019e82610090565b91506101a983610090565b92508282019050808211156101c1576101c0610164565b5b92915050565b61095f806101d66000396000f3fe6080604052600436106100705760003560e01c80633ccfd60b1161004e5780633ccfd60b146100c15780634b449cba146100ec57806391f9015714610117578063d57bde791461014257610070565b80631998aeef146100755780632a24f46c1461007f57806338af3eed14610096575b600080fd5b61007d61016d565b005b34801561008b57600080fd5b506100946102fe565b005b3480156100a257600080fd5b506100ab610476565b6040516100b89190610631565b60405180910390f35b3480156100cd57600080fd5b506100d661049a565b6040516100e39190610667565b60405180910390f35b3480156100f857600080fd5b506101016105be565b60405161010e919061069b565b60405180910390f35b34801561012357600080fd5b5061012c6105c4565b6040516101399190610631565b60405180910390f35b34801561014e57600080fd5b506101576105ea565b604051610164919061069b565b60405180910390f35b6001544211156101b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a990610713565b60405180910390fd5b60035434116101f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed9061077f565b60405180910390fd5b60006003541461027b5760035460046000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461027391906107ce565b925050819055505b33600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550346003819055507ff4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad30033346040516102f4929190610802565b60405180910390a1565b600154421015610343576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033a90610877565b60405180910390fd5b600560009054906101000a900460ff1615610393576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038a90610909565b60405180910390fd5b6001600560006101000a81548160ff0219169083151502179055507fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600354604051610403929190610802565b60405180910390a160008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6003549081150290604051600060405180830381858888f19350505050158015610473573d6000803e3d6000fd5b50565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008111156105b5576000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050506105b45780600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060009150506105bb565b5b60019150505b90565b60015481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061061b826105f0565b9050919050565b61062b81610610565b82525050565b60006020820190506106466000830184610622565b92915050565b60008115159050919050565b6106618161064c565b82525050565b600060208201905061067c6000830184610658565b92915050565b6000819050919050565b61069581610682565b82525050565b60006020820190506106b0600083018461068c565b92915050565b600082825260208201905092915050565b7f41756374696f6e20616c726561647920656e6465642e00000000000000000000600082015250565b60006106fd6016836106b6565b9150610708826106c7565b602082019050919050565b6000602082019050818103600083015261072c816106f0565b9050919050565b7f546865726520616c7265616479206973206120686967686572206269642e0000600082015250565b6000610769601e836106b6565b915061077482610733565b602082019050919050565b600060208201905081810360008301526107988161075c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006107d982610682565b91506107e483610682565b92508282019050808211156107fc576107fb61079f565b5b92915050565b60006040820190506108176000830185610622565b610824602083018461068c565b9392505050565b7f41756374696f6e206e6f742079657420656e6465642e00000000000000000000600082015250565b60006108616016836106b6565b915061086c8261082b565b602082019050919050565b6000602082019050818103600083015261089081610854565b9050919050565b7f61756374696f6e456e642068617320616c7265616479206265656e2063616c6c60008201527f65642e0000000000000000000000000000000000000000000000000000000000602082015250565b60006108f36023836106b6565b91506108fe82610897565b604082019050919050565b60006020820190508181036000830152610922816108e6565b905091905056fea26469706673582212204a7a90dab0fc572f92733ddd36d21ace895ee9235b9d54a5c20ca355e7460cf764736f6c634300081c0033",
}

// SimpleAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleAuctionMetaData.ABI instead.
var SimpleAuctionABI = SimpleAuctionMetaData.ABI

// SimpleAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleAuctionMetaData.Bin instead.
var SimpleAuctionBin = SimpleAuctionMetaData.Bin

// DeploySimpleAuction deploys a new Ethereum contract, binding an instance of SimpleAuction to it.
func DeploySimpleAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _biddingTime *big.Int, _beneficiary common.Address) (common.Address, *types.Transaction, *SimpleAuction, error) {
	parsed, err := SimpleAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleAuctionBin), backend, _biddingTime, _beneficiary)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleAuction{SimpleAuctionCaller: SimpleAuctionCaller{contract: contract}, SimpleAuctionTransactor: SimpleAuctionTransactor{contract: contract}, SimpleAuctionFilterer: SimpleAuctionFilterer{contract: contract}}, nil
}

// SimpleAuction is an auto generated Go binding around an Ethereum contract.
type SimpleAuction struct {
	SimpleAuctionCaller     // Read-only binding to the contract
	SimpleAuctionTransactor // Write-only binding to the contract
	SimpleAuctionFilterer   // Log filterer for contract events
}

// SimpleAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleAuctionSession struct {
	Contract     *SimpleAuction    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleAuctionCallerSession struct {
	Contract *SimpleAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleAuctionTransactorSession struct {
	Contract     *SimpleAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleAuctionRaw struct {
	Contract *SimpleAuction // Generic contract binding to access the raw methods on
}

// SimpleAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleAuctionCallerRaw struct {
	Contract *SimpleAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleAuctionTransactorRaw struct {
	Contract *SimpleAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleAuction creates a new instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuction(address common.Address, backend bind.ContractBackend) (*SimpleAuction, error) {
	contract, err := bindSimpleAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleAuction{SimpleAuctionCaller: SimpleAuctionCaller{contract: contract}, SimpleAuctionTransactor: SimpleAuctionTransactor{contract: contract}, SimpleAuctionFilterer: SimpleAuctionFilterer{contract: contract}}, nil
}

// NewSimpleAuctionCaller creates a new read-only instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionCaller(address common.Address, caller bind.ContractCaller) (*SimpleAuctionCaller, error) {
	contract, err := bindSimpleAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionCaller{contract: contract}, nil
}

// NewSimpleAuctionTransactor creates a new write-only instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleAuctionTransactor, error) {
	contract, err := bindSimpleAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionTransactor{contract: contract}, nil
}

// NewSimpleAuctionFilterer creates a new log filterer instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleAuctionFilterer, error) {
	contract, err := bindSimpleAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionFilterer{contract: contract}, nil
}

// bindSimpleAuction binds a generic wrapper to an already deployed contract.
func bindSimpleAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAuction *SimpleAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAuction.Contract.SimpleAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAuction *SimpleAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.Contract.SimpleAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAuction *SimpleAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAuction.Contract.SimpleAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAuction *SimpleAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAuction *SimpleAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAuction *SimpleAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAuction.Contract.contract.Transact(opts, method, params...)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SimpleAuction *SimpleAuctionCaller) AuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleAuction.contract.Call(opts, &out, "auctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SimpleAuction *SimpleAuctionSession) AuctionEndTime() (*big.Int, error) {
	return _SimpleAuction.Contract.AuctionEndTime(&_SimpleAuction.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SimpleAuction *SimpleAuctionCallerSession) AuctionEndTime() (*big.Int, error) {
	return _SimpleAuction.Contract.AuctionEndTime(&_SimpleAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SimpleAuction *SimpleAuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAuction.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SimpleAuction *SimpleAuctionSession) Beneficiary() (common.Address, error) {
	return _SimpleAuction.Contract.Beneficiary(&_SimpleAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_SimpleAuction *SimpleAuctionCallerSession) Beneficiary() (common.Address, error) {
	return _SimpleAuction.Contract.Beneficiary(&_SimpleAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SimpleAuction *SimpleAuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleAuction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SimpleAuction *SimpleAuctionSession) HighestBid() (*big.Int, error) {
	return _SimpleAuction.Contract.HighestBid(&_SimpleAuction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_SimpleAuction *SimpleAuctionCallerSession) HighestBid() (*big.Int, error) {
	return _SimpleAuction.Contract.HighestBid(&_SimpleAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SimpleAuction *SimpleAuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAuction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SimpleAuction *SimpleAuctionSession) HighestBidder() (common.Address, error) {
	return _SimpleAuction.Contract.HighestBidder(&_SimpleAuction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_SimpleAuction *SimpleAuctionCallerSession) HighestBidder() (common.Address, error) {
	return _SimpleAuction.Contract.HighestBidder(&_SimpleAuction.CallOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_SimpleAuction *SimpleAuctionTransactor) AuctionEnd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.contract.Transact(opts, "auctionEnd")
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_SimpleAuction *SimpleAuctionSession) AuctionEnd() (*types.Transaction, error) {
	return _SimpleAuction.Contract.AuctionEnd(&_SimpleAuction.TransactOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_SimpleAuction *SimpleAuctionTransactorSession) AuctionEnd() (*types.Transaction, error) {
	return _SimpleAuction.Contract.AuctionEnd(&_SimpleAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SimpleAuction *SimpleAuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SimpleAuction *SimpleAuctionSession) Bid() (*types.Transaction, error) {
	return _SimpleAuction.Contract.Bid(&_SimpleAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_SimpleAuction *SimpleAuctionTransactorSession) Bid() (*types.Transaction, error) {
	return _SimpleAuction.Contract.Bid(&_SimpleAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SimpleAuction *SimpleAuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SimpleAuction *SimpleAuctionSession) Withdraw() (*types.Transaction, error) {
	return _SimpleAuction.Contract.Withdraw(&_SimpleAuction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_SimpleAuction *SimpleAuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _SimpleAuction.Contract.Withdraw(&_SimpleAuction.TransactOpts)
}

// SimpleAuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the SimpleAuction contract.
type SimpleAuctionAuctionEndedIterator struct {
	Event *SimpleAuctionAuctionEnded // Event containing the contract specifics and raw log

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
func (it *SimpleAuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleAuctionAuctionEnded)
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
		it.Event = new(SimpleAuctionAuctionEnded)
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
func (it *SimpleAuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleAuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleAuctionAuctionEnded represents a AuctionEnded event raised by the SimpleAuction contract.
type SimpleAuctionAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*SimpleAuctionAuctionEndedIterator, error) {

	logs, sub, err := _SimpleAuction.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionAuctionEndedIterator{contract: _SimpleAuction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *SimpleAuctionAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _SimpleAuction.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleAuctionAuctionEnded)
				if err := _SimpleAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) ParseAuctionEnded(log types.Log) (*SimpleAuctionAuctionEnded, error) {
	event := new(SimpleAuctionAuctionEnded)
	if err := _SimpleAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the SimpleAuction contract.
type SimpleAuctionHighestBidIncreasedIterator struct {
	Event *SimpleAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *SimpleAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleAuctionHighestBidIncreased)
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
		it.Event = new(SimpleAuctionHighestBidIncreased)
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
func (it *SimpleAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the SimpleAuction contract.
type SimpleAuctionHighestBidIncreased struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*SimpleAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _SimpleAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionHighestBidIncreasedIterator{contract: _SimpleAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *SimpleAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _SimpleAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleAuctionHighestBidIncreased)
				if err := _SimpleAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0xf4757a49b326036464bec6fe419a4ae38c8a02ce3e68bf0809674f6aab8ad300.
//
// Solidity: event HighestBidIncreased(address bidder, uint256 amount)
func (_SimpleAuction *SimpleAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*SimpleAuctionHighestBidIncreased, error) {
	event := new(SimpleAuctionHighestBidIncreased)
	if err := _SimpleAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
