package validate

import (
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// User 用户相关请求的验证器结构体
type User struct{}

// NewUser 创建一个新的User验证器实例
// 返回：
//   - *User: 用户验证器实例的指针
func NewUser() *User {
	return &User{}
}

// Login 验证用户登录请求
// 检查登录请求参数是否合法，包括用户名和密码的存在性
// 参数：
//   - c: Gin上下文，包含HTTP请求信息
//   - req: 登录请求结构体指针，将被绑定到请求体
//
// 返回：
//   - int: 状态码，表示验证结果
//   - statecode.CommonSuccess: 验证成功
//   - statecode.ParameterEmptyErr: 请求体为空
//   - statecode.PNameEmpty: 用户名为空
//   - statecode.CommonErrServerErr: 其他错误
func (v *User) Login(c *gin.Context, req *request.Login) int {

	err := c.ShouldBind(req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "Name" && e.Tag() == "required" {
				return statecode.PNameEmpty
			}
			if e.Field() == "Password" && e.Tag() == "required" {
				return statecode.PNameEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	return statecode.CommonSuccess
}
