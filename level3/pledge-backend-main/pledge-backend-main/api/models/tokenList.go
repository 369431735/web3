package models

import (
	"errors"
	"pledge-backend/api/models/request"
	"pledge-backend/db"
)

// TokenInfo 代币信息数据模型
// 用于存储基本的代币信息，包括符号、合约地址和所属链ID
type TokenInfo struct {
	Id      int32  `json:"-" gorm:"column:id;primaryKey"`   // 主键ID，不在JSON中显示
	Symbol  string `json:"symbol" gorm:"column:symbol"`     // 代币符号，如ETH、BTC等
	Token   string `json:"token" gorm:"column:token"`       // 代币合约地址
	ChainId int    `json:"chain_id" gorm:"column:chain_id"` // 区块链ID，表示代币所在的区块链
}

// TokenList 扩展的代币列表数据模型
// 包含比TokenInfo更详细的代币信息，额外添加了精度和Logo
type TokenList struct {
	Id       int32  `json:"-" gorm:"column:id;primaryKey"`   // 主键ID，不在JSON中显示
	Symbol   string `json:"symbol" gorm:"column:symbol"`     // 代币符号
	Decimals int    `json:"decimals" gorm:"column:decimals"` // 代币精度，如18表示最小单位为10^-18
	Token    string `json:"token" gorm:"column:token"`       // 代币合约地址
	Logo     string `json:"logo" gorm:"column:logo"`         // 代币Logo的URL
	ChainId  int    `json:"chain_id" gorm:"column:chain_id"` // 区块链ID，表示代币所在的区块链
}

// NewTokenInfo 创建一个新的TokenInfo实例
// 返回：
//   - *TokenInfo: 代币信息模型实例的指针
func NewTokenInfo() *TokenInfo {
	return &TokenInfo{}
}

// GetTokenInfo 获取指定链上的所有代币信息
// 根据链ID查询所有代币的基本信息
// 参数：
//   - req: 代币列表请求结构体指针，包含链ID
//
// 返回：
//   - error: 错误信息，如果有的话；nil表示操作成功
//   - []TokenInfo: 代币信息数组，查询结果
func (m *TokenInfo) GetTokenInfo(req *request.TokenList) (error, []TokenInfo) {
	var tokenInfo = make([]TokenInfo, 0)
	err := db.Mysql.Table("token_info").Where("chain_id", req.ChainId).Find(&tokenInfo).Debug().Error
	if err != nil {
		return errors.New("record select err " + err.Error()), nil
	}
	return nil, tokenInfo
}

// GetTokenList 获取指定链上的所有扩展代币信息
// 根据链ID查询所有代币的详细信息，包括精度和Logo
// 参数：
//   - req: 代币列表请求结构体指针，包含链ID
//
// 返回：
//   - error: 错误信息，如果有的话；nil表示操作成功
//   - []TokenList: 扩展代币信息数组，查询结果
func (m *TokenInfo) GetTokenList(req *request.TokenList) (error, []TokenList) {
	var tokenList = make([]TokenList, 0)
	err := db.Mysql.Table("token_info").Where("chain_id", req.ChainId).Find(&tokenList).Debug().Error
	if err != nil {
		return errors.New("record select err " + err.Error()), nil
	}
	return nil, tokenList
}
