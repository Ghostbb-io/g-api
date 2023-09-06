# gormx
- 支持連線多種資料庫，目前支援mssql，mysql，pgsql。  
- 使用redis進行查詢緩存
- 使用zap進行log輸出

## 使用方法
```go
package main

import (
    "github.com/Ghostbb-io/g-api/pkg/gormx"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/config"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/gzap"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/storage"
)

func main() {
	zapConfig := gzap.Config{
		Level:         "info",
		Prefix:        "[Gorm] ",
		Format:        "console", // console or json
		Director:      "/log",
		EncodeLevel:   "LowercaseLevelEncoder",
		StacktraceKey: "stacktrace",
		MaxAge:        5,
		ShowLine:      true,
		LogInConsole:  false,
	}

	// database 設定
	dbConfig := &gormx.Config{
		DBType:       gormx.Mysql,
		Dsn:          "gormx:123456@tcp(localhost:3306)/dbname",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		LogMode:      "info",
		UseZap:       true,
		Zap:          gormx.Zap{Config: zapConfig},
	}

	// 緩存設定
	cacheConfig := &config.CacheConfig{
		CacheLevel: config.CacheLevelAll,
		CacheStorage: storage.NewRedis(&storage.RedisStoreConfig{
			Client: redis.client,
		}),
		InvalidateWhenUpdate: true,
		CacheTTL:             5000,
		CacheMaxItemCnt:      50,
		DebugMode:            true,
	}

	db, err := gormx.New(dbConfig, cacheConfig)
}
```