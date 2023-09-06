package table

import (
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/google/uuid"
)

type SysUser struct {
	global.GB_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:使用者UUID"`
	Username string    `json:"username" gorm:"index;comment:帳號"`
	Password string    `json:"-" gorm:"comment:使用者登入密碼"`
	NickName string    `json:"nickName" gorm:"comment:使用者暱稱"`
	Email    string    `json:"email"  gorm:"comment:使用者信箱"`
	Remark   string    `json:"remark" gorm:"comment:備註"`
	Roles    []SysRole `json:"roles" gorm:"many2many:sys_user_role;"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
