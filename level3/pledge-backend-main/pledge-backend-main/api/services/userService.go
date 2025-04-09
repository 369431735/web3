package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/log"
	"pledge-backend/utils"
)

// UserService 用户服务结构体
// 提供用户相关的业务逻辑处理，如登录验证、用户管理等
// 在质押池系统中，用户服务负责身份验证和授权，确保只有合法用户
// 能够访问和操作系统中的敏感功能，如质押池管理和多签配置
type UserService struct{}

// NewUser 创建一个新的UserService实例
// 工厂方法模式，用于获取用户服务的单例实例
// 返回：
//   - *UserService: 用户服务实例的指针
func NewUser() *UserService {
	return &UserService{}
}

// Login 处理用户登录请求
// 验证用户名和密码，生成JWT令牌，并将登录状态保存到Redis
// JWT令牌用于后续API请求的身份验证，Redis存储确保会话状态可以跨服务实例共享
// 参数：
//   - req: 登录请求结构体指针，包含用户名和密码
//   - result: 登录响应结构体指针，用于返回JWT令牌等信息
//
// 返回：
//   - int: 状态码，表示登录处理结果
//   - statecode.CommonSuccess: 登录成功，已生成令牌并保存会话
//   - statecode.NameOrPasswordErr: 用户名或密码错误，验证失败
//   - statecode.CommonErrServerErr: 服务器错误，如令牌生成失败或Redis操作异常
func (s *UserService) Login(req *request.Login, result *response.Login) int {
	// 记录登录尝试的日志信息
	log.Logger.Sugar().Info("contractService", req)

	// 验证用户凭据 (注: 实际生产环境应该使用数据库存储用户信息，并进行密码哈希比对)
	if req.Name == "admin" && req.Password == "password" {
		// 生成JWT令牌，包含用户标识信息
		token, err := utils.CreateToken(req.Name)
		if err != nil {
			log.Logger.Error("CreateToken" + err.Error())
			return statecode.CommonErrServerErr
		}

		// 将令牌保存到响应结构中
		result.TokenId = token

		// 将登录状态保存到Redis，用于会话管理
		// 使用配置的过期时间，确保令牌自动失效
		_ = db.RedisSet(req.Name, "login_ok", config.Config.Jwt.ExpireTime)

		return statecode.CommonSuccess
	} else {
		// 用户名或密码不匹配
		return statecode.NameOrPasswordErr
	}
}
