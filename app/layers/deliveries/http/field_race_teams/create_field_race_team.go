package fieldraceteams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/field_race_teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateFieldRaceTeam(c *gin.Context) {
	request, err := new(dtos.CreateFieldRaceTeamRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.Field_race_teams.CreateFieldRaceTeam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.CreateFieldRaceTeamResponsJSON).Parse(res)
	utils.JSONSuccessResponse(c, responseData)
}
