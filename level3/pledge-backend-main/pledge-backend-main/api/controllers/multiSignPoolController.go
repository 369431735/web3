package controllers

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/api/services"
	"pledge-backend/api/validate"
	"pledge-backend/log"

	"github.com/gin-gonic/gin"
)

// MultiSignPoolController 多重签名池控制器
// 处理多重签名钱包相关的API请求，包括设置和获取多签信息
type MultiSignPoolController struct {
}

// SetMultiSign 设置多重签名信息
// 处理设置多签地址和相关参数的请求
// ctx: Gin上下文
func (c *MultiSignPoolController) SetMultiSign(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.SetMultiSign{} // 多签设置请求参数
	log.Logger.Sugar().Info("SetMultiSign req ", req)

	// 验证请求参数
	errCode := validate.NewMutiSign().SetMultiSign(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 参数验证失败，返回错误
		return
	}

	// 调用服务设置多签信息
	errCode, err := services.NewMutiSign().SetMultiSign(&req)
	if errCode != statecode.CommonSuccess {
		log.Logger.Error(err.Error())   // 记录错误日志
		res.Response(ctx, errCode, nil) // 设置失败，返回错误
		return
	}

	// 设置成功，返回成功状态
	res.Response(ctx, statecode.CommonSuccess, nil)
	return
}

// GetMultiSign 获取多重签名信息
// 根据链ID获取对应的多签地址和参数信息
// ctx: Gin上下文
func (c *MultiSignPoolController) GetMultiSign(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.GetMultiSign{}  // 获取多签请求参数
	result := response.MultiSign{} // 多签响应结果
	log.Logger.Sugar().Info("GetMultiSign req ", nil)

	// 验证请求参数
	errCode := validate.NewMutiSign().GetMultiSign(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 参数验证失败，返回错误
		return
	}

	// 调用服务获取多签信息
	errCode, err := services.NewMutiSign().GetMultiSign(&result, req.ChainId)
	if errCode != statecode.CommonSuccess {
		log.Logger.Error(err.Error())   // 记录错误日志
		res.Response(ctx, errCode, nil) // 获取失败，返回错误
		return
	}

	// 获取成功，返回多签信息
	res.Response(ctx, statecode.CommonSuccess, result)
	return
}
