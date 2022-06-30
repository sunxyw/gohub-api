// Package auth 处理用户注册、登录、密码重置
package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}
