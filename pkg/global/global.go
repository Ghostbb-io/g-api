package global

import (
	"github.com/Ghostbb-io/g-api/pkg/global/config"
	"github.com/Ghostbb-io/g-api/pkg/redisx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GB_DB     *gorm.DB
	GB_DBS    map[string]*gorm.DB
	GB_CONFIG config.Server
	GB_LOG    *zap.Logger
	GB_VP     *viper.Viper
	GB_REDIS  *redisx.Redis
	GB_SF     = &singleflight.Group{}
)
