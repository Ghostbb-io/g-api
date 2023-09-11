package table

import "github.com/Ghostbb-io/g-api/pkg/global"

type SysMenu struct {
	global.GB_MODEL
	Type      string `json:"type" gorm:"comment:dir menu iframe"`
	ParentID  uint   `json:"-"`
	Path      string `json:"path"`
	Name      string `json:"name" gorm:"comment:路由名稱"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Sort      int    `json:"sort"`
	Status    bool   `json:"status" gorm:"default:true;comment:是否啟用"`
	Meta      `json:"meta" gorm:"embedded;comment:附加屬性"`
	Children  []*SysMenu `json:"children" gorm:"-"`
	Roles     []SysRole  `json:"roles" gorm:"many2many:sys_role_menu;"`
}

type Meta struct {
	// 路由title  一般必填
	Title string `json:"title"`
	// 動態路由可打開Tab頁數
	DynamicLevel int `json:"dynamicLevel"`
	// 動態路由的實際Path, 即去除路由的動態部分;
	RealPath string `json:"realPath"`
	// 是否忽略KeepAlive緩存
	IgnoreKeepAlive bool `json:"ignoreKeepAlive"`
	// 是否固定標簽
	Affix bool `json:"affix"`
	// 圖標，也是菜單圖標
	Icon string `json:"icon"`
	// 內嵌iframe的地址
	FrameSrc string `json:"frameSrc"`
	// 指定該路由切換的動畫名
	TransitionName string `json:"transitionName"`
	// 隱藏該路由在面包屑上面的顯示
	HideBreadcrumb bool `json:"hideBreadcrumb"`
	// 如果該路由會攜帶參數，且需要在tab頁上面顯示。則需要設置為true
	CarryParam bool `json:"carryParam"`
	// 隱藏所有子菜單
	HideChildrenInMenu bool `json:"hideChildrenInMenu"`
	// 當前激活的菜單。用於配置詳情頁時左側激活的菜單路徑
	CurrentActiveMenu string `json:"currentActiveMenu"`
	// 當前路由不再標簽頁顯示
	HideTab bool `json:"hideTab"`
	// 當前路由不再菜單顯示
	HideMenu bool `json:"hideMenu"`
	// 忽略路由。用於在ROUTE_MAPPING以及BACK權限模式下，生成對應的菜單而忽略路由。2.5.3以上版本有效
	IgnoreRoute bool `json:"ignoreRoute"`
	// 是否在子級菜單的完整path中忽略本級path。2.5.3以上版本有效
	HidePathForChildren bool `json:"hidePathForChildren"`
}

func (SysMenu) TableName() string {
	return "sys_menus"
}
