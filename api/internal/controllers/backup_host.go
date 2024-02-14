package controllers

import (
	"VizMigrateX/internal/pkg/lg"
	"VizMigrateX/internal/pkg/response"
	"VizMigrateX/internal/pkg/utils"
	"VizMigrateX/internal/services"
	"github.com/gin-gonic/gin"
)

type BackupHostControllers struct{}

// BackupHostCreatePost
//
//	@Description: 备份主机创建
//	@receiver c
//	@param ctx
func (c BackupHostControllers) BackupHostCreatePost(ctx *gin.Context) {
	var createJson services.BackupHostJsonStruct
	if err := ctx.ShouldBindJSON(&createJson); err != nil {
		lg.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	value := ctx.MustGet("USER").(*utils.Claims)

	rowID, err := createJson.Create(value.Uid)
	if err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(gin.H{"rowID": rowID}))
}

// BackupHostListGet
//
//	@Description: 备份主机列表
//	@receiver c
//	@param ctx
func (c BackupHostControllers) BackupHostListGet(ctx *gin.Context) {
	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// BackupHostInfoGet
//
//	@Description: 备份主机详情
//	@receiver c
//	@param ctx
func (c BackupHostControllers) BackupHostInfoGet(ctx *gin.Context) {
	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}
