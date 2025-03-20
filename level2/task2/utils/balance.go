package utils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ... rest of the code ...

// GetBalance 获取指定地址的余额
func GetBalance(address string) (*big.Int, error) {
	client, err := GetEthClientHTTP()
	if err != nil {
		return nil, fmt.Errorf("连接到以太坊节点失败: %v", err)
	}
	// 使用的是单例客户端，不需要关闭

	// 将地址字符串转换为以太坊地址
	ethAddress := common.HexToAddress(address)

	// 获取当前余额
	balance, err := client.BalanceAt(context.Background(), ethAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("获取余额失败: %v", err)
	}

	return balance, nil
}

// ... rest of the code ...
