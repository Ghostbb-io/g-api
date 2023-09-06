package gormx

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func openMssql(cfg *Config) (*gorm.DB, error) {
	mssqlConfig := sqlserver.Config{
		DSN:               cfg.Dsn, // DSN data source name
		DefaultStringSize: 191,     // string 類型字段默認長度
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), gormx.config(cfg)); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		return db, nil
	}
}
