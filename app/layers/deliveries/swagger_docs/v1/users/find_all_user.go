package users

// FindAllUser Find All User
// @Summary  Find All User
// @Description API For Find All User
// @ID get-all-user
// @Accept json
// @Produce json
// @Tags USERS
// @Param Authorization header string true "for authentication"
// @Param year query string false "year of EVAT eMC"
// @Param page query string false "Offset for search users"
// @Param pageSize query string false "PageSize of users"
// @Success 200 {object} dtos.FindAllUserResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/users [get]
func FindAllUser() {}
