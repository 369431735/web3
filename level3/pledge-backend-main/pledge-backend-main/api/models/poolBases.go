package models

import (
	"encoding/json"
	"pledge-backend/db"
)

// PoolBaseInfo 质押池基础信息响应结构体
// 用于API响应，包含质押池的基本信息和代币信息
type PoolBaseInfo struct {
	PoolID                 int             `json:"pool_id"`                // 质押池ID
	AutoLiquidateThreshold string          `json:"autoLiquidateThreshold"` // 自动清算阈值
	BorrowSupply           string          `json:"borrowSupply"`           // 借入供应量
	BorrowToken            string          `json:"borrowToken"`            // 借入代币
	BorrowTokenInfo        BorrowTokenInfo `json:"borrowTokenInfo"`        // 借入代币详细信息
	EndTime                string          `json:"endTime"`                // 结束时间
	InterestRate           string          `json:"interestRate"`           // 利率
	JpCoin                 string          `json:"jpCoin"`                 // 联合提供者币种
	LendSupply             string          `json:"lendSupply"`             // 借出供应量
	LendToken              string          `json:"lendToken"`              // 借出代币
	LendTokenInfo          LendTokenInfo   `json:"lendTokenInfo"`          // 借出代币详细信息
	MartgageRate           string          `json:"martgageRate"`           // 抵押率
	MaxSupply              string          `json:"maxSupply"`              // 最大供应量
	SettleTime             string          `json:"settleTime"`             // 结算时间
	SpCoin                 string          `json:"spCoin"`                 // 服务提供者币种
	State                  string          `json:"state"`                  // 状态
}

// PoolBases 质押池基础信息数据库模型
// 用于数据库操作，存储质押池的基本配置信息
type PoolBases struct {
	Id                     int    `json:"-" gorm:"column:id;primaryKey"`                                  // 主键ID
	PoolID                 int    `json:"pool_id" gorm:"column:pool_id;"`                                 // 质押池ID
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold" gorm:"column:auto_liquidata_threshold;"` // 自动清算阈值
	BorrowSupply           string `json:"borrowSupply" gorm:"column:borrow_supply;"`                      // 借入供应量
	BorrowToken            string `json:"borrowToken" gorm:"column:pool_id;"`                             // 借入代币（注意:字段名可能有错，应该是borrow_token而不是pool_id）
	BorrowTokenInfo        string `json:"borrowTokenInfo" gorm:"column:borrow_token_info;"`               // 借入代币详细信息（JSON字符串）
	EndTime                string `json:"endTime" gorm:"end_time;"`                                       // 结束时间
	InterestRate           string `json:"interestRate" gorm:"column:interest_rate;"`                      // 利率
	JpCoin                 string `json:"jpCoin" gorm:"column:jp_coin;"`                                  // 联合提供者币种
	LendSupply             string `json:"lendSupply" gorm:"column:lend_supply;"`                          // 借出供应量
	LendToken              string `json:"lendToken" gorm:"column:lend_token;"`                            // 借出代币
	LendTokenInfo          string `json:"lendTokenInfo" gorm:"column:lend_token_info;"`                   // 借出代币详细信息（JSON字符串）
	MartgageRate           string `json:"martgageRate" gorm:"column:martgage_rate;"`                      // 抵押率
	MaxSupply              string `json:"maxSupply" gorm:"column:max_supply;"`                            // 最大供应量
	SettleTime             string `json:"settleTime" gorm:"column:settle_time;"`                          // 结算时间
	SpCoin                 string `json:"spCoin" gorm:"column:sp_coin;"`                                  // 服务提供者币种
	State                  string `json:"state" gorm:"column:state;"`                                     // 状态
}

// BorrowTokenInfo 借入代币信息结构体
// 用于表示借入代币的详细信息，包括费用、图标、名称和价格
type BorrowTokenInfo struct {
	BorrowFee  string `json:"borrowFee"`  // 借入费用
	TokenLogo  string `json:"tokenLogo"`  // 代币图标URL
	TokenName  string `json:"tokenName"`  // 代币名称
	TokenPrice string `json:"tokenPrice"` // 代币价格
}

// LendTokenInfo 借出代币信息结构体
// 用于表示借出代币的详细信息，包括费用、图标、名称和价格
type LendTokenInfo struct {
	LendFee    string `json:"lendFee"`    // 借出费用
	TokenLogo  string `json:"tokenLogo"`  // 代币图标URL
	TokenName  string `json:"tokenName"`  // 代币名称
	TokenPrice string `json:"tokenPrice"` // 代币价格
}

// PoolBaseInfoRes 质押池基础信息响应结构体
// 用于API响应，包含质押池基础数据和索引信息
type PoolBaseInfoRes struct {
	Index    int          `json:"index"`     // 索引，通常为poolID减1
	PoolData PoolBaseInfo `json:"pool_data"` // 质押池基础数据
}

// NewPoolBases 创建一个新的PoolBases实例
// 返回：
//   - *PoolBases: 质押池基础信息模型实例的指针
func NewPoolBases() *PoolBases {
	return &PoolBases{}
}

// TableName 指定表名
// 返回数据库表名，用于GORM映射
// 返回：
//   - string: 数据库表名
func (p *PoolBases) TableName() string {
	return "poolbases"
}

// PoolBaseInfo 获取指定链上的所有质押池基础信息
// 根据链ID查询并组装质押池基础信息响应
// 参数：
//   - chainId: 区块链ID，用于筛选特定链上的质押池
//   - res: 质押池基础信息响应数组指针，用于存储查询结果
//
// 返回：
//   - error: 错误信息，如果有的话；nil表示操作成功
func (p *PoolBases) PoolBaseInfo(chainId int, res *[]PoolBaseInfoRes) error {
	var poolBases []PoolBases

	// 查询指定链ID的质押池基础信息，按池ID升序排序
	err := db.Mysql.Table("poolbases").Where("chain_id=?", chainId).Order("pool_id asc").Find(&poolBases).Debug().Error
	if err != nil {
		return err
	}

	// 遍历池基础信息，解析代币信息并组装响应数据
	for _, v := range poolBases {
		// 解析借入代币信息（从JSON字符串转为结构体）
		borrowTokenInfo := BorrowTokenInfo{}
		_ = json.Unmarshal([]byte(v.BorrowTokenInfo), &borrowTokenInfo)
		// 解析借出代币信息（从JSON字符串转为结构体）
		lendTokenInfo := LendTokenInfo{}
		_ = json.Unmarshal([]byte(v.LendTokenInfo), &lendTokenInfo)
		// 组装完整的池基础信息响应
		*res = append(*res, PoolBaseInfoRes{
			Index: v.PoolID - 1,
			PoolData: PoolBaseInfo{
				PoolID:                 v.PoolID,
				AutoLiquidateThreshold: v.AutoLiquidateThreshold,
				BorrowSupply:           v.BorrowSupply,
				BorrowToken:            v.BorrowToken,
				BorrowTokenInfo:        borrowTokenInfo,
				EndTime:                v.EndTime,
				InterestRate:           v.InterestRate,
				JpCoin:                 v.JpCoin,
				LendSupply:             v.LendSupply,
				LendToken:              v.LendToken,
				LendTokenInfo:          lendTokenInfo,
				MartgageRate:           v.MartgageRate,
				MaxSupply:              v.MaxSupply,
				SettleTime:             v.SettleTime,
				SpCoin:                 v.SpCoin,
				State:                  v.State,
			},
		})
	}
	return nil
}
