package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateAssignmentImage(c *gin.Context) {
	request, err := new(dtos.UpdateAssignmentImageRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.AssignmentsUseCase.UpdateAssignmentImage(c.Request.Context(), request.AssignmentUUID, request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FileResponse).Parse(c, res)
	utils.JSONSuccessResponse(c, responseData)
}
