package dto

/* Res */

// GithubProfileRes github用户信息
type GithubProfileRes struct {
	AvatarUrl string `json:"avatar_url"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
	Followers int    `json:"followers"`
}

/* Req */

// LoginParamsData 登录参数
type LoginParamsData struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
