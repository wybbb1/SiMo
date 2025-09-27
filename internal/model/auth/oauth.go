package model

type GithubInfo struct {
	Login    string `json:"login"`
	Avatar   string `json:"avatar_url"`
	Name     string `json:"name"`
	// 用户邮箱(需要设置为公开)
	Email    string `json:"email"`
	Html_url string `json:"html_url"`
}

const (
	CallSuccess = "OAuth Call Success"
	CallFailed  = "OAuth Call Failed"

	LoginSuccess = "OAuth Login Success"
	LoginFailed  = "OAuth Login Failed"
)
