package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	versionMap = map[int]*gin.RouterGroup{}
)

func version(v int, m ...gin.HandlerFunc) *gin.RouterGroup {
	if _, ok := versionMap[v]; !ok {
		versionMap[v] = _engine.Group(fmt.Sprintf("%s/v%d", _engine.prefix, v))
	}
	g := versionMap[v].Group("")
	for _, middleware := range m {
		g.Use(middleware)
	}
	return g
}
