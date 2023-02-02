package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FinOneTeam(c *gin.Context) {
	request, err := new(dtos.FindOneTeamRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	team, err := h.TeamsUseCase.FindOneTeam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindOneTeamResponseJSON).Parse(c, team)
	utils.JSONSuccessResponse(c, responseData)
}
