package teams

// SendAssignmentTeamPushDocument SendAssignmentTeam Push Document
// @Summary Send AssignmentTeam Push Document
// @Description API For Send Assignment Team Push Document
// @ID post-team-assignment-teams-documents
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param team_uuid path string true "team_uuid of teams"
// @Param assignment_uuid path string true "uuid of assignments"
// @Param document formData file true "file document upload"
// @Success 200 {object} dtos.SendAssignmentTeamPushDocumentJSONJSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams/:team_uuid/assignemnts/:assignemnt_uuid/documents [post]
func SendAssignmentTeamPushDocument() {}
