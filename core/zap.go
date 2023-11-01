package core

import (
	"github.com/Ghostbb-io/g-api/core/log"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitZap 獲取 zap.Logger
func InitZap() (logger *zap.Logger) {
	// 判斷 config.yaml 設定的 Director 資料夾是否存在
	utils.MkdirIfNotExist(global.GB_CONFIG.Zap.Director)

	cores := log.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GB_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger)
	logger.Info("----------------------------------------------------------------------------------")
	return logger
}
