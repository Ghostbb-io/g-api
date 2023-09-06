package config

import (
	"github.com/Ghostbb-io/g-api/pkg/gormx"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/config"
)

type Gorm struct {
	Type        string `mapstructure:"type" json:"type" yaml:"type"`
	DBConfig    `yaml:",inline" mapstructure:",squash"`
	CacheConfig `yaml:",inline" mapstructure:",squash"`
	LogMode     string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	LogZap      bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
	MultipleDbs []MDB  `mapstructure:"multiple-dbs" json:"multiple-dbs" yaml:"multiple-dbs"`
}

type DBConfig struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高級配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 資料庫名稱
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 資料庫使用者名稱
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 資料庫密碼
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空閒中的最大連線數
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打開到資料庫的最大連線數
	Cache        bool   `mapstructure:"cache" json:"cache" yaml:"cache"`
}

type CacheConfig struct {
	CacheLevel           string `mapstructure:"cache-level" json:"cache-level" yaml:"cache-level"`
	InvalidateWhenUpdate bool   `mapstructure:"invalidate-when-update" json:"invalidate-when-update" yaml:"invalidate-when-update"`
	CacheTTL             int64  `mapstructure:"cache-ttl" json:"cache-ttl" yaml:"cache-ttl"`
	CacheMaxItemCnt      int64  `mapstructure:"cache-max-item-cnt" json:"cache-max-item-cnt" yaml:"cache-max-item-cnt"`
	DebugMode            bool   `mapstructure:"debug-mode" json:"debug-mode" yaml:"debug-mode"`
}

type MDB struct {
	Enable    bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	DBConfig  `yaml:",inline" mapstructure:",squash"`
}

func (g *Gorm) Dsn() string {
	switch g.Type {
	case "mssql":
		return "sqlserver://" + g.Username + ":" + g.Password + "@" + g.Path + ":" + g.Port + "?database=" + g.Dbname + "&encrypt=disable"
	case "mysql":
		return g.Username + ":" + g.Password + "@tcp(" + g.Path + ":" + g.Port + ")/" + g.Dbname + "?" + g.Config
	case "pgsql":
		return "host=" + g.Path + " user=" + g.Username + " password=" + g.Password + " dbname=" + g.Dbname + " port=" + g.Port + " " + g.Config
	default:
		return ""
	}
}

func (g *Gorm) DBType() gormx.DbType {
	switch g.Type {
	case "mssql":
		return gormx.Mssql
	case "mysql":
		return gormx.Mysql
	case "pgsql":
		return gormx.Pgsql
	default:
		return ""
	}
}

func (g *Gorm) CacheLevel() config.CacheLevel {
	switch g.CacheConfig.CacheLevel {
	case "CacheLevelOnlyPrimary":
		return config.CacheLevelOnlyPrimary
	case "CacheLevelOnlySearch":
		return config.CacheLevelOnlySearch
	case "CacheLevelAll":
		return config.CacheLevelAll
	default:
		return config.CacheLevelAll
	}
}

func (m *MDB) DBType() gormx.DbType {
	switch m.Type {
	case "mssql":
		return gormx.Mssql
	case "mysql":
		return gormx.Mysql
	case "pgsql":
		return gormx.Pgsql
	default:
		return ""
	}
}

func (m *MDB) Dsn() string {
	switch m.Type {
	case "mssql":
		return "sqlserver://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "?database=" + m.Dbname + "&encrypt=disable"
	case "mysql":
		return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
	case "pgsql":
		return "host=" + m.Path + " user=" + m.Username + " password=" + m.Password + " dbname=" + m.Dbname + " port=" + m.Port + " " + m.Config
	default:
		return ""
	}
}
