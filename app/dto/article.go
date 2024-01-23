package dto

import "gorm.io/gorm"

/* Res */

type BlogList struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

/* Req */

type GetBlogListParams struct {
	PageSize    int `json:"pageSize"`
	CurrentPage int `json:"currentPage"`
}

type NewArticleParams struct {
	// binding:"required,dive" 当比如 title 是数组时，需要required验证，就需要使用 dive
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
	PWD         string `json:"pwd" binding:"required"`
}
