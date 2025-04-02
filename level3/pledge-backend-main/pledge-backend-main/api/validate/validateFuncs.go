package validate

import (
	"fmt"
	"pledge-backend/db"
	"regexp"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

// 缓存用于OnlyOne验证的参数值
var oneofValsCache = map[string][]string{}

// 读写锁，用于保护oneofValsCache的并发访问
var oneofValsCacheRWLock = sync.RWMutex{}

// 用于分割参数的正则表达式字符串，匹配引号内的内容或非空白字符
var splitParamsRegexString = `'[^']*'|\S+`

// 预编译的正则表达式，用于分割参数
var splitParamsRegex = regexp.MustCompile(splitParamsRegexString)

// CheckUserNicknameLength 检查用户昵称长度是否合法
// 限制1-20个字符，不同用户的昵称可以重复，第一个字符不可以是空格
// 参数：
//   - fl: 验证器字段级别，包含要验证的字段值
//
// 返回：
//   - bool: 如果昵称长度在1-20个字符之间则返回true，否则返回false
func CheckUserNicknameLength(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		//if isOk, _ := regexp.MatchString(`^\w{1,20}$`, fl.Field().Interface().(string)); isOk {
		if isOk, _ := regexp.MatchString(`^.{1,20}$`, fl.Field().Interface().(string)); isOk {
			return isOk
		}
	}
	return false
}

// CheckUserAccount 检查用户账号是否合法
// 用户账号必须以字母开头，由字母和数字组成，长度为6-20个字符
// 参数：
//   - fl: 验证器字段级别，包含要验证的字段值
//
// 返回：
//   - bool: 如果账号格式正确则返回true，否则返回false
func CheckUserAccount(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^[A-Za-z][A-Za-z0-9]{5,19}$`, fl.Field().Interface().(string)); isOk {
			return isOk
		}
		return false
	}
	return false
}

// IsPassword 判断是否为合法密码
// 密码必须由字母、数字或特定特殊字符组成，长度为6-20个字符
// 参数：
//   - fl: 验证器字段级别，包含要验证的字段值
//
// 返回：
//   - bool: 如果密码格式正确则返回true，否则返回false
func IsPassword(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9!@#￥%^&*]{6,20}$`, fl.Field().Interface().(string)); isOk {
			return isOk
		}
	}
	return false
}

// IsPhoneNumber 检查手机号码字段是否合法
// 手机号码必须是以1开头的11位数字
// 参数：
//   - fl: 验证器字段级别，包含要验证的字段值
//
// 返回：
//   - bool: 如果手机号码格式正确则返回true，否则返回false
func IsPhoneNumber(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^1[0-9]{10}$`, fl.Field().Interface().(string)); isOk {
			return isOk
		}
	}
	return false
}

// IsEmail 检查邮箱地址是否合法
// 邮箱必须符合标准邮箱格式，包含@和域名部分
// 参数：
//   - fl: 验证器字段级别，包含要验证的字段值
//
// 返回：
//   - bool: 如果邮箱格式正确则返回true，否则返回false
func IsEmail(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, fl.Field().Interface().(string)); isOk {
			return isOk
		}
	}
	return false
}

// dataStruct 用于数据库查询结果的结构体
type dataStruct struct {
	DataCount int // 这个结构体用来保存查询到的记录条数
}

// OnlyOne 字段唯一性约束
// 检查字段值在指定表的指定列中是否已存在，用于确保数据唯一性
// 参数格式: 'tableName' 'fieldName'，例如: 'users' 'email'
// 参数：
//   - fl: 验证器字段级别，包含要验证的字段值和参数
//
// 返回：
//   - bool: 如果字段值在数据库中不存在则返回true，否则返回false
func OnlyOne(fl validator.FieldLevel) bool {
	vals := parseOnlyOneParam(fl.Param())
	if len(vals) <= 0 {
		panic("OnlyOne parameter err")
	}
	tableName := vals[0]
	fieldName := vals[1]

	var data dataStruct // 用于接收结构体
	sqlStr := fmt.Sprintf("`%s`=?", fieldName)
	db.Mysql.Table(tableName).Select("COUNT(*)").Where(sqlStr, fl.Field().Interface()).Scan(&data.DataCount)

	if data.DataCount > 0 {
		// 如果记录总数大于0 说明存在记录，直接返回false即可
		return false
	}
	// 没触发false就说明没有重复记录，返回true
	return true
}

// parseOnlyOneParam 解析OnlyOne验证器的参数
// 将参数字符串解析为表名和字段名
// 参数：
//   - s: 参数字符串，格式为'tableName' 'fieldName'
//
// 返回：
//   - []string: 解析后的参数数组，包含表名和字段名
func parseOnlyOneParam(s string) []string {
	oneofValsCacheRWLock.RLock()
	vals, ok := oneofValsCache[s]
	oneofValsCacheRWLock.RUnlock()
	if !ok {
		oneofValsCacheRWLock.Lock()
		vals = splitParamsRegex.FindAllString(s, -1)
		for i := 0; i < len(vals); i++ {
			vals[i] = strings.Replace(vals[i], "'", "", -1)
		}
		oneofValsCache[s] = vals
		oneofValsCacheRWLock.Unlock()
	}
	return vals
}
