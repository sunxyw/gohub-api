package v1

import (
	"gohub/app/models/category"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, gin.H{
			"category": categoryModel,
		})
	} else {
		response.ServerError(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Update(c *gin.Context) {

	// 验证 url 参数 id 是否正确
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.NotFound(c)
		return
	}

	// 表单验证
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	// 保存数据
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.SuccessWithData(c, gin.H{
			"category": categoryModel,
		})
	} else {
		response.ServerError(c)
	}
}

func (ctrl *CategoriesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	// TODO: implement pager
	data, _ := category.Paginate(c, 10)
	response.SuccessWithData(c, gin.H{
		"categories": data,
	})
}

func (ctrl *CategoriesController) Delete(c *gin.Context) {

	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.NotFound(c)
		return
	}

	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.ServerError(c, "删除失败，请稍后尝试~")
}
