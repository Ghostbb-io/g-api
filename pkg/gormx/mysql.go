package gormx

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openMysql(cfg *Config) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN:                       cfg.Dsn,
		DefaultStringSize:         191,   // string 類型默認長度
		SkipInitializeWithVersion: false, // 根據版本自動配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormx.config(cfg)); err != nil {
		return nil, err
	} else {
		if cfg.Engine != "" {
			db.InstanceSet("gorm:table_options", "ENGINE="+cfg.Engine)
		} else {
			db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		}
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		return db, nil
	}
}
