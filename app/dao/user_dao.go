package dao

import (
	"encoding/json"
	"fmt"
	"log"
	"my-blog/app/dto"
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
	return data.Phone == "123" && data.Password == "123"
}

func Login(data *dto.LoginParamsData) string {
	return fmt.Sprintf("token %s-%s-%s", data.Email, data.Phone, data.Password)
}
