package auth

import (
	"errors"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

// RefreshToken 刷新 Access Token
func (lc *LoginController) RefreshToken(c *gin.Context) {

	// TODO: implement refresh token
	// token, err := jwt.NewJWT().RefreshToken(c)
	token := "hello"
	err := errors.New("hello")

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func (lc *LoginController) Logout(c *gin.Context) {
	// TODO: implement revoke token
	// jwt.NewJWT().RevokeToken(c)
	response.Success(c)
}
