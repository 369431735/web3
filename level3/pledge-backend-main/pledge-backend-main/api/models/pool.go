package models

import (
	"encoding/json"
	"pledge-backend/api/models/request"
	"pledge-backend/db"
	"pledge-backend/schedule/models"
)

// Pool 质押池模型
// 包含质押池的基本信息和详细数据，用于API响应
type Pool struct {
	PoolID                 int      `json:"pool_id"`                // 质押池ID
	SettleTime             string   `json:"settleTime"`             // 结算时间
	EndTime                string   `json:"endTime"`                // 结束时间
	InterestRate           string   `json:"interestRate"`           // 利率
	MaxSupply              string   `json:"maxSupply"`              // 最大供应量
	LendSupply             string   `json:"lendSupply"`             // 借出供应量
	BorrowSupply           string   `json:"borrowSupply"`           // 借入供应量
	MartgageRate           string   `json:"martgageRate"`           // 抵押率
	LendToken              string   `json:"lendToken"`              // 借出代币
	LendTokenSymbol        string   `json:"lend_token_symbol"`      // 借出代币符号
	BorrowToken            string   `json:"borrowToken"`            // 借入代币
	BorrowTokenSymbol      string   `json:"borrow_token_symbol"`    // 借入代币符号
	State                  string   `json:"state"`                  // 状态
	SpCoin                 string   `json:"spCoin"`                 // 服务提供者币种
	JpCoin                 string   `json:"jpCoin"`                 // 联合提供者币种
	AutoLiquidateThreshold string   `json:"autoLiquidateThreshold"` // 自动清算阈值
	Pooldata               PoolData `json:"pooldata"`               // 池数据详情
}

// NewPool 创建一个新的Pool实例
// 返回：
//   - *Pool: 质押池模型实例的指针
func NewPool() *Pool {
	return &Pool{}
}

// Pagination 分页查询质押池信息
// 根据搜索条件和过滤条件获取分页后的质押池数据
// 参数：
//   - req: 搜索请求结构体指针，包含分页参数和链ID
//   - whereCondition: SQL的WHERE条件字符串，用于过滤结果
//
// 返回：
//   - error: 错误信息，如果有的话；nil表示操作成功
//   - int64: 符合条件的总记录数
//   - []Pool: 分页后的质押池数组，包含池基础信息和数据信息
func (p *Pool) Pagination(req *request.Search, whereCondition string) (error, int64, []Pool) {
	var total int64
	pools := []Pool{}
	poolBase := []models.PoolBase{}

	// 获取符合条件的总记录数
	db.Mysql.Table("poolbases").Where(whereCondition).Count(&total)

	// 分页查询池基础信息
	err := db.Mysql.Table("poolbases").Where(whereCondition).Order("pool_id desc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&poolBase).Debug().Error
	if err != nil {
		return err, 0, nil
	}

	// 遍历池基础信息，获取相应的池数据信息并组装完整的池信息
	for _, b := range poolBase {
		poolData := PoolData{}
		err = db.Mysql.Table("pooldata").Where("chain_id=?", req.ChainID).First(&poolData).Debug().Error
		if err != nil {
			return err, 0, nil
		}
		// 解析借出代币信息
		var lendToken models.LendToken
		_ = json.Unmarshal([]byte(b.LendTokenInfo), &lendToken)
		// 解析借入代币信息
		var borrowToken models.BorrowToken
		_ = json.Unmarshal([]byte(b.BorrowTokenInfo), &borrowToken)
		// 组装完整的池信息
		pools = append(pools, Pool{
			PoolID:                 b.PoolId,
			SettleTime:             b.SettleTime,
			EndTime:                b.EndTime,
			InterestRate:           b.InterestRate,
			MaxSupply:              b.MaxSupply,
			LendSupply:             b.LendSupply,
			BorrowSupply:           b.BorrowSupply,
			MartgageRate:           b.MartgageRate,
			LendToken:              lendToken.TokenName,
			BorrowToken:            borrowToken.TokenName,
			State:                  b.State,
			SpCoin:                 b.SpCoin,
			JpCoin:                 b.JpCoin,
			AutoLiquidateThreshold: b.AutoLiquidateThreshold,
			Pooldata:               poolData,
		})
	}
	return nil, total, pools
}
