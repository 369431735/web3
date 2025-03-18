package utils

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"task2/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// TransferTokens 发送代币交易
func TransferTokens(toAddress string, amount *big.Int) (string, error) {
	log.Println("=== 发送代币交易演示 ===")

	// 连接到以太坊网络
	client, err := InitClient()
	if err != nil {
		return "", fmt.Errorf("连接以太坊网络失败: %v", err)
	}

	// 获取实际的链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取链ID失败: %v", err)
	}

	network := config.GetCurrentNetwork()
	log.Printf("当前网络: %s (Chain ID: %d)", network.NetworkName, chainID)

	// 获取私钥
	privateKey, err := GetPrivateKey(network.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("解析私钥失败: %v", err)
	}

	// 获取发送者地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	log.Printf("发送者地址: %s", fromAddress.Hex())

	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("获取nonce失败: %v", err)
	}
	log.Printf("当前nonce: %d", nonce)

	// 获取gas limit
	gasLimit := uint64(21000)
	log.Printf("Gas limit: %d", gasLimit)

	// 获取gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取gas price失败: %v", err)
	}
	log.Printf("Gas price: %s", gasPrice.String())

	// 创建交易
	to := common.HexToAddress(toAddress)
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)
	log.Println("交易创建成功")

	// 使用网络配置中的签名器
	signer := network.GetSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return "", fmt.Errorf("签名交易失败: %v", err)
	}
	log.Println("交易签名成功")

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("发送交易失败: %v", err)
	}

	txHash := signedTx.Hash().Hex()
	log.Printf("交易发送成功，交易哈希: %s", txHash)
	return txHash, nil
}

// TransferTokensWithData 发送带数据的代币交易
func TransferTokensWithData(toAddress string, amount *big.Int, data []byte) (string, error) {
	log.Println("=== 发送带数据的代币交易演示 ===")

	// 连接到以太坊网络
	client, err := InitClient()
	if err != nil {
		return "", fmt.Errorf("连接以太坊网络失败: %v", err)
	}

	// 获取实际的链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取链ID失败: %v", err)
	}

	network := config.GetCurrentNetwork()
	log.Printf("当前网络: %s (Chain ID: %d)", network.NetworkName, chainID)

	// 获取私钥
	privateKey, err := GetPrivateKey(network.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("解析私钥失败: %v", err)
	}

	// 获取发送者地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	log.Printf("发送者地址: %s", fromAddress.Hex())

	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("获取nonce失败: %v", err)
	}
	log.Printf("当前nonce: %d", nonce)

	// 获取gas limit
	gasLimit := uint64(21000)
	log.Printf("Gas limit: %d", gasLimit)

	// 获取gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取gas price失败: %v", err)
	}
	log.Printf("Gas price: %s", gasPrice.String())

	// 创建交易
	to := common.HexToAddress(toAddress)
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	log.Println("交易创建成功")

	// 使用网络配置中的签名器
	signer := network.GetSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return "", fmt.Errorf("签名交易失败: %v", err)
	}
	log.Println("交易签名成功")

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("发送交易失败: %v", err)
	}

	txHash := signedTx.Hash().Hex()
	log.Printf("交易发送成功，交易哈希: %s", txHash)
	return txHash, nil
}
