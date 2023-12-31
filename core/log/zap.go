package log

import (
	"github.com/Ghostbb-io/g-api/pkg/global"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap = new(_zap)

type _zap struct{}

// GetZapCores 根據配置文件的Level獲取 []zapcore.Core
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.GB_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetEncoder 獲取 zapcore.Encoder
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.GB_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig 獲取 zapcore.EncoderConfig
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GB_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.GB_CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// CustomTimeEncoder 自定義log輸出時間格式
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.GB_CONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// GetEncoderCore 獲取Encoder的 zapcore.Core
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := z.GetWriteSyncer(l.String()) // 日誌分割
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// GetLevelPriority 根據 zapcore.Level 獲取 zap.LevelEnablerFunc
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
func (z *_zap) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := NewCutter(global.GB_CONFIG.Zap.Director, level, WithCutterFormat("2006-01-02"))
	if global.GB_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
