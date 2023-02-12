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
	_, _, err = h.TeamsUseCase.RegisterTeam(c.Request.Context(), team, user)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	utils.JSONSuccessResponse(c, nil)
}
