package utils

import (
	"context"
	"fmt"

	"task2/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// InitClient 初始化以太坊客户端
func InitClient() (*ethclient.Client, error) {
	network := config.GetCurrentNetwork()
	client, err := ethclient.Dial(network.NodeURL)
	if err != nil {
		return nil, fmt.Errorf("连接以太坊节点失败: %v", err)
	}
	return client, nil
}

// SetAccountBalance 设置账户余额
func SetAccountBalance() error {
	client, err := InitClient()
	if err != nil {
		return err
	}
	defer client.Close()

	network := config.GetCurrentNetwork()
	cfg := config.GetConfig()

	// 获取私钥
	privateKey, err := crypto.HexToECDSA(network.PrivateKey)
	if err != nil {
		return fmt.Errorf("解析私钥失败: %v", err)
	}

	// 获取链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("获取链ID失败: %v", err)
	}

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("创建交易选项失败: %v", err)
	}

	// 设置余额
	auth.Value = cfg.Accounts.DefaultBalance

	return nil
}
