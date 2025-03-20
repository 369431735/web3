package utils

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"task2/config"
	"task2/contracts"
)

var currentNetwork = "local" // 默认使用本地网络

// DeployContracts 部署所有合约
func DeployContracts() error {
	log.Println("=== 开始部署合约 ===")

	// 获取网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return fmt.Errorf("获取网络配置失败")
	}

	// 连接以太坊客户端
	client, err := ethclient.Dial(network.RPCURL)
	if err != nil {
		return fmt.Errorf("连接以太坊客户端失败: %v", err)
	}

	// 获取默认账户的私钥
	account, ok := network.Accounts["default"]
	if !ok {
		return fmt.Errorf("未找到默认账户")
	}

	privateKey, err := GetPrivateKey(account.PrivateKey)
	if err != nil {
		return fmt.Errorf("获取私钥失败: %v", err)
	}

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(network.ChainID))
	if err != nil {
		return fmt.Errorf("创建交易选项失败: %v", err)
	}

	// 部署 SimpleStorage 合约
	log.Println("正在部署 SimpleStorage 合约...")
	simpleStorageAddr, _, _, err := contracts.DeploySimpleStorage(auth, client)
	if err != nil {
		return fmt.Errorf("部署 SimpleStorage 合约失败: %v", err)
	}
	log.Printf("SimpleStorage 合约已部署到: %s", simpleStorageAddr.Hex())
	SaveContractAddress("simple_storage", simpleStorageAddr)

	// 部署 Lock 合约
	log.Println("正在部署 Lock 合约...")
	unlockTime := time.Now().Add(24 * time.Hour).Unix()
	lockAddr, _, _, err := contracts.DeployLock(auth, client, big.NewInt(unlockTime))
	if err != nil {
		return fmt.Errorf("部署 Lock 合约失败: %v", err)
	}
	log.Printf("Lock 合约已部署到: %s", lockAddr.Hex())
	SaveContractAddress("lock", lockAddr)

	log.Println("=== 所有合约部署完成 ===")
	return nil
}

// SaveContractAddress 保存合约地址
func SaveContractAddress(name string, address common.Address) error {
	cfg := config.GetConfig()
	if cfg == nil {
		return fmt.Errorf("配置未初始化")
	}

	network := config.GetCurrentNetwork()
	if network == nil {
		return fmt.Errorf("未找到网络配置")
	}

	if network.Contracts == nil {
		network.Contracts = make(map[string]config.ContractConfig)
	}

	contract := network.Contracts[name]
	contract.Address = address.Hex()
	network.Contracts[name] = contract
	cfg.Ethereum.Networks[currentNetwork] = *network

	return config.SaveConfig()
}

// GetContractAddress 获取合约地址
func GetContractAddress(name string) (common.Address, error) {
	network := config.GetCurrentNetwork()
	if network == nil {
		return common.Address{}, fmt.Errorf("未找到网络配置")
	}

	contract, ok := network.Contracts[name]
	if !ok {
		return common.Address{}, fmt.Errorf("找不到合约 %s 的地址", name)
	}

	return common.HexToAddress(contract.Address), nil
}

// GetContractABI 获取合约 ABI
func GetContractABI(contractName string) (string, error) {
	network := config.GetCurrentNetwork()
	if network == nil {
		return "", fmt.Errorf("未找到网络配置")
	}

	contract, ok := network.Contracts[contractName]
	if !ok {
		return "", fmt.Errorf("未找到合约: %s", contractName)
	}

	return contract.ABI, nil
}
