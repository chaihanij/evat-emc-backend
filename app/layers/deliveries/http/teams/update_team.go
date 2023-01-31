package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateTeam(c *gin.Context) {
	request, err := new(dtos.UpdateTeamRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	team, err := h.TeamsUseCase.UpdateTeam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindOneTeamResponseJSON).Parse(team)
	utils.JSONSuccessResponse(c, responseData)
}
