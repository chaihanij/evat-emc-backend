package members

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/members"
)

type Handler struct {
	MemberUseCase members.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, membersUseCase members.UseCase) {
	handler := &Handler{
		MemberUseCase: membersUseCase,
	}
	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/members", handler.FindAllMember)
		v1Auth.POST("/members", handler.CreateMember)

		//
		v1Auth.GET("/members/:member_uuid", handler.FindOneMember)
		v1Auth.PUT("/members/:member_uuid", handler.UpdateMember)
		v1Auth.DELETE("/members/:member_uuid", handler.DeleteMember)

		// members
		v1Auth.POST("/members/:member_uuid/image", handler.UpdateMemberImage)
		v1Auth.POST("/members/:member_uuid/documents", handler.UpdateMemberPushDocument)
		v1Auth.POST("/members/:member_uuid/documents/:document_uuid", handler.UpdateMemberPullDocument)
	}
}
