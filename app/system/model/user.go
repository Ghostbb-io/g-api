package model

import "github.com/google/uuid"

type UserInfoResponse struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	NickName string    `json:"nickName"`
	Email    string    `json:"email"`
	Remark   string    `json:"remark"`
	Roles    []string  `json:"roles"`
}

type ChangePassRequest struct {
	OldPass string `json:"oldPassword"`
	NewPass string `json:"newPassword"`
}

type RolesRequest struct {
	Roles []string `json:"roles"`
}
