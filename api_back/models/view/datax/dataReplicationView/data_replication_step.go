package dataReplicationView

import (
	"dataxWeb/models/model/dataReplicationModel"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool/db"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func (thisStep2 dataReplicationStep2JsonStruct) update(c *gin.Context, taskInfo dataReplicationModel.DataReplicationModel) error {
	// 数据校验
	if err := c.ShouldBindJSON(&thisStep2); err != nil {
		lg.Error(err)
		return err
	}

	sourceObjectBy, err := json.Marshal(thisStep2)
	if err != nil {
		lg.Error(err)
		return err
	}

	if taskInfo.EditingSteps < 2 {
		taskInfo.EditingSteps = 2
	}

	// 更新数据
	putSQL := `update dataReplication set sourceObject = ?,editingSteps = ? where id = ?;`
	_, err = db.Db.Exec(putSQL, string(sourceObjectBy), taskInfo.EditingSteps, taskInfo.Id)
	if err != nil {
		lg.Error(err)
		return err
	}

	return nil
}
