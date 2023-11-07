package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Default() *Engine {
	return &Engine{
		Engine:     gin.Default(),
		versionMap: make(map[int]*gin.RouterGroup),
	}
}

type Engine struct {
	// origin engine
	*gin.Engine

	// route prefix
	prefix string

	// version group
	versionMap map[int]*gin.RouterGroup
}

// SetPrefix is the prefix for setting the route
func (e *Engine) SetPrefix(prefix string) {
	if prefix[0:1] != "/" {
		prefix = "/" + prefix
	}
	e.prefix = prefix
}

// version function
func (e *Engine) version(v int, m ...gin.HandlerFunc) *gin.RouterGroup {
	// Whether version exists in versionMap
	if _, ok := e.versionMap[v]; !ok {
		// "/{prefix}/v1" or "/{prefix}/v2" ...
		e.versionMap[v] = e.Group(fmt.Sprintf("%s/v%d", e.prefix, v))
	}

	// Create a group from an existing version
	g := e.versionMap[v].Group("")

	// Install middleware
	for _, middleware := range m {
		g.Use(middleware)
	}
	return g
}
