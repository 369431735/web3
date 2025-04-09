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
// 支持多条件组合查询，如按链ID、代币符号、池状态等过滤
// 用于前端展示和筛选质押池列表，是用户浏览质押池的重要入口
type SearchService struct{}

// NewSearch 创建一个新的SearchService实例
// 工厂方法模式，用于获取搜索服务的实例
// 返回：
//   - *SearchService: 搜索服务实例的指针
func NewSearch() *SearchService {
	return &SearchService{}
}

// Search 搜索质押池信息
// 根据提供的搜索条件过滤质押池，支持链ID、借贷代币符号和状态过滤
// 同时支持分页查询，提高大量数据时的查询效率和用户体验
// 参数：
//   - req: 搜索请求结构体指针，包含以下字段：
//   - ChainID: 区块链ID，必填项，用于指定搜索哪个区块链上的质押池
//   - LendTokenSymbol: 借贷代币符号，可选项，用于筛选特定代币的质押池
//   - State: 质押池状态，可选项，例如"活跃"、"已结束"等状态过滤
//   - Page: 分页参数，指定当前页码
//   - PageSize: 分页参数，指定每页记录数
//
// 返回：
//   - int: 状态码，表示操作结果
//   - 0: 搜索成功
//   - statecode.CommonErrServerErr: 服务器错误，如数据库查询失败
//   - int64: 符合条件的总记录数，用于前端分页组件展示
//   - []models.Pool: 搜索结果，符合条件的质押池数组，包含池的详细信息
func (c *SearchService) Search(req *request.Search) (int, int64, []models.Pool) {
	// 构建WHERE条件，首先按链ID筛选，这是必选项
	whereCondition := fmt.Sprintf(`chain_id='%v'`, req.ChainID)

	// 如果指定了借贷代币符号，添加到查询条件中
	if req.LendTokenSymbol != "" {
		whereCondition += fmt.Sprintf(` and lend_token_symbol='%v'`, req.LendTokenSymbol)
	}

	// 如果指定了质押池状态，添加到查询条件中
	if req.State != "" {
		whereCondition += fmt.Sprintf(` and state='%v'`, req.State)
	}

	// 调用模型层的分页查询方法，获取总记录数和分页数据
	err, total, data := models.NewPool().Pagination(req, whereCondition)
	if err != nil {
		// 记录错误并返回服务器错误状态码
		log.Logger.Error(err.Error())
		return statecode.CommonErrServerErr, 0, nil
	}

	// 返回成功状态码、总记录数和查询结果
	return 0, total, data
}
