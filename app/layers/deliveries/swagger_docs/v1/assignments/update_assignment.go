package assignments

// UpdateAssignment Update Assignment
// @Summary Update Assignment
// @Description API For Update Assignment
// @ID put-assignments
// @Accept json
// @Produce json
// @Tags ASSIGNMENTS
// @Param Authorization header string true "for authentication"
// @Param assignment_uuid path string true "assignment_uuid of assignments"
// @Param body body dtos.UpdateAssignmentRequestJSON true "All params related to assignment"
// @Success 200 {object} dtos.UpdateAssignmentResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/assignment/:assignment_uuid [put]
func UpdateAssignment() {}
