package v1

import (
	"gohub/app/models/link"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.SuccessWithData(c, gin.H{
		"links": link.AllCached(),
	})
}
