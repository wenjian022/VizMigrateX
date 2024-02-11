package controllers

import (
	"VizMigrateX/internal/pkg/lg"
	"VizMigrateX/internal/pkg/response"
	"VizMigrateX/internal/services"
	"github.com/gin-gonic/gin"
)

var dataSourceService = new(services.DataSourceService)

type DataSourceControllers struct{}

// DataSourceListGet
//
//	@Description: 获取数据源列表
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourceListGet(ctx *gin.Context) {
	var query services.DataSourceListQueryStruct
	_ = ctx.ShouldBindQuery(&query)

	// 查询
	res, err := query.Query()
	if err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(res))
}

// DataSourceCreatePost
//
//	@Description: 创建数据源
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourceCreatePost(ctx *gin.Context) {
	var createJson services.DataSourceJsonStruct
	if err := ctx.ShouldBindJSON(&createJson); err != nil {
		lg.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	// 创建
	rowID, err := createJson.Create()
	if err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}
	ctx.JSON(response.ReturnStatus{}.Success(gin.H{"rowID": rowID}))
}

// DataSourceInfoGet
//
//	@Description: 获取数据源详情
//	@receiver c
//	@param ctx
func (c *DataSourceControllers) DataSourceInfoGet(ctx *gin.Context) {
	var uri services.DataSourceUriIDStruct
	if err := ctx.ShouldBindUri(&uri); err != nil {
		lg.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	//  查询id是否存在
	if err := uri.Query(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(uri.DataSourceModel))
}

// DataSourcePut
//
//	@Description: 修改数据源
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourcePut(ctx *gin.Context) {
	var uri services.DataSourceUriIDStruct
	if err := ctx.ShouldBindUri(&uri); err != nil {
		lg.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	var dataSourceJson services.DataSourceJsonStruct
	if err := ctx.ShouldBindJSON(&dataSourceJson); err != nil {
		lg.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	//  查询id是否存在
	if err := uri.Query(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	// 更新
	if err := uri.Put(dataSourceJson); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// DataSourceDelete
//
//	@Description: 删除数据源
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourceDelete(ctx *gin.Context) {
	var uri services.DataSourceUriIDStruct
	if err := ctx.ShouldBindUri(&uri); err != nil {
		lg.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	//  查询id是否存在
	if err := uri.Query(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	// 删除
	if err := uri.Del(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// DataSourceTestConnectionPost
//
//	@Description: 测试连接
//	@receiver c
//	@param context
func (c *DataSourceControllers) DataSourceTestConnectionPost(ctx *gin.Context) {
	var testConnectionJson services.DataSourceTestConnectionJsonStruct
	if err := ctx.ShouldBindJSON(&testConnectionJson); err != nil {
		lg.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	// 连接数据库
	if err := testConnectionJson.TestConnection(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// LabelListGet
//
//	@Description: 获取标签列表
//	@receiver c
//	@param ctx
func (c *DataSourceControllers) LabelListGet(ctx *gin.Context) {
	var query services.LabelQueryStruct
	_ = ctx.ShouldBindQuery(&query)

	res, err := query.Query()
	if err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(res))
}

// LabelCreatePost
//
//	@Description: 创建标签列表
//	@receiver c
//	@param ctx
func (c *DataSourceControllers) LabelCreatePost(ctx *gin.Context) {
	var createJson services.LabelCreateJsonStruct
	if err := ctx.ShouldBindJSON(&createJson); err != nil {
		lg.Logger.Errorln(err)
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	if err := createJson.Create(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// LabelDel
//
//	@Description: 删除标签列表
//	@receiver c
//	@param ctx
func (c *DataSourceControllers) LabelDel(ctx *gin.Context) {
	var uri services.LabelUriLabelIdStruct
	if err := ctx.ShouldBindUri(&uri); err != nil {
		lg.Logger.Errorln(err)
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	if err := uri.Del(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// EnvListGet
//
//	@Description: 获取环境列表
//	@receiver c
//	@param ctx
func (c *DataSourceControllers) EnvListGet(ctx *gin.Context) {
	var query services.EnvQueryStruct
	_ = ctx.ShouldBindQuery(&query)

	res, err := query.Query()
	if err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(res))
}

// EnvCreatePost
//
//	@Description: 创建环境列表
//	@receiver c
//	@param ctx
func (c *DataSourceControllers) EnvCreatePost(ctx *gin.Context) {
	var createJson services.EnvCreateJsonStruct
	if err := ctx.ShouldBindJSON(&createJson); err != nil {
		lg.Logger.Errorln(err)
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	if err := createJson.Create(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// EnvDel
//
//	@Description: 删除环境列表
//	@receiver c
//	@param ctx
func (c *DataSourceControllers) EnvDel(ctx *gin.Context) {
	var uri services.EnvUriEnvIdStruct
	if err := ctx.ShouldBindUri(&uri); err != nil {
		lg.Logger.Errorln(err)
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	if err := uri.Del(); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}
