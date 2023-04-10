package fieldraces

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	fieldraceteams "gitlab.com/chaihanij/evat/app/layers/usecase/field_races"
)

type Handler struct {
	TeamFieldRacesUseCase fieldraceteams.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, fieldraceteamsUseCase fieldraceteams.UseCase) {
	handler := &Handler{
		TeamFieldRacesUseCase: fieldraceteamsUseCase,
	}
	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/fieldraces/:uuid", handler.FindAllTeamFieldRaces)
		v1Auth.PUT("/fieldraces/:uuid/updatescore", handler.UploadScoreFieldRace)
	}
}
