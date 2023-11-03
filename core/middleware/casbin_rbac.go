package middleware

import (
	"github.com/Ghostbb-io/g-api/pkg/casbinx"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// Casbin 攔截器
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cb, _ := casbinx.New(global.GB_DB)
		e := cb.GetSyncedCachedEnforcer()
		userId := ginx.GetUserID(c)

		// 放行有root角色的使用者
		if isRoot, _ := cb.HasRoot(userId); !isRoot {
			obj := strings.TrimPrefix(c.Request.URL.Path, global.GB_CONFIG.System.RouterPrefix)
			act := c.Request.Method
			sub := strconv.Itoa(int(userId))

			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(c, gin.H{}, "no authority")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
