package v1

import (
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type RoleApi struct {
	service.Role
}

func (r *RoleApi) Register(ver ginx.VersionFunc) {
	v1 := ver(1, middleware.Auth()).Group("role")
	{
		v1.GET("", r.roleList)
		v1.GET("page", r.roleListByPage)
	}
	v1Private := v1.Group("", middleware.Casbin())
	{
		v1Private.PATCH(":role/status", r.setStatus)
		v1Private.PUT(":role", r.editRole)
		v1Private.POST("", r.addRole)
		v1Private.DELETE(":role", r.delRole)
	}
}

// @Tags      角色
// @Summary   獲取所有角色
// @Produce   application/json
// @Security  BearerToken
// @Success   200  {object}  response.Response{data=[]model.RoleListResponse,msg=string}  "操作成功"
// @Router    /v1/role [get]
func (r *RoleApi) roleList(c *gin.Context) {
	result, err := r.RoleList()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      角色
// @Summary   根據頁數獲取所有角色
// @Produce   application/json
// @Security  BearerToken
// @Param     page query int true "頁數"
// @Param     pageSize query int true "每頁幾筆"
// @Param     roleName query string false "角色名稱"
// @Param     status query string false "ture or false"
// @Success   200  {object}  response.Response{data=model.BasicFetchResult[model.RoleItem],msg=string}  "操作成功"
// @Router    /v1/role/page [get]
func (r *RoleApi) roleListByPage(c *gin.Context) {
	var query model.RolePageParams
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	result, err := r.RoleListByPage(query)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}

// @Tags      角色
// @Summary   設定狀態
// @Produce   application/json
// @Security  BearerToken
// @Param     role path string true "角色"
// @Param     in body model.SetStatusRequest true "設定"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/role/{role}/status [patch]
func (r *RoleApi) setStatus(c *gin.Context) {
	role := ginx.ParseParam(c, "role")
	var json model.SetStatusRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := r.SetStatus(role, json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      角色
// @Summary   更新角色
// @Produce   application/json
// @Security  BearerToken
// @Param     role path string true "角色"
// @Param     role body model.EditRoleRequest true "角色資訊"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/role/{role} [put]
func (r *RoleApi) editRole(c *gin.Context) {
	var json model.EditRoleRequest
	role := ginx.ParseParam(c, "role")
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := r.Update(role, json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      角色
// @Summary   更新角色
// @Produce   application/json
// @Security  BearerToken
// @Param     role body model.AddRoleRequest true "角色資訊"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/role [post]
func (r *RoleApi) addRole(c *gin.Context) {
	var json model.AddRoleRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := r.AddRole(json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Tags      角色
// @Summary   刪除角色
// @Produce   application/json
// @Security  BearerToken
// @Param     role path string true "角色"
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/role/{role} [delete]
func (r *RoleApi) delRole(c *gin.Context) {
	role := ginx.ParseParam(c, "role")
	err := r.DelRole(role)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.Ok(c)
}
