package user

// userLoginJsonStrict 用户登录需要的请求体
type userLoginJsonStrict struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// userInsertStruct
// @Description: 添加用户结构体
type userInsertJsonStruct struct {
	UserName    string `json:"userName" binding:"required,isUserNameValidator"`
	DisplayName string `json:"displayName" binding:"required,isDisplayNameValidator"`
	Password    string `json:"password" binding:"required"`
	Role        int    `json:"role" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
}
