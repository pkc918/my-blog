package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	C          *gin.Context
	HttpStatus int
	Code       int
	Data       interface{}
	Msg        string
}

func Response(res *Res) {
	res.C.JSON(res.HttpStatus, gin.H{
		"code": res.Code,
		"data": res.Data,
		"msg":  res.Msg,
	})
}

// Success 成功响应信息
func Success(c *gin.Context, data interface{}) {
	res := Res{
		C:          c,
		HttpStatus: http.StatusOK,
		Code:       10000,
		Data:       data,
		Msg:        "操作成功",
	}
	Response(&res)
}

func Fail(c *gin.Context, data interface{}, msg string) {
	res := Res{
		C:          c,
		HttpStatus: http.StatusOK,
		Code:       10001,
		Data:       data,
		Msg:        msg,
	}
	Response(&res)
}
