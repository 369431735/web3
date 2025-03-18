package main

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 初始化以太坊客户端和账户
func setupTest(t *testing.T) (*ethclient.Client, *bind.TransactOpts) {
	// 连接到本地节点
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("连接到以太坊节点失败: %v", err)
	}

	// 使用测试账户私钥（Hardhat默认的第一个账户）
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		t.Fatalf("解析私钥失败: %v", err)
	}

	// 创建交易选项
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatalf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		t.Fatalf("创建交易选项失败: %v", err)
	}

	return client, auth
}

// 测试SimpleStorage合约
func TestSimpleStorage(t *testing.T) {
	client, auth := setupTest(t)
	defer client.Close()

	// 部署合约
	address, tx, instance, err := DeploySimpleStorage(auth, client)
	if err != nil {
		t.Fatalf("部署合约失败: %v", err)
	}
	fmt.Printf("SimpleStorage合约已部署到: %s\n", address.Hex())

	// 等待交易被确认
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待交易确认失败: %v", err)
	}

	// 测试Store方法
	value := big.NewInt(42)
	tx, err = instance.Store(auth, value)
	if err != nil {
		t.Fatalf("存储值失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待Store交易确认失败: %v", err)
	}

	// 测试Retrieve方法
	result, err := instance.Retrieve(nil)
	if err != nil {
		t.Fatalf("获取值失败: %v", err)
	}
	if result.Cmp(value) != 0 {
		t.Errorf("期望值为 %s，实际值为 %s", value, result)
	}
}

// 测试Lock合约
func TestLock(t *testing.T) {
	client, auth := setupTest(t)
	defer client.Close()

	// 设置解锁时间为1小时后
	unlockTime := big.NewInt(time.Now().Unix() + 3600)
	auth.Value = big.NewInt(1000000000000000000) // 发送1 ETH

	// 部署合约
	address, tx, instance, err := DeployLock(auth, client, unlockTime)
	if err != nil {
		t.Fatalf("部署合约失败: %v", err)
	}
	fmt.Printf("Lock合约已部署到: %s\n", address.Hex())

	// 等待交易被确认
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待交易确认失败: %v", err)
	}

	// 测试Withdraw方法（应该失败，因为还未到解锁时间）
	tx, err = instance.Withdraw(auth)
	if err == nil {
		t.Error("提取应该失败，因为还未到解锁时间")
	}
}

// 测试ERC20MinerReward合约
func TestERC20MinerReward(t *testing.T) {
	client, auth := setupTest(t)
	defer client.Close()

	// 部署合约
	address, tx, instance, err := DeployERC20MinerReward(auth, client)
	if err != nil {
		t.Fatalf("部署合约失败: %v", err)
	}
	fmt.Printf("ERC20MinerReward合约已部署到: %s\n", address.Hex())

	// 等待交易被确认
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待交易确认失败: %v", err)
	}

	// 测试_reward方法
	tx, err = instance.Reward(auth)
	if err != nil {
		t.Fatalf("调用reward方法失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待reward交易确认失败: %v", err)
	}

	// 测试balanceOf方法
	balance, err := instance.BalanceOf(nil, auth.From)
	if err != nil {
		t.Fatalf("获取余额失败: %v", err)
	}
	fmt.Printf("账户余额: %s\n", balance.String())
}

// 测试Shipping合约
func TestShipping(t *testing.T) {
	client, auth := setupTest(t)
	defer client.Close()

	// 部署合约
	address, tx, instance, err := DeployShipping(auth, client)
	if err != nil {
		t.Fatalf("部署合约失败: %v", err)
	}
	fmt.Printf("Shipping合约已部署到: %s\n", address.Hex())

	// 等待交易被确认
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待交易确认失败: %v", err)
	}

	// 测试Status方法
	status, err := instance.Status(nil)
	if err != nil {
		t.Fatalf("获取状态失败: %v", err)
	}
	fmt.Printf("初始状态: %v\n", status)

	// 测试Shipped方法
	tx, err = instance.Shipped(auth)
	if err != nil {
		t.Fatalf("调用Shipped方法失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待Shipped交易确认失败: %v", err)
	}

	// 测试Delivered方法
	tx, err = instance.Delivered(auth)
	if err != nil {
		t.Fatalf("调用Delivered方法失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待Delivered交易确认失败: %v", err)
	}
}

// 测试SimpleAuction合约
func TestSimpleAuction(t *testing.T) {
	client, auth := setupTest(t)
	defer client.Close()

	// 设置拍卖时间为1小时
	biddingTime := big.NewInt(3600)
	beneficiary := auth.From

	// 部署合约
	address, tx, instance, err := DeploySimpleAuction(auth, client, biddingTime, beneficiary)
	if err != nil {
		t.Fatalf("部署合约失败: %v", err)
	}
	fmt.Printf("SimpleAuction合约已部署到: %s\n", address.Hex())

	// 等待交易被确认
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待交易确认失败: %v", err)
	}

	// 测试bid方法
	auth.Value = big.NewInt(1000000000000000000) // 1 ETH
	tx, err = instance.Bid(auth)
	if err != nil {
		t.Fatalf("投标失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待bid交易确认失败: %v", err)
	}

	// 测试highestBid方法
	highestBid, err := instance.HighestBid(nil)
	if err != nil {
		t.Fatalf("获取最高出价失败: %v", err)
	}
	fmt.Printf("最高出价: %s\n", highestBid.String())
}

// 测试Purchase合约
func TestPurchase(t *testing.T) {
	client, auth := setupTest(t)
	defer client.Close()

	// 部署合约
	auth.Value = big.NewInt(2000000000000000000) // 2 ETH
	address, tx, instance, err := DeployPurchase(auth, client)
	if err != nil {
		t.Fatalf("部署合约失败: %v", err)
	}
	fmt.Printf("Purchase合约已部署到: %s\n", address.Hex())

	// 等待交易被确认
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待交易确认失败: %v", err)
	}

	// 测试confirmPurchase方法
	auth.Value = big.NewInt(2000000000000000000) // 2 ETH
	tx, err = instance.ConfirmPurchase(auth)
	if err != nil {
		t.Fatalf("确认购买失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待confirmPurchase交易确认失败: %v", err)
	}

	// 测试confirmReceived方法
	tx, err = instance.ConfirmReceived(auth)
	if err != nil {
		t.Fatalf("确认收货失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		t.Fatalf("等待confirmReceived交易确认失败: %v", err)
	}
}

// 测试生成合约绑定
func TestGenerateBindings(t *testing.T) {
	if err := GenerateBindings(); err != nil {
		t.Fatalf("生成合约绑定失败: %v", err)
	}
}
