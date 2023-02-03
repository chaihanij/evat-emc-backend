package teams

// DeleteTeam Delete Team
// @Summary Delete Team
// @Description API For Delete Team
// @ID delete-team
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param uuid path string true "uuid of teams"
// @Success 200 {object} dtos.DeleteTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams/:uuid [delete]
func DeleteTeam() {}
