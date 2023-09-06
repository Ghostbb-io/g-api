package app

import (
	"github.com/Ghostbb-io/g-api/app/system"
)

func All() [][]any {
	// 註冊插件
	return [][]any{
		system.New(),
	}
}
