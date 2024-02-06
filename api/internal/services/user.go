package services

import (
	"VizMigrateX/internal/models"
	Logger "VizMigrateX/internal/pkg/lg"
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
		Logger.Logger.Errorln(res.Error)
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
		Logger.Logger.Errorln(err)
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
		Logger.Logger.Errorln(res.Error)
		return user, errors.New("查询用户失败")
	}

	if res.RowsAffected == 0 {
		return user, errors.New("未查询到该用户")
	}

	return user, nil
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

	return "", nil
}
