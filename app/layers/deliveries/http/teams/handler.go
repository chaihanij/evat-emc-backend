package teams

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/teams"
)

type Handler struct {
	TeamsUseCase teams.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine,
	authMiddleware middlewares.AuthMiddleware,
	teamsUseCase teams.UseCase) {
	handler := &Handler{
		TeamsUseCase: teamsUseCase,
	}

	v1 := ginEngine.Group("v1")
	{
		v1.POST("/register-teams", handler.RegisterTeam)
		v1.POST("/teams/:team_uuid/slip", handler.UpdateTeamSlip)
		v1.GET("/teams/export", handler.ExportTeams)
	}

	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/teams", handler.FinAllTeam)
		v1Auth.POST("/teams", handler.CreateTeam)
		//
		v1Auth.GET("/teams/:team_uuid", handler.FinOneTeam)
		v1Auth.PUT("/teams/:team_uuid", handler.UpdateTeam)
		v1Auth.DELETE("/teams/:team_uuid", handler.DeleteTeam)

		//
		v1Auth.GET("/teams/:team_uuid/assignments/:assignment_uuid", handler.FindOneAssignmentTeam)
		v1Auth.POST("/teams/:team_uuid/assignments/:assignment_uuid", handler.SendAssignmentTeam)
		v1Auth.POST("/teams/:team_uuid/assignments/:assignment_uuid/documents/:topic", handler.SendAssignmentTeamPushDocument)
		v1Auth.DELETE("/teams/:team_uuid/assignments/:assignment_uuid/documents/:document_uuid", handler.SendAssignmentTeamPullhDocument)
		//search
		// v1Auth.GET("/teams/search", handler.FindAllSearchTeam)
	}
}
