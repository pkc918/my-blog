package dao

import (
	"fmt"
	"my-blog/app/dto"
)

func GetBlogs(params *dto.GetBlogListParams) []*dto.BlogList {
	// Mysql
	fmt.Println(params)
	var blogList []*dto.BlogList
	blogList = append(blogList, &dto.BlogList{
		Title:       "你好",
		Description: "我需求那得不到的东西，我得到了我所没有寻求的东西",
	})
	return blogList
}

func PostNewArticle(params *dto.NewArticleParams) error {
	// Mysql
	return nil
}

func GetArticleById(params *dto.ArticleDetailParams) (*dto.NewArticleParams, error) {
	// Mysql
	var article = &dto.NewArticleParams{
		Title:       "测试",
		Content:     "### 测试用的",
		Description: "I love three things",
		Author:      "陪我去看海吧",
		PWD:         "juejin",
	}
	return article, nil
}
