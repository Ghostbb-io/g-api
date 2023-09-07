package casbinx

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type ApiInfo struct {
	Method string // 方法
	Path   string // 路由
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

type Casbinx struct {
	db                   *gorm.DB
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
}

func New(db *gorm.DB) (Casbinx, error) {
	var err error = nil
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(db)
		if err != nil {
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && keyMatch2(r.obj,p.obj) && r.act == p.act || g(r.sub, "root")
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	if err != nil {
		return Casbinx{}, err
	}
	return Casbinx{
		db:                   db,
		syncedCachedEnforcer: syncedCachedEnforcer,
	}, nil
}

// GetSyncedCachedEnforcer 暴露原生syncedCachedEnforcer
func (c *Casbinx) GetSyncedCachedEnforcer() *casbin.SyncedCachedEnforcer {
	return c.syncedCachedEnforcer
}

// GetRolesForUser 獲取使用者所有角色
func (c *Casbinx) GetRolesForUser(userID uint) ([]string, error) {
	e := c.syncedCachedEnforcer
	return e.GetUsersForRole(strconv.Itoa(int(userID)))
}

// HasRoot 判斷使用者是否有root權限
func (c *Casbinx) HasRoot(userID uint) (bool, error) {
	e := c.syncedCachedEnforcer
	return e.HasRoleForUser(strconv.Itoa(int(userID)), "root")
}

// AddRoleForUser 為使用者添加角色, 如果使用者已經擁有該角色, 則返回false。
func (c *Casbinx) AddRoleForUser(userID uint, role string) (bool, error) {
	e := c.syncedCachedEnforcer
	return e.AddRoleForUser(strconv.Itoa(int(userID)), role)
}

// UpdateRolesForUser 更新使用者角色
func (c *Casbinx) UpdateRolesForUser(userID uint, roles []string) error {
	if _, err := c.DeleteRolesForUser(userID); err != nil {
		return err
	}
	for _, role := range roles {
		if _, err := c.AddRoleForUser(userID, role); err != nil {
			return err
		}
	}
	return nil
}

// DeleteRoleForUser 刪除使用者的角色, 如果使用者沒有該角色, 則返回false。
func (c *Casbinx) DeleteRoleForUser(userID uint, role string) (bool, error) {
	e := c.syncedCachedEnforcer
	return e.DeleteRoleForUser(strconv.Itoa(int(userID)), role)
}

// DeleteRolesForUser 刪除使用者所有角色, 如果使用者沒有任何角色, 則返回false。
func (c *Casbinx) DeleteRolesForUser(userID uint) (bool, error) {
	e := c.syncedCachedEnforcer
	return e.DeleteRolesForUser(strconv.Itoa(int(userID)))
}

// DeleteUser 刪除使用者, 如果使用者不存在返回false
func (c *Casbinx) DeleteUser(userID uint) (bool, error) {
	e := c.syncedCachedEnforcer
	return e.DeleteUser(strconv.Itoa(int(userID)))
}

// DeleteRole 刪除角色, 如果使用者不存在返回false
func (c *Casbinx) DeleteRole(role string) (bool, error) {
	e := c.syncedCachedEnforcer
	return e.DeleteRole(role)
}

// AddPolicy 新增一筆權限
func (c *Casbinx) AddPolicy(role string, apiInfo ApiInfo) error {
	e := c.syncedCachedEnforcer
	success, _ := e.AddPolicy(role, apiInfo.Path, apiInfo.Method)
	if !success {
		return errors.New("failed to add, the same api already exists")
	}
	return nil
}

// Update 更新權限, 會把權限全部清除, 新增新的權限
func (c *Casbinx) Update(role string, ApiInfos []ApiInfo) error {
	c.ClearCasbin(0, role)
	var rules [][]string
	for _, v := range ApiInfos {
		rules = append(rules, []string{role, v.Path, v.Method})
	}
	e := c.syncedCachedEnforcer
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("failed to add, the same api already exists")
	}
	return nil
}

// UpdateApi Api更新
func (c *Casbinx) UpdateApi(old ApiInfo, new ApiInfo) error {
	err := c.db.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", old.Path, old.Method).Updates(map[string]interface{}{
		"v1": new.Path,
		"v2": new.Method,
	}).Error
	e := c.syncedCachedEnforcer
	err = e.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

// ClearCasbin 清除匹配的權限
func (c *Casbinx) ClearCasbin(v int, p ...string) bool {
	e := c.syncedCachedEnforcer
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// GetPolicyPathByRole 獲取權限列表
func (c *Casbinx) GetPolicyPathByRole(role string) (pathMaps []ApiInfo) {
	e := c.syncedCachedEnforcer
	list := e.GetFilteredPolicy(0, role)
	for _, v := range list {
		pathMaps = append(pathMaps, ApiInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}
