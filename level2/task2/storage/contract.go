package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	instance *ContractStorage
	once     sync.Once
)

// ContractStorage 合约存储结构
type ContractStorage struct {
	filePath  string
	addresses map[string]string
	mu        sync.RWMutex
}

// GetInstance 获取合约存储的单例实例
func GetInstance() *ContractStorage {
	once.Do(func() {
		instance = &ContractStorage{
			filePath:  "data/contracts.json",
			addresses: make(map[string]string),
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

	if err := json.Unmarshal(data, &s.addresses); err != nil {
		return fmt.Errorf("解析合约地址数据失败: %v", err)
	}

	return nil
}

// save 保存合约地址到文件
func (s *ContractStorage) save() error {
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
	s.mu.Lock()
	defer s.mu.Unlock()
	s.addresses[contractType] = address
	return s.save()
}

// GetAddress 获取合约地址
func (s *ContractStorage) GetAddress(contractType string) (string, error) {
	if value, ok := s.addresses[contractType]; ok {
		return value, nil
	}
	return "", fmt.Errorf("未知的合约类型: %s", contractType)

}

// GetAllAddresses 获取所有合约地址
func (s *ContractStorage) GetAllAddresses() map[string]string {
	return s.addresses
}
