package token

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	DefaultExpiration time.Duration = 1 * time.Minute // default time, 1 min
	Prefix            string        = "jwtx"
	whiteKey          string        = Prefix + "-white:%s"
	blackKey          string        = Prefix + "-black:%s"
	refreshKey        string        = Prefix + "-refresh:%s"
)

// SetWhite 設定白名單token
func SetWhite(client *redis.Client, token string) error {
	return client.Set(context.TODO(), fmt.Sprintf(whiteKey, token), true, DefaultExpiration).Err()
}

// IsWhite 判斷是否為白名單
func IsWhite(client *redis.Client, token string) (bool, error) {
	exists, err := client.Exists(context.TODO(), fmt.Sprintf(whiteKey, token)).Result()
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

// SetBlack 設置黑名單token
func SetBlack(client *redis.Client, token string) error {
	return client.Set(context.TODO(), fmt.Sprintf(blackKey, token), true, 10*time.Minute).Err()
}

// IsBlack 判斷是否為黑名單
func IsBlack(client *redis.Client, token string) (bool, error) {
	exists, err := client.Exists(context.TODO(), fmt.Sprintf(blackKey, token)).Result()
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

// SetRefreshToken 設置refresh token到redis緩存
func SetRefreshToken(client *redis.Client, key string, token string, ep time.Duration) error {
	return client.Set(context.TODO(), fmt.Sprintf(refreshKey, key), token, ep).Err()
}

// DelRefreshToken 刪除使用者refresh token
func DelRefreshToken(client *redis.Client, key string) error {
	_, err := client.Del(context.TODO(), fmt.Sprintf(refreshKey, key)).Result()
	return err
}

// GetRefreshToken 從redis獲取refresh token
func GetRefreshToken(client *redis.Client, key string) (string, error) {
	token, err := client.Get(context.TODO(), fmt.Sprintf(refreshKey, key)).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}
