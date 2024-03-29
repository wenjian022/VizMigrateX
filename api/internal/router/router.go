package router

import (
	"VizMigrateX/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	// Global middlewares
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())
	Router.Use(middlewares.JWTAuthMiddleware())
	api := Router.Group("/api")
	{
		// public routes, no auth required
		LoadPublicRoutes(api)

		// user routes
		LoadUserRoutes(api)

		// 数据源管理
		LoadDataSourceRoutes(api)

		// 数据复制
		LoadDataReplicationRoutes(api)

		// 备份与恢复
		LoadBackupAndRecoveryRoutes(api)

	}
}
