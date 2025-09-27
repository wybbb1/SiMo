package service

import (
	"math/rand"
	"net/http"
	"errors"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"go.uber.org/zap"

	"github.com/wybbb1/SiMo/internal/config"
	auth "github.com/wybbb1/SiMo/internal/model/auth"
	"github.com/wybbb1/SiMo/internal/repo"
	"github.com/wybbb1/SiMo/internal/log"
	login "github.com/wybbb1/SiMo/internal/model/err"
	jwt "github.com/wybbb1/SiMo/internal/util"
)

var state string
var githubOauthConfig = &oauth2.Config{
	ClientID:     config.Config.OAuth.ClientID,
	ClientSecret: config.Config.OAuth.ClientSecret,
	RedirectURL:  config.Config.OAuth.RedirectURL,
	Scopes:       config.Config.OAuth.Scopes,
	Endpoint:     github.Endpoint,
}

type LoginService struct {
	repo repo.LoginRepository
}

func NewLoginService(repo repo.LoginRepository) *LoginService {
	return &LoginService{repo: repo}
}

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
// OAuthCall 处理 OAuth 授权请求
func (s *LoginService) OAuthCall(c *gin.Context) error {
	state = randString(16) // 生成随机字符串以防止 CSRF 攻击

	// 生成授权 URL
	url := githubOauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)

	return nil
}
// OAuthCallback 处理 OAuth 回调逻辑
func (s *LoginService) OAuthCallback(c *gin.Context) error {
	stateFromCallback := c.Query("state")
	if stateFromCallback != state {
		log.Logger.Error("异常 CSRF 攻击",
			zap.String("ip", c.ClientIP()),
		)
		return errors.New("invalid state")
	}
	// 处理 OAuth 回调逻辑
	code := c.Query("code")
	if code == "" {
		return errors.New("code not found")
	}
	// 执行令牌交换
	ctx := c.Request.Context()
	token, err := githubOauthConfig.Exchange(ctx, code)
	if err != nil {
		return errors.New("failed to exchange token")
	}
	// 创建一个 HTTP 客户端来使用获取的令牌
	client := githubOauthConfig.Client(ctx, token)

	userResp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return errors.New("failed to get user info")
	}
	defer userResp.Body.Close()

	// 解析用户信息
	var userInfo auth.GithubInfo
	if err := json.NewDecoder(userResp.Body).Decode(&userInfo); err != nil {
		return errors.New("failed to parse user info")
	}

	return nil
}

func (s *LoginService) Login(username, password string) (string, error) {
	user := s.repo.GetUserByUsername(username)
	if (user == nil || user.Password != password) {
		return "", errors.New(login.AuthError)
	}

	claims := jwt.NewClaims(*user)
	token, err := jwt.GenerateToken(claims)
	if err != nil {
		log.Logger.Error("GenerateToken failed",
			zap.String("error", err.Error()),
			zap.String("username", username),
		)
		return "", err
	}

	// todo :写入缓存

	return token, nil
}