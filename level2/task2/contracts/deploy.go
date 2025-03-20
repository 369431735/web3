package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DeployLock 使用提供的交易者在区块链上部署一个新的Lock合约
// 参数 _unlockTime 指定合约解锁时间
func DeployLock(auth *bind.TransactOpts, backend bind.ContractBackend, _unlockTime *big.Int) (common.Address, *types.Transaction, *Lock, error) {
	// 使用合约的元数据创建交易
	parsed, err := abi.JSON(strings.NewReader(LockMetaData.ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockMetaData.Bin), backend, _unlockTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	// 创建正确的合约实例
	instance := &Lock{
		LockCaller:     LockCaller{contract: contract},
		LockTransactor: LockTransactor{contract: contract},
		LockFilterer:   LockFilterer{contract: contract},
	}

	return address, tx, instance, nil
}

// DeploySimpleStorage 使用提供的交易者在区块链上部署一个新的SimpleStorage合约
func DeploySimpleStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleStorage, error) {
	// 使用合约的元数据创建交易
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageMetaData.ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleStorageMetaData.Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	// 创建正确的合约实例
	instance := &SimpleStorage{
		SimpleStorageCaller:     SimpleStorageCaller{contract: contract},
		SimpleStorageTransactor: SimpleStorageTransactor{contract: contract},
		SimpleStorageFilterer:   SimpleStorageFilterer{contract: contract},
	}

	return address, tx, instance, nil
}
