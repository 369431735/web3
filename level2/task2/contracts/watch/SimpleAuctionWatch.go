package watch

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sync"
	"task2/utils"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"task2/contracts/bindings" // 替换为你的合约绑定路径
)

const (
	maxRetries         = 5
	baseRetryDelay     = 1 * time.Second
	maxRetryDelay      = 30 * time.Second
	eventChanBuffer    = 1000
	reconnectThreshold = 3
)

// AuctionWatcher 封装拍卖监控逻辑
type AuctionWatcher struct {
	client       *ethclient.Client
	contract     *bindings.SimpleAuction
	currentBlock *big.Int
	sink         chan *bindings.SimpleAuctionHighestBidIncreased
	mu           sync.RWMutex
}

// NewAuctionWatcher 创建新的监控实例
func NewAuctionWatcher(contractAddr string) (*AuctionWatcher, error) {
	client, err := utils.GetEthClientWS()
	if err != nil {
		return nil, fmt.Errorf("连接节点失败: %w", err)
	}

	contract, err := bindings.NewSimpleAuction(common.HexToAddress(contractAddr), client)
	if err != nil {
		return nil, fmt.Errorf("初始化合约失败: %w", err)
	}

	// 获取当前区块号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("获取初始区块失败: %w", err)
	}

	return &AuctionWatcher{
		client:       client,
		contract:     contract,
		currentBlock: header.Number,
		sink:         make(chan *bindings.SimpleAuctionHighestBidIncreased, eventChanBuffer),
	}, nil
}

// WatchHighestBidIncreased 启动事件监听
func (w *AuctionWatcher) WatchHighestBidIncreased(ctx context.Context, eventChan chan<- *bindings.SimpleAuctionHighestBidIncreased) error {
	var reconnectCount int
	retryDelay := baseRetryDelay

	for {
		select {
		case <-ctx.Done():
			close(eventChan)
			return nil

		default:
			sub, err := w.createSubscription(ctx)
			if err != nil {
				log.Printf("订阅失败: %v", err)

				if reconnectCount >= reconnectThreshold {
					retryDelay = w.calculateBackoff(reconnectCount)
				}
				select {
				case <-time.After(retryDelay):
					reconnectCount++
					continue
				case <-ctx.Done():
					return nil
				}
			}

			reconnectCount = 0
			retryDelay = baseRetryDelay

			if err := w.handleSubscription(ctx, sub, eventChan); err != nil {
				return fmt.Errorf("处理订阅时发生错误: %w", err)
			}
		}
	}
}

// 创建事件订阅
func (w *AuctionWatcher) createSubscription(ctx context.Context) (ethereum.Subscription, error) {
	startBlock := w.getStartBlock()

	// 转换为安全的uint64
	blockUint64, err := safeUint64FromBig(startBlock)
	if err != nil {
		return nil, fmt.Errorf("区块号转换失败: %w", err)
	}

	opts := &bind.WatchOpts{
		Context: ctx,
		Start:   &blockUint64,
	}

	sub, err := w.contract.WatchHighestBidIncreased(opts, w.sink)
	if err != nil {
		return nil, fmt.Errorf("创建订阅失败: %w", err)
	}

	return sub, nil
}

// 处理订阅事件
func (w *AuctionWatcher) handleSubscription(
	ctx context.Context,
	sub ethereum.Subscription,
	eventChan chan<- *bindings.SimpleAuctionHighestBidIncreased,
) error {
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case err := <-sub.Err():
			if errors.Is(err, ethereum.NotFound) {
				return fmt.Errorf("订阅丢失: %w", err)
			}
			return fmt.Errorf("订阅错误: %w", err)

		case event := <-w.sink:
			w.updateBlockNumber(event.Raw.BlockNumber)
			w.sendEvent(event, eventChan)
		}
	}
}

// 更新处理的最高区块号
func (w *AuctionWatcher) updateBlockNumber(newBlock uint64) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if newBlock > w.currentBlock.Uint64() {
		w.currentBlock = new(big.Int).SetUint64(newBlock)
	}
}

// 获取安全的起始区块号
func (w *AuctionWatcher) getStartBlock() *big.Int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.currentBlock
}

// 非阻塞发送事件
func (w *AuctionWatcher) sendEvent(event *bindings.SimpleAuctionHighestBidIncreased, ch chan<- *bindings.SimpleAuctionHighestBidIncreased) {
	select {
	case ch <- event:
	default:
		log.Printf("事件通道已满，丢弃事件 (当前大小: %d/%d)", len(ch), cap(ch))
	}
}

// 计算指数退避时间
func (w *AuctionWatcher) calculateBackoff(attempt int) time.Duration {
	delay := baseRetryDelay * time.Duration(1<<attempt)
	if delay > maxRetryDelay {
		return maxRetryDelay
	}
	return delay
}

// 安全转换big.Int到uint64
func safeUint64FromBig(i *big.Int) (uint64, error) {
	if i == nil {
		return 0, errors.New("nil指针异常")
	}
	if !i.IsUint64() {
		return 0, fmt.Errorf("数值溢出: %s > uint64最大值", i.String())
	}
	return i.Uint64(), nil
}

// 使用示例
func main() {
	// 配置信息
	const (
		rpcURL       = "wss://mainnet.infura.io/ws/v3/YOUR_PROJECT_ID"
		contractAddr = "0xYourContractAddress"
	)

	// 初始化监控器
	watcher, err := NewAuctionWatcher(contractAddr)
	if err != nil {
		log.Fatal("初始化失败: ", err)
	}

	// 创建事件通道
	eventChan := make(chan *bindings.SimpleAuctionHighestBidIncreased, eventChanBuffer)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动监听协程
	go func() {
		if err := watcher.WatchHighestBidIncreased(ctx, eventChan); err != nil {
			log.Fatal("监听异常: ", err)
		}
	}()

	// 处理事件
	for {
		select {
		case event := <-eventChan:
			handleEvent(event)
		case <-ctx.Done():
			return
		}
	}
}

func handleEvent(event *bindings.SimpleAuctionHighestBidIncreased) {
	fmt.Printf("[New Bid] Block:%-6d Bidder:%-42s Amount:%s Wei\n",
		event.Raw.BlockNumber,
		event.Bidder.Hex(),
		event.Amount.String(),
	)
}
