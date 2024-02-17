package user

import (
	model "dataxWeb/models/model/userModel"
	"dataxWeb/models/util"
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/myValidate"
	"dataxWeb/models/util/tool"
	"dataxWeb/models/util/tool/db"
	"dataxWeb/models/view/user/user/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Url 路由表
func Url(api gin.RouterGroup) {
	// 获取登录用户基础信息
	api.GET("/user/info", userInfoGet)
	// 用户登录
	api.POST("/auth/login", userLoginPost)
	// 用户注销
	api.POST("/auth/logout", userLogoutPost)
	// 添加用户
	api.POST("/user/add", userAddPost)
	// 修改用户
	api.POST("/user/edit/:id", userEditPost)
	// 查看所有用户
	api.GET("/users", userListGet)
}

// userInfoGet 获取登录用户基础信息
func userInfoGet(c *gin.Context) {
	user, err := c.Get("USER")
	if !err {
		c.JSON(http.StatusOK, util.UtilError("登陆过期"))
		return
	}
	c.JSON(http.StatusOK, util.UtilSuccess(user))
}

// userLoginPost 用户登录
func userLoginPost(c *gin.Context) {
	type respStruct struct {
		Token string `json:"token"`
	}
	var u userLoginJsonStrict
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusOK, util.UtilError("请输入用户名与密码"))
		return
	}
	userInfo, err := userWhereUserAndPwd(u)
	if err != nil {
		if db.NoRows(err) {
			c.JSON(http.StatusOK, util.UtilError("用户名或密码错误"))
			return
		}
		lg.Error(err)
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}
	// 生成JWT
	var resp respStruct
	resp.Token, _ = jwt.GenToken(userInfo)
	c.JSON(http.StatusOK, util.UtilSuccess(resp))
}

// userLogoutPost 用户注销
func userLogoutPost(c *gin.Context) {
	c.JSON(http.StatusOK, util.UtilSuccess("ok"))
}

// userAddPost 添加用户
func userAddPost(c *gin.Context) {
	var userJson userInsertJsonStruct
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusOK, myValidate.ErrResp(err))
		return
	}
	// 添加用户
	if err := userInsert(userJson); err != nil {
		lg.Error(err)
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.UtilSuccess("添加成功"))
}

// userEditPost
//
//	@Description: 修改用户
//	@param c
func userEditPost(c *gin.Context) {

	c.JSON(http.StatusOK, util.UtilSuccess("修改成功"))
}

// userListGet
//
//	@Description: 查看用户列表
//	@param c
func userListGet(c *gin.Context) {
	//var query tool.PagingDataResStruct
	//_ = c.ShouldBindQuery(&query)
	rel, err := model.UsersAll()
	if err != nil {
		lg.Error(err)
		c.JSON(http.StatusOK, util.UtilError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.UtilSuccess(tool.PagingDataResStruct{
		Data:  rel,
		Total: 0,
		Size:  0,
		Page:  0,
	}))
}
