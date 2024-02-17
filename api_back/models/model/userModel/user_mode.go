package userModel

// UsersModelDb
//
//	@Description: 用户表
type UsersModelDb struct {
	Id                     int64  `json:"id" db:"id"`
	UserName               string `json:"userName" db:"userName" sqlLite:"type=text"`                                          // 用户名
	DisplayName            string `json:"displayName" db:"displayName" sqlLite:"type=text"`                                    // 显示名
	Password               string `json:"password" db:"password" sqlLite:"type=varchar(24)"`                                   // 密码
	Role                   int    `json:"role" db:"role" sqlLite:"type=int"`                                                   // 角色
	Phone                  string `json:"phone" db:"phone" sqlLite:"type=varchar(12),default=''"`                              // 手机号
	MoreContactInformation string `json:"moreContactInformation" db:"moreContactInformation" sqlLite:"type=text,default='{}'"` // 其他联系方式
	CreateTime             string `json:"createTime" db:"createTime" sqlLite:"type=datetime"`                                  //  创建时间
}
