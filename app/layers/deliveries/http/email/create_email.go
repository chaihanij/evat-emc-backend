package email

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/email/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateEmail(c *gin.Context) {
	requset, err := new(dtos.CreateEmailRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	res, err := h.EmailUseCase.CreateEmail(c.Request.Context(), requset.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	responseData := new(dtos.CreateEmailResponseJSON).Parse(res)

	utils.JSONSuccessResponse(c, responseData)

}
