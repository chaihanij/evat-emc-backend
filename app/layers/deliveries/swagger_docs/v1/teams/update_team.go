package teams

// UpdateTeam Update Team
// @Summary Update Team
// @Description API For Update Team
// @ID put-team
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param uuid path string true "uuid of teams"
// @Param body body dtos.UpdateTeamRequestJSON true "All params related to teams"
// @Success 200 {object} dtos.UpdateTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams/:uuid [put]
func UpdateTeam() {}
