package assetDatabasesModel

// AssetsDatabasesModel
// @Description: 数据库资产表
type AssetsDatabasesModel struct {
	Id                int64  `json:"id" db:"id" sqlLite:"type=INTEGER"`
	UpdateTime        string `json:"updateTime" db:"updateTime" sqlLite:"type=DATA,default='2006-01-02 15:04:05'"`                             // 更新时间
	AssetName         string `json:"assetName" db:"assetName" sqlLite:"type=TEXT"  binding:"required"`                                         // 资产名称
	DatabasesType     string `json:"databasesType" db:"databasesType" sqlLite:"type=TEXT,default=''" binding:"required,databaseTypeValidator"` // 数据库类型
	DatabasesAddress  string `json:"databasesAddress" db:"databasesAddress" sqlLite:"type=TEXT" binding:"required"`                            // 资产类型
	DatabasesPort     int    `json:"databasesPort" db:"databasesPort" sqlLite:"type=INTEGER" binding:"required"`                               // 资产端口
	DatabasesUserName string `json:"databasesUserName" db:"databasesUserName" sqlLite:"type=TEXT,default=''" binding:"required"`               // 数据库登录用户
	DatabasesPassword string `json:"databasesPassword" db:"databasesPassword" sqlLite:"type=TEXT,default=''" binding:"required"`               // 数据库密码
	Database          string `json:"database" db:"database" sqlLite:"type=TEXT" binding:"defaultDatabaseValidator"`                            // 默认连接库
}
