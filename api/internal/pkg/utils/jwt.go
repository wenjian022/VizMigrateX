package utils

import (
	"fmt"
	"time"

	"VizMigrateX/configs"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims

	Uid      uint   `json:"uid"`
	Username string `json:"username"`
}

// GenerateToken
//
//	@Description: 生成JWT
//	@param claims
//	@return string
func GenerateToken(claims *Claims) (string, error) {
	envConfig := configs.EnvConfig

	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 30)), // 登录30 天过期
		Issuer:    envConfig.Jwt.Secret,
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(envConfig.Jwt.Secret))
}

// JwtVerify
//
//	@Description:
//	@param tokenStr
//	@return *Claims
//	@return error
func JwtVerify(tokenStr string) (*Claims, error) {
	envConfig := configs.EnvConfig

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(envConfig.Jwt.Secret), nil
	})

	if !token.Valid || err != nil {
		return nil, fmt.Errorf("token invalid")
	}
	claims, ok := token.Claims.(*Claims)

	if float64(claims.ExpiresAt.Unix()) < float64(time.Now().Unix()) {
		return nil, fmt.Errorf("token expired")
	}

	if !ok {
		return nil, err
	}
	return claims, err

}
