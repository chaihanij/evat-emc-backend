package assignments

// DeleteAssignment Delete Assignment
// @Summary Delete Assignment
// @Description API For Delete Assignment
// @ID delete-assignment
// @Accept json
// @Produce json
// @Tags ASSIGNMENTS
// @Param Authorization header string true "for authentication"
// @Param assignment_uuid path string true "assignment_uuid of assignments"
// @Success 200 {object} dtos.DeleteAssignmentResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/assignments/:assignment_uuid [delete]
func DeleteAssignment() {}
