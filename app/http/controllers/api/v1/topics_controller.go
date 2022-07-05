package v1

import (
	"gohub/app/models/topic"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, gin.H{
			"topic": topicModel,
		})
	} else {
		response.ServerError(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.NotFound(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Forbidden(c)
		return
	}

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		response.SuccessWithData(c, gin.H{
			"topic": topicModel,
		})
	} else {
		response.ServerError(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Delete(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.NotFound(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Forbidden(c)
		return
	}

	rowsAffected := topicModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.ServerError(c, "删除失败，请稍后尝试~")
}

func (ctrl *TopicsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	// TODO: implement pager
	data, _ := topic.Paginate(c, 10)
	response.SuccessWithData(c, gin.H{
		"topics": data,
	})
}

func (ctrl *TopicsController) Show(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.NotFound(c)
		return
	}
	response.SuccessWithData(c, gin.H{
		"topic": topicModel,
	})
}
