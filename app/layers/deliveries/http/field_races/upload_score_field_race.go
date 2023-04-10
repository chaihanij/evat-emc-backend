package fieldraces

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/field_races/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UploadScoreFieldRace(c *gin.Context) {
	request, err := new(dtos.UploadScoreFieldRace).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	response, err := h.TeamFieldRacesUseCase.UploadScoreFieldRace(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.UploadScoreFieldRaceResponseJSON).Parse(c, response)
	utils.JSONSuccessResponse(c, responseData)
}
