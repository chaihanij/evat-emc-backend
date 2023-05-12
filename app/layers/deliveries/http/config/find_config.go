package config

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/config/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindOneConfig(c *gin.Context) {
	// fmt.Println("request :", c)
	request, err := new(dtos.ConfigReq).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	config, err := h.ConfigUseCase.FindOneConfig(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.ConfigRes).Parse(c, config)

	utils.JSONSuccessResponse(c, responseData)

}
