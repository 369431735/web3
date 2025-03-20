package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"task2/contracts"
	"task2/utils"
)

// EventController 处理与区块链事件相关的请求
type EventController struct{}

// 订阅事件请求的参数
type SubscribeParams struct {
	Address string `json:"address" binding:"required"`
}

// SubscribeToEvents 处理订阅事件的HTTP请求
func (e *EventController) SubscribeToEvents(c *gin.Context) {
	var params SubscribeParams
	contractType := c.Param("contractType")

	// 尝试从请求体获取合约地址
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}

	// 验证合约地址
	if !common.IsHexAddress(params.Address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的合约地址"})
		return
	}

	// 连接到以太坊节点
	client, err := ConnectToEthereum()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法连接到以太坊网络: " + err.Error()})
		return
	}

	// 根据合约类型订阅相应的事件
	switch contractType {
	case "lock":
		go subscribeToLockEvents(client, common.HexToAddress(params.Address))
		c.JSON(http.StatusOK, gin.H{"message": "成功订阅Lock合约事件"})
	case "simplestorage":
		go subscribeToSimpleStorageEvents(client, common.HexToAddress(params.Address))
		c.JSON(http.StatusOK, gin.H{"message": "成功订阅SimpleStorage合约事件"})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的合约类型"})
	}
}

// ConnectToEthereum 连接到以太坊网络
// 专门针对事件订阅，使用WebSocket连接
func ConnectToEthereum() (*ethclient.Client, error) {
	// 对于事件订阅，必须使用WebSocket连接
	return utils.GetEthClientWS()
}

// 订阅Lock合约的事件
func subscribeToLockEvents(client *ethclient.Client, address common.Address) {
	// 初始化Lock合约实例
	lockContract, err := contracts.NewLock(address, client)
	if err != nil {
		log.Println("无法初始化Lock合约:", err)
		return
	}

	// 创建用于接收事件的通道
	withdrawalCh := make(chan *contracts.LockWithdrawal)

	// 设置过滤器选项
	opts := &bind.WatchOpts{
		Context: context.Background(),
	}

	// 订阅Withdrawal事件
	sub, err := lockContract.WatchWithdrawal(opts, withdrawalCh)
	if err != nil {
		log.Println("无法订阅Lock合约的Withdrawal事件:", err)
		return
	}
	defer sub.Unsubscribe()

	// 持续监听事件
	for {
		select {
		case err := <-sub.Err():
			log.Println("Lock事件订阅出错:", err)
			return
		case event := <-withdrawalCh:
			// 记录提款事件
			eventJSON, _ := json.Marshal(map[string]interface{}{
				"contractAddress": address.Hex(),
				"eventType":       "Withdrawal",
				"amount":          event.Amount.String(),
				"when":            event.When.String(),
				"timestamp":       time.Now().Format(time.RFC3339),
			})
			log.Printf("收到Lock合约事件: %s\n", string(eventJSON))
		}
	}
}

// 订阅SimpleStorage合约的事件
func subscribeToSimpleStorageEvents(client *ethclient.Client, address common.Address) {
	// 初始化SimpleStorage合约实例
	ssContract, err := contracts.NewSimpleStorage(address, client)
	if err != nil {
		log.Println("无法初始化SimpleStorage合约:", err)
		return
	}

	// 创建用于接收事件的通道
	valueCh := make(chan *contracts.SimpleStorageValueChanged)

	// 设置过滤器选项
	opts := &bind.WatchOpts{
		Context: context.Background(),
	}

	// 订阅ValueChanged事件
	sub, err := ssContract.WatchValueChanged(opts, valueCh)
	if err != nil {
		log.Println("无法订阅SimpleStorage合约的ValueChanged事件:", err)
		return
	}
	defer sub.Unsubscribe()

	// 持续监听事件
	for {
		select {
		case err := <-sub.Err():
			log.Println("SimpleStorage事件订阅出错:", err)
			return
		case event := <-valueCh:
			// 记录值变更事件
			eventJSON, _ := json.Marshal(map[string]interface{}{
				"contractAddress": address.Hex(),
				"eventType":       "ValueChanged",
				"value":           event.Value.String(),
				"timestamp":       time.Now().Format(time.RFC3339),
			})
			log.Printf("收到SimpleStorage合约事件: %s\n", string(eventJSON))
		}
	}
}
