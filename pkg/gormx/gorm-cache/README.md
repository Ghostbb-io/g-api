# gorm-cache

`gorm-cache` 旨在為gorm v2用戶提供一個即插即用的旁路緩存解決方案。本緩存只適用於數據庫表單主鍵時的場景。

## 特性
- 即插即用
- 旁路緩存
- 穿透防護
- 擊穿防護
- 多儲存介質（内存/redis）

## 使用說明

```go
import (
	"context"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/cache"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/storage"
	"github.com/redis/go-redis/v9"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	cache, _ := cache.NewGorm2Cache(&config.CacheConfig{
		CacheLevel:           config.CacheLevelAll,
		CacheStorage:         storage.NewRedis(&storage.RedisStoreConfig{Client: redisClient}),
		InvalidateWhenUpdate: true, // when you create/update/delete objects, invalidate cache
		CacheTTL:             5000, // 5000 ms
		CacheMaxItemCnt:      50,   // if length of objects retrieved one single time 
		// exceeds this number, then don't cache
	})
	// More options in `config/config.go`
	db.Use(cache) // use gorm plugin
	// cache.AttachToDB(db)

	var users []User

	db.Where("value > ?", 123).Find(&users) // search cache not hit, objects cached
	db.Where("value > ?", 123).Find(&users) // search cache hit

	db.Where("id IN (?)", []int{1, 2, 3}).Find(&users) // primary key cache not hit, users cached
	db.Where("id IN (?)", []int{1, 3}).Find(&users)    // primary key cache hit
}
```

在gorm中主要有5種操作（括號中是gorm中對應函數名）:

1. Query (First/Take/Last/Find/FindInBatches/FirstOrInit/FirstOrCreate/Count/Pluck)
2. Create (Create/CreateInBatches/Save)
3. Delete (Delete)
4. Update (Update/Updates/UpdateColumn/UpdateColumns/Save)
5. Row (Row/Rows/Scan)

本庫不支持Row操作的緩存。

## 存儲介質細節

本庫支持使用2種 cache 存儲介質：

1. 内存 (ccache/gcache)
2. Redis (所有數據存儲在redis中，如果你有多個實例使用本緩存，那麽他們不共享redis存儲空間)

並且允許多個gorm-cache公用一個存儲池，以確保同一數據庫的多個gorm實例共享緩存。
