package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindAllSearchTeam(c *gin.Context) {
	request, err := new(dtos.FindAllTeamRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	totalRecords, teams, err := h.TeamsUseCase.FindAllSearchTeam(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindAllTeamSearchResponseJSON).ParseTeamSearch(teams)
	metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)
}
