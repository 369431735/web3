package response

import (
	"pledge-backend/api/common/statecode"

	"github.com/gin-gonic/gin"
)

// Gin 响应处理封装
// 封装Gin上下文，提供统一的响应方法
type Gin struct {
	Res *gin.Context // 当前请求的Gin上下文
}

// Page 分页响应结构
// 适用于需要返回总数据量的分页数据
type Page struct {
	Code  int         `json:"code"`    // 状态码
	Msg   string      `json:"message"` // 状态消息
	Total int         `json:"total"`   // 数据总量
	Data  interface{} `json:"data"`    // 响应数据
}

// ResponsePages 返回分页响应
// 生成统一的分页响应格式，包含状态码、消息、总数据量和分页数据
// c: Gin上下文
// code: 状态码
// totalCount: 数据总量
// data: 分页数据
func (g *Gin) ResponsePages(c *gin.Context, code int, totalCount int, data interface{}) {
	// 获取请求语言类型，默认为简体中文
	lang := statecode.LangZh
	langInf, hasLang := c.Get("lang")
	if hasLang {
		lang = langInf.(int)
	}
	// 构造分页响应结构
	rsp := Page{
		Code:  code,                         // 状态码
		Msg:   statecode.GetMsg(code, lang), // 根据状态码和语言获取消息
		Total: totalCount,                   // 数据总量
		Data:  data,                         // 分页数据
	}
	// 返回JSON响应
	g.Res.JSON(200, rsp)
	return
}

// Response 返回标准响应
// 生成统一的响应格式，包含状态码、消息和数据
// c: Gin上下文
// code: 状态码
// data: 响应数据
// httpStatus: 可选的HTTP状态码，默认为200
func (g *Gin) Response(c *gin.Context, code int, data interface{}, httpStatus ...int) {
	// 获取请求语言类型，默认为英文
	lang := statecode.LangEn
	langInf, hasLang := c.Get("lang")
	if hasLang {
		lang = langInf.(int)
	}
	// 构造标准响应结构
	rsp := Response{
		Code: code,                         // 状态码
		Msg:  statecode.GetMsg(code, lang), // 根据状态码和语言获取消息
		Data: data,                         // 响应数据
	}
	// 设置HTTP状态码，默认为200
	HttpStatus := 200
	if len(httpStatus) > 0 {
		HttpStatus = httpStatus[0]
	}
	// 返回JSON响应
	g.Res.JSON(HttpStatus, rsp)
	return
}

// Response 标准响应结构
// 包含状态码、消息和数据的通用响应格式
type Response struct {
	Code int         `json:"code"`    // 状态码
	Msg  string      `json:"message"` // 状态消息
	Data interface{} `json:"data"`    // 响应数据
}
