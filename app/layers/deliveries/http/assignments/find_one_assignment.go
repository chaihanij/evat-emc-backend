package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindOneAssignment(c *gin.Context) {
	request, err := new(dtos.FindOneAssignmentRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	assignment, err := h.AssignmentsUseCase.FindOneAssignment(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindOneAssignmentResponseJSON).Parse(c, assignment)
	utils.JSONSuccessResponse(c, responseData)

}
