package services

import (
	"encoding/json"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
)

// MutiSignService 多重签名服务结构体
// 提供与多重签名相关的业务逻辑处理，包括设置和获取多签账户信息
type MutiSignService struct{}

// NewMutiSign 创建一个新的MutiSignService实例
// 返回：
//   - *MutiSignService: 多重签名服务实例的指针
func NewMutiSign() *MutiSignService {
	return &MutiSignService{}
}

// SetMultiSign 设置多重签名信息
// 将多重签名配置保存到数据库中
// 参数：
//   - mutiSign: 包含多重签名配置的请求结构体指针
//
// 返回：
//   - int: 状态码，表示操作结果
//   - error: 错误信息，如果有的话
//   - nil: 操作成功
//   - 非nil: 操作失败的错误信息
func (c *MutiSignService) SetMultiSign(mutiSign *request.SetMultiSign) (int, error) {
	//db set
	err := models.NewMultiSign().Set(mutiSign)
	if err != nil {
		return statecode.CommonErrServerErr, err
	}
	return statecode.CommonSuccess, nil
}

// GetMultiSign 获取多重签名信息
// 从数据库中获取指定链ID的多重签名配置，并填充到响应结构体中
// 参数：
//   - mutiSign: 用于存储多重签名信息的响应结构体指针
//   - chainId: 区块链ID，用于筛选特定链上的多重签名配置
//
// 返回：
//   - int: 状态码，表示操作结果
//   - error: 错误信息，如果有的话
//   - nil: 操作成功
//   - 非nil: 操作失败的错误信息
func (c *MutiSignService) GetMultiSign(mutiSign *response.MultiSign, chainId int) (int, error) {
	//db get
	multiSignModel := models.NewMultiSign()
	err := multiSignModel.Get(chainId)
	if err != nil {
		return statecode.CommonErrServerErr, err
	}
	var multiSignAccount []string
	_ = json.Unmarshal([]byte(multiSignModel.MultiSignAccount), &multiSignAccount)

	mutiSign.SpName = multiSignModel.SpName
	mutiSign.SpToken = multiSignModel.SpToken
	mutiSign.JpName = multiSignModel.JpName
	mutiSign.JpToken = multiSignModel.JpToken
	mutiSign.SpAddress = multiSignModel.SpAddress
	mutiSign.JpAddress = multiSignModel.JpAddress
	mutiSign.SpHash = multiSignModel.SpHash
	mutiSign.JpHash = multiSignModel.JpHash
	mutiSign.MultiSignAccount = multiSignAccount
	return statecode.CommonSuccess, nil
}
