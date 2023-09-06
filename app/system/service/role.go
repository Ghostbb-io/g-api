package service

import (
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/model/table"
	"github.com/Ghostbb-io/g-api/pkg/global"
)

var RoleService Role = new(role)

type Role interface {
	RoleList() ([]model.RoleListResponse, error)
}

type role struct{}

// RoleList 獲取角色列表
func (role) RoleList() ([]model.RoleListResponse, error) {
	var roles []table.SysRole
	if err := global.GB_DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	result := make([]model.RoleListResponse, 0)
	for _, role := range roles {
		result = append(result, model.RoleListResponse{Role: role.Role, RoleName: role.RoleName})
	}
	return result, nil
}
