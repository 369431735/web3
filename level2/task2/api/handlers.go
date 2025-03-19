package api

import (
	"net/http"
	"task2/config"
	"task2/utils"

	"github.com/gin-gonic/gin"
)

// GetContractAddresses 获取所有合约地址
func GetContractAddresses(c *gin.Context) {
	network := config.GetCurrentNetwork()
	if network == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "未找到网络配置",
		})
		return
	}

	contractAddresses := make(map[string]string)
	for name, contract := range network.Contracts {
		contractAddresses[name] = contract.Address
	}

	c.JSON(http.StatusOK, contractAddresses)
}

// GetNetworkInfo 获取网络信息
func GetNetworkInfo(c *gin.Context) {
	network := config.GetCurrentNetwork()
	if network == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "未找到网络配置",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    network.NetworkName,
		"chainId": network.ChainID,
		"nodeUrl": network.RPCURL,
	})
}

// GetBalance 获取账户余额
func GetBalance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "地址参数不能为空",
		})
		return
	}

	balance, err := utils.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"balance": balance.String(),
	})
}

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
