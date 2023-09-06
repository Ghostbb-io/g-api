package router

import (
	"github.com/Ghostbb-io/g-api/app"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Init() {
	engine := ginx.Default()
	engine.Use(middleware.CorsByRules())
	global.GB_LOG.Info("[GIN] Use middleware cors")

	// 使用zap
	if global.GB_CONFIG.Gin.LogZap {
		engine.Use(middleware.ZapLogger())
		global.GB_LOG.Info("[GIN] Use middleware zap")
	}

	// 註冊 Swagger
	engine.GET(global.GB_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康監測
	engine.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })

	// 設定前綴
	engine.SetPrefix(global.GB_CONFIG.System.RouterPrefix)

	// -----批量註冊-----
	for _, plugins := range app.All() {
		ginx.Register(plugins)
	}
}
