package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/users/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

// TODO
func (h *Handler) ChangePassword(c *gin.Context) {
	request, err := new(dtos.ChangePasswordRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	// err = h.UsersUseCase.ChangePassword(c.Request.Context(), request.ToEntity())
	// if err != nil {
	// 	utils.JSONErrorResponse(c, err)
	// 	return
	// }

	utils.JSONSuccessResponse(c, request.ToEntity())
}
