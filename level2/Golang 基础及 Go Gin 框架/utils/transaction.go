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

// 创建并发送交易
func CreateAndSendTransaction() (string, error) {
	log.Println("=== 创建并发送交易演示 ===")

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
	toAddress := common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	log.Printf("接收者地址: %s", toAddress.Hex())

	// 获取发送者的 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("获取nonce失败: %v", err)
	}
	log.Printf("当前 nonce: %d", nonce)

	// 设置交易金额（根据是否是测试网络调整）
	value := big.NewInt(1000000000000000) // 0.001 ETH for testnet
	log.Printf("发送金额: %s Wei", value.String())

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
		To:       &toAddress,
		Value:    value,
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

// 查询交易
func QueryTransaction(txHash string) error {
	log.Println("=== 交易查询演示 ===")

	client, err := InitClient()
	if err != nil {
		return fmt.Errorf("连接以太坊网络失败: %v", err)
	}

	// 将字符串转换为common.Hash类型
	hash := common.HexToHash(txHash)

	// 获取交易详情
	tx, isPending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return fmt.Errorf("获取交易信息失败: %v", err)
	}

	// 获取实际的链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("获取链ID失败: %v", err)
	}

	// 使用网络配置中的签名器
	signer := CurrentNetwork.GetSigner(chainID)
	sender, err := signer.Sender(tx)
	if err != nil {
		return fmt.Errorf("获取交易发送者失败: %v", err)
	}

	// 输出交易基本信息
	log.Printf("交易哈希: %s", tx.Hash().Hex())
	log.Printf("交易是否待处理: %v", isPending)
	log.Printf("交易发送者: %s", sender.Hex())
	if tx.To() != nil {
		log.Printf("交易接收者: %s", tx.To().Hex())
	} else {
		log.Println("交易接收者: 合约创建交易")
	}
	log.Printf("交易金额: %s Wei", tx.Value().String())
	log.Printf("交易Gas限制: %d", tx.Gas())
	log.Printf("交易Gas价格: %s Wei", tx.GasPrice().String())
	log.Printf("交易Nonce: %d", tx.Nonce())

	// 如果交易已完成，获取交易收据
	if !isPending {
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err != nil {
			return fmt.Errorf("获取交易收据失败: %v", err)
		}

		log.Printf("交易状态: %d", receipt.Status)
		log.Printf("交易所在区块号: %d", receipt.BlockNumber)
		log.Printf("交易实际Gas使用量: %d", receipt.GasUsed)
		log.Printf("交易累计Gas使用量: %d", receipt.CumulativeGasUsed)
	}

	return nil
}
