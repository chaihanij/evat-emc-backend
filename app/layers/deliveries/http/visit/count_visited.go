package visit

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/visit/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CountVisit(c *gin.Context) {

	res, err := h.VisitUseCase.FindVisit(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	responseData := new(dtos.VisitResponseJSON).Parse(res)
	utils.JSONSuccessResponse(c, responseData)

}
