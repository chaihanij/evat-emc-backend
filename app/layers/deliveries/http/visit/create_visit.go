package visit

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/visit/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateVisit(c *gin.Context) {
	request, err := new(dtos.CreateVisitRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	res, err := h.VisitUseCase.CreateVisit(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.CreateVisitResponseJSON).Parse(res)
	utils.JSONSuccessResponse(c, responseData)

}
