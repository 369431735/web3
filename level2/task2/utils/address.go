package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 常量定义
const (
	PrivateKey2     = "0x6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1"
	AccountAddress  = "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"
	AccountAddress2 = "0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0"
)

// 根据私钥计算地址
func AddressCul() error {
	log.Println("=== 根据私钥计算地址 ===")

	// 将私钥字符串转换为字节
	privateKey, err := GetPrivateKey(PrivateKey2)
	if err != nil {
		return fmt.Errorf("私钥解码失败: %v", err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("转换公钥类型失败")
	}

	// 从公钥计算地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("计算得到的地址: %s", address.Hex())

	return nil
}

// 地址转换演示
func Address() {
	log.Println("=== 地址转换演示 ===")
	// 将十六进制字符串转换为地址
	address := common.HexToAddress(AccountAddress)

	// 输出地址的不同表示形式
	log.Printf("十六进制表示: %s", address.Hex())
	log.Printf("字节表示: %v", address.Bytes())

	// 将地址转换为哈希
	var hash common.Hash
	hash.SetBytes(address.Bytes())
	log.Printf("哈希表示: %s", hash.Hex())
}

// 余额查询演示
func Balance() error {
	log.Println("=== 余额查询演示 ===")

	// 连接到本地以太坊网络
	client, err := InitClient()
	if err != nil {
		return fmt.Errorf("连接失败: %v", err)
	}

	// 查询账户信息
	account := common.HexToAddress(AccountAddress)
	log.Printf("查询账户: %s", account.Hex())

	// 获取当前余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return fmt.Errorf("获取当前余额失败: %v", err)
	}
	log.Printf("当前余额: %s wei", balance.String())

	// 获取最新区块信息
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("获取区块头失败: %v", err)
	}
	blockNumber := header.Number
	log.Printf("当前区块号: %s", blockNumber.String())

	// 获取指定区块的余额
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		return fmt.Errorf("获取指定区块余额失败: %v", err)
	}
	log.Printf("区块 %s 的余额: %s wei", blockNumber.String(), balanceAt.String())

	// 转换为 ETH
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	log.Printf("ETH 余额: %v ETH", ethValue)

	// 获取待处理余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		return fmt.Errorf("获取待处理余额失败: %v", err)
	}
	log.Printf("待处理余额: %s wei", pendingBalance.String())

	return nil
}

// 检查地址是否有效
func AddressCheck() error {
	log.Println("=== 地址检查演示 ===")
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	// 检查有效地址
	validAddr := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	log.Printf("检查地址 %s", validAddr)
	log.Printf("格式是否有效: %v", re.MatchString(validAddr))

	// 检查无效地址
	invalidAddr := "0xZYXb5d4c32345ced77393b3530b1eed0f346429d"
	log.Printf("检查地址 %s", invalidAddr)
	log.Printf("格式是否有效: %v", re.MatchString(invalidAddr))

	// 连接到以太坊网络
	client, err := InitClient()
	if err != nil {
		return fmt.Errorf("连接以太坊网络失败: %v", err)
	}

	// 检查智能合约地址
	contractAddr := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	log.Printf("检查合约地址: %s", contractAddr.Hex())
	bytecode, err := client.CodeAt(context.Background(), contractAddr, nil)
	if err != nil {
		return fmt.Errorf("获取合约代码失败: %v", err)
	}
	isContract := len(bytecode) > 0
	log.Printf("是否是合约地址: %v", isContract)

	// 检查普通用户地址
	userAddr := common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	log.Printf("检查用户地址: %s", userAddr.Hex())
	bytecode, err = client.CodeAt(context.Background(), userAddr, nil)
	if err != nil {
		return fmt.Errorf("获取用户地址代码失败: %v", err)
	}
	isContract = len(bytecode) > 0
	log.Printf("是否是合约地址: %v", isContract)

	return nil
}
