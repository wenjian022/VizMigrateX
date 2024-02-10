package router

import (
	"VizMigrateX/internal/controllers"
	"github.com/gin-gonic/gin"
)

var dataSourceController = new(controllers.DataSourceControllers)

func LoadDataSourceRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	dataSource := r.Group("/dataSource")
	{
		// 获取数据源列表
		dataSource.GET("", dataSourceController.DataSourceListGet)
		// 创建数据源
		dataSource.POST("", dataSourceController.DataSourceCreatePost)
		// 获取数据源详情
		dataSource.GET("/:dataSourceID", dataSourceController.DataSourceInfoGet)
		// 修改数据源
		dataSource.PUT("/:dataSourceID", dataSourceController.DataSourcePut)
		// 删除数据源
		dataSource.DELETE("/:dataSourceID", dataSourceController.DataSourceDelete)
		// 测试连接
		dataSource.POST("/testConnection", dataSourceController.DataSourceTestConnectionPost)
	}
	return dataSource
}
