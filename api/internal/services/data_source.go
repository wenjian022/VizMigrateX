package services

import (
	"VizMigrateX/internal/models"
	"VizMigrateX/internal/pkg/dataSource"
	"VizMigrateX/internal/pkg/lg"
	"VizMigrateX/internal/pkg/page"
	"errors"
	"fmt"
	GormLogger "gorm.io/gorm/logger"
	"strings"
)

type DataSourceService struct{}

// DataSourceListQueryStruct
// @Description: 数据源列表查询
type DataSourceListQueryStruct struct {
	page.PagingQueryStruct

	DataSourceName    string `form:"dataSourceName"`    // 数据源名称
	ConnectionAddress string `form:"connectionAddress"` // 连接地址
	DatabaseType      string `form:"databaseType"`      // 数据库类型
	EnvironmentName   string `form:"environmentName"`   // 环境
}

// DataSourceJsonStruct
// @Description: 创建数据源
type DataSourceJsonStruct struct {
	DataSourceName       string   `json:"dataSourceName" binding:"required"`    // 数据源名称
	ConnectionAddress    string   `json:"connectionAddress" binding:"required"` // 连接地址
	ConnectionPort       int      `json:"connectionPort" binding:"required"`    // 连接端口
	DatabaseAccount      string   `json:"databaseAccount" binding:"required"`   // 数据库账号
	DatabasePassword     string   `json:"databasePassword" binding:"required"`  // 数据库密码
	DatabaseType         string   `json:"databaseType" binding:"required"`      // 数据库类型
	Environment          uint     `json:"environment" binding:"required"`       // 环境
	Label                []string `json:"label"`                                // 标签
	Explain              string   `json:"explain"`                              // 说明
	AdditionalParameters string   `json:"additionalParameters"`                 // 额外参数
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

	DataSourceModel models.DataSource
}

type dataSourceRespJsonStruct struct {
	Id                int                  `json:"id"`
	CreatedAt         models.LocalTime     `json:"created_at"`
	DataSourceName    string               `json:"data_source_name"`
	ConnectionAddress string               `json:"connection_address"`
	ConnectionPort    int                  `json:"connection_port"`
	DatabaseAccount   string               `json:"database_account"`
	DatabaseType      string               `json:"database_type"`
	Label             models.LabelListType `json:"label" gorm:"column:label"`
	CreatorName       string               `json:"creator_name"`
	EnvName           string               `json:"environment_name" gorm:"column:environment_name"`
	EnvColor          string               `json:"environment_color" gorm:"column:environment_color"`
	ConnectionLog     string               `json:"connection_log"`
	State             int                  `json:"state"`
}

func (thisQuery *DataSourceListQueryStruct) Query() (page.PagingDataResStruct, error) {
	dataSourceDB := models.DB.Model(&models.DataSource{})

	if thisQuery.EnvironmentName != "all" && thisQuery.EnvironmentName != "" {
		dataSourceDB.Where("environment_name = ?", thisQuery.EnvironmentName)
	}

	if thisQuery.DatabaseType != "all" && thisQuery.DatabaseType != "" {
		dataSourceDB.Where("database_type = ?", thisQuery.DatabaseType)
	}

	if thisQuery.ConnectionAddress != "" {
		dataSourceDB.Where("connection_address = ?", thisQuery.ConnectionAddress)
	}

	if thisQuery.DataSourceName != "" {
		dataSourceDB.Where("data_source_name = ?", thisQuery.DataSourceName)
	}

	var dataSourceModels []dataSourceRespJsonStruct
	var total int64
	if err := dataSourceDB.Scopes(page.Paginate(&thisQuery.PagingQueryStruct)).
		Select("ds.id,ds.created_at,data_source_name, connection_address,connection_port,database_account,database_type,connection_log,label,state,u.name AS creator_name, e.name AS environment_name,e.color AS environment_color").
		Joins("as ds LEFT JOIN  users u ON ds.creator_id = u.id LEFT JOIN environments e ON ds.environment_id = e.id").Find(&dataSourceModels).
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

// envDoesItExist
//
//	@Description:
//	@receiver thisDataSourceJson
func (thisDataSourceJson *DataSourceJsonStruct) envDoesItExist() error {
	err := models.DB.Where("id = ?", thisDataSourceJson.Environment).Take(&models.Environment{}).Error
	if err != nil {
		if errors.Is(GormLogger.ErrRecordNotFound, err) {
			return fmt.Errorf("环境不存在")
		}
		lg.Logger.Errorln(err.Error())
		return err
	}
	return nil
}

// labelDoesItExist
//
//	@Description:
//	@receiver thisDataSourceJson
func (thisDataSourceJson *DataSourceJsonStruct) labelDoesItExist() error {
	var labelModels []models.Label
	err := models.DB.Where("name in ?", thisDataSourceJson.Label).Find(&labelModels).Error
	if err != nil {
		lg.Logger.Errorln(err.Error())
		return err
	}

	// 只校验个数
	if len(labelModels) != len(thisDataSourceJson.Label) {
		return fmt.Errorf("标签不存在")
	}
	return nil
}

// Create
//
//	@Description:
//	@receiver thisDataSourceJson
//	@return error
func (thisDataSourceJson *DataSourceJsonStruct) Create(creatorID uint) (int64, error) {
	// 唯一校验
	if err := uniqueVerification(0, thisDataSourceJson.DataSourceName); err != nil {
		lg.Logger.Errorln(err.Error())
		return 0, err
	}

	// 校验环境是否存在
	if err := thisDataSourceJson.envDoesItExist(); err != nil {
		return 0, err
	}

	// 校验标签是否存在
	if err := thisDataSourceJson.labelDoesItExist(); err != nil {
		return 0, err
	}

	// 创建
	res := models.DB.Omit().Create(&models.DataSource{
		DataSourceName:       thisDataSourceJson.DataSourceName,
		ConnectionAddress:    thisDataSourceJson.ConnectionAddress,
		ConnectionPort:       thisDataSourceJson.ConnectionPort,
		DatabaseAccount:      thisDataSourceJson.DatabaseAccount,
		DatabasePassword:     thisDataSourceJson.DatabasePassword,
		DatabaseType:         thisDataSourceJson.DatabaseType,
		Environment:          models.Environment{BasicModel: models.BasicModel{ID: thisDataSourceJson.Environment}},
		Creator:              models.User{BasicModel: models.BasicModel{ID: creatorID}},
		Label:                thisDataSourceJson.Label,
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

	thisUri.DataSourceModel = dataSourceModel
	return nil
}

// Put
//
//	@Description: 更新数据源
//	@receiver thisDataSourceJson
//	@return error
func (thisUri *DataSourceUriIDStruct) Put(dataSourceJson DataSourceJsonStruct) error {
	// 唯一校验
	if err := uniqueVerification(thisUri.DataSourceID, dataSourceJson.DataSourceName); err != nil {
		lg.Logger.Errorln(err.Error())
		return err
	}

	if err := models.DB.Model(&thisUri.DataSourceModel).Updates(&dataSourceJson).Error; err != nil {
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
	if err := models.DB.Model(&thisUri.DataSourceModel).Delete("id = ?", thisUri.DataSourceID).Error; err != nil {
		lg.Logger.Errorln(err.Error())
		return err
	}

	return nil
}

// testConnection
//
//	@Description: 测试数据源连接
//	@param dataSourceModel
//	@return error
func testConnection(dataSourceModel models.DataSource) error {
	basics := dataSource.BasicsDatabases{
		ConnectionAddress:    dataSourceModel.ConnectionAddress,
		ConnectionPort:       dataSourceModel.ConnectionPort,
		DatabaseAccount:      dataSourceModel.DatabaseAccount,
		DatabasePassword:     dataSourceModel.DatabasePassword,
		AdditionalParameters: dataSourceModel.AdditionalParameters,
	}

	var database dataSource.DatabaseInterface
	// 匹配对应的连接方式
	switch dataSourceModel.DatabaseType {
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

func (thisUri *DataSourceUriIDStruct) TestConnection() error {
	thisUri.DataSourceModel.DatabaseType = strings.Split(thisUri.DataSourceModel.DatabaseType, "_")[1]
	err := testConnection(thisUri.DataSourceModel)

	if err != nil {
		thisUri.DataSourceModel.ChangeState(1, err.Error())
		return err
	}

	thisUri.DataSourceModel.ChangeState(0, "数据源可用,连接成功")
	return nil
}

func (thisDataSourceJson *DataSourceTestConnectionJsonStruct) TestConnection() error {
	return testConnection(models.DataSource{
		ConnectionAddress:    thisDataSourceJson.ConnectionAddress,
		ConnectionPort:       thisDataSourceJson.ConnectionPort,
		DatabaseAccount:      thisDataSourceJson.DatabaseAccount,
		DatabasePassword:     thisDataSourceJson.DatabasePassword,
		AdditionalParameters: thisDataSourceJson.AdditionalParameters,
		DatabaseType:         thisDataSourceJson.DatabaseType,
	})
}
