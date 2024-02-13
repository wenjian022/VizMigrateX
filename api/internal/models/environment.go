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
	Color     string `json:"color"`                      // 颜色
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
	initEnvMap := []map[string]string{{"name": "生产", "color": "#F56C6C"}, {"name": "预发", "color": "#E6A23C"}}
	for _, s := range initEnvMap {
		var existingEnv Environment
		result := DB.Where("name = ?", s["name"]).First(&existingEnv)
		if result.Error != nil && !errors.Is(GormLogger.ErrRecordNotFound, result.Error) {
			// 处理查询错误
			lg.Logger.Errorln(result.Error.Error())
			panic(fmt.Sprintf("初始化环境表失败 err: %s", result.Error.Error()))
		} else {
			if result.RowsAffected == 0 {
				// 记录不存在，执行插入操作
				newEnv := Environment{Name: s["name"], Color: s["color"]}
				if DB.Create(&newEnv).Error != nil {
					panic(fmt.Sprintf("初始化环境表失败 err: %s", result.Error.Error()))
				}
			}
		}
	}

}
