package utils

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"task2/abi/bindings"
	"task2/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// DeployContracts 部署所有合约
func DeployContracts() error {
	// 连接到本地节点
	client, err := InitClient()
	if err != nil {
		return fmt.Errorf("连接到以太坊节点失败: %v", err)
	}
	defer client.Close()

	// 获取网络配置
	network := config.GetCurrentNetwork()

	// 获取私钥
	privateKey, err := GetPrivateKey(network.PrivateKey)
	if err != nil {
		return fmt.Errorf("解析私钥失败: %v", err)
	}

	// 创建交易选项
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("创建交易选项失败: %v", err)
	}

	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg := config.GetConfig()

	// 部署SimpleStorage合约
	auth.Value = big.NewInt(0)
	address, tx, _, err := bindings.DeploySimpleStorage(auth, client)
	if err != nil {
		return fmt.Errorf("部署SimpleStorage合约失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待SimpleStorage部署交易确认失败: %v", err)
	}
	fmt.Printf("SimpleStorage合约已部署到: %s\n", address.Hex())
	cfg.Contracts.SimpleStorageAddress = address.Hex()

	// 部署Lock合约
	unlockTime := big.NewInt(time.Now().Unix() + int64(cfg.Contracts.LockTime))
	lockValue := cfg.Contracts.LockValue
	auth.Value = lockValue
	address, tx, _, err = bindings.DeployLock(auth, client, unlockTime)
	if err != nil {
		return fmt.Errorf("部署Lock合约失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待Lock部署交易确认失败: %v", err)
	}
	fmt.Printf("Lock合约已部署到: %s\n", address.Hex())
	cfg.Contracts.LockAddress = address.Hex()

	// 部署Shipping合约
	auth.Value = big.NewInt(0)
	address, tx, _, err = bindings.DeployShipping(auth, client)
	if err != nil {
		return fmt.Errorf("部署Shipping合约失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待Shipping部署交易确认失败: %v", err)
	}
	fmt.Printf("Shipping合约已部署到: %s\n", address.Hex())
	cfg.Contracts.ShippingAddress = address.Hex()

	// 部署SimpleAuction合约
	auth.Value = big.NewInt(0)
	biddingTime := big.NewInt(int64(cfg.Contracts.AuctionTime))
	beneficiary := auth.From
	address, tx, _, err = bindings.DeploySimpleAuction(auth, client, biddingTime, beneficiary)
	if err != nil {
		return fmt.Errorf("部署SimpleAuction合约失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待SimpleAuction部署交易确认失败: %v", err)
	}
	fmt.Printf("SimpleAuction合约已部署到: %s\n", address.Hex())
	cfg.Contracts.SimpleAuctionAddress = address.Hex()

	// 部署Purchase合约
	purchaseValue := cfg.Contracts.PurchaseValue
	auth.Value = purchaseValue
	address, tx, _, err = bindings.DeployPurchase(auth, client)
	if err != nil {
		return fmt.Errorf("部署Purchase合约失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待Purchase部署交易确认失败: %v", err)
	}
	fmt.Printf("Purchase合约已部署到: %s\n", address.Hex())
	cfg.Contracts.PurchaseAddress = address.Hex()

	// 保存合约地址到配置文件
	if err := config.SaveContractAddresses(cfg.Contracts); err != nil {
		return fmt.Errorf("保存合约地址失败: %v", err)
	}

	return nil
}
