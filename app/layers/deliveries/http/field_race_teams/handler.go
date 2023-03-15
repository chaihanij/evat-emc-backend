package fieldraceteams


import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/field_race_teams"
)


type Handler struct {
	Field_race_teams  fieldraceteams.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine,
	authMiddleware middlewares.AuthMiddleware,
	fieldraceteamssUseCase fieldraceteams.UseCase) {
	handler := &Handler{
		Field_race_teams: fieldraceteamssUseCase,
	}



	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/field_races/teams", handler.FindAllFindAllField_race_teams)
	}
}
