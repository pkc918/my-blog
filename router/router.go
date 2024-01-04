package router

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/controller"
)

func AddRoutes(r *gin.Engine) {
	r.GET("/user", controller.GetGithubProfile)
}
