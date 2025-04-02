package controllers

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/api/services"
	"pledge-backend/api/validate"
	"pledge-backend/config"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// PoolController 质押池控制器
// 处理质押池相关的API请求，包括获取池基础信息、数据信息、代币列表等
type PoolController struct {
}

// PoolBaseInfo 获取质押池基础信息
// 根据链ID获取质押池的基础配置信息
// ctx: Gin上下文
func (c *PoolController) PoolBaseInfo(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.PoolBaseInfo{}       // 质押池基础信息请求参数
	var result []models.PoolBaseInfoRes // 质押池基础信息响应结果

	// 验证请求参数
	errCode := validate.NewPoolBaseInfo().PoolBaseInfo(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 参数验证失败，返回错误
		return
	}

	// 调用服务获取质押池基础信息
	errCode = services.NewPool().PoolBaseInfo(req.ChainId, &result)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 获取失败，返回错误
		return
	}

	// 返回质押池基础信息
	res.Response(ctx, statecode.CommonSuccess, result)
	return
}

// PoolDataInfo 获取质押池数据信息
// 根据链ID获取质押池的实时数据信息
// ctx: Gin上下文
func (c *PoolController) PoolDataInfo(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.PoolDataInfo{}       // 质押池数据信息请求参数
	var result []models.PoolDataInfoRes // 质押池数据信息响应结果

	// 验证请求参数
	errCode := validate.NewPoolDataInfo().PoolDataInfo(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 参数验证失败，返回错误
		return
	}

	// 调用服务获取质押池数据信息
	errCode = services.NewPool().PoolDataInfo(req.ChainId, &result)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 获取失败，返回错误
		return
	}

	// 返回质押池数据信息
	res.Response(ctx, statecode.CommonSuccess, result)
	return
}

// TokenList 获取代币列表
// 获取支持的代币列表信息，包括名称、符号、精度、地址等
// ctx: Gin上下文
func (c *PoolController) TokenList(ctx *gin.Context) {
	req := request.TokenList{}     // 代币列表请求参数
	result := response.TokenList{} // 代币列表响应结果

	// 验证请求参数
	errCode := validate.NewTokenList().TokenList(ctx, &req)
	if errCode != statecode.CommonSuccess {
		ctx.JSON(200, map[string]string{
			"error": "chainId error",
		}) // 参数验证失败，返回错误
		return
	}

	// 调用服务获取代币列表
	errCode, data := services.NewTokenList().GetTokenList(&req)
	if errCode != statecode.CommonSuccess {
		ctx.JSON(200, map[string]string{
			"error": "chainId error",
		}) // 获取失败，返回错误
		return
	}

	// 构建代币列表响应结果
	var BaseUrl = c.GetBaseUrl()                                     // 获取基础URL
	result.Name = "Pledge Token List"                                // 设置列表名称
	result.LogoURI = BaseUrl + "storage/img/Pledge-project-logo.png" // 设置项目Logo地址
	result.Timestamp = time.Now()                                    // 设置时间戳
	result.Version = response.Version{                               // 设置版本信息
		Major: 2,
		Minor: 16,
		Patch: 12,
	}
	// 遍历代币数据，构建响应结构
	for _, v := range data {
		result.Tokens = append(result.Tokens, response.Token{
			Name:     v.Symbol,   // 代币名称
			Symbol:   v.Symbol,   // 代币符号
			Decimals: v.Decimals, // 代币精度
			Address:  v.Token,    // 代币合约地址
			ChainID:  v.ChainId,  // 所属链ID
			LogoURI:  v.Logo,     // 代币Logo地址
		})
	}

	// 返回代币列表
	ctx.JSON(200, result)
	return
}

// Search 搜索质押池
// 根据搜索条件查询质押池列表
// ctx: Gin上下文
func (c *PoolController) Search(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.Search{}     // 搜索请求参数
	result := response.Search{} // 搜索响应结果

	// 验证请求参数
	errCode := validate.NewSearch().Search(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 参数验证失败，返回错误
		return
	}

	// 调用服务执行搜索
	errCode, count, pools := services.NewSearch().Search(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 搜索失败，返回错误
		return
	}

	// 构建搜索响应结果
	result.Rows = pools  // 搜索结果列表
	result.Count = count // 结果总数
	res.Response(ctx, statecode.CommonSuccess, result)
	return
}

// DebtTokenList 获取债务代币列表
// 获取平台支持的债务代币列表信息
// ctx: Gin上下文
func (c *PoolController) DebtTokenList(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.TokenList{} // 代币列表请求参数

	// 验证请求参数
	errCode := validate.NewTokenList().TokenList(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 参数验证失败，返回错误
		return
	}

	// 调用服务获取债务代币列表
	errCode, result := services.NewTokenList().DebtTokenList(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil) // 获取失败，返回错误
		return
	}

	// 返回债务代币列表
	res.Response(ctx, statecode.CommonSuccess, result)
	return
}

// GetBaseUrl 获取基础URL
// 根据配置的域名和端口，生成完整的基础URL
// 返回: 基础URL字符串
func (c *PoolController) GetBaseUrl() string {
	domainName := config.Config.Env.DomainName
	domainNameSlice := strings.Split(domainName, "")
	pattern := "\\d+"
	isNumber, _ := regexp.MatchString(pattern, domainNameSlice[0])
	// 如果域名以数字开头，可能是IP地址，添加端口号
	if isNumber {
		return config.Config.Env.Protocol + "://" + config.Config.Env.DomainName + ":" + config.Config.Env.Port + "/"
	}
	// 否则认为是正常域名，不添加端口号
	return config.Config.Env.Protocol + "://" + config.Config.Env.DomainName + "/"
}
