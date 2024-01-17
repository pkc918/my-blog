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
