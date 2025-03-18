package main

/*
config.go - 网络配置相关
client.go - 客户端初始化和基础操作
wallet.go - 钱包相关操作
transaction.go - 交易相关操作
block.go - 区块相关操作
address.go - 地址相关操作
block_subscribe.go - 区块订阅相关操作
*/
import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"task2/abi/bindings"
	"task2/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 初始化以太坊客户端和账户
func setupClient() (*ethclient.Client, *bind.TransactOpts, error) {
	// 连接到本地节点
	client, err := utils.InitClient()
	if err != nil {
		return nil, nil, fmt.Errorf("连接到以太坊节点失败: %v", err)
	}

	// 使用测试账户私钥（Hardhat默认的第一个账户）
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		return nil, nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	// 创建交易选项
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, nil, fmt.Errorf("创建交易选项失败: %v", err)
	}

	return client, auth, nil
}

// 部署所有合约
func deployContracts() error {
	client, auth, err := setupClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	config := utils.GetConfig()

	// 部署SimpleStorage合约
	address, tx, _, err := bindings.DeploySimpleStorage(auth, client)
	if err != nil {
		return fmt.Errorf("部署SimpleStorage合约失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待SimpleStorage部署交易确认失败: %v", err)
	}
	fmt.Printf("SimpleStorage合约已部署到: %s\n", address.Hex())

	// 部署Lock合约
	unlockTime := big.NewInt(time.Now().Unix() + int64(config.Contracts.LockTime))
	lockValue := config.Contracts.LockValue
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

	// 部署ERC20MinerReward合约
	auth.Value = big.NewInt(0)
	address, tx, _, err = bindings.DeployERC20MinerReward(auth, client)
	if err != nil {
		return fmt.Errorf("部署ERC20MinerReward合约失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待ERC20MinerReward部署交易确认失败: %v", err)
	}
	fmt.Printf("ERC20MinerReward合约已部署到: %s\n", address.Hex())

	// 部署Shipping合约
	address, tx, _, err = bindings.DeployShipping(auth, client)
	if err != nil {
		return fmt.Errorf("部署Shipping合约失败: %v", err)
	}
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待Shipping部署交易确认失败: %v", err)
	}
	fmt.Printf("Shipping合约已部署到: %s\n", address.Hex())

	// 部署SimpleAuction合约
	biddingTime := big.NewInt(int64(config.Contracts.AuctionTime))
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

	// 部署Purchase合约
	purchaseValue := config.Contracts.PurchaseValue
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

	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("=== 程序开始执行 ===")
	defer func() {
		if r := recover(); r != nil {
			log.Printf("程序发生错误: %v", r)
		}
	}()

	// 创建区块订阅实例
	log.Println("\n1. 区块订阅演示")
	blockSub, err := utils.NewBlockSubscription()
	if err != nil {
		log.Printf("创建区块订阅失败: %v", err)
		return
	}

	// 订阅新区块
	err = blockSub.Subscribe(1, func(block *types.Block) {
		log.Printf("\n收到新区块回调:")
		log.Printf("区块号: %d", block.Number().Uint64())
		log.Printf("区块时间: %v", time.Unix(int64(block.Time()), 0))
		log.Printf("交易数量: %d", len(block.Transactions()))

		// 遍历区块中的所有交易
		for i, tx := range block.Transactions() {
			log.Printf("\n交易 %d:", i+1)
			log.Printf("  哈希: %s", tx.Hash().Hex())
			if tx.To() != nil {
				log.Printf("  接收者: %s", tx.To().Hex())
			}
			log.Printf("  金额: %s Wei", tx.Value().String())
		}
	})
	if err != nil {
		log.Printf("订阅新区块失败: %v", err)
		return
	}
	log.Println("成功订阅新区块")

	// 等待一段时间确保订阅已经开始
	time.Sleep(2 * time.Second)

	log.Println("开始执行以太坊操作演示...")

	// 设置账户余额
	if err := utils.SetAccountBalance(); err != nil {
		log.Printf("设置账户余额失败: %v", err)
		return
	}
	log.Println("账户余额设置成功")

	log.Println("\n2. 地址转换演示")
	utils.Address()

	log.Println("\n3. 余额查询演示")
	if err := utils.Balance(); err != nil {
		log.Printf("余额查询失败: %v", err)
		return
	}

	log.Println("\n4. 创建新钱包")
	if err := utils.NewWallet(); err != nil {
		log.Printf("创建新钱包失败: %v", err)
		return
	}

	log.Println("\n5. 创建 Keystore")
	if err := utils.CreateKs(); err != nil {
		log.Printf("创建 Keystore 失败: %v", err)
		return
	}

	log.Println("\n6. HD 钱包演示")
	if err := utils.Chdwallet(); err != nil {
		log.Printf("HD 钱包操作失败: %v", err)
		return
	}

	log.Println("\n7. 地址检查演示")
	if err := utils.AddressCheck(); err != nil {
		log.Printf("地址检查失败: %v", err)
		return
	}

	log.Println("\n8. 区块信息查询")
	if err := utils.BlockInfo(); err != nil {
		log.Printf("区块信息查询失败: %v", err)
		return
	}

	log.Println("\n9. 创建并发送交易")
	txHash, err := utils.CreateAndSendTransaction()
	if err != nil {
		log.Printf("创建并发送交易失败: %v", err)
		return
	}
	log.Printf("交易发送成功，交易哈希: %s", txHash)

	log.Println("\n10. 创建原始交易")
	if err := utils.CreateRawTransaction(); err != nil {
		log.Printf("创建原始交易失败: %v", err)
		return
	}

	log.Println("\n11. 部署智能合约")
	if err := deployContracts(); err != nil {
		log.Printf("部署智能合约失败: %v", err)
		return
	}
	log.Println("智能合约部署成功")

	// 等待一段时间以观察新区块
	time.Sleep(5 * time.Second)

	// 取消订阅
	blockSub.Unsubscribe()
	log.Println("区块订阅已取消")
	log.Println("程序结束")
}
