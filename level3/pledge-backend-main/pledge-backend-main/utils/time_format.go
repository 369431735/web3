package utils

import (
	"strconv"
	"time"
)

// 时间常量定义
const (
	TimeOffset = 8 * 3600  // 8小时时区偏移（东八区）
	HalfOffset = 12 * 3600 // 半天时间偏移（12小时）
)

// GetCurrentTimestampBySecond 获取当前秒级时间戳
// 返回当前时间的Unix时间戳（秒）
// 返回: 秒级时间戳
func GetCurrentTimestampBySecond() int64 {
	return time.Now().Unix()
}

// UnixSecondToTime 将秒级时间戳转换为time.Time类型
// second: 秒级时间戳
// 返回: 对应的time.Time对象
func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

// UnixNanoSecondToTime 将纳秒级时间戳转换为time.Time类型
// nanoSecond: 纳秒级时间戳
// 返回: 对应的time.Time对象
func UnixNanoSecondToTime(nanoSecond int64) time.Time {
	return time.Unix(0, nanoSecond)
}

// GetCurrentTimestampByNano 获取当前纳秒级时间戳
// 返回当前时间的纳秒级时间戳
// 返回: 纳秒级时间戳
func GetCurrentTimestampByNano() int64 {
	return time.Now().UnixNano()
}

// GetCurrentTimestampByMill 获取当前毫秒级时间戳
// 返回当前时间的毫秒级时间戳
// 返回: 毫秒级时间戳
func GetCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetCurDayZeroTimestamp 获取当天零点时间戳
// 计算当前日期的零点（00:00:00）对应的UTC时间戳
// 返回: 当天零点的秒级时间戳
func GetCurDayZeroTimestamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.Unix() - TimeOffset
}

// GetCurDayHalfTimestamp 获取当天中午12点时间戳
// 计算当前日期的中午12点对应的UTC时间戳
// 返回: 当天中午12点的秒级时间戳
func GetCurDayHalfTimestamp() int64 {
	return GetCurDayZeroTimestamp() + HalfOffset
}

// GetCurDayZeroTimeFormat 获取当天零点的格式化时间
// 将当天零点的时间戳格式化为"2006-01-02_00-00-00"格式
// 返回: 格式化后的时间字符串
func GetCurDayZeroTimeFormat() string {
	return time.Unix(GetCurDayZeroTimestamp(), 0).Format("2006-01-02_15-04-05")
}

// GetCurDayHalfTimeFormat 获取当天中午12点的格式化时间
// 将当天中午12点的时间戳格式化为"2006-01-02_12-00-00"格式
// 返回: 格式化后的时间字符串
func GetCurDayHalfTimeFormat() string {
	return time.Unix(GetCurDayZeroTimestamp()+HalfOffset, 0).Format("2006-01-02_15-04-05")
}

// GetTimeStampByFormat 根据日期时间字符串获取时间戳
// 将"2006-01-02 15:04:05"格式的日期时间字符串转换为时间戳字符串
// datetime: 日期时间字符串，格式为"2006-01-02 15:04:05"
// 返回: 时间戳字符串
func GetTimeStampByFormat(datetime string) string {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix()
	return strconv.FormatInt(timestamp, 10)
}

// TimeStringFormatTimeUnix 将指定格式的时间字符串转换为Unix时间戳
// timeFormat: 时间格式，如"2006-01-02 15:04:05"
// timeSrc: 时间字符串
// 返回: 对应的Unix时间戳
func TimeStringFormatTimeUnix(timeFormat string, timeSrc string) int64 {
	tm, _ := time.Parse(timeFormat, timeSrc)
	return tm.Unix()
}

// GetCurDateTimeFormat 获取当前时间的标准格式字符串
// 返回格式为"2006-01-02 15:04:05"的当前时间字符串
// 返回: 格式化的当前时间字符串
func GetCurDateTimeFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// NowDataTime 获取当前时间的标准格式字符串（同GetCurDateTimeFormat）
// 返回格式为"2006-01-02 15:04:05"的当前时间字符串
// 返回: 格式化的当前时间字符串
func NowDataTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
