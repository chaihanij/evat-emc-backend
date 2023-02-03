package assignments

// UpdateAssignmentImage Update Assignment Image
// @Summary Update Assignment Image
// @Description API For Update Assignment Image
// @ID post-assignment-image
// @Accept mpfd
// @Produce json
// @Tags ASSIGNMENTS
// @Param Authorization header string true "for authentication"
// @Param assignment_uuid path string true "assignment_uuid of assignments"
// @Param image formData file true "file image upload"
// @Success 200 {object} dtos.UpdateAssignmentImageResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/assignments/:assignment_uuid/image [post]
func UpdateAssignmentImage() {}
