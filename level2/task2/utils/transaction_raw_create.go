package utils

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"task2/config"
)

// CreateRawTransaction 创建原始交易
func CreateRawTransaction() error {
	client, err := GetEthClientHTTP()
	if err != nil {
		return err
	}
	// 使用的是单例客户端，不需要关闭

	// 获取网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return fmt.Errorf("未找到网络配置")
	}

	// 获取默认账户
	account, ok := network.Accounts["default"]
	if !ok {
		return fmt.Errorf("未找到默认账户")
	}

	// 获取私钥
	privateKey, err := crypto.HexToECDSA(account.PrivateKey[2:]) // 移除 "0x" 前缀
	if err != nil {
		return fmt.Errorf("私钥解析失败: %v", err)
	}

	// 获取发送方地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("获取 nonce 失败: %v", err)
	}

	// 获取 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("获取 gas 价格失败: %v", err)
	}

	// 创建交易
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress("0x1234567890123456789012345678901234567890"),
		big.NewInt(1000000000000000000), // 1 ETH
		uint64(21000),
		gasPrice,
		nil,
	)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(network.ChainID)), privateKey)
	if err != nil {
		return fmt.Errorf("签名交易失败: %v", err)
	}

	// 获取原始交易数据
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("序列化交易失败: %v", err)
	}

	// 输出原始交易数据
	log.Printf("原始交易数据: %x", rawTxBytes)

	return nil
}
