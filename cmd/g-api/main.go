package main

import (
	"github.com/Ghostbb-io/g-api/core"
	"github.com/Ghostbb-io/g-api/core/server"
	_ "github.com/Ghostbb-io/g-api/core/swagger"
	"github.com/Ghostbb-io/g-api/pkg/global"
)

// @title g-api
// @version 1.0.0
// @description api框架
// @basePath /api
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
// @schemes http https
// @contact.name Ghostbb
// @contact.email mucgll0328@ghostbb.net
func main() {
	global.GB_VP = core.InitViper()
	global.GB_LOG = core.InitZap()     // 初始化Zap
	global.GB_REDIS = core.InitRedis() // 初始化Redis
	global.GB_DB = core.InitGorm()     // 初始化Gorm
	global.GB_DBS = core.InitMGorm()   // 初始化多資料庫
	global.GB_ENG = core.InitEngine()  // 初始化server engine

	defer func() {
		// 程式關閉前，關閉所有資料庫
		if global.GB_DB != nil {
			db, _ := global.GB_DB.DB()
			_ = db.Close()
		}
		for _, v := range global.GB_DBS {
			if v != nil {
				db, _ := v.DB()
				_ = db.Close()
			}
		}
	}()

	server.Run()
}
