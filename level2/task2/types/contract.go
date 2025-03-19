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
	ContractName string `json:"contract_name" binding:"required"`
}

// ContractResponse 合约部署响应
type ContractResponse struct {
	Address string `json:"address"`
	TxHash  string `json:"tx_hash"`
}

// ContractBytecodeRequest 获取合约字节码请求
type ContractBytecodeRequest struct {
	Address string `json:"address" binding:"required"`
}

// ContractBytecodeResponse 获取合约字节码响应
type ContractBytecodeResponse struct {
	Bytecode string `json:"bytecode"`
}
