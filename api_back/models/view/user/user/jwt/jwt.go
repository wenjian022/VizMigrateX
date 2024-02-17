package jwt

import (
	model "dataxWeb/models/model/userModel"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type Claims struct {
	model.UsersModelDb
	jwt.RegisteredClaims
}

var secret = []byte("Centimani") // 指定加密密钥

// GenToken 生成JWT
func GenToken(userInfo model.UsersModelDb) (string, error) {
	// 创建一个我们自己的声明
	c := Claims{
		userInfo,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 30)), // 登录30 天过期
			Issuer:    "Newings",
		},
		// jwt.StandardClaims{
		// 	ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // 登录30 天过期
		// 	Issuer:    "Newings",                                  // 签发人
		// },
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("无效的token")
}
