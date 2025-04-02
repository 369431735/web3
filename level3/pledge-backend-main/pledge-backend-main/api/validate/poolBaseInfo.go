package validate

import (
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// PoolBaseInfo 质押池基础信息相关请求的验证器结构体
// 负责验证获取质押池基础信息请求参数的有效性
type PoolBaseInfo struct{}

// NewPoolBaseInfo 创建一个新的PoolBaseInfo验证器实例
// 返回：
//   - *PoolBaseInfo: 质押池基础信息验证器实例的指针
func NewPoolBaseInfo() *PoolBaseInfo {
	return &PoolBaseInfo{}
}

// PoolBaseInfo 验证质押池基础信息请求参数
// 检查质押池基础信息请求的必要参数是否存在，并验证链ID是否有效
// 参数：
//   - c: Gin上下文，包含HTTP请求信息
//   - req: 质押池基础信息请求结构体指针，将被绑定到请求体
//
// 返回：
//   - int: 状态码，表示验证结果
//   - statecode.CommonSuccess: 验证成功
//   - statecode.ParameterEmptyErr: 请求体为空
//   - statecode.ChainIdEmpty: 链ID为空
//   - statecode.ChainIdErr: 链ID无效（不是97或56）
//   - statecode.CommonErrServerErr: 其他错误
func (v *PoolBaseInfo) PoolBaseInfo(c *gin.Context, req *request.PoolBaseInfo) int {
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
