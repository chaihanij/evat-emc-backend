package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/users/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateActivateCode(c *gin.Context) {
	request, err := new(dtos.UpdateActivateCode).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.UsersUseCase.UpdateUser(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.UpdateUserResponseJSON).Parse(res)
	utils.JSONSuccessResponse(c, responseData)
}
