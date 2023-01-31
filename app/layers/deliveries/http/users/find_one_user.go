package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/users/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FineOneUser(c *gin.Context) {
	request, err := new(dtos.FindOneUserRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	user, err := h.UsersUseCase.FindOneUser(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindOneUserResponseJSON).Parse(user)
	utils.JSONSuccessResponse(c, responseData)
}
