package service

import (
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/pkg/casbinx"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"strings"
)

var CasbinService Casbin = new(casbin)

type Casbin interface {
	ApiList() []model.ApiResponse
	UpdateAuth(string, []model.ApiRequest) error
	ApiListByRole(string) []model.ApiRequest
	RemoveAuth([]model.ApiRequest) error
	AddAuth(string, model.ApiRequest) error
}

type casbin struct{}

// ApiList 獲取Api列表
func (casbin) ApiList() []model.ApiResponse {
	engine := ginx.GetEngine()
	routes := make([]model.ApiResponse, 0)
	for _, route := range engine.Routes() {
		routes = append(routes, model.ApiResponse{
			Method: route.Method,
			Path:   strings.TrimPrefix(route.Path, global.GB_CONFIG.System.RouterPrefix),
		})
	}
	return routes
}

// UpdateAuth 更新權限
func (casbin) UpdateAuth(role string, in []model.ApiRequest) error {
	apis := make([]casbinx.ApiInfo, 0)
	for _, api := range in {
		apis = append(apis, casbinx.ApiInfo{Path: api.Path, Method: api.Method})
	}
	cb, _ := casbinx.New(global.GB_DB)
	return cb.Update(role, apis)
}

// ApiListByRole 根據role取的api列表
func (casbin) ApiListByRole(role string) []model.ApiRequest {
	cb, _ := casbinx.New(global.GB_DB)
	paths := cb.GetPolicyPathByRole(role)
	apis := make([]model.ApiRequest, 0)
	for _, path := range paths {
		apis = append(apis, model.ApiRequest{Path: path.Path, Method: path.Method})
	}
	return apis
}

// RemoveAuth 刪除api
func (casbin) RemoveAuth(in []model.ApiRequest) error {
	cb, _ := casbinx.New(global.GB_DB)
	for _, api := range in {
		cb.ClearCasbin(1, api.Path, api.Method)
	}
	return nil
}

// AddAuth 新增一筆權限
func (casbin) AddAuth(role string, in model.ApiRequest) error {
	cb, _ := casbinx.New(global.GB_DB)
	return cb.AddPolicy(role, casbinx.ApiInfo{Path: in.Path, Method: in.Method})
}
