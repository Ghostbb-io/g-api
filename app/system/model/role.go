package model

type AddRoleRequest struct {
	RoleItem
}

type RoleListResponse struct {
	Role     string `json:"role"`
	RoleName string `json:"roleName"`
	Status   bool   `json:"status"`
}

type SetStatusRequest struct {
	Status bool `json:"status"`
}

type RoleItem struct {
	RoleName  string   `json:"roleName"`
	Role      string   `json:"role"`
	Status    bool     `json:"status"`
	Remark    string   `json:"remark"`
	CreatedAt string   `json:"createdAt"`
	Menu      []uint   `json:"menu"`
	Api       []string `json:"api"`
}

type RoleParams struct {
	RoleName string `form:"roleName"`
	Status   string `form:"status"`
}

type RolePageParams struct {
	BasicPageParams
	RoleParams
}

type EditRoleRequest struct {
	RoleItem
}
