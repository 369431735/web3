package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var account_ys string = "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"
var url string = "http://localhost:8545"

func main() {
	adress()
	yue()
}

// 初始化客户端
func client() (*ethclient.Client, error) {
	runtime.Breakpoint() // 添加断点
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	return client, err
}

// 地址
func adress() {
	// 将十六进制字符串转换为地址
	address := common.HexToAddress(account_ys)

	// 输出地址的不同表示形式
	fmt.Println("Hex:", address.Hex())     // 十六进制表示
	fmt.Println("Bytes:", address.Bytes()) // 字节表示

	// 将地址转换为哈希
	var hash common.Hash
	hash.SetBytes(address.Bytes())
	fmt.Println("Hash:", hash.Hex()) // 哈希表示
}

// 余额
func yue() {
	fmt.Println("=== fmt.Println 的输出 ===")

	// 连接到本地 Ganache 网络
	client, err := client()
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}

	fmt.Println("成功连接到本地以太坊网络")

	// 使用 Ganache 提供的第一个账户
	account := common.HexToAddress(account_ys)
	fmt.Printf("正在查询账户: %s\n", account.Hex())

	fmt.Printf("正在获取当前余额...")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalf("获取当前余额失败: %v", err)
	}
	fmt.Printf("当前余额: %s wei\n", balance.String())

	fmt.Printf("正在获取最新区块信息...")
	// 获取最新区块号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("获取区块头失败: %v", err)
	}
	blockNumber := header.Number
	fmt.Printf("当前区块号: %s\n", blockNumber.String())

	log.Printf("正在获取区块 %s 的余额...\n", blockNumber.String())
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatalf("获取指定区块余额失败: %v", err)
	}
	fmt.Printf("区块 %s 的余额: %s wei\n", blockNumber.String(), balanceAt.String())

	log.Println("正在转换余额为 ETH...")
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("转换为 ETH: %v ETH\n", ethValue)

	fmt.Printf("正在获取待处理余额...")
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatalf("获取待处理余额失败: %v", err)
	}
	fmt.Printf("待处理余额: %s wei\n", pendingBalance.String())

	fmt.Printf("程序执行完成!")
}
