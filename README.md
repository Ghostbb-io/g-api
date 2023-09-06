
<div align=center>
    <h1>g-api</h1>
</div>
<div align=center>
    <img src="https://img.shields.io/badge/golang-1.21-blue" alt="golang"/>
    <img src="https://img.shields.io/badge/gin-1.9.1-00BB00" alt="gin"/>
    <img src="https://img.shields.io/badge/gorm-1.25.2-red" alt="gorm"/>
    <img src="https://img.shields.io/badge/redis-9.0.5-FF2D2D" alt="gorm"/>
    <img src="https://img.shields.io/badge/jwt-5.0.0-8A2BE2" alt="jwt"/>
</div>

## 1. 基本介绍
> g-api 框架是一個基於 [gin](https://gin-gonic.com) 開發的後端基礎平台，集成jwt，gorm，viper，cors，zap，go-redis等常用功能。

## 2. 使用說明
```
- golang版本 >= v1.21
- IDE推薦：Goland
```

### 2.1 g-api項目
使用 `Goland，vscode` 等編輯工具，打開g-api目錄

```bash
# 克隆項目
git clone https://github.com/YuWeiGhostbb/g-api.git
# 進入g-api目錄
cd g-api

# 使用 go mod 並安裝go依賴包
go generate

# 編譯
go build -o server main.go (windows編譯命令為go build -o server.exe main.go )

# 運行二進制
./server (windows運行命令為 server.exe)

# 打包成Docker
make docker_build

# 運行image
docker run --rm -d  g-api:latest
```

### 2.2 swagger自動化API文檔
#### 2.2.1 安裝 swagger

````
go get -u github.com/swaggo/swag/cmd/swag
````

#### 2.2.2 生成API文檔

```` shell
cd g-api
make swag
````

> 執行上面的命令後，g-api/core/swagger目錄下會出現docs，資料夾裡的 `docs.go`, `swagger.json`, `swagger.yaml` 三個文件更新，啟動go服務之後, 在瀏覽器輸入 [https://127.0.0.1:9000/api/swagger/index.html](https://127.0.0.1:9000/api/swagger/index.html) 即可查看swagger文檔

## 3. 技術選用
- Web框架：用 [Gin](https://gin-gonic.com/) 快速搭建基礎restful風格API，[Gin](https://gin-gonic.com/) 是一個go語言編寫的Web框架。
- 資料庫：用 [gorm](http://gorm.cn) 實現對資料庫的基本操作，支持mysql，mssql，pgsql。
- 緩存：使用`Redis`實現緩存，常用資料先對Redis進行查詢，如果沒查到才會到資料庫裡進行查詢。
- API文檔：使用`Swagger`構建自動化文檔。
- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 實現`yaml`格式的配置文件。
- 日誌：使用 [zap](https://github.com/uber-go/zap) 實現日誌紀錄。

## 4. 項目架構
### 4.1 系统架構圖
有空再補
![系統架構圖(有空再補)]()

### 4.2 Layout

#### 4.2.1 目錄結構
```
└── g-api
    ├── app             (應用層)
    ├── cmd             
    │   └── g-api          (main.go存放位置)
    ├── core            (核心文件)
    ├── pkg             (包)
    └── script          (腳本)
```
#### 4.2.2 core結構
```
└── core
    ├── log             
    │   ├── file_rotatelogs.go      (日誌分割)
    │   └── zap.go                  (zap內部操作)
    ├── middleware             
    │   ├── auth.go                 (JWT驗證中間件)
    │   ├── casbin_rbac.go          (rbac中間件)
    │   ├── cors.go                 (跨域處理中間件)
    │   ├── logger.go               (gin log輸出中間件)
    │   └── timeout.go              (timeout處理中間件)
    ├── router           
    │   └── router.go               (路由設定)
    ├── server          
    │   └── server.go               (web server)
    ├── swagger          (swagger文檔)
    │   ├── docs.go                 
    │   ├── swagger.json             
    │   └── swagger.yaml            
    ├── gorm.go          (gorm設定)
    ├── redis.go         (redis設定)
    ├── viper.go         (viper設定)
    └── zap.go           (zap設定)
```
#### 4.2.3 app結構
```
└── app
    ├── system 
    │   ├── api
    │   │   ├── v1        (v1版本接口)
    │   │   └── ...       (v2，v3...)  
    │   ├── model     
    │   │   └── table     (資料庫表)        
    │   └── service    (service層)
    └── app.go      (註冊插件)     
```

## 5. 主要功能
- 權限管理：基於`jwt`實現的權限管理，生產accessToken和refreshToken雙刷機制達到使用者無感刷新。
- 日誌：使用Zap進行日誌輸出，利用file-rotatelogs進行日誌分割。
- 設定檔：config.yaml用於開發階段，config-docker.yaml用於生產階段。
- 資料庫：支援多種，在config裡設定。
- 資料庫緩存：將資料庫查詢結果緩存到redis，資料庫查詢前先檢查redis，沒結果才會進行資料庫操作。三種模式對應緩存
  - `CacheLevelOnlyPrimary：只緩存primary key操作`
  - `CacheLevelOnlySearch: 指緩存搜尋`
  - `CacheLevelAll: 緩存全部`
- Cors：設計三種模式對應跨域處理
  - `allow-all：放行全部;`
  - `whitelist：白名單模式, 來自白名單內域名的請求添加 cors 頭;`
  - `strict-whitelist：嚴格白名單模式, 白名單外的請求一律拒絕;`
- restful示例：参考Swagger 文檔。

## 6. app
### 6.1 插件
>每個插件目錄裡必須有`enter.go`，並且由`New()`回傳所有路由  
  
system為例  
`system/enter.go`
```
func New() []any {
	err := global.GB_DB.AutoMigrate(
		// 自動建立表
		&table.SysUser{},
		&table.SysRole{},
	)
	if err != nil {
		global.GB_LOG.Error("create table error", zap.Error(err))
		os.Exit(0)
	}
	return []any{
		// 回傳Api群組
		&v1.BaseApi{service.BaseService},
		&v1.UserApi{service.UserService},
		&v1.CasbinApi{service.CasbinService},
		&v1.RoleApi{service.RoleService},
		&v1.CacheApi{service.CacheService},
	}
}
```
`app.go`
```
func All() [][]any {
	// 註冊插件
	return [][]any{
		system.New(),
	}
}
```
### 6.2 api層
>每個api群組必須用`struct`，並且要有`Register(ver ginx.VersionFunc)`方法  
每個api對應一個service
  
base.go為例  
`api/v1/base.go`
```
type BaseApi struct {
	service.Base
}

func (b *BaseApi) Register(ver ginx.VersionFunc) {
	v1 := ver(1).Group("")
	// v2 := ver(2).Group("")
	// ...
	{
		v1.POST("login", b.login)
		v1.POST("register", b.registerUser)
	}
	v1Private := v1.Group("", middleware.Auth())
	{
		v1Private.DELETE("logout", b.logout)
	}
}

// @Tags      系統
// @Summary   登入
// @Produce   application/json
// @Param     Info body model.LoginRequest true "帳號&密碼"
// @Success   200  {object}  response.Response{data=model.LoginResponse,msg=string}  "操作成功"
// @Router    /v1/login [post]
func (b *BaseApi) login(c *gin.Context) {
	var json model.LoginRequest
	if err := ginx.ParseJSON(c, &json); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	result, err := b.Login(json)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, result)
}
```
### 6.2 service層
>每個service必須用struct，再由interface進行包裝，最後用變數進行公開  
  
base.go為例  
`service/base.go`
```
var BaseService Base = new(base)

type Base interface {
	Login(model.LoginRequest) (model.LoginResponse, error)
	RegisterUser(model.RegisterRequest) error
	Logout(uuid.UUID, string) error
}

type base struct{}

// Login 登入
func (base) Login(in model.LoginRequest) (model.LoginResponse, error) {
	...
	return model.LoginResponse{Token: accessToken}, nil
}
```

## 7. global
`pkg/global/global.go`
```
var (
	GB_DB     *gorm.DB  
	GB_DBS    map[string]*gorm.DB
	GB_CONFIG config.Server
	GB_LOG    *zap.Logger
	GB_VP     *viper.Viper
	GB_REDIS  *redisx.Redis
	GB_SF     = &singleflight.Group{}
)
```
```
global.GB_DB.First(&user, "username = ?", username)
```
