package models

import (
	"VizMigrateX/internal/pkg/lg"
	"errors"
	"fmt"
	GormLogger "gorm.io/gorm/logger"
)

// Environment
// @Description: 环境表
type Environment struct {
	BasicModel

	Name      string `json:"name" gorm:"unique"`         // 环境名称
	CreatorID uint   `json:"creatorID" gorm:"default:0"` // 创建人ID
}

// Label
// @Description: 标签表
type Label struct {
	BasicModel

	Name      string `json:"name" gorm:"unique"`         // 标签名称
	CreatorID uint   `json:"creatorID" gorm:"default:0"` // 创建人ID
}

// initializeTableDataEnvironment
//
//	@Description: 初始化环境表
func initializeTableDataEnvironment() {
	for _, s := range []string{"预发", "生产"} {
		var existingEnv Environment
		result := DB.Where("name = ?", s).First(&existingEnv)
		if result.Error != nil && !errors.Is(GormLogger.ErrRecordNotFound, result.Error) {
			// 处理查询错误
			lg.Logger.Errorln(result.Error.Error())
			panic(fmt.Sprintf("初始化环境表失败 err: %s", result.Error.Error()))
		} else {
			if result.RowsAffected == 0 {
				// 记录不存在，执行插入操作
				newEnv := Environment{Name: s}
				if DB.Create(&newEnv).Error != nil {
					panic(fmt.Sprintf("初始化环境表失败 err: %s", result.Error.Error()))
				}
			}
		}
	}

}
