package consideration

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/consideration/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) AllScoreConsideration(c *gin.Context) {

	request, err := new(dtos.AllScoreRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	response, err := h.ConsiderationUseCase.AllScore(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	responseData := new(dtos.AllScoresResponseJSON).Parse(c, response)

	utils.JSONSuccessResponse(c, responseData)

}
