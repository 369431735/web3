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
type UserService struct{}

// NewUser 创建一个新的UserService实例
// 返回：
//   - *UserService: 用户服务实例的指针
func NewUser() *UserService {
	return &UserService{}
}

// Login 处理用户登录请求
// 验证用户名和密码，生成JWT令牌，并将登录状态保存到Redis
// 参数：
//   - req: 登录请求结构体指针，包含用户名和密码
//   - result: 登录响应结构体指针，用于返回令牌
//
// 返回：
//   - int: 状态码，表示登录处理结果
//   - statecode.CommonSuccess: 登录成功
//   - statecode.NameOrPasswordErr: 用户名或密码错误
//   - statecode.CommonErrServerErr: 服务器错误
func (s *UserService) Login(req *request.Login, result *response.Login) int {
	log.Logger.Sugar().Info("contractService", req)
	if req.Name == "admin" && req.Password == "password" {
		token, err := utils.CreateToken(req.Name)
		if err != nil {
			log.Logger.Error("CreateToken" + err.Error())
			return statecode.CommonErrServerErr
		}
		result.TokenId = token
		//save to redis
		_ = db.RedisSet(req.Name, "login_ok", config.Config.Jwt.ExpireTime)
		return statecode.CommonSuccess
	} else {
		return statecode.NameOrPasswordErr
	}
}
