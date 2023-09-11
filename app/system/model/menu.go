package model

import "github.com/Ghostbb-io/g-api/app/system/model/table"

type AddMenuRequest struct {
	MenuItem
}

type EditMenuRequest struct {
	MenuItem
}

type MenuParams struct {
	MenuName string `form:"menuName"`
	Status   string `form:"status"`
}

type MenuTreeParams struct {
	Dir bool `form:"dir"`
}

type MenuItem struct {
	ID         uint   `json:"id"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	Type       string `json:"type"`
	ParentID   uint   `json:"parentID"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	Sort       int    `json:"sort"`
	Status     bool   `json:"status"`
	table.Meta `json:",inline"`
	Children   []*MenuItem `json:"children"`
}

type MenuTree struct {
	ParentID uint   `json:"-"`
	Icon     string `json:"icon"`
	TreeNode[uint, string, MenuTree]
}
