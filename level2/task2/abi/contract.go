package abi

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"task2/abi/bindings"
	"task2/config"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetContractByteCode 根据合约地址获取合约字节码
func GetContractByteCode(client *ethclient.Client, address common.Address) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bytecode, err := client.CodeAt(ctx, address, nil)
	if err != nil {
		return "", fmt.Errorf("获取合约字节码失败: %v", err)
	}

	if len(bytecode) == 0 {
		return "", fmt.Errorf("地址 %s 不是合约地址或合约不存在", address.Hex())
	}

	return common.Bytes2Hex(bytecode), nil
}

// DeploySimpleStorage 部署 SimpleStorage 合约
func DeploySimpleStorage(client *ethclient.Client) (common.Address, common.Hash, *bindings.SimpleStorage, error) {
	network := config.GetCurrentNetwork()
	privateKey, err := utils.GetPrivateKey(network.Accounts["default"].PrivateKey)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	address, tx, contract, err := bindings.DeploySimpleStorage(auth, client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("部署合约失败: %v", err)
	}
	return address, tx.Hash(), contract, nil
}

// DeployLock 部署 Lock 合约
func DeployLock(client *ethclient.Client) (common.Address, common.Hash, *bindings.Lock, error) {
	network := config.GetCurrentNetwork()
	privateKey, err := utils.GetPrivateKey(network.Accounts["default"].PrivateKey)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	unlockTime := big.NewInt(time.Now().Unix() + 3600) // 1小时后解锁
	address, tx, contract, err := bindings.DeployLock(auth, client, unlockTime)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("部署合约失败: %v", err)
	}

	return address, tx.Hash(), contract, nil
}

// DeployShipping 部署 Shipping 合约
func DeployShipping(client *ethclient.Client) (common.Address, common.Hash, *bindings.Shipping, error) {
	network := config.GetCurrentNetwork()
	privateKey, err := utils.GetPrivateKey(network.Accounts["default"].PrivateKey)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	address, tx, contract, err := bindings.DeployShipping(auth, client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("部署合约失败: %v", err)
	}

	return address, tx.Hash(), contract, nil
}

// DeploySimpleAuction 部署 SimpleAuction 合约
func DeploySimpleAuction(client *ethclient.Client) (common.Address, common.Hash, *bindings.SimpleAuction, error) {
	network := config.GetCurrentNetwork()
	privateKey, err := utils.GetPrivateKey(network.Accounts["default"].PrivateKey)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	biddingTime := big.NewInt(3600) // 1小时
	beneficiary := auth.From
	address, tx, contract, err := bindings.DeploySimpleAuction(auth, client, biddingTime, beneficiary)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("部署合约失败: %v", err)
	}

	return address, tx.Hash(), contract, nil
}

// DeployArrayDemo 部署 ArrayDemo 合约
func DeployArrayDemo(client *ethclient.Client) (common.Address, common.Hash, *bindings.ArrayDemo, error) {
	network := config.GetCurrentNetwork()
	privateKey, err := utils.GetPrivateKey(network.Accounts["default"].PrivateKey)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	address, tx, contract, err := bindings.DeployArrayDemo(auth, client)
	if err != nil {
		return common.Address{}, common.Hash{}, nil, fmt.Errorf("部署合约失败: %v", err)
	}

	return address, tx.Hash(), contract, nil
}

// GetContractABI 获取合约 ABI
func GetContractABI(name string) (*abi.ABI, error) {
	var abiStr string

	switch strings.ToLower(name) {
	case "simplestorage":
		abiStr = SimpleStorageABI
	case "lock":
		abiStr = LockABI
	case "shipping":
		abiStr = ShippingABI
	case "simpleauction":
		abiStr = SimpleAuctionABI
	case "arraydemo":
		abiStr = ArrayDemoABI
	default:
		return nil, fmt.Errorf("未知的合约类型: %s", name)
	}

	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, fmt.Errorf("解析 ABI 失败: %v", err)
	}

	return &parsed, nil
}

// SimpleStorageABI SimpleStorage 合约的 ABI
const SimpleStorageABI = `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "x",
				"type": "uint256"
			}
		],
		"name": "set",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "get",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

// LockABI Lock 合约的 ABI
const LockABI = `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_unlockTime",
				"type": "uint256"
			}
		],
		"stateMutability": "payable",
		"type": "constructor"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "unlockTime",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdraw",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

// ShippingABI Shipping 合约的 ABI
const ShippingABI = `[
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_itemName",
				"type": "string"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"inputs": [],
		"name": "itemName",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "shipped",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "ship",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

// SimpleAuctionABI SimpleAuction 合约的 ABI
const SimpleAuctionABI = `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_biddingTime",
				"type": "uint256"
			},
			{
				"internalType": "address payable",
				"name": "_beneficiary",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"inputs": [],
		"name": "bid",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "auctionEnd",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "highestBidder",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "highestBid",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

// ArrayDemoABI ArrayDemo 合约的 ABI
const ArrayDemoABI = `[
	{
		"inputs": [],
		"name": "getArray",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "",
				"type": "uint256[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_value",
				"type": "uint256"
			}
		],
		"name": "addValue",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "removeValue",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
