package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/log"
)

// poolService 质押池服务结构体
// 提供与质押池相关的业务逻辑处理，包括获取池基本信息和动态数据
// 质押池是系统的核心功能，允许用户质押和借贷资产
// 本服务提供对池信息的查询和管理接口
type poolService struct{}

// NewPool 创建一个新的poolService实例
// 工厂方法模式，用于获取质押池服务的实例
// 返回：
//   - *poolService: 质押池服务实例的指针
func NewPool() *poolService {
	return &poolService{}
}

// PoolBaseInfo 获取质押池基本信息
// 根据链ID获取质押池的静态基本信息，如池名称、代币信息、创建时间等
// 这些信息通常在池创建时确定，不会频繁变化
// 参数：
//   - req: 请求结构体指针，包含链ID等查询参数
//   - result: 响应结构体指针，用于返回池基本信息
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 查询成功
//   - statecode.CommonErrServerErr: 服务器错误，如数据库操作失败
func (p *poolService) PoolBaseInfo(req *request.PoolBaseInfo, result *response.PoolBaseInfoList) int {
	// 记录查询请求的日志信息
	log.Logger.Sugar().Info("poolService.PoolBaseInfo", req)

	// 从数据库获取质押池基本信息
	poolBaseInfoList, err := models.NewPledgePoolInfoBase().PoolBaseInfo(req.ChainId)
	if err != nil {
		// 记录错误并返回服务器错误状态码
		log.Logger.Error("poolService.PoolBaseInfo error :" + err.Error())
		return statecode.CommonErrServerErr
	}

	// 填充响应结构体
	result.PoolBaseInfoList = poolBaseInfoList
	return statecode.CommonSuccess
}

// PoolDataInfo 获取质押池动态数据信息
// 根据链ID获取质押池的动态数据，如当前质押量、借款量、利率等
// 这些数据会随着用户交互和市场变化而频繁更新
// 参数：
//   - req: 请求结构体指针，包含链ID等查询参数
//   - result: 响应结构体指针，用于返回池动态数据
//
// 返回：
//   - int: 状态码，表示操作结果
//   - statecode.CommonSuccess: 查询成功
//   - statecode.CommonErrServerErr: 服务器错误，如数据库操作失败
func (p *poolService) PoolDataInfo(req *request.PoolDataInfo, result *response.PoolDataInfoList) int {
	// 记录查询请求的日志信息
	log.Logger.Sugar().Info("poolService.PoolDataInfo", req)

	// 从数据库获取质押池动态数据
	poolDataInfo, err := models.NewPledgePoolInfo().PoolDataInfo(req.ChainId)
	if err != nil {
		// 记录错误并返回服务器错误状态码
		log.Logger.Error("poolService.PoolDataInfo error :" + err.Error())
		return statecode.CommonErrServerErr
	}

	// 填充响应结构体
	result.PoolDataInfoList = poolDataInfo
	return statecode.CommonSuccess
}
