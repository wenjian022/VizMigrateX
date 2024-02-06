package router

import (
	"VizMigrateX/internal/controllers"

	"github.com/gin-gonic/gin"
)

var userController = new(controllers.UserController)

func LoadUserRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	user := r.Group("/user")
	{
		// 获取登录用户基础信息
		user.GET("/info", userController.UserInfoGet)
		// 用户登录
		user.POST("/auth/login", userController.UserLoginPost)
		// 用户注销
		user.POST("/auth/logout", userController.UserLogoutPost)
		// 查看是否需要用户初始化
		user.GET("/initialization", userController.UserInitializationGet)
		// 初始化用户
		user.POST("/initialization", userController.UserInitializationPost)
	}
	return user
}
