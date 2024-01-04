package controller

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/service"
	"my-blog/tool"
	"net/http"
)

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
