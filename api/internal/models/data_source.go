package models

type DataSource struct {
	BasicModel

	DataSourceName       string      `gorm:"unique" json:"dataSourceName"`                                         // 数据源名称
	ConnectionAddress    string      `json:"connectionAddress"`                                                    // 连接地址
	ConnectionPort       int         `json:"connectionPort"`                                                       // 连接端口
	DatabaseAccount      string      `json:"databaseAccount"`                                                      // 数据库账号
	DatabasePassword     string      `json:"databasePassword"`                                                     // 数据库密码
	DatabaseType         string      `json:"databaseType"`                                                         // 数据库类型
	Environment          Environment `json:"environment" gorm:"ForeignKey:EnvironmentID;AssociationForeignKey:id"` // 环境外键声明
	EnvironmentID        uint        `json:"environmentID"`                                                        // 环境ID
	Label                []string    `json:"label"  gorm:"type:bytes;serializer:json"`                             // 标签
	Creator              User        `json:"creator" gorm:"ForeignKey:CreatorID;AssociationForeignKey:id"`         // 创建人外键声明
	CreatorID            uint        `json:"creatorID"`                                                            // 创建人id
	Explain              string      `json:"explain"`                                                              // 说明
	AdditionalParameters string      `json:"additionalParameters"`                                                 // 额外参数
}
