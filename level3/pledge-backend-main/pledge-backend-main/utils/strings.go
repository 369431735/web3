package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
)

// IntToString 整数转字符串
// 将整数转换为十进制字符串表示
// i: 整数值
// 返回: 字符串表示
func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

// StringToInt 字符串转整数
// 将字符串解析为整数值
// i: 待转换的字符串
// 返回: 整数值，忽略错误
func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}

// IsContain 判断字符串是否在列表中
// 检查目标字符串是否存在于给定的字符串列表中
// target: 待查找的目标字符串
// List: 字符串列表
// 返回: 如果找到则返回true，否则返回false
func IsContain(target string, List []string) bool {
	for _, element := range List {
		if target == element {
			return true
		}
	}
	return false
}

// InterfaceArrayToStringArray 接口数组转字符串数组
// 将接口数组转换为字符串数组
// data: 接口数组
// 返回: 字符串数组
func InterfaceArrayToStringArray(data []interface{}) (i []string) {
	for _, param := range data {
		i = append(i, param.(string))
	}
	return i
}

// StructToJsonString 结构体转JSON字符串
// 将结构体序列化为JSON字符串
// param: 结构体对象
// 返回: JSON字符串
func StructToJsonString(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

// JsonStringToStruct JSON字符串转结构体
// 将JSON字符串反序列化为结构体
// s: JSON字符串
// args: 接收结果的结构体指针
// 返回: 可能的错误信息
func JsonStringToStruct(s string, args interface{}) error {
	err := json.Unmarshal([]byte(s), args)
	return err
}

// GetMsgID 生成消息ID
// 使用发送者ID、当前时间戳和随机数生成唯一的消息ID
// sendID: 发送者ID
// 返回: 生成的消息ID
func GetMsgID(sendID string) string {
	t := int64ToString(GetCurrentTimestampByNano())
	return Md5(t + sendID + int64ToString(rand.Int63n(GetCurrentTimestampByNano())))
}

// int64ToString 64位整数转字符串
// 将64位整数转换为十进制字符串表示
// i: 64位整数值
// 返回: 字符串表示
func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
