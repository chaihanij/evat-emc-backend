package announcements

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/announcements/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) DeleteAnnouncements(c *gin.Context) {
	request, err := new(dtos.DeleteAnnouncementRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	err = h.AnnouncementUseCase.DeleteAnnouncements(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	utils.JSONSuccessResponse(c, nil)

}
