package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/usecase/users"
)

type Handler struct {
	UsersUC users.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, usersUC users.UseCase) {
	handler := &Handler{
		UsersUC: usersUC,
	}

	v1 := ginEngine.Group("v1")
	{
		v1.POST("/users.superAdmin", handler.CreateSuperAdmin)
	}
}
