package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"go.uber.org/zap"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// LogLayout 日誌layout
type LogLayout struct {
	Time      time.Time
	Metadata  map[string]interface{} // 儲存自定義原數據
	Path      string                 // 訪問路徑
	Method    string                 // 請求方法
	Query     string                 // 攜帶query
	Body      string                 // 攜帶body數據
	IP        string                 // ip地址
	UserAgent string                 // 代理
	Error     string                 // 錯誤
	Cost      time.Duration          // 花費時間
	Source    string                 // 來源
}

type Logger struct {
	// Filter 自定義過濾
	Filter func(c *gin.Context) bool
	// FilterKeyword 關鍵字過濾(key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess 權限處理
	AuthProcess func(c *gin.Context, layout *LogLayout)
	// Print 日誌處理
	Print func(LogLayout)
	// Source 服務唯一指標
	Source string
}

func (l Logger) SetLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		query := c.Request.URL.RawQuery
		var body []byte
		if l.Filter != nil && !l.Filter(c) {
			body, _ = c.GetRawData()
			// 將原body塞回去
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
		cost := time.Since(start)
		layout := LogLayout{
			Time:      time.Now(),
			Path:      path,
			Method:    method,
			Query:     query,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Error:     strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
			Cost:      cost,
			Source:    l.Source,
		}
		if l.Filter != nil && !l.Filter(c) {
			layout.Body = string(body)
		}
		if l.AuthProcess != nil {
			// 處理權限需要訊息
			l.AuthProcess(c, &layout)
		}
		if l.FilterKeyword != nil {
			// 自行判斷key/value
			l.FilterKeyword(&layout)
		}
		// 自行處理日誌
		l.Print(layout)
	}
}

func DefaultLogger() gin.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			// 標準輸出,k8s做收集
			v, _ := json.Marshal(layout)
			fmt.Println(string(v))
		},
		Source: "GB",
	}.SetLoggerMiddleware()
}

func ZapLogger() gin.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			msg := fmt.Sprintf("[Gin] %s \"%s\"", layout.Method, layout.Path)
			if layout.Error != "" {
				global.GB_LOG.Error(msg,
					zap.String("cost", fmt.Sprintf("%.2fs", layout.Cost.Seconds())),
					zap.String("error", layout.Error),
					zap.String("query", layout.Query),
					zap.String("body", layout.Body),
					zap.String("ip", layout.IP),
					zap.String("user-agent", layout.UserAgent),
					zap.String("source", layout.Source),
				)
			} else {
				global.GB_LOG.Info(msg,
					zap.String("cost", fmt.Sprintf("%.2fs", layout.Cost.Seconds())),
					zap.String("query", layout.Query),
					zap.String("body", layout.Body),
					zap.String("ip", layout.IP),
					zap.String("user-agent", layout.UserAgent),
					zap.String("source", layout.Source),
				)
			}
		},
		Source: "GB",
	}.SetLoggerMiddleware()
}
