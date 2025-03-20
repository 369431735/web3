package utils

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sync"

	"task2/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
)

var (
	// HTTP客户端实例
	httpClientInstance *ethclient.Client
	httpClientMutex    sync.Mutex
	httpInitialized    bool

	// WebSocket客户端实例
	wsClientInstance *ethclient.Client
	wsClientMutex    sync.Mutex
	wsInitialized    bool

	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Be careful with this in production
		},
	}
)

// GetConfig 获取全局配置
func GetConfig() *config.Config {
	return config.GetConfig()
}

// GetCurrentNetwork 获取当前网络配置
func GetCurrentNetwork() *config.NetworkConfig {
	return config.GetCurrentNetwork()
}

// 如果客户端尚未初始化，则会初始化一个新的HTTP客户端
func GetEthClientHTTP() (*ethclient.Client, error) {
	if httpInitialized && httpClientInstance != nil {
		return httpClientInstance, nil
	}

	// 加锁确保只有一个goroutine可以初始化客户端
	httpClientMutex.Lock()
	defer httpClientMutex.Unlock()

	// 双重检查，防止在等待锁期间已有其他goroutine完成了初始化
	if httpInitialized && httpClientInstance != nil {
		return httpClientInstance, nil
	}

	// 获取网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return nil, fmt.Errorf("未找到网络配置")
	}

	// 确保RPCURL已配置
	if network.RPCURL == "" {
		return nil, fmt.Errorf("未配置HTTP RPC URL")
	}

	// 尝试HTTP连接
	client, err := ethclient.Dial(network.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP连接失败: %v", err)
	}

	log.Println("成功使用HTTP连接以太坊网络:", network.RPCURL)
	httpClientInstance = client
	httpInitialized = true
	return client, nil
}

// GetEthClientWS 返回使用WebSocket连接的以太坊客户端单例
// 如果客户端尚未初始化，则会初始化一个新的WebSocket客户端
func GetEthClientWS() (*ethclient.Client, error) {
	if wsInitialized && wsClientInstance != nil {
		return wsClientInstance, nil
	}

	// 加锁确保只有一个goroutine可以初始化客户端
	wsClientMutex.Lock()
	defer wsClientMutex.Unlock()

	// 双重检查，防止在等待锁期间已有其他goroutine完成了初始化
	if wsInitialized && wsClientInstance != nil {
		return wsClientInstance, nil
	}

	// 获取网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return nil, fmt.Errorf("未找到网络配置")
	}

	// 确保WSURL已配置
	if network.WSURL == "" {
		return nil, fmt.Errorf("未配置WebSocket URL")
	}

	// 尝试WebSocket连接
	client, err := ethclient.Dial(network.WSURL)
	if err != nil {
		return nil, fmt.Errorf("WebSocket连接失败: %v", err)
	}

	log.Println("成功使用WebSocket连接以太坊网络:", network.WSURL)
	wsClientInstance = client
	wsInitialized = true
	return client, nil
}

// ResetEthClientHTTP 重置HTTP客户端实例
func ResetEthClientHTTP() {
	httpClientMutex.Lock()
	defer httpClientMutex.Unlock()

	if httpClientInstance != nil {
		httpClientInstance.Close()
		httpClientInstance = nil
	}
	httpInitialized = false
}

// ResetEthClientWS 重置WebSocket客户端实例
func ResetEthClientWS() {
	wsClientMutex.Lock()
	defer wsClientMutex.Unlock()

	if wsClientInstance != nil {
		wsClientInstance.Close()
		wsClientInstance = nil
	}
	wsInitialized = false
}

// ResetEthClient 重置所有客户端连接
func ResetEthClient() {
	ResetEthClientHTTP()
	ResetEthClientWS()
}

// SetAccountBalance 设置账户余额
func SetAccountBalance(address string) error {
	// 这里不需要获取客户端，因为当前功能只打印日志

	// 设置默认余额为 1 ETH
	value := new(big.Int)
	value.SetString("1000000000000000000", 10) // 1 ETH

	// TODO: 实现设置账户余额的逻辑
	log.Printf("设置账户 %s 的余额为 1 ETH", address)
	return nil
}
