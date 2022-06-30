// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub/app/models"
	"gohub/app/models/user_auth"
	"gohub/pkg/database"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	Rank   int    `json:"rank,omitempty"`

	Auths []user_auth.UserAuth `json:"auths"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
