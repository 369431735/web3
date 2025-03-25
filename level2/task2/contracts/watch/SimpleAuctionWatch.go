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
	"task2/contracts/bindings"
)

const (
	eventChanBuffer    = 2000 // 增大事件通道缓冲区
	baseRetryDelay     = 2 * time.Second
	maxRetryDelay      = 30 * time.Second
	reconnectThreshold = 3
	eventSendTimeout   = 500 * time.Millisecond // 新增事件发送超时
)

type AuctionWatcher struct {
	client       *ethclient.Client
	contract     *bindings.SimpleAuction
	currentBlock *big.Int
	sub          ethereum.Subscription
	mu           sync.RWMutex
	ctx          context.Context
	cancelCtx    context.CancelFunc
	eventChan    chan *bindings.SimpleAuctionHighestBidIncreased
}

func NewAuctionWatcher(contractAddr string) (*AuctionWatcher, error) {
	client, err := utils.GetEthClientWS()
	if err != nil {
		return nil, fmt.Errorf("连接节点失败: %w", err)
	}

	contract, err := bindings.NewSimpleAuction(common.HexToAddress(contractAddr), client)
	if err != nil {
		return nil, fmt.Errorf("初始化合约失败: %w", err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("获取初始区块失败: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &AuctionWatcher{
		client:       client,
		contract:     contract,
		currentBlock: header.Number,
		ctx:          ctx,
		cancelCtx:    cancel,
		eventChan:    make(chan *bindings.SimpleAuctionHighestBidIncreased, eventChanBuffer),
	}, nil
}

// 新增关闭方法确保资源释放
func (w *AuctionWatcher) Close() {
	w.cancelCtx()
	w.closeSubscription()
}

func (w *AuctionWatcher) WatchHighestBidIncreased() error {
	go w.watchEvents()
	return nil
}

func (w *AuctionWatcher) watchEvents() {
	defer func() {
		w.closeSubscription()
		close(w.eventChan) // 唯一关闭通道的位置
	}()

	retryCount := 0

	for {
		select {
		case <-w.ctx.Done():
			return
		default:
			// 加锁检查订阅状态
			if w.isSubscriptionActive() {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			if err := w.createAndHandleSubscription(&retryCount); err != nil {
				log.Printf("[ERROR] 订阅失败: %v", err)
				w.handleRetry(&retryCount)
			} else {
				retryCount = 0
			}
		}
	}
}

// 封装订阅状态检查（加锁）
func (w *AuctionWatcher) isSubscriptionActive() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.sub != nil && w.sub.Err() == nil
}

func (w *AuctionWatcher) createAndHandleSubscription(retryCount *int) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.closeSubscription()

	startBlock := w.getStartBlock()
	blockUint64, err := safeUint64FromBig(startBlock)
	if err != nil {
		return fmt.Errorf("区块号转换失败: %w", err)
	}

	opts := &bind.WatchOpts{
		Context: w.ctx,
		Start:   &blockUint64,
	}

	sink := make(chan *bindings.SimpleAuctionHighestBidIncreased, eventChanBuffer)
	sub, err := w.contract.WatchHighestBidIncreased(opts, sink)
	if err != nil {
		return fmt.Errorf("创建订阅失败: %w", err)
	}

	w.sub = sub
	*retryCount = 0

	go w.processEvents(sink)
	return nil
}

func (w *AuctionWatcher) processEvents(sink <-chan *bindings.SimpleAuctionHighestBidIncreased) {
	for {
		select {
		case <-w.ctx.Done():
			return
		case event, ok := <-sink:
			if !ok {
				log.Println("[INFO] 事件源通道关闭")
				return
			}
			w.handleEvent(event)
		}
	}
}

// 改进事件处理（增加超时机制）
func (w *AuctionWatcher) handleEvent(event *bindings.SimpleAuctionHighestBidIncreased) {
	w.updateBlockNumber(event.Raw.BlockNumber)

	select {
	case w.eventChan <- event:
		for event1 := range w.eventChan {
			log.Printf("[EVENT] 收到事件: 区块=%d 出价人=%s 金额=%s",
				event1.Raw.BlockNumber,
				event1.Bidder.Hex(),
				event1.Amount.String(),
			)
		}
	case <-time.After(eventSendTimeout):
		log.Printf("[WARN] 发送事件超时，丢弃区块 %d 的事件", event.Raw.BlockNumber)
	case <-w.ctx.Done():
	}
}

func (w *AuctionWatcher) updateBlockNumber(newBlock uint64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if newBlock > w.currentBlock.Uint64() {
		w.currentBlock = new(big.Int).SetUint64(newBlock)
	}
}

func (w *AuctionWatcher) getStartBlock() *big.Int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.currentBlock
}

// 改进重试逻辑（统一退避处理）
func (w *AuctionWatcher) handleRetry(retryCount *int) {
	if *retryCount >= reconnectThreshold {
		delay := w.calculateBackoff(*retryCount)
		log.Printf("[WARN] 达到重试阈值，等待 %v 后重试", delay)
		time.Sleep(delay)
	} else {
		time.Sleep(baseRetryDelay)
	}
	*retryCount++
}

func (w *AuctionWatcher) closeSubscription() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.sub != nil {
		w.sub.Unsubscribe()
		w.sub = nil
	}
}

func (w *AuctionWatcher) calculateBackoff(attempt int) time.Duration {
	delay := baseRetryDelay * time.Duration(1<<uint(attempt))
	if delay > maxRetryDelay {
		return maxRetryDelay
	}
	return delay
}

func safeUint64FromBig(i *big.Int) (uint64, error) {
	if i == nil {
		return 0, errors.New("nil 指针异常")
	}
	if !i.IsUint64() {
		return 0, fmt.Errorf("数值溢出: %s > uint64 最大值", i.String())
	}
	return i.Uint64(), nil
}
