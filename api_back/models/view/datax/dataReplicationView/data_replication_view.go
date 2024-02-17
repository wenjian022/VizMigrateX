package dataReplicationView

import (
	"dataxWeb/models/util"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/myValidate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DataReplicationTaskListGet
//
//	@Description: 数据复制任务列表
//	@param c
func DataReplicationTaskListGet(c *gin.Context) {
	var dataReplicationTaskQuery dataReplicationTaskQueryStruct
	_ = c.ShouldBindQuery(&dataReplicationTaskQuery)

	taskList, err := dataReplicationTaskQuery.query()
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.UtilSuccess(taskList))
}

// DataReplicationTaskInfoGet
//
//	@Description: 获取任务详情
//	@param c
func DataReplicationTaskInfoGet(c *gin.Context) {
	var dataReplicationTaskIDUri dataReplicationTaskIDUriStruct
	if err := c.ShouldBindUri(&dataReplicationTaskIDUri); err != nil {
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}
	// 获取详情
	taskInfo, err := dataReplicationTaskIDUri.query()
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.UtilSuccess(taskInfo))
}

// DataReplicationTaskPost
//
//	@Description: 添加数据复制任务
//	@param c
func DataReplicationTaskPost(c *gin.Context) {
	var dataReplicationStep1Json dataReplicationStep1JsonStruct
	if err := c.ShouldBindJSON(&dataReplicationStep1Json); err != nil {
		lg.Error(err)
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}

	rowID, err := dataReplicationStep1Json.create()
	if err != nil {
		lg.Error(err)
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.UtilSuccess(gin.H{"rowID": rowID}))
}

// DataReplicationTaskPut
//
//	@Description: 修改数据复制任务
//	@param c
func DataReplicationTaskPut(c *gin.Context) {
	var uri dataReplicationStepTaskIDUriStruct
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}
	taskInfo, err := uri.query()
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	var dataReplicationStep dataReplicationStepInterface
	//
	switch uri.Step {
	case 1:
	case 2:
		dataReplicationStep = dataReplicationStep2JsonStruct{}
	default:
		c.JSON(http.StatusOK, util.UtilError("流程错误"))
		return
	}

	// 更新数据
	if err := dataReplicationStep.update(c, taskInfo); err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.UtilSuccess("ok"))
}
