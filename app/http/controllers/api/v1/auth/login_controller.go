package auth

import (
	"errors"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/logger"
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
		logger.LogIf(err)
		response.Fail(c, "令牌刷新失败")
	} else {
		response.SuccessWithData(c, gin.H{
			"token": token,
		})
	}
}

func (lc *LoginController) Logout(c *gin.Context) {
	// TODO: implement revoke token
	// jwt.NewJWT().RevokeToken(c)
	response.Success(c)
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	user, err := auth.Attempt(request.Identifier, request.Password)

	if err != nil {
		response.Unauthorized(c, "账号或密码错误")
	} else {
		token := jwt.IssueToken(user.GetStringID())
		response.SuccessWithData(c, gin.H{
			"token": token,
		})
	}
}
