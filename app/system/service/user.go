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
}

type user struct{}

// UserInfo 獲取使用者訊息
func (user) UserInfo(userID uint) (model.UserInfoResponse, error) {
	var user table.SysUser
	if err := global.GB_DB.Where("id = ?", userID).Preload("Roles").First(&user).Error; err != nil {
		return model.UserInfoResponse{}, err
	}
	roles := make([]string, 0)
	for _, role := range user.Roles {
		roles = append(roles, role.Role)
	}
	return model.UserInfoResponse{
		UUID:     user.UUID,
		Username: user.Username,
		NickName: user.NickName,
		Email:    user.Email,
		Remark:   user.Remark,
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
