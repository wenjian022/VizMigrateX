package assetDatabasesModel

import (
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool"
	"dataxWeb/models/util/tool/db"
	"dataxWeb/models/util/tool/db/dbConnection"
	"fmt"
)

// AssetsDatabases
//
//	@Description: 获取所有数据库资产
//	@param databasesType 数据库类型 'all' 全部类型
//	@return databases
//	@return err
func AssetsDatabases(databasesType string) (databases []AssetsDatabasesModel, err error) {
	var querySQL string
	if databasesType == "all" {
		querySQL = `select * from assetsDatabases`
	} else {
		querySQL = fmt.Sprintf("select * from assetsDatabases where databasesType = '%s' ", databasesType)
	}
	err = db.Db.Select(&databases, querySQL)
	if err != nil {
		lg.SqlLogDebug(querySQL)
	}
	return
}

// uniqueQuery
//
//	@Description: 资产唯一查询
func uniqueQuery(databasesAddress, databasesUserName, assetName string, databasesPort int, id int64) (rel AssetsDatabasesModel, err error) {
	querySQL := `SELECT * FROM assetsDatabases where assetName = ? or (databasesAddress = ? and databasesPort = ? and databasesUserName = ?) and id != ?`
	err = db.Db.Get(&rel, querySQL, assetName, databasesAddress, databasesPort, databasesUserName, id)
	if err != nil {
		lg.SqlLogDebug(querySQL, assetName, databasesAddress, databasesPort, databasesUserName, id)
		lg.Error(err)
	}
	return
}

// WhereIdQuery
//
//	@Description: 通过id查询资产信息
func WhereIdQuery(id int64) (rel AssetsDatabasesModel, err error) {
	querySQL := `select * from assetsDatabases where id = ?`
	err = db.Db.Get(&rel, querySQL, id)
	if err != nil {
		lg.SqlLogDebug(querySQL, id)
	}
	return
}

// Delete
//
//	@Description: 删除资产
//	@receiver assets
//	@return error
func (databases AssetsDatabasesModel) Delete() error {
	DeleteSQL := `delete from assetsDatabases where id = ?`
	_, err := db.Db.Exec(DeleteSQL, databases.Id)
	if err != nil {
		lg.SqlLogDebug(DeleteSQL, databases.Id)
	}
	return err
}

// Crate
//
//	@Description: 创建数据库资产
//	@receiver databases
//	@return error
func (databases AssetsDatabasesModel) Crate() error {
	_, err := uniqueQuery(databases.DatabasesAddress, databases.DatabasesUserName, databases.AssetName, databases.DatabasesPort, 0)
	if db.NoRows(err) == false {
		return tool.StringToErr("重复添加")
	} else if err != nil && db.NoRows(err) == false {
		return err
	}

	// 创建
	createSQL := `insert into assetsDatabases (assetName, databasesType, databasesAddress, databasesPort, databasesUserName,databasesPassword,database,updateTime) 
						values (:assetName,:databasesType,:databasesAddress,:databasesPort,:databasesUserName,:databasesPassword,:database,datetime('now', 'localtime'))`
	_, err = db.Db.NamedExec(createSQL, databases)
	if err != nil {
		lg.SqlLogDebug(createSQL, databases)
	}
	return err
}

// Update
//
//	@Description: 更新数据库资产
//	@receiver databases
//	@return error
func (databases AssetsDatabasesModel) Update(assetName, databasesAddress, databasesUserName, databasesPassword, database string, databasesPort int) error {
	// 查重
	_, err := uniqueQuery(databasesAddress, databasesUserName, assetName, databasesPort, databases.Id)
	if db.NoRows(err) == false {
		return tool.StringToErr("重复添加")
	} else if err != nil && db.NoRows(err) == false {
		return err
	}

	// 更新
	updateSQL := `update assetsDatabases set  assetName = ? , databasesAddress = ? , databasesUserName = ?, databasesPassword = ? ,database = ? ,databasesPort = ? where id = ?;`
	_, err = db.Db.Exec(updateSQL, assetName, databasesAddress, databasesUserName, databasesPassword, database, databasesPort, databases.Id)
	if err != nil {
		lg.SqlLogDebug(updateSQL, assetName, databasesAddress, databasesUserName, databasesPassword, database, databasesPort, databases.Id)
		lg.Error(err)
	}
	return err
}

// Basics
//
//	@Description: 连接资产基础
//	@receiver databases
//	@return dbConnection.BasicsDatabases
//	@return error
func (databases AssetsDatabasesModel) Basics() (dbConnection.DatabaseInterface, error) {
	basis := dbConnection.BasicsDatabases{
		UserName:    databases.DatabasesUserName,
		Password:    databases.DatabasesPassword,
		HostAddress: databases.DatabasesAddress,
		Port:        databases.DatabasesPort,
		Databases:   databases.Database,
		Timeout:     30,
	}

	var databasesOpen dbConnection.DatabaseInterface
	switch databases.DatabasesType {
	case "mysql":
		databasesOpen = &dbConnection.MysqlOpenStruct{BasicsDatabases: basis}
	default:
		return nil, tool.StringToErr("不支持数据库类型")
	}

	return databasesOpen, nil
}
