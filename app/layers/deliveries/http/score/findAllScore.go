package score

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/score/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FinAllScore(c *gin.Context) {
	fmt.Println("/score")
	request, err := new(dtos.FindAllScoreRequesJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	totalRecords, score, err := h.ScoreUseCase.FindAllScore(c.Request.Context(), request.ToEntity())
	fmt.Println("totalRecords :", totalRecords)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	responseData := new(dtos.FindAllScoreResponseJSON).Parse(score)
	metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)

}
