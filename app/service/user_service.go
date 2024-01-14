package service

import (
	"errors"
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
func LogIn(data *dto.LoginParamsData) (token string, err error) {
	if !dao.ValidateCredentials(data) {
		err = errors.New("账号密码不一致")
		return token, err
	}
	token = dao.Login(data)
	return token, nil
}
