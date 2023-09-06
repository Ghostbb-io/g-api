package gormx

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openPgsql(cfg *Config) (*gorm.DB, error) {
	pgsqlConfig := postgres.Config{
		DSN:                  cfg.Dsn, // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), gormx.config(cfg)); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		return db, nil
	}
}
