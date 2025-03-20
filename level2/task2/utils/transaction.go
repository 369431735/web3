package utils

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"task2/config"
)

// CreateAndSendTransaction 创建并发送交易
func CreateAndSendTransaction(from, to string, amount *big.Int) (string, error) {
	client, err := GetEthClientHTTP()
	if err != nil {
		return "", err
	}
	// 使用的是单例客户端，不需要关闭

	// 获取网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return "", fmt.Errorf("未找到网络配置")
	}

	// 获取发送方账户
	account, ok := network.Accounts[from]
	if !ok {
		return "", fmt.Errorf("未找到账户: %s", from)
	}

	// 获取私钥
	privateKey, err := crypto.HexToECDSA(account.PrivateKey[2:]) // 移除 "0x" 前缀
	if err != nil {
		return "", fmt.Errorf("私钥解析失败: %v", err)
	}

	// 获取发送方地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("获取 nonce 失败: %v", err)
	}

	// 获取 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取 gas 价格失败: %v", err)
	}

	// 创建交易
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(to),
		amount,
		uint64(21000),
		gasPrice,
		nil,
	)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(network.ChainID)), privateKey)
	if err != nil {
		return "", fmt.Errorf("签名交易失败: %v", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("发送交易失败: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

// GetTransactionReceipt 获取交易收据
func GetTransactionReceipt(txHash string) (*types.Receipt, error) {
	client, err := GetEthClientHTTP()
	if err != nil {
		return nil, err
	}
	// 使用的是单例客户端，不需要关闭

	hash := common.HexToHash(txHash)
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, fmt.Errorf("获取交易收据失败: %v", err)
	}

	return receipt, nil
}

// GetTransactionByHash 获取交易信息
func GetTransactionByHash(txHash string) (*types.Transaction, error) {
	client, err := GetEthClientHTTP()
	if err != nil {
		return nil, err
	}
	// 使用的是单例客户端，不需要关闭

	hash := common.HexToHash(txHash)
	tx, isPending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return nil, fmt.Errorf("获取交易信息失败: %v", err)
	}

	log.Printf("交易状态: %v", isPending)
	return tx, nil
}
