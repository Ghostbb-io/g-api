package gormx

import (
	"github.com/Ghostbb-io/g-api/pkg/gormx/zapgorm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var gormx = new(_gormx)

type _gormx struct{}

func (g *_gormx) config(cfg *Config) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.Prefix,
			SingularTable: cfg.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	_default := logInterface(cfg)

	switch cfg.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}

func logInterface(cfg *Config) logger.Interface {
	if cfg.UseZap {
		return zapLogger(cfg.Zap)
	}
	return baseLogger()
}

func zapLogger(zap Zap) logger.Interface {
	if zap.Logger == nil {
		panic("Zap logger not found!")
	}
	l := zapgorm.New(zap.Logger, zapgorm.Config{
		SlowThreshold:             1 * time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		SkipCallerLookup:          false,
	})
	l.SetFolderName(zap.LogFolderName)
	return l
}

func baseLogger() logger.Interface {
	return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             1 * time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	})
}
