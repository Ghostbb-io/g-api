package table

// SysRoleMenu 是 sysRole 和 sysMenu 的連接表
type SysRoleMenu struct {
	SysMenuID   uint   `gorm:"column:sys_menu_id"`
	SysRoleRole string `gorm:"column:sys_role_role"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
