package models

import (
	"VizMigrateX/internal/pkg/lg"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// LabelListType 标签类型
// 在gorm中不支持使用[]string,需要自定义类型,并实现Scan，Value方法
type LabelListType []string

type DataSource struct {
	BasicModel

	DataSourceName       string        `gorm:"unique" json:"dataSourceName"`                                         // 数据源名称
	ConnectionAddress    string        `json:"connectionAddress"`                                                    // 连接地址
	ConnectionPort       int           `json:"connectionPort"`                                                       // 连接端口
	DatabaseAccount      string        `json:"databaseAccount"`                                                      // 数据库账号
	DatabasePassword     string        `json:"databasePassword"`                                                     // 数据库密码
	DatabaseType         string        `json:"databaseType"`                                                         // 数据库类型
	Environment          Environment   `json:"environment" gorm:"ForeignKey:EnvironmentID;AssociationForeignKey:id"` // 环境外键声明
	EnvironmentID        uint          `json:"environmentID"`                                                        // 环境ID
	Label                LabelListType `json:"label" gorm:"type:string"`                                             // 标签
	Creator              User          `json:"creator" gorm:"ForeignKey:CreatorID;AssociationForeignKey:id"`         // 创建人外键声明
	CreatorID            uint          `json:"creatorID"`                                                            // 创建人id
	Explain              string        `json:"explain"`                                                              // 说明
	AdditionalParameters string        `json:"additionalParameters"`                                                 // 额外参数
	State                int           `json:"state" gorm:"default:0"`                                               // 0 正常,1 异常
	ConnectionLog        string        `json:"connectionLog" gorm:"default:''"`                                      // 连接日志
}

func (j *LabelListType) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var s []string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	*j = LabelListType(s)
	return nil
}

func (j LabelListType) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// ChangeState
//
//	@Description: 改变状态
//	@receiver thisDataSource
//	@param state
func (thisDataSource DataSource) ChangeState(state int, connectionLog string) {
	if err := DB.Model(&DataSource{}).Where("id = ?", thisDataSource.ID).Update("state", state).Update("connection_log", connectionLog).Error; err != nil {
		lg.Logger.Errorln(err.Error())
	}
}
