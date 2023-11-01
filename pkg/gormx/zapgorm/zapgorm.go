package zapgorm

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"runtime"
	"strings"
	"time"
)

type ContextFn func(ctx context.Context) []zapcore.Field

type Logger struct {
	ZapLogger *zap.Logger
	Config
	Context  ContextFn
	business zapcore.Field // custom log path
}

type Config struct {
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
	LogLevel                  gormlogger.LogLevel
}

func New(zapLogger *zap.Logger, config Config) *Logger {
	l := &Logger{
		ZapLogger: zapLogger,
		Config:    config,
		Context:   nil,
	}
	gormlogger.Default = l
	return l
}

func (l *Logger) SetFolderName(name string) {
	l.business = zap.String("business", name)
}

func (l *Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	l.LogLevel = level
	if l.business.Key == "" {
		l.business = zap.String("business", "gorm")
	}
	return l
}

func (l *Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.logger(ctx).Sugar().Infof(str, args...)
}

func (l *Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.logger(ctx).Sugar().Warnf(str, args...)
}

func (l *Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.logger(ctx).Sugar().Errorf(str, args...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	logger := l.logger(ctx)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		logger.Error("trace",
			l.business,
			zap.Error(err),
			zap.String("elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		logger.Warn("trace",
			l.business,
			zap.String("elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		logger.Info("trace",
			l.business,
			zap.String("elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	}
}

var (
	gormPackage = "gorm.io"
)

func (l *Logger) logger(ctx context.Context) *zap.Logger {
	logger := l.ZapLogger
	if l.Context != nil {
		fields := l.Context(ctx)
		logger = logger.With(fields...)
	}

	if l.SkipCallerLookup {
		return logger
	}

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		if ok && (!strings.Contains(file, gormPackage) || strings.HasSuffix(file, "_test.go")) {
			return logger.WithOptions(zap.AddCallerSkip(i - 1))
		}
	}
	return logger
}
