package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReturnStatus struct {
	HttpCode int         `json:"-"`
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

// Success
//
//	@Description: 成功返回
//	@receiver thisStatus
//	@param result
//	@return gin.H
func (thisStatus ReturnStatus) Success(result interface{}) (int, gin.H) {
	return http.StatusOK, gin.H{"code": 0, "result": result}
}

// Error
//
//	@Description: 错误返回
//	@receiver thisStatus
//	@param errStr
//	@return gin.H
func (thisStatus ReturnStatus) Error(errStr string) (int, gin.H) {
	return http.StatusOK, gin.H{"code": 1, "result": errStr}
}
