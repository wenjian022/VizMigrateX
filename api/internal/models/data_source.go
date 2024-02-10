package models

type DataSource struct {
	BasicModel

	DataSourceName    string `gorm:"unique" json:"dataSourceName"` // 数据源名称
	ConnectionAddress string `json:"connectionAddress"`            // 连接地址
	ConnectionPort    int    `json:"connectionPort"`               // 连接端口
	DatabaseAccount   string `json:"databaseAccount"`              // 数据库账号
	DatabasePassword  string `json:"databasePassword"`             // 数据库密码
	DatabaseType      string `json:"databaseType"`                 // 数据库类型
	//Environment          string // 环境
	Explain              string `json:"explain"`              // 说明
	AdditionalParameters string `json:"additionalParameters"` // 额外参数
}
