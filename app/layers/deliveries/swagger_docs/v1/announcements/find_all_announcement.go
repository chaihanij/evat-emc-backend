package announcements

// FindAllAnnouncements Find All Announcements
// @Summary Find All Announcements
// @Description API For Find All Announcements
// @ID get-all-announcements
// @Accept json
// @Produce json
// @Tags ANNOUNCEMENTS
// @Param Authorization header string true "for authentication"
// @Param year query string false "year of EVAT eMC"
// @Param page query string false "Offset for search assignments"
// @Param pageSize query string false "PageSize of assignments"
// @Success 200 {object} dtos.FindAllAnnouncementResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/announcement [get]
func FindAllAnnouncement() {}
