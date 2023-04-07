package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UploadScoreAssignment(c *gin.Context) {

	request, err := new(dtos.UploadScoreAssingment).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.AssignmentsUseCase.UploadScoreAssignment(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseDat := new(dtos.UploadScoreAssignmentResponseJSON).Parse(c, res)
	utils.JSONSuccessResponse(c, responseDat)

}
