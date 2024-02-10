package controllers

import (
	"VizMigrateX/internal/pkg/response"
	"VizMigrateX/internal/services"
	"github.com/gin-gonic/gin"
)

var dataSourceService = new(services.DataSourceService)

type DataSourceControllers struct{}

// DataSourceListGet
//
//	@Description: 获取数据源列表
//	@receiver *DataSourceControllers
//	@param context
func (*DataSourceControllers) DataSourceListGet(ctx *gin.Context) {
	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// DataSourceCreatePost
//
//	@Description: 创建数据源
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourceCreatePost(ctx *gin.Context) {
	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// DataSourcePut
//
//	@Description: 修改数据源
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourcePut(ctx *gin.Context) {
	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// DataSourceDelete
//
//	@Description: 删除数据源
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourceDelete(ctx *gin.Context) {
	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// DataSourceTestConnectionPost
//
//	@Description: 测试连接
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourceTestConnectionPost(ctx *gin.Context) {
	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}
