package middleware

import (
	"dataxWeb/models/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PanicMiddleware
//
//	@Description: 异常捕获中间件，放着
func PanicMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, util.UtilError("系统异常，请联系管理员！"))
				return
			}
		}()
		c.Next()
	}
}
