package controller

// ContractMethodRequest 调用合约方法的通用请求类型
type ContractMethodRequest struct {
	ContractName string        `json:"contract_name" binding:"required" example:"SimpleStorage"`
	Method       string        `json:"method" binding:"required" example:"set"`
	Params       []interface{} `json:"params" example:"[123]"`
}

// LockWithdrawRequest Lock合约提取资金的请求
type LockWithdrawRequest struct {
	// 该请求暂时不需要参数，保留结构体用于API文档生成和未来扩展
}

// 以下是已在其他文件中定义的类型和函数的注释

// SimpleStorageSetRequest SimpleStorage合约设置值的请求
// SimpleStorageGetResponse SimpleStorage合约获取值的响应
// SimpleAuctionBidRequest SimpleAuction合约竞拍的请求
// ArrayDemoAddValueRequest ArrayDemo合约添加值的请求

// getTransactOpts 获取交易选项的函数
// getContractAddress 根据合约类型获取已部署的合约地址，如果合约未部署则返回错误
