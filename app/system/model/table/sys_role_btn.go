package table

// SysRoleBtn 是 sysRole 和 sysBtn 的連接表
type SysRoleBtn struct {
	SysBtnPermission string `gorm:"column:sys_btn_permission"`
	SysRoleRole      string `gorm:"column:sys_role_role"`
}

func (SysRoleBtn) TableName() string {
	return "sys_role_btn"
}
