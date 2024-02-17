package dataReplicationModel

// DataReplicationModel
// @Description: 数据复制表
type DataReplicationModel struct {
	Id               int64  `json:"id" db:"id" sqlLite:"type=INTEGER"`
	CreateTime       string `json:"createTime" db:"createTime" sqlLite:"type=DATA,default='2006-01-02 15:04:05'"`        // 创建时间
	UpdateTime       string `json:"updateTime" db:"updateTime" sqlLite:"type=DATA,default='2006-01-02 15:04:05'"`        // 更新时间
	TaskName         string `json:"taskName" db:"taskName" sqlLite:"type=TEXT"  binding:"required"`                      // 任务名称
	SourceDataSource int64  `json:"sourceDataSource" db:"sourceDataSource" sqlLite:"type=INTEGER" binding:"required"`    // 源数据源
	TargetDataSource int64  `json:"targetDataSource" db:"targetDataSource" sqlLite:"type=INTEGER" binding:"required"`    // 目标数据源
	CopyMethod       int    `json:"copyMethod" db:"copyMethod" sqlLite:"type=INTEGER" binding:"required"`                // 复制方式
	CopyType         string `json:"copyType" db:"copyType" sqlLite:"type=TEXT" binding:"required"`                       // 复制类型
	SourceObject     string `json:"sourceObject" db:"sourceObject" sqlLite:"type=TEXT,default='[]'"  binding:"required"` // 源对象，可能是库或者表
	TargetObject     string `json:"targetObject" db:"targetObject" sqlLite:"type=TEXT,default='[]'"  binding:"required"` // 目标对象，可能是库或者表
	EditingSteps     int    `json:"editingSteps" db:"editingSteps" sqlLite:"type=INTEGER,default=1"  binding:"required"` // 编辑步骤
}
