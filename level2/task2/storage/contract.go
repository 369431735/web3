package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// ContractAddresses 存储合约地址的结构
type ContractAddresses struct {
	SimpleStorage string `json:"simpleStorage"`
	Lock          string `json:"lock"`
	Shipping      string `json:"shipping"`
	SimpleAuction string `json:"simpleAuction"`
	ArrayDemo     string `json:"arrayDemo"`
	Ballot        string `json:"ballot"`
	Lottery       string `json:"lottery"`
	Purchase      string `json:"purchase"`
}

var (
	instance *ContractStorage
	once     sync.Once
)

// ContractStorage 合约存储结构
type ContractStorage struct {
	filePath  string
	addresses *ContractAddresses
	mu        sync.RWMutex
}

// GetInstance 获取合约存储的单例实例
func GetInstance() *ContractStorage {
	once.Do(func() {
		instance = &ContractStorage{
			filePath:  "data/contracts.json",
			addresses: &ContractAddresses{},
		}
		// 确保数据目录存在
		if err := os.MkdirAll(filepath.Dir(instance.filePath), 0755); err != nil {
			panic(fmt.Sprintf("创建数据目录失败: %v", err))
		}
		// 加载已存在的合约地址
		if err := instance.load(); err != nil {
			fmt.Printf("加载合约地址失败: %v\n", err)
		}
	})
	return instance
}

// load 从文件加载合约地址
func (s *ContractStorage) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("读取合约地址文件失败: %v", err)
	}

	if err := json.Unmarshal(data, s.addresses); err != nil {
		return fmt.Errorf("解析合约地址数据失败: %v", err)
	}

	return nil
}

// save 保存合约地址到文件
func (s *ContractStorage) save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保目录存在
	dirPath := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Printf("创建目录失败 %s: %v\n", dirPath, err)
		return fmt.Errorf("创建目录失败: %v", err)
	}
	fmt.Printf("确保目录存在: %s\n", dirPath)

	data, err := json.MarshalIndent(s.addresses, "", "  ")
	if err != nil {
		fmt.Printf("序列化合约地址数据失败: %v\n", err)
		return fmt.Errorf("序列化合约地址数据失败: %v", err)
	}
	fmt.Printf("序列化数据成功，准备写入文件: %s\n", s.filePath)

	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		fmt.Printf("保存合约地址文件失败 %s: %v\n", s.filePath, err)
		return fmt.Errorf("保存合约地址文件失败: %v", err)
	}
	fmt.Printf("成功写入合约地址到文件: %s\n", s.filePath)

	return nil
}

// SetAddress 设置合约地址
func (s *ContractStorage) SetAddress(contractType string, address string) error {

	switch contractType {
	case "SimpleStorage":
		s.addresses.SimpleStorage = address
	case "Lock":
		s.addresses.Lock = address
	case "Shipping":
		s.addresses.Shipping = address
	case "SimpleAuction":
		s.addresses.SimpleAuction = address
	case "ArrayDemo":
		s.addresses.ArrayDemo = address
	case "Ballot":
		s.addresses.Ballot = address
	case "Lottery":
		s.addresses.Lottery = address
	case "Purchase":
		s.addresses.Purchase = address
	default:
		return fmt.Errorf("未知的合约类型: %s", contractType)
	}

	return s.save()
}

// GetAddress 获取合约地址
func (s *ContractStorage) GetAddress(contractType string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	switch contractType {
	case "SimpleStorage":
		return s.addresses.SimpleStorage, nil
	case "Lock":
		return s.addresses.Lock, nil
	case "Shipping":
		return s.addresses.Shipping, nil
	case "SimpleAuction":
		return s.addresses.SimpleAuction, nil
	case "ArrayDemo":
		return s.addresses.ArrayDemo, nil
	case "Ballot":
		return s.addresses.Ballot, nil
	case "Lottery":
		return s.addresses.Lottery, nil
	case "Purchase":
		return s.addresses.Purchase, nil
	default:
		return "", fmt.Errorf("未知的合约类型: %s", contractType)
	}
}

// GetAllAddresses 获取所有合约地址
func (s *ContractStorage) GetAllAddresses() *ContractAddresses {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 创建副本以避免并发访问问题
	addresses := *s.addresses
	return &addresses
}

// GetAllContractAddresses returns a map of contract names to their addresses
func GetAllContractAddresses() map[string]string {
	// TODO: Implement actual storage logic
	return make(map[string]string)
}
