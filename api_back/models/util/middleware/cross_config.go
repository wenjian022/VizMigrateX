package middleware

import (
	"github.com/gin-contrib/cors"
	"time"
)

// CrossConfig 跨域配置
var CrossConfig = cors.New(cors.Config{
	//允许跨域的网站
	AllowOrigins: []string{"*"},
	// 允许跨域的原点网站
	AllowOriginFunc: func(origin string) bool {
		return true
	},
	AllowMethods:     []string{"GET", "PUT", "DELETE", "POST", "OPTIONS"},
	AllowHeaders:     []string{"Authorization", "Content-Type", "Origin", "Access-Token"},
	AllowCredentials: true,
	ExposeHeaders:    []string{"Content-Type"},
	// 超时时间
	MaxAge: 1 * time.Hour,
	// websocket 协议
	AllowWebSockets: true,
})
