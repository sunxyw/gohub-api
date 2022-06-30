//Package user_auth 模型
package user_auth

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
	"os/user"
)

type UserAuth struct {
	models.BaseModel

	UserID     uint64 `json:"user_id,omitempty"`
	Type       string `json:"type,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Credential string `json:"-"`
	Verified   bool   `json:"verified,omitempty"`

	User user.User `json:"user"`

	models.CommonTimestampsField
}

func (userAuth *UserAuth) Create() {
	database.DB.Create(&userAuth)
}

func (userAuth *UserAuth) Save() (rowsAffected int64) {
	result := database.DB.Save(&userAuth)
	return result.RowsAffected
}

func (userAuth *UserAuth) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&userAuth)
	return result.RowsAffected
}

// ComparePassword 密码是否正确
func (userAuth *UserAuth) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userAuth.Credential)
}
