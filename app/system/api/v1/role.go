package v1

import (
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
