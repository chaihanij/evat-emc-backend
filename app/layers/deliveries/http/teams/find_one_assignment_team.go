package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindOneAssignmentTeam(c *gin.Context) {
	request, err := new(dtos.FindOneAssignmentTeamRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	assignmentTeam, err := h.TeamsUseCase.FindOneAssignmentTeam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindOneAssignmentTeamResponseJSON).Parse(c, assignmentTeam)
	utils.JSONSuccessResponse(c, responseData)
}
