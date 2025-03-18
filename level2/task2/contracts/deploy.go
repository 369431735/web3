package contracts

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"task2/abi/bindings"
	"task2/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetTransactOpts 获取交易选项
func GetTransactOpts(client *ethclient.Client) (*bind.TransactOpts, error) {
	network := config.GetCurrentNetwork()
	privateKey, err := crypto.HexToECDSA(network.PrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth, nil
}

// DeploySimpleStorage 部署 SimpleStorage 合约
func DeploySimpleStorage(client *ethclient.Client) (common.Address, common.Hash, *bindings.SimpleStorage, error) {
	auth, err := GetTransactOpts(client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	address, tx, instance, err := bindings.DeploySimpleStorage(auth, client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	return address, tx.Hash(), instance, nil
}

// DeployLock 部署 Lock 合约
func DeployLock(client *ethclient.Client) (common.Address, common.Hash, *bindings.Lock, error) {
	auth, err := GetTransactOpts(client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}

	// 设置锁定时间为 1 小时
	unlockTime := big.NewInt(3600)
	address, tx, instance, err := bindings.DeployLock(auth, client, unlockTime)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	return address, tx.Hash(), instance, nil
}

// DeployShipping 部署 Shipping 合约
func DeployShipping(client *ethclient.Client) (common.Address, common.Hash, *bindings.Shipping, error) {
	auth, err := GetTransactOpts(client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	address, tx, instance, err := bindings.DeployShipping(auth, client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	return address, tx.Hash(), instance, nil
}

// DeploySimpleAuction 部署 SimpleAuction 合约
func DeploySimpleAuction(client *ethclient.Client) (common.Address, common.Hash, *bindings.SimpleAuction, error) {
	auth, err := GetTransactOpts(client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}

	// 设置拍卖时间为 1 小时
	biddingTime := big.NewInt(3600)
	privateKey, err := crypto.HexToECDSA(config.GetCurrentNetwork().PrivateKey)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, common.Hash{}, nil, err
	}
	beneficiary := crypto.PubkeyToAddress(*publicKeyECDSA)

	address, tx, instance, err := bindings.DeploySimpleAuction(auth, client, biddingTime, beneficiary)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	return address, tx.Hash(), instance, nil
}

// DeployArrayDemo 部署 ArrayDemo 合约
func DeployArrayDemo(client *ethclient.Client) (common.Address, common.Hash, *bindings.ArrayDemo, error) {
	auth, err := GetTransactOpts(client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	address, tx, instance, err := bindings.DeployArrayDemo(auth, client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, err
	}
	return address, tx.Hash(), instance, nil
}

// GetContractBytecode 获取合约字节码
func GetContractBytecode(client *ethclient.Client, address string) ([]byte, error) {
	return client.CodeAt(context.Background(), common.HexToAddress(address), nil)
}
