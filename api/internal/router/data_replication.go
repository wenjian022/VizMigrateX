package router

import (
	"VizMigrateX/internal/controllers"
	"github.com/gin-gonic/gin"
)

var dataReplicationController = new(controllers.DataReplicationControllers)

func LoadDataReplicationRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	dataReplication := r.Group("/dataReplication")
	{
		//
	}
	return dataReplication
}
