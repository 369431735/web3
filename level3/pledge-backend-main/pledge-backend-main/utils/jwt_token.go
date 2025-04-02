package utils

import (
	"pledge-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken 创建JWT令牌
// 为指定用户名生成包含过期时间的JWT令牌
// username: 用户名，作为令牌的主体
// 返回: 生成的JWT令牌字符串和可能的错误
func CreateToken(username string) (string, error) {
	// 创建一个带有声明的JWT令牌
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,                                   // 设置用户名声明
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(), // 设置过期时间为30天后
	})
	// 使用密钥对令牌进行签名
	token, err := at.SignedString([]byte(config.Config.Jwt.SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解析JWT令牌
// 验证并解析JWT令牌，提取其中的用户名
// token: JWT令牌字符串
// secret: 用于验证令牌的密钥
// 返回: 令牌中的用户名和可能的错误
func ParseToken(token string, secret string) (string, error) {
	// 解析令牌，验证签名
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	// 提取并返回令牌中的用户名
	return claim.Claims.(jwt.MapClaims)["username"].(string), nil
}
