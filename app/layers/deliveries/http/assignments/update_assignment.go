package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateAssignment(c *gin.Context) {
	request, err := new(dtos.UpdateAssignmentRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.AssignmentsUseCase.UpdateAssignment(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.UpdateAssignmentResponseJSON).Parse(c, res)
	utils.JSONSuccessResponse(c, responseData)
}
