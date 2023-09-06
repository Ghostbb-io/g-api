package gzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Config struct {
	Level         string // 級別
	Prefix        string // log前綴
	Format        string // console 輸出普通文本，json輸出json格式
	Director      string // log資料夾
	EncodeLevel   string // 編碼級別
	StacktraceKey string // 堆疊名稱
	MaxAge        int    // log留存時間，天為單位
	ShowLine      bool   // 顯示行
	LogInConsole  bool   // 輸出控制台
}

// ZapEncodeLevel 根據 EncodeLevel 返回 zapcore.LevelEncoder
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (c *Config) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case c.EncodeLevel == "LowercaseLevelEncoder": // 小寫編碼器(默認)
		return zapcore.LowercaseLevelEncoder
	case c.EncodeLevel == "LowercaseColorLevelEncoder": // 小寫編碼器帶顏色
		return zapcore.LowercaseColorLevelEncoder
	case c.EncodeLevel == "CapitalLevelEncoder": // 大寫編碼器
		return zapcore.CapitalLevelEncoder
	case c.EncodeLevel == "CapitalColorLevelEncoder": // 大寫編碼器帶顏色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 根據Level轉化為 zapcore.Level
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (c *Config) TransportLevel() zapcore.Level {
	c.Level = strings.ToLower(c.Level)
	switch c.Level {
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

func New(cfg Config) (logger *zap.Logger) {
	cores := Zapx.Set(cfg).GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if cfg.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
