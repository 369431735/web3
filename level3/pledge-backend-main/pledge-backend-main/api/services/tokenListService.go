package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
)

// TokenList 代币列表服务结构体
// 提供与代币信息相关的业务逻辑处理，包括获取债务代币和所有代币列表
type TokenList struct{}

// NewTokenList 创建一个新的TokenList服务实例
// 返回：
//   - *TokenList: 代币列表服务实例的指针
func NewTokenList() *TokenList {
	return &TokenList{}
}

// DebtTokenList 获取债务代币列表
// 根据请求条件查询债务代币信息
// 参数：
//   - req: 代币列表请求结构体指针，包含筛选条件
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 查询成功
//   - statecode.CommonErrServerErr: 服务器错误
//   - []models.TokenInfo: 代币信息数组，查询结果
func (c *TokenList) DebtTokenList(req *request.TokenList) (int, []models.TokenInfo) {
	err, res := models.NewTokenInfo().GetTokenInfo(req)
	if err != nil {
		return statecode.CommonErrServerErr, nil
	}
	return statecode.CommonSuccess, res

}

// GetTokenList 获取所有代币列表
// 根据请求条件查询所有可用代币
// 参数：
//   - req: 代币列表请求结构体指针，包含筛选条件
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 查询成功
//   - statecode.CommonErrServerErr: 服务器错误
//   - []models.TokenList: 代币列表数组，查询结果
func (c *TokenList) GetTokenList(req *request.TokenList) (int, []models.TokenList) {
	err, tokenList := models.NewTokenInfo().GetTokenList(req)
	if err != nil {
		return statecode.CommonErrServerErr, nil
	}
	return statecode.CommonSuccess, tokenList

}
