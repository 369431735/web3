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

// LotteryMetaData contains all meta data concerning the Lottery contract.
var LotteryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LotteryReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PlayerEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WinnerPicked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DRAW_COOLDOWN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_ENTRY_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlayersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeUntilNextDraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDrawTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pickWinner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"players\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600281905550611283806100666000396000f3fe6080604052600436106100a75760003560e01c8063a76594bf11610064578063a76594bf146101ae578063db2e21bc146101d9578063e8936cc6146101f0578063e97dcb621461021b578063f71d96cb14610225578063fa6f5f1e14610262576100a7565b806312065fe0146100ac578063481c6a75146100d757806349debdbc146101025780635cbbd6351461012d5780635d495aea146101585780638b5b9ccc14610183575b600080fd5b3480156100b857600080fd5b506100c161028d565b6040516100ce9190610a2e565b60405180910390f35b3480156100e357600080fd5b506100ec610295565b6040516100f99190610a8a565b60405180910390f35b34801561010e57600080fd5b506101176102b9565b6040516101249190610a2e565b60405180910390f35b34801561013957600080fd5b506101426102c4565b60405161014f9190610a2e565b60405180910390f35b34801561016457600080fd5b5061016d6102ca565b60405161017a9190610a8a565b60405180910390f35b34801561018f57600080fd5b506101986104ea565b6040516101a59190610b63565b60405180910390f35b3480156101ba57600080fd5b506101c3610578565b6040516101d09190610a2e565b60405180910390f35b3480156101e557600080fd5b506101ee610585565b005b3480156101fc57600080fd5b5061020561071d565b6040516102129190610a2e565b60405180910390f35b610223610723565b005b34801561023157600080fd5b5061024c60048036038101906102479190610bb6565b61080b565b6040516102599190610a8a565b60405180910390f35b34801561026e57600080fd5b5061027761084a565b6040516102849190610a2e565b60405180910390f35b600047905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b662386f26fc1000081565b60025481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461035b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035290610c66565b60405180910390fd5b6000600180549050116103a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039a90610cd2565b60405180910390fd5b610e106002546103b39190610d21565b4210156103f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ec90610dc7565b60405180910390fd5b6000600180549050610405610883565b61040f9190610e16565b905060006001828154811061042757610426610e47565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060004790508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561049f573d6000803e3d6000fd5b507f64791dbae5677392ba76761a5273633cec8f1d9d8cfe808da7bac6ef16a880be82826040516104d1929190610e76565b60405180910390a16104e16108ce565b81935050505090565b6060600180548060200260200160405190810160405280929190818152602001828054801561056e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610524575b5050505050905090565b6000600180549050905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610613576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060a90610c66565b60405180910390fd5b60004711610656576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064d90610eeb565b60405180910390fd5b6018610e106106659190610f0b565b6002546106729190610d21565b4210156106b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ab90610f99565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561071a573d6000803e3d6000fd5b50565b610e1081565b662386f26fc1000034101561076d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076490611005565b60405180910390fd5b6001339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fc16481c9484122696a660b57df0bf6839645c6b620cb1704ac2437de13be92253334604051610801929190610e76565b60405180910390a1565b6001818154811061081b57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080610e1060025461085d9190610d21565b9050804210610870576000915050610880565b428161087c9190611025565b9150505b90565b600080444243600180436108979190611025565b406040516020016108ac9594939291906111c3565b6040516020818303038152906040528051906020012090508060001c91505090565b600067ffffffffffffffff8111156108e9576108e861121e565b5b6040519080825280602002602001820160405280156109175781602001602082028036833780820191505090505b506001908051906020019061092d92919061096e565b50426002819055507f6564d21026f38ac4bc4ef082b01864f3ae6869e139a373fdf664a27a6c6dbebb426040516109649190610a2e565b60405180910390a1565b8280548282559060005260206000209081019282156109e7579160200282015b828111156109e65782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019061098e565b5b5090506109f491906109f8565b5090565b5b80821115610a115760008160009055506001016109f9565b5090565b6000819050919050565b610a2881610a15565b82525050565b6000602082019050610a436000830184610a1f565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a7482610a49565b9050919050565b610a8481610a69565b82525050565b6000602082019050610a9f6000830184610a7b565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610ada81610a69565b82525050565b6000610aec8383610ad1565b60208301905092915050565b6000602082019050919050565b6000610b1082610aa5565b610b1a8185610ab0565b9350610b2583610ac1565b8060005b83811015610b56578151610b3d8882610ae0565b9750610b4883610af8565b925050600181019050610b29565b5085935050505092915050565b60006020820190508181036000830152610b7d8184610b05565b905092915050565b600080fd5b610b9381610a15565b8114610b9e57600080fd5b50565b600081359050610bb081610b8a565b92915050565b600060208284031215610bcc57610bcb610b85565b5b6000610bda84828501610ba1565b91505092915050565b600082825260208201905092915050565b7f4f6e6c79206d616e616765722063616e2063616c6c20746869732066756e637460008201527f696f6e0000000000000000000000000000000000000000000000000000000000602082015250565b6000610c50602383610be3565b9150610c5b82610bf4565b604082019050919050565b60006020820190508181036000830152610c7f81610c43565b9050919050565b7f4e6f20706c617965727320696e20746865206c6f747465727900000000000000600082015250565b6000610cbc601983610be3565b9150610cc782610c86565b602082019050919050565b60006020820190508181036000830152610ceb81610caf565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d2c82610a15565b9150610d3783610a15565b9250828201905080821115610d4f57610d4e610cf2565b5b92915050565b7f506c65617365207761697420666f722074686520636f6f6c646f776e2070657260008201527f696f640000000000000000000000000000000000000000000000000000000000602082015250565b6000610db1602383610be3565b9150610dbc82610d55565b604082019050919050565b60006020820190508181036000830152610de081610da4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610e2182610a15565b9150610e2c83610a15565b925082610e3c57610e3b610de7565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000604082019050610e8b6000830185610a7b565b610e986020830184610a1f565b9392505050565b7f4e6f2066756e647320746f207769746864726177000000000000000000000000600082015250565b6000610ed5601483610be3565b9150610ee082610e9f565b602082019050919050565b60006020820190508181036000830152610f0481610ec8565b9050919050565b6000610f1682610a15565b9150610f2183610a15565b9250828202610f2f81610a15565b91508282048414831517610f4657610f45610cf2565b5b5092915050565b7f4d757374207761697420323420636f6f6c646f776e20706572696f6473000000600082015250565b6000610f83601d83610be3565b9150610f8e82610f4d565b602082019050919050565b60006020820190508181036000830152610fb281610f76565b9050919050565b7f4d696e696d756d20656e7472792066656520697320302e303120657468657200600082015250565b6000610fef601f83610be3565b9150610ffa82610fb9565b602082019050919050565b6000602082019050818103600083015261101e81610fe2565b9050919050565b600061103082610a15565b915061103b83610a15565b925082820390508181111561105357611052610cf2565b5b92915050565b6000819050919050565b61107461106f82610a15565b611059565b82525050565b600081549050919050565b600081905092915050565b60008190508160005260206000209050919050565b6110ae81610a69565b82525050565b60006110c083836110a5565b60208301905092915050565b60008160001c9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061110c611107836110cc565b6110d9565b9050919050565b600061111f82546110f9565b9050919050565b6000600182019050919050565b600061113e8261107a565b6111488185611085565b935061115383611090565b8060005b8381101561118b5761116882611113565b61117288826110b4565b975061117d83611126565b925050600181019050611157565b5085935050505092915050565b6000819050919050565b6000819050919050565b6111bd6111b882611198565b6111a2565b82525050565b60006111cf8288611063565b6020820191506111df8287611063565b6020820191506111ef8286611063565b6020820191506111ff8285611133565b915061120b82846111ac565b6020820191508190509695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea2646970667358221220007cffac7999120342366cf6eb65703a5f9bb7d457ff4e10990c31fe55ca456264736f6c634300081c0033",
}

// LotteryABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryMetaData.ABI instead.
var LotteryABI = LotteryMetaData.ABI

// LotteryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LotteryMetaData.Bin instead.
var LotteryBin = LotteryMetaData.Bin

// DeployLottery deploys a new Ethereum contract, binding an instance of Lottery to it.
func DeployLottery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lottery, error) {
	parsed, err := LotteryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LotteryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// Lottery is an auto generated Go binding around an Ethereum contract.
type Lottery struct {
	LotteryCaller     // Read-only binding to the contract
	LotteryTransactor // Write-only binding to the contract
	LotteryFilterer   // Log filterer for contract events
}

// LotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotterySession struct {
	Contract     *Lottery          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryCallerSession struct {
	Contract *LotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryTransactorSession struct {
	Contract     *LotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryRaw struct {
	Contract *Lottery // Generic contract binding to access the raw methods on
}

// LotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryCallerRaw struct {
	Contract *LotteryCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryTransactorRaw struct {
	Contract *LotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLottery creates a new instance of Lottery, bound to a specific deployed contract.
func NewLottery(address common.Address, backend bind.ContractBackend) (*Lottery, error) {
	contract, err := bindLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// NewLotteryCaller creates a new read-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryCaller(address common.Address, caller bind.ContractCaller) (*LotteryCaller, error) {
	contract, err := bindLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryCaller{contract: contract}, nil
}

// NewLotteryTransactor creates a new write-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryTransactor, error) {
	contract, err := bindLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryTransactor{contract: contract}, nil
}

// NewLotteryFilterer creates a new log filterer instance of Lottery, bound to a specific deployed contract.
func NewLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryFilterer, error) {
	contract, err := bindLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryFilterer{contract: contract}, nil
}

// bindLottery binds a generic wrapper to an already deployed contract.
func bindLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LotteryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.LotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transact(opts, method, params...)
}

// DRAWCOOLDOWN is a free data retrieval call binding the contract method 0xe8936cc6.
//
// Solidity: function DRAW_COOLDOWN() view returns(uint256)
func (_Lottery *LotteryCaller) DRAWCOOLDOWN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "DRAW_COOLDOWN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DRAWCOOLDOWN is a free data retrieval call binding the contract method 0xe8936cc6.
//
// Solidity: function DRAW_COOLDOWN() view returns(uint256)
func (_Lottery *LotterySession) DRAWCOOLDOWN() (*big.Int, error) {
	return _Lottery.Contract.DRAWCOOLDOWN(&_Lottery.CallOpts)
}

// DRAWCOOLDOWN is a free data retrieval call binding the contract method 0xe8936cc6.
//
// Solidity: function DRAW_COOLDOWN() view returns(uint256)
func (_Lottery *LotteryCallerSession) DRAWCOOLDOWN() (*big.Int, error) {
	return _Lottery.Contract.DRAWCOOLDOWN(&_Lottery.CallOpts)
}

// MINIMUMENTRYFEE is a free data retrieval call binding the contract method 0x49debdbc.
//
// Solidity: function MINIMUM_ENTRY_FEE() view returns(uint256)
func (_Lottery *LotteryCaller) MINIMUMENTRYFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "MINIMUM_ENTRY_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMENTRYFEE is a free data retrieval call binding the contract method 0x49debdbc.
//
// Solidity: function MINIMUM_ENTRY_FEE() view returns(uint256)
func (_Lottery *LotterySession) MINIMUMENTRYFEE() (*big.Int, error) {
	return _Lottery.Contract.MINIMUMENTRYFEE(&_Lottery.CallOpts)
}

// MINIMUMENTRYFEE is a free data retrieval call binding the contract method 0x49debdbc.
//
// Solidity: function MINIMUM_ENTRY_FEE() view returns(uint256)
func (_Lottery *LotteryCallerSession) MINIMUMENTRYFEE() (*big.Int, error) {
	return _Lottery.Contract.MINIMUMENTRYFEE(&_Lottery.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Lottery *LotteryCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Lottery *LotterySession) GetBalance() (*big.Int, error) {
	return _Lottery.Contract.GetBalance(&_Lottery.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Lottery *LotteryCallerSession) GetBalance() (*big.Int, error) {
	return _Lottery.Contract.GetBalance(&_Lottery.CallOpts)
}

// GetPlayers is a free data retrieval call binding the contract method 0x8b5b9ccc.
//
// Solidity: function getPlayers() view returns(address[])
func (_Lottery *LotteryCaller) GetPlayers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getPlayers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPlayers is a free data retrieval call binding the contract method 0x8b5b9ccc.
//
// Solidity: function getPlayers() view returns(address[])
func (_Lottery *LotterySession) GetPlayers() ([]common.Address, error) {
	return _Lottery.Contract.GetPlayers(&_Lottery.CallOpts)
}

// GetPlayers is a free data retrieval call binding the contract method 0x8b5b9ccc.
//
// Solidity: function getPlayers() view returns(address[])
func (_Lottery *LotteryCallerSession) GetPlayers() ([]common.Address, error) {
	return _Lottery.Contract.GetPlayers(&_Lottery.CallOpts)
}

// GetPlayersCount is a free data retrieval call binding the contract method 0xa76594bf.
//
// Solidity: function getPlayersCount() view returns(uint256)
func (_Lottery *LotteryCaller) GetPlayersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getPlayersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPlayersCount is a free data retrieval call binding the contract method 0xa76594bf.
//
// Solidity: function getPlayersCount() view returns(uint256)
func (_Lottery *LotterySession) GetPlayersCount() (*big.Int, error) {
	return _Lottery.Contract.GetPlayersCount(&_Lottery.CallOpts)
}

// GetPlayersCount is a free data retrieval call binding the contract method 0xa76594bf.
//
// Solidity: function getPlayersCount() view returns(uint256)
func (_Lottery *LotteryCallerSession) GetPlayersCount() (*big.Int, error) {
	return _Lottery.Contract.GetPlayersCount(&_Lottery.CallOpts)
}

// GetTimeUntilNextDraw is a free data retrieval call binding the contract method 0xfa6f5f1e.
//
// Solidity: function getTimeUntilNextDraw() view returns(uint256)
func (_Lottery *LotteryCaller) GetTimeUntilNextDraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getTimeUntilNextDraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeUntilNextDraw is a free data retrieval call binding the contract method 0xfa6f5f1e.
//
// Solidity: function getTimeUntilNextDraw() view returns(uint256)
func (_Lottery *LotterySession) GetTimeUntilNextDraw() (*big.Int, error) {
	return _Lottery.Contract.GetTimeUntilNextDraw(&_Lottery.CallOpts)
}

// GetTimeUntilNextDraw is a free data retrieval call binding the contract method 0xfa6f5f1e.
//
// Solidity: function getTimeUntilNextDraw() view returns(uint256)
func (_Lottery *LotteryCallerSession) GetTimeUntilNextDraw() (*big.Int, error) {
	return _Lottery.Contract.GetTimeUntilNextDraw(&_Lottery.CallOpts)
}

// LastDrawTime is a free data retrieval call binding the contract method 0x5cbbd635.
//
// Solidity: function lastDrawTime() view returns(uint256)
func (_Lottery *LotteryCaller) LastDrawTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "lastDrawTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDrawTime is a free data retrieval call binding the contract method 0x5cbbd635.
//
// Solidity: function lastDrawTime() view returns(uint256)
func (_Lottery *LotterySession) LastDrawTime() (*big.Int, error) {
	return _Lottery.Contract.LastDrawTime(&_Lottery.CallOpts)
}

// LastDrawTime is a free data retrieval call binding the contract method 0x5cbbd635.
//
// Solidity: function lastDrawTime() view returns(uint256)
func (_Lottery *LotteryCallerSession) LastDrawTime() (*big.Int, error) {
	return _Lottery.Contract.LastDrawTime(&_Lottery.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Lottery *LotteryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Lottery *LotterySession) Manager() (common.Address, error) {
	return _Lottery.Contract.Manager(&_Lottery.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Lottery *LotteryCallerSession) Manager() (common.Address, error) {
	return _Lottery.Contract.Manager(&_Lottery.CallOpts)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_Lottery *LotteryCaller) Players(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "players", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_Lottery *LotterySession) Players(arg0 *big.Int) (common.Address, error) {
	return _Lottery.Contract.Players(&_Lottery.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_Lottery *LotteryCallerSession) Players(arg0 *big.Int) (common.Address, error) {
	return _Lottery.Contract.Players(&_Lottery.CallOpts, arg0)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Lottery *LotteryTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Lottery *LotterySession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Lottery.Contract.EmergencyWithdraw(&_Lottery.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Lottery *LotteryTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Lottery.Contract.EmergencyWithdraw(&_Lottery.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0xe97dcb62.
//
// Solidity: function enter() payable returns()
func (_Lottery *LotteryTransactor) Enter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "enter")
}

// Enter is a paid mutator transaction binding the contract method 0xe97dcb62.
//
// Solidity: function enter() payable returns()
func (_Lottery *LotterySession) Enter() (*types.Transaction, error) {
	return _Lottery.Contract.Enter(&_Lottery.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0xe97dcb62.
//
// Solidity: function enter() payable returns()
func (_Lottery *LotteryTransactorSession) Enter() (*types.Transaction, error) {
	return _Lottery.Contract.Enter(&_Lottery.TransactOpts)
}

// PickWinner is a paid mutator transaction binding the contract method 0x5d495aea.
//
// Solidity: function pickWinner() returns(address)
func (_Lottery *LotteryTransactor) PickWinner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "pickWinner")
}

// PickWinner is a paid mutator transaction binding the contract method 0x5d495aea.
//
// Solidity: function pickWinner() returns(address)
func (_Lottery *LotterySession) PickWinner() (*types.Transaction, error) {
	return _Lottery.Contract.PickWinner(&_Lottery.TransactOpts)
}

// PickWinner is a paid mutator transaction binding the contract method 0x5d495aea.
//
// Solidity: function pickWinner() returns(address)
func (_Lottery *LotteryTransactorSession) PickWinner() (*types.Transaction, error) {
	return _Lottery.Contract.PickWinner(&_Lottery.TransactOpts)
}

// LotteryLotteryResetIterator is returned from FilterLotteryReset and is used to iterate over the raw logs and unpacked data for LotteryReset events raised by the Lottery contract.
type LotteryLotteryResetIterator struct {
	Event *LotteryLotteryReset // Event containing the contract specifics and raw log

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
func (it *LotteryLotteryResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryLotteryReset)
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
		it.Event = new(LotteryLotteryReset)
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
func (it *LotteryLotteryResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryLotteryResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryLotteryReset represents a LotteryReset event raised by the Lottery contract.
type LotteryLotteryReset struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLotteryReset is a free log retrieval operation binding the contract event 0x6564d21026f38ac4bc4ef082b01864f3ae6869e139a373fdf664a27a6c6dbebb.
//
// Solidity: event LotteryReset(uint256 timestamp)
func (_Lottery *LotteryFilterer) FilterLotteryReset(opts *bind.FilterOpts) (*LotteryLotteryResetIterator, error) {

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "LotteryReset")
	if err != nil {
		return nil, err
	}
	return &LotteryLotteryResetIterator{contract: _Lottery.contract, event: "LotteryReset", logs: logs, sub: sub}, nil
}

// WatchLotteryReset is a free log subscription operation binding the contract event 0x6564d21026f38ac4bc4ef082b01864f3ae6869e139a373fdf664a27a6c6dbebb.
//
// Solidity: event LotteryReset(uint256 timestamp)
func (_Lottery *LotteryFilterer) WatchLotteryReset(opts *bind.WatchOpts, sink chan<- *LotteryLotteryReset) (event.Subscription, error) {

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "LotteryReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryLotteryReset)
				if err := _Lottery.contract.UnpackLog(event, "LotteryReset", log); err != nil {
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

// ParseLotteryReset is a log parse operation binding the contract event 0x6564d21026f38ac4bc4ef082b01864f3ae6869e139a373fdf664a27a6c6dbebb.
//
// Solidity: event LotteryReset(uint256 timestamp)
func (_Lottery *LotteryFilterer) ParseLotteryReset(log types.Log) (*LotteryLotteryReset, error) {
	event := new(LotteryLotteryReset)
	if err := _Lottery.contract.UnpackLog(event, "LotteryReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryPlayerEnteredIterator is returned from FilterPlayerEntered and is used to iterate over the raw logs and unpacked data for PlayerEntered events raised by the Lottery contract.
type LotteryPlayerEnteredIterator struct {
	Event *LotteryPlayerEntered // Event containing the contract specifics and raw log

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
func (it *LotteryPlayerEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryPlayerEntered)
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
		it.Event = new(LotteryPlayerEntered)
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
func (it *LotteryPlayerEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryPlayerEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryPlayerEntered represents a PlayerEntered event raised by the Lottery contract.
type LotteryPlayerEntered struct {
	Player common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlayerEntered is a free log retrieval operation binding the contract event 0xc16481c9484122696a660b57df0bf6839645c6b620cb1704ac2437de13be9225.
//
// Solidity: event PlayerEntered(address player, uint256 amount)
func (_Lottery *LotteryFilterer) FilterPlayerEntered(opts *bind.FilterOpts) (*LotteryPlayerEnteredIterator, error) {

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "PlayerEntered")
	if err != nil {
		return nil, err
	}
	return &LotteryPlayerEnteredIterator{contract: _Lottery.contract, event: "PlayerEntered", logs: logs, sub: sub}, nil
}

// WatchPlayerEntered is a free log subscription operation binding the contract event 0xc16481c9484122696a660b57df0bf6839645c6b620cb1704ac2437de13be9225.
//
// Solidity: event PlayerEntered(address player, uint256 amount)
func (_Lottery *LotteryFilterer) WatchPlayerEntered(opts *bind.WatchOpts, sink chan<- *LotteryPlayerEntered) (event.Subscription, error) {

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "PlayerEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryPlayerEntered)
				if err := _Lottery.contract.UnpackLog(event, "PlayerEntered", log); err != nil {
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

// ParsePlayerEntered is a log parse operation binding the contract event 0xc16481c9484122696a660b57df0bf6839645c6b620cb1704ac2437de13be9225.
//
// Solidity: event PlayerEntered(address player, uint256 amount)
func (_Lottery *LotteryFilterer) ParsePlayerEntered(log types.Log) (*LotteryPlayerEntered, error) {
	event := new(LotteryPlayerEntered)
	if err := _Lottery.contract.UnpackLog(event, "PlayerEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryWinnerPickedIterator is returned from FilterWinnerPicked and is used to iterate over the raw logs and unpacked data for WinnerPicked events raised by the Lottery contract.
type LotteryWinnerPickedIterator struct {
	Event *LotteryWinnerPicked // Event containing the contract specifics and raw log

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
func (it *LotteryWinnerPickedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryWinnerPicked)
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
		it.Event = new(LotteryWinnerPicked)
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
func (it *LotteryWinnerPickedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryWinnerPickedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryWinnerPicked represents a WinnerPicked event raised by the Lottery contract.
type LotteryWinnerPicked struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWinnerPicked is a free log retrieval operation binding the contract event 0x64791dbae5677392ba76761a5273633cec8f1d9d8cfe808da7bac6ef16a880be.
//
// Solidity: event WinnerPicked(address winner, uint256 amount)
func (_Lottery *LotteryFilterer) FilterWinnerPicked(opts *bind.FilterOpts) (*LotteryWinnerPickedIterator, error) {

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "WinnerPicked")
	if err != nil {
		return nil, err
	}
	return &LotteryWinnerPickedIterator{contract: _Lottery.contract, event: "WinnerPicked", logs: logs, sub: sub}, nil
}

// WatchWinnerPicked is a free log subscription operation binding the contract event 0x64791dbae5677392ba76761a5273633cec8f1d9d8cfe808da7bac6ef16a880be.
//
// Solidity: event WinnerPicked(address winner, uint256 amount)
func (_Lottery *LotteryFilterer) WatchWinnerPicked(opts *bind.WatchOpts, sink chan<- *LotteryWinnerPicked) (event.Subscription, error) {

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "WinnerPicked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryWinnerPicked)
				if err := _Lottery.contract.UnpackLog(event, "WinnerPicked", log); err != nil {
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

// ParseWinnerPicked is a log parse operation binding the contract event 0x64791dbae5677392ba76761a5273633cec8f1d9d8cfe808da7bac6ef16a880be.
//
// Solidity: event WinnerPicked(address winner, uint256 amount)
func (_Lottery *LotteryFilterer) ParseWinnerPicked(log types.Log) (*LotteryWinnerPicked, error) {
	event := new(LotteryWinnerPicked)
	if err := _Lottery.contract.UnpackLog(event, "WinnerPicked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
