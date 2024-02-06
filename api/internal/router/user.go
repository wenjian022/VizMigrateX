package router

import (
	"VizMigrateX/internal/controllers"

	"github.com/gin-gonic/gin"
)

var userController = new(controllers.UserController)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {

	user := r.Group("/users")
	{
		// 获取登录用户基础信息
		user.GET("/user/info", userController.UserInfoGet)
		// 用户登录
		user.POST("/auth/login", userController.UserLoginPost)
		// 用户注销
		user.POST("/auth/logout", userController.UserLogoutPost)
		// 初始化用户
		user.POST("/user/initialization", userController.UserInitializationPost)
	}
	return user
}
