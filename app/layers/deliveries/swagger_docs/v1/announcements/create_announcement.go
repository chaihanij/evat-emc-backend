package announcements

// CreateAnnouncement CreateAnnouncement
// @Summary Create Announcement
// @Description API For Create Announcement
// @ID post-assignment
// @Accept json
// @Produce json
// @Tags ANNOUNCEMENTS
// @Param Authorization header string true "for authentication"
// @Param announcement_uuid path string true "uuid of announcements"
// @Param body body dtos.CreateAnnouncementRequestJSON true "All params related to announcements"
// @Success 200 {object} dtos.CreateAnnouncementResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/announcement [post]
func CreateAnnouncement() {}
