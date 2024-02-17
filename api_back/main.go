package main

import (
	"dataxWeb/models/model"
	"dataxWeb/models/util/middleware"
	"dataxWeb/models/view"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"strings"
)

//go:embed web/*
var webStatic embed.FS

func main() {
	r := gin.Default()
	// 中间件配置
	//r.Use(middleware.CrossConfig, middleware.PanicMiddleware())
	r.Use(middleware.CrossConfig)
	// 路由
	api := r.Group("/api")
	view.Route(*api)
	//  当 API 不存在时，返回静态文件
	r.NoRoute(func(c *gin.Context) {
		webPath := c.Request.URL.Path                                      // 获取请求路径
		s := strings.Split(webPath, ".")                                   // 分割路径，获取文件后缀
		prefix := "web"                                                    // 前缀路径
		if data, err := webStatic.ReadFile(prefix + webPath); err != nil { // 读取文件内容
			// 如果文件不存在，返回首页 index.html
			if data, err = webStatic.ReadFile(prefix + "/index.html"); err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"err": err,
				})
			} else {
				c.Data(http.StatusOK, mime.TypeByExtension(".html"), data)
			}
		} else {
			// 如果文件存在，根据请求的文件后缀，设置正确的mime type，并返回文件内容
			c.Data(http.StatusOK, mime.TypeByExtension(fmt.Sprintf(".%s", s[len(s)-1])), data)
		}
	})

	// 数据库初始化
	model.RegisterInit()
	_ = r.Run(":8888")
}
