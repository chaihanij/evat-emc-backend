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
	v1 := ginEngine.Group("v1")
	{
		v1.GET("/assignment/topic/:assignment_uuid", handler.FindTopic)
	}
	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/assignments", handler.FindAllAssignment)

		// v1Auth.GET("/assignment/topic/:assignment_uuid", handler.FindTopic)

		v1Auth.POST("/assignments", handler.CreateAssignment)

		//
		v1Auth.GET("/assignments/:assignment_uuid", handler.FindOneAssignment)
		v1Auth.PUT("/assignments/:assignment_uuid", handler.UpdateAssignment)
		v1Auth.DELETE("/assignments/:assignment_uuid", handler.DeleteAssignment)

		// assignments
		v1Auth.POST("/assignments/:assignment_uuid/image", handler.UpdateAssignmentImage)
		v1Auth.POST("/assignments/:assignment_uuid/document", handler.UpdateAssignmentDocument)

		v1Auth.GET("/assignment/team/:team_uuid", handler.FindTeamAssignment)

		v1Auth.POST("/assignment/:assignment_uuid/doc", handler.UploadScoreAssignment)

		v1Auth.PUT("/assignment/:assignment_uuid/updatescors", handler.UploadScoreAssignment)
		v1Auth.PUT("/assignment/:assignment_uuid/uploadfile", handler.UploadFileScore)

	}
}
