package models

import (
	"pledge-backend/db"
)

// PoolData 质押池数据模型
// 用于存储质押池的动态数据，包括借贷金额、清算金额和结算金额等
type PoolData struct {
	Id                     int    `json:"-" gorm:"column:id;primaryKey;autoIncrement"`                     // 主键ID，自增
	PoolID                 int    `json:"pool_id" gorm:"column:pool_id;"`                                  // 质押池ID
	ChainId                string `json:"chain_id" gorm:"column:chain_id"`                                 // 区块链ID
	FinishAmountBorrow     string `json:"finish_amount_borrow" gorm:"column:finish_amount_borrow"`         // 完成的借入金额
	FinishAmountLend       string `json:"finish_amount_lend" gorm:"column:finish_amount_lend"`             // 完成的借出金额
	LiquidationAmounBorrow string `json:"liquidation_amoun_borrow" gorm:"column:liquidation_amoun_borrow"` // 清算的借入金额
	LiquidationAmounLend   string `json:"liquidation_amoun_lend" gorm:"column:liquidation_amoun_lend"`     // 清算的借出金额
	SettleAmountBorrow     string `json:"settle_amount_borrow" gorm:"column:settle_amount_borrow"`         // 结算的借入金额
	SettleAmountLend       string `json:"settle_amount_lend" gorm:"column:settle_amount_lend"`             // 结算的借出金额
	CreatedAt              string `json:"created_at" gorm:"column:created_at"`                             // 创建时间
	UpdatedAt              string `json:"updated_at" gorm:"column:updated_at"`                             // 更新时间
}

// PoolDataInfoRes 质押池数据信息响应结构体
// 用于API响应，包含质押池数据和索引信息
type PoolDataInfoRes struct {
	Index    int      `json:"index"`     // 索引，通常为poolID减1
	PoolData PoolData `json:"pool_data"` // 质押池数据
}

// NewPoolData 创建一个新的PoolData实例
// 返回：
//   - *PoolData: 质押池数据模型实例的指针
func NewPoolData() *PoolData {
	return &PoolData{}
}

// TableName 指定表名
// 返回数据库表名，用于GORM映射
// 返回：
//   - string: 数据库表名
func (p *PoolData) TableName() string {
	return "pooldata"
}

// PoolDataInfo 获取指定链上的所有质押池数据信息
// 根据链ID查询并组装质押池数据信息响应
// 参数：
//   - chainId: 区块链ID，用于筛选特定链上的质押池数据
//   - res: 质押池数据信息响应数组指针，用于存储查询结果
//
// 返回：
//   - error: 错误信息，如果有的话；nil表示操作成功
func (p *PoolData) PoolDataInfo(chainId int, res *[]PoolDataInfoRes) error {
	var poolData []PoolData

	// 查询指定链ID的质押池数据，按池ID升序排序
	err := db.Mysql.Table("pooldata").Where("chain_id=?", chainId).Order("pool_id asc").Find(&poolData).Debug().Error
	if err != nil {
		return err
	}

	// 组装响应数据
	for _, v := range poolData {
		*res = append(*res, PoolDataInfoRes{
			Index:    v.PoolID - 1,
			PoolData: v,
		})
	}
	return nil
}
