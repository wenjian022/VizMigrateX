package models

// BackupHost
// @Description: 备份主机
type BackupHost struct {
	BasicModel

	Creator           User   `json:"creator" gorm:"ForeignKey:CreatorID;AssociationForeignKey:id"` // 创建人外键声明
	CreatorID         uint   `json:"creatorID"`                                                    // 创建人id
	HostName          string `gorm:"unique"`                                                       // 主机名称
	ConnectionAddress string // 连接地址
	ConnectionPort    int    // 连接端口
	LineType          string // 服务器类型，如: FTP,SFTP,SSH
	HostAccount       string // 账号
	HostPassword      string // 密码
	BackupDirectory   string // 备份目录
}
