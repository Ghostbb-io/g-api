package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	versionMap = map[int]*gin.RouterGroup{}
)

func version(v int, m ...gin.HandlerFunc) *gin.RouterGroup {
	// Whether version exists in versionMap
	if _, ok := versionMap[v]; !ok {
		// "/prefix/v1" or "/prefix/v2" ...
		versionMap[v] = _engine.Group(fmt.Sprintf("%s/v%d", _engine.getPrefix(), v))
	}

	// Create a group from an existing version
	g := versionMap[v].Group("")

	// Install middleware
	for _, middleware := range m {
		g.Use(middleware)
	}
	return g
}
