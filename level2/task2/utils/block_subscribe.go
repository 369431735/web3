package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockSubscription 区块订阅结构体
type BlockSubscription struct {
	client       *ethclient.Client  // 以太坊客户端连接
	ticker       *time.Ticker       // 定时器，用于定期检查新区块
	lastBlockNum uint64             // 最后一个已处理的区块号
	isSubscribed bool               // 是否正在订阅状态
	callback     func(*types.Block) // 新区块的回调函数
	ctx          context.Context    // 上下文，用于取消订阅
	cancel       context.CancelFunc // 取消函数
}

// NewBlockSubscription 创建新的区块订阅实例
func NewBlockSubscription() (*BlockSubscription, error) {
	client, err := InitClient()
	if err != nil {
		return nil, fmt.Errorf("连接以太坊网络失败: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &BlockSubscription{
		client:       client,
		isSubscribed: false,
		ctx:          ctx,
		cancel:       cancel,
	}, nil
}

// Subscribe 开始订阅新区块
// interval: 轮询间隔（秒）
// callback: 新区块回调函数
func (bs *BlockSubscription) Subscribe(interval int, callback func(*types.Block)) error {
	if bs.isSubscribed {
		return fmt.Errorf("已经在订阅区块")
	}

	log.Println("=== 开始订阅新区块 ===")

	// 获取初始区块号
	lastBlock, err := bs.client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("获取最新区块失败: %v", err)
	}
	bs.lastBlockNum = lastBlock.Number().Uint64()
	log.Printf("初始区块号: %d", bs.lastBlockNum)
	bs.callback = callback

	// 开始轮询新区块，使用100毫秒的间隔
	bs.ticker = time.NewTicker(100 * time.Millisecond)
	bs.isSubscribed = true

	go bs.startPolling()

	return nil
}

// startPolling 开始轮询新区块
func (bs *BlockSubscription) startPolling() {
	log.Println("开始轮询新区块...")
	for {
		select {
		case <-bs.ctx.Done():
			log.Println("订阅已停止，退出轮询")
			return
		case <-bs.ticker.C:
			if !bs.isSubscribed {
				log.Println("订阅已取消，退出轮询")
				return
			}

			// 获取最新区块
			block, err := bs.client.BlockByNumber(context.Background(), nil)
			if err != nil {
				log.Printf("获取区块失败: %v", err)
				continue
			}

			currentBlockNum := block.Number().Uint64()

			// 如果区块号大于上一个区块号，说明有新区块
			if currentBlockNum > bs.lastBlockNum {
				log.Printf("发现新区块: %d (上一个区块: %d)", currentBlockNum, bs.lastBlockNum)

				// 获取区块中的所有交易
				txCount := len(block.Transactions())
				log.Printf("区块 %d 包含 %d 笔交易", currentBlockNum, txCount)

				// 打印区块信息
				log.Printf("区块时间戳: %v", time.Unix(int64(block.Time()), 0))
				log.Printf("区块哈希: %s", block.Hash().Hex())
				log.Printf("父区块哈希: %s", block.ParentHash().Hex())
				log.Printf("区块Gas使用量: %d", block.GasUsed())

				// 更新最新区块号
				bs.lastBlockNum = currentBlockNum

				// 调用回调函数
				if bs.callback != nil {
					bs.callback(block)
				}
			}
		}
	}
}

// Unsubscribe 取消订阅
func (bs *BlockSubscription) Unsubscribe() {
	if bs.isSubscribed {
		bs.isSubscribed = false
		bs.ticker.Stop()
		bs.cancel()
		log.Println("已取消订阅新区块")
	}
}

// GetLastBlockNumber 获取最新区块号
func (bs *BlockSubscription) GetLastBlockNumber() uint64 {
	return bs.lastBlockNum
}

// IsSubscribed 检查是否正在订阅
func (bs *BlockSubscription) IsSubscribed() bool {
	return bs.isSubscribed
}
