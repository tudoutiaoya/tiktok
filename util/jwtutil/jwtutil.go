package jwtutil

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"tiktok/config"
	"time"
)

var secretKey []byte

// tokenExpireDuration 过期时间
const tokenExpireDuration = time.Hour * 24 * 30

// InitJwtSecretKey 初始化密钥
func InitJwtSecretKey(config *config.Configuration) {
	secretKey = []byte(config.JWTSettings.SecretKey)
}

type CustomClaims struct {
	// 可根据需要自行添加字段
	ID                   int64 `json:"id"`
	jwt.RegisteredClaims       // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(ID int64) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		ID, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)),
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secretKey)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("token不合法或token已过期")
}
