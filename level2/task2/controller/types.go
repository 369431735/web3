package controller

// AccountBalance 账户余额请求模型
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
	ContractType string `json:"contract_type,omitempty" example:"SimpleStorage"`
	Address      string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	TxHash       string `json:"txHash" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Error        string `json:"error,omitempty" example:"交易失败"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"服务器内部错误"`
}

// TransactionResponse 交易响应
type TransactionResponse struct {
	Hash        string `json:"hash" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	From        string `json:"from" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	To          string `json:"to" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	Value       string `json:"value" example:"1000000000000000000"`
	BlockNumber string `json:"blockNumber" example:"12345"`
	Timestamp   uint64 `json:"timestamp" example:"1634567890"`
	Gas         uint64 `json:"gas" example:"21000"`
	GasPrice    string `json:"gasPrice" example:"20000000000"`
	BlockHash   string `json:"blockHash" example:"0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"`
	Status      string `json:"status" example:"成功"`
}

// ContractBytecodeRequest 合约字节码请求
type ContractBytecodeRequest struct {
	ContractType string `json:"contract_type" binding:"required" example:"SimpleStorage"`
}

// ContractBytecodeResponse 合约字节码响应
type ContractBytecodeResponse struct {
	ContractType string `json:"contract_type" example:"SimpleStorage"`
	Address      string `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"`
	Bytecode     string `json:"bytecode" example:"0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea26469706673582212209a159a4f3847890f10d0e3a00307a5b6bc6608c2f89fd2d8bd8d7e72a7c1d7d064736f6c63430008070033"`
}

// SuccessResponse 通用成功响应
type SuccessResponse struct {
	Code   int         `json:"code" example:"200"`
	Data   interface{} `json:"data"`
	ErrMsg string      `json:"errMsg" example:""`
}
