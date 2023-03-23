package assignmentteams


import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/assignment_teams"
)
type Handler struct {
	AssignmentTeamUseCase assignmentteams.UseCase
}
func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, assignmentTeamsUseCase assignmentteams.UseCase) {
	handler := &Handler{
		AssignmentTeamUseCase: assignmentTeamsUseCase,
	}
	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/assignmentteam/:team_uuid", handler.FindAllAssignmentTeamScore)
	}
}