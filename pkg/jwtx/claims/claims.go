package claims

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ClaimType string

const (
	Access  ClaimType = "access"
	Refresh ClaimType = "refresh"
)

type AccessClaims struct {
	Key uuid.UUID
	jwt.RegisteredClaims
	BaseClaims
}

type BaseClaims struct {
	UUID     uuid.UUID
	ID       uint
	Username string
	Roles    []string
}

type RefreshClaims struct {
	Key uuid.UUID
	jwt.RegisteredClaims
}
