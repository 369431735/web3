package controller

// AccountBalance 账户余额请求
type AccountBalance struct {
	Address string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	Balance string `json:"balance" example:"1000000000000000000"`
}

// AccountResponse 账户响应
type AccountResponse struct {
	Address string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	Balance string `json:"balance" example:"1000000000000000000"`
}

// BlockResponse 区块响应
type BlockResponse struct {
	Number     string `json:"number" example:"123456"`
	Hash       string `json:"hash" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	ParentHash string `json:"parentHash" example:"0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"`
	Timestamp  uint64 `json:"timestamp" example:"1634567890"`
	TxCount    int    `json:"txCount" example:"10"`
}

// ContractDeployRequest 合约部署请求
type ContractDeployRequest struct {
	ContractName string `json:"contractName" example:"SimpleStorage"`
}

// ContractResponse 合约响应
type ContractResponse struct {
	Address string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	TxHash  string `json:"txHash" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"服务器内部错误"`
}
