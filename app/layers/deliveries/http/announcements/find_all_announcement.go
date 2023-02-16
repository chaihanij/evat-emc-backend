package announcements

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/announcements/dtos"
	"gitlab.com/chaihanij/evat/app/utils"
)

func (h *Handler) FindAllAnnouncements(c *gin.Context) {
	request, err := new(dtos.FindAllAnnouncementRequestJSON).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	totalRecords, announcements, err := h.AnnouncementUseCase.FindAllAnnouncement(c.Request.Context(), request.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	responseData := new(dtos.FindAllAnnouncementResponseJSON).Parse(announcements)
	metaData := new(dtos.MetaDataResponse).Parse(request.Page, request.PageSize, totalRecords)
	utils.JSONSuccessCodeWithMetaDataResponse(c, responseData, metaData)

}
