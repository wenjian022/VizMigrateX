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
	// 环境路由
	env := dataSource.Group("/env")
	{
		// 获取环境列表
		env.GET("", dataSourceController.EnvListGet)
		// 创建环境列表
		env.POST("", dataSourceController.EnvCreatePost)
		// 删除环境列表
		env.DELETE("/:envID", dataSourceController.EnvDel)
	}

	// 标签路由
	label := dataSource.Group("/label")
	{
		// 获取标签列表
		label.GET("", dataSourceController.LabelListGet)
		// 创建标签列表
		label.POST("", dataSourceController.LabelCreatePost)
		// 删除标签列表
		label.DELETE("/:labelID", dataSourceController.LabelDel)

	}
	return dataSource
}
