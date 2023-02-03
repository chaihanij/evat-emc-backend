package assignments

// FindOneAssignment Find One Assignment
// @Summary Find One Assignment
// @Description API For Find One Assignment
// @ID get-one-assignment
// @Accept json
// @Produce json
// @Tags ASSIGNMENTS
// @Param Authorization header string true "for authentication"
// @Param uuid path string true "uuid of assignments"
// @Success 200 {object} dtos.FindOneAssignmentResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/assignments [get]
func FindOneAssignment() {}
