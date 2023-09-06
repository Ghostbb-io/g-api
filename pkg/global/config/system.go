package config

type System struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Https        bool   `mapstructure:"https" json:"https" yaml:"https"`
}
