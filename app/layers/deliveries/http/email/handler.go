package email

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/usecase/email"
)

type Handler struct {
	EmailUseCase email.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, emailUseCase email.UseCase) {
	handler := &Handler{

		EmailUseCase: emailUseCase,
	}

	v1 := ginEngine.Group("v1")
	{
		v1.POST("/email", handler.CreateEmail)
	}

}
