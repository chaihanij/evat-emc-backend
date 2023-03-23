package assignmentteams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/utils"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/assignment_teams/dtos"

)

func (h *Handler) FindAllAssignmentTeamScore(c *gin.Context) {
	request, err := new(dtos.FindAllAssignmentRequestTeamJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	totalRecords, assignments, err := h.AssignmentTeamUseCase.FindAllAssignmentTeamscore(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindAllAssignmentTeamResponseJSON).Parse(assignments)
	metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)

}