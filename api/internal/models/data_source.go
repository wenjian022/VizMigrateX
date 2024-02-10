package models

type DataSource struct {
	BasicModel

	DataSourceName       string `gorm:"unique"` // 数据源名称
	ConnectionAddress    string // 连接地址
	ConnectionPort       int    // 连接端口
	DatabaseAccount      string // 数据库账号
	DatabasePassword     string // 数据库密码
	Environment          string // 环境
	Explain              string // 说明
	AdditionalParameters string // 额外参数
}
