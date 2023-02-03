package assignments

// UpdateAssignmentDocument Update Assignment Document
// @Summary Update Assignment Document
// @Description API For Update Assignment Document
// @ID post-assignment-document
// @Accept mpfd
// @Produce json
// @Tags ASSIGNMENTS
// @Param Authorization header string true "for authentication"
// @Param assignment_uuid path string true "assignment_uuid of assignments"
// @Param document formData file true "file document upload"
// @Success 200 {object} dtos.UpdateAssignmentDocumentResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/assignments/:assignment_uuid/document [post]
func UpdateAssignmentDocument() {}
