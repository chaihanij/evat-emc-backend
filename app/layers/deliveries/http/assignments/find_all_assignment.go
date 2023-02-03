package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignments/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindAllAssignment(c *gin.Context) {
	request, err := new(dtos.FindAllAssignmentRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	totalRecords, assignments, err := h.AssignmentsUseCase.FindAllAssignment(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindAllAssignmentResponseJSON).Parse(assignments)
	metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)

}
