package util

import (
	"fmt"
	"time"

	"github.com/wybbb1/SiMo/internal/config"
	userModel "github.com/wybbb1/SiMo/internal/model/user"
	model "github.com/wybbb1/SiMo/internal/model/auth"

	"github.com/golang-jwt/jwt/v5"
)

func NewClaims(user userModel.User) *model.JWTClaims {
	return &model.JWTClaims{
		UserID:   user.ID, 
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Config.Jwt.Issuer,
			Subject:   user.Username,
			Audience:  jwt.ClaimStrings{config.Config.Jwt.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Config.Jwt.ExpireHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
}

func GenerateToken(claim *model.JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(config.Config.Jwt.SecretKey))
}

func ParseToken(tokenString string) (*model.JWTClaims, error) {
	claims := &model.JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法，防止算法混淆攻击
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名算法: %v", token.Header["alg"])
		}
		return []byte(config.Config.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok {
		return nil, fmt.Errorf("无效的token claims")
	}
	return claims, nil
}
