package teams

// FindOneTeam Find One Team
// @Summary  Find One Team
// @Description API For Find One Team
// @ID get-one-team
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param uuid path string true "uuid of team"
// @Success 200 {object} dtos.FindOneTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams/:uid [get]
func FindOneTeam() {}
