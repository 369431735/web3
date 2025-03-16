package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

const (
	accountAddress = "0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0"
	nodeURL        = "http://localhost:8545"
	privateKey     = "0x6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("=== 程序开始执行 ===")
	address_cul()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("程序发生错误: %v", r)
		}
	}()

	log.Println("\n1. 地址转换演示")
	address()

	log.Println("\n2. 余额查询演示")
	if err := balance(); err != nil {
		log.Printf("余额查询失败: %v", err)
	}

	log.Println("\n3. 创建新钱包")
	if err := newWallet(); err != nil {
		log.Printf("创建钱包失败: %v", err)
	}

	log.Println("\n4. 创建 Keystore")
	if err := createKs(); err != nil {
		log.Printf("创建 Keystore 失败: %v", err)
	}

	log.Println("\n5. HD 钱包演示")
	if err := chdwallet(); err != nil {
		log.Printf("HD 钱包操作失败: %v", err)
	}

	log.Println("\n6. 地址检查演示")
	if err := address_check(); err != nil {
		log.Printf("地址检查失败: %v", err)
	}

	log.Println("\n=== 程序执行完成 ===")
}

// 根据私钥计算地址
func address_cul() error {
	log.Println("=== 根据私钥计算地址 ===")

	// 将私钥字符串转换为字节
	privateKeyBytes, err := hexutil.Decode(privateKey)
	if err != nil {
		return fmt.Errorf("私钥解码失败: %v", err)
	}

	// 从字节创建私钥对象
	privateKeyECDSA, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return fmt.Errorf("创建私钥对象失败: %v", err)
	}

	// 从私钥获取公钥
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("转换公钥类型失败")
	}

	// 从公钥计算地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("计算得到的地址: %s", address.Hex())

	return nil
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
func balance() error {
	log.Println("=== 余额查询演示 ===")

	// 连接到本地以太坊网络
	client, err := initClient()
	if err != nil {
		return fmt.Errorf("连接失败: %v", err)
	}

	// 查询账户信息
	account := common.HexToAddress(accountAddress)
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

// 创建新钱包
func newWallet() error {
	log.Println("=== 创建新钱包演示 ===")

	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("生成私钥失败: %v", err)
	}

	// 获取私钥的十六进制表示
	privateKeyBytes := crypto.FromECDSA(privateKey)
	log.Printf("私钥: %s", hexutil.Encode(privateKeyBytes)[2:])

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("无法将公钥转换为 ECDSA 格式")
	}

	// 获取公钥的十六进制表示
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	log.Printf("公钥: %s", hexutil.Encode(publicKeyBytes)[4:])

	// 从公钥生成以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("地址: %s", address.Hex())

	// 使用 Keccak-256 哈希函数计算地址
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	log.Printf("通过哈希计算的地址: %s", hexutil.Encode(hash.Sum(nil)[12:]))

	return nil
}

func createKs() error {
	log.Println("=== 创建 Keystore 演示 ===")

	// 创建 keystore 目录
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	log.Println("成功创建 Keystore 目录")

	// 创建新账户
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		return fmt.Errorf("创建账户失败: %v", err)
	}
	log.Printf("成功创建新账户，地址: %s", account.Address.Hex())

	return nil
}

func chdwallet() error {
	log.Println("=== HD 钱包演示 ===")

	// 使用助记词创建 HD 钱包
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return fmt.Errorf("创建 HD 钱包失败: %v", err)
	}
	log.Println("成功创建 HD 钱包")

	// 派生第一个地址
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return fmt.Errorf("派生第一个地址失败: %v", err)
	}
	log.Printf("第一个派生地址: %s", account.Address.Hex())

	// 派生第二个地址
	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		return fmt.Errorf("派生第二个地址失败: %v", err)
	}
	log.Printf("第二个派生地址: %s", account.Address.Hex())

	return nil
}

// 检查地址是否有效
func address_check() error {
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
	client, err := ethclient.Dial(nodeURL)
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

// 查询区块信息
func block_info() error {
	log.Println("=== 区块信息查询演示 ===")

	// 连接到以太坊网络
	client, err := initClient()
	if err != nil {
		return fmt.Errorf("连接以太坊网络失败: %v", err)
	}
	// 检查客户端连接状态
	if client == nil {
		return fmt.Errorf("以太坊客户端未连接")
	}

	// 检查网络ID
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("获取网络ID失败: %v", err)
	}
	log.Printf("当前网络ID: %d", networkID)

	// 检查同步状态
	sync, err := client.SyncProgress(context.Background())
	if err != nil {
		return fmt.Errorf("获取同步状态失败: %v", err)
	}
	if sync != nil {
		log.Printf("节点正在同步中，当前区块: %d，最高区块: %d", sync.CurrentBlock, sync.HighestBlock)
	} else {
		log.Println("节点已完成同步")
	}

	// 获取最新区块号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("获取区块头失败: %v", err)
	}
	log.Printf("当前最新区块号: %d", header.Number.Uint64())

	// 获取区块详细信息
	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		return fmt.Errorf("获取区块信息失败: %v", err)
	}

	// 输出区块信息
	log.Printf("区块哈希: %s", block.Hash().Hex())
	log.Printf("父区块哈希: %s", block.ParentHash().Hex())
	log.Printf("区块中交易数量: %d", len(block.Transactions()))
	log.Printf("区块时间戳: %v", time.Unix(int64(block.Time()), 0))
	log.Printf("区块难度: %s", block.Difficulty().String())
	log.Printf("区块Gas限制: %d", block.GasLimit())
	log.Printf("区块Gas使用量: %d", block.GasUsed())

	return nil
}
