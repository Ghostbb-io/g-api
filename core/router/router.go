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

	//engine.Use(static.Serve("/", static.LocalFile("resource/html", true)))
	//engine.NoRoute(func(c *gin.Context) {
	//	accept := c.Request.Header.Get("Accept")
	//	flag := strings.Contains(accept, "text/html")
	//	if flag {
	//		content, err := ioutil.ReadFile("resource/html/index.html")
	//		if (err) != nil {
	//			c.Writer.WriteHeader(404)
	//			c.Writer.WriteString("Not Found")
	//			return
	//		}
	//		c.Writer.WriteHeader(200)
	//		c.Writer.Header().Add("Accept", "text/html")
	//		c.Writer.Write(content)
	//		c.Writer.Flush()
	//	}
	//})

	// cors
	engine.Use(middleware.CorsByRules())
	global.GB_LOG.Info("[GIN] Use middleware cors")

	// zap
	if global.GB_CONFIG.Gin.LogZap {
		engine.Use(middleware.ZapLogger())
		global.GB_LOG.Info("[GIN] Use middleware zap")
	}

	// Ip limit
	engine.Use(middleware.DefaultLimit())
	global.GB_LOG.Info("[GIN] Use middleware limit-ip")

	// Register swagger
	engine.GET(global.GB_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Server test
	engine.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })

	// Set prefix
	engine.SetPrefix(global.GB_CONFIG.System.RouterPrefix)

	// -----Batch registration-----
	for _, plugins := range app.All() {
		ginx.Register(plugins...)
	}
}
