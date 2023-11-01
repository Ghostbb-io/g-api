package core

import (
	"fmt"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/gormx"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/config"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/storage"
	"github.com/Ghostbb-io/g-api/pkg/gormx/zapgorm"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func InitGorm() *gorm.DB {
	var db *gorm.DB

	// database 設定
	dbConfig := &gormx.Config{
		DBType:       global.GB_CONFIG.Gorm.DBType(),
		Dsn:          global.GB_CONFIG.Gorm.Dsn(),
		MaxIdleConns: global.GB_CONFIG.Gorm.MaxIdleConns,
		MaxOpenConns: global.GB_CONFIG.Gorm.MaxOpenConns,
		LogMode:      global.GB_CONFIG.Gorm.LogMode,
		UseZap:       global.GB_CONFIG.Gorm.LogZap,
		Zap: gormx.Zap{
			LogFolderName: fmt.Sprintf("gorm(%s)", "main"),
			Logger:        global.GB_LOG,
		},
	}

	// 緩存設定
	if global.GB_CONFIG.Gorm.Cache {
		cacheConfig := &config.CacheConfig{
			CacheLevel: global.GB_CONFIG.Gorm.CacheLevel(),
			CacheStorage: storage.NewRedis(&storage.RedisStoreConfig{
				Client: global.GB_REDIS.GetClient(),
			}),
			InvalidateWhenUpdate: global.GB_CONFIG.Gorm.InvalidateWhenUpdate,
			CacheTTL:             global.GB_CONFIG.Gorm.CacheTTL,
			CacheMaxItemCnt:      global.GB_CONFIG.Gorm.CacheMaxItemCnt,
			DebugMode:            global.GB_CONFIG.Gorm.DebugMode,
		}
		var err error
		db, err = gormx.New(dbConfig, cacheConfig)
		if err != nil {
			global.GB_LOG.Error("[GORM] connection error", zap.Error(err))
			os.Exit(0)
		}
	} else {
		var err error
		db, err = gormx.New(dbConfig)
		if err != nil {
			global.GB_LOG.Error("[GORM] connection error", zap.Error(err))
			os.Exit(0)
		}
	}
	global.GB_LOG.Info("[GORM] main connection success")
	return db
}

func InitMGorm() map[string]*gorm.DB {
	result := make(map[string]*gorm.DB)
	for _, mdb := range global.GB_CONFIG.Gorm.MultipleDbs {
		if !mdb.Enable {
			continue
		}
		// 獲取主要資料庫logger
		var logger *zap.Logger
		if global.GB_CONFIG.Gorm.LogZap {
			logger = global.GB_DB.Logger.(*zapgorm.Logger).ZapLogger
		}
		// database 設定
		dbConfig := &gormx.Config{
			DBType:       mdb.DBType(),
			Dsn:          mdb.Dsn(),
			MaxIdleConns: mdb.MaxIdleConns,
			MaxOpenConns: mdb.MaxOpenConns,
			LogMode:      global.GB_CONFIG.Gorm.LogMode,
			UseZap:       global.GB_CONFIG.Gorm.LogZap,
			Zap: gormx.Zap{
				LogFolderName: fmt.Sprintf("gorm(%s)", mdb.AliasName),
				Logger:        logger,
			},
		}
		var db *gorm.DB
		if global.GB_CONFIG.Gorm.Cache {
			cacheConfig := &config.CacheConfig{
				CacheLevel: global.GB_CONFIG.Gorm.CacheLevel(),
				CacheStorage: storage.NewRedis(&storage.RedisStoreConfig{
					Client: global.GB_REDIS.GetClient(),
				}),
				InvalidateWhenUpdate: global.GB_CONFIG.Gorm.InvalidateWhenUpdate,
				CacheTTL:             global.GB_CONFIG.Gorm.CacheTTL,
				CacheMaxItemCnt:      global.GB_CONFIG.Gorm.CacheMaxItemCnt,
				DebugMode:            global.GB_CONFIG.Gorm.DebugMode,
			}
			var err error
			db, err = gormx.New(dbConfig, cacheConfig)
			if err != nil {
				global.GB_LOG.Error("[GORM] "+mdb.AliasName+" connection error", zap.Error(err))
				os.Exit(0)
			}
		} else {
			var err error
			db, err = gormx.New(dbConfig)
			if err != nil {
				global.GB_LOG.Error("[GORM] "+mdb.AliasName+" connection error", zap.Error(err))
				os.Exit(0)
			}
		}
		global.GB_LOG.Info("[GORM] " + mdb.AliasName + " connection success")
		result[mdb.AliasName] = db
	}
	return result
}
