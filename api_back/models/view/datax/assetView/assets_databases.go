package assetView

import (
	"dataxWeb/models/model/assetDatabasesModel"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool"
	"dataxWeb/models/util/tool/db/dbConnection"
)

// connectionTest
//
//	@Description: 数据库连接测试
//	@param assetDatabasesJson
//	@return error
func connectionTest(assetDatabasesJson assetDatabasesModel.AssetsDatabasesModel) error {
	basis := dbConnection.BasicsDatabases{
		UserName:    assetDatabasesJson.DatabasesUserName,
		Password:    assetDatabasesJson.DatabasesPassword,
		HostAddress: assetDatabasesJson.DatabasesAddress,
		Port:        assetDatabasesJson.DatabasesPort,
		Databases:   assetDatabasesJson.Database,
		Timeout:     30,
	}
	var databases dbConnection.DatabaseInterface
	switch assetDatabasesJson.DatabasesType {
	case "mysql":
		databases = &dbConnection.MysqlOpenStruct{BasicsDatabases: basis}
	default:
		return tool.StringToErr("不支持数据库类型")
	}

	if err := databases.Open(); err != nil {
		lg.Error(err, assetDatabasesJson.DatabasesType+"测试连接失败")
		return err
	}
	defer func() { databases.Close() }()
	return nil
}
