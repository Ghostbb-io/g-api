package model

type RoleListResponse struct {
	Role     string `json:"role"`
	RoleName string `json:"roleName"`
	Status   bool   `json:"status"`
}
