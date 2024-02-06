package controllers

import (
	Logger "VizMigrateX/internal/pkg/lg"
	"VizMigrateX/internal/pkg/response"
	"VizMigrateX/internal/pkg/utils"
	"VizMigrateX/internal/services"
	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

type UserController struct{}

type userLoginJsonStruct struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// userInitializationJsonStruct
// @Description: 用户初始化需要的参数
type userInitializationJsonStruct struct {
	userLoginJsonStruct
}

// UserInfoGet
//
//	@Description: 获取用户信息
//	@receiver userController
//	@param ctx
func (userController *UserController) UserInfoGet(ctx *gin.Context) {
	value := ctx.MustGet("USER").(*utils.Claims)
	user, err := userService.UserQueries(value.Uid)

	if err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(user))
}

// UserLoginPost
//
//	@Description: 用户登录
//	@receiver userController
//	@param ctx
func (userController *UserController) UserLoginPost(ctx *gin.Context) {
	var userLoginJson userLoginJsonStruct
	if err := ctx.ShouldBindJSON(&userLoginJson); err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	token, err := userService.LoginByUsernamePassword(userLoginJson.UserName, userLoginJson.Password)
	if err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(map[string]string{"token": token}))
}

// UserLogoutPost
//
//	@Description: 用户退出
//	@receiver userController
//	@param ctx
func (userController *UserController) UserLogoutPost(ctx *gin.Context) {
	ctx.JSON(response.ReturnStatus{}.Success("ok"))
}

// UserInitializationGet
//
//	@Description: 查看是否需要用户初始化
//	@receiver userController
func (userController *UserController) UserInitializationGet(ctx *gin.Context) {
	ok, err := userService.UserInitializationQuery()
	if err != nil {
		Logger.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(ok))
}

// UserInitializationPost
//
//	@Description: 初始化用户
//	@receiver userController
//	@param ctx
func (userController *UserController) UserInitializationPost(ctx *gin.Context) {
	var userInitializationJson userInitializationJsonStruct
	if err := ctx.ShouldBindJSON(&userInitializationJson); err != nil {
		Logger.Logger.Errorln(err.Error())
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	//
	token, err := userService.UserInitialization(userInitializationJson.UserName, userInitializationJson.Password)
	if err != nil {
		ctx.JSON(response.ReturnStatus{}.Error(err.Error()))
		return
	}

	ctx.JSON(response.ReturnStatus{}.Success(gin.H{"token": token}))
}
