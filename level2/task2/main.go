package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"task2/events"
	"time"

	"task2/api"
	"task2/config"
	"task2/storage"
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
// @swagger         2.0

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
	// 连接到以太坊节点
	client, err := utils.GetEthClientHTTP()
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

// 捕获崩溃并记录到日志
func recoverAndLog() {
	if r := recover(); r != nil {
		// 获取堆栈信息
		buf := make([]byte, 8192)
		n := runtime.Stack(buf, false)
		stackInfo := string(buf[:n])

		// 记录详细错误信息
		log.Printf("程序发生致命错误: %v\n堆栈信息:\n%s", r, stackInfo)

		// 尝试将错误信息同时写入到独立的崩溃日志文件
		crashLogFile := "crash_" + time.Now().Format("20060102_150405") + ".log"
		crashLog := fmt.Sprintf("时间: %s\n错误: %v\n堆栈信息:\n%s\n",
			time.Now().Format("2006-01-02 15:04:05"), r, stackInfo)

		// 忽略写入崩溃日志的错误，不能因为写入日志失败而导致程序无法处理其他错误
		_ = os.WriteFile(crashLogFile, []byte(crashLog), 0644)

		// 在控制台打印错误信息以便于调试
		fmt.Printf("程序崩溃: %v\n请查看日志文件获取详细信息\n", r)
	}
}

func main() {
	// 设置基本的日志格式
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("=== 程序开始执行 ===")

	// 使用defer函数捕获所有panic并记录到日志
	defer recoverAndLog()

	// 加载配置文件
	if err := config.LoadConfig(); err != nil {
		log.Printf("加载配置文件失败: %v", err)
		os.Exit(1)
	}
	log.Println("配置文件加载成功")

	// 获取日志配置
	cfg := config.GetConfig()
	if cfg == nil {
		log.Printf("错误: 无法获取配置，配置对象为nil")
		os.Exit(1)
	}

	// 设置日志输出到文件
	if cfg.Log.Filename != "" {
		log.Printf("正在配置日志文件: %s", cfg.Log.Filename)

		// 确保日志目录存在
		logDir := filepath.Dir(cfg.Log.Filename)
		if logDir != "." {
			if err := os.MkdirAll(logDir, 0755); err != nil {
				log.Printf("创建日志目录失败: %v", err)
			} else {
				log.Printf("日志目录已创建/确认: %s", logDir)
			}
		}

		// 打开日志文件
		logFile, err := os.OpenFile(cfg.Log.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("打开日志文件失败: %v", err)
		} else {
			// 设置日志输出为多写入器，同时输出到控制台和文件
			multiWriter := io.MultiWriter(os.Stdout, logFile)
			log.SetOutput(multiWriter)
			log.Printf("日志已成功重定向到文件和控制台: %s", cfg.Log.Filename)
		}
	} else {
		log.Printf("警告: 未配置日志文件名，将使用标准输出")
	}

	// 初始化合约地址
	contractStorage := storage.GetInstance()
	log.Println("合约地址已从存储中加载")
	for key, value := range contractStorage.GetAllAddresses() {
		events.InitializeEventHandlersByAdress(key, value)
	}
	// 初始化事件处理器

	// 初始化路由
	log.Println("初始化路由和API服务")

	// 创建路由实例
	router := api.SetupRouter()
	if router == nil {
		log.Printf("错误: 路由初始化失败，router对象为nil")
		return
	}

	// 获取服务器配置
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
