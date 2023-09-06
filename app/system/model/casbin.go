package model

type ApiResponse struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type ApiRequest struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type AddRoleForUserRequest struct {
	UserID uint     `json:"userID"`
	Roles  []string `json:"role"`
}
