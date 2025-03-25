package api

import (
	"task2/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 设置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 提供docs目录的静态文件服务，直接使用swag生成的文件
	r.Static("/docs", "./docs")

	// API文档 - 使用docs目录中的swagger.json
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/docs/swagger.json")))

	// 区块相关路由
	blocks := r.Group("/block")
	{
		blocks.GET("/latest", controller.GetLatestBlock)     // 获取最新区块信息
		blocks.GET("/:number", controller.GetBlockByNumber)  // 通过区块号获取区块信息
		blocks.GET("/hash/:hash", controller.GetBlockByHash) // 通过区块哈希获取区块信息
		blocks.GET("/info", controller.GetBlockInfo)         // 获取区块详细信息

	}

	// 账户相关路由
	accounts := r.Group("/accounts")
	{
		accounts.GET("/:address/balance", controller.GetAccountBalance)           // 获取账户余额
		accounts.GET("/:address/transactions", controller.GetAccountTransactions) // 获取账户交易历史
		accounts.GET("/:address/nonce", controller.GetAccountNonce)               // 获取账户nonce值
		accounts.GET("/:address/code", controller.GetAccountCode)                 // 获取账户合约代码
	}

	// 钱包相关路由
	wallet := r.Group("/account")
	{
		wallet.POST("/wallet", controller.CreateWallet)     // 创建新的以太坊钱包
		wallet.POST("/keystore", controller.CreateKeystore) // 创建新的以太坊密钥库文件
		wallet.POST("/hdwallet", controller.CreateHDWallet) // 创建新的分层确定性钱包
	}

	// 交易相关路由
	transactions := r.Group("/transaction")
	{
		transactions.POST("/create", controller.CreateTransaction) // 创建并发送交易
		transactions.POST("/raw", controller.CreateRawTransaction) // 创建原始交易

	}

	// 合约相关路由
	contracts := r.Group("/contracts")
	{
		contracts.POST("/deploy-all", controller.DeployAllContracts)    // 部署所有合约
		contracts.GET("/allAddresses", controller.GetContractAddresses) // 获取所有已部署合约地址
		contracts.POST("/bytecode", controller.GetContractBytecode)     // 获取合约字节码

		// SimpleStorage合约方法
		simpleStorage := contracts.Group("/SimpleStorage")
		{
			simpleStorage.POST("/set", controller.SimpleStorageSet) // 设置SimpleStorage合约的存储值
			simpleStorage.GET("/get", controller.SimpleStorageGet)  // 获取SimpleStorage合约的存储值
		}

		// Lock合约方法
		lock := contracts.Group("/Lock")
		{
			lock.POST("/withdraw", controller.LockWithdraw)

		}

		// Shipping合约方法
		shipping := contracts.Group("/Shipping")
		{
			shipping.POST("/advance-state", controller.ShippingAdvanceState)
			shipping.GET("/get-state", controller.ShippingGetState)
		}

		// SimpleAuction合约方法
		simpleAuction := contracts.Group("/SimpleAuction")
		{
			simpleAuction.POST("/bid", controller.SimpleAuctionBid)
			simpleAuction.POST("/withdraw", controller.SimpleAuctionWithdraw)
			simpleAuction.POST("/watchHighestBidIncreased", controller.WatchHighestBidIncreased)
		}

		// ArrayDemo合约方法
		arrayDemo := contracts.Group("/arraydemo")
		{
			arrayDemo.POST("/add-value", controller.ArrayDemoAddValue)
			arrayDemo.GET("/get-values", controller.ArrayDemoGetValues)
		}

		// Ballot合约方法
		ballot := contracts.Group("/Ballot")
		{
			ballot.POST("/vote", controller.BallotVote)
			ballot.GET("/winner", controller.BallotWinningProposal)
			ballot.GET("/winner-name", controller.BallotWinnerName)
		}

		// Lottery合约方法
		lottery := contracts.Group("/Lottery")
		{
			lottery.POST("/enter", controller.LotteryEnter)
			lottery.POST("/draw", controller.LotteryPickWinner)
			lottery.GET("/players", controller.LotteryGetPlayers)
			lottery.GET("/balance", controller.LotteryGetBalance)
		}

		// Purchase合约方法
		purchase := contracts.Group("/Purchase")
		{
			purchase.POST("/abort", controller.PurchaseAbort)
			purchase.POST("/confirm", controller.PurchaseConfirmPurchase)
			purchase.POST("/confirm-received", controller.PurchaseConfirmReceived)
			purchase.GET("/value", controller.PurchaseGetValue)
			purchase.GET("/state", controller.PurchaseGetState)
		}
	}

	// 事件相关路由
	events := r.Group("/events")
	{
		events.GET("/subscribe", controller.SubscribeContractEvents) // 通过WebSocket订阅合约事件
	}

	return r
}
