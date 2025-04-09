package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
)

// TokenList 代币列表服务结构体
// 提供与代币信息相关的业务逻辑处理，包括获取债务代币和所有代币列表
// 代币信息是质押池系统的基础数据，为用户提供可质押和可借贷的代币选择
// 支持按链ID和代币类型进行筛选，满足不同区块链网络的需求
type TokenList struct{}

// NewTokenList 创建一个新的TokenList服务实例
// 工厂方法模式，用于获取代币列表服务的实例
// 返回：
//   - *TokenList: 代币列表服务实例的指针
func NewTokenList() *TokenList {
	return &TokenList{}
}

// DebtTokenList 获取债务代币列表
// 根据请求条件查询债务代币信息，通常用于展示可借贷的代币
// 债务代币是用户在质押池中可以借入的代币类型
// 参数：
//   - req: 代币列表请求结构体指针，包含以下字段：
//   - ChainID: 区块链ID，用于指定查询哪个区块链上的代币
//   - TokenType: 代币类型，可选项，用于进一步筛选代币类别
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 查询成功
//   - statecode.CommonErrServerErr: 服务器错误，如数据库查询失败
//   - []models.TokenInfo: 代币信息数组，包含代币名称、符号、地址、小数位数等详细信息
func (c *TokenList) DebtTokenList(req *request.TokenList) (int, []models.TokenInfo) {
	// 调用模型层获取代币信息
	err, res := models.NewTokenInfo().GetTokenInfo(req)
	if err != nil {
		// 返回服务器错误状态码
		return statecode.CommonErrServerErr, nil
	}
	// 返回成功状态码和查询结果
	return statecode.CommonSuccess, res
}

// GetTokenList 获取所有代币列表
// 根据请求条件查询所有可用代币，包括可质押和可借贷的代币
// 提供完整的代币信息，用于前端展示代币选择下拉列表或代币详情页
// 参数：
//   - req: 代币列表请求结构体指针，包含以下字段：
//   - ChainID: 区块链ID，用于指定查询哪个区块链上的代币
//   - TokenType: 代币类型，可选项，用于筛选特定类型的代币（如ERC20、ERC721等）
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 查询成功
//   - statecode.CommonErrServerErr: 服务器错误，如数据库查询失败
//   - []models.TokenList: 代币列表数组，包含代币详细信息，如代币ID、名称、符号、合约地址等
func (c *TokenList) GetTokenList(req *request.TokenList) (int, []models.TokenList) {
	// 调用模型层获取完整代币列表
	err, tokenList := models.NewTokenInfo().GetTokenList(req)
	if err != nil {
		// 返回服务器错误状态码
		return statecode.CommonErrServerErr, nil
	}
	// 返回成功状态码和查询结果
	return statecode.CommonSuccess, tokenList
}
