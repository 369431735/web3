package validate

import (
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TokenList 代币列表相关请求的验证器结构体
// 负责验证代币列表请求参数的有效性
type TokenList struct{}

// NewTokenList 创建一个新的TokenList验证器实例
// 返回：
//   - *TokenList: 代币列表验证器实例的指针
func NewTokenList() *TokenList {
	return &TokenList{}
}

// TokenList 验证代币列表请求参数
// 检查代币列表请求的必要参数是否存在，并验证链ID是否有效
// 参数：
//   - c: Gin上下文，包含HTTP请求信息
//   - req: 代币列表请求结构体指针，将被绑定到请求体
//
// 返回：
//   - int: 状态码，表示验证结果
//   - statecode.CommonSuccess: 验证成功
//   - statecode.ParameterEmptyErr: 请求体为空
//   - statecode.ChainIdEmpty: 链ID为空
//   - statecode.ChainIdErr: 链ID无效（不是97或56）
//   - statecode.CommonErrServerErr: 其他错误
func (v *TokenList) TokenList(c *gin.Context, req *request.TokenList) int {

	err := c.ShouldBind(req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "ChainId" && e.Tag() == "required" {
				return statecode.ChainIdEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	// 验证链ID是否为有效值（仅支持BSC主网56和测试网97）
	if req.ChainId != 97 && req.ChainId != 56 {
		return statecode.ChainIdErr
	}

	return statecode.CommonSuccess
}
