package controller

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"task2/events"
	"task2/storage"
	"task2/utils"
)

// SubscribeContractEvents godoc
// @Summary      订阅合约事件
// @Description  开始监听和处理合约事件
// @Tags         事件
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /events/subscribe [post]
func SubscribeContractEvents(c *gin.Context) {
	client, err := utils.GetEthClientWS()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket连接失败: " + err.Error()})
		return
	}
	// 获取合约存储实例
	contractStorage := storage.GetInstance()
	contractAddresses := contractStorage.GetAllAddresses()
	if len(contractAddresses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有找到已部署的合约"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "订阅事件失败: " + err.Error()})
		return
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				utils.LogError("合约事件订阅出错", err)
				return
			case vLog := <-logs:
				// 获取事件签名（第一个topic就是事件签名）
				if len(vLog.Topics) > 0 {
					eventSignature := vLog.Topics[0]

					// 查找对应的事件处理器
					if handler, ok := events.EventHandlers[eventSignature]; ok {
						// 在新的goroutine中处理事件，避免阻塞主事件循环
						go handler.Handle(vLog)
					} else {
						utils.LogInfo("未找到事件处理器", map[string]interface{}{
							"eventSignature": eventSignature.Hex(),
							"contract":       vLog.Address.Hex(),
						})
					}
				}
			}
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "合约事件订阅已启动",
	})
}
