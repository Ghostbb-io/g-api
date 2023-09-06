package middleware

import (
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/global/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Cors 直接放行所有跨域請求並放行所有 OPTIONS 方法
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, Content-Disposition")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 處理請求
		c.Next()
	}
}

// CorsByRules 按照配置處理跨域請求
func CorsByRules() gin.HandlerFunc {
	// 放行全部
	if global.GB_CONFIG.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(c *gin.Context) {
		whitelist := checkCors(c.GetHeader("origin"))

		// 通過檢查，添加請求頭
		if whitelist != nil {
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		// 嚴格白名單模式且未通過檢查，直接拒絕處理請求
		if whitelist == nil && global.GB_CONFIG.Cors.Mode == "strict-whitelist" && !(c.Request.Method == "GET" && (c.Request.URL.Path == "/health" || strings.Contains(c.Request.URL.Path, "/api/swagger"))) {
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			// 非嚴格白名單模式，無論是否通過檢查均放行所有 OPTIONS 方法
			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}

		// 處理請求
		c.Next()
	}
}

func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.GB_CONFIG.Cors.Whitelist {
		// 遍歷配置中的跨域頭，尋找匹配項
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
