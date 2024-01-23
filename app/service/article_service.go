package service

import (
	"my-blog/app/dao"
	"my-blog/app/dto"
)

/* client */

// GetBlogs 获取博客列表
func GetBlogs(params *dto.GetBlogListParams) []*dto.BlogList {
	return dao.GetBlogs(params)
}

// GetArticleById 获取文章内容
func GetArticleById(params *dto.ArticleDetailParams) (*dto.NewArticleParams, error) {
	article, err := dao.GetArticleById(params)
	if err != nil {
		return nil, err
	}
	return article, nil
}

/* admin */

// PostNewArticle 写新文章
func PostNewArticle(article *dto.NewArticleParams) error {
	return dao.PostNewArticle(article)
}
