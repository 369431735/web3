package services

import (
	"fmt"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/log"
)

// SearchService 搜索服务结构体
// 提供质押池搜索功能的业务逻辑处理
type SearchService struct{}

// NewSearch 创建一个新的SearchService实例
// 返回：
//   - *SearchService: 搜索服务实例的指针
func NewSearch() *SearchService {
	return &SearchService{}
}

// Search 搜索质押池信息
// 根据提供的搜索条件过滤质押池，支持链ID、借贷代币符号和状态过滤
// 参数：
//   - req: 搜索请求结构体指针，包含搜索条件和分页信息
//
// 返回：
//   - int: 状态码，表示操作结果
//   - 0: 搜索成功
//   - statecode.CommonErrServerErr: 服务器错误
//   - int64: 符合条件的总记录数
//   - []models.Pool: 搜索结果，符合条件的质押池数组
func (c *SearchService) Search(req *request.Search) (int, int64, []models.Pool) {

	whereCondition := fmt.Sprintf(`chain_id='%v'`, req.ChainID)
	if req.LendTokenSymbol != "" {
		whereCondition += fmt.Sprintf(` and lend_token_symbol='%v'`, req.LendTokenSymbol)
	}
	if req.State != "" {
		whereCondition += fmt.Sprintf(` and state='%v'`, req.State)
	}
	err, total, data := models.NewPool().Pagination(req, whereCondition)
	if err != nil {
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr, 0, nil
	}
	return 0, total, data
}
