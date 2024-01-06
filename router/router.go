package router

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/controller"
)

func AddRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/user", controller.GetGithubProfile)
}
