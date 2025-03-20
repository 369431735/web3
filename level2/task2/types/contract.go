package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// DeployedContracts 存储已部署的合约地址
var DeployedContracts = make(map[string]common.Address)

// RegisterContract 注册合约地址
func RegisterContract(name string, address common.Address) {
	DeployedContracts[name] = address
}

// GetDeployedContracts 获取所有已部署的合约地址
func GetDeployedContracts() map[string]string {
	result := make(map[string]string)
	for name, address := range DeployedContracts {
		result[name] = address.Hex()
	}
	return result
}

// ContractDeployRequest 合约部署请求
type ContractDeployRequest struct {
	ContractType string `json:"contractType" binding:"required" example:"SimpleStorage"`
}

// ContractResponse 合约部署响应
type ContractResponse struct {
	ContractType string `json:"contractType" example:"SimpleStorage"`
	Address      string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	TxHash       string `json:"txHash" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Error        string `json:"error,omitempty" example:"部署失败: 交易确认超时"`
}

// ContractBytecodeRequest 合约字节码请求
type ContractBytecodeRequest struct {
	ContractType string `json:"contractType" binding:"required" example:"SimpleStorage"`
}

// ContractBytecodeResponse 合约字节码响应
type ContractBytecodeResponse struct {
	ContractType string `json:"contractType" example:"SimpleStorage"`
	Address      string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	Bytecode     string `json:"bytecode" example:"0x608060405234801561001057600080fd5b50..."`
}
