package members

// UpdateMemberPushDocument Update Member Push Document
// @Summary Update Member Push Document
// @Description API For Update Member  Push Document
// @ID post-member-documents
// @Accept mpfd
// @Produce json
// @Tags MEMBERS
// @Param Authorization header string true "for authentication"
// @Param member_uuid path string true "member_uuid of member"
// @Param document formData file true "file document upload"
// @Success 200 {object} dtos.UpdateMemberPushDocumentResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members/:member_uuid/documents [post]
func UpdateMemberPushDocument() {}
