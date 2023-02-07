package teams

// SendAssignmentTeam Send Assignment Team
// @Summary Send Assignment Team
// @Description API For Send Assignment Team
// @ID post-team-assignment-teams
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param team_uuid path string true "team_uuid of teams"
// @Param assignment_uuid path string team_uuid "uuid of assignments"
// @Param body body dtos.SendAssignmentTeamRequestJSON true "All params related to teams"
// @Success 200 {object} dtos.SendAssignmentTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams/:team_uuid/assignemnts/:assignemnt_uuid [post]
func SendAssignmentTeam() {}
