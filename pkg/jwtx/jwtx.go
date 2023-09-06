package jwtx

import (
	"errors"
	"github.com/Ghostbb-io/g-api/pkg/jwtx/claims"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	TokenExpired    = errors.New("token is expired")
	RefreshTokenErr = errors.New("wrong refresh token")
)

type Jwtx struct {
	signingKey  []byte
	issuer      string
	expiresTime time.Duration
}

func New(signingKey string, issuer string, expiresTime time.Duration) *Jwtx {
	return &Jwtx{
		signingKey:  []byte(signingKey),
		issuer:      issuer,
		expiresTime: expiresTime,
	}
}

// CreateToken 創建Token
func (j *Jwtx) CreateToken(baseClaims claims.BaseClaims) (accessToken, refreshToken string, err error) {
	key, _ := uuid.NewUUID()

	accessClaims := claims.AccessClaims{
		BaseClaims: baseClaims,
		Key:        key,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)), // 過期時間，根據配置
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       // 簽發時間
			NotBefore: jwt.NewNumericDate(time.Now()),                       // 生效時間
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(j.signingKey)
	if err != nil {
		return "", "", err
	}

	refreshClaims := claims.RefreshClaims{
		Key: key,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiresTime)), // 過期時間，根據配置
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // 簽發時間
			NotBefore: jwt.NewNumericDate(time.Now()),                    // 生效時間
		},
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(j.signingKey)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// ParseAccessToken 解析AccessToken
func (j *Jwtx) ParseAccessToken(tokenString string) (*claims.AccessClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims.AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})

	if _claims, ok := token.Claims.(*claims.AccessClaims); ok && token.Valid {
		return _claims, nil
	} else {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return _claims, TokenExpired
		}
		return nil, err
	}
}

// ParseRefreshToken 解析refreshToken
func (j *Jwtx) ParseRefreshToken(tokenString string) (*claims.RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})

	if _claims, ok := token.Claims.(*claims.RefreshClaims); ok && token.Valid {
		return _claims, nil
	} else {
		return nil, err
	}
}

// Refreshing 刷新token
func (j *Jwtx) Refreshing(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	accessClaims, _, ok, err := j.compareAccessAndRefresh(accessToken, refreshToken)
	if err != nil {
		return "", "", err
	}
	if !ok {
		return "", "", RefreshTokenErr
	}
	newAccessToken, newRefreshToken, err = j.CreateToken(accessClaims.BaseClaims)
	if err != nil {
		return "", "", err
	}
	return newAccessToken, newRefreshToken, nil
}

func (j *Jwtx) compareAccessAndRefresh(accessToken, refreshToken string) (*claims.AccessClaims, *claims.RefreshClaims, bool, error) {
	accessClaims, err := j.ParseAccessToken(accessToken)
	if err != nil {
		if !errors.Is(err, TokenExpired) {
			return nil, nil, false, err
		}
	}
	refreshClaims, err := j.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, nil, false, err
	}
	if accessClaims.Key == refreshClaims.Key {
		return accessClaims, refreshClaims, true, nil
	} else {
		return accessClaims, refreshClaims, false, nil
	}
}
