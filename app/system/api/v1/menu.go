package v1

import (
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type MenuApi struct {
	service.Menu
}

func (m *MenuApi) Register(ver ginx.VersionFunc) {
	v1 := ver(1, middleware.Auth()).Group("menu")
	{
		v1.GET("tree", m.menuTree)
		v1.GET("", m.menuList)
	}
	v1Private := v1.Group("", middleware.Casbin())
	{
		v1Private.POST("", m.addMenu)
		v1Private.DELETE(":id", m.delMenu)
		v1Private.PUT(":id", m.editMenu)
	}
}

// @Tags      目錄
// @Summary   獲取所有菜單tree
// @Produce   application/json
// @Security  BearerToken
// @Param     dir query bool false "是否只查詢目錄"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/menu/tree [get]
func (m *MenuApi) menuTree(c *gin.Context) {
	var query model.MenuTreeParams
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	result, err := m.MenuTree(query)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      目錄
// @Summary   獲取所有目錄
// @Produce   application/json
// @Security  BearerToken
// @Param     menuName query string false "菜單名稱"
// @Param     status query string false "true or false"
// @Success   200  {object}  response.Response{data=model.BasicFetchResult[model.MenuItem],msg=string}  "操作成功"
// @Router    /v1/menu [get]
func (m *MenuApi) menuList(c *gin.Context) {
	var query model.MenuParams
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	result, err := m.MenuList(query)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      目錄
// @Summary   新增菜單
// @Produce   application/json
// @Security  BearerToken
// @Param     menu body model.AddMenuRequest true "菜單"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/menu [post]
func (m *MenuApi) addMenu(c *gin.Context) {
	var json model.AddMenuRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := m.AddMenu(json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      目錄
// @Summary   刪除菜單
// @Produce   application/json
// @Security  BearerToken
// @Param     id path string true "菜單ID"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/menu/{id} [delete]
func (m *MenuApi) delMenu(c *gin.Context) {
	id := ginx.ParseParamID(c, "id")
	err := m.DelMenu(uint(id))
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      目錄
// @Summary   更新菜單
// @Produce   application/json
// @Security  BearerToken
// @Param     id path string true "菜單ID"
// @Param     menu body model.EditMenuRequest true "目錄"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/menu/{id} [put]
func (m *MenuApi) editMenu(c *gin.Context) {
	id := ginx.ParseParamID(c, "id")
	var json model.EditMenuRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := m.EditMenu(uint(id), json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}
