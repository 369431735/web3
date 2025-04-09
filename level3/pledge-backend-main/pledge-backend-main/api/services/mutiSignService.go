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
// 多重签名是一种安全机制，要求交易必须由多个私钥签名才能执行
// 在本系统中用于保障质押池操作的安全性，特别是涉及资金的操作
type MutiSignService struct{}

// NewMutiSign 创建一个新的MutiSignService实例
// 工厂方法模式，用于获取多签服务的实例
// 返回：
//   - *MutiSignService: 多重签名服务实例的指针
func NewMutiSign() *MutiSignService {
	return &MutiSignService{}
}

// SetMultiSign 设置多重签名信息
// 将多重签名配置保存到数据库中，包括SP代币、JP代币信息和多签账户列表
// 这些配置用于后续的质押池交易验证和执行
// 参数：
//   - mutiSign: 包含多重签名配置的请求结构体指针，包括代币名称、地址、交易哈希等信息
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 设置成功
//   - statecode.CommonErrServerErr: 服务器错误
//   - error: 错误信息，如果有的话
//   - nil: 操作成功
//   - 非nil: 操作失败的错误信息
func (c *MutiSignService) SetMultiSign(mutiSign *request.SetMultiSign) (int, error) {
	//db set - 将多签配置持久化到数据库
	err := models.NewMultiSign().Set(mutiSign)
	if err != nil {
		return statecode.CommonErrServerErr, err
	}
	return statecode.CommonSuccess, nil
}

// GetMultiSign 获取多重签名信息
// 从数据库中获取指定链ID的多重签名配置，并填充到响应结构体中
// 用于前端展示当前的多签配置或者智能合约调用前的验证
// 参数：
//   - mutiSign: 用于存储多重签名信息的响应结构体指针，将被填充各项配置信息
//   - chainId: 区块链ID，用于筛选特定链上的多重签名配置，不同链可能有不同配置
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 获取成功
//   - statecode.CommonErrServerErr: 服务器错误，如数据库查询失败
//   - error: 错误信息，如果有的话
//   - nil: 操作成功
//   - 非nil: 详细的错误信息，包括数据库操作错误
func (c *MutiSignService) GetMultiSign(mutiSign *response.MultiSign, chainId int) (int, error) {
	// 从数据库获取多签配置信息
	multiSignModel := models.NewMultiSign()
	err := multiSignModel.Get(chainId)
	if err != nil {
		return statecode.CommonErrServerErr, err
	}

	// 解析多签账户JSON字符串为字符串数组
	var multiSignAccount []string
	_ = json.Unmarshal([]byte(multiSignModel.MultiSignAccount), &multiSignAccount)

	// 填充响应结构体
	// SpName/JpName: 供应代币/借款代币名称
	mutiSign.SpName = multiSignModel.SpName
	mutiSign.SpToken = multiSignModel.SpToken // 供应代币地址
	mutiSign.JpName = multiSignModel.JpName
	mutiSign.JpToken = multiSignModel.JpToken     // 借款代币地址
	mutiSign.SpAddress = multiSignModel.SpAddress // 供应代币合约地址
	mutiSign.JpAddress = multiSignModel.JpAddress // 借款代币合约地址
	mutiSign.SpHash = multiSignModel.SpHash       // 供应代币交易哈希
	mutiSign.JpHash = multiSignModel.JpHash       // 借款代币交易哈希
	mutiSign.MultiSignAccount = multiSignAccount  // 多签账户列表

	return statecode.CommonSuccess, nil
}
