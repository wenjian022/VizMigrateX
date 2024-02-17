package model

import (
	"dataxWeb/models/model/assetDatabasesModel"
	"dataxWeb/models/model/dataReplicationModel"
	"dataxWeb/models/model/userModel"
	"dataxWeb/models/util/tool/db"
)

func RegisterInit() {
	db.Registration("assetsDatabases", assetDatabasesModel.AssetsDatabasesModel{})
	db.Registration("users", userModel.UsersModelDb{})
	db.Registration("dataReplication", dataReplicationModel.DataReplicationModel{})
	db.TabTask()
	userModel.UserInit() // 用户初始化
}
