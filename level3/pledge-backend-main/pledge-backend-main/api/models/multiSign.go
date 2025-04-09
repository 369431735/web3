package models

import (
	"encoding/json"
	"errors"
	"pledge-backend/api/models/request"
	"pledge-backend/db"

	"gorm.io/gorm"
)

// MultiSign 多重签名数据模型
// 用于存储和管理多重签名相关的信息，包括服务提供者(SP)和联合提供者(JP)的信息
type MultiSign struct {
	Id               int32  `gorm:"column:id;primaryKey"`                                // 主键ID
	SpName           string `json:"sp_name" gorm:"column:sp_name"`                       // 服务提供者名称
	ChainId          int    `json:"chain_id" gorm:"column:chain_id"`                     // 区块链ID
	SpToken          string `json:"_spToken" gorm:"column:sp_token"`                     // 服务提供者代币
	JpName           string `json:"jp_name" gorm:"column:jp_name"`                       // 联合提供者名称
	JpToken          string `json:"_jpToken" gorm:"column:jp_token"`                     // 联合提供者代币
	SpAddress        string `json:"sp_address" gorm:"column:sp_address"`                 // 服务提供者地址
	JpAddress        string `json:"jp_address" gorm:"column:jp_address"`                 // 联合提供者地址
	SpHash           string `json:"spHash" gorm:"column:sp_hash"`                        // 服务提供者哈希
	JpHash           string `json:"jpHash" gorm:"column:jp_hash"`                        // 联合提供者哈希
	MultiSignAccount string `json:"multi_sign_account" gorm:"column:multi_sign_account"` // 多重签名账户列表，JSON字符串格式
}

// NewMultiSign 创建一个新的MultiSign实例
// 返回：
//   - *MultiSign: 多重签名模型实例的指针
func NewMultiSign() *MultiSign {
	return &MultiSign{}
}

// Set 设置多重签名信息
// 将多重签名配置保存到数据库中，先删除旧记录然后创建新记录
// 参数：
//   - multiSign: 包含多重签名配置的请求结构体指针
//
// 返回：
//   - error: 错误信息，如果有的话；nil表示操作成功
func (m *MultiSign) Set(multiSign *request.SetMultiSign) error {

	// 将多重签名账户列表转换为JSON字符串
	MultiSignAccountByteArr, _ := json.Marshal(multiSign.MultiSignAccount)
	// 删除同一链ID的旧记录
	err := db.Mysql.Table("multi_sign").Where("chain_id", multiSign.ChainId).Delete(&m).Debug().Error
	if err != nil {
		return errors.New("record select err " + err.Error())
	}
	// 创建新的多重签名记录
	err = db.Mysql.Table("multi_sign").Where("id=?", m.Id).Create(&MultiSign{
		ChainId:          multiSign.ChainId,
		SpName:           multiSign.SpName,
		SpToken:          multiSign.SpToken,
		JpName:           multiSign.JpName,
		JpToken:          multiSign.JpToken,
		SpAddress:        multiSign.SpAddress,
		JpAddress:        multiSign.JpAddress,
		SpHash:           multiSign.SpHash,
		JpHash:           multiSign.JpHash,
		MultiSignAccount: string(MultiSignAccountByteArr),
	}).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

// Get 获取多重签名信息
// 从数据库中获取指定链ID的多重签名配置
// 参数：
//   - chainId: 区块链ID，用于筛选特定链上的多重签名配置
//
// 返回：
//   - error: 错误信息，如果有的话；nil表示操作成功或记录不存在
func (m *MultiSign) Get(chainId int) error {
	err := db.Mysql.Table("multi_sign").Where("chain_id", chainId).First(&m).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // 记录不存在时不返回错误，而是返回空的结构体
		} else {
			return errors.New("record select err " + err.Error())
		}
	}
	return nil
}
