package routes

import (
	"pledge-backend/api/controllers"
	"pledge-backend/api/middlewares"
	"pledge-backend/config"

	"github.com/gin-gonic/gin"
)

// InitRoute 初始化API路由配置
// 配置所有API端点及其对应的控制器方法
// e: Gin引擎实例
// 返回: 配置好路由的Gin引擎实例
func InitRoute(e *gin.Engine) *gin.Engine {

	// 创建版本分组，如 /api/v2.1
	v2Group := e.Group("/api/v" + config.Config.Env.Version)

	// 质押池相关接口路由配置
	poolController := controllers.PoolController{}
	v2Group.GET("/poolBaseInfo", poolController.PoolBaseInfo)                                   // 获取质押池基础信息
	v2Group.GET("/poolDataInfo", poolController.PoolDataInfo)                                   // 获取质押池数据信息
	v2Group.GET("/token", poolController.TokenList)                                             // 获取代币列表信息
	v2Group.POST("/pool/debtTokenList", middlewares.CheckToken(), poolController.DebtTokenList) // 获取债务代币列表（需要认证）
	v2Group.POST("/pool/search", middlewares.CheckToken(), poolController.Search)               // 搜索质押池（需要认证）

	// 代币价格相关接口路由配置
	priceController := controllers.PriceController{}
	v2Group.GET("/price", priceController.NewPrice) // 从Kucoin交易所获取最新PLGR-USDT价格

	// 多重签名池相关接口路由配置
	multiSignPoolController := controllers.MultiSignPoolController{}
	v2Group.POST("/pool/setMultiSign", middlewares.CheckToken(), multiSignPoolController.SetMultiSign) // 设置多重签名（需要认证）
	v2Group.POST("/pool/getMultiSign", middlewares.CheckToken(), multiSignPoolController.GetMultiSign) // 获取多重签名信息（需要认证）

	// 用户管理相关接口路由配置
	userController := controllers.UserController{}
	v2Group.POST("/user/login", userController.Login)                             // 用户登录
	v2Group.POST("/user/logout", middlewares.CheckToken(), userController.Logout) // 用户登出（需要认证）

	return e
}
