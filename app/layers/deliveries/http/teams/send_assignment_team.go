package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) SendAssignmentTeam(c *gin.Context) {
	request, err := new(dtos.SendAssignmentTeamRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	assignmentTeam, err := h.TeamsUseCase.SendAssignmentTeam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.SendAssignmentTeamResponseJSON).Parse(assignmentTeam)
	utils.JSONSuccessResponse(c, responseData)
}
