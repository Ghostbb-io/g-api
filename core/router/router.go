package router

import (
	"github.com/Ghostbb-io/g-api/app"
	"github.com/Ghostbb-io/g-api/core/middleware"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/ioutil"
	"net/http"
	"strings"
)

func Init() {
	engine := ginx.Default()

	// 載入api管理系統
	engine.Use(static.Serve("/", static.LocalFile("resource/html", true)))
	engine.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("resource/html/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		}
	})

	engine.Use(middleware.CorsByRules())
	global.GB_LOG.Info("[GIN] Use middleware cors")

	// 使用zap
	if global.GB_CONFIG.Gin.LogZap {
		engine.Use(middleware.ZapLogger())
		global.GB_LOG.Info("[GIN] Use middleware zap")
	}

	// 限流
	engine.Use(middleware.DefaultLimit())
	global.GB_LOG.Info("[GIN] Use middleware limit-ip")

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
