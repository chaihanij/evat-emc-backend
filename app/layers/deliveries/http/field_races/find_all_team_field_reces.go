package fieldraces

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/field_races/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindAllTeamFieldRaces(c *gin.Context) {
	request, err := new(dtos.FindAllTeamFieldRacestRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	totalRecords, assignments, err := h.TeamFieldRacesUseCase.FindAllTeamFieldRace(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindAllTeamFieldRecestResponseJSON).Parse(assignments)
	metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)

}
