package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt簽名
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 過期時間
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 簽發者
}
