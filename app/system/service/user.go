package service

import (
	"errors"
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/model/table"
	"github.com/Ghostbb-io/g-api/pkg/casbinx"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/jwtx/token"
	"github.com/Ghostbb-io/g-api/pkg/utils"
	"gorm.io/gorm"
)

var UserService User = new(user)

type User interface {
	UserInfo(uint) (model.UserInfoResponse, error)
	ChangePassword(uint, model.ChangePassRequest, string) error
	UpdateRoles(uint, model.RolesRequest) error
	PermList(uint) ([]string, error)
	MenuList(uint) ([]*model.MenuListResponse, error)
}

type user struct{}

// UserInfo 獲取使用者訊息
func (user) UserInfo(userID uint) (model.UserInfoResponse, error) {
	var user table.SysUser
	if err := global.GB_DB.Where("id = ?", userID).Preload("Roles").First(&user).Error; err != nil {
		return model.UserInfoResponse{}, err
	}
	roles := make([]model.RoleInfo, 0)
	for _, role := range user.Roles {
		roles = append(roles, model.RoleInfo{Role: role.Role, RoleName: role.RoleName})
	}
	return model.UserInfoResponse{
		UserID:   user.ID,
		UUID:     user.UUID.String(),
		Username: user.Username,
		NickName: user.NickName,
		RealName: user.RealName,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Avatar:   user.Avatar,
		Desc:     user.Desc,
		Roles:    roles,
	}, nil
}

// ChangePassword 更改密碼
func (user) ChangePassword(userID uint, in model.ChangePassRequest, accToken string) error {
	var user table.SysUser
	if err := global.GB_DB.First(&user, userID).Error; err != nil {
		return err
	}
	// 原密碼比對
	if !utils.BcryptCheck(user.Password, in.OldPass) {
		return errors.New("password wrong")
	}
	// 密碼更新
	user.Password = utils.BcryptHash(in.NewPass)
	if err := global.GB_DB.Save(&user).Error; err != nil {
		return err
	}
	// 需要重新登入, 將access token 列入黑名單
	if err := token.SetBlack(global.GB_REDIS.GetClient(), accToken); err != nil {
		return err
	}
	return nil
}

// UpdateRoles 更新使用者角色
func (user) UpdateRoles(userID uint, roles model.RolesRequest) error {
	return global.GB_DB.Transaction(func(tx *gorm.DB) error {
		// 刪除使用者所有角色
		if err := tx.Delete(&table.SysUserRole{}, "sys_user_id", userID).Error; err != nil {
			return err
		}
		// 將新的角色存入
		userRole := make([]table.SysUserRole, 0)
		for _, role := range roles.Roles {
			userRole = append(userRole, table.SysUserRole{SysUserId: userID, SysRoleRole: role})
		}
		if err := tx.Create(&userRole).Error; err != nil {
			return err
		}
		// 更新casbin
		cb, _ := casbinx.New(global.GB_DB)
		if err := cb.UpdateRolesForUser(userID, roles.Roles); err != nil {
			return err
		}
		return nil
	})
}

// PermList 獲取按鈕權限列表
func (user) PermList(userID uint) ([]string, error) {
	var perms []string
	err := global.GB_DB.Select("sys_btn_permission").Table("sys_users a").
		Joins("join sys_user_role b on (a.id = b.sys_user_id)").
		Joins("join sys_role_btn c on (b.sys_role_role = c.sys_role_role)").
		Where("a.id = ?", userID).Find(&perms).Error
	if err != nil {
		return nil, err
	}
	return perms, nil
}

// MenuList 根據user id取得menu
func (user) MenuList(userID uint) ([]*model.MenuListResponse, error) {
	var menus []table.SysMenu
	err := global.GB_DB.Select("distinct d.*").Table("sys_users a").
		Joins("join sys_user_role b on (a.id = b.sys_user_id)").
		Joins("join sys_role_menu c on (b.sys_role_role = c.sys_role_role)").
		Joins("join sys_menus d on (c.sys_menu_id = d.id)").
		Where("a.id = ?", userID).Order("d.sort").Find(&menus).Error
	if err != nil {
		return nil, err
	}

	menuMap := make(map[uint]*model.MenuListResponse)
	result := make([]*model.MenuListResponse, 0)
	item := make([]*model.MenuListResponse, 0)
	for _, menu := range menus {
		temps := &model.MenuListResponse{
			ParentID:  menu.ParentID,
			Path:      menu.Path,
			Name:      menu.Path,
			Component: menu.Component,
			Redirect:  menu.Redirect,
			Meta:      menu.Meta,
			Children:  make([]*model.MenuListResponse, 0),
		}
		if menu.ParentID == 0 {
			result = append(result, temps)
		} else {
			item = append(item, temps)
		}
		menuMap[menu.ID] = temps
	}

	// 組裝
	for _, menu := range item {
		// 判斷父親menu是否存在
		if _, ok := menuMap[menu.ParentID]; ok {
			menuMap[menu.ParentID].Children = append(menuMap[menu.ParentID].Children, menu)
		}
	}
	return result, nil
}
