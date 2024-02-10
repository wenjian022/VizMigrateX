package dataSource

import (
	"VizMigrateX/internal/pkg/lg"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
	"time"
)

var NotSupportedErr = errors.New("不支持这个数据源")

type DatabaseInterface interface {
	// Open 建立连接
	Open() error
	// Close 关闭连接
	Close()
	// Select 执行查询
	Select(dest interface{}, query string, args ...interface{}) error
}

// BasicsDatabases
// @Description: 基础结构体
type BasicsDatabases struct {
	ConnectionAddress    string // 连接地址
	ConnectionPort       int    // 连接端口
	DatabaseAccount      string // 数据库账号
	DatabasePassword     string // 数据库密码
	AdditionalParameters string // 额外参数

	Timeout time.Duration // 超时时间

	Db *sqlx.DB
}

// MysqlStruct Mysql 连接
type MysqlStruct struct {
	BasicsDatabases
}

func (thisMysql *MysqlStruct) Open() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/?%s",
		thisMysql.DatabaseAccount, thisMysql.DatabasePassword,
		thisMysql.ConnectionAddress, thisMysql.ConnectionPort,
		thisMysql.AdditionalParameters)

	// 连接数据库
	if open, err := sqlx.Open("mysql", dsn); err != nil {
		lg.Logger.Errorln(err.Error())
		return err
	} else {
		thisMysql.Db = open
	}
	return thisMysql.Db.Ping()
}

func (thisMysql *MysqlStruct) Close() {
	if thisMysql.Db != nil {
		_ = thisMysql.Db.Close()
	}
}

func (thisMysql *MysqlStruct) Select(dest interface{}, query string, args ...interface{}) error {
	return thisMysql.Db.Select(dest, query, args...)
}
