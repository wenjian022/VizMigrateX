package middlewares

import (
	"VizMigrateX/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Jwt middleware
// whiteList 不需要jwt验证的白名单
var whiteList = []string{"/api/user/auth/login", "/api/user/initialization"}

// JWTAuthMiddleware
//
//	@Description: jwt认证
//	@return func(c *gin.Context)
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// API白名单
		for _, s := range whiteList {
			if s == c.Request.RequestURI {
				c.Next()
				return
			}
		}
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := utils.JwtVerify(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("USER", mc)
		c.Next() // 后续的处理函数可以用过c.Get("USER")来获取当前请求的用户信息
	}

}
