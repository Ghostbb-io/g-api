package cache

import (
	"fmt"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/config"
)

func New(cacheConfig *config.CacheConfig) (Cache, error) {
	if cacheConfig == nil {
		return nil, fmt.Errorf("you pass a nil config")
	}
	cache := &Gorm2Cache{
		Config: cacheConfig,
		stats:  &stats{},
	}
	err := cache.Init()
	if err != nil {
		return nil, err
	}
	return cache, nil
}
