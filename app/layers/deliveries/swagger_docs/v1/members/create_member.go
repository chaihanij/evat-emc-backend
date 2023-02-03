package members

// CreateMember Create Member
// @Summary Create Member
// @Description API For Create Member
// @ID post-member
// @Accept json
// @Produce json
// @Tags MEMBERS
// @Param Authorization header string true "for authentication"
// @Param uuid path string true "uuid of member"
// @Param body body dtos.CreateMemberRequestJSON true "All params related to members"
// @Success 200 {object} dtos.CreateMemberResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members [post]
func CreateMember() {}
