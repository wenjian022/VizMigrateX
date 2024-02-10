package services

import (
	"VizMigrateX/internal/models"
	"VizMigrateX/internal/pkg/dataSource"
	"VizMigrateX/internal/pkg/lg"
	"VizMigrateX/internal/pkg/page"
	"errors"
	GormLogger "gorm.io/gorm/logger"
)

type DataSourceService struct{}

// DataSourceListQueryStruct
// @Description: 数据源列表查询
type DataSourceListQueryStruct struct {
	page.PagingQueryStruct

	DataSourceName    string `form:"dataSourceName"`    // 数据源名称
	ConnectionAddress string `form:"connectionAddress"` // 连接地址
	DatabaseType      string `form:"databaseType"`      // 数据库类型
	//Environment       string `form:"environment"`       // 环境
}

// DataSourceJsonStruct
// @Description: 创建数据源
type DataSourceJsonStruct struct {
	DataSourceName    string `json:"dataSourceName" binding:"required"`    // 数据源名称
	ConnectionAddress string `json:"connectionAddress" binding:"required"` // 连接地址
	ConnectionPort    int    `json:"connectionPort" binding:"required"`    // 连接端口
	DatabaseAccount   string `json:"databaseAccount" binding:"required"`   // 数据库账号
	DatabasePassword  string `json:"databasePassword" binding:"required"`  // 数据库密码
	DatabaseType      string `json:"databaseType" binding:"required"`      // 数据库类型
	//Environment          string `json:"environment" binding:"required"`       // 环境
	Explain              string `json:"explain"`              // 说明
	AdditionalParameters string `json:"additionalParameters"` // 额外参数
}

// DataSourceTestConnectionJsonStruct
// @Description: 测试连接
type DataSourceTestConnectionJsonStruct struct {
	ConnectionAddress    string `json:"connectionAddress" binding:"required"` // 连接地址
	ConnectionPort       int    `json:"connectionPort" binding:"required"`    // 连接端口
	DatabaseAccount      string `json:"databaseAccount" binding:"required"`   // 数据库账号
	DatabasePassword     string `json:"databasePassword" binding:"required"`  // 数据库密码
	DatabaseType         string `json:"databaseType" binding:"required"`      // 数据库类型
	AdditionalParameters string `json:"additionalParameters"`                 // 额外参数
}

type DataSourceUriIDStruct struct {
	DataSourceID uint `uri:"dataSourceID" binding:"required"`

	dataSourceModel models.DataSource
}

func (thisQuery *DataSourceListQueryStruct) Query() (page.PagingDataResStruct, error) {
	dataSourceDB := models.DB.Model(&models.DataSource{})

	//if thisQuery.Environment != "all" && thisQuery.Environment != "" {
	//	dataSourceDB.Where("environment = ?", thisQuery.Environment)
	//}

	if thisQuery.DatabaseType != "all" && thisQuery.DatabaseType != "" {
		dataSourceDB.Where("database_type = ?", thisQuery.DatabaseType)
	}

	if thisQuery.ConnectionAddress != "" {
		dataSourceDB.Where("connection_address = ?", thisQuery.ConnectionAddress)
	}

	if thisQuery.DataSourceName != "" {
		dataSourceDB.Where("data_source_name = ?", thisQuery.DataSourceName)
	}

	var dataSourceModels models.DataSource
	var total int64
	if err := dataSourceDB.Scopes(page.Paginate(&thisQuery.PagingQueryStruct)).Select("id").Find(&dataSourceModels).
		Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		lg.Logger.Errorln(err.Error())
		return page.PagingDataResStruct{}, err
	}

	return page.PagingDataResStruct{
		Data:              dataSourceModels,
		Total:             total,
		PagingQueryStruct: thisQuery.PagingQueryStruct,
	}, nil
}

// uniqueVerification
//
//	@Description: 唯一校验
//	@param id ID
//	@param dataSourceName 数据源名称
func uniqueVerification(id uint, dataSourceName string) error {
	err := models.DB.Model(&models.DataSource{}).Where("data_source_name = ? and id != ?", dataSourceName, id).Take(&models.DataSource{}).Error
	if err != nil {
		if errors.Is(err, GormLogger.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return errors.New("数据源名称重复")
}

// Create
//
//	@Description:
//	@receiver thisDataSourceJson
//	@return error
func (thisDataSourceJson *DataSourceJsonStruct) Create() (int64, error) {
	// 唯一校验
	if err := uniqueVerification(0, thisDataSourceJson.DataSourceName); err != nil {
		lg.Logger.Errorln(err.Error())
		return 0, err
	}

	// 创建
	res := models.DB.Model(&models.DataSource{}).Create(&models.DataSource{
		DataSourceName:    thisDataSourceJson.DataSourceName,
		ConnectionAddress: thisDataSourceJson.ConnectionAddress,
		ConnectionPort:    thisDataSourceJson.ConnectionPort,
		DatabaseAccount:   thisDataSourceJson.DatabaseAccount,
		DatabasePassword:  thisDataSourceJson.DatabasePassword,
		DatabaseType:      thisDataSourceJson.DatabaseType,
		//Environment:          thisDataSourceJson.Environment,
		Explain:              thisDataSourceJson.Explain,
		AdditionalParameters: thisDataSourceJson.AdditionalParameters,
	})
	if res.Error != nil {
		lg.Logger.Errorln(res.Error)
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

// Query
//
//	@Description: uriID 查询并获取数据源对象
//	@receiver thisUri
//	@return error
func (thisUri *DataSourceUriIDStruct) Query() error {
	var dataSourceModel models.DataSource
	if err := models.DB.Model(&dataSourceModel).Where("id = ?", thisUri.DataSourceID).Take(&dataSourceModel).Error; err != nil {
		if errors.Is(err, GormLogger.ErrRecordNotFound) {
			return errors.New("没有查询到这个数据源")
		}
		lg.Logger.Errorln(err.Error())
		return err
	}

	thisUri.dataSourceModel = dataSourceModel
	return nil
}

// Put
//
//	@Description: 更新数据源
//	@receiver thisDataSourceJson
//	@return error
func (thisUri *DataSourceUriIDStruct) Put(dataSourceJson DataSourceJsonStruct) error {
	// 唯一校验
	if err := uniqueVerification(0, dataSourceJson.DataSourceName); err != nil {
		lg.Logger.Errorln(err.Error())
		return err
	}

	if err := models.DB.Model(&thisUri.dataSourceModel).Updates(&dataSourceJson).Error; err != nil {
		lg.Logger.Errorln(err.Error())
		return err
	}

	return nil
}

// Del
//
//	@Description: 删除数据源
//	@receiver thisDataSourceJson
//	@return error
func (thisUri *DataSourceUriIDStruct) Del() error {
	if err := models.DB.Model(&thisUri.dataSourceModel).Delete("id = ?", thisUri.DataSourceID).Error; err != nil {
		lg.Logger.Errorln(err.Error())
		return err
	}

	return nil
}

// TestConnection
//
//	@Description 测试数据源连接
//	@receiver *DataSourceService
//	@param dataSourceJson
//	@return error
func (thisDataSourceJson *DataSourceTestConnectionJsonStruct) TestConnection() error {
	basics := dataSource.BasicsDatabases{
		ConnectionAddress:    thisDataSourceJson.ConnectionAddress,
		ConnectionPort:       thisDataSourceJson.ConnectionPort,
		DatabaseAccount:      thisDataSourceJson.DatabaseAccount,
		DatabasePassword:     thisDataSourceJson.DatabasePassword,
		AdditionalParameters: thisDataSourceJson.AdditionalParameters,
	}

	var database dataSource.DatabaseInterface
	// 匹配对应的连接方式
	switch thisDataSourceJson.DatabaseType {
	case "mysql":
		database = &dataSource.MysqlStruct{BasicsDatabases: basics}
	default:
		return dataSource.NotSupportedErr
	}

	if err := database.Open(); err != nil {
		return err
	}

	defer func() { database.Close() }()

	return nil
}
