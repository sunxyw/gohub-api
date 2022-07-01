// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
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

func (sc *SignupController) Signup(c *gin.Context) {
	// 验证表单
	request := requests.SignupRequest{}
	if ok := requests.Validate(c, &request, requests.Signup); !ok {
		return
	}

	// 创建用户
	userModel := user.User{
		Name: request.Name,
		Auths: []user_auth.UserAuth{
			{
				Type:       request.Type,
				Identifier: request.Identifier,
				Credential: request.Password,
			},
		},
	}
	userModel.Create()
	if userModel.ID > 0 {

		// 注册成功
		response.Created(c, gin.H{
			"user": userModel,
		})
		return
	}

	response.Abort500(c, "创建用户失败，请稍后再试~")
}
