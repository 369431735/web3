package utils

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetPrivateKey 从私钥字符串获取私钥对象
func GetPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	// 移除 0x 前缀（如果存在）
	privateKeyStr = strings.TrimPrefix(privateKeyStr, "0x")
	privateKeyBytes, err := hexutil.Decode("0x" + privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("创建私钥对象失败: %v", err)
	}

	return privateKey, nil
}
