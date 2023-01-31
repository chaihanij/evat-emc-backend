package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) DeleteTeam(c *gin.Context) {
	request, err := new(dtos.DeleteTeamRequest).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	err = h.TeamsUseCase.DeleteTeam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	utils.JSONSuccessResponse(c, nil)
}
