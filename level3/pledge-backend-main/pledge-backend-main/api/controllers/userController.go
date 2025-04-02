package controllers

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/api/services"
	"pledge-backend/api/validate"
	"pledge-backend/db"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
// 处理用户登录、登出等用户账户相关的请求
type UserController struct {
}

// Login 用户登录处理
// 验证用户登录请求，生成JWT令牌并返回
// ctx: Gin上下文
func (c *UserController) Login(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.Login{}     // 登录请求参数
	result := response.Login{} // 登录响应结果

	// 验证登录请求参数
	errCode := validate.NewUser().Login(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 参数验证失败，返回错误
		return
	}

	// 处理登录业务逻辑
	errCode = services.NewUser().Login(&req, &result)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 登录处理失败，返回错误
		return
	}

	// 登录成功，返回JWT令牌等信息
	res.Response(ctx, statecode.CommonSuccess, result)
	return
}

// Logout 用户登出处理
// 从Redis中删除用户登录状态
// ctx: Gin上下文
func (c *UserController) Logout(ctx *gin.Context) {
	res := response.Gin{Res: ctx}

	// 从上下文获取用户名
	usernameIntf, _ := ctx.Get("username")

	// 从Redis中删除用户登录状态
	_, _ = db.RedisDelete(usernameIntf.(string))

	// 返回登出成功
	res.Response(ctx, statecode.CommonSuccess, nil)
	return
}
