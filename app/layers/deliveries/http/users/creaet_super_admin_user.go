package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/users/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateSuperAdmin(c *gin.Context) {
	request, err := new(dtos.CreateSuperAdminRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.UsersUC.CreateUser(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.CreateSuperAdminResponseJSON).Parse(res)
	utils.JSONSuccessResponse(c, responseData)
}
