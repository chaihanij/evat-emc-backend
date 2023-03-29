package visit

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/visit"
)

type Handler struct {
	VisitUseCase visit.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, visitUseCase visit.UseCase) {

	handler := &Handler{
		VisitUseCase: visitUseCase,
	}
	v1 := ginEngine.Group("v1")

	{
		v1.GET("/visit", handler.CountVisit)
		v1.GET("/create/visit", handler.CreateVisit)
	}

}
