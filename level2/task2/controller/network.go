package controller

import (
	"task2/config"

	"github.com/gin-gonic/gin"
)

// NetworkInfo 网络信息
type NetworkInfo struct {
	Name    string `json:"name" example:"Sepolia"`
	ChainID int64  `json:"chainId" example:"11155111"`
	NodeURL string `json:"nodeUrl" example:"https://sepolia.infura.io/v3/..."`
}

// GetNetworkInfo godoc
// @Summary      获取网络信息
// @Description  获取当前连接的网络信息
// @Tags         网络
// @Accept       json
// @Produce      json
// @Success      200  {object}  NetworkInfo
// @Router       /network [get]
func GetNetworkInfo(c *gin.Context) {
	// 从配置中获取当前网络信息
	networkConfig := config.GetCurrentNetwork()
	if networkConfig == nil {
		c.JSON(500, gin.H{"error": "无法获取网络配置"})
		return
	}

	info := NetworkInfo{
		Name:    networkConfig.NetworkName,
		ChainID: networkConfig.ChainID,
		NodeURL: networkConfig.NodeURL,
	}
	c.JSON(200, info)
}
