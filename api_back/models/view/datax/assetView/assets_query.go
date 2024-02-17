package assetView

import (
	"dataxWeb/models/model"
	"dataxWeb/models/model/assetDatabasesModel"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool"
	"dataxWeb/models/util/tool/db"
	"fmt"
)

func (receiver assetsQueryStruct) query() (rel tool.PagingDataResStruct, err error) {
	querySQL := `SELECT * FROM assetsDatabases WHERE TRUE`
	totalSQL := `SELECT COUNT(*) AS total FROM assetsDatabases WHERE TRUE`
	var args []interface{}
	var WHERE string
	// 条件查询
	if receiver.DatabasesType != "" && receiver.DatabasesType != "all" {
		WHERE = WHERE + " and databasesType = ? "
		args = append(args, receiver.DatabasesType)
	}
	if receiver.AssetName != "" {
		WHERE = WHERE + " and assetName like ?"
		args = append(args, "%"+receiver.AssetName+"%")
	}
	if receiver.DatabasesAddress != "" {
		WHERE = WHERE + " and databasesAddress like ?"
		args = append(args, "%"+receiver.DatabasesAddress+"%")
	}

	// 查询统计
	total, err := model.SqlTotal(totalSQL+WHERE, args...)
	if err != nil {
		lg.Error(err, "获取资产查询统计失败")
		return
	}

	// 分页查询
	querySQL = querySQL + WHERE + " limit ?,?"
	args = append(args, receiver.Size*(receiver.Page-1), receiver.Size)

	var assets []assetDatabasesModel.AssetsDatabasesModel
	err = db.Db.Select(&assets, querySQL, args...)
	if err != nil {
		lg.SqlLogDebug(querySQL, args...)
		lg.Error(err, "获取资产查询失败")
	}

	rel = tool.PagingDataResStruct{
		Data:  assets,
		Total: total,
		Size:  receiver.Size,
		Page:  receiver.Page,
	}
	return
}

// query
//
//	@Description: 查询
//	@receiver assetsId
//	@return assetDatabasesModel.AssetsDatabasesModel
//	@return error
func (assetsId assetsIdUriStruct) query() (assetDatabasesModel.AssetsDatabasesModel, error) {
	return assetDatabasesModel.WhereIdQuery(assetsId.AssetsId)
}

// assetsSelectQuery
//
//	@Description:
//	@return []assetsDatabasesSelectJsonStruct
//	@return error
func assetsSelectQuery() ([]assetsDatabasesSelectJsonStruct, error) {
	var assetDatabases []assetDatabasesModel.AssetsDatabasesModel
	querySQL := `SELECT id,assetName,databasesAddress,databasesPort,databasesType FROM assetsDatabases `
	if err := db.Db.Select(&assetDatabases, querySQL); err != nil {
		lg.Error(err)
		return nil, err
	}
	// 重定向数据
	var assetsDatabasesSelectJson []assetsDatabasesSelectJsonStruct
	for _, database := range assetDatabases {
		assetsDatabasesSelectJson = append(assetsDatabasesSelectJson, assetsDatabasesSelectJsonStruct{
			Value:        database.Id,
			Label:        fmt.Sprintf("%s  %s:%v", database.AssetName, database.DatabasesAddress, database.DatabasesPort),
			DatabaseType: database.DatabasesType,
		})
	}

	return assetsDatabasesSelectJson, nil
}
