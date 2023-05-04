package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/users/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) Login(c *gin.Context) {
	request, err := new(dtos.LoginRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.UsersUseCase.Login(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	ip := entities.LastLogin{IP: c.ClientIP(), Email: request.Email}
	_, err = h.UsersUseCase.CreatLastLogin(c.Request.Context(), &ip)

	responseData := new(dtos.LoginResponseJSON).Parse(res)
	utils.JSONSuccessResponse(c, responseData)
}
