package dataReplicationView

import (
	"dataxWeb/models/model/dataReplicationModel"
	"dataxWeb/models/util/tool"
	"github.com/gin-gonic/gin"
)

type dataReplicationTaskQueryStruct struct {
	tool.PageQueryStruct
	TaskName string `form:"taskName"`
}
type dataReplicationStepInterface interface {
	update(*gin.Context, dataReplicationModel.DataReplicationModel) error
}

// dataReplicationStep1JsonStruct
// @Description: 添加数据复制任务流程一
type dataReplicationStep1JsonStruct struct {
	TaskName         string `json:"taskName"  binding:"required" db:"taskName"`                // 任务名称
	SourceDataSource int64  `json:"sourceDataSource" binding:"required" db:"sourceDataSource"` // 源数据源
	TargetDataSource int64  `json:"targetDataSource" binding:"required" db:"targetDataSource"` // 目标数据源
	CopyMethod       int    `json:"copyMethod" db:"copyMethod"`                                // 复制方式
	CopyType         []int  `json:"copyType" binding:"required" db:"copyType"`                 // 复制类型
}

// dataReplicationStep2JsonStruct
// @Description: 添加数据复制任务流程二
type dataReplicationStep2JsonStruct struct {
	SourceObject []struct {
		SchemaName string `json:"schemaName" binding:"required"`
		TableName  string `json:"tableName" binding:"required"`
	} `json:"sourceObject" binding:"required"`
}

// dataReplicationTaskInfoResStruct
// @Description: 数据复制任务列表返回结构体
type dataReplicationTaskInfoResStruct struct {
	Id               int64               `json:"id"`
	TaskName         string              `json:"taskName" `         // 任务名称
	SourceDataSource sourceDetailsStruct `json:"sourceDataSource" ` // 源数据源
	TargetDataSource sourceDetailsStruct `json:"targetDataSource"`  // 目标数据源
	CopyMethod       int                 `json:"copyMethod" `       // 复制方式
	EditingSteps     int                 `json:"editingSteps"`      // 编辑步骤
}

// sourceDetailsStruct
// @Description: 数据复制任务源详情信息
type sourceDetailsStruct struct {
	AssetName        string `json:"assetName" `       // 资产名称
	DatabasesType    string `json:"databasesType" `   // 数据库类型
	DatabasesAddress string `json:"databasesAddress"` // 资产类型
	DatabasesPort    int    `json:"databasesPort"`    // 资产端口
}

type dataReplicationTaskIDUriStruct struct {
	TaskID int64 `uri:"taskID"` //
}

type dataReplicationStepTaskIDUriStruct struct {
	dataReplicationTaskIDUriStruct
	Step int `uri:"step"` // 步骤
}
