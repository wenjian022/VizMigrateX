package view

import (
	"dataxWeb/models/view/datax"
	"dataxWeb/models/view/user/user"
	"dataxWeb/models/view/user/user/jwt"
	"github.com/gin-gonic/gin"
)

// Route 路由
func Route(api gin.RouterGroup) {
	api.Use(jwt.JWTAuthMiddleware())
	user.Url(api)
	datax.Url(api)
}
