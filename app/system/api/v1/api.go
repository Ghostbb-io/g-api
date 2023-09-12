package v1

import (
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type ApiApi struct {
	service.Api
}

func (a *ApiApi) Register(ver ginx.VersionFunc) {
	v1 := ver(1, middleware.Auth()).Group("api")
	{
		v1.GET("page", a.apiListByPage)
		v1.GET("tree", a.apiTree)
	}
	v1Private := v1.Group("", middleware.Casbin())
	{
		v1Private.POST("", a.addApi)
		v1Private.PUT(":id", a.editApi)
		v1Private.DELETE(":id", a.delApi)
	}
}

// @Tags      api
// @Summary   根據頁數獲取所有api
// @Produce   application/json
// @Security  BearerToken
// @Param     page query int true "頁數"
// @Param     pageSize query int true "每頁幾筆"
// @Param     path query string false "路徑"
// @Param     group query string false "分組"
// @Param     method query string false "請求"
// @Success   200  {object}  response.Response{data=model.BasicFetchResult[model.ApiItem],msg=string}  "操作成功"
// @Router    /v1/api/page [get]
func (a *ApiApi) apiListByPage(c *gin.Context) {
	var query model.ApiPageParams
	if err := ginx.ParseQuery(c, &query); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	result, err := a.ApiListByPage(query)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      api
// @Summary   新增api
// @Produce   application/json
// @Security  BearerToken
// @Param     api body model.AddApiRequest true "api資訊"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/api [post]
func (a *ApiApi) addApi(c *gin.Context) {
	var json model.AddApiRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := a.AddApi(json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      api
// @Summary   編輯api
// @Produce   application/json
// @Security  BearerToken
// @Param     api body model.EditApiRequest true "api資訊"
// @Param     id path string true "api id"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/api/{id} [put]
func (a *ApiApi) editApi(c *gin.Context) {
	var json model.EditApiRequest
	id := uint(ginx.ParseParamID(c, "id"))
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := a.EditApi(id, json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      api
// @Summary   刪除api
// @Produce   application/json
// @Security  BearerToken
// @Param     id path string true "api id"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/api/{id} [delete]
func (a *ApiApi) delApi(c *gin.Context) {
	id := uint(ginx.ParseParamID(c, "id"))
	err := a.DelApi(id)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      api
// @Summary   獲取api tree
// @Produce   application/json
// @Security  BearerToken
// @Success   200  {object}  response.Response{data=[]model.ApiTree,msg=string}  "操作成功"
// @Router    /v1/api/tree [get]
func (a *ApiApi) apiTree(c *gin.Context) {
	result, err := a.ApiTree()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}
