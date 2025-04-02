package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/log"
)

// poolService 质押池服务结构体
// 提供与质押池相关的业务逻辑处理，包括获取质押池基础信息和数据信息
type poolService struct{}

// NewPool 创建一个新的poolService实例
// 返回：
//   - *poolService: 质押池服务实例的指针
func NewPool() *poolService {
	return &poolService{}
}

// PoolBaseInfo 获取指定链ID的质押池基础信息
// 参数：
//   - chainId: 区块链ID，用于筛选特定链上的质押池
//   - result: 质押池基础信息结果数组指针，用于存储查询结果
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 查询成功
//   - statecode.CommonErrServerErr: 服务器错误
func (s *poolService) PoolBaseInfo(chainId int, result *[]models.PoolBaseInfoRes) int {

	err := models.NewPoolBases().PoolBaseInfo(chainId, result)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr
	}
	return statecode.CommonSuccess
}

// PoolDataInfo 获取指定链ID的质押池数据信息
// 包括质押数量、收益率等动态数据
// 参数：
//   - chainId: 区块链ID，用于筛选特定链上的质押池
//   - result: 质押池数据信息结果数组指针，用于存储查询结果
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 查询成功
//   - statecode.CommonErrServerErr: 服务器错误
func (s *poolService) PoolDataInfo(chainId int, result *[]models.PoolDataInfoRes) int {

	err := models.NewPoolData().PoolDataInfo(chainId, result)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr
	}
	return statecode.CommonSuccess
}
