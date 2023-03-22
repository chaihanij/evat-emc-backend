package fieldraceteams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	fieldraceteams "gitlab.com/chaihanij/evat/app/layers/usecase/field_race_teams"
)

type Handler struct {
	Field_race_teams fieldraceteams.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine,
	authMiddleware middlewares.AuthMiddleware,
	fieldraceteamssUseCase fieldraceteams.UseCase) {
	handler := &Handler{
		Field_race_teams: fieldraceteamssUseCase,
	}

	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/field_race_teams", handler.FindAllFindAllFieldRaceTeams)
		v1Auth.POST("/fild_race_team/update", handler.CreateFieldRaceTeam)
	}
}
