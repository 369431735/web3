package middlewares

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/response"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/utils"

	"github.com/gin-gonic/gin"
)

// CheckToken 身份验证中间件
// 验证请求头中的认证令牌(authCode)是否有效
// 验证流程：
// 1. 从请求头获取authCode
// 2. 解析JWT令牌
// 3. 验证用户名是否为管理员
// 4. 检查Redis中是否存在有效的登录状态
// 返回: Gin中间件处理函数
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := response.Gin{Res: c}
		token := c.Request.Header.Get("authCode") // 从请求头获取认证令牌

		// 解析JWT令牌，获取用户名
		username, err := utils.ParseToken(token, config.Config.Jwt.SecretKey)
		if err != nil {
			// 令牌解析失败，返回Token错误
			res.Response(c, statecode.TokenErr, nil)
			c.Abort() // 中止请求处理流程
			return
		}

		// 验证用户名是否为管理员用户名
		if username != config.Config.DefaultAdmin.Username {
			res.Response(c, statecode.TokenErr, nil)
			c.Abort()
			return
		}

		// 从Redis检查用户是否处于登录状态
		resByteArr, err := db.RedisGet(username)
		if string(resByteArr) != `"login_ok"` {
			// 未找到有效的登录状态，返回Token错误
			res.Response(c, statecode.TokenErr, nil)
			c.Abort()
			return
		}

		// 在上下文中保存用户名，便于后续处理函数使用
		c.Set("username", username)

		// 继续处理请求
		c.Next()
	}
}
