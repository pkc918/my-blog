package controller

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/service"
	"my-blog/tool"
	"net/http"
)

/* client */

// GetGithubProfile 从Github获取信息
func GetGithubProfile(c *gin.Context) {
	githubProfile := service.GetGithubProfile()
	res := tool.Res{
		C:          c,
		Code:       10000,
		HttpStatus: http.StatusOK,
		Data:       &githubProfile,
		Msg:        "操作成功",
	}
	tool.Response(&res)
}

/* admin */

// LogIn 登录
func LogIn(c *gin.Context) {
	service.LogIn()
}
