package v1

import (
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CasbinApi struct {
	service.Casbin
}

func (a *CasbinApi) Register(ver ginx.VersionFunc) {
	v1 := ver(1, middleware.Auth(), middleware.Casbin()).Group("casbin")
	{
		v1.GET("", a.apiList)
		v1.DELETE("", a.removeAuth)

		v1.GET(":role", a.apiListByRole)
		v1.POST(":role", a.addAuth)
		v1.PUT(":role", a.updateAuth)
	}
}

// @Tags      casbin
// @Summary   獲取所有Api
// @Produce   application/json
// @Security  BearerToken
// @Success   200  {object}  response.Response{data=[]model.ApiResponse,msg=string}  "操作成功"
// @Router    /v1/casbin [get]
func (a *CasbinApi) apiList(c *gin.Context) {
	result := a.ApiList()
	response.OkWithData(c, result)
}

// @Tags      casbin
// @Summary   新增一筆權限
// @Produce   application/json
// @Security  BearerToken
// @Param     role path string true "角色"
// @Param     ApiInfo body model.ApiRequest true "路由資訊"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/casbin/{role} [post]
func (a *CasbinApi) addAuth(c *gin.Context) {
	role := ginx.ParseParam(c, "role")
	var json model.ApiRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := a.AddAuth(role, json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      casbin
// @Summary   更新api權限
// @Produce   application/json
// @Security  BearerToken
// @Param     role path string true "角色"
// @Param     ApiInfo body []model.ApiRequest true "路由資訊"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/casbin/{role} [put]
func (a *CasbinApi) updateAuth(c *gin.Context) {
	role := ginx.ParseParam(c, "role")
	var json []model.ApiRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := a.UpdateAuth(role, json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      casbin
// @Summary   根據角色獲取api列表
// @Produce   application/json
// @Security  BearerToken
// @Param     role path string true "角色"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/casbin/{role} [get]
func (a *CasbinApi) apiListByRole(c *gin.Context) {
	role := ginx.ParseParam(c, "role")
	result := a.ApiListByRole(role)
	response.OkWithData(c, result)
}

// @Tags      casbin
// @Summary   移除api
// @Produce   application/json
// @Security  BearerToken
// @Param     ApiInfo body []model.ApiRequest true "路由資訊"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/casbin [delete]
func (a *CasbinApi) removeAuth(c *gin.Context) {
	var json []model.ApiRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := a.RemoveAuth(json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}
