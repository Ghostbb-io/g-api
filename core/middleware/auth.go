package middleware

import (
	"errors"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"github.com/Ghostbb-io/g-api/pkg/jwtx"
	"github.com/Ghostbb-io/g-api/pkg/jwtx/claims"
	"github.com/Ghostbb-io/g-api/pkg/jwtx/token"
	"github.com/Ghostbb-io/g-api/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ep, _ := jwtx.ParseDuration(global.GB_CONFIG.JWT.ExpiresTime)
		jwt := jwtx.New(global.GB_CONFIG.JWT.SigningKey, global.GB_CONFIG.JWT.Issuer, ep)
		accessToken := ginx.GetToken(c)
		if accessToken == "" {
			response.UnAuth(c, "token not valid")
			c.Abort()
			return
		}
		// 判斷token是否在黑名單
		if exist, _ := token.IsBlack(global.GB_REDIS.GetClient(), accessToken); exist {
			response.UnAuth(c, "token not valid")
			c.Abort()
			return
		}
		accessClaims, err := jwt.ParseAccessToken(accessToken)
		if err != nil {
			if errors.Is(err, jwtx.TokenExpired) {
				// 如果在白名單，直接放行
				if exist, _ := token.IsWhite(global.GB_REDIS.GetClient(), accessToken); !exist {
					newAccessToken, err, _ := refresh(jwt, accessToken, accessClaims, ep)
					if err != nil {
						response.UnAuth(c, err.Error())
						c.Abort()
						return
					}
					// 將新accessToken放進header
					if newAccessToken != nil {
						// 把舊的Token放進白名單
						_ = token.SetWhite(global.GB_REDIS.GetClient(), accessToken)
						c.Header("new-token", newAccessToken.(string))
					}
				}
			} else {
				response.UnAuth(c, err.Error())
				c.Abort()
				return
			}
		}
		// 寫入上下文
		ginx.SetUserID(c, accessClaims.BaseClaims.ID)
		ginx.SetUserUUID(c, accessClaims.UUID)
		ginx.SetUserName(c, accessClaims.Username)
		ginx.SetRole(c, accessClaims.Roles)
		c.Next()
	}
}

func refresh(jwt *jwtx.Jwtx, accessToken string, accessClaims *claims.AccessClaims, ep time.Duration) (interface{}, error, bool) {
	// 使用singleflight避免併發
	return global.GB_SF.Do("JWT:"+accessToken, func() (interface{}, error) {
		// 從Redis獲取RefreshToken
		refreshToken, err := token.GetRefreshToken(global.GB_REDIS.GetClient(), accessClaims.BaseClaims.UUID.String())
		if err != nil {
			return nil, err
		}
		// 換發
		newAccessToken, newRefreshToken, err := jwt.Refreshing(accessToken, refreshToken)
		if err != nil {
			return nil, err
		}
		// 將新的RefreshToken放回去Redis
		if err = token.SetRefreshToken(global.GB_REDIS.GetClient(), accessClaims.BaseClaims.UUID.String(), newRefreshToken, ep); err != nil {
			return nil, err
		}
		return newAccessToken, nil
	})
}
