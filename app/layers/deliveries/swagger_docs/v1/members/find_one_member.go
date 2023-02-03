package members

// FindOneMember Find One Member
// @Summary Find One Member
// @Description API For Find One Member
// @ID get-one-member
// @Accept json
// @Produce json
// @Tags MEMBERS
// @Param Authorization header string true "for authentication"
// @Param uuid path string true "uuid of member"
// @Success 200 {object} dtos.FindAllMemberResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members [get]
func FindOneMember() {}
