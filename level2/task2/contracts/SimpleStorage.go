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

// SimpleStorageMetaData contains all meta data concerning the SimpleStorage contract.
var SimpleStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806360fe47b11461003b5780636d4ce63c14610057575b600080fd5b6100556004803603810190610050919061009d565b610075565b005b61005f6100b9565b60405161006c91906100d9565b60405180910390f35b806000819055507fb922f092a64f1a076de6f21e4d7c6400b6e55791cc935e7bb8e7e90f7652f15b816040516100aa91906100d9565b60405180910390a150565b60008054905090565b6000813590506100d381610103565b92915050565b6100e381610100565b82525050565b60006020820190506100fe60008301846100da565b92915050565b6000819050919050565b61010c81610100565b811461011757600080fd5b5056fea26469706673582212202a84c2a884a03d9b7a6cf0c3a7a3033d14e60e2b50ac9fbfc9d1d75c8c70d19a64736f6c63430008070033",
}

// SimpleStorageABI is the input ABI used to generate the binding from.
const SimpleStorageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValueChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SimpleStorage is an auto generated Go binding around an Ethereum contract.
type SimpleStorage struct {
	SimpleStorageCaller     // Read-only binding to the contract
	SimpleStorageTransactor // Write-only binding to the contract
	SimpleStorageFilterer   // Log filterer for contract events
}

// SimpleStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleStorageCaller struct {
	contract *bind.BoundContract
}

// SimpleStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleStorageTransactor struct {
	contract *bind.BoundContract
}

// SimpleStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleStorageFilterer struct {
	contract *bind.BoundContract
}

// SimpleStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleStorageSession struct {
	Contract     *SimpleStorage
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

// NewSimpleStorage creates a new instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) (*SimpleStorage, error) {
	contract, err := bindSimpleStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleStorage{
		SimpleStorageCaller:     SimpleStorageCaller{contract: contract},
		SimpleStorageTransactor: SimpleStorageTransactor{contract: contract},
		SimpleStorageFilterer:   SimpleStorageFilterer{contract: contract},
	}, nil
}

// bindSimpleStorage binds a generic wrapper to an already deployed contract.
func bindSimpleStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// SimpleStorageValueChanged represents a ValueChanged event raised by the SimpleStorage contract.
type SimpleStorageValueChanged struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchValueChanged is a free log subscription operation binding the contract event 0xb922f092a64f1a076de6f21e4d7c6400b6e55791cc935e7bb8e7e90f7652f15b.
//
// Solidity: event ValueChanged(uint256 value)
func (f *SimpleStorage) WatchValueChanged(opts *bind.WatchOpts, sink chan<- *SimpleStorageValueChanged) (event.Subscription, error) {
	logCh, sub, err := f.SimpleStorageFilterer.contract.WatchLogs(opts, "ValueChanged")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case log := <-logCh:
				event := new(SimpleStorageValueChanged)
				event.Raw = log

				if err := f.SimpleStorageFilterer.contract.UnpackLog(event, "ValueChanged", log); err != nil {
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
