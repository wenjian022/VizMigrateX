package datax

import (
	"dataxWeb/models/view/datax/assetView"
	"dataxWeb/models/view/datax/dataReplicationView"
	"github.com/gin-gonic/gin"
)

func Url(api gin.RouterGroup) {
	// 获取数据库资产列表
	api.GET("asset/database/list", assetView.AssetDatabaseListGet)
	// 资产数据库-选择器数据格式返回
	api.GET("asset/databases/select", assetView.AssetsDatabasesSelectGet)
	// 添加数据库资产
	api.POST("asset/databases", assetView.AssetDatabasesPost)
	// 更新数据库资产
	api.PUT("asset/database/:assetsID", assetView.AssetDatabasesPut)
	// 获取数据库资产-库信息
	api.GET("asset/database/information/:assetsID", assetView.AssetDatabasesInformationGet)
	// 获取资产库-表信息
	api.GET("asset/database/:assetsID/:database/tables", assetView.AssetsDatabasesTablesGet)
	// 数据库连接测试
	api.POST("asset/connection/test", assetView.AssetsConnectionTestPost)
	// 数据复制任务列表
	api.GET("data/replication/task/list", dataReplicationView.DataReplicationTaskListGet)
	// 获取任务详情
	api.GET("data/replication/task/:taskID", dataReplicationView.DataReplicationTaskInfoGet)
	// 添加数据复制任务
	api.POST("data/replication/task", dataReplicationView.DataReplicationTaskPost)
	// 修改数据复制任务
	api.PUT("data/replication/task/:step/:taskID", dataReplicationView.DataReplicationTaskPut)
}
