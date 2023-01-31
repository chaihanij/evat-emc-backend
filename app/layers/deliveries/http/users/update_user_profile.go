package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/users/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateUserProfile(c *gin.Context) {
	request, err := new(dtos.UpdateUserProfileRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.UsersUseCase.UpdateUser(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindOneUserProfileResponseJSON).Parse(res)
	utils.JSONSuccessResponse(c, responseData)
}
