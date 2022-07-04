package policies

import (
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

// func CanViewAuth(c *gin.Context, authModel auth.Auth) bool {}
// func CanCreateAuth(c *gin.Context, authModel auth.Auth) bool {}
// func CanUpdateAuth(c *gin.Context, authModel auth.Auth) bool {}
// func CanDeleteAuth(c *gin.Context, authModel auth.Auth) bool {}

var allowRegister bool
var allowTypes []string

func init() {
	allowRegister = config.Get[bool]("auth.register.enable")
	allowTypes = config.Get[[]string]("auth.login_types")
}

func CanRegisterUsing(c *gin.Context, method string) bool {
	return allowRegister && lo.Contains(allowTypes, method)
}
