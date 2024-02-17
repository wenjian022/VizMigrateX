package dataReplicationModel

import (
	"dataxWeb/models/model/assetDatabasesModel"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool"
	"dataxWeb/models/util/tool/db"
)

// UniqueQuery
//
//	@Description: 唯一校验
//	@param taskName
//	@param id
//	@return error
func UniqueQuery(taskName string, id int64) error {
	querySQL := `select id from dataReplication where taskName = ? and taskName != ?;`
	err := db.Db.Get(&DataReplicationModel{}, querySQL, taskName, id)
	if err != nil {
		if db.NoRows(err) {
			return nil
		}
		return err
	}
	return tool.StringToErr("任务名称不能重复")
}

// AssetInfo
//
//	@Description:
//	@receiver thisReplication
//	@return assetDatabasesModel.AssetsDatabasesModel 源数据源
//	@return assetDatabasesModel.AssetsDatabasesModel 目标数据源
//	@return error
func (thisReplication DataReplicationModel) AssetInfo() (assetDatabasesModel.AssetsDatabasesModel, assetDatabasesModel.AssetsDatabasesModel, error) {
	var sAsset, tAsset assetDatabasesModel.AssetsDatabasesModel
	// 源数据源
	sAsset, err := assetDatabasesModel.WhereIdQuery(thisReplication.SourceDataSource)
	if err != nil {
		lg.Error(err)
		return sAsset, tAsset, err
	}

	// 目标数据源
	tAsset, err = assetDatabasesModel.WhereIdQuery(thisReplication.TargetDataSource)
	if err != nil {
		lg.Error(err)
		return sAsset, tAsset, err
	}
	return sAsset, tAsset, nil
}

// WhereIdQuery
//
//	@Description: 通过id查询
//	@param id
//	@return DataReplicationModel
//	@return error
func WhereIdQuery(id int64) (DataReplicationModel, error) {
	querySQL := `select * from dataReplication where id = ?`
	var dataReplication DataReplicationModel
	err := db.Db.Get(&dataReplication, querySQL, id)
	return dataReplication, err
}
