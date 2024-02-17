package model

import (
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool/db"
)

// SqlTotal
//
//	@Description: 获取查询条件的统计
//	@param sqlTotal
//	@param args
//	@return int
//	@return error
func SqlTotal(selectSql string, args ...interface{}) (int, error) {
	// 统计数量
	var total struct {
		Total int `db:"total"`
	}
	err := db.Db.Get(&total, selectSql, args...)
	if err != nil {
		lg.SqlLogDebug(selectSql, args)
	}
	return total.Total, err
}
