package config

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// NetworkConfig 网络配置
type NetworkConfig struct {
	NetworkName string `yaml:"network_name"`
	ChainID     int64  `yaml:"chain_id"`
	NodeURL     string `yaml:"node_url"`
	PrivateKey  string `yaml:"private_key"`
}

// ContractConfig 合约配置
type ContractConfig struct {
	SimpleStorageAddress string   `yaml:"simpleStorageAddress"`
	LockAddress          string   `yaml:"lockAddress"`
	ShippingAddress      string   `yaml:"shippingAddress"`
	SimpleAuctionAddress string   `yaml:"simpleAuctionAddress"`
	PurchaseAddress      string   `yaml:"purchaseAddress"`
	LockTime             int64    `yaml:"lockTime"`
	LockValue            *big.Int `yaml:"lockValue"`
	AuctionTime          int64    `yaml:"auctionTime"`
	PurchaseValue        *big.Int `yaml:"purchaseValue"`
}

// AccountConfig 账户配置
type AccountConfig struct {
	DefaultBalance *big.Int `yaml:"default_balance"`
}

// Config 全局配置
type Config struct {
	Networks  map[string]NetworkConfig `yaml:"networks"`
	Contracts ContractConfig           `yaml:"contracts"`
	Accounts  AccountConfig            `yaml:"accounts"`
}

var (
	globalConfig   Config
	currentNetwork string
)

// GetConfig 获取全局配置
func GetConfig() *Config {
	return &globalConfig
}

// GetCurrentNetwork 获取当前网络配置
func GetCurrentNetwork() *NetworkConfig {
	if network, ok := globalConfig.Networks[currentNetwork]; ok {
		return &network
	}
	return nil
}

// GetSigner 获取签名器
func (n *NetworkConfig) GetSigner(chainID *big.Int) types.Signer {
	return types.NewLondonSigner(chainID)
}

// 初始化配置
func init() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Printf("警告: .env 文件未找到，使用默认配置")
	}

	// 创建大整数
	lockValue := new(big.Int)
	lockValue.SetString("1000000000000000000", 10) // 1 ETH

	purchaseValue := new(big.Int)
	purchaseValue.SetString("2000000000000000000", 10) // 2 ETH

	defaultBalance := new(big.Int)
	defaultBalance.SetString("10000000000000000000", 10) // 10 ETH

	// 设置默认配置
	globalConfig = Config{
		Networks: map[string]NetworkConfig{
			"mainnet": {
				NetworkName: "Mainnet",
				ChainID:     1,
				NodeURL:     "https://mainnet.infura.io/v3/your-project-id",
			},
			"sepolia": {
				NetworkName: "Sepolia",
				ChainID:     11155111,
				NodeURL:     "https://sepolia.infura.io/v3/your-project-id",
			},
			"local": {
				NetworkName: "Local",
				ChainID:     31337,
				NodeURL:     "http://127.0.0.1:8545",
				PrivateKey:  "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
			},
		},
		Contracts: ContractConfig{
			LockTime:      3600,          // 1小时
			LockValue:     lockValue,     // 1 ETH
			AuctionTime:   7200,          // 2小时
			PurchaseValue: purchaseValue, // 2 ETH
		},
		Accounts: AccountConfig{
			DefaultBalance: defaultBalance, // 10 ETH
		},
	}

	// 从配置文件加载
	configFile := "config.yml"
	if _, err := os.Stat(configFile); err == nil {
		data, err := os.ReadFile(configFile)
		if err != nil {
			log.Printf("读取配置文件失败: %v", err)
			return
		}

		var fileConfig Config
		if err := yaml.Unmarshal(data, &fileConfig); err != nil {
			log.Printf("解析配置文件失败: %v", err)
			return
		}

		// 合并配置
		mergeConfig(&globalConfig, &fileConfig)
	}

	// 设置当前网络
	network := os.Getenv("NETWORK")
	if network == "" {
		network = "local" // 默认使用本地网络
	}
	if _, ok := globalConfig.Networks[network]; !ok {
		log.Printf("警告: 网络 %s 未配置，使用本地网络", network)
		network = "local"
	}
	currentNetwork = network
}

// 合并配置
func mergeConfig(dst, src *Config) {
	// 合并网络配置
	for name, network := range src.Networks {
		if _, ok := dst.Networks[name]; !ok {
			dst.Networks[name] = network
		}
	}

	// 合并合约配置
	if src.Contracts.LockTime > 0 {
		dst.Contracts.LockTime = src.Contracts.LockTime
	}
	if src.Contracts.LockValue != nil {
		dst.Contracts.LockValue = src.Contracts.LockValue
	}
	if src.Contracts.AuctionTime > 0 {
		dst.Contracts.AuctionTime = src.Contracts.AuctionTime
	}
	if src.Contracts.PurchaseValue != nil {
		dst.Contracts.PurchaseValue = src.Contracts.PurchaseValue
	}

	// 合并账户配置
	if src.Accounts.DefaultBalance != nil {
		dst.Accounts.DefaultBalance = src.Accounts.DefaultBalance
	}
}

// SaveContractAddresses 保存合约地址到配置文件
func SaveContractAddresses(contracts ContractConfig) error {
	// 读取现有的配置文件
	data, err := os.ReadFile("config.yml")
	if err != nil {
		if os.IsNotExist(err) {
			// 如果文件不存在，创建一个新的配置
			data = []byte{}
		} else {
			return fmt.Errorf("读取配置文件失败: %v", err)
		}
	}

	// 解析现有的配置
	var config Config
	if len(data) > 0 {
		if err := yaml.Unmarshal(data, &config); err != nil {
			return fmt.Errorf("解析配置文件失败: %v", err)
		}
	}

	// 更新合约地址
	config.Contracts.SimpleStorageAddress = contracts.SimpleStorageAddress
	config.Contracts.LockAddress = contracts.LockAddress
	config.Contracts.ShippingAddress = contracts.ShippingAddress
	config.Contracts.SimpleAuctionAddress = contracts.SimpleAuctionAddress
	config.Contracts.PurchaseAddress = contracts.PurchaseAddress

	// 将配置写回文件
	newData, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	if err := os.WriteFile("config.yml", newData, 0644); err != nil {
		return fmt.Errorf("保存配置文件失败: %v", err)
	}

	return nil
}
