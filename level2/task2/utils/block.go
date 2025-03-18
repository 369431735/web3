package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogError 记录错误日志
func LogError(message string, err error) {
	log.Printf("错误: %s - %v", message, err)
}

// LogInfo 记录信息日志
func LogInfo(message string, data map[string]interface{}) {
	log.Printf("信息: %s - %+v", message, data)
}

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

// SubscribeNewBlock 订阅新区块
func SubscribeNewBlock(client *ethclient.Client) error {
	headers := make(chan *ethTypes.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Printf("订阅错误: %v", err)
			case header := <-headers:
				block, err := client.BlockByHash(context.Background(), header.Hash())
				if err != nil {
					log.Printf("获取区块错误: %v", err)
					continue
				}
				log.Printf("新区块: 区块号=%v, 时间戳=%v, 交易数=%v",
					block.Number().String(),
					block.Time(),
					len(block.Transactions()))
			}
		}
	}()

	return nil
}

// SubscribeContractEvents 订阅合约事件
func SubscribeContractEvents(client *ethclient.Client, contracts map[string]common.Address) {
	for name, address := range contracts {
		go subscribeContractEvents(client, name, address)
	}
}

// subscribeContractEvents 订阅单个合约的事件
func subscribeContractEvents(client *ethclient.Client, name string, address common.Address) {
	// 创建事件过滤器
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	// 订阅日志
	logs := make(chan ethTypes.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		LogError("订阅合约事件失败", err)
		return
	}

	// 处理事件
	for {
		select {
		case err := <-sub.Err():
			LogError("合约事件订阅错误", err)
		case vLog := <-logs:
			// 记录事件信息
			LogInfo("合约事件", map[string]interface{}{
				"contract": name,
				"address":  address.Hex(),
				"topics":   vLog.Topics,
				"data":     vLog.Data,
			})
		}
	}
}

// ContractEvent 合约事件结构
type ContractEvent struct {
	ContractName string
	EventName    string
	Address      common.Address
	Topics       []common.Hash
	Data         []byte
}
