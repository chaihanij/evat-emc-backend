package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindTopic(c *gin.Context) {

	request, err := new(dtos.FindTopicRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	topic, err := h.AssignmentsUseCase.FindTopicAssignment(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	responseData := new(dtos.ExportAssignmentTopicResponseJSON).Parse(c, topic)
	utils.JSONSuccessResponse(c, responseData)

}
