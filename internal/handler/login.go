package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	auth "github.com/wybbb1/SiMo/internal/model/err"
	user "github.com/wybbb1/SiMo/internal/model/user"
    model "github.com/wybbb1/SiMo/internal/model"
	login "github.com/wybbb1/SiMo/internal/service/login"
	"github.com/wybbb1/SiMo/internal/log"
)

type LoginHandler struct {
    service login.ILoginService
}

func NewLoginHandler(service login.ILoginService) *LoginHandler {
    return &LoginHandler{service: service}
}

func (h *LoginHandler) OAuthCall(c *gin.Context) {
    if err := h.service.OAuthCall(c); err != nil {

        log.Logger.Error("OAuthCall failed",
            zap.String("error", err.Error()),
            zap.String("ip", c.ClientIP()),
        )
        c.JSON(http.StatusOK, model.Error(auth.CallFailed))
    }

    log.Logger.Info("OAuthCall success",
        zap.String("ip", c.ClientIP()),
    )
    c.JSON(http.StatusOK, model.Success(auth.CallSuccess))
}

func (h *LoginHandler) OAuthCallback(c *gin.Context) {
    if err := h.service.OAuthCallback(c); err != nil {

        log.Logger.Error("OAuthLogin failed",
            zap.String("error", err.Error()),
            zap.String("ip", c.ClientIP()),
        )
        c.JSON(http.StatusOK, model.Error(auth.OAuthFailed))
    }

    log.Logger.Info("OAuthLogin success",
        zap.String("ip", c.ClientIP()),
    )
    c.JSON(http.StatusOK, model.Success(auth.OAuthSuccess))
}

func (h *LoginHandler) Login(c *gin.Context) {
    var loginUser user.User
    if err := c.ShouldBindJSON(&loginUser); err != nil {
        log.Logger.Error("Login bind json failed",
            zap.String("error", err.Error()),
            zap.String("ip", c.ClientIP()),
        )
        c.JSON(http.StatusOK, model.Error(auth.LoginFailed))
        return
    }

    if token, err := h.service.Login(loginUser.Username, loginUser.Password); err != nil {
        log.Logger.Error("Login failed",
            zap.String("error", err.Error()),
            zap.String("ip", c.ClientIP()),
        )
        c.JSON(http.StatusOK, model.Error(auth.AuthError))
    } else {
        log.Logger.Info("Login success",
            zap.String("ip", c.ClientIP()),
        )
        c.JSON(http.StatusOK, model.Success(token))
    }
}
    