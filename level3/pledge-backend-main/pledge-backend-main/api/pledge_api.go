package main

import (
	"pledge-backend/api/middlewares"
	"pledge-backend/api/models"
	"pledge-backend/api/models/kucoin"
	"pledge-backend/api/models/ws"
	"pledge-backend/api/routes"
	"pledge-backend/api/static"
	"pledge-backend/api/validate"
	"pledge-backend/config"
	"pledge-backend/db"

	"github.com/gin-gonic/gin"
)

// main API服务主入口函数
// 负责初始化数据库连接、验证器、WebSocket服务和RESTful API服务
// 这是质押池后端的核心入口点，协调所有子系统的启动和运行
func main() {

	// 初始化MySQL数据库连接
	// 建立与MySQL数据库的连接，用于存储质押池、用户、代币等持久化数据
	db.InitMysql()

	// 初始化Redis连接
	// 创建Redis连接池，用于缓存、会话管理和分布式锁等功能
	db.InitRedis()

	// 初始化数据表结构
	// 确保数据库中存在所有必要的表，并创建缺失的表或更新表结构
	models.InitTable()

	// 绑定gin框架的参数验证器
	// 注册自定义验证器，用于请求参数的验证和转换，提高API的安全性和稳定性
	validate.BindingValidator()

	// 启动WebSocket服务器（异步）
	// 在单独的协程中启动WebSocket服务，用于向客户端推送实时数据（如价格变动）
	go ws.StartServer()

	// 从Kucoin交易所获取PLGR代币价格（异步）
	// 在单独的协程中定期从交易所获取最新代币价格，确保系统使用的价格数据保持最新
	go kucoin.GetExchangePrice()

	// 启动Gin Web框架
	// 配置并初始化Web服务，处理来自客户端的HTTP请求
	gin.SetMode(gin.ReleaseMode)                    // 设置为发布模式，减少不必要的日志输出，提高性能
	app := gin.Default()                            // 创建默认的Gin应用，包含日志记录器和崩溃恢复功能
	staticPath := static.GetCurrentAbPathByCaller() // 获取静态资源目录的绝对路径，支持不同运行环境
	app.Static("/storage/", staticPath)             // 设置静态资源路由，用于提供图片、文档等静态内容
	app.Use(middlewares.Cors())                     // 启用跨域中间件，允许不同源的客户端访问API
	routes.InitRoute(app)                           // 初始化API路由，注册所有API端点和对应的处理函数
	_ = app.Run(":" + config.Config.Env.Port)       // 在配置的端口上启动Web服务，开始接收和处理HTTP请求

}

/*
 注意：如果修改版本，需要同时修改以下文件：
 config/init.go

 这确保配置文件和API服务保持版本一致性，避免兼容性问题。
*/
