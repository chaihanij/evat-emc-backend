package announcements

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/announcements"
)

type Handler struct {
	AnnouncementUseCase announcements.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, announcementsUseCase announcements.UseCase) {
	handler := &Handler{
		AnnouncementUseCase: announcementsUseCase,
	}
	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/announcements", handler.FindAllAnnouncements)                      // Get All
		v1Auth.POST("/announcements", handler.CreateAnnouncements)                      // Create
		v1Auth.GET("/announcements/:announcement_uuid", handler.FindOneAnnouncement)    // GET One
		v1Auth.PUT("/announcements/:announcement_uuid", handler.UpdateAnnouncement)     // update
		v1Auth.DELETE("/announcements/:announcement_uuid", handler.DeleteAnnouncements) // delete
	}
}
