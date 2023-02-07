package teams

// FindOneAssignmentTeam Find One AssignmentTeam
// @Summary Find One AssignmentTeam
// @Description API For Find One AssignmentTeam
// @ID get-team-assignment-teams
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param team_uuid path string true "team_uuid of teams"
// @Param assignment_uuid path string team_uuid "uuid of assignments"
// @Success 200 {object} dtos.FindOneAssignmentTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams/:team_uuid/assignemnts/:assignemnt_uuid [get]
func FindOneAssignmentTeam() {}
