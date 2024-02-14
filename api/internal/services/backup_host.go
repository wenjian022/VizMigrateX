package services

import (
	"VizMigrateX/internal/models"
	"VizMigrateX/internal/pkg/lg"
	"errors"
	"fmt"
	GormLogger "gorm.io/gorm/logger"
)

type BackupHostService struct{}

type BackupHostJsonStruct struct {
	HostName          string `json:"hostName" binding:"required"`
	ConnectionAddress string `json:"connectionAddress" binding:"required"`
	ConnectionPort    int    `json:"connectionPort" binding:"required"`
	LineType          string `json:"lineType" binding:"required"`
	HostAccount       string `json:"hostAccount" binding:"required"`
	HostPassword      string `json:"hostPassword" binding:"required"`
	BackupDirectory   string `json:"backupDirectory" binding:"required"`
}

// uniqueVerification
//
//	@Description: 唯一校验
//	@receiver thisJson
func (thisJson *BackupHostJsonStruct) uniqueVerification(id uint) error {
	err := models.DB.Where("host_name = ? and id != ?", thisJson.HostName, id).Take(&models.BackupHost{}).Error
	if err != nil {
		if errors.Is(GormLogger.ErrRecordNotFound, err) {
			return nil
		}
		lg.Logger.Errorln(err.Error())
		return err
	}

	return fmt.Errorf("备份存储主机名称重复")
}

// Create
//
//	@Description: 创建备份存储服务器
//	@receiver thisJson
//	@return uint
//	@return error
func (thisJson *BackupHostJsonStruct) Create(creatorId uint) (uint, error) {
	if err := thisJson.uniqueVerification(0); err != nil {
		return 0, err
	}

	//
	models.DB.Create(&models.BackupHost{
		Creator:           models.User{BasicModel: models.BasicModel{ID: creatorId}},
		HostName:          thisJson.HostName,
		ConnectionAddress: thisJson.ConnectionAddress,
		ConnectionPort:    thisJson.ConnectionPort,
		LineType:          thisJson.LineType,
		HostAccount:       thisJson.HostAccount,
		HostPassword:      thisJson.HostPassword,
		BackupDirectory:   thisJson.BackupDirectory,
	})

	return 0, nil
}
