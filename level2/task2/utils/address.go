package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"regexp"
	"strings"
	"task2/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 常量定义
const (
	AccountAddress = "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
)

// 根据私钥计算地址
func AddressCul() error {
	log.Println("=== 根据私钥计算地址 ===")

	// 从配置文件获取私钥
	network := config.GetCurrentNetwork()
	if network == nil {
		return fmt.Errorf("未找到网络配置")
	}

	account := network.Accounts["default"]
	privateKey, err := GetPrivateKey(account.PrivateKey)
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

// GetPrivateKey 从十六进制字符串获取私钥
func GetPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	log.Printf("处理私钥...")

	// 检查私钥格式并移除0x前缀
	if strings.HasPrefix(privateKeyHex, "0x") {
		log.Printf("检测到0x前缀，将被移除")
		privateKeyHex = privateKeyHex[2:]
	}

	// 检查私钥长度
	if len(privateKeyHex) != 64 {
		return nil, fmt.Errorf("私钥长度不正确: %d (应为64个字符)", len(privateKeyHex))
	}

	// 验证是否为有效的十六进制字符串
	for _, c := range privateKeyHex {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return nil, fmt.Errorf("私钥包含无效字符: %c", c)
		}
	}

	// 转换为ECDSA私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("私钥转换失败: %v", err)
		return nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	log.Printf("私钥处理成功")
	return privateKey, nil
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
	client, err := GetEthClientHTTP()
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
	client, err := GetEthClientHTTP()
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

// GetAccountAddress 获取账户地址
func GetAccountAddress(accountName string) (string, error) {
	network := config.GetCurrentNetwork()
	if network == nil {
		return "", fmt.Errorf("未找到网络配置")
	}

	account, ok := network.Accounts[accountName]
	if !ok {
		return "", fmt.Errorf("未找到账户: %s", accountName)
	}

	return account.Address, nil
}

// GetAccountPrivateKey 获取账户私钥
func GetAccountPrivateKey(accountName string) (string, error) {
	network := config.GetCurrentNetwork()
	if network == nil {
		return "", fmt.Errorf("未找到网络配置")
	}

	account, ok := network.Accounts[accountName]
	if !ok {
		return "", fmt.Errorf("未找到账户: %s", accountName)
	}

	return account.PrivateKey, nil
}
