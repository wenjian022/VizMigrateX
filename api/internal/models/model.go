package models

import (
	"VizMigrateX/internal/pkg/lg"
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	GormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"time"
)

type LocalTime struct {
	time.Time
}

type BasicModel struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt *LocalTime `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *LocalTime `json:"updated_at" gorm:"autoUpdateTime"`
}

var DB *gorm.DB

func Connect(dataSourceType, address string) {

	logger := lg.Logger

	var db *gorm.DB
	var err error

	switch dataSourceType {
	case "sqlite":
		// 创建目录
		_ = os.MkdirAll(path.Dir(address), 0666)
		db, err = gorm.Open(sqlite.Open(address), &gorm.Config{
			Logger: GormLogger.New(log.New(logger.Out, "\r\n", log.LstdFlags),
				GormLogger.Config{
					SlowThreshold: time.Second,
					LogLevel:      GormLogger.Info,
					Colorful:      false,
				},
			),
		})
		sqlDB, err := db.DB()
		if err != nil {
			panic(fmt.Sprintf("数据库连接失败 %s", err))
		}
		// 配置最大连接数
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetMaxIdleConns(10)
	case "":
		logger.Debugln("没有配置数据库信息,跳过数据库初始化")
		return
	default:
		panic(fmt.Sprintf("暂不支持该数据库类型 %s", dataSourceType))
	}
	// Migrate the schema
	migrateErr := db.AutoMigrate(&User{}, &DataSource{}, &Environment{}, &Label{}, &BackupHost{})

	if migrateErr != nil {
		panic(fmt.Sprintf("自动迁移失败，请使用检查您的 %s 地址: %s", dataSourceType, address))
	}

	DB = db

	logger.Printf(fmt.Sprintf("数据库连接成功 %s 地址: %s", dataSourceType, address))
	if err != nil {
		panic(fmt.Sprintf("连接失败，请使用检查您的:%s 地址: %s err: %s", dataSourceType, address, err.Error()))
	}

	// 初始化表数据
	initializeTableDataEnvironment()

}

func (lt *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := lt.Time
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Scan
//
//	@Description: 实现GORM的Scanner接口
//	@receiver lt
//	@param value
//	@return error
func (lt *LocalTime) Scan(value interface{}) error {
	if value == nil {
		lt.Time = time.Time{}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("failed to scan LocalTime field - source is not a time.Time")
	}
	lt.Time = t
	return nil
}

// Value
//
//	@Description: Valuer接口
//	@receiver lt
//	@return driver.Value
//	@return error
func (lt *LocalTime) Value() (driver.Value, error) {
	if lt.IsZero() {
		return nil, nil
	}
	return lt.Time, nil
}
