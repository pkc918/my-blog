package controller

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/dto"
	"my-blog/app/service"
	"my-blog/tool"
	"net/http"
)

/* client */

// GetGithubProfile 从Github获取信息
func GetGithubProfile(c *gin.Context) {
	githubProfile := service.GetGithubProfile()
	tool.Success(c, &githubProfile)
}

/* admin */

// LogIn 登录
func LogIn(c *gin.Context) {
	var loginParams *dto.LoginParamsData
	// 是否有参数
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		res := tool.Res{
			C:          c,
			Code:       10000,
			HttpStatus: http.StatusBadRequest,
			Data:       nil,
			Msg:        err.Error(),
		}
		tool.Response(&res)
		return
	}
	// 参数格式验证
	token, err := service.LogIn(loginParams)
	if err != nil {
		tool.Fail(c, nil, err.Error())
		return
	}
	tool.Success(c, token)
}
