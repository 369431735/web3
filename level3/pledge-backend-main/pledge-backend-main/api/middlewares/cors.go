package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域中间件
// 处理跨域资源共享(CORS)，允许来自不同源的HTTP请求访问API
// 设置允许的请求方法、请求头以及暴露的响应头
// 返回: Gin中间件处理函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			// 设置允许所有来源访问
			c.Header("Access-Control-Allow-Origin", "*")
			// 设置允许的HTTP请求方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			// 设置允许的HTTP请求头
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, authCode, token, Content-Type, Accept, Authorization")
			// 设置可以被客户端访问的响应头
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			// 设置是否允许携带凭证信息
			c.Header("Access-Control-Allow-Credentials", "false")
			// 设置内容类型
			c.Set("content-type", "application/json")
		}

		// 对OPTIONS预检请求直接返回204状态码（无内容）
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 继续处理请求
		c.Next()
	}
}
