package model

type key interface {
	uint | int | string
}

type BasicFetchResult[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

type BasicPageParams struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

type TreeNode[K key, T key, C any] struct {
	Key      K    `json:"key"`
	Title    T    `json:"title"`
	Children []*C `json:"children"`
}
