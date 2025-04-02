package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	mrand "math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
)

// GetMd5String 生成32位md5字串
// 对输入字符串进行MD5哈希计算，并返回32位的十六进制字符串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// UniqueId 生成Guid字串
// 创建一个全局唯一的ID字符串，包含随机数、时间戳和MD5哈希
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	encryptedString := GetMd5String(base64.URLEncoding.EncodeToString(b))
	return encryptedString[0:16] + Int64ToString(time.Now().Unix()) + encryptedString[26:]
}

// JsonToMap 将JSON字符串转换为Map
// 解析JSON字符串并返回对应的映射表结构
func JsonToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}

// GenerateCode 生成验证码
// 生成指定位数的随机数字验证码
func GenerateCode(figures int) (randNum string) {
	startNum := math.Pow(10, float64(figures))
	number := mrand.New(mrand.NewSource(time.Now().UnixNano())).Int31n(int32(startNum))
	return fmt.Sprintf("%06d", number)
}

// IsPhone 判断是否为手机号码
// 验证字符串是否符合中国大陆11位手机号格式（1开头）
func IsPhone(phoneNo string) bool {
	if phoneNo != "" {
		if isOk, _ := regexp.MatchString(`^1[0-9]{10}$`, phoneNo); isOk {
			return isOk
		}
	}
	return false
}

// IsNumb 判断是否为数字
// 验证字符串是否只包含数字
func IsNumb(num string) bool {
	if num != "" {
		if isOk, _ := regexp.MatchString(`^[0-9]*$`, num); isOk {
			return isOk
		}
	}
	return false
}

/*
CheckAccountFormat 判断账户是否为字母开头的字母和数字组合
字母开头，限制6-20位，可以使用数字和字母
验证用户账号格式是否合法
*/
func CheckAccountFormat(s string) bool {
	if s != "" {
		isOk, _ := regexp.MatchString(`^[A-Za-z][A-Za-z0-9]{5,19}$`, s)
		if isOk {
			return isOk
		}
	}
	return false
}

// IsPassword 判断是否为合法密码
// 验证密码是否符合要求：6-20位，可包含字母、数字和特定符号
func IsPassword(pwd string) bool {
	if pwd != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9!@#￥%^&*]{6,20}$`, pwd); isOk {
			return isOk
		}
	}
	return false
}

// IsEmail 判断是否为合法邮箱
// 验证邮箱地址格式是否符合标准
func IsEmail(email string) bool {
	if email != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email); isOk {
			return isOk
		}
		//if isOk, _ := regexp.MatchString(`^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email); isOk {
		//	return isOk
		//}
	}
	return false
}

// StringToInt64 字符串转int64
// 将字符串解析为64位整数
func StringToInt64(s string) int64 {
	int64Num, _ := strconv.ParseInt(s, 10, 64)
	return int64Num
}

//字符串转float64
//func StringToFloat64(s string) float64 {
//	float, _ := strconv.ParseFloat(s, 64)
//	return float
//}

// Int64ToString int64转字符串
// 将64位整数转换为字符串
func Int64ToString(n int64) string {
	i := int64(n)
	return strconv.FormatInt(i, 10)
}

// Int32ToString int32转字符串
// 将32位整数转换为字符串
func Int32ToString(n int32) string {
	i := int64(n)
	return strconv.FormatInt(i, 10)
}

// StringToInt32 字符串转int32
// 将字符串解析为32位整数
func StringToInt32(s string) int32 {
	var j int32
	int10, _ := strconv.ParseInt(s, 10, 32)
	j = int32(int10)
	return j
}

// Int64ToInt int64转int
// 将64位整数转换为标准int类型
func Int64ToInt(n int64) int {
	strInt64 := strconv.FormatInt(n, 10)
	id16, _ := strconv.Atoi(strInt64)
	return id16
}

// Wrap 将float64转成精确的int64
// 通过将浮点数乘以10的retain次方，并转为整数，实现精确表示
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// Unwrap 将int64恢复成正常的float64
// 通过将整数除以10的retain次方，恢复为原始浮点数
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

// WrapToFloat64 精准float64
// 通过乘以10的retain次方实现精确的浮点数表示
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

// UnwrapToInt64 精准int64
// 将精确表示的int64转回普通表示
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}

// PathExists 判断文件或文件夹是否存在
// 检查指定路径的文件或目录是否存在，返回是否存在及错误信息
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetRandomString 随机生成指定位数的大小写字母和数字的组合
// 生成包含字母和数字的随机字符串，长度为n
func GetRandomString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// CreateCaptcha 生成10位随机数字
// 创建一个10位的随机数字字符串，用于验证码等场景
func CreateCaptcha() string {
	randomInt := mrand.New(mrand.NewSource(time.Now().UnixNano())).Int63n(10000000000)
	if randomInt < 1000000000 {
		randomInt = randomInt * 10
	}
	return fmt.Sprintf("%d", randomInt)
}

// Encryption 生成20位随机数字转账密码串
func Encryption() string {
	return CreateCaptcha() + CreateCaptcha()
}

func HttpGet(url string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	defer func(Body io.ReadCloser) {
		if Body != nil {
			_ = Body.Close()
		}
	}(req.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}

func HttpPost(uri string, header map[string]string, data interface{}, args ...string) ([]byte, error) {

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(req.Body)

	req.Header.Add("content-type", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}

// Float64AddToString float64 相加返回 string
func Float64AddToString(fa, fb float64) string {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Add(decimalB)
	return decimalC.String()
}

// Float64SubToString float64 相减返回 string
func Float64SubToString(fa, fb float64) string {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Sub(decimalB)
	return decimalC.String()
}

// Float64MulToString float64 相乘返回 string
func Float64MulToString(fa, fb float64) string {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Mul(decimalB)
	return decimalC.String()
}

// Float64DivToString float64 相除返回 string
func Float64DivToString(fa, fb float64) string {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Div(decimalB)
	return decimalC.String()
}

// Float64AddToFloat64 float64 相加返回 float64
func Float64AddToFloat64(fa, fb float64) float64 {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Add(decimalB)
	res, _ := decimalC.Float64()
	return res
}

// Float64SubToFloat64 float64 相减返回 float64
func Float64SubToFloat64(fa, fb float64) float64 {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Sub(decimalB)
	res, _ := decimalC.Float64()
	return res
}

// Float64MulToFloat64 float64 相乘返回 float64
func Float64MulToFloat64(fa, fb float64) float64 {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Mul(decimalB)
	res, _ := decimalC.Float64()
	return res
}

// Float64DivToFloat64 float64 相除返回 float64
func Float64DivToFloat64(fa, fb float64) float64 {
	decimalA := decimal.NewFromFloat(fa)
	decimalB := decimal.NewFromFloat(fb)
	decimalC := decimalA.Div(decimalB)
	res, _ := decimalC.Float64()
	return res
}

// Float64SubToFloat64s float64 相减返回 float64
func Float64SubToFloat64s(args ...float64) float64 {
	totalAmount := decimal.NewFromFloat(0)
	for _, val := range args {
		decimalA := decimal.NewFromFloat(val)
		totalAmount = decimalA.Sub(totalAmount)
	}
	res, _ := totalAmount.Float64()
	return res
}

// StringAddToString string 相加返回 string
func StringAddToString(sa, sb string) (string, error) {
	decimalA, err := decimal.NewFromString(sa)
	if err != nil {
		return "", err
	}
	decimalB, err := decimal.NewFromString(sb)
	if err != nil {
		return "", err
	}
	decimalC := decimalA.Add(decimalB)
	return decimalC.String(), nil
}

// StringSubToString string 相减返回 string
func StringSubToString(sa, sb string) (string, error) {
	decimalA, err := decimal.NewFromString(sa)
	if err != nil {
		return "", err
	}
	decimalB, err := decimal.NewFromString(sb)
	if err != nil {
		return "", err
	}
	decimalC := decimalA.Sub(decimalB)
	return decimalC.String(), nil
}

// StringSubStrings string 相减返回 string 多个值
func StringSubStrings(args ...string) (string, error) {
	totalAmount, _ := decimal.NewFromString("0")
	for _, val := range args {
		decimalA, err := decimal.NewFromString(val)
		if err != nil {
			continue
		}
		totalAmount = decimalA.Sub(totalAmount)
	}
	return totalAmount.String(), nil
}

// StringMulToString string 相乘返回 string
func StringMulToString(sa, sb string) (string, error) {
	decimalA, err := decimal.NewFromString(sa)
	if err != nil {
		return "", err
	}
	decimalB, err := decimal.NewFromString(sb)
	if err != nil {
		return "", err
	}
	decimalC := decimalA.Mul(decimalB)
	return decimalC.String(), nil
}

// StringDivToString string 相除返回 string
func StringDivToString(sa, sb string) (string, error) {
	decimalA, err := decimal.NewFromString(sa)
	if err != nil {
		return "", err
	}
	decimalB, err := decimal.NewFromString(sb)
	if err != nil {
		return "", err
	}
	decimalC := decimalA.Div(decimalB)
	return decimalC.String(), nil
}

func StringToFloat64(s string) float64 {
	val, err := decimal.NewFromString(s)
	if err != nil {
		return 0
	}
	res, _ := val.Float64()
	return res
}

func Float64ToString(f float64) string {
	return decimal.NewFromFloat(f).String()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ToJsonString converts any value to JSON string.
func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
