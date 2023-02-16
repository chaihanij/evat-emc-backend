package announcements

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/announcements/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) CreateAnnouncements(c *gin.Context) {
	request, err := new(dtos.CreateAnnouncementRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	res, err := h.AnnouncementUseCase.CreateAnnouncements(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.CreateAnnouncementResponseJSON).Parse(c, res)
	utils.JSONSuccessResponse(c, responseData)
}
