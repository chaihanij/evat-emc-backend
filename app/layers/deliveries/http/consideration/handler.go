package consideration

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/consideration"
)

type Handler struct {
	ConsiderationUseCase consideration.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, considerationUseCase consideration.UseCase) {
	handler := &Handler{
		ConsiderationUseCase: considerationUseCase,
	}
	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/consideration/:consideration_uuid", handler.FindOneConsideration)

	}
}
