package emailcontact

import (
	"github.com/gin-gonic/gin"
	emailcontact "gitlab.com/chaihanij/evat/app/layers/usecase/emailcontact"
)

type Handler struct {
	EmailContact emailcontact.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, emailContactUseCase emailcontact.UseCase) {

	handler := &Handler{
		EmailContact: emailContactUseCase,
	}

	v1 := ginEngine.Group("v1")
	{
		v1.POST("/emailcontact", handler.CreateEmailContact)
	}

}
