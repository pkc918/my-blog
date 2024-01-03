package model

type GithubProfile struct {
	AvatarUrl string `json:"avatar_url"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
	Followers int    `json:"followers"`
}
