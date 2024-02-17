package dbConnection

import (
	"dataxWeb/models/util/lg"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
	"time"
)

type DatabaseInterface interface {
	// Open 建立连接
	Open() error
	// Close 关闭连接
	Close()
	// Select 执行查询
	Select(dest interface{}, query string, args ...interface{}) error
}

type BasicsDatabases struct {
	UserName    string        // 用户
	Password    string        // 密码
	HostAddress string        // 链接地址
	Port        int           // 端口
	Databases   string        // 数据库名
	Timeout     time.Duration // 超时时间
	Db          *sqlx.DB
}

// MysqlOpenStruct MySQL库链接方式
type MysqlOpenStruct struct {
	BasicsDatabases
}

func (mySQL *MysqlOpenStruct) Open() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)%s", mySQL.UserName, mySQL.Password, mySQL.HostAddress, mySQL.Port, mySQL.Databases)
	if open, err := sqlx.Open("mysql", dsn); err != nil {
		lg.ErrorStr(dsn, "连接[mysql]数据库失败 DSN:")
		return err
	} else {
		mySQL.Db = open
	}
	return mySQL.Db.Ping()
}

func (mySQL *MysqlOpenStruct) Close() {
	if mySQL.Db != nil {
		_ = mySQL.Db.Close()
	}
}

func (mySQL *MysqlOpenStruct) Select(dest interface{}, query string, args ...interface{}) error {
	return mySQL.Db.Select(dest, query, args...)
}
