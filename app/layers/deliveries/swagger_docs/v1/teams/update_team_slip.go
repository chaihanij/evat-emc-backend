package teams

// UpdateTeamSlip Update Team
// @Summary UpdateTeamSlip
// @Description API For UpdateTeamSlip
// @ID post-team-slip
// @Accept mpfd
// @Produce json
// @Tags TEAMS
// @Param Authorization header string true "for authentication"
// @Param uuid path string true "uuid of teams"
// @Param slip formData file true "image slip"
// @Success 200 {object} dtos.FindOneTeamResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/teams/:uuid/slip [POST]
func UpdateTeamSlip() {}
