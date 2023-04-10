package consideration

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/consideration/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FinConsiderationFieldRaceTeam(c *gin.Context) {
	request, err := new(dtos.ConsiderationFieldRaceTeamRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	consideration, err := h.ConsiderationUseCase.FindConsiderationFieldRaceTeam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	responseData := new(dtos.ConsiderationFieldRaceTeamResponeJSON).Parse(c, consideration)
	utils.JSONSuccessResponse(c, responseData)

}
