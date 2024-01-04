package service

import (
	"my-blog/app/dao"
	"my-blog/app/model"
)

func GetGithubProfile() *model.GithubProfile {
	githubProfile := dao.GetGithubProfile()
	return githubProfile
}
