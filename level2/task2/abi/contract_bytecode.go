package abi

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetContractBytecode 获取智能合约字节码
func GetContractBytecode(client *ethclient.Client, contractAddress string) ([]byte, error) {
	// 验证地址格式
	if !common.IsHexAddress(contractAddress) {
		return nil, fmt.Errorf("无效的合约地址格式")
	}

	// 转换为以太坊地址格式
	address := common.HexToAddress(contractAddress)

	// 获取字节码
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return nil, fmt.Errorf("获取合约字节码失败: %v", err)
	}

	// 如果字节码长度为0，说明该地址不是合约地址
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("该地址不是合约地址")
	}

	return bytecode, nil
}
