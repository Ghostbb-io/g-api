package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪個資料庫
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 伺服器網址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密碼
}
