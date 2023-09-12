package system

import (
	v1 "github.com/Ghostbb-io/g-api/app/system/api/v1"
	"github.com/Ghostbb-io/g-api/app/system/model/table"
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"go.uber.org/zap"
	"os"
)

func New() []any {
	err := global.GB_DB.AutoMigrate(
		// 自動建立表
		&table.SysUser{},
		&table.SysRole{},
		&table.SysMenu{},
		&table.SysBtn{},
		&table.SysApi{},
	)
	if err != nil {
		global.GB_LOG.Error("create table error", zap.Error(err))
		os.Exit(0)
	}
	return []any{
		// 回傳Api群組
		&v1.BaseApi{service.BaseService},
		&v1.UserApi{service.UserService},
		&v1.CasbinApi{service.CasbinService},
		&v1.MenuApi{service.MenuService},
		&v1.RoleApi{service.RoleService},
		&v1.CacheApi{service.CacheService},
		&v1.ApiApi{service.ApiService},
	}
}
