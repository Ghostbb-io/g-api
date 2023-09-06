package service

import (
	"github.com/Ghostbb-io/g-api/pkg/global"
	gormCache "github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/cache"
	"go.uber.org/zap"
)

var CacheService Cache = new(cache)

type Cache interface {
	ClearDbCache() error
}

type cache struct{}

// ClearDbCache 清除資料庫緩存
func (cache) ClearDbCache() error {
	gc := global.GB_DB.Config.Plugins["gormcache"].(gormCache.Cache)
	if err := gc.ResetCache(); err != nil {
		global.GB_LOG.Error("[GORM] clear cache error", zap.Error(err))
		return err
	}
	global.GB_LOG.Info("[GORM] clear cache")
	return nil
}
