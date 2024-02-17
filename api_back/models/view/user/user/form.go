package user

import (
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/myValidate"
	"dataxWeb/models/util/tool/db"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("isUserNameValidator", isUserNameValidator)
		_ = v.RegisterValidation("isDisplayNameValidator", isDisplayNameValidator)
	}
	myValidate.CustomizeError{Name: "userInsertJsonStruct.userName", Tag: "isUserNameValidator", ErrorStr: "用户名重复"}.Add()
	myValidate.CustomizeError{Name: "userInsertJsonStruct.displayName", Tag: "isDisplayNameValidator", ErrorStr: "显示名重复"}.Add()
}

// 查看用户名是否重复
var isUserNameValidator validator.Func = func(fl validator.FieldLevel) bool {
	if _, err := usersWhereUserName(fl.Field().String()); err != nil {
		if db.NoRows(err) {
			return true
		}
		lg.Error(err)
	}
	return false
}

// 查看显示名是否重复
var isDisplayNameValidator validator.Func = func(fl validator.FieldLevel) bool {
	if _, err := usersWhereDisplayName(fl.Field().String()); err != nil {
		if db.NoRows(err) {
			return true
		}
		lg.Error(err)
	}
	return false
}
