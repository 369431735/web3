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

// main API服务主入口函数，初始化数据库连接、启动Web服务和WebSocket服务
func main() {

	//初始化MySQL数据库连接
	db.InitMysql()

	//初始化Redis连接
	db.InitRedis()
	//初始化数据表结构
	models.InitTable()

	//绑定gin框架的参数验证器
	validate.BindingValidator()

	// 启动WebSocket服务器（异步）
	go ws.StartServer()

	// 从Kucoin交易所获取PLGR代币价格（异步）
	go kucoin.GetExchangePrice()

	// 启动Gin Web框架
	gin.SetMode(gin.ReleaseMode)                    // 设置为发布模式
	app := gin.Default()                            // 创建默认的Gin应用
	staticPath := static.GetCurrentAbPathByCaller() // 获取静态资源目录的绝对路径
	app.Static("/storage/", staticPath)             // 设置静态资源路由
	app.Use(middlewares.Cors())                     // 启用跨域中间件
	routes.InitRoute(app)                           // 初始化API路由
	_ = app.Run(":" + config.Config.Env.Port)       // 在配置的端口上启动Web服务

}

/*
 如果修改版本，需要同时修改以下文件：
 config/init.go
*/
