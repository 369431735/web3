package api

import (
	"net/http"
	"task2/config"
	"task2/utils"

	"github.com/gin-gonic/gin"
)

// GetContractAddresses 获取所有合约地址
func GetContractAddresses(c *gin.Context) {
	cfg := config.GetConfig()
	c.JSON(http.StatusOK, gin.H{
		"simpleStorage": cfg.Contracts.SimpleStorageAddress,
		"lock":          cfg.Contracts.LockAddress,
		"shipping":      cfg.Contracts.ShippingAddress,
		"simpleAuction": cfg.Contracts.SimpleAuctionAddress,
		"purchase":      cfg.Contracts.PurchaseAddress,
	})
}

// GetNetworkInfo 获取网络信息
func GetNetworkInfo(c *gin.Context) {
	network := config.GetCurrentNetwork()
	c.JSON(http.StatusOK, gin.H{
		"name":    network.NetworkName,
		"chainId": network.ChainID,
		"nodeUrl": network.NodeURL,
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
