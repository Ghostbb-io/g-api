package service

import (
	"errors"
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/model/table"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"gorm.io/gorm"
	"strings"
)

var MenuService Menu = new(menu)

type Menu interface {
	MenuTree(model.MenuTreeParams) ([]*model.MenuTree, error)
	MenuList(model.MenuParams) (model.BasicFetchResult[*model.MenuItem], error)
	AddMenu(model.AddMenuRequest) error
	DelMenu(uint) error
	EditMenu(uint, model.EditMenuRequest) error
}

type menu struct{}

// MenuTree 獲取menu tree列表
func (menu) MenuTree(in model.MenuTreeParams) ([]*model.MenuTree, error) {
	var menus []table.SysMenu
	if in.Dir {
		if err := global.GB_DB.Where("type = ?", "dir").Order("sort").Find(&menus).Error; err != nil {
			return nil, err
		}
	} else {
		if err := global.GB_DB.Order("sort").Find(&menus).Error; err != nil {
			return nil, err
		}
	}

	result := make([]*model.MenuTree, 0)
	item := make([]*model.MenuTree, 0)
	menuMap := make(map[uint]*model.MenuTree)
	for _, menu := range menus {
		temps := &model.MenuTree{
			ParentID: menu.ParentID,
			Icon:     menu.Icon,
			TreeNode: model.TreeNode[uint, string, model.MenuTree]{
				Key:   menu.ID,
				Title: menu.Title,
			},
		}
		if menu.ParentID == 0 {
			result = append(result, temps)
		} else {
			item = append(item, temps)
		}
		menuMap[menu.ID] = temps
	}

	for _, menu := range item {
		if _, ok := menuMap[menu.ParentID]; ok {
			if menuMap[menu.ParentID].Children == nil {
				menuMap[menu.ParentID].Children = make([]*model.MenuTree, 0)
			}
			menuMap[menu.ParentID].Children = append(menuMap[menu.ParentID].Children, menu)
		}
	}
	return result, nil
}

// MenuList 獲取menu列表
func (menu) MenuList(in model.MenuParams) (model.BasicFetchResult[*model.MenuItem], error) {
	// query
	var menus []table.SysMenu
	tx := global.GB_DB.Order("sort").Where("title like ?", "%"+in.MenuName+"%")
	if in.Status != "" {
		if in.Status == "true" {
			tx.Where("status = ?", true)
		} else {
			tx.Where("status = ?", false)
		}
	}
	if err := tx.Find(&menus).Error; err != nil {
		return model.BasicFetchResult[*model.MenuItem]{}, err
	}

	// menu組裝
	var result model.BasicFetchResult[*model.MenuItem]
	item := make([]*model.MenuItem, 0)
	result.Items = make([]*model.MenuItem, 0)
	menuMap := make(map[uint]*model.MenuItem)
	for _, menu := range menus {
		temps := &model.MenuItem{
			ID:        menu.ID,
			CreatedAt: menu.CreatedAt.Format("\"2006-01-02 - 15:04:05\""),
			UpdatedAt: menu.UpdatedAt.Format("\"2006-01-02 - 15:04:05\""),
			Type:      menu.Type,
			ParentID:  menu.ParentID,
			Path:      menu.Path,
			Name:      menu.Name,
			Component: menu.Component,
			Redirect:  menu.Redirect,
			Sort:      menu.Sort,
			Status:    menu.Status,
			Meta:      menu.Meta,
		}
		if menu.ParentID == 0 {
			result.Items = append(result.Items, temps)
		} else {
			item = append(item, temps)
		}
		menuMap[menu.ID] = temps
	}
	for _, menu := range item {
		if _, ok := menuMap[menu.ParentID]; ok {
			if menuMap[menu.ParentID].Children == nil {
				menuMap[menu.ParentID].Children = make([]*model.MenuItem, 0)
			}
			menuMap[menu.ParentID].Children = append(menuMap[menu.ParentID].Children, menu)
		} else {
			result.Items = append(result.Items, menu)
		}
	}

	// count
	if err := global.GB_DB.Model(&table.SysMenu{}).Count(&result.Total).Error; err != nil {
		return model.BasicFetchResult[*model.MenuItem]{}, err
	}

	// return
	return result, nil
}

// AddMenu 新增menu
func (menu) AddMenu(in model.AddMenuRequest) error {
	var menu table.SysMenu
	if in.ParentID == 0 {
		if !strings.HasPrefix(in.Path, "/") {
			return errors.New("頂級路由開頭必須有'/'")
		}
		if in.Type == "dir" {
			in.Component = "LAYOUT"
		}
	} else {
		if strings.HasPrefix(in.Path, "/") {
			return errors.New("二級以上路由開頭不能有'/'")
		}
	}
	if in.Type == "iframe" {
		in.Component = "IFrame"
	}
	{
		menu.Type = in.Type
		menu.ParentID = in.ParentID
		menu.Path = in.Path
		menu.Name = in.Name
		menu.Component = in.Component
		menu.Redirect = in.Redirect
		menu.Sort = in.Sort
		menu.Status = in.Status
		menu.Meta = in.Meta
	}
	return global.GB_DB.Create(&menu).Error
}

// DelMenu 刪除menu
func (menu) DelMenu(id uint) error {
	return global.GB_DB.Transaction(func(tx *gorm.DB) error {
		var menu table.SysMenu
		if !errors.Is(tx.Where("parent_id = ?", id).First(&menu).Error, gorm.ErrRecordNotFound) {
			return errors.New("不能刪除")
		}
		if err := tx.Where("id = ?", id).Delete(&table.SysMenu{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// EditMenu 更新menu
func (menu) EditMenu(id uint, in model.EditMenuRequest) error {
	var menu table.SysMenu
	if in.ParentID == 0 {
		if !strings.HasPrefix(in.Path, "/") {
			return errors.New("頂級路由開頭必須有'/'")
		}
		if in.Type == "dir" {
			in.Component = "LAYOUT"
		}
	} else {
		if strings.HasPrefix(in.Path, "/") {
			return errors.New("二級以上路由開頭不能有'/'")
		}
	}
	if in.Type == "iframe" {
		in.Component = "IFrame"
	}

	menu.ID = id
	menu.ParentID = in.ParentID
	menu.Title = in.Title
	menu.Component = in.Component
	menu.Path = in.Path
	menu.Name = in.Name
	menu.Sort = in.Sort
	menu.Icon = in.Icon
	menu.Status = in.Status
	menu.HideBreadcrumb = in.HideBreadcrumb
	menu.Type = in.Type
	switch in.Type {
	case "dir":
		menu.Redirect = in.Redirect
		menu.HideChildrenInMenu = in.HideChildrenInMenu
		menu.IgnoreRoute = in.IgnoreRoute
		menu.HidePathForChildren = in.HidePathForChildren
	case "menu":
		menu.TransitionName = in.TransitionName
		menu.IgnoreKeepAlive = in.IgnoreKeepAlive
		menu.Affix = in.Affix
		menu.HideTab = in.HideTab
		menu.IgnoreRoute = in.IgnoreRoute
	case "iframe":
		menu.FrameSrc = in.FrameSrc
		menu.IgnoreKeepAlive = in.IgnoreKeepAlive
		menu.Affix = in.Affix
		menu.HideTab = in.HideTab
	default:
		return errors.New("type error")
	}
	return global.GB_DB.Save(&menu).Error
}
