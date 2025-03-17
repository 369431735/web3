package utils

import (
	"context"
	"fmt"
	"log"
	"time"
)

// 查询区块信息
func BlockInfo() error {
	log.Println("=== 区块信息查询演示 ===")

	// 连接到以太坊网络
	client, err := InitClient()
	if err != nil {
		return fmt.Errorf("连接以太坊网络失败: %v", err)
	}
	// 检查客户端连接状态
	if client == nil {
		return fmt.Errorf("以太坊客户端未连接")
	}

	// 检查网络ID
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("获取网络ID失败: %v", err)
	}
	log.Printf("当前网络ID: %d", networkID)

	// 检查同步状态
	sync, err := client.SyncProgress(context.Background())
	if err != nil {
		return fmt.Errorf("获取同步状态失败: %v", err)
	}
	if sync != nil {
		log.Printf("节点正在同步中，当前区块: %d，最高区块: %d", sync.CurrentBlock, sync.HighestBlock)
	} else {
		log.Println("节点已完成同步")
	}

	// 获取最新区块号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("获取区块头失败: %v", err)
	}
	log.Printf("当前最新区块号: %d", header.Number.Uint64())

	// 获取区块详细信息
	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		return fmt.Errorf("获取区块信息失败: %v", err)
	}

	// 输出区块信息
	log.Printf("区块哈希: %s", block.Hash().Hex())
	log.Printf("父区块哈希: %s", block.ParentHash().Hex())
	log.Printf("区块中交易数量: %d", len(block.Transactions()))
	log.Printf("区块时间戳: %v", time.Unix(int64(block.Time()), 0))
	log.Printf("区块难度: %s", block.Difficulty().String())
	log.Printf("区块Gas限制: %d", block.GasLimit())
	log.Printf("区块Gas使用量: %d", block.GasUsed())

	return nil
}
