package assignments

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/assignments"
)

type Handler struct {
	AssignmentsUseCase assignments.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, assignmentsUseCase assignments.UseCase) {
	handler := &Handler{
		AssignmentsUseCase: assignmentsUseCase,
	}
	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/assignments", handler.FindAllAssignment)
		v1Auth.POST("/assignments", handler.CreateAssignment)

		//
		v1Auth.GET("/assignments/:assignment_uuid", handler.FindOneAssignment)
		v1Auth.PUT("/assignments/:assignment_uuid", handler.UpdateAssignment)
		v1Auth.DELETE("/assignments/:assignment_uuid", handler.DeleteAssignment)

		// assignments
		v1Auth.POST("/assignments/:assignment_uuid/image", handler.UpdateAssignmentImage)
		v1Auth.POST("/assignments/:assignment_uuid/document", handler.UpdateAssignmentDocument)

		v1Auth.GET("/assignment/team/:team_uuid", handler.FindTeamAssignment)
	}
}
