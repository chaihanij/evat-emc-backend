package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) RegisterTeam(c *gin.Context) {
	request, err := new(dtos.RegisterTeamRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	team, user := request.ToEntity()
	team, user, charge, err := h.TeamsUseCase.RegisterTeam(c.Request.Context(), team, user)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res := new(dtos.RegisterTeamResponseJSON).Parse(c, team, user, charge)
	utils.JSONSuccessResponse(c, res)
}
