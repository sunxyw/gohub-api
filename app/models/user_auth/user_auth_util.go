package user_auth

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (userAuth UserAuth) {
	database.DB.Where("id", idstr).First(&userAuth)
	return
}

func GetBy(field, value string) (userAuth UserAuth) {
	database.DB.Where("? = ?", field, value).First(&userAuth)
	return
}

func All() (userAuths []UserAuth) {
	database.DB.Find(&userAuths)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UserAuth{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (userAuths []UserAuth, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(UserAuth{}),
		&userAuths,
		app.V1URL(database.TableName(&UserAuth{})),
		perPage,
	)
	return
}

// GetByIdentifier 通过 Identifier 来获取用户
func GetByIdentifier(identifier string) (userAuth UserAuth) {
	return GetBy("identifier", identifier)
}

// IsIdentifierExist 判断 Identifier 已被注册
func IsIdentifierExist(identifier string) bool {
	return IsExist("identifier", identifier)
}
