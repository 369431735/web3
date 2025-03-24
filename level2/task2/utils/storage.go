package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// ContractStorage 合约存储结构体
type ContractStorage struct {
	mu        sync.RWMutex
	filePath  string
	Contracts map[string]string `json:"contracts"`
}

// NewContractStorage 创建新的合约存储实例
func NewContractStorage() (*ContractStorage, error) {
	// 确保存储目录存在
	storageDir := "storage"
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return nil, fmt.Errorf("创建存储目录失败: %v", err)
	}

	filePath := filepath.Join(storageDir, "contracts.json")
	storage := &ContractStorage{
		filePath:  filePath,
		Contracts: make(map[string]string),
	}

	// 尝试从文件加载已存在的合约地址
	if err := storage.Load(); err != nil {
		// 如果文件不存在，不视为错误，创建一个新的存储
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return storage, nil
}

// Load 从文件加载合约地址
func (s *ContractStorage) Load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.Contracts)
}

// Save 保存合约地址到文件
func (s *ContractStorage) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := json.MarshalIndent(s.Contracts, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化合约数据失败: %v", err)
	}

	return os.WriteFile(s.filePath, data, 0644)
}

// StoreContract 存储合约地址
func (s *ContractStorage) StoreContract(name string, address common.Address) error {
	s.mu.Lock()
	s.Contracts[name] = address.Hex()
	s.mu.Unlock()

	return s.Save()
}

// GetContract 获取合约地址
func (s *ContractStorage) GetContract(name string) (common.Address, bool) {

	addressHex, exists := s.Contracts[name]
	if !exists {
		return common.Address{}, false
	}

	return common.HexToAddress(addressHex), true
}

// GetAllContracts 获取所有合约地址
func (s *ContractStorage) GetAllContracts() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 创建副本避免并发问题
	result := make(map[string]string)
	for name, address := range s.Contracts {
		result[name] = address
	}

	return result
}

// ContractExists 检查合约是否存在
func (s *ContractStorage) ContractExists(name string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.Contracts[name]
	return exists
}

// Global contract storage instance
var (
	contractStorage *ContractStorage
	once            sync.Once
	initErr         error
)

// GetContractStorage 获取合约存储实例
func GetContractStorage() (*ContractStorage, error) {
	once.Do(func() {
		contractStorage, initErr = NewContractStorage()
	})
	return contractStorage, initErr
}
