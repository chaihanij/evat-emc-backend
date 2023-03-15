package score

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/score"
	// "gitlab.com/chaihanij/evat/app/layers/usecase/score"
)

type Handler struct {
	ScoreUseCase score.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, scoreUseCase score.UseCase) {

	Handler := &Handler{
		ScoreUseCase: scoreUseCase,
	}
	// v1 := ginEngine.Group("v1")

	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/score", Handler.FinAllScore)
	}

}
