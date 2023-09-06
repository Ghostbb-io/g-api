package gzap

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var Zapx = new(_zap)

type _zap struct {
	config Config
}

func (z *_zap) Set(c Config) *_zap {
	return &_zap{c}
}

// GetZapCores 根據配置文件的Level獲取 []zapcore.Core
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.config.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetEncoder 獲取 zapcore.Encoder
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *_zap) GetEncoder() zapcore.Encoder {
	if z.config.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig 獲取 zapcore.EncoderConfig
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  z.config.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    z.config.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// CustomTimeEncoder 自定義log輸出時間格式
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(z.config.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// GetEncoderCore 獲取Encoder的 zapcore.Core
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := z.GetWriteSyncer(l.String()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// GetLevelPriority 根據 zapcore.Level 獲取 zap.LevelEnablerFunc
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 測試級別
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日誌級別
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告級別
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 錯誤級別
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic級別
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic級別
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 終止級別
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 測試級別
			return level == zap.DebugLevel
		}
	}
}

// GetWriteSyncer 獲取 zapcore.WriteSyncer
// Author [YuWeiGhostbb](https://github.com/YuWeiGhostbb)
func (z *_zap) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(z.config.Director, "%Y-%m-%d", level+".log"),             // 路徑
		rotatelogs.WithClock(rotatelogs.Local),                             // 設定時區
		rotatelogs.WithMaxAge(time.Duration(z.config.MaxAge)*24*time.Hour), // log留存時間
		rotatelogs.WithRotationTime(time.Hour*24),                          // 根據天分割
	)
	if z.config.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
