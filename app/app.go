package app

import (
	"github.com/Ghostbb-io/g-api/app/system"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
)

func All() [][]ginx.Interface {
	// 註冊插件
	return [][]ginx.Interface{
		system.New(),
	}
}
