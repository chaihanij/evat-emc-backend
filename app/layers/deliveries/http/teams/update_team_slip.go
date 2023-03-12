package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateTeamSlip(c *gin.Context) {
	request, err := new(dtos.UpdateTeamSlipRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	teamUUID, file := request.ToEntity()
	team, err := h.TeamsUseCase.UpdateTeamSlip(c.Request.Context(), teamUUID, file)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindOneTeamResponseJSON).Parse(c, team)
	utils.JSONSuccessResponse(c, responseData)
}
