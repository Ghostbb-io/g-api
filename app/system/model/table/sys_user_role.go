package table

// SysUserRole 是 sysUser 和 sysRole 的連接表
type SysUserRole struct {
	SysUserId   uint   `gorm:"column:sys_user_id"`
	SysRoleRole string `gorm:"column:sys_role_role"`
}

func (s *SysUserRole) TableName() string {
	return "sys_user_role"
}
