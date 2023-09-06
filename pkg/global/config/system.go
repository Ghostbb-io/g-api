package config

type System struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Https        bool   `mapstructure:"https" json:"https" yaml:"https"`
	LimitTimeIP  int    `mapstructure:"limit-time-ip" json:"limit-time-ip" yaml:"limit-time-ip"`
	LimitCountIP int    `mapstructure:"limit-count-ip" json:"limit-count-ip" yaml:"limit-count-ip"`
}
