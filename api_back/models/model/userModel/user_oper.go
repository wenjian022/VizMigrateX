package userModel

import "dataxWeb/models/util/tool/db"

// UsersAll
//
//	@Description: 查看所有用户
func UsersAll() (rel []UsersModelDb, err error) {
	sqlStr := `select * from users`
	err = db.Db.Select(&rel, sqlStr)
	return
}

// userInit
//
//	@Description: 用户信息初始化

// UserInit
//
//	@Description:用户信息初始化
func UserInit() {
	sqlStr := `select * from users where id = 1`
	var userInfo UsersModelDb
	if err := db.Db.Get(&userInfo, sqlStr); err != nil {
		if db.NoRows(err) {
			// 生成admin用户
			sqlStr = `insert into users (id,userName, displayName, password, role, phone, createTime)  values (1,'admin','管理员','admin1234',0,'15928651434',datetime('now','localtime'))`
			if _, err := db.Db.Exec(sqlStr); err != nil {
				panic(err)
			}
		}
	}
}
