package members

// DeleteMember Delete Member
// @Summary Delete Member
// @Description API For Delete Member
// @ID delete-member
// @Accept json
// @Produce json
// @Tags MEMBERS
// @Param Authorization header string true "for authentication"
// @Param member_uuid path string true "member_uuid of member"
// @Success 200 {object} dtos.DeleteMemberResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members/:member_uuid [delete]
func DeleteMember() {}
