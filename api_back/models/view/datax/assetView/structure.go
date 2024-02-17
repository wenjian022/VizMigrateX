package assetView

import "dataxWeb/models/util/tool"

// assetsQueryStruct
// @Description: 资产查询结构体
type assetsQueryStruct struct {
	tool.PageQueryStruct
	DatabasesType    string `form:"databasesType"`
	AssetName        string `form:"assetName"`
	DatabasesAddress string `form:"databasesAddress"`
}

type assetsIdUriStruct struct {
	AssetsId int64 `uri:"assetsID" binding:"required"` // 资产id
}

// assetsDatabasesSelectJsonStruct
// @Description: 资产数据库-单选查询
type assetsDatabasesSelectJsonStruct struct {
	Value        int64  `json:"value"`
	Label        string `json:"label"`
	DatabaseType string `json:"databaseType"`
}

type libraryListJsonStruct struct {
	Id         string                  `json:"id"`
	Type       int                     `json:"type"` // 类型， 0 库、1 表
	Label      string                  `json:"label"`
	SchemaName string                  `json:"schemaName,omitempty"` // 库名
	Children   []libraryListJsonStruct `json:"children,omitempty"`
}
