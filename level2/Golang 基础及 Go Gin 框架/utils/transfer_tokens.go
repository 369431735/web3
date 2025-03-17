package utils

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	log.Printf("当前网络: %s (Chain ID: %d)", CurrentNetwork.NetworkName, chainID)

	// 解码私钥
	privateKeyBytes, err := hexutil.Decode(PrivateKey)
	if err != nil {
		return "", fmt.Errorf("私钥解码失败: %v", err)
	}

	// 创建私钥对象
	privateKeyECDSA, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("创建私钥对象失败: %v", err)
	}

	// 获取发送者地址
	fromAddress := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
	log.Printf("发送者地址: %s", fromAddress.Hex())

	// 获取接收者地址
	to := common.HexToAddress(toAddress)
	log.Printf("接收者地址: %s", to.Hex())

	// 获取发送者的 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("获取nonce失败: %v", err)
	}
	log.Printf("当前 nonce: %d", nonce)

	// 设置 gas 限制
	gasLimit := uint64(21000)
	log.Printf("Gas 限制: %d", gasLimit)

	// 获取当前建议的 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取gas价格失败: %v", err)
	}
	log.Printf("Gas 价格: %s Wei", gasPrice.String())

	// 创建基础交易数据
	data := &types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	}

	// 创建交易
	tx := types.NewTx(data)

	// 使用网络配置中的签名器
	signer := CurrentNetwork.GetSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKeyECDSA)
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
	log.Printf("当前网络: %s (Chain ID: %d)", CurrentNetwork.NetworkName, chainID)

	// 解码私钥
	privateKeyBytes, err := hexutil.Decode(PrivateKey)
	if err != nil {
		return "", fmt.Errorf("私钥解码失败: %v", err)
	}

	// 创建私钥对象
	privateKeyECDSA, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("创建私钥对象失败: %v", err)
	}

	// 获取发送者地址
	fromAddress := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
	log.Printf("发送者地址: %s", fromAddress.Hex())

	// 获取接收者地址
	to := common.HexToAddress(toAddress)
	log.Printf("接收者地址: %s", to.Hex())

	// 获取发送者的 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("获取nonce失败: %v", err)
	}
	log.Printf("当前 nonce: %d", nonce)

	// 设置 gas 限制（对于带数据的交易，需要更多的 gas）
	gasLimit := uint64(100000) // 增加 gas 限制
	log.Printf("Gas 限制: %d", gasLimit)

	// 获取当前建议的 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取gas价格失败: %v", err)
	}
	log.Printf("Gas 价格: %s Wei", gasPrice.String())

	// 创建基础交易数据
	txData := &types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	}

	// 创建交易
	tx := types.NewTx(txData)

	// 使用网络配置中的签名器
	signer := CurrentNetwork.GetSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKeyECDSA)
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
