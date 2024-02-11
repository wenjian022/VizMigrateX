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
	Name string `json:"name" binding:"required"`
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

	var envModels []models.Environment
	var total int64
	if err := envDB.Scopes(page.Paginate(&thisQuery.PagingQueryStruct)).Find(&envModels).
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
		labelDB.Where("name = ?", thisQuery.Name)
	}

	var LabelModels []models.Label
	var total int64
	if err := labelDB.Scopes(page.Paginate(&thisQuery.PagingQueryStruct)).Find(&LabelModels).
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
	err := models.DB.Where("name = ? and id != 0", thisJson.Name).Take(&models.Environment{}).Error
	if err != nil {
		if errors.Is(GormLogger.ErrRecordNotFound, err) {
			// 创建
			if _err := models.DB.Create(&models.Environment{Name: thisJson.Name, CreatorID: userID}).Error; _err != nil {
				lg.Logger.Errorln(err.Error())
				return fmt.Errorf("创建环境失败")
			}
			return nil
		}
		return err
	}
	return fmt.Errorf("环境名称重复")
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
