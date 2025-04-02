package validate

import (
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// MutiSign 多重签名相关请求的验证器结构体
// 负责验证多重签名设置和获取请求参数的有效性
type MutiSign struct{}

// NewMutiSign 创建一个新的MutiSign验证器实例
// 返回：
//   - *MutiSign: 多重签名验证器实例的指针
func NewMutiSign() *MutiSign {
	return &MutiSign{}
}

// SetMultiSign 验证设置多重签名请求参数
// 检查设置多重签名请求的必要参数是否存在，并验证链ID是否有效
// 参数：
//   - c: Gin上下文，包含HTTP请求信息
//   - req: 设置多重签名请求结构体指针，将被绑定到请求体
//
// 返回：
//   - int: 状态码，表示验证结果
//   - statecode.CommonSuccess: 验证成功
//   - statecode.ChainIdErr: 链ID无效（不是97或56）
//   - statecode.ParameterEmptyErr: 请求体为空
//   - statecode.PNameEmpty: SP名称为空
//   - statecode.CommonErrServerErr: 其他错误
func (v *MutiSign) SetMultiSign(c *gin.Context, req *request.SetMultiSign) int {

	err := c.ShouldBind(req)
	// 先验证链ID是否有效
	if req.ChainId != 97 && req.ChainId != 56 {
		return statecode.ChainIdErr
	}
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "SpName" && e.Tag() == "required" {
				return statecode.PNameEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	return statecode.CommonSuccess
}

// GetMultiSign 验证获取多重签名请求参数
// 检查获取多重签名请求的必要参数是否存在，并验证链ID是否有效
// 参数：
//   - c: Gin上下文，包含HTTP请求信息
//   - req: 获取多重签名请求结构体指针，将被绑定到请求体
//
// 返回：
//   - int: 状态码，表示验证结果
//   - statecode.CommonSuccess: 验证成功
//   - statecode.ChainIdErr: 链ID无效（不是97或56）
//   - statecode.ParameterEmptyErr: 请求体为空
//   - statecode.ChainIdEmpty: 链ID为空
//   - statecode.CommonErrServerErr: 其他错误
func (v *MutiSign) GetMultiSign(c *gin.Context, req *request.GetMultiSign) int {

	err := c.ShouldBind(req)
	// 先验证链ID是否有效
	if req.ChainId != 97 && req.ChainId != 56 {
		return statecode.ChainIdErr
	}
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

	return statecode.CommonSuccess
}
