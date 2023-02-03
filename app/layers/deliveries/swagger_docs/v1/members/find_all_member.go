package members

// FindAllMember Find All Member
// @Summary  Find All Member
// @Description API For Find All Member
// @ID get-all-member
// @Accept json
// @Produce json
// @Tags MEMBERS
// @Param Authorization header string true "for authentication"
// @Param year query string false "year of EVAT eMC"
// @Param page query string false "Offset for search teams"
// @Param pageSize query string false "PageSize of teams"
// @Success 200 {object} dtos.FindAllMemberResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/members [get]
func FindAllMember() {}
