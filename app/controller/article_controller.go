package controller

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/dto"
	"my-blog/app/service"
	"my-blog/tool"
	"net/http"
)

/* client */

// GetBlogs 获取博客列表
func GetBlogs(c *gin.Context) {
	var params *dto.GetBlogListParams
	if err := c.ShouldBindJSON(&params); err != nil {
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
	service.GetBlogs(params)
}

// GetArticleById 获取文章内容
func GetArticleById(c *gin.Context) {
	service.GetArticleById()
}

/* admin */

// PostNewArticle 写新文章
func PostNewArticle(c *gin.Context) {
	var article *dto.NewArticleParams
	if err := c.ShouldBindJSON(&article); err != nil {
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
	if err := service.PostNewArticle(article); err != nil {
		tool.Fail(c, nil, err.Error())
		return
	}
	tool.Success(c, "添加成功")
}
