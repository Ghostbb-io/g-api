package service

import (
	"errors"
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/model/table"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"gorm.io/gorm"
)

var RoleService Role = new(role)

type Role interface {
	RoleList() ([]model.RoleListResponse, error)
	RoleListByPage(model.RolePageParams) (model.BasicFetchResult[model.RoleItem], error)
	SetStatus(string, model.SetStatusRequest) error
	Update(string, model.EditRoleRequest) error
	AddRole(model.AddRoleRequest) error
	DelRole(string) error
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

// RoleListByPage 根據page, pageSize獲取角色列表
func (role) RoleListByPage(in model.RolePageParams) (model.BasicFetchResult[model.RoleItem], error) {
	tx := global.GB_DB.Model(&table.SysRole{}).Where("role_name like ?", "%"+in.RoleName+"%").Limit(in.PageSize).Offset((in.Page - 1) * in.PageSize)
	if in.Status != "" {
		if in.Status == "true" {
			tx.Where("status = ?", true)
		} else {
			tx.Where("status = ?", false)
		}
	}
	var roles []table.SysRole
	if err := tx.Preload("Menus").Find(&roles).Error; err != nil {
		return model.BasicFetchResult[model.RoleItem]{}, err
	}

	var result model.BasicFetchResult[model.RoleItem]
	for _, role := range roles {
		menusID := make([]uint, 0)
		for _, menu := range role.Menus {
			if menu.Component != "LAYOUT" && menu.Component != "" {
				menusID = append(menusID, menu.ID)
			}
		}
		result.Items = append(result.Items, model.RoleItem{
			Role:      role.Role,
			RoleName:  role.RoleName,
			Status:    role.Status,
			Remark:    role.Remark,
			CreatedAt: role.CreatedAt.Format("2006-01-02 - 15:04:05"),
			Menu:      menusID,
		})
	}
	if err := global.GB_DB.Model(&table.SysRole{}).Count(&result.Total).Error; err != nil {
		return model.BasicFetchResult[model.RoleItem]{}, err
	}
	return result, nil
}

// SetStatus 設定狀態
func (role) SetStatus(inRole string, in model.SetStatusRequest) error {
	var role table.SysRole
	if err := global.GB_DB.Where("role = ?", inRole).First(&role).Error; err != nil {
		return err
	}
	role.Status = in.Status
	if err := global.GB_DB.Save(&role).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新角色
func (role) Update(r string, in model.EditRoleRequest) error {
	var role table.SysRole
	if err := global.GB_DB.Where("role = ?", r).First(&role).Error; err != nil {
		return err
	}
	{
		role.RoleName = in.RoleName
		role.Remark = in.Remark
		role.Status = in.Status
	}
	return global.GB_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&role).Error; err != nil {
			return err
		}

		// menu更新
		sql := `with tree as (
					select id,
						   parent_id,
						   title, convert(nvarchar(150),CONCAT('.',id,'.')) as path
					from sys_menus where parent_id = 0
					union all
					select data.id,
						   data.parent_id,
						   data.title,
						   Convert(nvarchar(150),CONCAT(tree.Path , '-','.',data.id,'.')) as Path
					from sys_menus data
					join tree tree on data.parent_id = tree.id and tree.path not like '%' + CONCAT('.',data.Id,'.') + '%'
				)
				select distinct replace(s.value, '.', '') as string from tree t
				cross apply STRING_SPLIT(t.path, '-') s
				where t.id in ?`
		var menus []uint
		var roleMenu []table.SysRoleMenu
		if err := tx.Raw(sql, in.Menu).Find(&menus).Error; err != nil {
			return err
		}
		if err := tx.Where("sys_role_role = ?", in.Role).Delete(&table.SysRoleMenu{}).Error; err != nil {
			return err
		}
		for _, m := range menus {
			roleMenu = append(roleMenu, table.SysRoleMenu{
				SysRoleRole: in.Role,
				SysMenuID:   m,
			})
		}
		if err := tx.Create(&roleMenu).Error; err != nil {
			return err
		}
		return nil
	})
}

// AddRole 新增角色
func (role) AddRole(in model.AddRoleRequest) error {
	var role table.SysRole
	if !errors.Is(global.GB_DB.Where("role = ?", in.Role).First(&role).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色已存在")
	}
	{
		role.Role = in.Role
		role.RoleName = in.RoleName
		role.Status = in.Status
		role.Remark = in.Remark
	}
	for _, menu := range in.Menu {
		role.Menus = append(role.Menus, table.SysMenu{
			GB_MODEL: global.GB_MODEL{ID: menu},
		})
	}
	return global.GB_DB.Create(&role).Error
}

// DelRole 刪除角色
func (role) DelRole(role string) error {
	return global.GB_DB.Where("role = ?", role).Delete(&table.SysRole{}).Error
}
