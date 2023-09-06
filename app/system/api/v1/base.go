package v1

import (
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
	service.Base
}

func (b *BaseApi) Register(ver ginx.VersionFunc) {
	v1 := ver(1).Group("")
	{
		v1.POST("login", b.login)
		v1.POST("register", b.registerUser)
	}
	v1Private := v1.Group("", middleware.Auth())
	{
		v1Private.DELETE("logout", b.logout)
	}
}

// @Tags      系統
// @Summary   登入
// @Produce   application/json
// @Param     Info body model.LoginRequest true "帳號&密碼"
// @Success   200  {object}  response.Response{data=model.LoginResponse,msg=string}  "操作成功"
// @Router    /v1/login [post]
func (b *BaseApi) login(c *gin.Context) {
	var json model.LoginRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	result, err := b.Login(json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      系統
// @Summary   註冊
// @Produce   application/json
// @Param     Info body model.RegisterRequest true "使用者資訊"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/register [post]
func (b *BaseApi) registerUser(c *gin.Context) {
	var json model.RegisterRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := b.RegisterUser(json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      系統
// @Summary   登出
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/logout [delete]
func (b *BaseApi) logout(c *gin.Context) {
	uuid := ginx.GetUserUUID(c)
	tokenStr := ginx.GetToken(c)
	err := b.Logout(uuid, tokenStr)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}
