// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user_auth"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsIdentifierExist(c *gin.Context) {
	request := requests.SignupIdentifierExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupIdentifierExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user_auth.IsIdentifierExist(request.Identifier),
	})
}
