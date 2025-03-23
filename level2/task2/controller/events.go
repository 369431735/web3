package controller

import (
	"context"
	"net/http"
	"task2/events"
	"task2/storage"
	"task2/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

// SubscribeContractEvents godoc
// @Summary      订阅合约事件
// @Description  通过WebSocket连接订阅合约事件，实时接收事件数据
// @Tags         事件
// @Accept       application/json
// @Produce      application/json
// @Success      101  {string}  string  "Switching Protocols to websocket"
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /events/subscribe [get]
func SubscribeContractEvents(c *gin.Context) {
	client, err := utils.GetEthClientWS()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket连接创建失败: " + err.Error()})
		return
	}
	// 获取合约地址列表
	contractStorage := storage.GetInstance()
	contractAddresses := contractStorage.GetAllAddresses()
	if len(contractAddresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有可用的合约地址，请先部署合约"})
		return
	}

	query := ethereum.FilterQuery{
		Addresses: make([]common.Address, 0, len(contractAddresses)),
	}

	for _, addr := range contractAddresses {
		address := common.HexToAddress(addr)
		query.Addresses = append(query.Addresses, address)
	}

	logs := make(chan ethTypes.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "订阅合约事件失败: " + err.Error()})
		return
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				utils.LogError("合约事件订阅出错", err)
				return
			case vLog := <-logs:
				// 处理接收到的事件日志，根据Topic识别事件类型
				if len(vLog.Topics) > 0 {
					eventSignature := vLog.Topics[0]

					// 查找并调用对应的事件处理器
					if handler, ok := events.EventHandlers[eventSignature]; ok {
						// 使用单独的goroutine处理事件，避免阻塞主事件循环
						go handler.Handle(vLog)
					} else {
						utils.LogInfo("未知事件签名", map[string]interface{}{
							"eventSignature": eventSignature.Hex(),
							"contract":       vLog.Address.Hex(),
						})
					}
				}
			}
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "合约事件订阅已启动成功",
	})
}
