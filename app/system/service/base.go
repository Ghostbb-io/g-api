package service

import (
	"errors"
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/model/table"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/jwtx"
	"github.com/Ghostbb-io/g-api/pkg/jwtx/claims"
	"github.com/Ghostbb-io/g-api/pkg/jwtx/token"
	"github.com/Ghostbb-io/g-api/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var BaseService Base = new(base)

type Base interface {
	Login(model.LoginRequest) (model.LoginResponse, error)
	RegisterUser(model.RegisterRequest) error
	Logout(uuid.UUID, string) error
}

type base struct{}

// Login 登入
func (base) Login(in model.LoginRequest) (model.LoginResponse, error) {
	var user table.SysUser
	if err := global.GB_DB.Preload("Roles").First(&user, "username = ?", in.Username).Error; err != nil {
		return model.LoginResponse{}, err
	}
	if !utils.BcryptCheck(user.Password, in.Password) {
		return model.LoginResponse{}, errors.New("account or password incorrect")
	}

	// 建立Token
	expiresTime, _ := jwtx.ParseDuration(global.GB_CONFIG.JWT.ExpiresTime)
	roles := make([]string, 0)
	for _, role := range user.Roles {
		roles = append(roles, role.Role)
	}
	jwt := jwtx.New(
		global.GB_CONFIG.JWT.SigningKey,
		global.GB_CONFIG.JWT.Issuer,
		expiresTime,
	)
	accessToken, RefreshToken, err := jwt.CreateToken(claims.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Username: user.Username,
		Roles:    roles,
	})
	if err != nil {
		return model.LoginResponse{}, errors.New("create token error")
	}
	// 將refresh token放進緩存
	if err = token.SetRefreshToken(global.GB_REDIS.GetClient(), user.UUID.String(), RefreshToken, expiresTime); err != nil {
		return model.LoginResponse{}, errors.New("cache error")
	}

	return model.LoginResponse{Token: accessToken}, nil
}

// RegisterUser 註冊使用者
func (base) RegisterUser(in model.RegisterRequest) error {
	var user table.SysUser
	// 判斷使用者是否存在
	if !errors.Is(global.GB_DB.First(&user, "username = ?", in.Username).Error, gorm.ErrRecordNotFound) {
		return errors.New("user already exists")
	}
	{
		user.UUID, _ = uuid.NewUUID()
		user.Username = in.Username
		user.Password = utils.BcryptHash(in.Password)
		user.NickName = in.NickName
		user.Email = in.Email
		user.Roles = append([]table.SysRole{}, table.SysRole{Role: "user"})
	}
	if err := global.GB_DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Logout 登出
func (base) Logout(uuid uuid.UUID, accToken string) error {
	// 刪除refresh token
	if err := token.DelRefreshToken(global.GB_REDIS.GetClient(), uuid.String()); err != nil {
		return err
	}
	// 將access token 列入黑名單
	if err := token.SetBlack(global.GB_REDIS.GetClient(), accToken); err != nil {
		return err
	}
	return nil
}
