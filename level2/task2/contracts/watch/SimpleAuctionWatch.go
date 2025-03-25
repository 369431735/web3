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
	eventChanBuffer    = 2000
	baseRetryDelay     = 2 * time.Second
	maxRetryDelay      = 30 * time.Second
	reconnectThreshold = 3
	eventSendTimeout   = 500 * time.Millisecond
)

type AuctionWatcher struct {
	client       *ethclient.Client
	contract     *bindings.SimpleAuction
	currentBlock *big.Int
	sub          ethereum.Subscription
	subActive    bool
	mu           sync.RWMutex
	ctx          context.Context
	cancelCtx    context.CancelFunc
	eventChan    chan *bindings.SimpleAuctionHighestBidIncreased
}

// NewAuctionWatcher 创建监听器（带入口日志）
func NewAuctionWatcher(contractAddr string) (*AuctionWatcher, error) {
	log.Printf("[METHOD] NewAuctionWatcher 入口，合约地址: %s", contractAddr)

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

// Close 释放资源（带入口日志）
func (w *AuctionWatcher) Close() {
	log.Printf("[METHOD] Close 入口")

	w.cancelCtx()
	w.closeSubscription()
}

// WatchHighestBidIncreased 启动监听（带入口日志）
func (w *AuctionWatcher) WatchHighestBidIncreased() error {
	log.Printf("[METHOD] WatchHighestBidIncreased 入口")

	go w.watchEvents()
	return nil
}

// watchEvents 核心循环（带入口日志）
func (w *AuctionWatcher) watchEvents() {
	log.Printf("[METHOD] watchEvents 入口")

	defer func() {
		w.closeSubscription()
		close(w.eventChan)
	}()

	retryCount := 0

	for {
		select {
		case <-w.ctx.Done():
			log.Printf("[METHOD] watchEvents 退出")
			return
		default:
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

// isSubscriptionActive 检查订阅状态（带入口日志）
func (w *AuctionWatcher) isSubscriptionActive() bool {
	log.Printf("[METHOD] isSubscriptionActive 入口")

	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.sub != nil && w.subActive
}

// createAndHandleSubscription 创建订阅（带入口日志）
func (w *AuctionWatcher) createAndHandleSubscription(retryCount *int) error {
	log.Printf("[METHOD] createAndHandleSubscription 入口，重试次数: %d", *retryCount)

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
	w.subActive = true
	*retryCount = 0

	go w.monitorSubscription(sub)
	go w.processEvents(sink)
	return nil
}

// monitorSubscription 监控订阅（带入口日志）
func (w *AuctionWatcher) monitorSubscription(sub ethereum.Subscription) {
	log.Printf("[METHOD] monitorSubscription 入口")

	err := <-sub.Err()

	w.mu.Lock()
	w.subActive = false
	w.mu.Unlock()

	if err != nil {
		log.Printf("[WARN] 订阅异常: %v", err)
	}
}

// processEvents 处理事件流（带入口日志）
func (w *AuctionWatcher) processEvents(sink <-chan *bindings.SimpleAuctionHighestBidIncreased) {
	log.Printf("[METHOD] processEvents 入口")

	for {
		select {
		case <-w.ctx.Done():
			log.Printf("[METHOD] processEvents 退出")
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

// handleEvent 处理单个事件（带入口日志）
func (w *AuctionWatcher) handleEvent(event *bindings.SimpleAuctionHighestBidIncreased) {
	log.Printf("[METHOD] handleEvent 入口，区块: %d", event.Raw.BlockNumber)

	w.updateBlockNumber(event.Raw.BlockNumber)

	log.Printf("[EVENT] 收到事件: 区块=%d 出价人=%s 金额=%s wei",
		event.Raw.BlockNumber,
		event.Bidder.Hex(),
		event.Amount.String(),
	)

	select {
	case w.eventChan <- event:
	case <-time.After(eventSendTimeout):
		log.Printf("[WARN] 事件通道已满，丢弃区块 %d 的事件", event.Raw.BlockNumber)
	case <-w.ctx.Done():
	}
}

// updateBlockNumber 更新区块号（带入口日志）
func (w *AuctionWatcher) updateBlockNumber(newBlock uint64) {
	log.Printf("[METHOD] updateBlockNumber 入口，新区块: %d", newBlock)

	w.mu.Lock()
	defer w.mu.Unlock()
	if newBlock > w.currentBlock.Uint64() {
		w.currentBlock = new(big.Int).SetUint64(newBlock)
	}
}

// getStartBlock 获取起始区块（带入口日志）
func (w *AuctionWatcher) getStartBlock() *big.Int {
	log.Printf("[METHOD] getStartBlock 入口")
	return w.currentBlock
}

// handleRetry 处理重试（带入口日志）
func (w *AuctionWatcher) handleRetry(retryCount *int) {
	log.Printf("[METHOD] handleRetry 入口，当前重试次数: %d", *retryCount)

	if *retryCount >= reconnectThreshold {
		delay := w.calculateBackoff(*retryCount)
		log.Printf("[WARN] 达到重试阈值，等待 %v 后重试", delay)
		time.Sleep(delay)
	} else {
		time.Sleep(baseRetryDelay)
	}
	*retryCount += 1
}

// closeSubscription 关闭订阅（带入口日志）
func (w *AuctionWatcher) closeSubscription() {
	log.Printf("[METHOD] closeSubscription 入口")
	if w.sub != nil {
		w.sub.Unsubscribe()
		w.sub = nil
		w.subActive = false
	}
}

// calculateBackoff 计算退避时间（带入口日志）
func (w *AuctionWatcher) calculateBackoff(attempt int) time.Duration {
	log.Printf("[METHOD] calculateBackoff 入口，尝试次数: %d", attempt)

	delay := baseRetryDelay * time.Duration(1<<uint(attempt))
	if delay > maxRetryDelay {
		return maxRetryDelay
	}
	return delay
}

// safeUint64FromBig 转换大整数（带入口日志）
func safeUint64FromBig(i *big.Int) (uint64, error) {
	log.Printf("[METHOD] safeUint64FromBig 入口，输入值: %s", i.String())

	if i == nil {
		return 0, errors.New("nil 指针异常")
	}
	if !i.IsUint64() {
		return 0, fmt.Errorf("数值溢出: %s > uint64 最大值", i.String())
	}
	return i.Uint64(), nil
}
