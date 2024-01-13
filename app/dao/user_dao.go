package dao

import (
	"encoding/json"
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
