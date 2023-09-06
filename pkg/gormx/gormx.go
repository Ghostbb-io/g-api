package gormx

import (
	"errors"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/cache"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/config"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gzap"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	Mysql DbType = "mysql"
	Pgsql DbType = "pgsql"
	Mssql DbType = "mssql"
)

var (
	DBTypeErr = errors.New("database type error")
)

type DbType string

type Config struct {
	DBType       DbType
	Dsn          string
	MaxIdleConns int
	MaxOpenConns int
	Prefix       string
	Singular     bool
	LogMode      string
	UseZap       bool
	Engine       string
	Zap
}

type Zap struct {
	Logger *zap.Logger
	Config gzap.Config // if Logger is null
}

func New(cfg *Config, _cache ...*config.CacheConfig) (db *gorm.DB, err error) {
	switch cfg.DBType {
	case Mssql:
		db, err = openMssql(cfg)
	case Mysql:
		db, err = openMysql(cfg)
	case Pgsql:
		db, err = openPgsql(cfg)
	default:
		db, err = nil, DBTypeErr
	}

	// 緩存設定
	if len(_cache) != 0 {
		var gormCache cache.Cache
		gormCache, err = cache.New(_cache[0])
		if err != nil {
			return nil, err
		}
		err = db.Use(gormCache)
		if err != nil {
			return nil, err
		}
	}

	return db, err
}
