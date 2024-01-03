package dao

import (
	"encoding/json"
	"log"
	"my-blog/app/model"
	"my-blog/tool"
)

func GetGithubProfile() *model.GithubProfile {
	var githubProfile *model.GithubProfile
	res := tool.Fetch("https://api.github.com/users/pkc918")
	if err := json.Unmarshal(res, &githubProfile); err != nil {
		log.Fatal(err)
	}
	return githubProfile
}
