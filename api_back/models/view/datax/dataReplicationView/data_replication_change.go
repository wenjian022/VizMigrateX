package dataReplicationView

import (
	"dataxWeb/models/model/dataReplicationModel"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool/db"
	"encoding/json"
)

// create
//
//	@Description:
//	@receiver thisStep1
//	@return int64
//	@return error
func (thisStep1 dataReplicationStep1JsonStruct) create() (int64, error) {
	// 唯一校验
	if err := dataReplicationModel.UniqueQuery(thisStep1.TaskName, 0); err != nil {
		lg.Error(err)
		return 0, err
	}
	CopyTypeBy, err := json.Marshal(thisStep1.CopyType)
	if err != nil {
		lg.Error(err)
		return 0, err
	}

	// 创建
	createSQL := `
		insert into dataReplication(createTime, updateTime, taskName, sourceDataSource, targetDataSource, copyMethod, copyType)
		values (datetime('now', 'localtime'), datetime('now', 'localtime'), ?,?,?,?,?,?)
	`
	row, err := db.Db.Exec(createSQL, thisStep1.TaskName, thisStep1.SourceDataSource, thisStep1.TargetDataSource,
		thisStep1.CopyMethod, CopyTypeBy)
	if err != nil {
		lg.Error(err)
		return 0, err
	}

	// 获取添加ID
	rowId, err := row.LastInsertId()
	if err != nil {
		lg.Error(err)
		return 0, err
	}

	return rowId, nil
}
