package ginx

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

// GetToken Get jwt token from header (Authorization: Bearer xxx)
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// ParseJSON Parse body json data to struct
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.New(fmt.Sprintf("Parse request json failed: %s", err.Error()))
	}
	return nil
}

// ParseParamID Param returns the value of the URL param
func ParseParamID(c *gin.Context, key string) uint64 {
	val := c.Param(key)
	id, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// ParseParam Param returns the value of the URL param
func ParseParam(c *gin.Context, key string) string {
	val := c.Param(key)
	return val
}

// ParseQuery Parse query parameter to struct
func ParseQuery(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.New(fmt.Sprintf("Parse request query failed: %s", err.Error()))
	}
	return nil
}

// ParseForm Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.New(fmt.Sprintf("Parse request form failed: %s", err.Error()))
	}
	return nil
}

// SetUserID 將UserID寫入上下文
func SetUserID(c *gin.Context, userID uint) {
	c.Set("UserID", userID)
}

// GetUserID 獲取UserID
func GetUserID(c *gin.Context) uint {
	userID, _ := c.Get("UserID")
	return userID.(uint)
}

// SetUserUUID 將User UUID寫入上下文
func SetUserUUID(c *gin.Context, uuid uuid.UUID) {
	c.Set("UserUUID", uuid)
}

// GetUserUUID 獲取User UUID
func GetUserUUID(c *gin.Context) uuid.UUID {
	userUUID, _ := c.Get("UserUUID")
	return userUUID.(uuid.UUID)
}

// SetUserName 將user name寫入上下文
func SetUserName(c *gin.Context, username string) {
	c.Set("UserName", username)
}

// GetUserName 獲取user name
func GetUserName(c *gin.Context) string {
	username, _ := c.Get("UserName")
	return username.(string)
}

// SetRole 將role寫入上下文
func SetRole(c *gin.Context, roles []string) {
	c.Set("Roles", roles)
}

// GetRole 獲取role
func GetRole(c *gin.Context) []string {
	roles, _ := c.Get("Roles")
	return roles.([]string)
}
