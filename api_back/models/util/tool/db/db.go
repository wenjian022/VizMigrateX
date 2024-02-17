package db

import (
	"database/sql"
	"dataxWeb/models/util/lg"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
	"os"
	"path"
	"regexp"
)

var Db *sqlx.DB
var ErrNoRowsZh = errors.New("未查询到数据")

func init() {
	Db = openDb()
}

// openDb
//
//	@Description: 创建一个sqlite3连接
//	@return *sqlx.DB
func openDb() *sqlx.DB {
	_ = os.MkdirAll("db", 0666)
	db, err := sqlx.Open("sqlite", path.Join("db", "fcc.db"))
	// 只使用一个链接数，解决 database is locked 问题
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	if err != nil {
		panic(err)
	}
	return db
}

// NoRows
//
//	@Description: 判断查询结果是否无数据
//	@param err
//	@return bool
func NoRows(err error) bool {
	return errors.Is(sql.ErrNoRows, err)
}

// Registration
//
//	@Description: 注册表
func Registration(tabName string, structure interface{}) {
	formRegistration = append(formRegistration, tableStructure{tableName: tabName, structure: structure})
}

// RunInTransaction
//
//	@Description: 事务方法
//	@param db db实例
//	@param transactionName 事务名称
//	@param f
//	@return error
func RunInTransaction(db *sqlx.DB, transactionName string, f func(tx *sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			lg.ErrorStr("事务回滚", transactionName)
		} else if err != nil {
			_ = tx.Rollback()
			lg.ErrorStr("事务回滚", transactionName)
		} else {
			_ = tx.Commit()
			lg.Info(fmt.Sprintf("[%s] 事务提交", transactionName))
		}
	}()
	err = f(tx)
	return err
}

// IsSelect
//
//	@Description: 判断语句是否是select
//	@param sqlStatement
//	@return bool
func IsSelect(sqlStatement string) bool {
	selectPattern := `(?i)^\s*SELECT\s+`
	matched, err := regexp.MatchString(selectPattern, sqlStatement)
	if err != nil {
		lg.Error(err)
		return false
	}
	return matched
}
