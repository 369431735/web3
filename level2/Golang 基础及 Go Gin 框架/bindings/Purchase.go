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

// PurchaseMetaData contains all meta data concerning the Purchase contract.
var PurchaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ItemReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PurchaseConfirmed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPurchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumPurchase.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405233600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002346100529190610140565b6000819055503460005460026100689190610171565b146100a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161009f90610210565b60405180910390fd5b6000600260146101000a81548160ff021916908360028111156100ce576100cd610230565b5b021790555061025f565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061014b826100d8565b9150610156836100d8565b925082610166576101656100e2565b5b828204905092915050565b600061017c826100d8565b9150610187836100d8565b9250828202610195816100d8565b915082820484148315176101ac576101ab610111565b5b5092915050565b600082825260208201905092915050565b7f56616c75652068617320746f206265206576656e2e0000000000000000000000600082015250565b60006101fa6015836101b3565b9150610205826101c4565b602082019050919050565b60006020820190508181036000830152610229816101ed565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b610c508061026e6000396000f3fe60806040526004361061007b5760003560e01c80637150d8ae1161004e5780637150d8ae1461011857806373fac6f014610143578063c19d93fb1461015a578063d6960697146101855761007b565b806308551a531461008057806312065fe0146100ab57806335a063b4146100d65780633fa4f245146100ed575b600080fd5b34801561008c57600080fd5b5061009561018f565b6040516100a29190610873565b60405180910390f35b3480156100b757600080fd5b506100c06101b5565b6040516100cd91906108a7565b60405180910390f35b3480156100e257600080fd5b506100eb6101bd565b005b3480156100f957600080fd5b506101026103b3565b60405161010f91906108a7565b60405180910390f35b34801561012457600080fd5b5061012d6103b9565b60405161013a9190610873565b60405180910390f35b34801561014f57600080fd5b506101586103df565b005b34801561016657600080fd5b5061016f61068c565b60405161017c9190610939565b60405180910390f35b61018d61069f565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600047905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461024d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610244906109b1565b60405180910390fd5b6000806002811115610262576102616108c2565b5b600260149054906101000a900460ff166002811115610284576102836108c2565b5b146102c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102bb90610a1d565b60405180910390fd5b60028060146101000a81548160ff021916908360028111156102e9576102e86108c2565b5b0217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610356573d6000803e3d6000fd5b507f13c3922f06c44c3cac6a2c721f8ec8db793360b0180d5e560ca2028de2037eaa600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516103a89190610873565b60405180910390a150565b60005481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461046f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046690610a89565b60405180910390fd5b6001806002811115610484576104836108c2565b5b600260149054906101000a900460ff1660028111156104a6576104a56108c2565b5b146104e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dd90610a1d565b60405180910390fd5b60028060146101000a81548160ff0219169083600281111561050b5761050a6108c2565b5b021790555060008054905060008160005460026105289190610ad8565b6105329190610b1a565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f1935050505015801561059c573d6000803e3d6000fd5b50600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610605573d6000803e3d6000fd5b507f977d335cc268af836f9e8542f5c430c79e387ceb4a005d8aa05350b655fd65d3600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660005460405161067f93929190610b4e565b60405180910390a1505050565b600260149054906101000a900460ff1681565b60008060028111156106b4576106b36108c2565b5b600260149054906101000a900460ff1660028111156106d6576106d56108c2565b5b14610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070d90610a1d565b60405180910390fd5b60005460026107259190610ad8565b341480610767576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075e90610bd1565b60405180910390fd5b33600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600260146101000a81548160ff021916908360028111156107ce576107cd6108c2565b5b02179055507fa29484a0d47153a96f00a1369c6885bbc89e95e35273ee97c943c440f7c4c282600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1634604051610826929190610bf1565b60405180910390a15050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061085d82610832565b9050919050565b61086d81610852565b82525050565b60006020820190506108886000830184610864565b92915050565b6000819050919050565b6108a18161088e565b82525050565b60006020820190506108bc6000830184610898565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110610902576109016108c2565b5b50565b6000819050610913826108f1565b919050565b600061092382610905565b9050919050565b61093381610918565b82525050565b600060208201905061094e600083018461092a565b92915050565b600082825260208201905092915050565b7f4f6e6c792073656c6c65722063616e2063616c6c20746869732e000000000000600082015250565b600061099b601a83610954565b91506109a682610965565b602082019050919050565b600060208201905081810360008301526109ca8161098e565b9050919050565b7f496e76616c69642073746174652e000000000000000000000000000000000000600082015250565b6000610a07600e83610954565b9150610a12826109d1565b602082019050919050565b60006020820190508181036000830152610a36816109fa565b9050919050565b7f4f6e6c792062757965722063616e2063616c6c20746869732e00000000000000600082015250565b6000610a73601983610954565b9150610a7e82610a3d565b602082019050919050565b60006020820190508181036000830152610aa281610a66565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ae38261088e565b9150610aee8361088e565b9250828202610afc8161088e565b91508282048414831517610b1357610b12610aa9565b5b5092915050565b6000610b258261088e565b9150610b308361088e565b9250828203905081811115610b4857610b47610aa9565b5b92915050565b6000606082019050610b636000830186610864565b610b706020830185610864565b610b7d6040830184610898565b949350505050565b7f436f6e646974696f6e206e6f74206d6574000000000000000000000000000000600082015250565b6000610bbb601183610954565b9150610bc682610b85565b602082019050919050565b60006020820190508181036000830152610bea81610bae565b9050919050565b6000604082019050610c066000830185610864565b610c136020830184610898565b939250505056fea264697066735822122012b00e0975e7145ecfecc26a3fbba6a8426f4982331da8c9d18228f326620e3c64736f6c634300081c0033",
}

// PurchaseABI is the input ABI used to generate the binding from.
// Deprecated: Use PurchaseMetaData.ABI instead.
var PurchaseABI = PurchaseMetaData.ABI

// PurchaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PurchaseMetaData.Bin instead.
var PurchaseBin = PurchaseMetaData.Bin

// DeployPurchase deploys a new Ethereum contract, binding an instance of Purchase to it.
func DeployPurchase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Purchase, error) {
	parsed, err := PurchaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PurchaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Purchase{PurchaseCaller: PurchaseCaller{contract: contract}, PurchaseTransactor: PurchaseTransactor{contract: contract}, PurchaseFilterer: PurchaseFilterer{contract: contract}}, nil
}

// Purchase is an auto generated Go binding around an Ethereum contract.
type Purchase struct {
	PurchaseCaller     // Read-only binding to the contract
	PurchaseTransactor // Write-only binding to the contract
	PurchaseFilterer   // Log filterer for contract events
}

// PurchaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type PurchaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PurchaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PurchaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurchaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PurchaseSession struct {
	Contract     *Purchase         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PurchaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PurchaseCallerSession struct {
	Contract *PurchaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PurchaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PurchaseTransactorSession struct {
	Contract     *PurchaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PurchaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type PurchaseRaw struct {
	Contract *Purchase // Generic contract binding to access the raw methods on
}

// PurchaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PurchaseCallerRaw struct {
	Contract *PurchaseCaller // Generic read-only contract binding to access the raw methods on
}

// PurchaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PurchaseTransactorRaw struct {
	Contract *PurchaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPurchase creates a new instance of Purchase, bound to a specific deployed contract.
func NewPurchase(address common.Address, backend bind.ContractBackend) (*Purchase, error) {
	contract, err := bindPurchase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Purchase{PurchaseCaller: PurchaseCaller{contract: contract}, PurchaseTransactor: PurchaseTransactor{contract: contract}, PurchaseFilterer: PurchaseFilterer{contract: contract}}, nil
}

// NewPurchaseCaller creates a new read-only instance of Purchase, bound to a specific deployed contract.
func NewPurchaseCaller(address common.Address, caller bind.ContractCaller) (*PurchaseCaller, error) {
	contract, err := bindPurchase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PurchaseCaller{contract: contract}, nil
}

// NewPurchaseTransactor creates a new write-only instance of Purchase, bound to a specific deployed contract.
func NewPurchaseTransactor(address common.Address, transactor bind.ContractTransactor) (*PurchaseTransactor, error) {
	contract, err := bindPurchase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PurchaseTransactor{contract: contract}, nil
}

// NewPurchaseFilterer creates a new log filterer instance of Purchase, bound to a specific deployed contract.
func NewPurchaseFilterer(address common.Address, filterer bind.ContractFilterer) (*PurchaseFilterer, error) {
	contract, err := bindPurchase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PurchaseFilterer{contract: contract}, nil
}

// bindPurchase binds a generic wrapper to an already deployed contract.
func bindPurchase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PurchaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Purchase *PurchaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Purchase.Contract.PurchaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Purchase *PurchaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Purchase.Contract.PurchaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Purchase *PurchaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Purchase.Contract.PurchaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Purchase *PurchaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Purchase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Purchase *PurchaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Purchase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Purchase *PurchaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Purchase.Contract.contract.Transact(opts, method, params...)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Purchase *PurchaseCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Purchase.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Purchase *PurchaseSession) Buyer() (common.Address, error) {
	return _Purchase.Contract.Buyer(&_Purchase.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Purchase *PurchaseCallerSession) Buyer() (common.Address, error) {
	return _Purchase.Contract.Buyer(&_Purchase.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Purchase *PurchaseCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Purchase.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Purchase *PurchaseSession) GetBalance() (*big.Int, error) {
	return _Purchase.Contract.GetBalance(&_Purchase.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Purchase *PurchaseCallerSession) GetBalance() (*big.Int, error) {
	return _Purchase.Contract.GetBalance(&_Purchase.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Purchase *PurchaseCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Purchase.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Purchase *PurchaseSession) Seller() (common.Address, error) {
	return _Purchase.Contract.Seller(&_Purchase.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Purchase *PurchaseCallerSession) Seller() (common.Address, error) {
	return _Purchase.Contract.Seller(&_Purchase.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Purchase *PurchaseCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Purchase.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Purchase *PurchaseSession) State() (uint8, error) {
	return _Purchase.Contract.State(&_Purchase.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Purchase *PurchaseCallerSession) State() (uint8, error) {
	return _Purchase.Contract.State(&_Purchase.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Purchase *PurchaseCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Purchase.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Purchase *PurchaseSession) Value() (*big.Int, error) {
	return _Purchase.Contract.Value(&_Purchase.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_Purchase *PurchaseCallerSession) Value() (*big.Int, error) {
	return _Purchase.Contract.Value(&_Purchase.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Purchase *PurchaseTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Purchase.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Purchase *PurchaseSession) Abort() (*types.Transaction, error) {
	return _Purchase.Contract.Abort(&_Purchase.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Purchase *PurchaseTransactorSession) Abort() (*types.Transaction, error) {
	return _Purchase.Contract.Abort(&_Purchase.TransactOpts)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0xd6960697.
//
// Solidity: function confirmPurchase() payable returns()
func (_Purchase *PurchaseTransactor) ConfirmPurchase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Purchase.contract.Transact(opts, "confirmPurchase")
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0xd6960697.
//
// Solidity: function confirmPurchase() payable returns()
func (_Purchase *PurchaseSession) ConfirmPurchase() (*types.Transaction, error) {
	return _Purchase.Contract.ConfirmPurchase(&_Purchase.TransactOpts)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0xd6960697.
//
// Solidity: function confirmPurchase() payable returns()
func (_Purchase *PurchaseTransactorSession) ConfirmPurchase() (*types.Transaction, error) {
	return _Purchase.Contract.ConfirmPurchase(&_Purchase.TransactOpts)
}

// ConfirmReceived is a paid mutator transaction binding the contract method 0x73fac6f0.
//
// Solidity: function confirmReceived() returns()
func (_Purchase *PurchaseTransactor) ConfirmReceived(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Purchase.contract.Transact(opts, "confirmReceived")
}

// ConfirmReceived is a paid mutator transaction binding the contract method 0x73fac6f0.
//
// Solidity: function confirmReceived() returns()
func (_Purchase *PurchaseSession) ConfirmReceived() (*types.Transaction, error) {
	return _Purchase.Contract.ConfirmReceived(&_Purchase.TransactOpts)
}

// ConfirmReceived is a paid mutator transaction binding the contract method 0x73fac6f0.
//
// Solidity: function confirmReceived() returns()
func (_Purchase *PurchaseTransactorSession) ConfirmReceived() (*types.Transaction, error) {
	return _Purchase.Contract.ConfirmReceived(&_Purchase.TransactOpts)
}

// PurchaseAbortedIterator is returned from FilterAborted and is used to iterate over the raw logs and unpacked data for Aborted events raised by the Purchase contract.
type PurchaseAbortedIterator struct {
	Event *PurchaseAborted // Event containing the contract specifics and raw log

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
func (it *PurchaseAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseAborted)
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
		it.Event = new(PurchaseAborted)
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
func (it *PurchaseAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseAborted represents a Aborted event raised by the Purchase contract.
type PurchaseAborted struct {
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAborted is a free log retrieval operation binding the contract event 0x13c3922f06c44c3cac6a2c721f8ec8db793360b0180d5e560ca2028de2037eaa.
//
// Solidity: event Aborted(address seller)
func (_Purchase *PurchaseFilterer) FilterAborted(opts *bind.FilterOpts) (*PurchaseAbortedIterator, error) {

	logs, sub, err := _Purchase.contract.FilterLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return &PurchaseAbortedIterator{contract: _Purchase.contract, event: "Aborted", logs: logs, sub: sub}, nil
}

// WatchAborted is a free log subscription operation binding the contract event 0x13c3922f06c44c3cac6a2c721f8ec8db793360b0180d5e560ca2028de2037eaa.
//
// Solidity: event Aborted(address seller)
func (_Purchase *PurchaseFilterer) WatchAborted(opts *bind.WatchOpts, sink chan<- *PurchaseAborted) (event.Subscription, error) {

	logs, sub, err := _Purchase.contract.WatchLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseAborted)
				if err := _Purchase.contract.UnpackLog(event, "Aborted", log); err != nil {
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

// ParseAborted is a log parse operation binding the contract event 0x13c3922f06c44c3cac6a2c721f8ec8db793360b0180d5e560ca2028de2037eaa.
//
// Solidity: event Aborted(address seller)
func (_Purchase *PurchaseFilterer) ParseAborted(log types.Log) (*PurchaseAborted, error) {
	event := new(PurchaseAborted)
	if err := _Purchase.contract.UnpackLog(event, "Aborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchaseItemReceivedIterator is returned from FilterItemReceived and is used to iterate over the raw logs and unpacked data for ItemReceived events raised by the Purchase contract.
type PurchaseItemReceivedIterator struct {
	Event *PurchaseItemReceived // Event containing the contract specifics and raw log

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
func (it *PurchaseItemReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchaseItemReceived)
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
		it.Event = new(PurchaseItemReceived)
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
func (it *PurchaseItemReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchaseItemReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchaseItemReceived represents a ItemReceived event raised by the Purchase contract.
type PurchaseItemReceived struct {
	Buyer  common.Address
	Seller common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterItemReceived is a free log retrieval operation binding the contract event 0x977d335cc268af836f9e8542f5c430c79e387ceb4a005d8aa05350b655fd65d3.
//
// Solidity: event ItemReceived(address buyer, address seller, uint256 value)
func (_Purchase *PurchaseFilterer) FilterItemReceived(opts *bind.FilterOpts) (*PurchaseItemReceivedIterator, error) {

	logs, sub, err := _Purchase.contract.FilterLogs(opts, "ItemReceived")
	if err != nil {
		return nil, err
	}
	return &PurchaseItemReceivedIterator{contract: _Purchase.contract, event: "ItemReceived", logs: logs, sub: sub}, nil
}

// WatchItemReceived is a free log subscription operation binding the contract event 0x977d335cc268af836f9e8542f5c430c79e387ceb4a005d8aa05350b655fd65d3.
//
// Solidity: event ItemReceived(address buyer, address seller, uint256 value)
func (_Purchase *PurchaseFilterer) WatchItemReceived(opts *bind.WatchOpts, sink chan<- *PurchaseItemReceived) (event.Subscription, error) {

	logs, sub, err := _Purchase.contract.WatchLogs(opts, "ItemReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchaseItemReceived)
				if err := _Purchase.contract.UnpackLog(event, "ItemReceived", log); err != nil {
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

// ParseItemReceived is a log parse operation binding the contract event 0x977d335cc268af836f9e8542f5c430c79e387ceb4a005d8aa05350b655fd65d3.
//
// Solidity: event ItemReceived(address buyer, address seller, uint256 value)
func (_Purchase *PurchaseFilterer) ParseItemReceived(log types.Log) (*PurchaseItemReceived, error) {
	event := new(PurchaseItemReceived)
	if err := _Purchase.contract.UnpackLog(event, "ItemReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurchasePurchaseConfirmedIterator is returned from FilterPurchaseConfirmed and is used to iterate over the raw logs and unpacked data for PurchaseConfirmed events raised by the Purchase contract.
type PurchasePurchaseConfirmedIterator struct {
	Event *PurchasePurchaseConfirmed // Event containing the contract specifics and raw log

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
func (it *PurchasePurchaseConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurchasePurchaseConfirmed)
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
		it.Event = new(PurchasePurchaseConfirmed)
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
func (it *PurchasePurchaseConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurchasePurchaseConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurchasePurchaseConfirmed represents a PurchaseConfirmed event raised by the Purchase contract.
type PurchasePurchaseConfirmed struct {
	Buyer common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPurchaseConfirmed is a free log retrieval operation binding the contract event 0xa29484a0d47153a96f00a1369c6885bbc89e95e35273ee97c943c440f7c4c282.
//
// Solidity: event PurchaseConfirmed(address buyer, uint256 value)
func (_Purchase *PurchaseFilterer) FilterPurchaseConfirmed(opts *bind.FilterOpts) (*PurchasePurchaseConfirmedIterator, error) {

	logs, sub, err := _Purchase.contract.FilterLogs(opts, "PurchaseConfirmed")
	if err != nil {
		return nil, err
	}
	return &PurchasePurchaseConfirmedIterator{contract: _Purchase.contract, event: "PurchaseConfirmed", logs: logs, sub: sub}, nil
}

// WatchPurchaseConfirmed is a free log subscription operation binding the contract event 0xa29484a0d47153a96f00a1369c6885bbc89e95e35273ee97c943c440f7c4c282.
//
// Solidity: event PurchaseConfirmed(address buyer, uint256 value)
func (_Purchase *PurchaseFilterer) WatchPurchaseConfirmed(opts *bind.WatchOpts, sink chan<- *PurchasePurchaseConfirmed) (event.Subscription, error) {

	logs, sub, err := _Purchase.contract.WatchLogs(opts, "PurchaseConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurchasePurchaseConfirmed)
				if err := _Purchase.contract.UnpackLog(event, "PurchaseConfirmed", log); err != nil {
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

// ParsePurchaseConfirmed is a log parse operation binding the contract event 0xa29484a0d47153a96f00a1369c6885bbc89e95e35273ee97c943c440f7c4c282.
//
// Solidity: event PurchaseConfirmed(address buyer, uint256 value)
func (_Purchase *PurchaseFilterer) ParsePurchaseConfirmed(log types.Log) (*PurchasePurchaseConfirmed, error) {
	event := new(PurchasePurchaseConfirmed)
	if err := _Purchase.contract.UnpackLog(event, "PurchaseConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
