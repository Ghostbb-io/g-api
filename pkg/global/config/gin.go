package config

type Gin struct {
	LogZap bool `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"` // 是否通過zap寫入日志文件
}
