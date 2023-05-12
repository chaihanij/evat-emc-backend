package config

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/config/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateConfig(c *gin.Context) {
	req, err := new(dtos.RequsetConfig).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	_, err = h.ConfigUseCase.UpdateConfigs(c.Request.Context(), req.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	utils.JSONSuccessResponse(c, req)

}
