package service

import (
	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	// OAuth
	OAuthCall(c *gin.Context) error
	OAuthCallback(c *gin.Context) error
	// JWTLOGIN
	Login(username, password string) (string, error)
}