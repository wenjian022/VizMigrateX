package services

import (
	"VizMigrateX/internal/models"
	"VizMigrateX/internal/pkg/lg"
	"VizMigrateX/internal/pkg/page"
	"errors"
	"fmt"
	GormLogger "gorm.io/gorm/logger"
)

type EnvQueryStruct struct {
	page.PagingQueryStruct

	Name string `form:"name"`
}

type LabelQueryStruct struct {
	page.PagingQueryStruct

	Name string `json:"name" binding:"required"`
}

// EnvCreateJsonStruct
// @Description: 创建环境json
type EnvCreateJsonStruct struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

// LabelCreateJsonStruct
// @Description: 创建标签json
type LabelCreateJsonStruct struct {
	Name string `json:"name" binding:"required"`
}

type EnvUriEnvIdStruct struct {
	EnvId uint `uri:"envID"`
}

type LabelUriLabelIdStruct struct {
	LabelId uint `uri:"labelID"`
}

func (thisQuery *EnvQueryStruct) Query() (page.PagingDataResStruct, error) {
	envDB := models.DB.Model(&models.Environment{})

	if thisQuery.Name != "" {
		envDB.Where("name = ?", thisQuery.Name)
	}

	var envModels []struct {
		Id              uint             `json:"id"`
		CreatedAt       models.LocalTime `json:"created_at"`
		EnvironmentName string           `json:"environment_name"`
		CreatorName     interface{}      `json:"creator_name" default:"-"`
		DataSourceCount int              `json:"data_source_count" default:"0"`
		Color           string           `json:"color"`
	}

	var total int64
	if err := envDB.
		Select("e.created_at AS created_at,e.id AS id,e.name AS environment_name, u.name AS creator_name,COUNT(ds.id) AS data_source_count,color").
		Joins("as e LEFT JOIN users u ON e.creator_id = u.id LEFT JOIN data_sources ds ON e.id = ds.environment_id").
		Group("e.id, e.name, u.name").
		Scopes(page.Paginate(&thisQuery.PagingQueryStruct)).Find(&envModels).
		Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		lg.Logger.Errorln(err)
		return page.PagingDataResStruct{}, err
	}

	return page.PagingDataResStruct{
		Data:              envModels,
		Total:             total,
		PagingQueryStruct: thisQuery.PagingQueryStruct,
	}, nil
}

func (thisQuery *LabelQueryStruct) Query() (page.PagingDataResStruct, error) {
	labelDB := models.DB.Model(&models.Label{})

	if thisQuery.Name != "" {
		labelDB.Where("label_name = ?", thisQuery.Name)
	}

	var LabelModels []struct {
		CreatedAt       models.LocalTime `json:"created_at"`
		LabelId         int              `json:"label_id"`
		LabelName       string           `json:"label_name"`
		CreatorName     interface{}      `json:"creator_name" default:"-"`
		DataSourceCount int              `json:"data_source_count" default:"0"`
	}

	var total int64
	if err := labelDB.Select("label.created_at AS created_at, label.id AS label_id, label.name AS label_name, u.name AS creator_name, COUNT(ds.id) AS data_source_count").
		Joins("as label left join users u on label.creator_id = u.id LEFT JOIN data_sources ds ON INSTR(ds.label, label.name) > 0").
		Group("label.name").
		Scopes(page.Paginate(&thisQuery.PagingQueryStruct)).Find(&LabelModels).
		Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		lg.Logger.Errorln(err)
		return page.PagingDataResStruct{}, err
	}

	return page.PagingDataResStruct{
		Data:              LabelModels,
		Total:             total,
		PagingQueryStruct: thisQuery.PagingQueryStruct,
	}, nil
}

func (thisJson *EnvCreateJsonStruct) Create(userID uint) error {
	// 唯一校验
	err := models.DB.Where("name = ? or color = ? and id != 0", thisJson.Name, thisJson.Color).Take(&models.Environment{}).Error
	if err != nil {
		if errors.Is(GormLogger.ErrRecordNotFound, err) {
			// 创建
			if _err := models.DB.Create(&models.Environment{Name: thisJson.Name, CreatorID: userID, Color: thisJson.Color}).Error; _err != nil {
				lg.Logger.Errorln(err.Error())
				return fmt.Errorf("创建环境失败")
			}
			return nil
		}
		return err
	}
	return fmt.Errorf("环境名称或颜色重复")
}

func (thisJson *LabelCreateJsonStruct) Create(userID uint) error {
	// 唯一校验
	err := models.DB.Where("name = ? and id != 0", thisJson.Name).Take(&models.Label{}).Error
	if err != nil {
		if errors.Is(GormLogger.ErrRecordNotFound, err) {
			// 创建
			if _err := models.DB.Create(&models.Label{Name: thisJson.Name, CreatorID: userID}).Error; _err != nil {
				lg.Logger.Errorln(err.Error())
				return fmt.Errorf("创建标签失败")
			}
			return nil
		}
		return err
	}
	return fmt.Errorf("标签名称重复")
}

func (thisUri *EnvUriEnvIdStruct) Del() error {
	err := models.DB.Where("id = ?", thisUri.EnvId).Delete(&models.Environment{}).Error
	if err != nil {
		if errors.Is(GormLogger.ErrRecordNotFound, err) {
			return fmt.Errorf("删除失败,没有查询到这个环境ID")
		}
		lg.Logger.Errorln(err.Error())
		return err
	}

	return nil
}

func (thisUri *LabelUriLabelIdStruct) Del() error {
	err := models.DB.Where("id = ?", thisUri.LabelId).Delete(&models.Label{}).Error
	if err != nil {
		if errors.Is(GormLogger.ErrRecordNotFound, err) {
			return fmt.Errorf("删除失败,没有查询到这个标签ID")
		}
		lg.Logger.Errorln(err.Error())
		return err
	}

	return nil
}
