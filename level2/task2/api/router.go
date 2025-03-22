package api

import (
	"task2/config"
	"task2/controller"
	"task2/docs"

	_ "task2/docs" // 导入 swagger 文档

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// --------------------------
// Swagger 主配置
// --------------------------
// @title           Web3 区块链接口服务
// @version         1.0
// @description     提供以太坊区块链相关的API服务，包括账户管理、交易处理、合约部署等功能
// @host           localhost:8080
// @BasePath       /api/v1
// @schemes        http
// @contact.name   API Support
// @contact.email  support@example.com
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html

// --------------------------
// 模型定义
// --------------------------
type (
	// AccountBalance 账户余额请求参数
	AccountBalance struct {
		Address string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e" binding:"required"`
	}

	// AccountResponse 账户余额响应
	AccountResponse struct {
		Address string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
		Balance string `json:"balance" example:"1000000000000000000"`
	}

	// ErrorResponse 错误响应
	ErrorResponse struct {
		Code    int    `json:"code" example:"400"`
		Message string `json:"message" example:"Invalid address format"`
	}
)

// --------------------------
// 路由定义
// --------------------------

// @Summary      健康检查
// @Description  检查服务是否正常运行
// @Tags         系统
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health [get]
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 获取配置
	cfg := config.GetConfig()
	basePath := cfg.Server.BasePath

	// 设置 Swagger 文档基础路径
	docs.SwaggerInfo.BasePath = basePath
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API 路由组
	v1 := r.Group(basePath)
	{
		// 健康检查
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		// 网络相关路由
		v1.GET("/network", controller.GetNetworkInfo)

		// 账户相关路由
		account := v1.Group("/account")
		{
			account.GET("/balance", controller.GetAccountBalance)
			account.POST("/balance", controller.SetAccountBalance)
			account.POST("/wallet", controller.CreateWallet)
			account.POST("/keystore", controller.CreateKeystore)
		}

		// 区块相关路由
		block := v1.Group("/block")
		{
			block.GET("/info", controller.GetBlockInfo)
			block.GET("/latest", controller.GetLatestBlock)
			block.GET("/subscribe", controller.SubscribeBlocks)
		}

		// 合约相关路由
		contracts := v1.Group("/contracts")
		{
			contracts.POST("/deploy-all", controller.DeployAllContracts)
			contracts.GET("/allAddresses", controller.GetContractAddresses)
		}

		// 单个合约部署路由
		contract := v1.Group("/contract")
		{
			// 合约部署路由
			contract.POST("/bytecode", controller.GetContractBytecode)

			// SimpleStorage合约方法
			simplestorage := contract.Group("/simplestorage")
			{
				simplestorage.POST("/set", controller.SimpleStorageSet)
				simplestorage.GET("/get", controller.SimpleStorageGet)
			}

			// Lock合约方法
			lock := contract.Group("/lock")
			{
				lock.POST("/withdraw", controller.LockWithdraw)
			}

			// SimpleAuction合约方法
			simpleauction := contract.Group("/simpleauction")
			{
				simpleauction.POST("/bid", controller.SimpleAuctionBid)
				simpleauction.POST("/withdraw", controller.SimpleAuctionWithdraw)
			}

			// Shipping合约方法
			shipping := contract.Group("/shipping")
			{
				shipping.POST("/advance-state", controller.ShippingAdvanceState)
				shipping.GET("/get-state", controller.ShippingGetState)
			}

			// ArrayDemo合约方法
			arraydemo := contract.Group("/arraydemo")
			{
				arraydemo.POST("/add-value", controller.ArrayDemoAddValue)
				arraydemo.GET("/get-values", controller.ArrayDemoGetValues)
			}
		}

		// 交易相关路由
		transaction := v1.Group("/transaction")
		{
			transaction.POST("/create", controller.CreateTransaction)
			transaction.POST("/raw", controller.CreateRawTransaction)
		}

		// 事件相关路由
		events := v1.Group("/events")
		{
			events.POST("/subscribe", controller.SubscribeContractEvents)
		}
	}

	// 保留一个根路径健康检查（用于负载均衡器/监控系统）
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return r
}
