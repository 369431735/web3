package utils

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"golang.org/x/crypto/sha3"
)

// 创建新钱包
func NewWallet() error {
	log.Println("=== 创建新钱包演示 ===")

	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("生成私钥失败: %v", err)
	}

	// 获取私钥的十六进制表示
	privateKeyBytes := crypto.FromECDSA(privateKey)
	log.Printf("私钥: %s", hexutil.Encode(privateKeyBytes)[2:])

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("无法将公钥转换为 ECDSA 格式")
	}

	// 获取公钥的十六进制表示
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	log.Printf("公钥: %s", hexutil.Encode(publicKeyBytes)[4:])

	// 从公钥生成以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("地址: %s", address.Hex())

	// 使用 Keccak-256 哈希函数计算地址
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	log.Printf("通过哈希计算的地址: %s", hexutil.Encode(hash.Sum(nil)[12:]))

	return nil
}

// 创建 Keystore
func CreateKs() error {
	log.Println("=== 创建 Keystore 演示 ===")

	// 创建 keystore 目录
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	log.Println("成功创建 Keystore 目录")

	// 创建新账户
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Printf("创建账户失败: %v", err)
		return nil
	}
	log.Printf("成功创建新账户，地址: %s", account.Address.Hex())

	return nil
}

// HD 钱包演示
func Chdwallet() error {
	log.Println("=== HD 钱包演示 ===")

	// 使用助记词创建 HD 钱包
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return fmt.Errorf("创建 HD 钱包失败: %v", err)
	}
	log.Println("成功创建 HD 钱包")

	// 派生第一个地址
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return fmt.Errorf("派生第一个地址失败: %v", err)
	}
	log.Printf("第一个派生地址: %s", account.Address.Hex())

	// 派生第二个地址
	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		return fmt.Errorf("派生第二个地址失败: %v", err)
	}
	log.Printf("第二个派生地址: %s", account.Address.Hex())

	return nil
}
