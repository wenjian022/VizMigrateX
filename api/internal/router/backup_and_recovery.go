package router

import (
	"VizMigrateX/internal/controllers"
	"github.com/gin-gonic/gin"
)

var backupControllers = new(controllers.BackupControllers)
var backupHostControllers = new(controllers.BackupHostControllers)
var recoveryControllers = new(controllers.RecoveryControllers)

func LoadBackupAndRecoveryRoutes(r *gin.RouterGroup) {
	// 备份路由
	backupRoute := r.Group("/backup")
	{
		// 备份任务列表
		backupRoute.GET("/task/list")
		// 创建任务
		backupRoute.POST("/task")
		// 任务详情
		backupRoute.GET("/task/:taskID")
	}
	//  备份主机路由
	backupHostRoute := backupRoute.Group("/host")
	{
		// 备份主机创建
		backupHostRoute.POST("", backupHostControllers.BackupHostCreatePost)
		// 备份主机列表
		backupHostRoute.GET("/list", backupHostControllers.BackupHostListGet)
		// 备份主机详情
		backupHostRoute.GET("/:hostID", backupHostControllers.BackupHostInfoGet)
	}
	// 恢复路由
	recoveryRoute := r.Group("/recovery")
	{
		// 恢复任务列表
		recoveryRoute.GET("/task/list")
		// 创建任务
		recoveryRoute.POST("/task")
		// 任务详情
		recoveryRoute.GET("/task/:taskID")
	}
}
