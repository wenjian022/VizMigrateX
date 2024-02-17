package assetView

import (
	"dataxWeb/models/model/assetDatabasesModel"
	"dataxWeb/models/util"
	"dataxWeb/models/util/myValidate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AssetDatabaseListGet
//
//	@Description: 获取数据库资产列表
//	@param c
func AssetDatabaseListGet(c *gin.Context) {
	var query assetsQueryStruct
	_ = c.ShouldBindQuery(&query)

	rel, err := query.query()
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.UtilSuccess(rel))
}

// AssetDatabasesPost
//
//	@Description: 添加数据库资产
//	@param c
func AssetDatabasesPost(c *gin.Context) {
	var assetsDatabasesJson assetDatabasesModel.AssetsDatabasesModel
	if err := c.ShouldBindJSON(&assetsDatabasesJson); err != nil {
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}
	// 创建
	if err := assetsDatabasesJson.Crate(); err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.UtilSuccess("添加成功"))
}

// AssetDatabasesPut
//
//	@Description: 更新数据库资产
//	@param c
func AssetDatabasesPut(c *gin.Context) {
	var assetUri assetsIdUriStruct
	if err := c.ShouldBindUri(&assetUri); err != nil {
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}

	var assetsDatabasesJson assetDatabasesModel.AssetsDatabasesModel
	if err := c.ShouldBindJSON(&assetsDatabasesJson); err != nil {
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}

	assetInfo, err := assetUri.query()
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	err = assetInfo.Update(assetsDatabasesJson.AssetName, assetsDatabasesJson.DatabasesAddress,
		assetsDatabasesJson.DatabasesUserName, assetsDatabasesJson.DatabasesPassword,
		assetsDatabasesJson.Database, assetsDatabasesJson.DatabasesPort)
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.UtilSuccess("修改成功"))
}

// AssetDatabasesInformationGet
//
//	@Description: 获取数据库资产-库信息
//	@param c
func AssetDatabasesInformationGet(c *gin.Context) {
	var assetsIdUri assetsIdUriStruct
	if err := c.ShouldBindUri(&assetsIdUri); err != nil {
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}

	assetInfo, err := assetsIdUri.query()
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}
	databaseOpen, err := assetInfo.Basics()
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	libraryList, err := libraryListQuery(databaseOpen)
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.UtilSuccess(libraryList))
}

// AssetsDatabasesTablesGet
//
//	@Description: 获取资产库-表信息
//	@param c
func AssetsDatabasesTablesGet(c *gin.Context) {

}

// AssetsConnectionTestPost
//
//	@Description: 测试
//	@param c
func AssetsConnectionTestPost(c *gin.Context) {
	var assetsDatabasesJson assetDatabasesModel.AssetsDatabasesModel
	if err := c.ShouldBindJSON(&assetsDatabasesJson); err != nil {
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}
	// 测试数据库地址
	if err := connectionTest(assetsDatabasesJson); err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.UtilSuccess("ok"))
}

// AssetsDatabasesSelectGet
//
//	@Description: 资产数据库-选择器数据格式返回
//	@param c
func AssetsDatabasesSelectGet(c *gin.Context) {
	assetsDatabasesSelectJson, err := assetsSelectQuery()
	if err != nil {
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.UtilSuccess(assetsDatabasesSelectJson))
}
