package fieldraceteams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/field_race_teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindAllFieldRacesTeam(c *gin.Context) {

	// dtos.FindAllField_race_teamsRequestJSON
	request, err := new(dtos.FindAllField_race_teamsRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	totalRecords, teams, err := h.Field_race_teams.FindFieldracesteam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindAllField_race_teamsResponseJSON).Parse(teams)
	metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)

}
