package assetDatabasesModel

import (
	"dataxWeb/models/util/myValidate"
	"dataxWeb/models/util/tool"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("databaseTypeValidator", databaseTypeValidator)
		_ = v.RegisterValidation("defaultDatabaseValidator", defaultDatabaseValidator)
	}
	myValidate.CustomizeError{Name: "AssetsDatabasesModel.Databases", Tag: "defaultDatabaseValidator", ErrorStr: "默认库不能为空"}.Add()
	myValidate.CustomizeError{Name: "AssetsDatabasesModel.DatabasesType", Tag: "databaseTypeValidator", ErrorStr: "类型错误"}.Add()
}

var databasesTypeList = []string{"mysql"}

var databaseTypeValidator validator.Func = func(f validator.FieldLevel) bool {
	return tool.IsElementStr(databasesTypeList, f.Field().String())
}

var defaultDatabaseValidator validator.Func = func(fl validator.FieldLevel) bool {
	var field = fl.Field().String()
	if field == "" {
		return false
	}
	return true
}
