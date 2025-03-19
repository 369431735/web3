package utils

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 生成新的密钥对
func GenerateKeyPair() error {
	log.Println("=== 生成新的密钥对 ===")

	// 生成新的私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("生成私钥失败: %v", err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("转换公钥类型失败")
	}

	// 从公钥获取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 输出结果
	log.Printf("私钥: 0x%x", crypto.FromECDSA(privateKey))
	log.Printf("公钥: 0x%x", crypto.FromECDSAPub(publicKeyECDSA))
	log.Printf("地址: %s", address.Hex())

	return nil
}

// 签名消息
func SignMessage() error {
	log.Println("=== 消息签名演示 ===")

	// 要签名的消息
	message := []byte("Hello, Ethereum!")
	log.Printf("原始消息: %s", string(message))

	// 计算消息的 Keccak256 哈希
	hash := crypto.Keccak256Hash(message)
	log.Printf("消息哈希: %s", hash.Hex())

	// 生成新的私钥用于签名
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("生成私钥失败: %v", err)
	}

	// 签名消息
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return fmt.Errorf("签名消息失败: %v", err)
	}
	log.Printf("签名: 0x%x", signature)

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("转换公钥类型失败")
	}

	// 从公钥获取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("签名者地址: %s", address.Hex())

	// 验证签名
	signatureNoRecoverID := signature[:len(signature)-1] // 移除恢复ID
	verified := crypto.VerifySignature(
		crypto.FromECDSAPub(publicKeyECDSA),
		hash.Bytes(),
		signatureNoRecoverID,
	)
	log.Printf("签名验证结果: %v", verified)

	// 从签名恢复公钥
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return fmt.Errorf("从签名恢复公钥失败: %v", err)
	}
	log.Printf("从签名恢复的公钥: 0x%x", sigPublicKey)

	// 从签名恢复地址
	recoveredPubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return fmt.Errorf("从签名恢复公钥失败: %v", err)
	}
	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey)
	log.Printf("从签名恢复的地址: %s", recoveredAddress.Hex())

	// 验证恢复的地址是否正确
	matches := address == recoveredAddress
	log.Printf("地址匹配: %v", matches)

	return nil
}

// 验证签名
func VerifySignature() error {
	log.Println("=== 签名验证演示 ===")

	// 生成新的私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("生成私钥失败: %v", err)
	}

	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("转换公钥类型失败")
	}

	// 获取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("地址: %s", address.Hex())

	// 要签名的消息
	message := []byte("Hello, Ethereum!")
	log.Printf("消息: %s", string(message))

	// 计算消息的哈希
	hash := crypto.Keccak256Hash(message)
	log.Printf("消息哈希: %s", hash.Hex())

	// 签名消息
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return fmt.Errorf("签名消息失败: %v", err)
	}
	log.Printf("签名: 0x%x", signature)

	// 验证签名
	signatureNoRecoverID := signature[:len(signature)-1] // 移除恢复ID
	verified := crypto.VerifySignature(
		crypto.FromECDSAPub(publicKeyECDSA),
		hash.Bytes(),
		signatureNoRecoverID,
	)
	log.Printf("签名验证结果: %v", verified)

	// 使用错误的消息验证
	wrongMessage := []byte("Wrong message")
	wrongHash := crypto.Keccak256Hash(wrongMessage)
	wrongVerified := crypto.VerifySignature(
		crypto.FromECDSAPub(publicKeyECDSA),
		wrongHash.Bytes(),
		signatureNoRecoverID,
	)
	log.Printf("错误消息的验证结果: %v", wrongVerified)

	return nil
}

// 检查签名
func CheckSignature() error {
	log.Println("=== 签名检查演示 ===")

	// 生成新的私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("生成私钥失败: %v", err)
	}

	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("转换公钥类型失败")
	}

	// 获取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("地址: %s", address.Hex())

	// 要签名的消息
	message := []byte("Hello, Ethereum!")
	log.Printf("消息: %s", string(message))

	// 计算消息的哈希
	hash := crypto.Keccak256Hash(message)
	log.Printf("消息哈希: %s", hash.Hex())

	// 签名消息
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return fmt.Errorf("签名消息失败: %v", err)
	}
	log.Printf("签名: 0x%x", signature)

	// 从签名恢复地址
	recoveredPubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return fmt.Errorf("从签名恢复公钥失败: %v", err)
	}
	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey)
	log.Printf("从签名恢复的地址: %s", recoveredAddress.Hex())

	// 验证地址
	matches := address == recoveredAddress
	log.Printf("地址匹配: %v", matches)

	// 验证签名
	signatureNoRecoverID := signature[:len(signature)-1] // 移除恢复ID
	verified := crypto.VerifySignature(
		crypto.FromECDSAPub(publicKeyECDSA),
		hash.Bytes(),
		signatureNoRecoverID,
	)
	log.Printf("签名验证结果: %v", verified)

	return nil
}

// 生成以太坊地址
func GenerateAddress() error {
	log.Println("=== 生成以太坊地址演示 ===")

	// 生成新的私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("生成私钥失败: %v", err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("转换公钥类型失败")
	}

	// 从公钥生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 输出结果
	log.Printf("私钥: 0x%x", crypto.FromECDSA(privateKey))
	log.Printf("公钥: 0x%x", crypto.FromECDSAPub(publicKeyECDSA))
	log.Printf("地址: %s", address.Hex())

	// 验证地址格式
	if !common.IsHexAddress(address.Hex()) {
		return fmt.Errorf("生成的地址格式无效")
	}

	return nil
}
