package teams

// RegisterTeam Register Team
// @Summary Register Team
// @Description API For Register Team
// @ID post-egister-team
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param body body dtos.RegisterTeamRequestJSON true "All params related to teams"
// @Success 200 {object} dtos.RegisterTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/register-teams [post]
func RegisterTeam() {}
