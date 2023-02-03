package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateAssignment(c *gin.Context) {
	request, err := new(dtos.CreateAssignmentRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.AssignmentsUseCase.CreateAssignment(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.CreateAssignmentResponseJSON).Parse(c, res)
	utils.JSONSuccessResponse(c, responseData)
}
