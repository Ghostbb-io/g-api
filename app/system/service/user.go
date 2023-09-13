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
	RouteList(uint) ([]*model.RouteResponse, error)
	IsUsernameExist(string) bool
	UserListByPage(model.UserPageParams) (model.BasicFetchResult[model.UserItem], error)
	EditUser(uint, model.EditUserRequest) error
	AddUser(model.AddUserRequest) error
	DelUser(uint) error
	SetStatus(uint, model.SetUserStatusRequest) error
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
		Status:   user.Status,
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

// RouteList 根據user id取得route
func (user) RouteList(userID uint) ([]*model.RouteResponse, error) {
	var menus []table.SysMenu
	err := global.GB_DB.Select("distinct d.*").Table("sys_users a").
		Joins("join sys_user_role b on (a.id = b.sys_user_id)").
		Joins("join sys_role_menu c on (b.sys_role_role = c.sys_role_role)").
		Joins("join sys_menus d on (c.sys_menu_id = d.id)").
		Where("a.id = ?", userID).Where("d.status = ?", true).Order("d.sort").Find(&menus).Error
	if err != nil {
		return nil, err
	}

	menuMap := make(map[uint]*model.RouteResponse)
	result := make([]*model.RouteResponse, 0)
	item := make([]*model.RouteResponse, 0)
	for _, menu := range menus {
		temps := &model.RouteResponse{
			ParentID:  menu.ParentID,
			Path:      menu.Path,
			Name:      menu.Name,
			Component: menu.Component,
			Redirect:  menu.Redirect,
			Meta:      menu.Meta,
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
			if menuMap[menu.ParentID].Children == nil {
				menuMap[menu.ParentID].Children = make([]*model.RouteResponse, 0)
			}
			menuMap[menu.ParentID].Children = append(menuMap[menu.ParentID].Children, menu)
		}
	}
	return result, nil
}

// IsUsernameExist 判斷使用者是否存在
func (user) IsUsernameExist(username string) bool {
	return errors.Is(global.GB_DB.Where("username = ?", username).First(&table.SysUser{}).Error, gorm.ErrRecordNotFound)
}

// UserListByPage 根據Page, PageSize獲取使用者列表
func (user) UserListByPage(in model.UserPageParams) (model.BasicFetchResult[model.UserItem], error) {
	status := []bool{true, false}
	if in.Status != "" {
		if in.Status == "true" {
			status = []bool{true}
		} else {
			status = []bool{false}
		}
	}
	var users []table.SysUser
	if err := global.GB_DB.Preload("Roles").
		Where("username like ?", "%"+in.Username+"%").
		Where("nick_name like ?", "%"+in.NickName+"%").
		Where("status in ?", status).Find(&users).Error; err != nil {
		return model.BasicFetchResult[model.UserItem]{}, err
	}
	var result model.BasicFetchResult[model.UserItem]
	for _, user := range users {
		roles := make([]string, 0)
		for _, role := range user.Roles {
			roles = append(roles, role.Role)
		}
		result.Items = append(result.Items, model.UserItem{
			ID:        user.ID,
			Username:  user.Username,
			NickName:  user.NickName,
			RealName:  user.RealName,
			Email:     user.Email,
			Mobile:    user.Mobile,
			Remark:    user.Remark,
			CreatedAt: user.CreatedAt.Format("2006-01-02 - 15:04:05"),
			Status:    user.Status,
			Roles:     roles,
		})
	}
	if err := global.GB_DB.Model(&table.SysUser{}).Count(&result.Total).Error; err != nil {
		return model.BasicFetchResult[model.UserItem]{}, err
	}
	return result, nil
}

// EditUser 編輯使用者
func (user) EditUser(userID uint, in model.EditUserRequest) error {
	var user table.SysUser
	if err := global.GB_DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}
	{
		user.ID = userID
		user.NickName = in.NickName
		user.RealName = in.RealName
		user.Email = in.Email
		user.Mobile = in.Mobile
		user.Remark = in.Remark
	}
	for _, role := range in.Roles {
		user.Roles = append(user.Roles, table.SysRole{Role: role})
	}
	return global.GB_DB.Save(&user).Error
}

// AddUser 新增使用者
func (user) AddUser(in model.AddUserRequest) error {
	var user table.SysUser
	if !errors.Is(global.GB_DB.Where("username = ?", in.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("此使用者已存在")
	}
	{
		user.Username = in.Username
		user.Password = utils.BcryptHash(in.Password)
		user.NickName = in.NickName
		user.RealName = in.RealName
		user.Email = in.Email
		user.Mobile = in.Mobile
		user.Remark = in.Remark
	}
	for _, role := range in.Roles {
		user.Roles = append(user.Roles, table.SysRole{Role: role})
	}
	return global.GB_DB.Create(&user).Error
}

// DelUser 刪除使用者
func (user) DelUser(userID uint) error {
	return global.GB_DB.Where("id = ?", userID).Delete(&table.SysUser{}).Error
}

// SetStatus 更新狀態
func (user) SetStatus(userID uint, in model.SetUserStatusRequest) error {
	return global.GB_DB.Model(&table.SysUser{}).Where("id = ?", userID).Update("status", in.Status).Error
}
