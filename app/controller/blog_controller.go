package controller

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/service"
)

/* client */

// GetBlogs 获取博客列表
func GetBlogs(c *gin.Context) {
	service.GetBlogs()
}

// GetArticleById 获取文章内容
func GetArticleById(c *gin.Context) {
	service.GetArticleById()
}

/* admin */

// PostNewArticle 写新文章
func PostNewArticle(c *gin.Context) {
	service.PostNewArticle()
}
