package main

import (
	"context"
	"fmt"
	"log"

	"task2/api"
	"task2/config"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// @title           Web3 区块链接口服务
// @version         1.0
// @description     提供以太坊区块链相关的API服务，包括账户管理、交易处理、合约部署等功能
// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http
// @contact.name    API Support
// @contact.email   support@example.com
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

/*
config.go - 网络配置相关
client.go - 客户端初始化和基础操作
wallet.go - 钱包相关操作
transaction.go - 交易相关操作
block.go - 区块相关操作
address.go - 地址相关操作
block_subscribe.go - 区块订阅相关操作
*/

// AccountBalance 账户余额请求模型

// 初始化以太坊客户端和账户
func setupClient() (*ethclient.Client, *bind.TransactOpts, error) {
	// 连接到本地节点
	client, err := utils.InitClient()
	if err != nil {
		return nil, nil, fmt.Errorf("连接到以太坊节点失败: %v", err)
	}

	// 获取网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return nil, nil, fmt.Errorf("未找到网络配置")
	}

	// 使用默认账户
	defaultAccount, ok := network.Accounts["default"]
	if !ok {
		return nil, nil, fmt.Errorf("未找到默认账户")
	}

	// 解析私钥（处理可能的"0x"前缀）
	privateKeyStr := defaultAccount.PrivateKey
	if len(privateKeyStr) >= 2 && privateKeyStr[:2] == "0x" {
		privateKeyStr = privateKeyStr[2:] // 移除"0x"前缀
	}

	privateKey, err := utils.GetPrivateKey(privateKeyStr)
	if err != nil {
		return nil, nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	// 创建交易选项
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	return client, auth, nil
}

// 部署所有合约的函数已移动到abi包中
// 请参考abi包中的相关实现

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("=== 程序开始执行 ===")
	defer func() {
		if r := recover(); r != nil {
			log.Printf("程序发生错误: %v", r)
		}
	}()

	// 加载配置文件
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
	log.Println("配置文件加载成功")

	// 初始化路由
	log.Println("\n1. 初始化路由和API服务")

	// 创建路由实例
	router := api.SetupRouter()

	// 注意：区块订阅功能已移至路由中处理
	// 在路由初始化过程中会创建区块订阅实例并注册相关处理程序
	log.Println("路由初始化完成，区块订阅功能已集成")

	// 注意：实际的订阅逻辑将在API处理程序中实现
	// 客户端可以通过HTTP请求触发订阅功能

	// 获取服务器配置
	cfg := config.GetConfig()
	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	// 设置 Gin 模式
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Printf("启动 HTTP 服务器在 %s", addr)
	log.Printf("服务器模式: %s", cfg.Server.Mode)
	log.Printf("Swagger 文档地址: http://localhost%s/swagger/index.html", addr)

	// 启动 HTTP 服务器
	if err := router.Run(addr); err != nil {
		log.Printf("启动 HTTP 服务器失败: %v", err)
		return
	}
}
