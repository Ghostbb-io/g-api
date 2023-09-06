# casbinx
>`casbinx`實現對api權限控管，使用rbac

## 規則
> root為超級管理員，無視所有規則  
> `sub`->角色  
> `obj`->路由  
> `sub`->請求方法
```
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
```

## 範例
```go
// Casbin 攔截器
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cb, _ := casbinx.New(global.GB_DB)
		e := cb.GetSyncedCachedEnforcer()
		userId := ginx.GetUserID(c)

		sub := strconv.Itoa(int(userId))
		obj := c.Request.URL.Path
		act := c.Request.Method
		
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			response.FailWithDetailed(c, gin.H{}, "no auth")
			c.Abort()
			return
		}
		c.Next()
	}
}
```