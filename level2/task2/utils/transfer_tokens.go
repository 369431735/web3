package utils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"task2/config"
)

// TransferTokens 转账代币
func TransferTokens(from, to string, amount *big.Int) (string, error) {
	client, err := InitClient()
	if err != nil {
		return "", err
	}
	defer client.Close()

	// 获取网络配置
	network := config.GetCurrentNetwork()
	if network == nil {
		return "", fmt.Errorf("未找到网络配置")
	}

	// 获取发送方账户
	account, ok := network.Accounts[from]
	if !ok {
		return "", fmt.Errorf("未找到账户: %s", from)
	}

	// 获取私钥
	privateKey, err := crypto.HexToECDSA(account.PrivateKey[2:]) // 移除 "0x" 前缀
	if err != nil {
		return "", fmt.Errorf("私钥解析失败: %v", err)
	}

	// 获取发送方地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("获取 nonce 失败: %v", err)
	}

	// 获取 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取 gas 价格失败: %v", err)
	}

	// 创建交易
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(to),
		amount,
		uint64(21000),
		gasPrice,
		nil,
	)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(network.ChainID)), privateKey)
	if err != nil {
		return "", fmt.Errorf("签名交易失败: %v", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("发送交易失败: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

// GetTokenBalance 获取代币余额
func GetTokenBalance(address string) (*big.Int, error) {
	client, err := InitClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// 获取账户余额
	accountAddress := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("获取余额失败: %v", err)
	}

	return balance, nil
}
