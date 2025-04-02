package validate

import (
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Search 搜索相关请求的验证器结构体
// 负责验证搜索请求参数的有效性
type Search struct{}

// NewSearch 创建一个新的Search验证器实例
// 返回：
//   - *Search: 搜索验证器实例的指针
func NewSearch() *Search {
	return &Search{}
}

// Search 验证搜索请求参数
// 检查搜索请求的必要参数是否存在，并验证链ID是否有效
// 参数：
//   - c: Gin上下文，包含HTTP请求信息
//   - req: 搜索请求结构体指针，将被绑定到请求体
//
// 返回：
//   - int: 状态码，表示验证结果
//   - statecode.CommonSuccess: 验证成功
//   - statecode.ParameterEmptyErr: 请求体为空
//   - statecode.ChainIdEmpty: 链ID为空
//   - statecode.ChainIdErr: 链ID无效（不是97或56）
//   - statecode.CommonErrServerErr: 其他错误
func (s *Search) Search(c *gin.Context, req *request.Search) int {

	err := c.ShouldBindJSON(req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "ChainID" && e.Tag() == "required" {
				return statecode.ChainIdEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	// 验证链ID是否为有效值（仅支持BSC主网56和测试网97）
	if req.ChainID != 97 && req.ChainID != 56 {
		return statecode.ChainIdErr
	}

	return statecode.CommonSuccess
}
