package config

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/yaml.v2"
)

// Config 配置结构体
type Config struct {
	Server struct {
		Port           int    `yaml:"port"`
		Mode           string `yaml:"mode"`
		ReadTimeout    int    `yaml:"read_timeout"`
		WriteTimeout   int    `yaml:"write_timeout"`
		MaxHeaderBytes int    `yaml:"max_header_bytes"`
		BasePath       string `yaml:"base_path"`
	} `yaml:"server"`
	Ethereum struct {
		Networks map[string]NetworkConfig `yaml:"networks"`
	} `yaml:"ethereum"`
	Log struct {
		Level    string `yaml:"level"`
		Filename string `yaml:"filename"`
	} `yaml:"log"`
}

// NetworkConfig 网络配置结构体
type NetworkConfig struct {
	NetworkName string                    `yaml:"network_name"`
	RPCURL      string                    `yaml:"rpc_url"`
	WSURL       string                    `yaml:"ws_url"`
	ChainID     int64                     `yaml:"chain_id"`
	Accounts    map[string]AccountConfig  `yaml:"accounts"`
	Contracts   map[string]ContractConfig `yaml:"contracts"`
}

// GetSigner 获取签名器
func (n *NetworkConfig) GetSigner() types.Signer {
	return types.NewLondonSigner(big.NewInt(n.ChainID))
}

// ContractConfig 合约配置结构体
type ContractConfig struct {
	Address string `yaml:"address"`
	ABI     string `yaml:"abi"`
}

// AccountConfig 账户配置结构体
type AccountConfig struct {
	PrivateKey string `yaml:"private_key"`
	Address    string `yaml:"address"`
}

var (
	cfg            *Config
	currentNetwork string
)

// LoadConfig 加载配置文件
func LoadConfig() error {
	configPath := filepath.Join("config", "config.yml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	cfg = &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 设置默认网络
	if currentNetwork == "" {
		currentNetwork = "local"
	}

	return nil
}

// GetConfig 获取配置
func GetConfig() *Config {
	return cfg
}

// SaveConfig 保存配置
func SaveConfig() error {
	if cfg == nil {
		return fmt.Errorf("配置未初始化")
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	configPath := filepath.Join("config", "config.yml")
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("保存配置文件失败: %v", err)
	}

	return nil
}

// GetCurrentNetwork 获取当前网络配置
func GetCurrentNetwork() *NetworkConfig {
	if cfg == nil {
		return nil
	}
	if network, ok := cfg.Ethereum.Networks[currentNetwork]; ok {
		return &network
	}
	return nil
}

// SetCurrentNetwork 设置当前网络
func SetCurrentNetwork(network string) error {
	if cfg == nil {
		return fmt.Errorf("配置未初始化")
	}
	if _, ok := cfg.Ethereum.Networks[network]; !ok {
		return fmt.Errorf("网络 %s 未配置", network)
	}
	currentNetwork = network
	return SaveConfig()
}
