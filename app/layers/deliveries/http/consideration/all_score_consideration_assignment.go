package consideration

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/consideration/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) AllScoreConsiderationAssignment(c *gin.Context) {
	request, err := new(dtos.AllScoreConsiderationAssignmentRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	response, err := h.ConsiderationUseCase.AllScoreConsiderationAssignment(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.AllScoreConsiderationAssignmentResponseJSON).Parse(c, response)

	utils.JSONSuccessResponse(c, responseData)

}
