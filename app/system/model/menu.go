package model

import "github.com/Ghostbb-io/g-api/app/system/model/table"

type MenuListResponse struct {
	ParentID  uint                `json:"-"`
	Path      string              `json:"path"`
	Name      string              `json:"name"`
	Component string              `json:"component"`
	Redirect  string              `json:"redirect"`
	Meta      table.Meta          `json:"meta"`
	Children  []*MenuListResponse `json:"children"`
}
