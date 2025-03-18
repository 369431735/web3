package api

import (
	"task2/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 健康检查
	r.GET("/health", HealthCheck)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 账户相关
		v1.POST("/account/balance", controller.SetAccountBalance)
		v1.GET("/account/balance", controller.GetBalance)
		v1.POST("/account/wallet", controller.CreateWallet)
		v1.POST("/account/keystore", controller.CreateKeystore)
		v1.POST("/account/hdwallet", controller.CreateHDWallet)

		// 区块相关
		v1.GET("/block/info", controller.GetBlockInfo)
		v1.POST("/transaction", controller.CreateTransaction)
		v1.POST("/transaction/raw", controller.CreateRawTransaction)

		// 合约相关
		v1.POST("/contracts/deploy", controller.DeployContracts)

		// 地址相关
		v1.GET("/address/check", controller.CheckAddress)

		// 获取所有合约地址
		v1.GET("/contracts", GetContractAddresses)

		// 获取网络信息
		v1.GET("/network", GetNetworkInfo)
	}

	return r
}
