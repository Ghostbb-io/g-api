package model

type ApiPageParams struct {
	BasicPageParams
	ApiParams
}

type ApiParams struct {
	Path   string `form:"path"`
	Group  string `form:"group"`
	Method string `form:"method"`
}

type ApiItem struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Path      string `json:"path"`
	Group     string `json:"group"`
	Desc      string `json:"desc"`
	Method    string `json:"method"`
}

type AddApiRequest struct {
	ApiItem
}

type EditApiRequest struct {
	ApiItem
}

type ApiTree struct {
	TreeNode[string, string, ApiTree]
}
