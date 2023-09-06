package config

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 級別
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // log前綴
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // console 輸出普通文本，json輸出json格式
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // log資料夾
	GormDirector  string `mapstructure:"gorm-director" json:"gorm-director"  yaml:"gorm-director"`   // gorm log 資料夾
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 編碼級別
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 堆疊名稱

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // log留存時間，天為單位
	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 顯示行
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 輸出控制台
}

// ZapEncodeLevel 根據 EncodeLevel 返回 zapcore.LevelEncoder
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小寫編碼器(默認)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小寫編碼器帶顏色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大寫編碼器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大寫編碼器帶顏色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根據Level轉化為 zapcore.Level
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
