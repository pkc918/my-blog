package router

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/service"
)

func AddRoutes(r *gin.Engine) {
	r.GET("/user", service.GetGithubProfile)
}
