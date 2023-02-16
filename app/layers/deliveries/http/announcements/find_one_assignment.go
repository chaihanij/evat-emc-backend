package announcements

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/announcements/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindOneAnnouncement(c *gin.Context) {
	request, err := new(dtos.FindOneAnnouncementRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	assignment, err := h.AnnouncementUseCase.FindOneAnnouncements(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindOneAnnouncementResponseJSON).Parse(c, assignment)
	utils.JSONSuccessResponse(c, responseData)

}
