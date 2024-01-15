package dao

import (
	"encoding/json"
	"log"
	"my-blog/app/dto"
	"my-blog/app/middleware"
	"my-blog/tool"
)

func GetGithubProfile() *dto.GithubProfileRes {
	var githubProfile *dto.GithubProfileRes
	res := tool.Fetch("https://api.github.com/users/pkc918")
	if err := json.Unmarshal(res, &githubProfile); err != nil {
		log.Fatal(err)
	}
	return githubProfile
}

func ValidateCredentials(data *dto.LoginParamsData) bool {
	return data.Phone == "18608645531" && data.Password == "123456" && data.Email == "2489964425@qq.com"
}

func Login(data *dto.LoginParamsData) string {
	jwt, err := middleware.GenerateJWT(data)
	if err != nil {
		panic(err.Error())
	}
	return jwt
}
