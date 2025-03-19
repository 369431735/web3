package utils

import (
	"fmt"
	"log"
	"math/big"

	"task2/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

// GetConfig 获取全局配置
func GetConfig() *config.Config {
	return config.GetConfig()
}

// GetCurrentNetwork 获取当前网络配置
func GetCurrentNetwork() *config.NetworkConfig {
	return config.GetCurrentNetwork()
}

// InitClient 初始化以太坊客户端
func InitClient() (*ethclient.Client, error) {
	network := config.GetCurrentNetwork()
	if network == nil {
		return nil, fmt.Errorf("未找到网络配置")
	}

	// 连接到以太坊节点
	client, err := ethclient.Dial(network.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("连接以太坊节点失败: %v", err)
	}

	return client, nil
}

// SetAccountBalance 设置账户余额
func SetAccountBalance(address string) error {
	client, err := InitClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// 设置默认余额为 1 ETH
	value := new(big.Int)
	value.SetString("1000000000000000000", 10) // 1 ETH

	// TODO: 实现设置账户余额的逻辑
	log.Printf("设置账户 %s 的余额为 1 ETH", address)
	return nil
}
