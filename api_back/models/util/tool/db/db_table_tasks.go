package db

import (
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool"
	"fmt"
	"reflect"
	"strings"
)

var formRegistration []tableStructure

// fieldAttribute
//
//	@Description: 字段属性格式化
//	@return attribute
func fieldAttribute(attributeStr []string) (attribute fieldAttributeStruct) {
	attribute.Type = "TEXT" // 默认类型，类型不能为空
	for _, s := range attributeStr {
		if strings.HasPrefix(s, "type=") {
			attribute.Type = strings.Split(s, "=")[1]
		} else if strings.HasPrefix(s, "default=") {
			attribute.Default = strings.Split(s, "=")[1]
		}
	}
	return
}

// createTable
//
//	@Description: 创建表
//	@param tabName
//	@return error
func createTable(tabName string) error {
	//	生成表,默认有一个主键
	createTab := fmt.Sprintf("create table %s (id INTEGER PRIMARY KEY AUTOINCREMENT);", tabName)
	_, err := Db.Exec(createTab)
	return err
}

// dropTable
//
//	@Description: 删除表
//	@param tabName
//	@return error
func dropTable(tabName string) error {
	dropTab := fmt.Sprintf("drop table %s", tabName)
	_, err := Db.Exec(dropTab)
	return err
}

// changeTable
//
//	@Description: 修改表名
//	@param tabName 表名
//	@param changeName 修改后的表名
//	@return error
func changeTable(tabName, changeName string) error {
	changeTab := fmt.Sprintf("alter table %s rename to %s", tabName, changeName)
	_, err := Db.Exec(changeTab)
	return err
}

// isTableField
//
//	@Description: 判断表字段是否存在
func isTableField(tabInfo []tableInfoStruct, field string) (int, bool) {
	for index, info := range tabInfo {
		if info.Name == field {
			return index, true
		}
	}
	return 0, false
}

// createTableField
//
//	@Description: 创建表字段
//	@param tabName
//	@param attr
func createTableField(tabName, fieldName string, attr fieldAttributeStruct) error {
	alterSql := fmt.Sprintf("alter table %s add %s %s", tabName, fieldName, attr.Type)
	if attr.Default != "" {
		alterSql += fmt.Sprintf(" default %s", attr.Default)
	}
	lg.Info(fmt.Sprintf("%s 添加字段: %s \n sql: %s", tabName, fieldName, alterSql))
	_, err := Db.Exec(alterSql)
	return err
}

// oldTableFieldsFunc
//
//	@Description: 获取旧表字段
//	@param oldTableInfo
//	@return []string
func oldTableFieldsFunc(oldTableInfo []tableInfoStruct) []string {
	var oldTableFields []string
	for _, t := range oldTableInfo {
		oldTableFields = append(oldTableFields, t.Name)
	}
	return oldTableFields
}

// updateTable
//
//	@Description: 更新表字段
//
// @param oldTableFields 旧表字段
//
//	@param newTable 新表
func updateTable(oldTableFields []string, newTable tableStructure) {
	// SQLite 仅仅支持 ALTER TABLE 语句的一部分功能,不能直接修改表字段
	// 步骤一, 创建临时表
	tabNameTmp := newTable.tableName + "_" + tool.RandSeq(6) + "_tmp"
	// 创建临时表
	if err := createTable(tabNameTmp); err != nil {
		panic(err)
	}
	// 步骤二，创建表字段
	structure := reflect.TypeOf(newTable.structure)
	var args = []string{"id"}
	for i := 0; i < structure.NumField(); i++ {
		// 字段名称
		fieldName := structure.Field(i).Tag.Get("db")
		if fieldName == "" || structure.Field(i).Tag.Get("sqlLite") == "" || fieldName == "id" {
			continue
		}
		args = append(args, fieldName)
		// 将标签的内容转换成列表
		attribute := strings.Split(structure.Field(i).Tag.Get("sqlLite"), ",")
		attr := fieldAttribute(attribute)
		// 创建字段
		if err := createTableField(tabNameTmp, fieldName, attr); err != nil {
			panic(err)
		}
	}

	// 将老表里不存在的字段，踢掉
	for i, f := range args {
		if !tool.IsElementStr(oldTableFields, f) {
			tool.DeleteStringElement(&args, i)
		}
	}

	// 将字段列表转换为字符串
	argsStr := strings.Join(args, ",")

	// 步骤三，导入数据
	importSql := fmt.Sprintf("insert into %s(%s) select %s from %s", tabNameTmp, argsStr, argsStr, newTable.tableName)
	if _, err := Db.Exec(importSql); err != nil {
		panic(err)
	}
	// 步骤四，删除旧表，并改名
	// 删除旧表
	if err := dropTable(newTable.tableName); err != nil {
		panic(err)
	}
	// 新表改名
	if err := changeTable(tabNameTmp, newTable.tableName); err != nil {
		panic(err)
	}
}

// generator
//
//	@Description: 表结构生成器
//	@param field
//	@param OldTabInfo
//	@param newTable
func generator(structure reflect.Type, OldTabInfo []tableInfoStruct, newTable tableStructure) {
	for i := 0; i < structure.NumField(); i++ {
		var field = structure.Field(i)
		if field.Type.Kind() == reflect.Struct {
			generator(field.Type, OldTabInfo, newTable)
		}
		fieldName := field.Tag.Get("db")
		// 无db标签 与id字段 直接跳过
		if fieldName == "" || field.Tag.Get("sqlLite") == "" || fieldName == "id" {
			continue
		}
		// 将标签的内容转换成列表
		attribute := strings.Split(field.Tag.Get("sqlLite"), ",")
		// 获取表详情
		// 表中是否有这个字段
		if index, ok := isTableField(OldTabInfo, fieldName); ok {
			// 修改字段属性
			attr := fieldAttribute(attribute)
			// strings.ToLower 全部转为大写
			if strings.ToLower(attr.Default) != strings.ToLower(OldTabInfo[index].DfltValue.String) || strings.ToLower(attr.Type) != strings.ToLower(OldTabInfo[index].Type) {
				updateTable(oldTableFieldsFunc(OldTabInfo), newTable)
				break
			}
		} else {
			//	创建字段
			attr := fieldAttribute(attribute)
			if err := createTableField(newTable.tableName, fieldName, attr); err != nil {
				lg.Error(err)
				panic(err)
			}
		}
	}

}

// TabTask
//
//	@Description: 表任务
func TabTask() {
	for _, newTable := range formRegistration {
		var selectTabQueryStruct struct {
			Name string `db:"name"` // 表名称
		}
		// 判断表是否存在
		selectTab := `SELECT name FROM sqlite_master WHERE type = 'table' and name = ?;`
		if err := Db.Get(&selectTabQueryStruct, selectTab, newTable.tableName); err != nil {
			// 结果是否为空
			if NoRows(err) {
				if err := createTable(newTable.tableName); err != nil {
					panic(err)
				}
			} else {
				// 其他报错
				lg.Error(err)
				panic(err)
			}
		}
		// 校验表结构
		// OldTabInfo 旧表
		var OldTabInfo []tableInfoStruct
		if err := Db.Select(&OldTabInfo, fmt.Sprintf("PRAGMA TABLE_INFO (%s);", newTable.tableName), newTable.tableName); err != nil {
			lg.Error(err)
			panic(err)
		}
		// 生成表结构
		structure := reflect.TypeOf(newTable.structure)
		generator(structure, OldTabInfo, newTable)
	}
}
