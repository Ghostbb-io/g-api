package core

import (
	"fmt"
	"github.com/Ghostbb-io/g-api/pkg/global"

	"github.com/spf13/viper"
)

// InitViper 初始化Viper，把config.yaml的設定檔放進global.GB_CONFIG裡面。
func InitViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	if err = v.Unmarshal(&global.GB_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
