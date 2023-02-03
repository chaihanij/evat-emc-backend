package assignments

// FindAllAssignment Find All Assignment
// @Summary  Find All Assignment
// @Description API For Find All Assignment
// @ID get-all-assignment
// @Accept json
// @Produce json
// @Tags ASSIGNMENTS
// @Param Authorization header string true "for authentication"
// @Param year query string false "year of EVAT eMC"
// @Param page query string false "Offset for search assignments"
// @Param pageSize query string false "PageSize of assignments"
// @Success 200 {object} dtos.FindAllAssignmentResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/assignments [get]
func FindAllAssignment() {}
