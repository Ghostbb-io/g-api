package model

type UserInfoResponse struct {
	UserID   uint       `json:"userId"`
	UUID     string     `json:"uuid"`
	Username string     `json:"username"`
	NickName string     `json:"nickName"`
	RealName string     `json:"realName"`
	Email    string     `json:"email"`
	Mobile   string     `json:"mobile"`
	Avatar   string     `json:"avatar"`
	Desc     string     `json:"desc"`
	Roles    []RoleInfo `json:"roles"`
}

type ChangePassRequest struct {
	OldPass string `json:"oldPassword"`
	NewPass string `json:"newPassword"`
}

type RolesRequest struct {
	Roles []string `json:"roles"`
}

type RoleInfo struct {
	Role     string `json:"role"`
	RoleName string `json:"roleName"`
}
