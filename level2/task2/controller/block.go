package controller

import (
	"net/http"
	"task2/utils"

	"github.com/gin-gonic/gin"
)

// GetBlockInfo 获取区块信息
func GetBlockInfo(c *gin.Context) {
	if err := utils.BlockInfo(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "区块信息获取成功",
	})
}

// CreateTransaction 创建并发送交易
func CreateTransaction(c *gin.Context) {
	txHash, err := utils.CreateAndSendTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "交易发送成功",
		"txHash":  txHash,
	})
}

// CreateRawTransaction 创建原始交易
func CreateRawTransaction(c *gin.Context) {
	if err := utils.CreateRawTransaction(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "原始交易创建成功",
	})
}
