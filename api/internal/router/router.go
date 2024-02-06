package router

import (
	"VizMigrateX/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	// Global middlewares
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())
	Router.Use(middlewares.JWTAuthMiddleware())
	api := Router.Group("/api")
	{
		// public routes, no auth required
		LoadPublicRoutes(api)

		// user routes
		LoadUserRoutes(api)
	}
}
