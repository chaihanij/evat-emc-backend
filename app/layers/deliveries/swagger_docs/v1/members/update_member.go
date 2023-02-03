package members

// UpdateMember Update Member
// @Summary Update Member
// @Description API For Update Member
// @ID put-member
// @Accept json
// @Produce json
// @Tags MEMBERS
// @Param Authorization header string true "for authentication"
// @Param member_uuid path string true "member_uuid of member"
// @Param body body dtos.UpdateMemberRequest true "All params related to members"
// @Success 200 {object} dtos.UpdateMemberResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members/:uuid [put]
func UpdateMember() {}
