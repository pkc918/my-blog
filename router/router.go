package router

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/controller"
	"my-blog/app/middleware"
)

func AddRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/user", controller.GetGithubProfile)
	api.GET("/blogs", controller.GetBlogs)
	api.GET("/blogs/:id", controller.GetArticleById)
	api.GET("/photos", controller.GetPhotos)

	admin := r.Group("/admin")
	admin.POST("/login", controller.LogIn)
	admin.Use(middleware.JWT())
	admin.POST("/new", controller.PostNewArticle)

}
