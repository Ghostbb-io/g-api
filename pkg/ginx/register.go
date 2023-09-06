package ginx

import "github.com/gin-gonic/gin"

type VersionFunc = func(v int, m ...gin.HandlerFunc) *gin.RouterGroup

type Interface interface {
	Register(VersionFunc)
}

func Register(p []any) {
	for _, g := range p { // 遍歷插件裡路由群組
		g.(Interface).Register(version)
	}
}
