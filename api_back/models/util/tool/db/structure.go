package db

import "database/sql"

//
//  tableInfoStruct
//  @Description: 表详情
//
type tableInfoStruct struct {
	Cid       int64          `db:"cid"`
	Name      string         `db:"name"`       // 字段名称
	Type      string         `db:"type"`       // 字段类型
	NotNull   bool           `db:"notnull"`    // 是否为空
	DfltValue sql.NullString `db:"dflt_value"` // 默认字段
	Pk        int            `db:"pk"`         // 主键
}

type tableStructure struct {
	tableName string      // 表名
	structure interface{} // 表结构
}

//  fieldAttributeStruct 字段属性
type fieldAttributeStruct struct {
	Type    string // 字段类型
	Default string // 默认值
}
