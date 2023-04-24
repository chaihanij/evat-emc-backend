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
		v1Auth.GET("/consideration/assignment/:assignment_UUID/:id", handler.FindOneConsideration)
		v1Auth.GET("/consideration/field_race/:rield_race_UUID/:id", handler.FinConsiderationFieldRaceTeam)

		v1Auth.GET("/consideration/all", handler.AllScoreConsideration)
		v1Auth.GET("/consideration/allscore/assignment/:assignment_UUID", handler.AllScoreConsiderationAssignment)
	}
}
