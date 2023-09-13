package v1

import (
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"

	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	service.User
}

func (u *UserApi) Register(ver ginx.VersionFunc) {
	v1 := ver(1, middleware.Auth()).Group("user")
	{
		v1.GET("me", u.userInfoByMe)
		v1.GET("me/perm", u.permByMe)
		v1.GET("me/route", u.routeByMe)
		v1.PATCH("password", u.changePass)
		v1.GET("check/username", u.isUsernameExist)
		v1.GET("page", u.userListByPage)
	}
	v1Private := v1.Group("", middleware.Casbin())
	{
		v1Private.PUT(":id/roles", u.updateRoles)
		v1Private.GET(":id", u.userInfo)
		v1Private.PUT(":id", u.editUser)
		v1Private.POST("", u.addUser)
		v1Private.DELETE(":id", u.delUser)
		v1Private.PATCH(":id/status", u.setStatus)
	}
}

// @Tags      使用者
// @Summary   獲取使用者資訊
// @Produce   application/json
// @Security  BearerToken
// @Success   200  {object}  response.Response{data=model.UserInfoResponse,msg=string}  "操作成功"
// @Router    /v1/user/me [get]
func (u *UserApi) userInfoByMe(c *gin.Context) {
	userID := ginx.GetUserID(c)
	result, err := u.UserInfo(userID)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      使用者
// @Summary   修改密碼
// @Produce   application/json
// @Security  BearerToken
// @Param     Password body model.ChangePassRequest true "舊密碼&新密碼"
// @Success   200  {object}  response.Response{data=model.UserInfoResponse,msg=string}  "操作成功"
// @Router    /v1/user/password [patch]
func (u *UserApi) changePass(c *gin.Context) {
	var json model.ChangePassRequest
	userID := ginx.GetUserID(c)
	tokenStr := ginx.GetToken(c)

	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := u.ChangePassword(userID, json, tokenStr)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      使用者
// @Summary   更新角色
// @Produce   application/json
// @Security  BearerToken
// @Param     id path string true "使用者id"
// @Param     roles body model.RolesRequest true "角色"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/user/{id}/roles [put]
func (u *UserApi) updateRoles(c *gin.Context) {
	var json model.RolesRequest
	userID := ginx.ParseParamID(c, "id")
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := u.UpdateRoles(uint(userID), json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      使用者
// @Summary   獲取自身按鈕權限
// @Produce   application/json
// @Security  BearerToken
// @Success   200  {object}  response.Response{data=[]string,msg=string} "操作成功"
// @Router    /v1/user/me/perm [get]
func (u *UserApi) permByMe(c *gin.Context) {
	userID := ginx.GetUserID(c)
	result, err := u.PermList(userID)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      使用者
// @Summary   獲取自身menu
// @Produce   application/json
// @Security  BearerToken
// @Success   200  {object}  response.Response{data=[]model.RouteResponse,msg=string} "操作成功"
// @Router    /v1/user/me/route [get]
func (u *UserApi) routeByMe(c *gin.Context) {
	userID := ginx.GetUserID(c)
	result, err := u.RouteList(userID)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      使用者
// @Summary   確認使用者是否存在
// @Produce   application/json
// @Security  BearerToken
// @Param     username query string true "使用者名稱"
// @Success   200  {object}  response.Response{msg=string} "操作成功"
// @Router    /v1/user/check/username [get]
func (u *UserApi) isUsernameExist(c *gin.Context) {
	username := c.Query("username")
	if u.IsUsernameExist(username) {
		response.Ok(c)
	} else {
		response.FailWithMessage(c, "使用者已存在")
	}
}

// @Tags      使用者
// @Summary   根據page, pageSize獲取使用者列表
// @Produce   application/json
// @Security  BearerToken
// @Param     page query int true "頁數"
// @Param     pageSize query int true "每頁幾筆"
// @Param     username query string false "使用者名稱"
// @Param     nickName query string false "暱稱"
// @Success   200  {object}  response.Response{data=model.BasicFetchResult[model.UserItem],msg=string} "操作成功"
// @Router    /v1/user/page [get]
func (u *UserApi) userListByPage(c *gin.Context) {
	var query model.UserPageParams
	if err := ginx.ParseQuery(c, &query); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	result, err := u.UserListByPage(query)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      使用者
// @Summary   編輯使用者
// @Produce   application/json
// @Security  BearerToken
// @Param     id path string true "使用者id"
// @Success   200  {object}  response.Response{msg=string} "操作成功"
// @Router    /v1/user/{id} [put]
func (u *UserApi) editUser(c *gin.Context) {
	userID := uint(ginx.ParseParamID(c, "id"))
	var json model.EditUserRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := u.EditUser(userID, json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      使用者
// @Summary   新增使用者
// @Produce   application/json
// @Security  BearerToken
// @Param     user body model.AddUserRequest true "使用者"
// @Success   200  {object}  response.Response{msg=string} "操作成功"
// @Router    /v1/user [post]
func (u *UserApi) addUser(c *gin.Context) {
	var json model.AddUserRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := u.AddUser(json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      使用者
// @Summary   刪除使用者
// @Produce   application/json
// @Security  BearerToken
// @Param     id path string true "使用者id"
// @Success   200  {object}  response.Response{msg=string} "操作成功"
// @Router    /v1/user/{id} [delete]
func (u *UserApi) delUser(c *gin.Context) {
	userID := uint(ginx.ParseParamID(c, "id"))
	err := u.DelUser(userID)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      使用者
// @Summary   獲取使用者資訊
// @Produce   application/json
// @Security  BearerToken
// @Param     id path string true "使用者id"
// @Success   200  {object}  response.Response{data=model.UserInfoResponse,msg=string} "操作成功"
// @Router    /v1/user/{id} [get]
func (u *UserApi) userInfo(c *gin.Context) {
	id := uint(ginx.ParseParamID(c, "id"))
	result, err := u.UserInfo(id)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      使用者
// @Summary   更新使用者狀態
// @Produce   application/json
// @Security  BearerToken
// @Param     id path string true "使用者id"
// @Param     status body model.SetUserStatusRequest true "狀態"
// @Success   200  {object}  response.Response{msg=string} "操作成功"
// @Router    /v1/user/{id}/status [patch]
func (u *UserApi) setStatus(c *gin.Context) {
	id := uint(ginx.ParseParamID(c, "id"))
	var json model.SetUserStatusRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := u.SetStatus(id, json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}
