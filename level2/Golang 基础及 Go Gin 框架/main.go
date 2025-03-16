package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	accountAddress = "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"
	nodeURL        = "http://localhost:8545"
)

func main() {
	log.Println("=== 程序开始执行 ===")
	//address()
	//balance()
	newWallet()
	log.Println("=== 程序执行完成 ===")
}

// 初始化以太坊客户端
func initClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}
	log.Println("成功连接到本地以太坊网络")
	return client, nil
}

// 地址转换演示
func address() {
	log.Println("=== 地址转换演示 ===")
	// 将十六进制字符串转换为地址
	address := common.HexToAddress(accountAddress)

	// 输出地址的不同表示形式
	log.Printf("十六进制表示: %s", address.Hex())
	log.Printf("字节表示: %v", address.Bytes())

	// 将地址转换为哈希
	var hash common.Hash
	hash.SetBytes(address.Bytes())
	log.Printf("哈希表示: %s", hash.Hex())
}

// 余额查询演示
func balance() {
	log.Println("=== 余额查询演示 ===")

	// 连接到本地以太坊网络
	client, err := initClient()
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}

	// 查询账户信息
	account := common.HexToAddress(accountAddress)
	log.Printf("查询账户: %s", account.Hex())

	// 获取当前余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalf("获取当前余额失败: %v", err)
	}
	log.Printf("当前余额: %s wei", balance.String())

	// 获取最新区块信息
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("获取区块头失败: %v", err)
	}
	blockNumber := header.Number
	log.Printf("当前区块号: %s", blockNumber.String())

	// 获取指定区块的余额
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatalf("获取指定区块余额失败: %v", err)
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
		log.Fatalf("获取待处理余额失败: %v", err)
	}
	log.Printf("待处理余额: %s wei", pendingBalance.String())
}

// 创建新钱包
func newWallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
