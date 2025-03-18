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
func GetDeployedContracts() map[string]common.Address {
	return DeployedContracts
}

// ContractDeployRequest 合约部署请求
type ContractDeployRequest struct {
	ContractName string `json:"contract_name" example:"SimpleStorage" binding:"required"`
}

// ContractResponse 合约部署响应
type ContractResponse struct {
	Address string `json:"address" example:"0x123..."`
	TxHash  string `json:"tx_hash" example:"0x456..."`
}

// ContractBytecodeRequest 获取合约字节码请求
type ContractBytecodeRequest struct {
	Address string `json:"address" binding:"required" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
}

// ContractBytecodeResponse 获取合约字节码响应
type ContractBytecodeResponse struct {
	Bytecode string `json:"bytecode" example:"0x608060405234801561001057600080fd..."`
}
