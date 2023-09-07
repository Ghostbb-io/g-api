package table

import (
	"gorm.io/gorm"
	"time"
)

type SysBtn struct {
	Permission     string `json:"permission" gorm:"not null;primarykey;comment:權限標示"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	PermissionName string         `json:"permissionName" gorm:"comment:展示值"`
	Remark         string         `json:"remark" gorm:"comment:備註"`
	Status         bool           `json:"status" gorm:"default:true;comment:是否啟用"`
	Roles          []SysRole      `json:"-" gorm:"many2many:sys_role_btn;"`
}

func (SysBtn) TableName() string {
	return "sys_btns"
}
