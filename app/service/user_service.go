package service

import (
	"my-blog/app/dao"
	"my-blog/app/dto"
)

/* client */

// GetGithubProfile 从Github获取信息
func GetGithubProfile() *dto.GithubProfileRes {
	githubProfile := dao.GetGithubProfile()
	return githubProfile
}

/* admin */

// LogIn 登录
func LogIn() {

}
