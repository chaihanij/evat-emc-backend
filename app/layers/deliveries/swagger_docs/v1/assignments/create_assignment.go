package assignments

// CreateAssignments Create Assignments
// @Summary Create Assignments
// @Description API For Create Assignments
// @ID post-assignment
// @Accept json
// @Produce json
// @Tags ASSIGNMENTS
// @Param Authorization header string true "for authentication"
// @Param uuid path string true "uuid of member"
// @Param body body dtos.CreateAssignmentRequestJSON true "All params related to assignments"
// @Success 200 {object} dtos.CreateAssignmentResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/assignments [post]
func CreateAssignment() {}
