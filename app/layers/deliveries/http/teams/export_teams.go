package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/teams/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) ExportTeams(c *gin.Context) {
	request, err := new(dtos.ExportTeamRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	teams, members, err := h.TeamsUseCase.ExportTeamMember(c.Request.Context(), request.ToEntity(), &entities.MemberFilter{})

	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.ExportAllTeamResponseJSON).Parse(c, teams, members)
	utils.JSONSuccessResponse(c, responseData)
	// metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	// utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)
}
