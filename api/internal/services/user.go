package services

import (
	"VizMigrateX/internal/models"
	"VizMigrateX/internal/pkg/lg"
	"VizMigrateX/internal/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

type UserService struct{}

// LoginByUsernamePassword
//
//	@Description: 用户登录
//	@receiver userService
//	@param username
//	@param password
//	@return string
//	@return error
func (userService *UserService) LoginByUsernamePassword(username string, password string) (string, error) {

	user := models.User{
		Username: username,
	}
	res := models.DB.First(&user, "username = ? and password = ?", username, password)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		lg.Logger.Errorln(res.Error)
		return "", errors.New("查询用户失败")
	}

	if res.RowsAffected == 0 {
		return "", errors.New("用户名或密码错误")
	}

	token, err := utils.GenerateToken(&utils.Claims{
		Uid:      user.ID,
		Username: user.Username,
	})
	if err != nil {
		lg.Logger.Errorln(err)
		return "", errors.New("获取token失败")
	}

	return token, nil
}

// UserQueries
//
//	@Description: 用户查询
//	@receiver userService
func (userService *UserService) UserQueries(ID uint) (models.User, error) {
	var user models.User
	res := models.DB.First(&user, "id = ?", ID)
	if res.Error != nil {
		lg.Logger.Errorln(res.Error)
		return user, errors.New("查询用户失败")
	}

	if res.RowsAffected == 0 {
		return user, errors.New("未查询到该用户")
	}

	return user, nil
}

// UserInitializationQuery
//
//	@Description:
//	@receiver userService
//	@return bool
func (userService *UserService) UserInitializationQuery() (bool, error) {
	// 通过查看用户表，来做用户初始化
	res := models.DB.Take(&models.User{})
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		lg.Logger.Errorln(res.Error.Error())
		return false, errors.New("错误:" + res.Error.Error())
	}

	return res.RowsAffected >= 1, nil
}

// UserInitialization
//
//	@Description: 用户初始化
//	@receiver userService
//	@param userName
//	@param password
//	@return string
//	@return error
func (userService *UserService) UserInitialization(userName, password string) (string, error) {

	ok, err := userService.UserInitializationQuery()
	if err != nil {
		return "", err
	}

	if ok {
		return "", errors.New("已经初始化过了")
	}

	var token string
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		user := models.User{Name: userName, Username: userName, Password: password, Role: 1}
		res := models.DB.Create(&user)
		if res.Error != nil {
			lg.Logger.Errorln(res.Error)
			return errors.New("初始化失败" + res.Error.Error())
		}
		// 用户登录
		var _err error
		token, _err = utils.GenerateToken(&utils.Claims{
			Uid:      user.ID,
			Username: user.Username,
		})
		if _err != nil {
			lg.Logger.Errorln(res.Error)
			return errors.New("登录失败: " + _err.Error())
		}
		return nil
	})

	return token, err
}
