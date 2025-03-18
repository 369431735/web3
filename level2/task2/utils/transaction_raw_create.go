package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"task2/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// CreateRawTransaction 创建一个原始交易
func CreateRawTransaction() error {
	log.Println("=== 创建原始交易演示 ===")

	// 获取客户端连接
	client, err := InitClient()
	if err != nil {
		return fmt.Errorf("连接以太坊网络失败: %v", err)
	}
	log.Println("成功连接到以太坊网络")

	network := config.GetCurrentNetwork()

	// 获取私钥
	privateKeyBytes, err := hexutil.Decode(network.PrivateKey)
	if err != nil {
		return fmt.Errorf("解析私钥失败: %v", err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return fmt.Errorf("创建私钥对象失败: %v", err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("无法获取公钥")
	}

	// 从公钥获取发送方地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("发送方地址: %s", fromAddress.Hex())

	// 获取nonce值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("获取nonce失败: %v", err)
	}
	log.Printf("当前 nonce: %d", nonce)

	// 设置交易参数
	value := big.NewInt(1000000000000000000) // 1 ETH
	log.Printf("发送金额: %s Wei", value.String())

	gasLimit := uint64(21000) // 标准ETH转账的gas limit
	log.Printf("Gas 限制: %d", gasLimit)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("获取gas价格失败: %v", err)
	}
	log.Printf("Gas 价格: %s Wei", gasPrice.String())

	// 接收方地址（使用第二个测试账户地址）
	toAddress := common.HexToAddress(AccountAddress2)
	log.Printf("接收方地址: %s", toAddress.Hex())

	// 创建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// 获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("获取链ID失败: %v", err)
	}
	log.Printf("当前网络: %s (Chain ID: %d)", network.NetworkName, chainID)

	// 签名交易
	signer := network.GetSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return fmt.Errorf("签名交易失败: %v", err)
	}
	log.Println("交易签名成功")

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("发送交易失败: %v", err)
	}

	log.Printf("交易已发送，交易哈希: %s", signedTx.Hash().Hex())
	// 等待交易被打包确认
	log.Println("等待交易被打包确认...")
	// 广播交易到网络
	txHash := signedTx.Hash()
	log.Printf("广播交易，交易哈希: %s", txHash.Hex())

	// 创建一个10秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return fmt.Errorf("等待交易确认超时，交易哈希: %s", signedTx.Hash().Hex())
		}
		return fmt.Errorf("等待交易确认失败: %v", err)
	}

	log.Printf("交易已确认，区块号: %d", receipt.BlockNumber)
	log.Printf("Gas 使用量: %d", receipt.GasUsed)
	if receipt.Status == 1 {
		log.Println("交易执行成功")
	} else {
		log.Println("交易执行失败")
	}
	return nil
}
