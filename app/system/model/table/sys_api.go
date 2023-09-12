package table

import "github.com/Ghostbb-io/g-api/pkg/global"

type SysApi struct {
	global.GB_MODEL
	Path   string `json:"path" gorm:"comment:api路徑"`
	Desc   string `json:"desc" gorm:"comment:api中文描述"`
	Group  string `json:"group" gorm:"comment:api組"`
	Method string `json:"method" gorm:"default:GET;comment:方法"`
}

func (SysApi) TableName() string {
	return "sys_apis"
}
