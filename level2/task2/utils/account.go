package utils

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Account 账户结构
type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

// CreateAccount 创建新账户
func CreateAccount() (*Account, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("生成私钥失败: %v", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return &Account{
		Address:    address,
		PrivateKey: privateKey,
	}, nil
}

// CreateKeystore 创建密钥库
func CreateKeystore() (*Account, error) {
	// 创建临时目录用于存储密钥文件
	tmpDir := filepath.Join(os.TempDir(), "ethereum-keystore")
	if err := os.MkdirAll(tmpDir, 0700); err != nil {
		return nil, fmt.Errorf("创建密钥库目录失败: %v", err)
	}

	// 创建密钥库实例
	ks := keystore.NewKeyStore(tmpDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// 创建新账户
	account, err := ks.NewAccount("password")
	if err != nil {
		return nil, fmt.Errorf("创建账户失败: %v", err)
	}

	// 从密钥库中获取私钥
	json, err := os.ReadFile(account.URL.Path)
	if err != nil {
		return nil, fmt.Errorf("读取密钥文件失败: %v", err)
	}

	key, err := keystore.DecryptKey(json, "password")
	if err != nil {
		return nil, fmt.Errorf("解密密钥失败: %v", err)
	}

	return &Account{
		Address:    account.Address,
		PrivateKey: key.PrivateKey,
	}, nil
}

// GetAccount 获取账户信息
func GetAccount(address common.Address) (*Account, error) {
	// 这里应该实现从本地存储或其他地方获取账户信息的逻辑
	// 目前仅返回错误
	return nil, fmt.Errorf("未找到账户信息")
}
