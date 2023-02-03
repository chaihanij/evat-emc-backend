package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) DeleteAssignment(c *gin.Context) {
	request, err := new(dtos.DeleteAssignmentRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	err = h.AssignmentsUseCase.DeleteAssignment(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	utils.JSONSuccessResponse(c, nil)

}
