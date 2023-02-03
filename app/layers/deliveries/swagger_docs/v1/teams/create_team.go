package teams

// CreateTeam Create Team
// @Summary Create Team
// @Description API For Create Team
// @ID post-team
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param body body dtos.CreateTeamRequestJSON true "All params related to teams"
// @Success 200 {object} dtos.CreateTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams [post]
func CreateTeam() {}
