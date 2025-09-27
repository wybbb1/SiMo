package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/wybbb1/SiMo/internal/config"
	jwt "github.com/wybbb1/SiMo/internal/util"
	response "github.com/wybbb1/SiMo/internal/model"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader(config.Config.Jwt.Subject)
		// 如果 Authorization 头部信息为空，或者格式不正确，或者 token 为空，则返回错误
		if auth == "" || auth == "null" || auth == "undefined" {
			// 如果 Authorization 头部信息为空，或者格式不正确，或者 token 为空，则返回错误
			ctx.JSON(http.StatusOK, response.ErrorWithCode(401, "账号未登录"))
			ctx.Abort()
			return
		}


		tokenString, err := jwt.ParseToken(strings.TrimSpace(auth))
		if err != nil || tokenString == nil {
			ctx.JSON(http.StatusOK, response.ErrorWithCode(401, "无效的token"))
			ctx.Abort()
			return
		}
		ctx.Set("userid", tokenString.UserID)
		ctx.Next()
	}
}