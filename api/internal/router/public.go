package router

import (
	"VizMigrateX/internal/controllers"

	"github.com/gin-gonic/gin"
)

var publicController = new(controllers.PublicController)

func LoadPublicRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	public := r.Group("/public")
	{
		public.GET("/ping", publicController.Ping)
	}
	return public
}
