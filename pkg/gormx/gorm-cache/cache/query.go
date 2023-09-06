package cache

import (
	"errors"
	"fmt"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/config"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/storage"
	"github.com/Ghostbb-io/g-api/pkg/gormx/gorm-cache/util"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// singleFlight 流程設計
// 根據key lock住，等待結果。query before之前，會先判斷是否有key，如果有，就等待結果，如果沒有，就執行query before，然後執行query，然後把結果放到key里面，然後unlock，然後返回結果。
// 等待完成後，進行一手返回，然後err設置為err.singleflightHit，afterQuery結束的時候進行一手檢查

func newQueryHandler(c *Gorm2Cache) *queryHandler {
	return &queryHandler{cache: c}
}

type queryHandler struct {
	cache        *Gorm2Cache
	singleFlight Group
}

func (h *queryHandler) Bind(db *gorm.DB) error {
	err := db.Callback().Query().Before("gorm:query").Register("gorm:cache:before_query", h.BeforeQuery())
	if err != nil {
		return err
	}
	err = db.Callback().Query().After("gorm:after_query").Register("gorm:cache:after_query", h.AfterQuery())
	if err != nil {
		return err
	}
	return nil
}

func (h *queryHandler) BeforeQuery() func(db *gorm.DB) {
	cache := h.cache
	return func(db *gorm.DB) {
		callbacks.BuildQuerySQL(db)
		tableName := ""
		if db.Statement.Schema != nil {
			tableName = db.Statement.Schema.Table
		} else {
			tableName = db.Statement.Table
		}
		ctx := db.Statement.Context

		sql := db.Statement.SQL.String()
		db.InstanceSet("gorm:cache:sql", sql)
		db.InstanceSet("gorm:cache:vars", db.Statement.Vars)

		if util.ShouldCache(tableName, cache.Config.Tables) {
			hit := false
			defer func() {
				if hit {
					cache.IncrHitCount()
				} else {
					cache.IncrMissCount()
				}
			}()

			// singleFlight Check
			singleFlightKey := util.GenSingleFlightKey(tableName, sql, db.Statement.Vars...)
			h.singleFlight.mu.Lock()
			if h.singleFlight.m == nil {
				h.singleFlight.m = make(map[string]*call)
			}
			if c, ok := h.singleFlight.m[singleFlightKey]; ok {
				c.dups++
				h.singleFlight.mu.Unlock()
				c.wg.Wait()

				// 临时糊一个拷贝在这里 性能可能并不是那么好
				d, err := json.Marshal(c.dest)
				if err != nil {
					_ = db.AddError(err)
					return
				}
				err = json.Unmarshal(d, db.Statement.Dest)
				if err != nil {
					_ = db.AddError(err)
					return
				}
				hit = true
				db.RowsAffected = c.rowsAffected
				db.Error = multierror.Append(util.SingleFlightHit) // 为保证后续流程不走，必须设一个error
				if c.err != nil {
					db.Error = multierror.Append(db.Error, c.err)
				}
				h.cache.Logger.CtxInfo(ctx, "[BeforeQuery] single flight hit for key %v", singleFlightKey)
				return
			}
			c := &call{key: singleFlightKey}
			c.wg.Add(1)
			h.singleFlight.m[singleFlightKey] = c
			h.singleFlight.mu.Unlock()
			db.InstanceSet("gorm:cache:query:single_flight_call", c)

			tryPrimaryCache := func() (hit bool) {
				primaryKeys := getPrimaryKeysFromWhereClause(db)
				cache.Logger.CtxInfo(ctx, "[BeforeQuery] parse primary keys = %v", primaryKeys)

				if len(primaryKeys) == 0 {
					return
				}

				// if (IN primaryKeys)/(Eq primaryKey) are the only clauses
				hasOtherClauseInWhere := hasOtherClauseExceptPrimaryField(db)
				if hasOtherClauseInWhere {
					// if query has other clauses, it can only query the database
					return
				}

				// primary cache hit
				cacheValues, err := cache.BatchGetPrimaryCache(ctx, tableName, primaryKeys)
				if err != nil {
					cache.Logger.CtxError(ctx, "[BeforeQuery] get primary cache value for key %v error: %v", primaryKeys, err)
					db.Error = nil
					return
				}
				if len(cacheValues) != len(primaryKeys) {
					db.Error = nil
					return
				}
				finalValue := ""

				destKind := reflect.Indirect(reflect.ValueOf(db.Statement.Dest)).Kind()
				if destKind == reflect.Struct && len(cacheValues) == 1 {
					finalValue = cacheValues[0]
				} else if (destKind == reflect.Array || destKind == reflect.Slice) && len(cacheValues) >= 1 {
					finalValue = "[" + strings.Join(cacheValues, ",") + "]"
				}
				if len(finalValue) == 0 {
					cache.Logger.CtxError(ctx, "[BeforeQuery] length of cache values and dest not matched")
					db.Error = util.ErrCacheUnmarshal
					return
				}

				err = json.Unmarshal([]byte(finalValue), db.Statement.Dest)
				if err != nil {
					cache.Logger.CtxError(ctx, "[BeforeQuery] unmarshal final value error: %v", err)
					db.Error = util.ErrCacheUnmarshal
					return
				}
				db.Error = util.PrimaryCacheHit
				hit = true
				return
			}

			trySearchCache := func() (hit bool) {
				// search cache hit
				cacheValue, err := cache.GetSearchCache(ctx, tableName, sql, db.Statement.Vars...)
				if err != nil {
					if !errors.Is(err, storage.ErrCacheNotFound) {
						cache.Logger.CtxError(ctx, "[BeforeQuery] get cache value for sql %s error: %v", sql, err)
					}
					db.Error = nil
					return
				}
				cache.Logger.CtxInfo(ctx, "[BeforeQuery] get value: %s", cacheValue)
				if cacheValue == "recordNotFound" { // 應對緩存穿透
					db.Error = util.RecordNotFoundCacheHit
					hit = true
					return
				}
				rowsAffectedPos := strings.Index(cacheValue, "|")
				db.RowsAffected, err = strconv.ParseInt(cacheValue[:rowsAffectedPos], 10, 64)
				if err != nil {
					cache.Logger.CtxError(ctx, "[BeforeQuery] unmarshal rows affected cache error: %v", err)
					db.Error = nil
					return
				}
				err = json.Unmarshal([]byte(cacheValue[rowsAffectedPos+1:]), db.Statement.Dest)
				if err != nil {
					cache.Logger.CtxError(ctx, "[BeforeQuery] unmarshal search cache error: %v", err)
					db.Error = nil
					return
				}
				db.Error = util.SearchCacheHit
				hit = true
				return
			}

			if cache.Config.CacheLevel == config.CacheLevelAll || cache.Config.CacheLevel == config.CacheLevelOnlyPrimary {
				if tryPrimaryCache() {
					hit = true
					return
				}
			}
			if cache.Config.CacheLevel == config.CacheLevelAll || cache.Config.CacheLevel == config.CacheLevelOnlySearch {
				if !hit && trySearchCache() {
					hit = true
				}
			}
		}
	}
}

func (h *queryHandler) AfterQuery() func(db *gorm.DB) {
	cache := h.cache
	return func(db *gorm.DB) {
		func() {
			tableName := ""
			if db.Statement.Schema != nil {
				tableName = db.Statement.Schema.Table
			} else {
				tableName = db.Statement.Table
			}
			ctx := db.Statement.Context
			sqlObj, _ := db.InstanceGet("gorm:cache:sql")
			sql := sqlObj.(string)
			varObj, _ := db.InstanceGet("gorm:cache:vars")
			vars := varObj.([]interface{})

			if !util.ShouldCache(tableName, cache.Config.Tables) {
				return
			}

			if db.Error == nil {
				destValue := reflect.Indirect(reflect.ValueOf(db.Statement.Dest))
				// 如果是结构体应该能提主键出来
				// 如果是数组需要判断内部元素是不是结构体，不是结构体的都提不了主键
				if destValue.Kind() == reflect.Slice || destValue.Kind() == reflect.Array {
					if (destValue.Type().Elem().Kind() == reflect.Pointer && destValue.Type().Elem().Elem().Kind() != reflect.Struct) ||
						(destValue.Type().Elem().Kind() != reflect.Pointer && destValue.Type().Elem().Kind() != reflect.Struct) {
						return
					}
				}

				// error is nil -> cache not hit, we cache newly retrieved data
				primaryKeys, objects := getObjectsAfterLoad(db)

				var wg sync.WaitGroup
				wg.Add(2)

				go func() {
					defer wg.Done()

					if cache.Config.CacheLevel == config.CacheLevelAll || cache.Config.CacheLevel == config.CacheLevelOnlySearch {
						// cache search data
						if cache.Config.CacheMaxItemCnt != 0 && int64(len(objects)) > cache.Config.CacheMaxItemCnt {
							return
						}

						cache.Logger.CtxInfo(ctx, "[AfterQuery] start to set search cache for sql: %s", sql)
						cacheBytes, err := json.Marshal(db.Statement.Dest)
						if err != nil {
							cache.Logger.CtxError(ctx, "[AfterQuery] cannot marshal cache for sql: %s, not cached", sql)
							return
						}
						cache.Logger.CtxInfo(ctx, "[AfterQuery] set cache: %v", string(cacheBytes))
						err = cache.SetSearchCache(ctx, fmt.Sprintf("%d|", db.RowsAffected)+string(cacheBytes), tableName, sql, vars...)
						if err != nil {
							cache.Logger.CtxError(ctx, "[AfterQuery] set search cache for sql: %s error: %v", sql, err)
							return
						}
						cache.Logger.CtxInfo(ctx, "[AfterQuery] sql %s cached", sql)
					}
				}()

				go func() {
					defer wg.Done()

					if cache.Config.CacheLevel == config.CacheLevelAll || cache.Config.CacheLevel == config.CacheLevelOnlyPrimary {
						// cache primary cache data
						if len(primaryKeys) != len(objects) {
							return
						}
						if cache.Config.CacheMaxItemCnt != 0 && int64(len(objects)) > cache.Config.CacheMaxItemCnt {
							cache.Logger.CtxInfo(ctx, "[AfterQuery] objects length is more than max item count, not cached")
							return
						}
						kvs := make([]util.Kv, 0, len(objects))
						for i := 0; i < len(objects); i++ {
							jsonStr, err := json.Marshal(objects[i])
							if err != nil {
								cache.Logger.CtxError(ctx, "[AfterQuery] object %v cannot marshal, not cached", objects[i])
								continue
							}
							kvs = append(kvs, util.Kv{
								Key:   primaryKeys[i],
								Value: string(jsonStr),
							})
						}
						cache.Logger.CtxInfo(ctx, "[AfterQuery] start to set primary cache for kvs: %+v", kvs)
						err := cache.BatchSetPrimaryKeyCache(ctx, tableName, kvs)
						if err != nil {
							cache.Logger.CtxError(ctx, "[AfterQuery] batch set primary key cache for key %v error: %v",
								primaryKeys, err)
						}
					}
				}()
				if !cache.Config.AsyncWrite {
					wg.Wait()
				}
				return
			}

			// 對應緩存穿透
			if errors.Is(db.Error, gorm.ErrRecordNotFound) && !cache.Config.DisableCachePenetrationProtect {
				cache.Logger.CtxInfo(ctx, "[AfterQuery] set cache: %v", "recordNotFound")
				err := cache.SetSearchCache(ctx, "recordNotFound", tableName, sql, vars...)
				if err != nil {
					cache.Logger.CtxError(ctx, "[AfterQuery] set search cache for sql: %s error: %v", sql, err)
					return
				}
				cache.Logger.CtxInfo(ctx, "[AfterQuery] sql %s cached", sql)
				return
			}
		}()
		// 之所以將上面的部分包在一個匿名函數中是為了方便
		// 上面的cache完成後直接傳播給其他等待中的goroutine
		// 上面只處理非singleflight且無錯誤或記錄不存在的情況
		h.fillCallAfterQuery(db)

		// 下面處理命中了緩存的情況
		// 有以下幾種err是專門用來傳狀態的：正常的cacheHit 這種情況不存在error
		// RecordNotFoundCacheHit 這種情況只會在notfound之後出現
		// SingleFlightHit 這種情況下error中除了SingleFlightHit還可能會存在其他error來自gorm的error
		// 且遇到任何一種hit我們都可以認為是命中了緩存 同時只可能命中至多兩個hit（single+其他
		if merr, ok := db.Error.(*multierror.Error); ok {
			errs := merr.WrappedErrors()
			if errors.Is(errs[0], util.SingleFlightHit) {
				if len(errs) > 1 {
					db.Error = errs[1]
				} else {
					db.Error = nil
				}
			}
		}

		switch db.Error {
		case util.RecordNotFoundCacheHit:
			db.Error = gorm.ErrRecordNotFound
		case util.SearchCacheHit, util.PrimaryCacheHit:
			db.Error = nil
		}
	}
}

func (h *queryHandler) fillCallAfterQuery(db *gorm.DB) {
	if singleFlightCallObj, exist := db.InstanceGet("gorm:cache:query:single_flight_call"); exist {
		c := singleFlightCallObj.(*call)
		c.dest = db.Statement.Dest
		c.rowsAffected = db.RowsAffected
		c.err = db.Error
		c.wg.Done()

		h.singleFlight.mu.Lock()
		if !c.forgotten {
			delete(h.singleFlight.m, c.key)
		}
		h.singleFlight.mu.Unlock()
	}
}
