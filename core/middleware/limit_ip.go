package middleware

import (
	"context"
	"errors"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type LimitConfig struct {
	// GenerationKey 根據業務生成key 下面CheckOrMark查詢生成
	GenerationKey func(c *gin.Context) string
	// CheckOrMark 檢查函數,用戶可修改具體邏輯,更加靈活
	CheckOrMark func(key string, expire int, limit int) error
	// Expire key 過期時間
	Expire int
	// Limit 周期時間
	Limit int
}

// LimitWithTime Generate middleware on Config
func (l LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

// DefaultGenerationKey 默認生成key
func DefaultGenerationKey(c *gin.Context) string {
	return "GB_Limit:" + c.ClientIP()
}

// DefaultCheckOrMark 檢查
func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
		global.GB_LOG.Error("limit", zap.Error(err))
	}
	return err
}

// SetLimitWithTime Set number of visits
func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	redis := global.GB_REDIS.GetClient()
	count, err := redis.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		pipe := redis.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err = pipe.Exec(context.Background())
		return err
	} else {
		// 次數
		if times, err := redis.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := redis.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("the request is too frequent, please try again later")
				} else {
					return errors.New("the request is too frequent, please try again in " + t.String() + " seconds.")
				}
			} else {
				return redis.Incr(context.Background(), key).Err()
			}
		}
	}
}

func DefaultLimit() gin.HandlerFunc {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   DefaultCheckOrMark,
		Expire:        global.GB_CONFIG.System.LimitTimeIP,
		Limit:         global.GB_CONFIG.System.LimitCountIP,
	}.LimitWithTime()
}
