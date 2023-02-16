package announcements

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/announcements/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) UpdateAnnouncement(c *gin.Context) {
	request, err := new(dtos.UpdateAnnouncementRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.AnnouncementUseCase.UpdateAnnouncements(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.UpdateAnnouncementResponseJSON).Parse(c, res)
	utils.JSONSuccessResponse(c, responseData)
}
