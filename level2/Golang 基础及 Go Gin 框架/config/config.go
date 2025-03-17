package config

import (
	"math/big"
	"os"
	"strconv"

	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// NetworkConfig 网络环境配置
type NetworkConfig struct {
	ChainID      *big.Int
	NetworkName  string
	NodeURL      string
	IsTestnet    bool
	ABIPath      string
	BindingsPath string
	PrivateKey   string
	GetSigner    func(*big.Int) types.Signer
}

// ContractConfig 合约配置
type ContractConfig struct {
	LockValue     *big.Int
	PurchaseValue *big.Int
	AuctionTime   int64
	LockTime      int64
}

// AccountConfig 账户配置
type AccountConfig struct {
	DefaultBalance *big.Int
	Addresses      map[string]common.Address
	PrivateKeys    map[string]string
}

// Config 全局配置
type Config struct {
	Networks   map[string]NetworkConfig
	Contracts  ContractConfig
	Accounts   AccountConfig
	CurrentEnv string
}

var (
	// 预定义网络配置
	Networks = map[string]NetworkConfig{
		"mainnet": {
			ChainID:      big.NewInt(1),
			NetworkName:  "Mainnet",
			NodeURL:      "https://mainnet.infura.io/v3/YOUR-PROJECT-ID",
			IsTestnet:    false,
			ABIPath:      "./artifacts/contracts",
			BindingsPath: "abi/bindings",
			PrivateKey:   os.Getenv("ETH_MAINNET_PRIVATE_KEY"),
			GetSigner:    func(chainID *big.Int) types.Signer { return types.NewLondonSigner(chainID) },
		},
		"sepolia": {
			ChainID:      big.NewInt(11155111),
			NetworkName:  "Sepolia",
			NodeURL:      "https://sepolia.infura.io/v3/YOUR-PROJECT-ID",
			IsTestnet:    true,
			ABIPath:      "./artifacts/contracts",
			BindingsPath: "abi/bindings",
			PrivateKey:   os.Getenv("ETH_TESTNET_PRIVATE_KEY"),
			GetSigner:    func(chainID *big.Int) types.Signer { return types.NewLondonSigner(chainID) },
		},
		"local": {
			ChainID:      big.NewInt(31337),
			NetworkName:  "Local",
			NodeURL:      "http://localhost:8545",
			IsTestnet:    true,
			ABIPath:      "D:/work/gitspace/web3/web3/hardhat-project/artifacts/contracts",
			BindingsPath: "abi/bindings",
			PrivateKey:   "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
			GetSigner:    func(chainID *big.Int) types.Signer { return types.NewLondonSigner(chainID) },
		},
	}

	// 合约配置
	Contracts = ContractConfig{
		LockValue:     new(big.Int),
		PurchaseValue: new(big.Int),
		AuctionTime:   3600,
		LockTime:      3600,
	}

	// 账户配置
	Accounts = AccountConfig{
		DefaultBalance: new(big.Int),
		Addresses: map[string]common.Address{
			"default": common.HexToAddress("0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"),
			"second":  common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0"),
		},
		PrivateKeys: map[string]string{
			"default": "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
			"second":  "0x6cbed15c793ce57650b9877cf6fa156fbef513c4e6134f022a85b1ffdd59b2a1",
		},
	}

	// GlobalConfig 全局配置实例
	GlobalConfig = &Config{
		Networks:   Networks,
		Contracts:  Contracts,
		Accounts:   Accounts,
		CurrentEnv: "local", // 默认使用本地网络
	}
)

func init() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Printf("警告：无法加载.env文件: %v", err)
	}

	// 设置默认值
	GlobalConfig.Contracts.LockValue.SetString("1000000000000000000", 10)       // 1 ETH
	GlobalConfig.Contracts.PurchaseValue.SetString("2000000000000000000", 10)   // 2 ETH
	GlobalConfig.Accounts.DefaultBalance.SetString("100000000000000000000", 10) // 100 ETH

	// 首先从环境变量加载配置
	loadFromEnv()

	// 然后尝试从配置文件加载（会覆盖环境变量的值）
	loadFromFile()

	// 设置环境变量
	network := GetCurrentNetwork()
	os.Setenv("ABI_PATH", network.ABIPath)
}

// loadFromEnv 从环境变量加载配置
func loadFromEnv() {
	// 加载当前环境
	if env := os.Getenv("ETH_ENV"); env != "" {
		GlobalConfig.CurrentEnv = env
	}

	// 加载合约配置
	if lockValue := os.Getenv("CONTRACT_LOCK_VALUE"); lockValue != "" {
		if value, ok := new(big.Int).SetString(lockValue, 10); ok {
			GlobalConfig.Contracts.LockValue = value
		}
	}
	if purchaseValue := os.Getenv("CONTRACT_PURCHASE_VALUE"); purchaseValue != "" {
		if value, ok := new(big.Int).SetString(purchaseValue, 10); ok {
			GlobalConfig.Contracts.PurchaseValue = value
		}
	}
	if auctionTime := os.Getenv("CONTRACT_AUCTION_TIME"); auctionTime != "" {
		if value, err := strconv.ParseInt(auctionTime, 10, 64); err == nil {
			GlobalConfig.Contracts.AuctionTime = value
		}
	}
	if lockTime := os.Getenv("CONTRACT_LOCK_TIME"); lockTime != "" {
		if value, err := strconv.ParseInt(lockTime, 10, 64); err == nil {
			GlobalConfig.Contracts.LockTime = value
		}
	}

	// 加载账户配置
	if balance := os.Getenv("ACCOUNT_DEFAULT_BALANCE"); balance != "" {
		if value, ok := new(big.Int).SetString(balance, 10); ok {
			GlobalConfig.Accounts.DefaultBalance = value
		}
	}
}

// loadFromFile 从配置文件加载配置
func loadFromFile() {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Printf("无法读取配置文件: %v", err)
		return
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Printf("解析配置文件失败: %v", err)
		return
	}

	// 合并配置
	if config.CurrentEnv != "" {
		GlobalConfig.CurrentEnv = config.CurrentEnv
	}
	if config.Contracts.LockValue != nil {
		GlobalConfig.Contracts.LockValue = config.Contracts.LockValue
	}
	if config.Contracts.PurchaseValue != nil {
		GlobalConfig.Contracts.PurchaseValue = config.Contracts.PurchaseValue
	}
	if config.Contracts.AuctionTime != 0 {
		GlobalConfig.Contracts.AuctionTime = config.Contracts.AuctionTime
	}
	if config.Contracts.LockTime != 0 {
		GlobalConfig.Contracts.LockTime = config.Contracts.LockTime
	}
	if config.Accounts.DefaultBalance != nil {
		GlobalConfig.Accounts.DefaultBalance = config.Accounts.DefaultBalance
	}
}

// GetCurrentNetwork 获取当前网络配置
func GetCurrentNetwork() NetworkConfig {
	return GlobalConfig.Networks[GlobalConfig.CurrentEnv]
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	return GlobalConfig
}
