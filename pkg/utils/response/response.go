package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 需跟前端協調成功碼和錯誤碼
const (
	SUCCESS = 0
	ERROR   = 7
	UNAUTH  = 10
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(c *gin.Context, httpCode int, code int, data interface{}, msg string) {
	c.JSON(httpCode, Response{code, data, msg})
}

func Ok(c *gin.Context) {
	Result(c, http.StatusOK, SUCCESS, map[string]interface{}{}, "操作成功")
}

func OkWithMessage(c *gin.Context, msg string) {
	Result(c, http.StatusOK, SUCCESS, map[string]interface{}{}, msg)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, http.StatusOK, SUCCESS, data, "操作成功")
}

func Fail(c *gin.Context) {
	Result(c, http.StatusOK, ERROR, map[string]interface{}{}, "操作失败")
}

func FailWithMessage(c *gin.Context, msg string) {
	Result(c, http.StatusOK, ERROR, map[string]interface{}{}, msg)
}

func FailWithDetailed(c *gin.Context, data interface{}, msg string) {
	Result(c, http.StatusOK, ERROR, data, msg)
}

func UnAuth(c *gin.Context, msg string) {
	Result(c, http.StatusUnauthorized, UNAUTH, map[string]interface{}{}, msg)
}
