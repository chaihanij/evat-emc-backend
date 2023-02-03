package teams

// FindAllUser Find All Team
// @Summary  Find All Team
// @Description API For Find All Team
// @ID get-all-teams
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param year query string false "year of EVAT eMC"
// @Param page query string false "Offset for search teams"
// @Param pageSize query string false "PageSize of teams"
// @Success 200 {object} dtos.FindAllTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams [get]
func FindAllTeam() {}
