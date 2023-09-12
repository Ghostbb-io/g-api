package service

import (
	"fmt"
	"github.com/Ghostbb-io/g-api/app/system/model"
	"github.com/Ghostbb-io/g-api/app/system/model/table"
	"github.com/Ghostbb-io/g-api/pkg/global"
)

var ApiService Api = new(api)

type Api interface {
	ApiListByPage(model.ApiPageParams) (model.BasicFetchResult[model.ApiItem], error)
	AddApi(model.AddApiRequest) error
	EditApi(uint, model.EditApiRequest) error
	DelApi(uint) error
	ApiTree() ([]*model.ApiTree, error)
}

type api struct{}

// ApiListByPage 根據page, pageSize獲取Api列表
func (api) ApiListByPage(in model.ApiPageParams) (model.BasicFetchResult[model.ApiItem], error) {
	methods := []string{"GET", "POST", "PATCH", "PUT", "DELETE"}
	if in.Method != "" {
		methods = []string{in.Method}
	}
	var apis []table.SysApi
	err := global.GB_DB.
		Where("path like ?", "%"+in.Path+"%").
		Where("[group] like ?", "%"+in.Group+"%").
		Where("method in ?", methods).
		Limit(in.PageSize).Offset((in.Page - 1) * in.PageSize).Find(&apis).Error
	if err != nil {
		return model.BasicFetchResult[model.ApiItem]{}, err
	}

	var result model.BasicFetchResult[model.ApiItem]
	for _, api := range apis {
		result.Items = append(result.Items, model.ApiItem{
			ID:        api.ID,
			CreatedAt: api.CreatedAt.Format("2006-01-02 - 15:04:05"),
			UpdatedAt: api.UpdatedAt.Format("2006-01-02 - 15:04:05"),
			Path:      api.Path,
			Group:     api.Group,
			Desc:      api.Desc,
			Method:    api.Method,
		})
	}
	if err = global.GB_DB.Model(&table.SysApi{}).Count(&result.Total).Error; err != nil {
		return model.BasicFetchResult[model.ApiItem]{}, err
	}
	return result, nil
}

// AddApi 新增api
func (api) AddApi(in model.AddApiRequest) error {
	var api table.SysApi
	{
		api.Path = in.Path
		api.Group = in.Group
		api.Method = in.Method
		api.Desc = in.Desc
	}
	return global.GB_DB.Create(&api).Error
}

// EditApi 編輯api
func (api) EditApi(id uint, in model.EditApiRequest) error {
	var api table.SysApi
	{
		api.ID = id
		api.Path = in.Path
		api.Group = in.Group
		api.Desc = in.Desc
		api.Method = in.Method
	}
	return global.GB_DB.Save(&api).Error
}

// DelApi 刪除api
func (api) DelApi(id uint) error {
	return global.GB_DB.Where("id = ?", id).Delete(&table.SysApi{}).Error
}

// ApiTree 獲取api tree
func (api) ApiTree() ([]*model.ApiTree, error) {
	var apis []table.SysApi
	if err := global.GB_DB.Find(&apis).Error; err != nil {
		return nil, err
	}
	result := make([]*model.ApiTree, 0)
	resMap := make(map[string]*model.ApiTree)
	for _, api := range apis {
		if _, ok := resMap[api.Group]; !ok {
			temps := &model.ApiTree{
				TreeNode: model.TreeNode[string, string, model.ApiTree]{
					Key:      api.Group,
					Title:    api.Group + "組",
					Children: make([]*model.ApiTree, 0),
				},
			}
			resMap[api.Group] = temps
			result = append(result, temps)
		}
		resMap[api.Group].Children = append(resMap[api.Group].Children, &model.ApiTree{
			TreeNode: model.TreeNode[string, string, model.ApiTree]{
				Key:   fmt.Sprintf("%s[%s]", api.Path, api.Method),
				Title: api.Desc,
			},
		})
	}
	return result, nil
}
