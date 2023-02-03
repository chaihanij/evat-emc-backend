package users

// FindOneUser Find One User
// @Summary  Find One User
// @Description API For Find One User
// @ID get-one-user
// @Accept json
// @Produce json
// @Tags USERS
// @Param Authorization header string true "for authentication"
// @Param uid path string true "uid of user"
// @Success 200 {object} dtos.FindOneUserResponseSwagger
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /v1/users/:uid [get]
func FindOneUser() {}
