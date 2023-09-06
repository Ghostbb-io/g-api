package ginx

import "github.com/gin-gonic/gin"

var _engine = new(Engine)

type Engine struct {
	*gin.Engine
	prefix string
}

func (e *Engine) SetPrefix(prefix string) {
	if prefix[0:1] != "/" {
		prefix = "/" + prefix
	}
	e.prefix = prefix
}

func (e *Engine) getPrefix() string {
	return e.prefix
}

func Default() *Engine {
	ginEngine := gin.Default()
	_engine.Engine = ginEngine
	return _engine
}

func GetEngine() *Engine {
	return _engine
}
