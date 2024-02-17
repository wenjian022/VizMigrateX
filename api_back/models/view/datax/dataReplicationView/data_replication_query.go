package dataReplicationView

import (
	"dataxWeb/models/model"
	"dataxWeb/models/model/dataReplicationModel"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool"
	"dataxWeb/models/util/tool/db"
)

func (receiver dataReplicationTaskQueryStruct) query() (rel tool.PagingDataResStruct, err error) {
	querySQL := `SELECT * FROM dataReplication WHERE TRUE`
	totalSQL := `SELECT COUNT(*) AS total FROM dataReplication WHERE TRUE`
	var args []interface{}
	var WHERE string
	// 条件查询
	if receiver.TaskName != "" {
		WHERE = WHERE + " and taskName = ? "
		args = append(args, receiver.TaskName)
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

	var taskList []dataReplicationModel.DataReplicationModel
	err = db.Db.Select(&taskList, querySQL, args...)
	if err != nil {
		lg.SqlLogDebug(querySQL, args...)
		lg.Error(err, "获取资产查询失败")
	}

	// 数据处理
	var dataReplicationTaskInfoRes []dataReplicationTaskInfoResStruct
	for _, replication := range taskList {
		sAsset, tAsset, _err := replication.AssetInfo()
		if err != nil {
			lg.Error(_err)
			return rel, _err
		}

		dataReplicationTaskInfoRes = append(dataReplicationTaskInfoRes, dataReplicationTaskInfoResStruct{
			Id:       replication.Id,
			TaskName: replication.TaskName,
			SourceDataSource: sourceDetailsStruct{
				AssetName:        sAsset.AssetName,
				DatabasesType:    sAsset.DatabasesType,
				DatabasesAddress: sAsset.DatabasesAddress,
				DatabasesPort:    sAsset.DatabasesPort,
			},
			TargetDataSource: sourceDetailsStruct{
				AssetName:        tAsset.AssetName,
				DatabasesType:    tAsset.DatabasesType,
				DatabasesAddress: tAsset.DatabasesAddress,
				DatabasesPort:    tAsset.DatabasesPort,
			},
			CopyMethod:   replication.CopyMethod,
			EditingSteps: replication.EditingSteps,
		})

	}

	rel = tool.PagingDataResStruct{
		Data:  dataReplicationTaskInfoRes,
		Total: total,
		Size:  receiver.Size,
		Page:  receiver.Page,
	}
	return
}

func (thisTaskID dataReplicationTaskIDUriStruct) query() (dataReplicationModel.DataReplicationModel, error) {
	return dataReplicationModel.WhereIdQuery(thisTaskID.TaskID)
}
