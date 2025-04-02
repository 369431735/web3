package statecode

// 定义系统中使用的状态码常量
// 状态码用于标识API请求的处理结果，分为不同的业务类别
const (
	// 语言类型常量
	LangZh   = 111 // 简体中文
	LangEn   = 112 // 英文
	LangZhTw = 113 // 繁体中文

	// 通用状态码
	CommonSuccess      = 0    // 操作成功
	CommonErrServerErr = 1000 // 服务器错误
	ParameterEmptyErr  = 1001 // 参数为空错误

	// 认证相关状态码
	TokenErr = 1102 // 令牌错误

	// 多重签名相关状态码
	PNameEmpty   = 1201 // 签名名称为空
	ChainIdEmpty = 1202 // 链ID为空
	ChainIdErr   = 1203 // 链ID错误

	// 用户相关状态码
	NameOrPasswordErr = 1303 // 用户名或密码错误
)

// Msg 状态码消息映射
// 为每个状态码定义不同语言的错误消息
// 第一层键为状态码，第二层键为语言类型，值为对应的消息文本
var Msg = map[int]map[int]string{
	0: {
		LangZh:   "成功",
		LangZhTw: "成功",
		LangEn:   "success",
	},
	1000: {
		LangZh:   "服务器繁忙，请稍后重试",
		LangZhTw: "服務器繁忙，請稍後重試",
		LangEn:   "server is busy, please try again later",
	},
	1001: {
		LangZh:   "参数不能为空",
		LangZhTw: "参数不能為空",
		LangEn:   "parameter is empty",
	},
	1101: {
		LangZh:   "token 不能为空",
		LangZhTw: "token 不能為空",
		LangEn:   "token required",
	},
	1102: {
		LangZh:   "token错误",
		LangZhTw: "token錯誤",
		LangEn:   "token invalid",
	},
	1201: {
		LangZh:   "sp_name 不能为空",
		LangZhTw: "sp_name 不能為空",
		LangEn:   "sp_name required",
	},
	1202: {
		LangZh:   "chain_id 不能为空",
		LangZhTw: "chain_id 不能為空",
		LangEn:   "chain_id required",
	},
	1203: {
		LangZh:   "chain_id 错误",
		LangZhTw: "chain_id 錯誤",
		LangEn:   "chain_id error",
	},
	1301: {
		LangZh:   "name 不能为空",
		LangZhTw: "name 不能為空",
		LangEn:   "name required",
	},
	1302: {
		LangZh:   "password 不能为空",
		LangZhTw: "password 不能為空",
		LangEn:   "password required",
	},
	1303: {
		LangZh:   "用户名或密码错误",
		LangZhTw: "用戶名或密碼錯誤",
		LangEn:   "name or password error",
	},
}

// GetMsg 获取状态码对应的消息
// 根据状态码和语言类型，返回对应的错误消息
// c: 状态码
// lang: 语言类型
// 返回: 对应语言的错误消息，如果未找到则返回服务器错误的默认消息
func GetMsg(c int, lang int) string {
	_, ok := Msg[c]
	if ok {
		msg, ok := Msg[c][lang]
		if ok {
			return msg
		}
		return Msg[CommonErrServerErr][lang]
	}
	return Msg[CommonErrServerErr][lang]
}
