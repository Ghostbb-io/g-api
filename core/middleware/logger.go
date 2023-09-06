package middleware

import (
	"fmt"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 先往下執行商業邏輯處理
		c.Next()

		// 計算執行多久
		cost := time.Since(start)
		global.GB_LOG.Info(fmt.Sprintf("[Gin] \"%s\"", path),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
