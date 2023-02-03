package members

// UpdateMemberPullDocument Update Member Pull Document
// @Summary Update Member Pull Document
// @Description API For Update Member Pull Document
// @ID delete-member-documents
// @Accept json
// @Produce json
// @Tags MEMBERS
// @Param Authorization header string true "for authentication"
// @Param member_uuid path string true "member_uuid of member"
// @Param document_uuid path string true "document_uuid of files"
// @Success 200 {object} dtos.UpdateMemberPullDocumentResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members/:member_uuid/documents/:document_uuid [delete]
func UpdateMemberPullDocument() {}
