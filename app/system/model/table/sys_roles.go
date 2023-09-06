package table

import (
	"gorm.io/gorm"
	"time"
)

type SysRole struct {
	Role      string `json:"role" gorm:"not null;primarykey;comment:角色"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	RoleName  string         `json:"roleName" gorm:"comment:角色展示值"`
	Status    bool           `json:"-" gorm:"default:true;comment:是否啟用該角色"`
	DataScope int            `json:"dataScope" gorm:"default:3;comment:數據權限 0全部數據 1部門及子部門數據 2本部門數據 3本人數據"`
	Remark    string         `json:"-" gorm:"comment:備註"`
	Users     []SysUser      `json:"-" gorm:"many2many:sys_user_role;"`
}

func (SysRole) TableName() string {
	return "sys_roles"
}
