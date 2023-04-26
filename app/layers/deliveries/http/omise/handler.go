package omise

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/usecase/teams"
)

type Handler struct {
	TeamsUseCase teams.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine,
	teamsUseCase teams.UseCase) {
	handler := &Handler{
		TeamsUseCase: teamsUseCase,
	}

	v1 := ginEngine.Group("v1")
	{
		v1.POST("/omise/web/hooks", handler.WebHooks)
	}
}
