package user

import (
	model "dataxWeb/models/model/userModel"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool/db"
)

// userInsert
//
//	@Description: 添加用户
//	@param userJson
//	@return error
func userInsert(userJson userInsertJsonStruct) error {
	sqlStr := `insert into users(userName, displayName, password, role, phone, createTime) values (?,?,?,?,?,datetime('now', 'localtime'))`
	lg.SqlLogDebug(sqlStr, userJson.UserName, userJson.DisplayName, userJson.Password, userJson.Role)
	_, err := db.Db.Exec(sqlStr, userJson.UserName, userJson.DisplayName, userJson.Password, userJson.Role, userJson.Phone)
	return err
}

// usersWhereUserName
//
//	@Description: 通过用户名查看用户
//	@param userName
//	@return rel
//	@return err
func usersWhereUserName(userName string) (rel model.UsersModelDb, err error) {
	sqlStr := `select * from users where userName = ?`
	err = db.Db.Get(&rel, sqlStr, userName)
	return
}

// usersWhereDisplayName
//
//	@Description: 通过显示名查看用户
//	@param displayName
//	@return rel
//	@return err
func usersWhereDisplayName(displayName string) (rel model.UsersModelDb, err error) {
	sqlStr := `select * from users where  displayName = ?`
	err = db.Db.Get(&rel, sqlStr, displayName)
	return
}

// userWhereUserAndPwd
//
//	@Description: 查看用户名或密码是否正常
//	@param userLogin
//	@return rel
//	@return err
func userWhereUserAndPwd(userLogin userLoginJsonStrict) (rel model.UsersModelDb, err error) {
	sqlStr := `select * from  users where  userName = ? and password = ?`
	err = db.Db.Get(&rel, sqlStr, userLogin.UserName, userLogin.Password)
	return
}
