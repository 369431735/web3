// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// LockMetaData contains all meta data concerning the Lock contract.
var LockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"when\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"unlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040516102a73803806102a783398181016040528101906100259190610105565b8042106100685760405162461bcd60e51b815260206004820152601460248201527f556e6c6f636b2074696d65206973206e6f74207365740000000000000000000060448201526064015b60405180910390fd5b600055610132565b600080fd5b600080fd5b600080fd5b6000819050919050565b6100948161008a565b81146100a857600080fd5b50565b6000815190506100b782610082565b92915050565b6000815190506100cc81610082565b92915050565b6000604082840312156100e2576100e161007e565b5b60006100f084828501610082565b91505092915050565b60006020828403121561010457600080fd5b600082013590506101158161008a565b92915050565b60006020828403121561012b57600080fd5b600061012984828501610082565b91505092915050565b600061010084840312156101485761014761007e565b5b600061015684828501610082565b9150506020610167848285016100c8565b90509250929050565b610167806101416000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e1a7d4d1461003b5780633292cf6014610057575b600080fd5b610043610075565b6040516100509190610103565b60405180910390f35b61005f610083565b60405161006c9190610103565b60405180910390f35b60008054905090565b6000544210156100f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100e7906101bb565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff164760405160006040518083038185875af1925050503d8060008114610147576040519150601f19603f3d011682016040523d82523d6000602084013e61014c565b606091505b5050905060004790507f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5814260405161009d9291906101e1565b60008190505b801561011d57600061011a565b90565b6000819050919050565b61011e81610111565b82525050565b60006020820190506101396000830184610115565b92915050565b600082825260208201905092915050565b7f596f752063616e27742077697468647261772079657400000000000000000000600082015250565b60006101a560148361013f565b91506101b08261014f565b602082019050919050565b600060208201905081810360008301526101d481610199565b9050919050565b60006040820190506101f06000830185610115565b6101fd602083018461011556b939250505056fea26469706673582212201af7996fee82a80e9c964516f4b486d3ca07e3cabc0af0b2358ef5db39c12a3f64736f6c634300080d0033",
}

// LockABI is the input ABI used to generate the binding from.
const LockABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"when\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"unlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Lock is an auto generated Go binding around an Ethereum contract.
type Lock struct {
	LockCaller     // Read-only binding to the contract
	LockTransactor // Write-only binding to the contract
	LockFilterer   // Log filterer for contract events
}

// LockCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockCaller struct {
	contract *bind.BoundContract
}

// LockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockTransactor struct {
	contract *bind.BoundContract
}

// LockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockFilterer struct {
	contract *bind.BoundContract
}

// LockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockSession struct {
	Contract     *Lock
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

// NewLock creates a new instance of Lock, bound to a specific deployed contract.
func NewLock(address common.Address, backend bind.ContractBackend) (*Lock, error) {
	contract, err := bindLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lock{
		LockCaller:     LockCaller{contract: contract},
		LockTransactor: LockTransactor{contract: contract},
		LockFilterer:   LockFilterer{contract: contract},
	}, nil
}

// bindLock binds a generic wrapper to an already deployed contract.
func bindLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// LockWithdrawal represents a Withdrawal event raised by the Lock contract.
type LockWithdrawal struct {
	Amount *big.Int
	When   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawal(uint256 amount, uint256 when)
func (f *Lock) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *LockWithdrawal) (event.Subscription, error) {
	logCh, sub, err := f.LockFilterer.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case log := <-logCh:
				event := new(LockWithdrawal)
				event.Raw = log

				if err := f.LockFilterer.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					continue
				}

				select {
				case sink <- event:
				case <-sub.Err():
					return
				}
			case <-sub.Err():
				return
			}
		}
	}()

	return sub, nil
}
