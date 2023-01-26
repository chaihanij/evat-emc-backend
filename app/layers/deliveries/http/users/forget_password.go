package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/users/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

// Todo
func (h *Handler) ForgetPassword(c *gin.Context) {
	request, err := new(dtos.ForgotPasswordRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	// err = h.UsersUseCase.ForgetPassword(c.Request.Context(), request.ToEntity())
	// if err != nil {
	// 	utils.JSONErrorResponse(c, err)
	// 	return
	// }

	utils.JSONSuccessResponse(c, request.ToEntity())
}
