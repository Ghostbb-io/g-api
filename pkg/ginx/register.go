package ginx

import "github.com/gin-gonic/gin"

// VersionFunc v: version, m: middleware
type VersionFunc = func(v int, m ...gin.HandlerFunc) *gin.RouterGroup

type Interface interface {
	Register(VersionFunc)
}

func (e *Engine) Register(p ...Interface) {
	// Traverse routing groups in a plugin
	for _, g := range p {
		g.Register(e.version)
	}
}
