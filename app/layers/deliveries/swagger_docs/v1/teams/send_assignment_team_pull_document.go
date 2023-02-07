package teams

// SendAssignmentTeamPullDocument Send Assignment Team Pull Document
// @Summary  Send Assignment Team Pull Document
// @Description API For Send Assignment Team
// @ID delete-team-assignment-teams-documents
// @Accept json
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param team_uuid path string true "team_uuid of teams"
// @Param assignment_uuid path string true "uuid of assignments"
// @Param document_uuid path string true "document_uuid of documents"
// @Success 200 {object} dtos.SendAssignmentTeamPushDocumentJSONJSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams/:team_uuid/assignemnts/:assignemnt_uuid/documents/:document_uuid [delete]
func SendAssignmentTeamPullDocumentRequestJSON() {}
