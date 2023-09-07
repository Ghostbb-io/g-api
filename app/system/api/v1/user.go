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
		v1.GET("me/menu", u.menuByMe)
		v1.PATCH("password", u.changePass)
	}
	v1Private := v1.Group("", middleware.Casbin())
	{
		v1Private.PUT(":id/roles", u.updateRoles)
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
// @Success   200  {object}  response.Response{data=[]model.MenuListResponse,msg=string} "操作成功"
// @Router    /v1/user/me/menu [get]
func (u *UserApi) menuByMe(c *gin.Context) {
	userID := ginx.GetUserID(c)
	result, err := u.MenuList(userID)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}
