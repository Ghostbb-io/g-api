package model

import "github.com/Ghostbb-io/g-api/app/system/model/table"

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
	Status   bool       `json:"status"`
	Remark   string     `json:"remark"`
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

type RouteResponse struct {
	ParentID  uint             `json:"-"`
	Path      string           `json:"path"`
	Name      string           `json:"name"`
	Component string           `json:"component"`
	Redirect  string           `json:"redirect"`
	Meta      table.Meta       `json:"meta"`
	Children  []*RouteResponse `json:"children"`
}

type UserPageParams struct {
	BasicPageParams
	UserParams
}

type UserParams struct {
	Username string `form:"username"`
	NickName string `form:"nickName"`
	Status   string `form:"status"`
}

type UserItem struct {
	ID        uint     `json:"id"`
	Username  string   `json:"username"`
	NickName  string   `json:"nickName"`
	RealName  string   `json:"realName"`
	Email     string   `json:"email"`
	Mobile    string   `json:"mobile"`
	Remark    string   `json:"remark"`
	CreatedAt string   `json:"createdAt"`
	Status    bool     `json:"status"`
	Roles     []string `json:"roles"`
}

type EditUserRequest struct {
	UserItem
}

type AddUserRequest struct {
	UserItem
	Password string `json:"password"`
}

type SetUserStatusRequest struct {
	Status bool `json:"status"`
}

type ResetPassRequest struct {
	Password string `json:"password"`
}
