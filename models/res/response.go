package res

import (
	CODE "Goblog/models/res/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 封装响应
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	SUCCESS = 0
	ERR     = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// ResultOK 成功响应
func ResultOK(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func ResultOkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func ResultOkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

// ResultFail 失败响应
func ResultFail(data any, msg string, c *gin.Context) {
	Result(ERR, data, msg, c)
}

func ResultFailWithMsg(msg string, c *gin.Context) {
	Result(ERR, map[string]any{}, msg, c)
}

func ResultFailWithCode(code CODE.ErrorCode, c *gin.Context) {
	if msg, ok := CODE.ErrorMap[code]; ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}

	Result(ERR, map[string]any{}, "未知错误", c)

}
