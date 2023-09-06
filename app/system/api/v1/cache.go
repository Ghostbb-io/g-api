package v1

import (
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CacheApi struct {
	service.Cache
}

func (c *CacheApi) Register(ver ginx.VersionFunc) {
	v1 := ver(1).Group("cache", middleware.Auth(), middleware.Casbin())
	{
		v1.DELETE("db", c.clearDbCache)
	}
}

// @Tags      緩存
// @Summary   清除資料庫緩存
// @Produce   application/json
// @Security  BearerToken
// @Success   200  {object}  response.Response{msg=string}  "操作成功"
// @Router    /v1/cache/db [delete]
func (c *CacheApi) clearDbCache(ctx *gin.Context) {
	err := c.ClearDbCache()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
